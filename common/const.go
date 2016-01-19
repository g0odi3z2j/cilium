package common

import "net"

const (
	PluginPath          = "/run/docker/plugins/"
	DriverSock          = PluginPath + "cilium.sock"
	CiliumPath          = "/var/run/cilium/"
	CiliumSock          = CiliumPath + "cilium.sock"
	DefaultContainerMAC = "AA:BB:CC:DD:EE:FF"
	BPFMap              = "/sys/fs/bpf/tc/globals/cilium_lxc"
)

var (
	// Default addressing schema
	//
	// cluster:		    beef:beef:beef:beef::/64
	// loadbalancer:	beef:beef:beef:beef:<lb>::/80
	// node:		    beef:beef:beef:beef:<lb>:<node>::/96
	// lxc:			    beef:beef:beef:beef:<lb>:<node>:<lxc>:<lxc>/128
	ClusterIPv6Mask      = net.CIDRMask(64, 128)
	LoadBalancerIPv6Mask = net.CIDRMask(80, 128)
	NodeIPv6Mask         = net.CIDRMask(96, 128)
	ContainerIPv6Mask    = net.CIDRMask(128, 128)
)
