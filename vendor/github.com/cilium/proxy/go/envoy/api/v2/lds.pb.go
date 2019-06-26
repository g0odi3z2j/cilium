// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/lds.proto

package v2

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	listener "github.com/cilium/proxy/go/envoy/api/v2/listener"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Listener_DrainType int32

const (
	// Drain in response to calling /healthcheck/fail admin endpoint (along with the health check
	// filter), listener removal/modification, and hot restart.
	Listener_DEFAULT Listener_DrainType = 0
	// Drain in response to listener removal/modification and hot restart. This setting does not
	// include /healthcheck/fail. This setting may be desirable if Envoy is hosting both ingress
	// and egress listeners.
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

// [#comment:next free field: 16]
type Listener struct {
	// The unique name by which this listener is known. If no name is provided,
	// Envoy will allocate an internal UUID for the listener. If the listener is to be dynamically
	// updated or removed via :ref:`LDS <config_listeners_lds>` a unique name must be provided.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address that the listener should listen on. In general, the address must be unique, though
	// that is governed by the bind rules of the OS. E.g., multiple listeners can listen on port 0 on
	// Linux as the actual port will be allocated by the OS.
	Address *core.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// A list of filter chains to consider for this listener. The
	// :ref:`FilterChain <envoy_api_msg_listener.FilterChain>` with the most specific
	// :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` criteria is used on a
	// connection.
	//
	// Example using SNI for filter chain selection can be found in the
	// :ref:`FAQ entry <faq_how_to_setup_sni>`.
	FilterChains []*listener.FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	// If a connection is redirected using *iptables*, the port on which the proxy
	// receives it might be different from the original destination address. When this flag is set to
	// true, the listener hands off redirected connections to the listener associated with the
	// original destination address. If there is no listener associated with the original destination
	// address, the connection is handled by the listener that receives it. Defaults to false.
	//
	// .. attention::
	//
	//   This field is deprecated. Use :ref:`an original_dst <config_listener_filters_original_dst>`
	//   :ref:`listener filter <envoy_api_field_Listener.listener_filters>` instead.
	//
	//   Note that hand off to another listener is *NOT* performed without this flag. Once
	//   :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` is implemented this flag
	//   will be removed, as filter chain matching can be used to select a filter chain based on the
	//   restored destination address.
	UseOriginalDst *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst,proto3" json:"use_original_dst,omitempty"` // Deprecated: Do not use.
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Listener metadata.
	Metadata *core.Metadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	// Listener filters have the opportunity to manipulate and augment the connection metadata that
	// is used in connection filter chain matching, for example. These filters are run before any in
	// :ref:`filter_chains <envoy_api_field_Listener.filter_chains>`. Order matters as the
	// filters are processed sequentially right after a socket has been accepted by the listener, and
	// before a connection is created.
	ListenerFilters []*listener.ListenerFilter `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	// The timeout to wait for all listener filters to complete operation. If the timeout is reached,
	// the accepted socket is closed without a connection being created. Specify 0 to disable the
	// timeout. If not specified, a default timeout of 15s is used.
	ListenerFiltersTimeout *duration.Duration `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	// Whether the listener should be set as a transparent socket.
	// When this flag is set to true, connections can be redirected to the listener using an
	// *iptables* *TPROXY* target, in which case the original source and destination addresses and
	// ports are preserved on accepted connections. This flag should be used in combination with
	// :ref:`an original_dst <config_listener_filters_original_dst>` :ref:`listener filter
	// <envoy_api_field_Listener.listener_filters>` to mark the connections' local addresses as
	// "restored." This can be used to hand off each redirected connection to another listener
	// associated with the connection's destination address. Direct connections to the socket without
	// using *TPROXY* cannot be distinguished from connections redirected using *TPROXY* and are
	// therefore treated as if they were redirected.
	// When this flag is set to false, the listener's socket is explicitly reset as non-transparent.
	// Setting this flag requires Envoy to run with the *CAP_NET_ADMIN* capability.
	// When this flag is not set (default), the socket is not modified, i.e. the transparent option
	// is neither set nor reset.
	Transparent *wrappers.BoolValue `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	// Whether the listener should set the *IP_FREEBIND* socket option. When this
	// flag is set to true, listeners can be bound to an IP address that is not
	// configured on the system running Envoy. When this flag is set to false, the
	// option *IP_FREEBIND* is disabled on the socket. When this flag is not set
	// (default), the socket is not modified, i.e. the option is neither enabled
	// nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions []*core.SocketOption `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	// Whether the listener should accept TCP Fast Open (TFO) connections.
	// When this flag is set to a value greater than 0, the option TCP_FASTOPEN is enabled on
	// the socket, with a queue length of the specified size
	// (see `details in RFC7413 <https://tools.ietf.org/html/rfc7413#section-5.1>`_).
	// When this flag is set to 0, the option TCP_FASTOPEN is disabled on the socket.
	// When this flag is not set (default), the socket is not modified,
	// i.e. the option is neither enabled nor disabled.
	//
	// On Linux, the net.ipv4.tcp_fastopen kernel parameter must include flag 0x2 to enable
	// TCP_FASTOPEN.
	// See `ip-sysctl.txt <https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt>`_.
	//
	// On macOS, only values of 0, 1, and unset are valid; other values may result in an error.
	// To set the queue length on macOS, set the net.inet.tcp.fastopen_backlog kernel parameter.
	TcpFastOpenQueueLength *wrappers.UInt32Value `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

// Deprecated: Do not use.
func (m *Listener) GetUseOriginalDst() *wrappers.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetListenerFiltersTimeout() *duration.Duration {
	if m != nil {
		return m.ListenerFiltersTimeout
	}
	return nil
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*core.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

// [#not-implemented-hide:]
type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn't
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// [V2-API-DIFF] This is deprecated in v2, all Listeners will bind to their
	// port. An additional filter chain must be created for every original
	// destination port this listener may redirect to in v2, with the original
	// port specified in the FilterChainMatch destination_port field.
	//
	// [#comment:TODO(PiotrSikora): Remove this once verified that we no longer need it.]
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
}

func init() { proto.RegisterFile("envoy/api/v2/lds.proto", fileDescriptor_34e2cd84a105bcd1) }

var fileDescriptor_34e2cd84a105bcd1 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x33, 0x8e, 0xdb, 0x38, 0x63, 0xc7, 0xb6, 0x46, 0x28, 0x5d, 0x4c, 0x88, 0x8d, 0x01,
	0xc9, 0x70, 0x58, 0x53, 0x57, 0x02, 0xa9, 0xaa, 0x84, 0xba, 0x35, 0x56, 0x8b, 0x5c, 0x1c, 0x36,
	0x69, 0x68, 0x4e, 0xab, 0xf1, 0xee, 0xb3, 0x33, 0x62, 0x3d, 0x33, 0x9d, 0x99, 0x35, 0xf8, 0xca,
	0x09, 0x71, 0x84, 0x0b, 0x1f, 0x81, 0xcf, 0xc0, 0x89, 0x23, 0x77, 0xee, 0x20, 0x21, 0x2e, 0x88,
	0xcf, 0x80, 0x84, 0x76, 0xbc, 0x6b, 0xec, 0x26, 0xa6, 0x1c, 0x7a, 0x7b, 0x33, 0xef, 0xff, 0x7e,
	0x7e, 0xfb, 0x7f, 0x6f, 0xd7, 0xf8, 0x10, 0xf8, 0x5c, 0x2c, 0xba, 0x54, 0xb2, 0xee, 0xbc, 0xd7,
	0x8d, 0x23, 0xed, 0x4a, 0x25, 0x8c, 0x20, 0x15, 0x7b, 0xef, 0x52, 0xc9, 0xdc, 0x79, 0xaf, 0xd1,
	0xdc, 0x50, 0x85, 0x42, 0x41, 0x97, 0x46, 0x91, 0x02, 0x9d, 0xc9, 0x1b, 0x47, 0x57, 0x05, 0x63,
	0xaa, 0xe1, 0xda, 0x6c, 0xc4, 0x74, 0x28, 0xe6, 0xa0, 0x16, 0x59, 0xf6, 0xad, 0xcd, 0x16, 0x98,
	0x36, 0xc0, 0x41, 0xad, 0x82, 0x9c, 0x31, 0x15, 0x62, 0x1a, 0x83, 0x95, 0x51, 0xce, 0x85, 0xa1,
	0x86, 0x09, 0x9e, 0xff, 0xfe, 0x71, 0x96, 0xb5, 0xa7, 0x71, 0x32, 0xe9, 0x46, 0x89, 0xb2, 0x82,
	0x6d, 0xf9, 0x2f, 0x14, 0x95, 0x12, 0x54, 0x5e, 0x7f, 0x6b, 0x4e, 0x63, 0x16, 0x51, 0x03, 0xdd,
	0x3c, 0xc8, 0x12, 0xaf, 0x4c, 0xc5, 0x54, 0xd8, 0xb0, 0x9b, 0x46, 0xcb, 0xdb, 0xf6, 0xdf, 0x25,
	0x5c, 0x1a, 0x66, 0xfd, 0x11, 0x82, 0x8b, 0x9c, 0xce, 0xc0, 0x41, 0x2d, 0xd4, 0xd9, 0xf7, 0x6d,
	0x4c, 0xfa, 0x78, 0x2f, 0x33, 0xc8, 0x29, 0xb4, 0x50, 0xa7, 0xdc, 0x6b, 0xb8, 0xeb, 0x86, 0xba,
	0xa9, 0x43, 0xee, 0xfd, 0xa5, 0xc2, 0xab, 0xfe, 0xfc, 0x6b, 0x73, 0xe7, 0xc7, 0x3f, 0x7f, 0xda,
	0xbd, 0xf1, 0x0d, 0x2a, 0xd4, 0x91, 0x9f, 0x97, 0x92, 0xcf, 0xf0, 0xc1, 0x84, 0xc5, 0x06, 0x54,
	0x10, 0x5e, 0x52, 0xc6, 0xb5, 0xb3, 0xdb, 0xda, 0xed, 0x94, 0x7b, 0xed, 0x4d, 0xd6, 0xca, 0xa8,
	0x81, 0xd5, 0x3e, 0x48, 0xa5, 0x6b, 0xcc, 0x6f, 0x51, 0xa1, 0x84, 0xfc, 0xca, 0xe4, 0xdf, 0xa4,
	0x26, 0x0f, 0x71, 0x3d, 0xd1, 0x10, 0x08, 0xc5, 0xa6, 0x8c, 0xd3, 0x38, 0x88, 0xb4, 0x71, 0x8a,
	0x59, 0x9f, 0x4b, 0xa7, 0xdc, 0xdc, 0x29, 0xd7, 0x13, 0x22, 0x3e, 0xa7, 0x71, 0x02, 0x5e, 0xc1,
	0x41, 0x7e, 0x35, 0xd1, 0x30, 0xca, 0xca, 0xfa, 0xda, 0x90, 0x09, 0x7e, 0x43, 0xa6, 0xfd, 0x09,
	0xce, 0x21, 0x4c, 0x0d, 0x0f, 0xc6, 0xc9, 0x64, 0x02, 0x2a, 0x88, 0xd9, 0x8c, 0x99, 0x60, 0xbc,
	0x30, 0xa0, 0x9d, 0x1b, 0x16, 0x7d, 0x74, 0x05, 0xfd, 0xe4, 0x11, 0x37, 0x77, 0x7a, 0x16, 0xee,
	0xbf, 0x2e, 0x41, 0x3d, 0x58, 0x51, 0x3c, 0x0b, 0x19, 0xa6, 0x0c, 0x2f, 0x45, 0x90, 0x0f, 0x70,
	0x69, 0x06, 0x86, 0x46, 0xd4, 0x50, 0xe7, 0xa6, 0xc5, 0xbd, 0x76, 0x8d, 0xa3, 0x8f, 0x33, 0x89,
	0xbf, 0x12, 0x93, 0x87, 0xf8, 0x20, 0x02, 0xa9, 0x20, 0xa4, 0x06, 0xa2, 0x60, 0x7e, 0xdb, 0xd9,
	0xb3, 0xd5, 0x6f, 0x6e, 0x56, 0xe7, 0xc3, 0x74, 0xfb, 0x2b, 0xed, 0xf9, 0x6d, 0xbf, 0x12, 0xad,
	0x9d, 0xc8, 0x87, 0x18, 0x47, 0x8a, 0x32, 0x1e, 0x98, 0x85, 0x04, 0xa7, 0xd4, 0x42, 0x9d, 0x6a,
	0xaf, 0xb5, 0x0d, 0x93, 0x0a, 0xcf, 0x16, 0x12, 0xfc, 0xfd, 0x28, 0x0f, 0xc9, 0x39, 0xae, 0xe7,
	0xb3, 0x0a, 0x96, 0xe3, 0xd0, 0xce, 0xbe, 0x9d, 0xe8, 0xdb, 0x5b, 0x26, 0x9a, 0xf3, 0x96, 0x93,
	0xf5, 0x8a, 0xe9, 0x50, 0xfd, 0x5a, 0xbc, 0x71, 0xab, 0xc9, 0x05, 0x76, 0x9e, 0xe7, 0x06, 0x86,
	0xcd, 0x40, 0x24, 0xc6, 0xa9, 0xd9, 0xa7, 0x7d, 0xf5, 0x8a, 0xf5, 0xfd, 0xec, 0xfd, 0xf0, 0x8a,
	0xdf, 0xff, 0xd6, 0x44, 0xfe, 0xe1, 0x73, 0xcc, 0xb3, 0x65, 0x39, 0xb9, 0x87, 0xcb, 0x46, 0x51,
	0xae, 0x25, 0x55, 0xc0, 0x8d, 0x83, 0x5f, 0xb4, 0x23, 0xfe, 0xba, 0x9c, 0xbc, 0x8f, 0x4b, 0x13,
	0x05, 0x30, 0x66, 0x3c, 0x72, 0xca, 0x2f, 0x2c, 0x5d, 0x69, 0xc9, 0x00, 0x57, 0xb5, 0x08, 0x3f,
	0x07, 0x13, 0x08, 0x69, 0xdf, 0x72, 0xe7, 0xc0, 0xda, 0xd4, 0xbc, 0x66, 0xe4, 0xa7, 0x56, 0x38,
	0xb2, 0x3a, 0xff, 0x40, 0xaf, 0x9d, 0x34, 0x79, 0x8a, 0x1b, 0x26, 0x94, 0xc1, 0x84, 0xea, 0x94,
	0x04, 0x3c, 0x78, 0x96, 0x40, 0x02, 0x41, 0x0c, 0x7c, 0x6a, 0x2e, 0x9d, 0xca, 0xff, 0xd8, 0xca,
	0x43, 0x13, 0xca, 0x01, 0xd5, 0x66, 0x24, 0x81, 0x7f, 0x9a, 0x16, 0x0f, 0x6d, 0x6d, 0x63, 0x88,
	0x2b, 0xeb, 0x9b, 0x42, 0xee, 0xe1, 0x4a, 0xda, 0x79, 0x60, 0x44, 0x20, 0x85, 0x32, 0xf6, 0x5b,
	0xf0, 0xdf, 0x4f, 0x8b, 0x53, 0xfd, 0x99, 0x38, 0x11, 0xca, 0xb4, 0xdf, 0xc1, 0xfb, 0xab, 0x85,
	0x21, 0x65, 0xbc, 0xd7, 0xff, 0x68, 0x70, 0xff, 0xc9, 0xf0, 0xac, 0xbe, 0x43, 0x6a, 0xb8, 0xfc,
	0x78, 0xd4, 0x7f, 0x34, 0xb8, 0x08, 0x46, 0x9f, 0x0c, 0x2f, 0xea, 0xe8, 0xe3, 0x62, 0xa9, 0x5a,
	0xaf, 0xf5, 0xfe, 0x42, 0xd8, 0xc9, 0x77, 0xa3, 0x9f, 0x7f, 0x4e, 0x4f, 0x41, 0xcd, 0x59, 0x08,
	0xe4, 0x29, 0xae, 0x9d, 0x1a, 0x05, 0x74, 0x96, 0x2b, 0x34, 0x39, 0xde, 0x34, 0x6e, 0x55, 0xe2,
	0xc3, 0xb3, 0x04, 0xb4, 0x69, 0x34, 0xb7, 0xe6, 0xb5, 0x14, 0x5c, 0x43, 0x7b, 0xa7, 0x83, 0xde,
	0x43, 0x24, 0xc1, 0xd5, 0x01, 0x98, 0xf0, 0xf2, 0x25, 0x82, 0xdb, 0x5f, 0xfd, 0xf2, 0xc7, 0x77,
	0x85, 0xa3, 0xf6, 0xad, 0x8d, 0x7f, 0x86, 0xbb, 0xf9, 0x3a, 0xea, 0xbb, 0xe8, 0x5d, 0xaf, 0xfb,
	0xc3, 0xef, 0xc7, 0x08, 0x37, 0x98, 0x58, 0xc2, 0xa4, 0x12, 0x5f, 0x2e, 0x36, 0xb8, 0x5e, 0x69,
	0x18, 0xe9, 0x93, 0xd4, 0xe3, 0x13, 0xf4, 0x35, 0x42, 0xe3, 0x9b, 0xd6, 0xef, 0x3b, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x82, 0x83, 0xd1, 0xdd, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (*UnimplementedListenerDiscoveryServiceServer) StreamListeners(srv ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) FetchListeners(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/lds.proto",
}
