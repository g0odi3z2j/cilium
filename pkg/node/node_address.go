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
	"github.com/cilium/cilium/pkg/cidr"
	"github.com/cilium/cilium/pkg/defaults"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"
)

var (
	ipv4ClusterCidrMaskSize = defaults.DefaultIPv4ClusterPrefixLen

	ipv4Loopback        net.IP
	ipv4ExternalAddress net.IP
	ipv4InternalAddress net.IP
	ipv6Address         net.IP
	ipv6RouterAddress   net.IP
	ipv4AllocRange      *cidr.CIDR
	ipv6AllocRange      *cidr.CIDR
)

func makeIPv6HostIP() net.IP {
	ipstr := "fc00::10CA:1"
	ip := net.ParseIP(ipstr)
	if ip == nil {
		log.WithField(logfields.IPAddr, ipstr).Fatal("Unable to parse IP")
	}

	return ip
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
		v4range := fmt.Sprintf(defaults.DefaultIPv4Prefix+"/%d",
			ip.To4()[3], defaults.DefaultIPv4PrefixLen)
		_, ip4net, err := net.ParseCIDR(v4range)
		if err != nil {
			log.WithError(err).WithField(logfields.V4Prefix, v4range).Panic("BUG: Invalid default IPv4 prefix")
		}

		ipv4AllocRange = cidr.NewCIDR(ip4net)
		log.WithField(logfields.V4Prefix, ipv4AllocRange).Info("Using autogenerated IPv4 allocation range")
	}

	if ipv6AllocRange == nil {
		// The IPv6 allocation should be derived from the IPv4 allocation.
		ip = ipv4AllocRange.IP
		v6range := fmt.Sprintf("%s%02x%02x:%02x%02x:0:0/%d",
			option.Config.IPv6ClusterAllocCIDRBase, ip[0], ip[1], ip[2], ip[3],
			defaults.IPv6NodePrefixLen)

		_, ip6net, err := net.ParseCIDR(v6range)
		if err != nil {
			log.WithError(err).WithField(logfields.V6Prefix, v6range).Panic("BUG: Invalid default IPv6 prefix")
		}

		ipv6AllocRange = cidr.NewCIDR(ip6net)
		log.WithField(logfields.V6Prefix, ipv6AllocRange).Info("Using autogenerated IPv6 allocation range")
	}
}

// GetIPv4ClusterRange returns the IPv4 prefix of the cluster
func GetIPv4ClusterRange() *net.IPNet {
	mask := net.CIDRMask(ipv4ClusterCidrMaskSize, 32)
	return &net.IPNet{
		IP:   ipv4AllocRange.IPNet.IP.Mask(mask),
		Mask: mask,
	}
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
	mask := net.CIDRMask(defaults.IPv6NodeAllocPrefixLen, 128)
	return cidr.NewCIDR(&net.IPNet{
		IP:   ipv6AllocRange.IPNet.IP.Mask(mask),
		Mask: mask,
	})
}

// GetIPv6NodeRange returns the IPv6 allocation prefix of this node
func GetIPv6NodeRange() *cidr.CIDR {
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

// SetIPv6NodeRange sets the IPv6 address pool to be used on this node
func SetIPv6NodeRange(net *net.IPNet) error {
	if ones, _ := net.Mask.Size(); ones != defaults.IPv6NodePrefixLen {
		return fmt.Errorf("prefix length must be /%d", defaults.IPv6NodePrefixLen)
	}

	copy := *net
	ipv6AllocRange = cidr.NewCIDR(&copy)

	return nil
}

// AutoComplete completes the parts of addressing that can be auto derived
func AutoComplete() error {
	if option.Config.EnableHostIPRestore {
		// Read the previous cilium host IPs from node_config.h for backward
		// compatibility
		ipv4GW, ipv6Router := getCiliumHostIPs()

		if ipv4GW != nil {
			log.Infof("Restored IPv4 internal node IP: %s", ipv4GW.String())
			SetInternalIPv4(ipv4GW)
		}

		if ipv6Router != nil {
			log.Infof("Restored IPv6 router IP: %s", ipv6Router.String())
			SetIPv6Router(ipv6Router)
		}
	}

	if option.Config.Device == "undefined" {
		InitDefaultPrefix("")
	}

	if option.Config.EnableIPv6 && ipv6AllocRange == nil {
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

	if option.Config.EnableIPv4 {
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
		if err := SetIPv6NodeRange(node.IPv6AllocCIDR.IPNet); err != nil {
			scopedLog.WithError(err).WithField(logfields.V6Prefix, node.IPv6AllocCIDR).Warn("k8s: Can't use IPv6 CIDR range from k8s")
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
func GetNodeAddressing() *models.NodeAddressing {
	return &models.NodeAddressing{
		IPV6: &models.NodeAddressingElement{
			Enabled:    option.Config.EnableIPv6,
			IP:         GetIPv6Router().String(),
			AllocRange: GetIPv6AllocRange().String(),
		},
		IPV4: &models.NodeAddressingElement{
			Enabled:    option.Config.EnableIPv4,
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
				if ipv4GWint64 != int64(0) {
					bs := make([]byte, net.IPv4len)
					byteorder.NetworkToHostPut(bs, uint32(ipv4GWint64))
					ipv4GW = net.IPv4(bs[0], bs[1], bs[2], bs[3])
				}
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

// getCiliumHostIPs returns the Cilium IPv4 gateway and router IPv6 address from
// the node_config.h file if is present; or by deriving it from
// defaults.HostDevice interface, on which only the IPv4 is possible to derive.
func getCiliumHostIPs() (ipv4GW, ipv6Router net.IP) {
	nodeConfig := option.Config.GetNodeConfigPath()
	ipv4GW, ipv6Router = getCiliumHostIPsFromFile(nodeConfig)
	if ipv4GW != nil {
		return ipv4GW, ipv6Router
	}
	return getCiliumHostIPsFromNetDev(option.Config.HostDevice)
}
