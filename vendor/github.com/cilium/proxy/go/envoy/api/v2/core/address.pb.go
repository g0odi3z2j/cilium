// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/address.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}

func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{1, 0}
}

type Pipe struct {
	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{0}
}

func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// [#next-free-field: 7]
type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.api.v2.core.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
	// to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
	// It is possible to distinguish a Listener address via the prefix/suffix matching
	// in :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>`.] When used
	// within an upstream :ref:`BindConfig <envoy_api_msg_core.BindConfig>`, the address
	// controls the source address of outbound connections. For :ref:`clusters
	// <envoy_api_msg_Cluster>`, the cluster type determines whether the
	// address must be an IP (*STATIC* or *EDS* clusters) or a hostname resolved by DNS
	// (*STRICT_DNS* or *LOGICAL_DNS* clusters). Address resolution can be customized
	// via :ref:`resolver_name <envoy_api_field_core.SocketAddress.resolver_name>`.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the custom resolver. This must have been registered with Envoy. If
	// this is empty, a context dependent default applies. If the address is a concrete
	// IP address, no resolution will occur. If address is a hostname this
	// should be set for resolution other than DNS. Specifying a custom resolver with
	// *STRICT_DNS* or *LOGICAL_DNS* will generate an error at runtime.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibility
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat           bool     `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{1}
}

func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

type TcpKeepalive struct {
	// Maximum number of keepalive probes to send without response before deciding
	// the connection is dead. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 9.)
	KeepaliveProbes *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes
	// start being sent. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 7200s (ie 2 hours.)
	KeepaliveTime *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. Default is to use the OS
	// level configuration (unless overridden, Linux defaults to 75s.)
	KeepaliveInterval    *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{2}
}

func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (m *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(m, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	// The address to bind to when creating a socket.
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Whether to set the *IP_FREEBIND* option when creating the socket. When this
	// flag is set to true, allows the :ref:`source_address
	// <envoy_api_field_UpstreamBindConfig.source_address>` to be an IP address
	// that is not configured on the system running Envoy. When this flag is set
	// to false, the option *IP_FREEBIND* is disabled on the socket. When this
	// flag is not set (default), the socket is not modified, i.e. the option is
	// neither enabled nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions        []*SocketOption `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{3}
}

func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (m *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(m, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

// CidrRange specifies an IP Address and a prefix length to construct
// the subnet mask for a `CIDR <https://tools.ietf.org/html/rfc4632>`_ range.
type CidrRange struct {
	// IPv4 or IPv6 address, e.g. ``192.0.0.0`` or ``2001:db8::``.
	AddressPrefix string `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	// Length of prefix, e.g. 0, 32.
	PrefixLen            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{5}
}

func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (m *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(m, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.api.v2.core.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.core.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.core.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.api.v2.core.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.core.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.core.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.core.CidrRange")
}

func init() { proto.RegisterFile("envoy/api/v2/core/address.proto", fileDescriptor_6906417f87bcce55) }

var fileDescriptor_6906417f87bcce55 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x26, 0x69, 0x93, 0x4c, 0x9b, 0x90, 0xae, 0x84, 0x6a, 0x85, 0x8a, 0x84, 0xf4, 0x12,
	0x2a, 0xe1, 0x88, 0x14, 0x71, 0xaf, 0xc3, 0x47, 0xab, 0x22, 0x6a, 0x4c, 0xcb, 0xd5, 0xda, 0x24,
	0x93, 0xb0, 0xaa, 0xe3, 0x5d, 0xad, 0x5d, 0xd3, 0xde, 0x50, 0x8f, 0x5c, 0xf9, 0x59, 0xfc, 0x08,
	0xce, 0xfc, 0x03, 0xd4, 0x4b, 0xd1, 0xae, 0x63, 0x57, 0x25, 0xa0, 0xc2, 0x6d, 0x77, 0xe6, 0xbd,
	0xb7, 0x6f, 0x3e, 0x16, 0xda, 0x18, 0x26, 0xe2, 0xa2, 0xcf, 0x24, 0xef, 0x27, 0x83, 0xfe, 0x58,
	0x28, 0xec, 0xb3, 0xc9, 0x44, 0x61, 0x14, 0xd9, 0x52, 0x89, 0x58, 0xd0, 0x0d, 0x03, 0xb0, 0x99,
	0xe4, 0x76, 0x32, 0xb0, 0x35, 0xa0, 0xb5, 0xb5, 0xcc, 0x19, 0xb1, 0x08, 0x53, 0x42, 0xeb, 0xe1,
	0x4c, 0x88, 0x59, 0x80, 0x7d, 0x73, 0x1b, 0x9d, 0x4d, 0xfb, 0x9f, 0x14, 0x93, 0x12, 0xd5, 0x42,
	0xb0, 0xb5, 0x99, 0xb0, 0x80, 0x4f, 0x58, 0x8c, 0xfd, 0xec, 0x90, 0x26, 0xba, 0xdb, 0x50, 0x76,
	0xb9, 0x44, 0xfa, 0x00, 0xca, 0x92, 0xc5, 0x1f, 0x2d, 0xd2, 0x21, 0xbd, 0x9a, 0x53, 0xb9, 0x72,
	0xca, 0xaa, 0xd8, 0x21, 0x9e, 0x09, 0x76, 0xbf, 0x15, 0xa1, 0xfe, 0x5e, 0x8c, 0x4f, 0x31, 0xde,
	0x4b, 0x6d, 0xd2, 0x23, 0xa8, 0x1a, 0xfe, 0x58, 0x04, 0x86, 0xd2, 0x18, 0x3c, 0xb6, 0x97, 0x3c,
	0xdb, 0xb7, 0x38, 0xb6, 0xbb, 0x20, 0x38, 0xd5, 0x2b, 0x67, 0xe5, 0x92, 0x14, 0x9b, 0xc4, 0xcb,
	0x45, 0xe8, 0x23, 0xa8, 0x2c, 0x5a, 0x60, 0x15, 0x6f, 0x5b, 0xc8, 0xe2, 0x74, 0x07, 0x40, 0x0a,
	0x15, 0xfb, 0x09, 0x0b, 0xce, 0xd0, 0x2a, 0x75, 0x48, 0xaf, 0xee, 0xd4, 0xae, 0x9c, 0xd5, 0x9d,
	0xb2, 0x75, 0x7d, 0x5d, 0xda, 0x2f, 0x78, 0x35, 0x9d, 0xfe, 0xa0, 0xb3, 0xb4, 0x0d, 0x10, 0xb2,
	0x39, 0x4e, 0x7c, 0x1d, 0xb2, 0xca, 0x5a, 0x51, 0x03, 0x4c, 0xcc, 0x15, 0x2a, 0xa6, 0xdb, 0x50,
	0x57, 0x18, 0x89, 0x20, 0x41, 0xe5, 0xeb, 0xa8, 0xb5, 0xa2, 0x31, 0xde, 0x7a, 0x16, 0x7c, 0xcb,
	0xe6, 0x5a, 0x65, 0x8d, 0xcb, 0xe4, 0x99, 0x3f, 0x16, 0x73, 0xc9, 0x62, 0x6b, 0xb5, 0x43, 0x7a,
	0x55, 0x0f, 0x74, 0x68, 0x68, 0x22, 0xdd, 0x2d, 0xa8, 0x66, 0x55, 0xd1, 0x0a, 0x94, 0x8e, 0x87,
	0x6e, 0xb3, 0xa0, 0x0f, 0x27, 0x2f, 0xdc, 0x26, 0x71, 0xee, 0x43, 0xc3, 0x18, 0x8e, 0x24, 0x8e,
	0xf9, 0x94, 0xa3, 0xa2, 0xa5, 0x9f, 0x0e, 0xe9, 0xfe, 0x20, 0xb0, 0x7e, 0x3c, 0x96, 0x87, 0x88,
	0x92, 0x05, 0x3c, 0x41, 0xfa, 0x1a, 0x9a, 0xa7, 0xd9, 0xc5, 0x97, 0x4a, 0x8c, 0x30, 0x32, 0x4d,
	0x5d, 0x1b, 0x6c, 0xd9, 0xe9, 0x5c, 0xed, 0x6c, 0xae, 0xf6, 0xc9, 0x41, 0x18, 0xef, 0x0e, 0x4c,
	0x91, 0xde, 0xbd, 0x9c, 0xe5, 0x1a, 0x12, 0x1d, 0x42, 0xe3, 0x46, 0x28, 0xe6, 0x73, 0x34, 0xbd,
	0xbc, 0x4b, 0xa6, 0x9e, 0x73, 0x8e, 0xf9, 0x1c, 0xe9, 0x21, 0xd0, 0x1b, 0x11, 0x1e, 0xc6, 0xa8,
	0x12, 0x16, 0x98, 0x76, 0xdf, 0x25, 0xb4, 0x91, 0xf3, 0x0e, 0x16, 0xb4, 0xee, 0x77, 0x02, 0xe0,
	0xf0, 0x70, 0x32, 0x14, 0xe1, 0x94, 0xcf, 0xe8, 0x3b, 0x68, 0x44, 0xe2, 0x4c, 0x8d, 0xd1, 0xcf,
	0x86, 0x9d, 0xd6, 0xd9, 0xb9, 0x6b, 0x79, 0xcc, 0xce, 0x7c, 0x31, 0x3b, 0x53, 0x4f, 0x15, 0xb2,
	0x4d, 0x7c, 0x0e, 0xd5, 0xa9, 0x42, 0x1c, 0xf1, 0x70, 0xb2, 0xa8, 0xb6, 0xb5, 0x64, 0xd2, 0x11,
	0x22, 0x48, 0x2d, 0xe6, 0x58, 0xfa, 0x4a, 0x5b, 0xd1, 0x2f, 0xf8, 0x42, 0xc6, 0x5c, 0x84, 0x91,
	0x55, 0xea, 0x94, 0x7a, 0x6b, 0x83, 0xf6, 0x5f, 0xad, 0x1c, 0x19, 0x9c, 0x7e, 0xff, 0xe6, 0x16,
	0x75, 0xbf, 0x12, 0xa8, 0x64, 0x5e, 0x0e, 0x72, 0xcd, 0xff, 0x2c, 0x6f, 0xbf, 0x90, 0xc9, 0x66,
	0x52, 0x4f, 0xa0, 0x2c, 0xb9, 0xcc, 0x06, 0xb8, 0xf9, 0x07, 0x01, 0xfd, 0x6d, 0xf7, 0x0b, 0x9e,
	0x81, 0x39, 0x8d, 0xfc, 0xfb, 0xa4, 0x3b, 0x76, 0x49, 0xa0, 0x36, 0xe4, 0x13, 0xe5, 0xb1, 0x70,
	0x86, 0xd4, 0x86, 0xc6, 0x22, 0xeb, 0x4b, 0x85, 0x53, 0x7e, 0xfe, 0xfb, 0x37, 0xaf, 0x2f, 0xd2,
	0xae, 0xc9, 0xd2, 0x97, 0x00, 0x29, 0xce, 0x0f, 0x30, 0xfc, 0x97, 0x1d, 0x32, 0xe3, 0xd9, 0x29,
	0x59, 0x9f, 0x89, 0x57, 0x4b, 0x99, 0x6f, 0x30, 0x74, 0x9e, 0x42, 0x9b, 0x8b, 0xd4, 0xb9, 0x54,
	0xe2, 0xfc, 0x62, 0xb9, 0x08, 0x67, 0x7d, 0x2f, 0x7b, 0x58, 0xc4, 0xc2, 0x25, 0xa3, 0x55, 0xa3,
	0xbe, 0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x8b, 0x95, 0x1d, 0x22, 0x05, 0x00, 0x00,
}
