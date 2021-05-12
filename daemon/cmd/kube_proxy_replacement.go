// Copyright 2019-2020 Authors of Cilium
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

// This module contains the kube-proxy replacement initialization helpers.

package cmd

import (
	"errors"
	"fmt"
	"math"
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/cilium/cilium/pkg/bpf"
	linuxdatapath "github.com/cilium/cilium/pkg/datapath/linux"
	"github.com/cilium/cilium/pkg/datapath/linux/probes"
	"github.com/cilium/cilium/pkg/datapath/loader"
	datapathOption "github.com/cilium/cilium/pkg/datapath/option"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/mac"
	"github.com/cilium/cilium/pkg/maglev"
	"github.com/cilium/cilium/pkg/node"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/probe"
	"github.com/cilium/cilium/pkg/sysctl"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func initKubeProxyReplacementOptions() (strict bool) {
	if option.Config.KubeProxyReplacement != option.KubeProxyReplacementStrict &&
		option.Config.KubeProxyReplacement != option.KubeProxyReplacementPartial &&
		option.Config.KubeProxyReplacement != option.KubeProxyReplacementProbe &&
		option.Config.KubeProxyReplacement != option.KubeProxyReplacementDisabled {
		log.Fatalf("Invalid value for --%s: %s", option.KubeProxyReplacement, option.Config.KubeProxyReplacement)
	}

	if option.Config.KubeProxyReplacement == option.KubeProxyReplacementDisabled {
		log.Infof("Auto-disabling %q, %q, %q, %q, %q features and falling back to %q",
			option.EnableNodePort, option.EnableExternalIPs,
			option.EnableHostReachableServices, option.EnableHostPort,
			option.EnableSessionAffinity, option.EnableHostLegacyRouting)

		disableNodePort()
		option.Config.EnableHostReachableServices = false
		option.Config.EnableHostServicesTCP = false
		option.Config.EnableHostServicesUDP = false
		option.Config.EnableSessionAffinity = false

		return
	}

	probesManager := probes.NewProbeManager()

	// strict denotes to panic if any to-be enabled feature cannot be enabled
	strict = option.Config.KubeProxyReplacement != option.KubeProxyReplacementProbe

	if option.Config.KubeProxyReplacement == option.KubeProxyReplacementProbe ||
		option.Config.KubeProxyReplacement == option.KubeProxyReplacementStrict {

		log.Infof("Trying to auto-enable %q, %q, %q, %q, %q features",
			option.EnableNodePort, option.EnableExternalIPs,
			option.EnableHostReachableServices, option.EnableHostPort,
			option.EnableSessionAffinity)

		option.Config.EnableHostPort = true
		option.Config.EnableNodePort = true
		option.Config.EnableExternalIPs = true
		option.Config.EnableHostReachableServices = true
		option.Config.EnableHostServicesTCP = true
		option.Config.EnableHostServicesUDP = true
		option.Config.EnableSessionAffinity = true
	}

	if option.Config.EnableNodePort {
		if option.Config.EnableIPSec {
			msg := "IPSec cannot be used with BPF NodePort."
			if strict {
				log.Fatal(msg)
			} else {
				disableNodePort()
				log.Warn(msg + " Disabling BPF NodePort feature.")
			}
		}

		if option.Config.NodePortMode != option.NodePortModeSNAT &&
			option.Config.NodePortMode != option.NodePortModeDSR &&
			option.Config.NodePortMode != option.NodePortModeHybrid {
			log.Fatalf("Invalid value for --%s: %s", option.NodePortMode, option.Config.NodePortMode)
		}

		if option.Config.NodePortMode == option.NodePortModeDSR &&
			option.Config.LoadBalancerDSRDispatch != option.DSRDispatchOption &&
			option.Config.LoadBalancerDSRDispatch != option.DSRDispatchIPIP ||
			option.Config.NodePortMode == option.NodePortModeHybrid &&
				option.Config.LoadBalancerDSRDispatch != option.DSRDispatchOption {
			log.Fatalf("Invalid value for --%s: %s", option.LoadBalancerDSRDispatch, option.Config.LoadBalancerDSRDispatch)
		}

		if option.Config.NodePortMode == option.NodePortModeDSR &&
			option.Config.LoadBalancerDSRL4Xlate != option.DSRL4XlateFrontend &&
			option.Config.LoadBalancerDSRL4Xlate != option.DSRL4XlateBackend {
			log.Fatalf("Invalid value for --%s: %s", option.LoadBalancerDSRL4Xlate, option.Config.LoadBalancerDSRL4Xlate)
		}

		if option.Config.LoadBalancerRSSv4CIDR != "" {
			ip, cidr, err := net.ParseCIDR(option.Config.LoadBalancerRSSv4CIDR)
			if ip.To4() == nil {
				err = fmt.Errorf("CIDR is not IPv4 based")
			}
			if err == nil {
				if ones, _ := cidr.Mask.Size(); ones == 0 {
					err = fmt.Errorf("CIDR length must be in (0,32]")
				}
			}
			if err != nil {
				log.WithError(err).Fatalf("Invalid value for --%s: %s",
					option.LoadBalancerRSSv4CIDR, option.Config.LoadBalancerRSSv4CIDR)
			}
			option.Config.LoadBalancerRSSv4 = *cidr
		}

		if option.Config.LoadBalancerRSSv6CIDR != "" {
			ip, cidr, err := net.ParseCIDR(option.Config.LoadBalancerRSSv6CIDR)
			if ip.To4() != nil {
				err = fmt.Errorf("CIDR is not IPv6 based")
			}
			if err == nil {
				if ones, _ := cidr.Mask.Size(); ones == 0 {
					err = fmt.Errorf("CIDR length must be in (0,128]")
				}
			}
			if err != nil {
				log.WithError(err).Fatalf("Invalid value for --%s: %s",
					option.LoadBalancerRSSv6CIDR, option.Config.LoadBalancerRSSv6CIDR)
			}
			option.Config.LoadBalancerRSSv6 = *cidr
		}

		if (option.Config.LoadBalancerRSSv4CIDR != "" || option.Config.LoadBalancerRSSv6CIDR != "") &&
			(option.Config.NodePortMode != option.NodePortModeDSR ||
				option.Config.LoadBalancerDSRDispatch != option.DSRDispatchIPIP) {
			log.Fatalf("Invalid value for --%s/%s: currently only supported under IPIP dispatch for DSR",
				option.LoadBalancerRSSv4CIDR, option.LoadBalancerRSSv6CIDR)
		}

		if option.Config.NodePortAlg != option.NodePortAlgRandom &&
			option.Config.NodePortAlg != option.NodePortAlgMaglev {
			log.Fatalf("Invalid value for --%s: %s", option.NodePortAlg, option.Config.NodePortAlg)
		}

		if option.Config.NodePortAcceleration != option.NodePortAccelerationDisabled &&
			option.Config.NodePortAcceleration != option.NodePortAccelerationGeneric &&
			option.Config.NodePortAcceleration != option.NodePortAccelerationNative {
			log.Fatalf("Invalid value for --%s: %s", option.NodePortAcceleration, option.Config.NodePortAcceleration)
		}

		if option.Config.KubeProxyReplacement == option.KubeProxyReplacementProbe {
			// We let kube-proxy do the less efficient bind-protection in
			// this case to avoid the latter throwing (harmless) warnings
			// to its log that bind request is rejected.
			option.Config.NodePortBindProtection = false
		} else if !option.Config.NodePortBindProtection {
			log.Warning("NodePort BPF configured without bind(2) protection against service ports")
		}

		if option.Config.NodePortAlg == option.NodePortAlgMaglev {
			// "Let N be the size of a VIP's backend pool." [...] "In practice, we choose M to be
			// larger than 100 x N to ensure at most a 1% difference in hash space assigned to
			// backends." (from Maglev paper, page 6)
			supportedPrimes := []int{251, 509, 1021, 2039, 4093, 8191, 16381, 32749, 65521, 131071}
			found := false
			for _, prime := range supportedPrimes {
				if option.Config.MaglevTableSize == prime {
					found = true
					break
				}
			}
			if !found {
				log.Fatalf("Invalid value for --%s: %d, supported values are: %v",
					option.MaglevTableSize, option.Config.MaglevTableSize, supportedPrimes)
			}
			if err := maglev.Init(
				option.Config.MaglevHashSeed,
				uint64(option.Config.MaglevTableSize),
			); err != nil {
				log.WithError(err).Fatalf("Failed to initialize maglev hash seeds")
			}
		}
	}

	if option.Config.EnableNodePort {
		found := false
		if h := probesManager.GetHelpers("sched_act"); h != nil {
			if _, ok := h["bpf_fib_lookup"]; ok {
				found = true
			}
		}
		if !found {
			msg := "BPF NodePort services needs kernel 4.17.0 or newer."
			if strict {
				log.Fatal(msg)
			} else {
				disableNodePort()
				log.Warn(msg + " Disabling BPF NodePort.")
			}
		}

		if err := checkNodePortAndEphemeralPortRanges(); err != nil {
			if strict {
				log.Fatal(err)
			} else {
				disableNodePort()
				log.WithError(err).Warn("Disabling BPF NodePort.")
			}
		}
	}

	if option.Config.EnableHostReachableServices {
		// Try to auto-load IPv6 module if it hasn't been done yet as there can
		// be v4-in-v6 connections even if the agent has v6 support disabled.
		probe.HaveIPv6Support()

		option.Config.EnableHostServicesPeer = true
		if option.Config.EnableIPv4 {
			if err := bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_INET4_GETPEERNAME); err != nil {
				option.Config.EnableHostServicesPeer = false
			}
		}
		if option.Config.EnableIPv6 {
			if err := bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_INET6_GETPEERNAME); err != nil {
				option.Config.EnableHostServicesPeer = false
			}
		}
		if option.Config.EnableHostServicesTCP && option.Config.EnableIPv4 {
			probeCgroupSupportTCP(strict, true)
		}
		if option.Config.EnableHostServicesTCP && option.Config.EnableIPv6 {
			probeCgroupSupportTCP(strict, false)
		}
		if option.Config.EnableHostServicesUDP && option.Config.EnableIPv4 {
			probeCgroupSupportUDP(strict, true)
		}
		if option.Config.EnableHostServicesUDP && option.Config.EnableIPv6 {
			probeCgroupSupportUDP(strict, false)
		}
		if !option.Config.EnableHostServicesTCP && !option.Config.EnableHostServicesUDP {
			option.Config.EnableHostReachableServices = false
		}
	} else {
		option.Config.EnableHostServicesTCP = false
		option.Config.EnableHostServicesUDP = false
	}

	if option.Config.EnableSessionAffinity {
		if !probesManager.GetMapTypes().HaveLruHashMapType {
			msg := "SessionAffinity feature requires BPF LRU maps"
			if strict {
				log.Fatal(msg)
			} else {
				log.Warnf("%s. Disabling the feature.", msg)
				option.Config.EnableSessionAffinity = false
			}

		}
	}
	if option.Config.EnableSessionAffinity && option.Config.EnableHostReachableServices {
		found1, found2 := false, false
		if h := probesManager.GetHelpers("cgroup_sock"); h != nil {
			_, found1 = h["bpf_get_netns_cookie"]
		}
		if h := probesManager.GetHelpers("cgroup_sock_addr"); h != nil {
			_, found2 = h["bpf_get_netns_cookie"]
		}
		if !(found1 && found2) {
			log.Warn("Session affinity for host reachable services needs kernel 5.7.0 or newer " +
				"to work properly when accessed from inside cluster: the same service endpoint " +
				"will be selected from all network namespaces on the host.")
		}
	}

	if option.Config.EnableNodePort {
		if option.Config.Tunnel != option.TunnelDisabled &&
			option.Config.NodePortMode != option.NodePortModeSNAT {

			log.Warnf("Disabling NodePort's %q mode feature due to tunneling mode being enabled",
				option.Config.NodePortMode)
			option.Config.NodePortMode = option.NodePortModeSNAT
		}

		if option.Config.NodePortAcceleration != option.NodePortAccelerationDisabled {
			if option.Config.Tunnel != option.TunnelDisabled {
				log.Fatalf("Cannot use NodePort acceleration with tunneling. Either run cilium-agent with --%s=%s or --%s=%s",
					option.NodePortAcceleration, option.NodePortAccelerationDisabled, option.TunnelName, option.TunnelDisabled)
			}

			if option.Config.EnableEgressGateway {
				log.Fatalf("Cannot use NodePort acceleration with the egress gateway. Run cilium-agent with either --%s=%s or %s=false",
					option.NodePortAcceleration, option.NodePortAccelerationDisabled, option.EnableEgressGateway)
			}
		}

		if option.Config.NodePortMode == option.NodePortModeDSR &&
			option.Config.LoadBalancerDSRDispatch == option.DSRDispatchIPIP {
			if option.Config.DatapathMode != datapathOption.DatapathModeLBOnly {
				log.Fatalf("DSR dispatch mode %s only supported for --%s=%s", option.Config.LoadBalancerDSRDispatch, option.DatapathMode, datapathOption.DatapathModeLBOnly)
			}
			if option.Config.NodePortAcceleration == option.NodePortAccelerationDisabled {
				log.Fatalf("DSR dispatch mode %s currently only available under XDP acceleration", option.Config.LoadBalancerDSRDispatch)
			}
		}

		if option.Config.EnableRecorder {
			if option.Config.DatapathMode != datapathOption.DatapathModeLBOnly {
				log.Fatalf("pcap recorder --%s currently only supported for --%s=%s", option.EnableRecorder, option.DatapathMode, datapathOption.DatapathModeLBOnly)
			}
			found := false
			if h := probesManager.GetHelpers("xdp"); h != nil {
				if _, ok := h["bpf_ktime_get_boot_ns"]; ok {
					found = true
				}
			}
			if !found {
				log.Fatalf("pcap recorder --%s datapath needs kernel 5.8.0 or newer", option.EnableRecorder)
			}
		}

		option.Config.EnableHealthDatapath =
			option.Config.DatapathMode == datapathOption.DatapathModeLBOnly &&
				option.Config.NodePortMode == option.NodePortModeDSR &&
				option.Config.LoadBalancerDSRDispatch == option.DSRDispatchIPIP
		if option.Config.EnableHealthDatapath {
			found := false
			if h := probesManager.GetHelpers("cgroup_sock_addr"); h != nil {
				if _, ok := h["bpf_getsockopt"]; ok {
					found = true
				}
			}
			if !found {
				option.Config.EnableHealthDatapath = false
				log.Info("BPF load-balancer health check datapath needs kernel 5.12.0 or newer. Disabling BPF load-balancer health check datapath.")
			}
		}
	}

	if option.Config.InstallNoConntrackIptRules {
		// InstallNoConntrackIptRules can only be enabled when Cilium is
		// running in full KPR mode as otherwise conntrack would be
		// required for NAT operations
		if !option.Config.KubeProxyReplacementFullyEnabled() {
			log.Fatalf("%s requires the agent to run with %s=%s.",
				option.InstallNoConntrackIptRules, option.KubeProxyReplacement, option.KubeProxyReplacementStrict)
		}

		if !option.Config.EnableBPFMasquerade {
			log.Fatalf("%s requires the agent to run with %s.",
				option.InstallNoConntrackIptRules, option.EnableBPFMasquerade)
		}
	}

	return
}

func probeCgroupSupportTCP(strict, ipv4 bool) {
	var err error

	if ipv4 {
		err = bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_INET4_CONNECT)
	} else {
		err = bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_INET6_CONNECT)
	}
	if err != nil {
		scopedLog := log.WithError(err)
		msg := "BPF host reachable services for TCP needs kernel 4.17.0 or newer."
		if errors.Is(err, unix.EPERM) {
			msg = "Cilium cannot load bpf programs. Security profiles like SELinux may be restricting permissions."
		}

		if strict {
			scopedLog.Fatal(msg)
		} else {
			option.Config.EnableHostServicesTCP = false
			scopedLog.Warn(msg + " Disabling the feature.")
		}
	}
}

func probeCgroupSupportUDP(strict, ipv4 bool) {
	var err error

	if ipv4 {
		err = bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_UDP4_RECVMSG)
	} else {
		err = bpf.TestDummyProg(bpf.ProgTypeCgroupSockAddr, bpf.BPF_CGROUP_UDP6_RECVMSG)
	}
	if err != nil {
		scopedLog := log.WithError(err)
		msg := fmt.Sprintf("BPF host reachable services for UDP needs kernel 4.19.57, 5.1.16, 5.2.0 or newer. If you run an older kernel and only need TCP, then specify: --%s=tcp and --%s=%s", option.HostReachableServicesProtos, option.KubeProxyReplacement, option.KubeProxyReplacementPartial)
		if errors.Is(err, unix.EPERM) {
			msg = "Cilium cannot load bpf programs. Security profiles like SELinux may be restricting permissions."
		}

		if strict {
			scopedLog.Fatal(msg)
		} else {
			option.Config.EnableHostServicesUDP = false
			scopedLog.Warn(msg + " Disabling the feature.")
		}
	}
}

// handleNativeDevices tries to detect bpf_host devices (if needed).
func handleNativeDevices(strict bool) {
	expandDevices()

	detectNodePortDevs := len(option.Config.Devices) == 0 &&
		(option.Config.EnableNodePort || option.Config.EnableHostFirewall || option.Config.EnableBandwidthManager)
	detectDirectRoutingDev := option.Config.EnableNodePort &&
		option.Config.DirectRoutingDevice == ""
	if detectNodePortDevs || detectDirectRoutingDev {
		if err := detectDevices(detectNodePortDevs, detectDirectRoutingDev); err != nil {
			msg := "Unable to detect devices for BPF NodePort."
			if strict {
				log.WithError(err).Fatal(msg)
			} else {
				disableNodePort()
				log.WithError(err).Warn(msg + " Disabling BPF NodePort.")
			}
		} else {
			l := log
			if detectNodePortDevs {
				l = l.WithField(logfields.Devices, option.Config.Devices)
			}
			if detectDirectRoutingDev {
				l = l.WithField(logfields.DirectRoutingDevice, option.Config.DirectRoutingDevice)
			}
			l.Info("Using auto-derived devices for BPF node port")
		}
	} else if option.Config.EnableNodePort { // both --devices and --direct-routing-device are specified by user
		// Check whether the DirectRoutingDevice (if specified) is
		// defined within devices and if not, add it.
		if option.Config.DirectRoutingDevice != "" {
			directDev := option.Config.DirectRoutingDevice
			directDevFound := false
			for _, iface := range option.Config.Devices {
				if iface == directDev {
					directDevFound = true
					break
				}
			}
			if !directDevFound {
				option.Config.Devices = append(option.Config.Devices, directDev)
			}
		}
	}
}

// finishKubeProxyReplacementInit finishes initialization of kube-proxy
// replacement after all devices are known.
func finishKubeProxyReplacementInit(isKubeProxyReplacementStrict bool) {
	if option.Config.EnableNodePort {
		if err := node.InitNodePortAddrs(option.Config.Devices, option.Config.LBDevInheritIPAddr); err != nil {
			msg := "Failed to initialize NodePort addrs."
			if isKubeProxyReplacementStrict {
				log.WithError(err).Fatal(msg)
			} else {
				disableNodePort()
				log.WithError(err).Warn(msg + " Disabling BPF NodePort.")
			}
		}
	}

	if !option.Config.EnableNodePort {
		// Make sure that NodePort dependencies are disabled
		disableNodePort()
		return
	}

	if option.Config.EnableSVCSourceRangeCheck && !probe.HaveFullLPM() {
		msg := fmt.Sprintf("--%s requires kernel 4.16 or newer.",
			option.EnableSVCSourceRangeCheck)
		if isKubeProxyReplacementStrict {
			log.Fatal(msg)
		} else {
			log.Warnf(msg + " Disabling the check.")
			option.Config.EnableSVCSourceRangeCheck = false
		}
	}

	// +-------------------------------------------------------+
	// | After this point, BPF NodePort should not be disabled |
	// +-------------------------------------------------------+

	if !option.Config.EnableHostLegacyRouting {
		msg := ""
		switch {
		// Needs host stack for packet handling.
		case option.Config.EnableIPSec:
			msg = fmt.Sprintf("BPF host routing is incompatible with %s.", option.EnableIPSecName)
		// Non-BPF masquerade requires netfilter and hence CT.
		case (option.Config.EnableIPv4Masquerade || option.Config.EnableIPv6Masquerade) &&
			!option.Config.EnableBPFMasquerade:
			msg = fmt.Sprintf("BPF host routing requires %s.", option.EnableBPFMasquerade)
		// All cases below still need to be implemented ...
		case option.Config.EnableEndpointRoutes:
			msg = fmt.Sprintf("BPF host routing is currently not supported with %s.", option.EnableEndpointRoutes)
		case !mac.HaveMACAddrs(option.Config.Devices):
			msg = "BPF host routing is currently not supported with devices without L2 addr."
		case option.Config.EnableWireguard:
			msg = fmt.Sprintf("BPF host routing is currently not compatible with Wireguard (--%s).", option.EnableWireguard)
		default:
			probesManager := probes.NewProbeManager()
			foundNeigh := false
			foundPeer := false
			if h := probesManager.GetHelpers("sched_cls"); h != nil {
				_, foundNeigh = h["bpf_redirect_neigh"]
				_, foundPeer = h["bpf_redirect_peer"]
			}
			if !foundNeigh || !foundPeer {
				msg = fmt.Sprintf("BPF host routing requires kernel 5.10 or newer.")
			}
		}
		if msg != "" {
			option.Config.EnableHostLegacyRouting = true
			log.Infof("%s Falling back to legacy host routing (%s=true).", msg, option.EnableHostLegacyRouting)
		}
	}

	if option.Config.NodePortAcceleration != option.NodePortAccelerationDisabled {
		if option.Config.XDPDevice != "undefined" &&
			(option.Config.DirectRoutingDevice == "" ||
				option.Config.XDPDevice != option.Config.DirectRoutingDevice) {
			log.Fatalf("Cannot set NodePort acceleration device: mismatch between Prefilter device %s and NodePort device %s",
				option.Config.XDPDevice, option.Config.DirectRoutingDevice)
		}
		option.Config.XDPDevice = option.Config.DirectRoutingDevice
		if err := loader.SetXDPMode(option.Config.NodePortAcceleration); err != nil {
			log.WithError(err).Fatal("Cannot set NodePort acceleration")
		}
	}

	for _, iface := range option.Config.Devices {
		link, err := netlink.LinkByName(iface)
		if err != nil {
			log.WithError(err).Fatalf("Cannot retrieve %s link", iface)
		}
		if strings.ContainsAny(iface, "=;") {
			// Because we pass IPV{4,6}_NODEPORT addresses to bpf/init.sh
			// in a form "$IFACE_NAME1=$IPV{4,6}_ADDR1;$IFACE_NAME2=...",
			// we need to restrict the iface names. Otherwise, bpf/init.sh
			// won't properly parse the mappings.
			log.Fatalf("%s link name contains '=' or ';' character which is not allowed",
				iface)
		}
		if idx := link.Attrs().Index; idx > math.MaxUint16 {
			log.Fatalf("%s link ifindex %d exceeds max(uint16)", iface, idx)
		}
	}

	if option.Config.EnableIPv4 &&
		option.Config.Tunnel == option.TunnelDisabled &&
		option.Config.NodePortMode != option.NodePortModeSNAT &&
		len(option.Config.Devices) > 1 {

		// In the case of the multi-dev NodePort DSR, if a request from an
		// external client was sent to a device which is not used for direct
		// routing, such request might be dropped by the destination node
		// if the destination node's direct routing device's rp_filter = 1
		// and the client IP is reachable via other device than the direct
		// routing one.

		iface := option.Config.DirectRoutingDevice
		if val, err := sysctl.Read(fmt.Sprintf("net.ipv4.conf.%s.rp_filter", iface)); err != nil {
			log.Warnf("Unable to read net.ipv4.conf.%s.rp_filter: %s. Ignoring the check",
				iface, err)
		} else {
			if val == "1" {
				log.Warnf(`DSR might not work for requests sent to other than %s device. `+
					`Run 'sysctl -w net.ipv4.conf.%s.rp_filter=2' (or set to '0') on each node to fix`,
					iface, iface)
			}
		}
	}

	option.Config.NodePortHairpin = len(option.Config.Devices) == 1
	if option.Config.NodePortAcceleration != option.NodePortAccelerationDisabled &&
		len(option.Config.Devices) != 1 {
		log.Fatalf("Cannot set NodePort acceleration due to multi-device setup (%q). Specify --%s with a single device to enable NodePort acceleration.", option.Config.Devices, option.Devices)
	}
}

// disableNodePort disables BPF NodePort and friends who are dependent from
// the latter.
func disableNodePort() {
	option.Config.EnableNodePort = false
	option.Config.EnableHostPort = false
	option.Config.EnableExternalIPs = false
	option.Config.EnableSVCSourceRangeCheck = false
	option.Config.EnableHostLegacyRouting = true
}

// detectDevices tries to detect device names which are going to be used for
// (a) NodePort BPF, (b) direct routing in NodePort BPF.
//
// (a) is determined from a default route and the k8s node IP addr.
// (b) is derived either from NodePort BPF devices (if only one is set) or
//     from the k8s node IP addr.
func detectDevices(detectNodePortDevs, detectDirectRoutingDev bool) error {
	var err error
	devSet := map[string]struct{}{} // iface name
	ifidxByAddr := map[string]int{} // str(ip addr) => ifindex

	if addrs, err := netlink.AddrList(nil, netlink.FAMILY_ALL); err != nil {
		// Do not return error, as a device with a default route can be used
		// later as a last resort
		log.WithError(err).Warn(
			"Cannot retrieve host IP addrs for BPF NodePort device detection")
	} else {
		for _, a := range addrs {
			// Any cilium_* interface will never be a valid NodePort or Direct Routing
			// interface. Skip interface if we cannot resolve it from Netlink via its
			// ifIndex or if its name begins with cilium_.
			if link, err := netlink.LinkByIndex(a.LinkIndex); err != nil {
				log.WithError(err).WithField(logfields.LinkIndex, a.LinkIndex).Warn(
					"Unable to resolve link from ifIndex, skipping interface for device detection")
			} else if strings.HasPrefix(link.Attrs().Name, "cilium_") {
				log.WithField(logfields.Device, link.Attrs().Name).Debug(
					"Skipping Cilium-generated interface for device detection")
			} else {
				ifidxByAddr[a.IP.String()] = a.LinkIndex
			}
		}
	}

	if detectNodePortDevs {
		if devSet, err = detectNodePortDevices(ifidxByAddr); err != nil {
			return fmt.Errorf("Unable to determine BPF NodePort devices: %s. Use --%s to specify them",
				err, option.Devices)
		}
	} else {
		for _, dev := range option.Config.Devices {
			devSet[dev] = struct{}{}
		}
	}

	if detectDirectRoutingDev {
		// If only single device was previously found, use it for direct routing.
		// Otherwise, use k8s Node IP addr to determine the device.
		if len(devSet) == 1 {
			for dev := range devSet {
				option.Config.DirectRoutingDevice = dev
			}
		} else {
			if option.Config.DirectRoutingDevice, err = detectNodeDevice(ifidxByAddr); err != nil {
				return fmt.Errorf("Unable to determine BPF NodePort direct routing device: %s. "+
					"Use --%s to specify it", err, option.DirectRoutingDevice)
			}
		}
	}
	if option.Config.DirectRoutingDevice != "" {
		devSet[option.Config.DirectRoutingDevice] = struct{}{}
	}

	l3DevOK := supportL3Dev()
	option.Config.Devices = make([]string, 0, len(devSet))
	for dev := range devSet {
		if !l3DevOK && !mac.HasMacAddr(dev) {
			log.WithField(logfields.Device, dev).
				Warn("Ignoring L3 device; >= 5.8 kernel is required.")
			continue
		}
		option.Config.Devices = append(option.Config.Devices, dev)
	}
	return nil
}

func detectNodePortDevices(ifidxByAddr map[string]int) (map[string]struct{}, error) {
	devSet := map[string]struct{}{}

	// Find a device with a default route (for backward compatibility)
	defaultRouteDevice, err := linuxdatapath.NodeDeviceNameWithDefaultRoute()
	if err == nil {
		devSet[defaultRouteDevice] = struct{}{}
	}

	// Derive a device from k8s Node IP
	if dev, err := detectNodeDevice(ifidxByAddr); err != nil {
		log.WithError(err).Warn(
			"Cannot determine a device from k8s Node IP addr for BPF NodePort device detection")
	} else {
		devSet[dev] = struct{}{}
	}

	if len(devSet) == 0 {
		return nil, fmt.Errorf("Cannot determine any device for BPF NodePort")
	}

	return devSet, nil
}

func detectNodeDevice(ifidxByAddr map[string]int) (string, error) {
	nodeIP := node.GetK8sNodeIP()
	if nodeIP == nil {
		return "", fmt.Errorf("K8s Node IP is not set")
	}

	ifindex, found := ifidxByAddr[nodeIP.String()]
	if !found {
		return "", fmt.Errorf("Cannot find device with %s addr", nodeIP)
	}

	link, err := netlink.LinkByIndex(ifindex)
	if err != nil {
		return "", fmt.Errorf("Cannot find device with %s addr by %d ifindex",
			nodeIP, ifindex)
	}

	return link.Attrs().Name, nil
}

// expandDevices expands all wildcard device names to concrete devices.
// e.g. device "eth+" expands to "eth0,eth1" etc. Non-matching wildcards are ignored.
func expandDevices() {
	allLinks, err := netlink.LinkList()
	if err != nil {
		log.WithError(err).Fatal("Cannot list network devices via netlink")
	}
	expandedDevices := make(map[string]bool)
	for _, iface := range option.Config.Devices {
		if strings.HasSuffix(iface, "+") {
			prefix := strings.TrimRight(iface, "+")
			for _, link := range allLinks {
				attrs := link.Attrs()
				if strings.HasPrefix(attrs.Name, prefix) {
					expandedDevices[attrs.Name] = true
				}
			}
		} else {
			expandedDevices[iface] = true
		}
	}
	option.Config.Devices = make([]string, 0, len(expandedDevices))
	for dev := range expandedDevices {
		option.Config.Devices = append(option.Config.Devices, dev)
	}
	sort.Strings(option.Config.Devices)
}

// checkNodePortAndEphemeralPortRanges checks whether the ephemeral port range
// does not clash with the nodeport range to prevent the BPF nodeport from
// hijacking an existing connection on the local host which source port is
// the same as a nodeport service.
//
// If it clashes, check whether the nodeport range is listed in ip_local_reserved_ports.
// If it isn't and EnableAutoProtectNodePortRange == false, then return an error
// making cilium-agent to stop.
// Otherwise, if EnableAutoProtectNodePortRange == true, then append the nodeport
// range to ip_local_reserved_ports.
func checkNodePortAndEphemeralPortRanges() error {
	ephemeralPortRangeStr, err := sysctl.Read("net.ipv4.ip_local_port_range")
	if err != nil {
		return fmt.Errorf("Unable to read net.ipv4.ip_local_port_range")
	}
	ephemeralPortRange := strings.Split(ephemeralPortRangeStr, "\t")
	if len(ephemeralPortRange) != 2 {
		return fmt.Errorf("Invalid ephemeral port range: %s", ephemeralPortRangeStr)
	}
	ephemeralPortMin, err := strconv.Atoi(ephemeralPortRange[0])
	if err != nil {
		return fmt.Errorf("Unable to parse min port value %s for ephemeral range", ephemeralPortRange[0])
	}
	ephemeralPortMax, err := strconv.Atoi(ephemeralPortRange[1])
	if err != nil {
		return fmt.Errorf("Unable to parse max port value %s for ephemeral range", ephemeralPortRange[1])
	}

	if option.Config.NodePortMax < ephemeralPortMin {
		// ephemeral port range does not clash with nodeport range
		return nil
	}

	nodePortRangeStr := fmt.Sprintf("%d-%d", option.Config.NodePortMin,
		option.Config.NodePortMax)

	if option.Config.NodePortMin > ephemeralPortMax {
		return fmt.Errorf("NodePort port range (%s) is not allowed to be after ephemeral port range (%s)",
			nodePortRangeStr, ephemeralPortRangeStr)
	}

	reservedPortsStr, err := sysctl.Read("net.ipv4.ip_local_reserved_ports")
	if err != nil {
		return fmt.Errorf("Unable to read net.ipv4.ip_local_reserved_ports")
	}
	for _, portRange := range strings.Split(reservedPortsStr, ",") {
		if portRange == "" {
			break
		}
		ports := strings.Split(portRange, "-")
		if len(ports) == 0 {
			return fmt.Errorf("Invalid reserved ports range")
		}
		from, err := strconv.Atoi(ports[0])
		if err != nil {
			return fmt.Errorf("Unable to parse reserved port %q", ports[0])
		}
		to := from
		if len(ports) == 2 {
			if to, err = strconv.Atoi(ports[1]); err != nil {
				return fmt.Errorf("Unable to parse reserved port %q", ports[1])
			}
		}

		if from <= option.Config.NodePortMin && to >= option.Config.NodePortMax {
			// nodeport range is protected by reserved port range
			return nil
		}

		if from > option.Config.NodePortMax {
			break
		}
	}

	if !option.Config.EnableAutoProtectNodePortRange {
		msg := `NodePort port range (%s) must not clash with ephemeral port range (%s). ` +
			`Adjust ephemeral range port with "sysctl -w net.ipv4.ip_local_port_range='MIN MAX'", or ` +
			`protect the NodePort range by appending it to "net.ipv4.ip_local_reserved_ports", or ` +
			`set --%s=true to auto-append the range to "net.ipv4.ip_local_reserved_ports"`
		return fmt.Errorf(msg, nodePortRangeStr, ephemeralPortRangeStr,
			option.EnableAutoProtectNodePortRange)
	}

	if reservedPortsStr != "" {
		reservedPortsStr += ","
	}
	reservedPortsStr += fmt.Sprintf("%d-%d", option.Config.NodePortMin, option.Config.NodePortMax)
	if err := sysctl.Write("net.ipv4.ip_local_reserved_ports", reservedPortsStr); err != nil {
		return fmt.Errorf("Unable to addend nodeport range (%s) to net.ipv4.ip_local_reserved_ports: %s",
			nodePortRangeStr, err)
	}

	return nil
}

func hasFullHostReachableServices() bool {
	return option.Config.EnableHostReachableServices &&
		option.Config.EnableHostServicesTCP &&
		option.Config.EnableHostServicesUDP
}

func supportL3Dev() bool {
	probesManager := probes.NewProbeManager()
	if h := probesManager.GetHelpers("sched_cls"); h != nil {
		_, found := h["bpf_skb_change_head"]
		return found
	}
	return false
}
