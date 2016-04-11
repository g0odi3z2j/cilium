package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"

	cnc "github.com/noironetworks/cilium-net/common/client"
	"github.com/noironetworks/cilium-net/common/plugins"
	ciliumtypes "github.com/noironetworks/cilium-net/common/types"

	log "github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/appc/cni/pkg/ns"
	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/appc/cni/pkg/skel"
	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/appc/cni/pkg/types"
	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/vishvananda/netlink"
)

const (
	hostInterfacePrefix      = "lxc"
	temporaryInterfacePrefix = "tmp"
)

type netConf struct {
	types.NetConf
	MTU int `json:"mtu"`
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stderr)
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func main() {
	skel.PluginMain(cmdAdd, cmdDel)
}

func loadNetConf(bytes []byte) (*netConf, error) {
	n := &netConf{}
	if err := json.Unmarshal(bytes, n); err != nil {
		return nil, fmt.Errorf("failed to load netconf: %s", err)
	}
	return n, nil
}

func removeIfFromNSIfExists(netns *os.File, ifName string) error {
	return ns.WithNetNS(netns, false, func(_ *os.File) error {
		l, err := netlink.LinkByName(ifName)
		if err != nil {
			if strings.Contains(err.Error(), "Link not found") {
				return nil
			}
			return err
		}
		return netlink.LinkDel(l)
	})
}

func renameLink(curName, newName string) error {
	link, err := netlink.LinkByName(curName)
	if err != nil {
		return err
	}

	return netlink.LinkSetName(link, newName)
}

func addIPConfigToLink(ipConfig *ciliumtypes.IPConfig, link netlink.Link, ifName string) error {
	addr := &netlink.Addr{IPNet: &ipConfig.IP}
	if err := netlink.AddrAdd(link, addr); err != nil {
		return fmt.Errorf("failed to add addr to %q: %v", ifName, err)
	}

	for _, r := range ipConfig.Routes {
		if err := netlink.RouteAdd(
			&netlink.Route{
				LinkIndex: link.Attrs().Index,
				Scope:     netlink.SCOPE_UNIVERSE,
				Dst:       &r.Destination,
				Gw:        r.NextHop,
			}); err != nil {
			if !os.IsExist(err) {
				return fmt.Errorf("failed to add route '%v via %v dev %v': %v",
					r.Destination, r.NextHop, ifName, err)
			}
		}
	}

	return nil
}

func configureIface(ifName string, config *ciliumtypes.IPAMConfig) error {
	link, err := netlink.LinkByName(ifName)
	if err != nil {
		return fmt.Errorf("failed to lookup %q: %v", ifName, err)
	}

	if err := netlink.LinkSetUp(link); err != nil {
		return fmt.Errorf("failed to set %q UP: %v", ifName, err)
	}

	if config.IP4 != nil {
		if err := addIPConfigToLink(config.IP4, link, ifName); err != nil {
			return fmt.Errorf("error configuring IPv4: %s", err.Error())
		}
	}
	if config.IP6 != nil {
		if err := addIPConfigToLink(config.IP6, link, ifName); err != nil {
			return fmt.Errorf("error configuring IPv6: %s", err.Error())
		}
	}

	return nil
}

func createCNIReply(ipamConf *ciliumtypes.IPAMConfig) error {
	routes := []types.Route{}
	for _, r := range ipamConf.IP6.Routes {
		newRoute := types.Route{
			Dst: r.Destination,
		}
		if r.NextHop != nil {
			newRoute.GW = r.NextHop
		}
		routes = append(routes, newRoute)
	}
	r := types.Result{
		IP6: &types.IPConfig{
			IP:      ipamConf.IP6.IP,
			Gateway: ipamConf.IP6.Gateway,
			Routes:  routes,
		},
	}
	return r.Print()
}

func cmdAdd(args *skel.CmdArgs) error {
	n, err := loadNetConf(args.StdinData)
	if err != nil {
		return err
	}

	c, err := cnc.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("error while starting cilium-client: %s", err)
	}

	netNsFile, err := os.Open(args.Netns)
	if err != nil {
		return fmt.Errorf("failed to open netns %q: %s", args.Netns, err)
	}
	defer netNsFile.Close()

	if err := removeIfFromNSIfExists(netNsFile, args.IfName); err != nil {
		return fmt.Errorf("failed removing interface %q from namespace %q: %s",
			args.IfName, args.Netns, err)
	}

	var ep ciliumtypes.Endpoint
	veth, peer, tmpIfName, err := plugins.SetupVeth(args.ContainerID, n.MTU, &ep)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if err = netlink.LinkDel(veth); err != nil {
				log.Warnf("failed to clean up veth %q: %s", veth.Name, err)
			}
		}
	}()

	if err = netlink.LinkSetNsFd(*peer, int(netNsFile.Fd())); err != nil {
		return fmt.Errorf("unable to move veth pair %q to netns: %s", peer, err)
	}

	err = ns.WithNetNS(netNsFile, false, func(_ *os.File) error {
		err := renameLink(tmpIfName, args.IfName)
		if err != nil {
			return fmt.Errorf("failed to rename %q to %q: %s", tmpIfName, args.IfName, err)
		}
		return nil
	})

	ipamConf, err := c.AllocateIPs(args.ContainerID)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if err = c.ReleaseIPs(args.ContainerID); err != nil {
				log.Warnf("failed to release allocated IP of container ID %q: %s", args.ContainerID, err)
			}
		}
	}()

	if err = ns.WithNetNS(netNsFile, false, func(_ *os.File) error {
		return configureIface(args.IfName, ipamConf)
	}); err != nil {
		return err
	}

	ep.LXCIP = ipamConf.IP6.IP.IP
	ep.NodeIP = ipamConf.IP6.Gateway
	ep.DockerID = args.ContainerID
	ep.SetID()
	if err = c.EndpointJoin(ep); err != nil {
		return fmt.Errorf("unable to create eBPF map: %s", err)
	}

	return createCNIReply(ipamConf)
}

func cmdDel(args *skel.CmdArgs) error {
	c, err := cnc.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("error while starting cilium-client: %s", err)
	}

	var containerIP net.IP
	// FIXME: We need to retrieve the IPv6 address somehow...
	ns.WithNetNSPath(args.Netns, false, func(hostNS *os.File) error {
		l, err := netlink.LinkByName(args.IfName)
		if err != nil {
			return err
		}
		addrs, err := netlink.AddrList(l, netlink.FAMILY_V6)
		if err != nil {
			return err
		}
		log.Debugf("IPv6 addresses found %+v\n", addrs)

		for _, addr := range addrs {
			if uint8(addr.Scope) == uint8(netlink.SCOPE_UNIVERSE) {
				containerIP = addr.IP
				break
			}
		}
		return nil
	})

	if err := c.ReleaseIPs(args.ContainerID); err != nil {
		log.Warnf("failed to release allocated IP of container ID %q: %s", args.ContainerID, err)
	}

	var ep ciliumtypes.Endpoint
	ep.LXCIP = containerIP
	ep.SetID()
	if err := c.EndpointLeave(ep.ID); err != nil {
		log.Warnf("leaving the endpoint failed: %s\n", err)
	}

	return ns.WithNetNSPath(args.Netns, false, func(hostNS *os.File) error {
		return plugins.DelLinkByName(args.IfName)
	})
}
