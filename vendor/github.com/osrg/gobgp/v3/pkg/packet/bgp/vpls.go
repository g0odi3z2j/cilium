package bgp

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// VPLSNLRI represents an NLRI for VPLS, as defined in [RFC 4761, section 3.2.2].
//
//	Path Attribute - MP_REACH_NLRI
//	  Flags: 0x90, Optional, Extended-Length, Non-transitive, Complete
//	  Type Code: MP_REACH_NLRI (14)
//	  Length: 28
//	  Address family identifier (AFI): Layer-2 VPN (25)
//	  Subsequent address family identifier (SAFI): VPLS (65)
//	  Next hop: 192.0.2.7
//	      IPv4 Address: 192.0.2.7
//	  Number of Subnetwork points of attachment (SNPA): 0
//	  Network Layer Reachability Information (NLRI)
//	      Length: 17
//	      RD: 65017:104
//	      CE-ID: 1
//	      Label Block Offset: 1
//	      Label Block Size: 8
//	      Label Block Base: 800000 (bottom)
//
// [RFC 4761, section 3.2.2]: https://www.rfc-editor.org/rfc/rfc4761.html#section-3.2.2.
type VPLSNLRI struct {
	PrefixDefault
	VEID          uint16
	VEBlockOffset uint16
	VEBlockSize   uint16
	LabelBase     MPLSLabelStack

	rd RouteDistinguisherInterface
}

func (n *VPLSNLRI) DecodeFromBytes(data []byte, options ...*MarshallingOption) error {
	if len(data) < 19 {
		return NewMessageError(BGP_ERROR_UPDATE_MESSAGE_ERROR, BGP_ERROR_SUB_MALFORMED_ATTRIBUTE_LIST, nil, "Not all VPLS NLRI bytes available")
	}
	length := int(binary.BigEndian.Uint16(data[0:2]))
	if len(data) < length+2 {
		return NewMessageError(BGP_ERROR_UPDATE_MESSAGE_ERROR, BGP_ERROR_SUB_MALFORMED_ATTRIBUTE_LIST, nil, "Not all VPLS NLRI bytes available")
	}
	n.rd = GetRouteDistinguisher(data[2:10])
	n.VEID = binary.BigEndian.Uint16(data[10:12])
	n.VEBlockOffset = binary.BigEndian.Uint16(data[12:14])
	n.VEBlockSize = binary.BigEndian.Uint16(data[14:16])
	return n.LabelBase.DecodeFromBytes(data[16:19], options...)
}

func (n *VPLSNLRI) Serialize(options ...*MarshallingOption) ([]byte, error) {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint16(buf[0:2], 17)
	rdbuf, err := n.rd.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[2:10], rdbuf[:8])
	binary.BigEndian.PutUint16(buf[10:12], n.VEID)
	binary.BigEndian.PutUint16(buf[12:14], n.VEBlockOffset)
	binary.BigEndian.PutUint16(buf[14:16], n.VEBlockSize)
	tbuf, err := n.LabelBase.Serialize(options...)
	if err != nil {
		return nil, err
	}
	return append(buf, tbuf...), nil
}

func (n *VPLSNLRI) AFI() uint16 {
	return AFI_L2VPN
}

func (n *VPLSNLRI) SAFI() uint8 {
	return SAFI_VPLS
}

func (n *VPLSNLRI) Len(options ...*MarshallingOption) int {
	return 19
}

func (n *VPLSNLRI) String() string {
	return fmt.Sprintf("%s:%d:%d (Block Size: %d, Label Base: %d)", n.rd, n.VEID, n.VEBlockOffset, n.VEBlockSize, n.LabelBase)
}

func (n *VPLSNLRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RD            RouteDistinguisherInterface `json:"rd"`
		VEID          uint16                      `json:"id"`
		VEBlockOffset uint16                      `json:"blockoffset"`
		VEBlockSize   uint16                      `json:"blocksize"`
		LabelBase     MPLSLabelStack              `json:"base"`
	}{
		RD:            n.rd,
		VEID:          n.VEID,
		VEBlockOffset: n.VEBlockOffset,
		VEBlockSize:   n.VEBlockSize,
		LabelBase:     n.LabelBase,
	})
}

func (n *VPLSNLRI) RD() RouteDistinguisherInterface {
	return n.rd
}

func (l *VPLSNLRI) Flat() map[string]string {
	return map[string]string{}
}

func NewVPLSNLRI(rd RouteDistinguisherInterface, id uint16, blockOffset uint16, blockSize uint16, labelBase MPLSLabelStack) *VPLSNLRI {
	return &VPLSNLRI{
		rd:            rd,
		VEID:          id,
		VEBlockOffset: blockOffset,
		VEBlockSize:   blockSize,
		LabelBase:     labelBase,
	}
}

// VPLSExtended repsents BGP VPLS Extended Community as described in [RFC 4761, section 3.2.4].
//
//	Path Attribute - EXTENDED_COMMUNITIES
//	  Flags: 0xc0, Optional, Transitive, Complete
//	  Type Code: EXTENDED_COMMUNITIES (16)
//	  Length: 16
//	  Carried extended communities: (2 communities)
//	      Route Target: 65017:104 [Transitive 2-Octet AS-Specific]
//	      Layer2 Info: [Generic Transitive Experimental Use]
//	          Type: Generic Transitive Experimental Use (0x80)
//	          Subtype (Experimental): Layer2 Info (0x0a)
//	          Encaps Type: VPLS (19)
//	          Control Flags: 0x00
//	          Layer-2 MTU: 0
//
// [RFC 4761, section 3.2.4]: https://www.rfc-editor.org/rfc/rfc4761.html#section-3.2.4
type VPLSExtended struct {
	SubType      ExtendedCommunityAttrSubType
	ControlFlags uint8
	MTU          uint16
}

func (e *VPLSExtended) Serialize() ([]byte, error) {
	buf := make([]byte, 8)
	buf[0] = byte(EC_TYPE_GENERIC_TRANSITIVE_EXPERIMENTAL)
	buf[1] = byte(EC_SUBTYPE_L2_INFO)
	buf[2] = byte(LAYER2ENCAPSULATION_TYPE_VPLS)
	buf[3] = byte(e.ControlFlags)
	binary.BigEndian.PutUint16(buf[4:6], e.MTU)
	// 6-8: reserved, but Juniper says this is "site preference"
	return buf, nil
}

func (e *VPLSExtended) String() string {
	return fmt.Sprintf("encaps: VPLS, control flags:0x%x, mtu: %d", e.ControlFlags, e.MTU)
}

func (e *VPLSExtended) MarshalJSON() ([]byte, error) {
	t, s := e.GetTypes()
	return json.Marshal(struct {
		Type    ExtendedCommunityAttrType    `json:"type"`
		Subtype ExtendedCommunityAttrSubType `json:"subtype"`
		Value   string                       `json:"value"`
	}{
		Type:    t,
		Subtype: s,
		Value:   e.String(),
	})
}

func (e *VPLSExtended) GetTypes() (ExtendedCommunityAttrType, ExtendedCommunityAttrSubType) {
	return EC_TYPE_GENERIC_TRANSITIVE_EXPERIMENTAL, EC_SUBTYPE_L2_INFO
}

func (e *VPLSExtended) Flat() map[string]string {
	return map[string]string{}
}

func NewVPLSExtended(flags uint8, mtu uint16) *VPLSExtended {
	return &VPLSExtended{
		SubType:      EC_SUBTYPE_L2_INFO,
		ControlFlags: flags,
		MTU:          mtu,
	}
}
