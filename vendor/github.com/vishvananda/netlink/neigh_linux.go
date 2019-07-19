package netlink

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"

	"github.com/vishvananda/netlink/nl"
	"github.com/vishvananda/netns"
	"golang.org/x/sys/unix"
)

const (
	NDA_UNSPEC = iota
	NDA_DST
	NDA_LLADDR
	NDA_CACHEINFO
	NDA_PROBES
	NDA_VLAN
	NDA_PORT
	NDA_VNI
	NDA_IFINDEX
	NDA_MAX = NDA_IFINDEX
)

// Neighbor Cache Entry States.
const (
	NUD_NONE       = 0x00
	NUD_INCOMPLETE = 0x01
	NUD_REACHABLE  = 0x02
	NUD_STALE      = 0x04
	NUD_DELAY      = 0x08
	NUD_PROBE      = 0x10
	NUD_FAILED     = 0x20
	NUD_NOARP      = 0x40
	NUD_PERMANENT  = 0x80
)

// Neighbor Flags
const (
	NTF_USE    = 0x01
	NTF_SELF   = 0x02
	NTF_MASTER = 0x04
	NTF_PROXY  = 0x08
	NTF_ROUTER = 0x80
)

// Ndmsg is for adding, removing or receiving information about a neighbor table entry
type Ndmsg struct {
	Family uint8
	Index  uint32
	State  uint16
	Flags  uint8
	Type   uint8
}

func deserializeNdmsg(b []byte) *Ndmsg {
	var dummy Ndmsg
	return (*Ndmsg)(unsafe.Pointer(&b[0:unsafe.Sizeof(dummy)][0]))
}

func (msg *Ndmsg) Serialize() []byte {
	return (*(*[unsafe.Sizeof(*msg)]byte)(unsafe.Pointer(msg)))[:]
}

func (msg *Ndmsg) Len() int {
	return int(unsafe.Sizeof(*msg))
}

// NeighAdd will add an IP to MAC mapping to the ARP table
// Equivalent to: `ip neigh add ....`
func NeighAdd(neigh *Neigh) error {
	return pkgHandle.NeighAdd(neigh)
}

// NeighAdd will add an IP to MAC mapping to the ARP table
// Equivalent to: `ip neigh add ....`
func (h *Handle) NeighAdd(neigh *Neigh) error {
	return h.neighAdd(neigh, unix.NLM_F_CREATE|unix.NLM_F_EXCL)
}

// NeighSet will add or replace an IP to MAC mapping to the ARP table
// Equivalent to: `ip neigh replace....`
func NeighSet(neigh *Neigh) error {
	return pkgHandle.NeighSet(neigh)
}

// NeighSet will add or replace an IP to MAC mapping to the ARP table
// Equivalent to: `ip neigh replace....`
func (h *Handle) NeighSet(neigh *Neigh) error {
	return h.neighAdd(neigh, unix.NLM_F_CREATE|unix.NLM_F_REPLACE)
}

// NeighAppend will append an entry to FDB
// Equivalent to: `bridge fdb append...`
func NeighAppend(neigh *Neigh) error {
	return pkgHandle.NeighAppend(neigh)
}

// NeighAppend will append an entry to FDB
// Equivalent to: `bridge fdb append...`
func (h *Handle) NeighAppend(neigh *Neigh) error {
	return h.neighAdd(neigh, unix.NLM_F_CREATE|unix.NLM_F_APPEND)
}

// NeighAppend will append an entry to FDB
// Equivalent to: `bridge fdb append...`
func neighAdd(neigh *Neigh, mode int) error {
	return pkgHandle.neighAdd(neigh, mode)
}

// NeighAppend will append an entry to FDB
// Equivalent to: `bridge fdb append...`
func (h *Handle) neighAdd(neigh *Neigh, mode int) error {
	req := h.newNetlinkRequest(unix.RTM_NEWNEIGH, mode|unix.NLM_F_ACK)
	return neighHandle(neigh, req)
}

// NeighDel will delete an IP address from a link device.
// Equivalent to: `ip addr del $addr dev $link`
func NeighDel(neigh *Neigh) error {
	return pkgHandle.NeighDel(neigh)
}

// NeighDel will delete an IP address from a link device.
// Equivalent to: `ip addr del $addr dev $link`
func (h *Handle) NeighDel(neigh *Neigh) error {
	req := h.newNetlinkRequest(unix.RTM_DELNEIGH, unix.NLM_F_ACK)
	return neighHandle(neigh, req)
}

func neighHandle(neigh *Neigh, req *nl.NetlinkRequest) error {
	var family int

	if neigh.Family > 0 {
		family = neigh.Family
	} else {
		family = nl.GetIPFamily(neigh.IP)
	}

	msg := Ndmsg{
		Family: uint8(family),
		Index:  uint32(neigh.LinkIndex),
		State:  uint16(neigh.State),
		Type:   uint8(neigh.Type),
		Flags:  uint8(neigh.Flags),
	}
	req.AddData(&msg)

	ipData := neigh.IP.To4()
	if ipData == nil {
		ipData = neigh.IP.To16()
	}

	dstData := nl.NewRtAttr(NDA_DST, ipData)
	req.AddData(dstData)

	if neigh.LLIPAddr != nil {
		llIPData := nl.NewRtAttr(NDA_LLADDR, neigh.LLIPAddr.To4())
		req.AddData(llIPData)
	} else if neigh.Flags != NTF_PROXY || neigh.HardwareAddr != nil {
		hwData := nl.NewRtAttr(NDA_LLADDR, []byte(neigh.HardwareAddr))
		req.AddData(hwData)
	}

	if neigh.Vlan != 0 {
		vlanData := nl.NewRtAttr(NDA_VLAN, nl.Uint16Attr(uint16(neigh.Vlan)))
		req.AddData(vlanData)
	}

	if neigh.VNI != 0 {
		vniData := nl.NewRtAttr(NDA_VNI, nl.Uint32Attr(uint32(neigh.VNI)))
		req.AddData(vniData)
	}

	_, err := req.Execute(unix.NETLINK_ROUTE, 0)
	return err
}

// NeighList returns a list of IP-MAC mappings in the system (ARP table).
// Equivalent to: `ip neighbor show`.
// The list can be filtered by link and ip family.
func NeighList(linkIndex, family int) ([]Neigh, error) {
	return pkgHandle.NeighList(linkIndex, family)
}

// NeighProxyList returns a list of neighbor proxies in the system.
// Equivalent to: `ip neighbor show proxy`.
// The list can be filtered by link and ip family.
func NeighProxyList(linkIndex, family int) ([]Neigh, error) {
	return pkgHandle.NeighProxyList(linkIndex, family)
}

// NeighList returns a list of IP-MAC mappings in the system (ARP table).
// Equivalent to: `ip neighbor show`.
// The list can be filtered by link and ip family.
func (h *Handle) NeighList(linkIndex, family int) ([]Neigh, error) {
	return h.NeighListExecute(Ndmsg{
		Family: uint8(family),
		Index:  uint32(linkIndex),
	})
}

// NeighProxyList returns a list of neighbor proxies in the system.
// Equivalent to: `ip neighbor show proxy`.
// The list can be filtered by link, ip family.
func (h *Handle) NeighProxyList(linkIndex, family int) ([]Neigh, error) {
	return h.NeighListExecute(Ndmsg{
		Family: uint8(family),
		Index:  uint32(linkIndex),
		Flags:  NTF_PROXY,
	})
}

// NeighListExecute returns a list of neighbour entries filtered by link, ip family, flag and state.
func NeighListExecute(msg Ndmsg) ([]Neigh, error) {
	return pkgHandle.NeighListExecute(msg)
}

// NeighListExecute returns a list of neighbour entries filtered by link, ip family, flag and state.
func (h *Handle) NeighListExecute(msg Ndmsg) ([]Neigh, error) {
	req := h.newNetlinkRequest(unix.RTM_GETNEIGH, unix.NLM_F_DUMP)
	req.AddData(&msg)

	msgs, err := req.Execute(unix.NETLINK_ROUTE, unix.RTM_NEWNEIGH)
	if err != nil {
		return nil, err
	}

	var res []Neigh
	for _, m := range msgs {
		ndm := deserializeNdmsg(m)
		if msg.Index != 0 && ndm.Index != msg.Index {
			// Ignore messages from other interfaces
			continue
		}

		neigh, err := NeighDeserialize(m)
		if err != nil {
			continue
		}

		res = append(res, *neigh)
	}

	return res, nil
}

func NeighDeserialize(m []byte) (*Neigh, error) {
	msg := deserializeNdmsg(m)

	neigh := Neigh{
		LinkIndex: int(msg.Index),
		Family:    int(msg.Family),
		State:     int(msg.State),
		Type:      int(msg.Type),
		Flags:     int(msg.Flags),
	}

	attrs, err := nl.ParseRouteAttr(m[msg.Len():])
	if err != nil {
		return nil, err
	}

	for _, attr := range attrs {
		switch attr.Attr.Type {
		case NDA_DST:
			neigh.IP = net.IP(attr.Value)
		case NDA_LLADDR:
			// BUG: Is this a bug in the netlink library?
			// #define RTA_LENGTH(len) (RTA_ALIGN(sizeof(struct rtattr)) + (len))
			// #define RTA_PAYLOAD(rta) ((int)((rta)->rta_len) - RTA_LENGTH(0))
			attrLen := attr.Attr.Len - unix.SizeofRtAttr
			if attrLen == 4 {
				neigh.LLIPAddr = net.IP(attr.Value)
			} else if attrLen == 16 {
				// Can be IPv6 or FireWire HWAddr
				link, err := LinkByIndex(neigh.LinkIndex)
				if err == nil && link.Attrs().EncapType == "tunnel6" {
					neigh.IP = net.IP(attr.Value)
				} else {
					neigh.HardwareAddr = net.HardwareAddr(attr.Value)
				}
			} else {
				neigh.HardwareAddr = net.HardwareAddr(attr.Value)
			}
		case NDA_VLAN:
			neigh.Vlan = int(native.Uint16(attr.Value[0:2]))
		case NDA_VNI:
			neigh.VNI = int(native.Uint32(attr.Value[0:4]))
		}
	}

	return &neigh, nil
}

// NeighSubscribe takes a chan down which notifications will be sent
// when neighbors are added or deleted. Close the 'done' chan to stop subscription.
func NeighSubscribe(ch chan<- NeighUpdate, done <-chan struct{}) error {
	return neighSubscribeAt(netns.None(), netns.None(), ch, done, nil, false)
}

// NeighSubscribeAt works like NeighSubscribe plus it allows the caller
// to choose the network namespace in which to subscribe (ns).
func NeighSubscribeAt(ns netns.NsHandle, ch chan<- NeighUpdate, done <-chan struct{}) error {
	return neighSubscribeAt(ns, netns.None(), ch, done, nil, false)
}

// NeighSubscribeOptions contains a set of options to use with
// NeighSubscribeWithOptions.
type NeighSubscribeOptions struct {
	Namespace     *netns.NsHandle
	ErrorCallback func(error)
	ListExisting  bool
}

// NeighSubscribeWithOptions work like NeighSubscribe but enable to
// provide additional options to modify the behavior. Currently, the
// namespace can be provided as well as an error callback.
func NeighSubscribeWithOptions(ch chan<- NeighUpdate, done <-chan struct{}, options NeighSubscribeOptions) error {
	if options.Namespace == nil {
		none := netns.None()
		options.Namespace = &none
	}
	return neighSubscribeAt(*options.Namespace, netns.None(), ch, done, options.ErrorCallback, options.ListExisting)
}

func neighSubscribeAt(newNs, curNs netns.NsHandle, ch chan<- NeighUpdate, done <-chan struct{}, cberr func(error), listExisting bool) error {
	s, err := nl.SubscribeAt(newNs, curNs, unix.NETLINK_ROUTE, unix.RTNLGRP_NEIGH)
	if err != nil {
		return err
	}
	if done != nil {
		go func() {
			<-done
			s.Close()
		}()
	}
	if listExisting {
		req := pkgHandle.newNetlinkRequest(unix.RTM_GETNEIGH,
			unix.NLM_F_DUMP)
		infmsg := nl.NewIfInfomsg(unix.AF_UNSPEC)
		req.AddData(infmsg)
		if err := s.Send(req); err != nil {
			return err
		}
	}
	go func() {
		defer close(ch)
		for {
			msgs, from, err := s.Receive()
			if err != nil {
				if cberr != nil {
					cberr(err)
				}
				return
			}
			if from.Pid != nl.PidKernel {
				if cberr != nil {
					cberr(fmt.Errorf("Wrong sender portid %d, expected %d", from.Pid, nl.PidKernel))
				}
				continue
			}
			for _, m := range msgs {
				if m.Header.Type == unix.NLMSG_DONE {
					continue
				}
				if m.Header.Type == unix.NLMSG_ERROR {
					native := nl.NativeEndian()
					error := int32(native.Uint32(m.Data[0:4]))
					if error == 0 {
						continue
					}
					if cberr != nil {
						cberr(syscall.Errno(-error))
					}
					return
				}
				neigh, err := NeighDeserialize(m.Data)
				if err != nil {
					if cberr != nil {
						cberr(err)
					}
					return
				}
				ch <- NeighUpdate{Type: m.Header.Type, Neigh: *neigh}
			}
		}
	}()

	return nil
}
