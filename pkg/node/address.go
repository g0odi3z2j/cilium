// Copyright 2016-2020 Authors of Cilium
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
	"github.com/cilium/cilium/pkg/byteorder"
	"github.com/cilium/cilium/pkg/cidr"
	"github.com/cilium/cilium/pkg/common"
	"github.com/cilium/cilium/pkg/defaults"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"

	"github.com/sirupsen/logrus"
)

var (
	ipv4Loopback        net.IP
	ipv4ExternalAddress net.IP
	ipv4InternalAddress net.IP
	ipv4NodePortAddrs   map[string]net.IP // iface name => ip addr
	ipv6Address         net.IP
	ipv6RouterAddress   net.IP
	ipv6NodePortAddrs   map[string]net.IP // iface name => ip addr
	ipv4AllocRange      *cidr.CIDR
	ipv6AllocRange      *cidr.CIDR

	// k8s Node IP (either InternalIP or ExternalIP or nil; the former is preferred)
	k8sNodeIP net.IP

	ipsecKeyIdentity uint8
)

func makeIPv6HostIP() net.IP {
	ipstr := "fc00::10CA:1"
	ip := net.ParseIP(ipstr)
	if ip == nil {
		log.WithField(logfields.IPAddr, ipstr).Fatal("Unable to parse IP")
	}

	return ip
}

// InitDefaultPrefix initializes the node address and allocation prefixes with
// default values derived from the system. device can be set to the primary
// network device of the system in which case the first address with global
// scope will be regarded as the system's node address.
func InitDefaultPrefix(device string) {
	if option.Config.EnableIPv4 {
		ip, err := firstGlobalV4Addr(device, GetInternalIPv4())
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
			v4range := fmt.Sprintf(defaults.DefaultIPv4Prefix+"/%d",
				ip.To4()[3], defaults.DefaultIPv4PrefixLen)
			_, ip4net, err := net.ParseCIDR(v4range)
			if err != nil {
				log.WithError(err).WithField(logfields.V4Prefix, v4range).Panic("BUG: Invalid default IPv4 prefix")
			}

			ipv4AllocRange = cidr.NewCIDR(ip4net)
			log.WithField(logfields.V4Prefix, ipv4AllocRange).Info("Using autogenerated IPv4 allocation range")
		}
	}

	if option.Config.EnableIPv6 {
		if ipv6Address == nil {
			// Find a IPv6 node address first
			ipv6Address, _ = firstGlobalV6Addr(device, GetIPv6Router())
			if ipv6Address == nil {
				ipv6Address = makeIPv6HostIP()
			}
		}

		if ipv6AllocRange == nil && ipv4AllocRange != nil {
			// The IPv6 allocation should be derived from the IPv4 allocation.
			ip := ipv4AllocRange.IP
			v6range := fmt.Sprintf("%s%02x%02x:%02x%02x:0:0/%d",
				option.Config.IPv6ClusterAllocCIDRBase, ip[0], ip[1], ip[2], ip[3], 96)

			_, ip6net, err := net.ParseCIDR(v6range)
			if err != nil {
				log.WithError(err).WithField(logfields.V6Prefix, v6range).Panic("BUG: Invalid default IPv6 prefix")
			}

			ipv6AllocRange = cidr.NewCIDR(ip6net)
			log.WithField(logfields.V6Prefix, ipv6AllocRange).Info("Using autogenerated IPv6 allocation range")
		}
	}
}

// InitNodePortAddrs initializes NodePort IPv{4,6} addrs from the given devices.
func InitNodePortAddrs(devices []string) error {
	if option.Config.EnableIPv4 {
		ipv4NodePortAddrs = make(map[string]net.IP, len(devices))
		for _, device := range devices {
			ip, err := firstGlobalV4Addr(device, GetK8sNodeIP())
			if err != nil {
				return fmt.Errorf("Failed to determine IPv4 of %s for NodePort", device)
			}

			ipv4NodePortAddrs[device] = ip
		}
	}

	if option.Config.EnableIPv6 {
		ipv6NodePortAddrs = make(map[string]net.IP, len(devices))
		for _, device := range devices {
			ip, err := firstGlobalV6Addr(device, GetK8sNodeIP())
			if err != nil {
				return fmt.Errorf("Failed to determine IPv6 of %s for NodePort", device)
			}
			ipv6NodePortAddrs[device] = ip
		}
	}

	return nil
}

// GetIPv4Loopback returns the loopback IPv4 address of this node.
func GetIPv4Loopback() net.IP {
	return ipv4Loopback
}

// SetIPv4Loopback sets the loopback IPv4 address of this node.
func SetIPv4Loopback(ip net.IP) {
	ipv4Loopback = ip
}

// GetIPv4AllocRange returns the IPv4 allocation prefix of this node
func GetIPv4AllocRange() *cidr.CIDR {
	return ipv4AllocRange
}

// GetIPv6AllocRange returns the IPv6 allocation prefix of this node
func GetIPv6AllocRange() *cidr.CIDR {
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
func SetIPv4AllocRange(net *cidr.CIDR) {
	ipv4AllocRange = net
}

// Uninitialize resets this package to the default state, for use in
// testsuite code.
func Uninitialize() {
	ipv4AllocRange = nil
	ipv6AllocRange = nil
}

// GetNodePortIPv4 returns the node-port IPv4 address for NAT
func GetNodePortIPv4Addrs() []net.IP {
	addrs := make([]net.IP, 0, len(ipv4NodePortAddrs))
	for _, addr := range ipv4NodePortAddrs {
		addrs = append(addrs, addr)
	}
	return addrs
}

// GetNodePortIPv4AddrsWithDevices returns the map iface => NodePort IPv4.
func GetNodePortIPv4AddrsWithDevices() map[string]net.IP {
	out := make(map[string]net.IP, len(ipv4NodePortAddrs))
	for iface, ip := range ipv4NodePortAddrs {
		dup := make(net.IP, len(ip))
		copy(dup, ip)
		out[iface] = dup
	}
	return out
}

// GetNodePortIPv6 returns the node-port IPv6 address for NAT
func GetNodePortIPv6Addrs() []net.IP {
	addrs := make([]net.IP, 0, len(ipv6NodePortAddrs))
	for _, addr := range ipv6NodePortAddrs {
		addrs = append(addrs, addr)
	}
	return addrs
}

// GetNodePortIPv4AddrsWithDevices returns the map iface => NodePort IPv6.
func GetNodePortIPv6AddrsWithDevices() map[string]net.IP {
	out := make(map[string]net.IP, len(ipv6NodePortAddrs))
	for iface, ip := range ipv6NodePortAddrs {
		dup := make(net.IP, len(ip))
		copy(dup, ip)
		out[iface] = dup
	}
	return out
}

// SetIPv6NodeRange sets the IPv6 address pool to be used on this node
func SetIPv6NodeRange(net *net.IPNet) {
	copy := *net
	ipv6AllocRange = cidr.NewCIDR(&copy)
}

// AutoComplete completes the parts of addressing that can be auto derived
func AutoComplete() error {
	if option.Config.EnableHostIPRestore {
		// Read the previous cilium host IPs from node_config.h for backward
		// compatibility
		ipv4GW, ipv6Router := getCiliumHostIPs()

		if ipv4GW != nil && option.Config.EnableIPv4 {
			SetInternalIPv4(ipv4GW)
		}

		if ipv6Router != nil && option.Config.EnableIPv6 {
			SetIPv6Router(ipv6Router)
		}
	}

	InitDefaultPrefix(option.Config.DirectRoutingDevice)

	if option.Config.EnableIPv6 && ipv6AllocRange == nil {
		return fmt.Errorf("IPv6 allocation CIDR is not configured. Please specificy --ipv6-range")
	}

	if option.Config.EnableIPv4 && ipv4AllocRange == nil {
		return fmt.Errorf("IPv4 allocation CIDR is not configured. Please specificy --ipv4-range")
	}

	return nil
}

// ValidatePostInit validates the entire addressing setup and completes it as
// required
func ValidatePostInit() error {
	if option.Config.EnableIPv4 || option.Config.Tunnel != option.TunnelDisabled {
		if ipv4ExternalAddress == nil {
			return fmt.Errorf("external IPv4 node address could not be derived, please configure via --ipv4-node")
		}
	}

	if option.Config.EnableIPv4 && ipv4InternalAddress == nil {
		return fmt.Errorf("BUG: Internal IPv4 node address was not configured")
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

// IsHostIPv4 returns true if the IP specified is a host IP
func IsHostIPv4(ip net.IP) bool {
	return ip.Equal(GetInternalIPv4()) || ip.Equal(GetExternalIPv4())
}

// IsHostIPv6 returns true if the IP specified is a host IP
func IsHostIPv6(ip net.IP) bool {
	return ip.Equal(GetIPv6()) || ip.Equal(GetIPv6Router())
}

// GetNodeAddressing returns the NodeAddressing model for the local IPs.
func GetNodeAddressing() *models.NodeAddressing {
	a := &models.NodeAddressing{}

	if option.Config.EnableIPv6 {
		a.IPV6 = &models.NodeAddressingElement{
			Enabled:    option.Config.EnableIPv6,
			IP:         GetIPv6Router().String(),
			AllocRange: GetIPv6AllocRange().String(),
		}
	}

	if option.Config.EnableIPv4 {
		a.IPV4 = &models.NodeAddressingElement{
			Enabled:    option.Config.EnableIPv4,
			IP:         GetInternalIPv4().String(),
			AllocRange: GetIPv4AllocRange().String(),
		}
	}

	return a
}

func getCiliumHostIPsFromFile(nodeConfig string) (ipv4GW, ipv6Router net.IP) {
	var hasIPv4, hasIPv6 bool
	f, err := os.Open(nodeConfig)
	switch {
	case err != nil:
	default:
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			txt := scanner.Text()
			switch {
			case !hasIPv6 && strings.Contains(txt, defaults.RestoreV6Addr):
				defineLine := strings.Split(txt, defaults.RestoreV6Addr)
				if len(defineLine) != 2 {
					continue
				}
				ipv6 := common.C2GoArray(defineLine[1])
				if len(ipv6) != net.IPv6len {
					continue
				}
				ipv6Router = net.IP(ipv6)
				hasIPv6 = true
			case !hasIPv4 && strings.Contains(txt, defaults.RestoreV4Addr):
				defineLine := strings.Split(txt, defaults.RestoreV4Addr)
				if len(defineLine) != 2 {
					continue
				}
				ipv4 := common.C2GoArray(defineLine[1])
				if len(ipv4) != net.IPv6len {
					continue
				}
				ipv4GW = net.IP(ipv4)
				hasIPv4 = true

			// Legacy cases based on the header defines:
			case !hasIPv4 && strings.Contains(txt, "IPV4_GATEWAY"):
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
				if ipv4GWint64 != int64(0) {
					bs := make([]byte, net.IPv4len)
					byteorder.NetworkToHostPut(bs, uint32(ipv4GWint64))
					ipv4GW = net.IPv4(bs[0], bs[1], bs[2], bs[3])
					hasIPv4 = true
				}
			case !hasIPv6 && strings.Contains(txt, " ROUTER_IP "):
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
				hasIPv6 = true
			}
		}
	}
	return ipv4GW, ipv6Router
}

// getCiliumHostIPs returns the Cilium IPv4 gateway and router IPv6 address from
// the node_config.h file if is present; or by deriving it from
// defaults.HostDevice interface, on which only the IPv4 is possible to derive.
func getCiliumHostIPs() (ipv4GW, ipv6Router net.IP) {
	nodeConfig := option.Config.GetNodeConfigPath()
	ipv4GW, ipv6Router = getCiliumHostIPsFromFile(nodeConfig)
	if ipv4GW != nil || ipv6Router != nil {
		log.WithFields(logrus.Fields{
			"ipv4": ipv4GW,
			"ipv6": ipv6Router,
			"file": nodeConfig,
		}).Info("Restored router address from node_config")
		return ipv4GW, ipv6Router
	}
	return getCiliumHostIPsFromNetDev(option.Config.HostDevice)
}

// SetIPsecKeyIdentity sets the IPsec key identity an opaque value used to
// identity encryption keys used on the node.
func SetIPsecKeyIdentity(id uint8) {
	ipsecKeyIdentity = id
}

// GetIPsecKeyIdentity returns the IPsec key identity of the node
func GetIPsecKeyIdentity() uint8 {
	return ipsecKeyIdentity
}

// GetK8sNodeIPs returns k8s Node IP addr.
func GetK8sNodeIP() net.IP {
	return k8sNodeIP
}

// SetK8sNodeIP sets k8s Node IP addr.
func SetK8sNodeIP(ip net.IP) {
	k8sNodeIP = ip
}
