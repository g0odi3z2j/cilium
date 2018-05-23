// Copyright 2016-2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package node

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/common"
	"github.com/cilium/cilium/pkg/byteorder"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

var (
	ipv4ClusterCidrMaskSize = DefaultIPv4ClusterPrefixLen

	ipv4ExternalAddress net.IP
	ipv4InternalAddress net.IP
	ipv6Address         net.IP
	ipv6RouterAddress   net.IP
	ipv4AllocRange      *net.IPNet
	ipv6AllocRange      *net.IPNet
	ipv4HealthAddress   net.IP
	ipv6HealthAddress   net.IP
)

func makeIPv6HostIP() net.IP {
	ipstr := "fc00::10CA:1"
	ip := net.ParseIP(ipstr)
	if ip == nil {
		log.WithField(logfields.IPAddr, ipstr).Fatal("Unable to parse IP")
	}

	return ip
}

func firstGlobalV4Addr(intf string) (net.IP, error) {
	var link netlink.Link
	var err error

	if intf != "" && intf != "undefined" {
		link, err = netlink.LinkByName(intf)
		if err != nil {
			return firstGlobalV4Addr("")
		}
	}

	addr, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return nil, err
	}

	for _, a := range addr {
		if a.Scope == unix.RT_SCOPE_UNIVERSE {
			if len(a.IP) >= 4 {
				return a.IP, nil
			}
		}
	}

	return nil, fmt.Errorf("No address found")
}

func findIPv6NodeAddr() net.IP {
	addr, err := netlink.AddrList(nil, netlink.FAMILY_V6)
	if err != nil {
		return nil
	}

	// prefer global scope address
	for _, a := range addr {
		if a.Scope == unix.RT_SCOPE_UNIVERSE {
			if len(a.IP) >= 16 {
				return a.IP
			}
		}
	}

	// fall back to anything wider than link (site, custom, ...)
	for _, a := range addr {
		if a.Scope < unix.RT_SCOPE_LINK {
			if len(a.IP) >= 16 {
				return a.IP
			}
		}
	}

	return nil
}

// SetIPv4ClusterCidrMaskSize sets the size of the mask of the IPv4 cluster prefix
func SetIPv4ClusterCidrMaskSize(size int) {
	ipv4ClusterCidrMaskSize = size
}

// InitDefaultPrefix initializes the node address and allocation prefixes with
// default values derived from the system. device can be set to the primary
// network device of the system in which case the first address with global
// scope will be regarded as the system's node address.
func InitDefaultPrefix(device string) {
	if ipv6Address == nil {
		// Find a IPv6 node address first
		ipv6Address = findIPv6NodeAddr()
		if ipv6Address == nil {
			ipv6Address = makeIPv6HostIP()
		}
	}

	ip, err := firstGlobalV4Addr(device)
	if err != nil {
		return
	}

	if ipv4ExternalAddress == nil {
		ipv4ExternalAddress = ip
	}

	if ipv4AllocRange == nil {
		// If the IPv6AllocRange is not nil then the IPv4 allocation should be
		// derived from the IPv6AllocRange.
		//                     vvvv vvvv
		// FD00:0000:0000:0000:0000:0000:0000:0000
		if ipv6AllocRange != nil {
			ip = net.IPv4(ipv6AllocRange.IP[8],
				ipv6AllocRange.IP[9],
				ipv6AllocRange.IP[10],
				ipv6AllocRange.IP[11])
		}
		v4range := fmt.Sprintf(DefaultIPv4Prefix+"/%d",
			ip.To4()[3], DefaultIPv4PrefixLen)
		_, ip4net, err := net.ParseCIDR(v4range)
		if err != nil {
			log.WithError(err).WithField(logfields.V4Prefix, v4range).Panic("BUG: Invalid default IPv4 prefix")
		}

		ipv4AllocRange = ip4net
		log.WithField(logfields.V4Prefix, ipv4AllocRange).Info("Using autogenerated IPv4 allocation range")
	}

	if ipv6AllocRange == nil {
		// The IPv6 allocation should be derived from the IPv4 allocation.
		ip = ipv4AllocRange.IP
		v6range := fmt.Sprintf("%s%02x%02x:%02x%02x:0:0/%d",
			DefaultIPv6Prefix, ip[0], ip[1], ip[2], ip[3],
			IPv6NodePrefixLen)

		_, ip6net, err := net.ParseCIDR(v6range)
		if err != nil {
			log.WithError(err).WithField(logfields.V6Prefix, v6range).Panic("BUG: Invalid default IPv6 prefix")
		}

		ipv6AllocRange = ip6net
		log.WithField(logfields.V6Prefix, ipv6AllocRange).Info("Using autogenerated IPv6 allocation range")
	}
}

// GetIPv4ClusterRange returns the IPv4 prefix of the cluster
func GetIPv4ClusterRange() *net.IPNet {
	mask := net.CIDRMask(ipv4ClusterCidrMaskSize, 32)
	return &net.IPNet{
		IP:   ipv4AllocRange.IP.Mask(mask),
		Mask: mask,
	}
}

// GetIPv4AllocRange returns the IPv4 allocation prefix of this node
func GetIPv4AllocRange() *net.IPNet {
	return ipv4AllocRange
}

// GetIPv6ClusterRange returns the IPv6 prefix of the clustr
func GetIPv6ClusterRange() *net.IPNet {
	mask := net.CIDRMask(DefaultIPv6ClusterPrefixLen, 128)
	return &net.IPNet{
		IP:   ipv6AllocRange.IP.Mask(mask),
		Mask: mask,
	}
}

// GetIPv6AllocRange returns the IPv6 allocation prefix of this node
func GetIPv6AllocRange() *net.IPNet {
	mask := net.CIDRMask(IPv6NodeAllocPrefixLen, 128)
	return &net.IPNet{
		IP:   ipv6AllocRange.IP.Mask(mask),
		Mask: mask,
	}
}

// GetIPv6NodeRange returns the IPv6 allocation prefix of this node
func GetIPv6NodeRange() *net.IPNet {
	return ipv6AllocRange
}

// SetExternalIPv4 sets the external IPv4 node address. It must be reachable on the network.
func SetExternalIPv4(ip net.IP) {
	ipv4ExternalAddress = ip
}

// GetExternalIPv4 returns the external IPv4 node address
func GetExternalIPv4() net.IP {
	return ipv4ExternalAddress
}

// SetInternalIPv4 sets the internal IPv4 node address, it is allocated from the node prefix
func SetInternalIPv4(ip net.IP) {
	ipv4InternalAddress = ip
}

// GetInternalIPv4 returns the internal IPv4 node address
func GetInternalIPv4() net.IP {
	return ipv4InternalAddress
}

// GetHostMasqueradeIPv4 returns the IPv4 address to be used for masquerading
// any traffic that is being forwarded from the host into the Cilium cluster.
func GetHostMasqueradeIPv4() net.IP {
	return ipv4InternalAddress
}

// SetIPv4AllocRange sets the IPv4 address pool to use when allocating
// addresses for local endpoints
func SetIPv4AllocRange(net *net.IPNet) {
	ipv4AllocRange = net
}

// Uninitialize resets this package to the default state, for use in
// testsuite code.
func Uninitialize() {
	ipv4AllocRange = nil
	ipv6AllocRange = nil
}

// SetIPv6NodeRange sets the IPv6 address pool to be used on this node
func SetIPv6NodeRange(net *net.IPNet) error {
	if ones, _ := net.Mask.Size(); ones != IPv6NodePrefixLen {
		return fmt.Errorf("prefix length must be /%d", IPv6NodePrefixLen)
	}

	copy := *net
	ipv6AllocRange = &copy

	return nil
}

// AutoComplete completes the parts of addressing that can be auto derived
func AutoComplete() error {
	// Read the previous cilium host IPs from node_config.h for backward
	// compatibility
	ipv4GW, ipv6Router := getCiliumHostIPs()
	SetInternalIPv4(ipv4GW)
	SetIPv6Router(ipv6Router)

	if option.Config.Device == "undefined" {
		InitDefaultPrefix("")
	}

	if ipv6AllocRange == nil {
		return fmt.Errorf("IPv6 per node allocation prefix is not configured. Please specificy --ipv6-range")
	}

	return nil
}

// ValidatePostInit validates the entire addressing setup and completes it as
// required
func ValidatePostInit() error {
	if ipv4ExternalAddress == nil {
		return fmt.Errorf("External IPv4 node address could not be derived, please configure via --ipv4-node")
	}

	if !option.Config.IPv4Disabled {
		if ipv4InternalAddress == nil {
			return fmt.Errorf("BUG: Internal IPv4 node address was not configured")
		}

		if !ipv4AllocRange.Contains(ipv4InternalAddress) {
			return fmt.Errorf("BUG: Internal IPv4 (%s) must be part of cluster prefix (%s)",
				ipv4InternalAddress, ipv4AllocRange)
		}

		ones, _ := ipv4AllocRange.Mask.Size()
		if ipv4ClusterCidrMaskSize > ones {
			return fmt.Errorf("IPv4 per node allocation prefix (%s) must be inside cluster prefix (%s)",
				ipv4AllocRange, GetIPv4ClusterRange())
		}
	}

	return nil
}

// SetIPv6 sets the IPv6 address of the node
func SetIPv6(ip net.IP) {
	ipv6Address = ip
}

// GetIPv6 returns the IPv6 address of the node
func GetIPv6() net.IP {
	return ipv6Address
}

// GetIPv6Router returns the IPv6 address of the node
func GetIPv6Router() net.IP {
	return ipv6RouterAddress
}

// SetIPv6Router returns the IPv6 address of the node
func SetIPv6Router(ip net.IP) {
	ipv6RouterAddress = ip
}

// GetIPv4HealthIP returns the IPv4 address of the local cilium-health endpoint
func GetIPv4HealthIP() net.IP {
	return ipv4HealthAddress
}

// SetIPv4HealthIP sets the IPv4 address of the local cilium-health endpoint
func SetIPv4HealthIP(ip net.IP) {
	ipv4HealthAddress = ip
}

// GetIPv6HealthIP returns the IPv6 address of the local cilium-health endpoint
func GetIPv6HealthIP() net.IP {
	return ipv6HealthAddress
}

// SetIPv6HealthIP sets the IPv6 address of the local cilium-health endpoint
func SetIPv6HealthIP(ip net.IP) {
	ipv6HealthAddress = ip
}

// GetIPv6NodeRoute returns a route pointing to the IPv6 node address
func GetIPv6NodeRoute() net.IPNet {
	return net.IPNet{
		IP:   ipv6RouterAddress,
		Mask: net.CIDRMask(128, 128),
	}
}

// GetIPv4NodeRoute returns a route pointing to the IPv4 node address
func GetIPv4NodeRoute() net.IPNet {
	return net.IPNet{
		IP:   ipv4InternalAddress,
		Mask: net.CIDRMask(32, 32),
	}
}

// UseNodeCIDR sets the ipv4-range and ipv6-range values values from the
// addresses defined in the given node.
func UseNodeCIDR(node *Node) error {
	scopedLog := log.WithField(logfields.Node, node.Name)
	if node.IPv4AllocCIDR != nil {
		scopedLog.WithField(logfields.V4Prefix, node.IPv4AllocCIDR).Info("Retrieved IPv4 allocation range for node. Using it for ipv4-range")
		SetIPv4AllocRange(node.IPv4AllocCIDR)
	}
	if node.IPv6AllocCIDR != nil {
		scopedLog.WithField(logfields.V6Prefix, node.IPv6AllocCIDR).Info("Retrieved IPv6 allocation range for node. Using it for ipv6-range")
		if err := SetIPv6NodeRange(node.IPv6AllocCIDR); err != nil {
			scopedLog.WithError(err).WithField(logfields.V4Prefix, node.IPv6AllocCIDR).Warn("k8s: Can't use IPv6 CIDR range from k8s")
		}
	}

	return nil
}

// UseNodeAddresses sets the local ipv4-node and ipv6-node values from the
// addresses defined in the given node.
func UseNodeAddresses(node *Node) error {
	scopedLog := log.WithField(logfields.Node, node.Name)
	nodeIP4 := node.GetNodeIP(false)
	if nodeIP4 != nil {
		scopedLog.WithField(logfields.IPAddr, nodeIP4).Info("Automatically retrieved IP for node. Using it for ipv4-node")
		SetExternalIPv4(nodeIP4)
	}
	nodeIP6 := node.GetNodeIP(true)
	if nodeIP6 != nil {
		scopedLog.WithField(logfields.IPAddr, nodeIP6).Info("Automatically retrieved IP for node. Using it for ipv6-node")
		SetIPv6(nodeIP6)
	}

	return nil
}

// IsHostIPv4 returns true if the IP specified is a host IP
func IsHostIPv4(ip net.IP) bool {
	return ip.Equal(GetInternalIPv4()) || ip.Equal(GetExternalIPv4())
}

// IsHostIPv6 returns true if the IP specified is a host IP
func IsHostIPv6(ip net.IP) bool {
	return ip.Equal(GetIPv6()) || ip.Equal(GetIPv6Router())
}

// GetNodeAddressing returns the NodeAddressing model for the local IPs.
func GetNodeAddressing(enableIPv4 bool) *models.NodeAddressing {
	return &models.NodeAddressing{
		IPV6: &models.NodeAddressingElement{
			Enabled:    true,
			IP:         GetIPv6Router().String(),
			AllocRange: GetIPv6AllocRange().String(),
		},
		IPV4: &models.NodeAddressingElement{
			Enabled:    enableIPv4,
			IP:         GetInternalIPv4().String(),
			AllocRange: GetIPv4AllocRange().String(),
		},
	}
}

func getCiliumHostIPsFromFile(nodeConfig string) (ipv4GW, ipv6Router net.IP) {
	f, err := os.Open(nodeConfig)
	switch {
	case err != nil:
	default:
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			txt := scanner.Text()
			switch {
			case strings.Contains(txt, "IPV4_GATEWAY"):
				// #define IPV4_GATEWAY 0xee1c000a
				defineLine := strings.Split(txt, " ")
				if len(defineLine) != 3 {
					continue
				}
				ipv4GWHex := strings.TrimPrefix(defineLine[2], "0x")
				ipv4GWint64, err := strconv.ParseInt(ipv4GWHex, 16, 0)
				if err != nil {
					continue
				}
				bs := make([]byte, net.IPv4len)
				byteorder.NetworkToHostPut(bs, uint32(ipv4GWint64))
				ipv4GW = net.IPv4(bs[0], bs[1], bs[2], bs[3])
			case strings.Contains(txt, " ROUTER_IP "):
				// #define ROUTER_IP 0xf0, 0xd, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8a, 0xd6
				defineLine := strings.Split(txt, " ROUTER_IP ")
				if len(defineLine) != 2 {
					continue
				}
				ipv6 := common.C2GoArray(defineLine[1])
				if len(ipv6) != net.IPv6len {
					continue
				}
				ipv6Router = net.IP(ipv6)
			}
		}
	}
	return ipv4GW, ipv6Router
}

// getCiliumHostIPsFromNetDev returns the first IPv4 link local and returns
// it
func getCiliumHostIPsFromNetDev(devName string) (ipv4GW, ipv6Router net.IP) {
	hostDev, err := netlink.LinkByName(devName)
	if err != nil {
		return nil, nil
	}
	addrs, err := netlink.AddrList(hostDev, netlink.FAMILY_ALL)
	if err != nil {
		return nil, nil
	}
	for _, addr := range addrs {
		if addr.IP.To4() != nil {
			if addr.Scope == int(netlink.SCOPE_LINK) {
				ipv4GW = addr.IP
			}
		} else {
			if addr.Scope != int(netlink.SCOPE_LINK) {
				ipv6Router = addr.IP
			}
		}
	}
	return nil, nil
}

// getCiliumHostIPs returns the Cilium IPv4 gateway and router IPv6 address from
// the node_config.h file if is present; or by deriving it from cilium_host
// interface, on which only the IPv4 is possible to derive.
func getCiliumHostIPs() (ipv4GW, ipv6Router net.IP) {
	nodeConfig := option.Config.GetNodeConfigPath()
	ipv4GW, ipv6Router = getCiliumHostIPsFromFile(nodeConfig)
	if ipv4GW != nil {
		return ipv4GW, ipv6Router
	}
	return getCiliumHostIPsFromNetDev(HostDevice)
}
