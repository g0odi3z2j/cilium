// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/listener.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/accesslog/v3"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	return fileDescriptor_1700162847fce94e, []int{0, 0}
}

// [#next-free-field: 23]
type Listener struct {
	// The unique name by which this listener is known. If no name is provided,
	// Envoy will allocate an internal UUID for the listener. If the listener is to be dynamically
	// updated or removed via :ref:`LDS <config_listeners_lds>` a unique name must be provided.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address that the listener should listen on. In general, the address must be unique, though
	// that is governed by the bind rules of the OS. E.g., multiple listeners can listen on port 0 on
	// Linux as the actual port will be allocated by the OS.
	Address *v3.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// A list of filter chains to consider for this listener. The
	// :ref:`FilterChain <envoy_api_msg_config.listener.v3.FilterChain>` with the most specific
	// :ref:`FilterChainMatch <envoy_api_msg_config.listener.v3.FilterChainMatch>` criteria is used on a
	// connection.
	//
	// Example using SNI for filter chain selection can be found in the
	// :ref:`FAQ entry <faq_how_to_setup_sni>`.
	FilterChains []*FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Listener metadata.
	Metadata *v3.Metadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.config.listener.v3.Listener_DrainType" json:"drain_type,omitempty"`
	// Listener filters have the opportunity to manipulate and augment the connection metadata that
	// is used in connection filter chain matching, for example. These filters are run before any in
	// :ref:`filter_chains <envoy_api_field_config.listener.v3.Listener.filter_chains>`. Order matters as the
	// filters are processed sequentially right after a socket has been accepted by the listener, and
	// before a connection is created.
	// UDP Listener filters can be specified when the protocol in the listener socket address in
	// :ref:`protocol <envoy_api_field_config.core.v3.SocketAddress.protocol>` is :ref:`UDP
	// <envoy_api_enum_value_config.core.v3.SocketAddress.Protocol.UDP>`.
	// UDP listeners currently support a single filter.
	ListenerFilters []*ListenerFilter `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	// The timeout to wait for all listener filters to complete operation. If the timeout is reached,
	// the accepted socket is closed without a connection being created unless
	// `continue_on_listener_filters_timeout` is set to true. Specify 0 to disable the
	// timeout. If not specified, a default timeout of 15s is used.
	ListenerFiltersTimeout *duration.Duration `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	// Whether a connection should be created when listener filters timeout. Default is false.
	//
	// .. attention::
	//
	//   Some listener filters, such as :ref:`Proxy Protocol filter
	//   <config_listener_filters_proxy_protocol>`, should not be used with this option. It will cause
	//   unexpected behavior when a connection is created.
	ContinueOnListenerFiltersTimeout bool `protobuf:"varint,17,opt,name=continue_on_listener_filters_timeout,json=continueOnListenerFiltersTimeout,proto3" json:"continue_on_listener_filters_timeout,omitempty"`
	// Whether the listener should be set as a transparent socket.
	// When this flag is set to true, connections can be redirected to the listener using an
	// *iptables* *TPROXY* target, in which case the original source and destination addresses and
	// ports are preserved on accepted connections. This flag should be used in combination with
	// :ref:`an original_dst <config_listener_filters_original_dst>` :ref:`listener filter
	// <envoy_api_field_config.listener.v3.Listener.listener_filters>` to mark the connections' local addresses as
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
	SocketOptions []*v3.SocketOption `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
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
	// Specifies the intended direction of the traffic relative to the local Envoy.
	TrafficDirection v3.TrafficDirection `protobuf:"varint,16,opt,name=traffic_direction,json=trafficDirection,proto3,enum=envoy.config.core.v3.TrafficDirection" json:"traffic_direction,omitempty"`
	// If the protocol in the listener socket address in :ref:`protocol
	// <envoy_api_field_config.core.v3.SocketAddress.protocol>` is :ref:`UDP
	// <envoy_api_enum_value_config.core.v3.SocketAddress.Protocol.UDP>`, this field specifies the actual udp
	// listener to create, i.e. :ref:`udp_listener_name
	// <envoy_api_field_config.listener.v3.UdpListenerConfig.udp_listener_name>` = "raw_udp_listener" for
	// creating a packet-oriented UDP listener. If not present, treat it as "raw_udp_listener".
	UdpListenerConfig *UdpListenerConfig `protobuf:"bytes,18,opt,name=udp_listener_config,json=udpListenerConfig,proto3" json:"udp_listener_config,omitempty"`
	// Used to represent an API listener, which is used in non-proxy clients. The type of API
	// exposed to the non-proxy application depends on the type of API listener.
	// When this field is set, no other field except for :ref:`name<envoy_api_field_config.listener.v3.Listener.name>`
	// should be set.
	//
	// .. note::
	//
	//  Currently only one ApiListener can be installed; and it can only be done via bootstrap config,
	//  not LDS.
	//
	// [#next-major-version: In the v3 API, instead of this messy approach where the socket
	// listener fields are directly in the top-level Listener message and the API listener types
	// are in the ApiListener message, the socket listener messages should be in their own message,
	// and the top-level Listener should essentially be a oneof that selects between the
	// socket listener and the various types of API listener. That way, a given Listener message
	// can structurally only contain the fields of the relevant type.]
	ApiListener *ApiListener `protobuf:"bytes,19,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
	// The listener's connection balancer configuration, currently only applicable to TCP listeners.
	// If no configuration is specified, Envoy will not attempt to balance active connections between
	// worker threads.
	ConnectionBalanceConfig *Listener_ConnectionBalanceConfig `protobuf:"bytes,20,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	// When this flag is set to true, listeners set the *SO_REUSEPORT* socket option and
	// create one socket for each worker thread. This makes inbound connections
	// distribute among worker threads roughly evenly in cases where there are a high number
	// of connections. When this flag is set to false, all worker threads share one socket.
	//
	// Before Linux v4.19-rc1, new TCP connections may be rejected during hot restart
	// (see `3rd paragraph in 'soreuseport' commit message
	// <https://github.com/torvalds/linux/commit/c617f398edd4db2b8567a28e89>`_).
	// This issue was fixed by `tcp: Avoid TCP syncookie rejected by SO_REUSEPORT socket
	// <https://github.com/torvalds/linux/commit/40a1227ea845a37ab197dd1caffb60b047fa36b1>`_.
	ReusePort bool `protobuf:"varint,21,opt,name=reuse_port,json=reusePort,proto3" json:"reuse_port,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by this listener.
	AccessLog            []*v31.AccessLog `protobuf:"bytes,22,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_1700162847fce94e, []int{0}
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

func (m *Listener) GetAddress() *v3.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *v3.Metadata {
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

func (m *Listener) GetListenerFilters() []*ListenerFilter {
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

func (m *Listener) GetContinueOnListenerFiltersTimeout() bool {
	if m != nil {
		return m.ContinueOnListenerFiltersTimeout
	}
	return false
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

func (m *Listener) GetSocketOptions() []*v3.SocketOption {
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

func (m *Listener) GetTrafficDirection() v3.TrafficDirection {
	if m != nil {
		return m.TrafficDirection
	}
	return v3.TrafficDirection_UNSPECIFIED
}

func (m *Listener) GetUdpListenerConfig() *UdpListenerConfig {
	if m != nil {
		return m.UdpListenerConfig
	}
	return nil
}

func (m *Listener) GetApiListener() *ApiListener {
	if m != nil {
		return m.ApiListener
	}
	return nil
}

func (m *Listener) GetConnectionBalanceConfig() *Listener_ConnectionBalanceConfig {
	if m != nil {
		return m.ConnectionBalanceConfig
	}
	return nil
}

func (m *Listener) GetReusePort() bool {
	if m != nil {
		return m.ReusePort
	}
	return false
}

func (m *Listener) GetAccessLog() []*v31.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// [#not-implemented-hide:]
type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn't
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// This is deprecated in v2, all Listeners will bind to their port. An
	// additional filter chain must be created for every original destination
	// port this listener may redirect to in v2, with the original port
	// specified in the FilterChainMatch destination_port field.
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
	return fileDescriptor_1700162847fce94e, []int{0, 0}
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

// Configuration for listener connection balancing.
type Listener_ConnectionBalanceConfig struct {
	// Types that are valid to be assigned to BalanceType:
	//	*Listener_ConnectionBalanceConfig_ExactBalance_
	BalanceType          isListener_ConnectionBalanceConfig_BalanceType `protobuf_oneof:"balance_type"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig) Reset()         { *m = Listener_ConnectionBalanceConfig{} }
func (m *Listener_ConnectionBalanceConfig) String() string { return proto.CompactTextString(m) }
func (*Listener_ConnectionBalanceConfig) ProtoMessage()    {}
func (*Listener_ConnectionBalanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1700162847fce94e, []int{0, 1}
}

func (m *Listener_ConnectionBalanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Size(m)
}
func (m *Listener_ConnectionBalanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig proto.InternalMessageInfo

type isListener_ConnectionBalanceConfig_BalanceType interface {
	isListener_ConnectionBalanceConfig_BalanceType()
}

type Listener_ConnectionBalanceConfig_ExactBalance_ struct {
	ExactBalance *Listener_ConnectionBalanceConfig_ExactBalance `protobuf:"bytes,1,opt,name=exact_balance,json=exactBalance,proto3,oneof"`
}

func (*Listener_ConnectionBalanceConfig_ExactBalance_) isListener_ConnectionBalanceConfig_BalanceType() {
}

func (m *Listener_ConnectionBalanceConfig) GetBalanceType() isListener_ConnectionBalanceConfig_BalanceType {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *Listener_ConnectionBalanceConfig) GetExactBalance() *Listener_ConnectionBalanceConfig_ExactBalance {
	if x, ok := m.GetBalanceType().(*Listener_ConnectionBalanceConfig_ExactBalance_); ok {
		return x.ExactBalance
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Listener_ConnectionBalanceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Listener_ConnectionBalanceConfig_ExactBalance_)(nil),
	}
}

// A connection balancer implementation that does exact balancing. This means that a lock is
// held during balancing so that connection counts are nearly exactly balanced between worker
// threads. This is "nearly" exact in the sense that a connection might close in parallel thus
// making the counts incorrect, but this should be rectified on the next accept. This balancer
// sacrifices accept throughput for accuracy and should be used when there are a small number of
// connections that rarely cycle (e.g., service mesh gRPC egress).
type Listener_ConnectionBalanceConfig_ExactBalance struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) Reset() {
	*m = Listener_ConnectionBalanceConfig_ExactBalance{}
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) String() string {
	return proto.CompactTextString(m)
}
func (*Listener_ConnectionBalanceConfig_ExactBalance) ProtoMessage() {}
func (*Listener_ConnectionBalanceConfig_ExactBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_1700162847fce94e, []int{0, 1, 0}
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Size(m)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.listener.v3.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.config.listener.v3.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.config.listener.v3.Listener.DeprecatedV1")
	proto.RegisterType((*Listener_ConnectionBalanceConfig)(nil), "envoy.config.listener.v3.Listener.ConnectionBalanceConfig")
	proto.RegisterType((*Listener_ConnectionBalanceConfig_ExactBalance)(nil), "envoy.config.listener.v3.Listener.ConnectionBalanceConfig.ExactBalance")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/listener.proto", fileDescriptor_1700162847fce94e)
}

var fileDescriptor_1700162847fce94e = []byte{
	// 1064 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x52, 0x1b, 0x37,
	0x14, 0x8e, 0xf9, 0xb5, 0x65, 0x1b, 0x8c, 0x68, 0x60, 0xe3, 0x01, 0xea, 0x30, 0x34, 0x35, 0x93,
	0x76, 0x3d, 0x31, 0x9d, 0x5e, 0x78, 0xb8, 0xb1, 0x21, 0x34, 0x50, 0x27, 0x90, 0xc5, 0x64, 0x9a,
	0xe9, 0xc5, 0x8e, 0xbc, 0x2b, 0x3b, 0x9a, 0x2e, 0x92, 0x2a, 0x69, 0x9d, 0x70, 0xd7, 0xcb, 0x4e,
	0x1f, 0xa1, 0x2f, 0xd0, 0x77, 0xe8, 0x7d, 0x67, 0x7a, 0xdb, 0xe7, 0xe8, 0x0b, 0x74, 0x7a, 0xd5,
	0x91, 0x76, 0xd7, 0xac, 0x61, 0x5d, 0x68, 0xef, 0x74, 0x74, 0xbe, 0xf3, 0xe9, 0xfc, 0x7c, 0x7b,
	0x16, 0x7c, 0x8a, 0xe9, 0x88, 0x5d, 0x35, 0x3c, 0x46, 0x07, 0x64, 0xd8, 0x08, 0x88, 0x54, 0x98,
	0x62, 0xd1, 0x18, 0xed, 0x8d, 0xcf, 0x36, 0x17, 0x4c, 0x31, 0x68, 0x19, 0xa0, 0x1d, 0x01, 0xed,
	0xb1, 0x73, 0xb4, 0x57, 0xdd, 0x9d, 0xa0, 0x40, 0x9e, 0x87, 0xa5, 0x0c, 0xd8, 0x50, 0x73, 0x8c,
	0x8d, 0x88, 0xa4, 0xba, 0x3d, 0x01, 0xf5, 0x98, 0xc0, 0x06, 0xe5, 0xfb, 0x02, 0x4b, 0x19, 0x63,
	0x3e, 0xce, 0xc4, 0xf4, 0x91, 0xc4, 0x31, 0xa0, 0x9e, 0x09, 0x90, 0xcc, 0xfb, 0x0e, 0x2b, 0x97,
	0x71, 0x45, 0x18, 0x8d, 0x91, 0x4f, 0xa7, 0x16, 0x87, 0x38, 0x71, 0x27, 0x0b, 0xac, 0x36, 0xef,
	0xec, 0x84, 0xeb, 0xb1, 0x4b, 0xce, 0x28, 0xa6, 0x4a, 0xde, 0x19, 0x13, 0xfa, 0xdc, 0x4d, 0xc5,
	0x99, 0xa6, 0x45, 0x31, 0x1b, 0x43, 0xc6, 0x86, 0x01, 0xd6, 0x29, 0x34, 0x10, 0xa5, 0x4c, 0x21,
	0x9d, 0x71, 0xc2, 0xb8, 0x15, 0x7b, 0x8d, 0xd5, 0x0f, 0x07, 0x0d, 0x3f, 0x14, 0x28, 0x55, 0xd2,
	0x2d, 0xff, 0x7b, 0x81, 0x38, 0xc7, 0x22, 0x89, 0xdf, 0x0c, 0x7d, 0x8e, 0xd2, 0xbc, 0x0d, 0xa9,
	0x90, 0x0a, 0x13, 0xf7, 0xe3, 0x5b, 0xee, 0x11, 0x16, 0x92, 0x30, 0x4a, 0x68, 0x92, 0xdf, 0xfa,
	0x08, 0x05, 0xc4, 0x47, 0x0a, 0x37, 0x92, 0x43, 0xe4, 0xd8, 0xfe, 0x73, 0x19, 0xe4, 0xbb, 0x71,
	0x49, 0x10, 0x82, 0x39, 0x8a, 0x2e, 0xb1, 0x95, 0xab, 0xe5, 0xea, 0x05, 0xc7, 0x9c, 0x61, 0x1b,
	0x2c, 0xc6, 0xa3, 0xb4, 0x66, 0x6a, 0xb9, 0x7a, 0xb1, 0xb9, 0x69, 0x4f, 0x88, 0x46, 0x8f, 0xca,
	0x1e, 0xed, 0xd9, 0xed, 0x08, 0xd4, 0xc9, 0xff, 0xdd, 0x99, 0xff, 0x29, 0x37, 0x53, 0xc9, 0x39,
	0x49, 0x1c, 0x3c, 0x01, 0xe5, 0x01, 0x09, 0x94, 0xee, 0xd9, 0x3b, 0x44, 0xa8, 0xb4, 0x66, 0x6b,
	0xb3, 0xf5, 0x62, 0xf3, 0x13, 0x7b, 0x9a, 0xfa, 0xec, 0x23, 0x03, 0x3f, 0xd0, 0x68, 0xa7, 0x34,
	0xb8, 0x36, 0x24, 0x1c, 0x80, 0xc7, 0x3c, 0x6a, 0x3e, 0xc5, 0x9e, 0x2e, 0xd6, 0xed, 0x87, 0x83,
	0x01, 0x16, 0x6e, 0x40, 0x2e, 0x89, 0x72, 0xfb, 0x57, 0x0a, 0x4b, 0x6b, 0xde, 0x24, 0xba, 0x61,
	0x47, 0x6d, 0xb5, 0x93, 0xb6, 0xda, 0x17, 0xc7, 0x54, 0xed, 0x35, 0xdf, 0xa0, 0x20, 0xc4, 0xce,
	0x26, 0xc7, 0xe2, 0x60, 0xcc, 0xd2, 0x31, 0x24, 0x5d, 0xcd, 0xd1, 0xd1, 0x14, 0xb0, 0x05, 0xf2,
	0x97, 0x58, 0x21, 0x1f, 0x29, 0x64, 0x2d, 0x18, 0xba, 0xad, 0xec, 0xba, 0x5f, 0xc6, 0x28, 0x67,
	0x8c, 0x87, 0x3d, 0x50, 0xf6, 0x31, 0x17, 0xd8, 0x43, 0x0a, 0xfb, 0xee, 0xe8, 0x99, 0xb5, 0x68,
	0x08, 0x1a, 0xd3, 0xeb, 0x4d, 0x26, 0x60, 0x1f, 0x8e, 0xe3, 0xde, 0x3c, 0x73, 0x4a, 0x7e, 0xca,
	0x82, 0x5f, 0x03, 0xe0, 0x0b, 0x44, 0xa8, 0xab, 0xae, 0x38, 0xb6, 0xf2, 0xb5, 0x5c, 0x7d, 0xa9,
	0xf9, 0xd9, 0x7d, 0x28, 0x75, 0x50, 0xef, 0x8a, 0x63, 0xa7, 0xe0, 0x27, 0x47, 0x78, 0x0e, 0x2a,
	0x63, 0x21, 0x47, 0xfd, 0x95, 0x56, 0xc1, 0x4c, 0xa5, 0x7e, 0x37, 0x65, 0x34, 0x1d, 0x67, 0x39,
	0x98, 0xb0, 0x25, 0x3c, 0x07, 0xd6, 0x4d, 0x52, 0x57, 0x91, 0x4b, 0xcc, 0x42, 0x65, 0x2d, 0x9b,
	0x16, 0x3c, 0xba, 0x35, 0x92, 0xc3, 0xf8, 0x4b, 0x70, 0xd6, 0x6e, 0xb0, 0xf5, 0xa2, 0x40, 0xf8,
	0x0a, 0xec, 0x78, 0x8c, 0x2a, 0x42, 0x43, 0xec, 0x32, 0xea, 0x4e, 0x7d, 0x60, 0xa5, 0x96, 0xab,
	0xe7, 0x9d, 0x5a, 0x82, 0x3d, 0xa5, 0xdd, 0x6c, 0xbe, 0x7d, 0x50, 0x54, 0x02, 0x51, 0xc9, 0x91,
	0xc0, 0x54, 0x59, 0xc0, 0xe4, 0x55, 0xbd, 0x95, 0x57, 0x87, 0xb1, 0x20, 0x12, 0x4a, 0x1a, 0x0e,
	0xbf, 0x04, 0xf9, 0x81, 0xc0, 0xb8, 0x4f, 0xa8, 0x6f, 0x15, 0xef, 0x0c, 0x1d, 0x63, 0xe1, 0x31,
	0x58, 0x9a, 0xd8, 0x65, 0xd2, 0x2a, 0x9b, 0x6e, 0x6f, 0x67, 0x8b, 0xea, 0xdc, 0x60, 0x4f, 0x0d,
	0xd4, 0x29, 0xcb, 0x94, 0x25, 0xe1, 0x37, 0xa0, 0xaa, 0x3c, 0xee, 0x0e, 0x90, 0xd4, 0x64, 0x98,
	0xba, 0xdf, 0x87, 0x38, 0xc4, 0x6e, 0x80, 0xe9, 0x50, 0xbd, 0xb3, 0x4a, 0xf7, 0x90, 0xfe, 0x9a,
	0xf2, 0xf8, 0x11, 0x92, 0xea, 0x94, 0x63, 0xfa, 0x5a, 0x07, 0x77, 0x4d, 0x2c, 0x3c, 0x07, 0x2b,
	0x4a, 0xa0, 0xc1, 0x80, 0x78, 0xae, 0x4f, 0x44, 0xf4, 0x61, 0x58, 0x15, 0x23, 0xb4, 0x27, 0xd9,
	0x79, 0xf6, 0x22, 0xf8, 0x61, 0x82, 0x76, 0x2a, 0xea, 0xc6, 0x0d, 0xfc, 0x16, 0xac, 0x66, 0xac,
	0x4d, 0x0b, 0x9a, 0x3c, 0x9f, 0x4e, 0x17, 0xdb, 0x85, 0xcf, 0x93, 0x09, 0x1e, 0x18, 0xaf, 0xb3,
	0x12, 0xde, 0xbc, 0x82, 0x2f, 0x40, 0x29, 0xbd, 0xf4, 0xad, 0x55, 0xc3, 0xfa, 0x2f, 0x8b, 0xa5,
	0xcd, 0x49, 0x42, 0xe1, 0x14, 0xd1, 0xb5, 0x01, 0x47, 0xe0, 0x51, 0x7a, 0xa7, 0xa0, 0x00, 0x51,
	0x0f, 0x27, 0xc9, 0x7e, 0x64, 0x68, 0x5b, 0xf7, 0xf8, 0xd8, 0x52, 0x1b, 0x25, 0xa2, 0x88, 0x73,
	0x5f, 0xf7, 0xb2, 0x1d, 0x70, 0x13, 0x00, 0x81, 0x43, 0x89, 0x5d, 0xce, 0x84, 0xb2, 0x1e, 0x1a,
	0x11, 0x17, 0xcc, 0xcd, 0x19, 0x13, 0x0a, 0x1e, 0x00, 0x10, 0xfd, 0x6e, 0xdd, 0x80, 0x0d, 0xad,
	0x35, 0xa3, 0x99, 0x9d, 0xc9, 0x3c, 0xae, 0x7f, 0xc7, 0xba, 0x3e, 0x63, 0x74, 0xd9, 0xd0, 0x29,
	0xa0, 0xe4, 0x58, 0x7d, 0x0f, 0x4a, 0xe9, 0xbd, 0x02, 0xf7, 0x41, 0x49, 0x8b, 0xd2, 0x55, 0x2c,
	0x7a, 0x35, 0x77, 0xa7, 0x90, 0x81, 0xc6, 0xf7, 0x98, 0x4e, 0xa9, 0xb5, 0xfb, 0xf3, 0x6f, 0x3f,
	0x6e, 0xed, 0x80, 0xe8, 0xaf, 0x6f, 0x23, 0x4e, 0xec, 0x51, 0x33, 0x7b, 0x81, 0x55, 0x7f, 0x99,
	0x01, 0xeb, 0x53, 0x3a, 0x02, 0x29, 0x28, 0xe3, 0x0f, 0xc8, 0x53, 0x49, 0xaf, 0xe3, 0x2c, 0xbe,
	0xfa, 0xff, 0x4d, 0xb6, 0x9f, 0x6b, 0xbe, 0xf8, 0xea, 0xc5, 0x03, 0xa7, 0x84, 0x53, 0x76, 0xf5,
	0x35, 0x28, 0xa5, 0xfd, 0xad, 0xb6, 0x2e, 0x63, 0x1f, 0xb4, 0xb2, 0xcb, 0xb8, 0xcf, 0x13, 0xad,
	0x2f, 0x34, 0x45, 0x03, 0x7c, 0xfe, 0x9f, 0x28, 0x3a, 0xab, 0xa0, 0x94, 0xc8, 0x4b, 0x6f, 0x72,
	0x38, 0xfb, 0x57, 0x27, 0xb7, 0xbd, 0x0b, 0x0a, 0xe3, 0x3d, 0x0d, 0x8b, 0x60, 0xf1, 0xf0, 0xf9,
	0x51, 0xfb, 0xa2, 0xdb, 0xab, 0x3c, 0x80, 0xcb, 0xa0, 0xf8, 0xf2, 0xf4, 0xf0, 0xf8, 0xe8, 0xad,
	0x7b, 0xfa, 0xaa, 0xfb, 0xb6, 0x92, 0x6b, 0x6d, 0xe8, 0x57, 0xd7, 0xc1, 0xc3, 0xcc, 0x57, 0x4f,
	0xe6, 0xf2, 0x4b, 0x95, 0xe5, 0x93, 0xb9, 0xfc, 0x5c, 0x65, 0xde, 0xa9, 0x68, 0x5d, 0x31, 0x41,
	0x86, 0x84, 0xa2, 0xc0, 0xf5, 0xa5, 0xea, 0xb4, 0x7f, 0xfd, 0xe1, 0xf7, 0x3f, 0x16, 0x66, 0x2a,
	0x33, 0xe0, 0x09, 0x61, 0x51, 0xa7, 0xb9, 0x60, 0x1f, 0xae, 0xa6, 0x36, 0xbd, 0x53, 0x4e, 0x98,
	0xcf, 0xb4, 0x2a, 0xce, 0x72, 0xfd, 0x05, 0x23, 0x8f, 0xbd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0x72, 0xb6, 0x1a, 0x64, 0x0a, 0x00, 0x00,
}
