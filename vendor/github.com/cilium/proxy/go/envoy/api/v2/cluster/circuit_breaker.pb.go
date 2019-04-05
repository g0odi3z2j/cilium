// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/circuit_breaker.proto

package cluster

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/gogo/protobuf/gogoproto"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_cluster.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_89bc8d4e21efdd79, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`.
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	// [#comment:TODO(htuch): add (validate.rules).enum.defined_only = true once
	// https://github.com/lyft/protoc-gen-validate/issues/42 is resolved.]
	Priority core.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.api.v2.core.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_89bc8d4e21efdd79, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v2.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v2.cluster.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/circuit_breaker.proto", fileDescriptor_89bc8d4e21efdd79)
}

var fileDescriptor_89bc8d4e21efdd79 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x40, 0xe5, 0xa6, 0xb4, 0xd5, 0x06, 0x25, 0x92, 0x89, 0x90, 0x15, 0x55, 0x55, 0x94, 0x0b,
	0x11, 0x87, 0x35, 0x72, 0xcf, 0x80, 0x48, 0xd4, 0x03, 0x97, 0xca, 0x32, 0xd0, 0x03, 0x17, 0x6b,
	0xe3, 0x0e, 0xee, 0xaa, 0xf6, 0xce, 0x32, 0xbb, 0x0e, 0xce, 0x1f, 0x21, 0x3e, 0x83, 0xef, 0xe0,
	0xc4, 0x97, 0x20, 0x7b, 0x9d, 0x84, 0x56, 0x3d, 0xf8, 0xb6, 0x9e, 0x99, 0xf7, 0x3c, 0x3b, 0x3b,
	0xec, 0x35, 0xa8, 0x0d, 0x6e, 0x43, 0xa1, 0x65, 0xb8, 0x89, 0xc2, 0xac, 0xa8, 0x8c, 0x05, 0x0a,
	0x33, 0x49, 0x59, 0x25, 0x6d, 0xba, 0x26, 0x10, 0xf7, 0x40, 0x5c, 0x13, 0x5a, 0xf4, 0x27, 0x6d,
	0x2d, 0x17, 0x5a, 0xf2, 0x4d, 0xc4, 0xbb, 0xda, 0xe9, 0xf9, 0x43, 0x03, 0x12, 0x84, 0x6b, 0x61,
	0xc0, 0x31, 0xd3, 0x8b, 0x1c, 0x31, 0x2f, 0x20, 0x6c, 0xbf, 0xd6, 0xd5, 0xb7, 0xf0, 0x07, 0x09,
	0xad, 0x81, 0x4c, 0x97, 0x9f, 0xe4, 0x98, 0x63, 0x7b, 0x0c, 0x9b, 0x93, 0x8b, 0xce, 0x7f, 0x1f,
	0xb3, 0xf1, 0xca, 0xf5, 0xb0, 0x74, 0x2d, 0x18, 0x3f, 0x66, 0xcc, 0xde, 0x11, 0x98, 0x3b, 0x2c,
	0x6e, 0x4d, 0xe0, 0xcd, 0x06, 0x8b, 0x61, 0xf4, 0x86, 0x3f, 0xd5, 0x12, 0x7f, 0x84, 0xf2, 0xcf,
	0x7b, 0x2e, 0xf9, 0xcf, 0x31, 0xfd, 0x33, 0x60, 0xec, 0x90, 0xf2, 0xdf, 0xb1, 0x33, 0x4d, 0x12,
	0x49, 0xda, 0x6d, 0xe0, 0xcd, 0xbc, 0xc5, 0x28, 0x9a, 0x3f, 0xd2, 0x23, 0x01, 0x4f, 0xb0, 0xb2,
	0x52, 0xe5, 0x71, 0x57, 0x99, 0xec, 0x19, 0xff, 0x8a, 0x8d, 0x4b, 0x51, 0xa7, 0x19, 0x2a, 0x05,
	0x99, 0x95, 0xa8, 0x4c, 0x70, 0x34, 0xf3, 0x16, 0xc3, 0xe8, 0x9c, 0xbb, 0x21, 0xf0, 0xdd, 0x10,
	0xf8, 0x97, 0x8f, 0xca, 0x5e, 0x46, 0x37, 0xa2, 0xa8, 0x20, 0x19, 0x95, 0xa2, 0x5e, 0x1d, 0x18,
	0xff, 0x9a, 0x4d, 0x1a, 0x8d, 0x06, 0x75, 0x2b, 0x55, 0x9e, 0x12, 0x7c, 0xaf, 0xc0, 0x58, 0x13,
	0x0c, 0x7a, 0xb8, 0xfc, 0x52, 0xd4, 0xb1, 0x03, 0x93, 0x8e, 0xf3, 0xdf, 0xb3, 0xe7, 0x8d, 0x6f,
	0xef, 0x39, 0xee, 0xe1, 0x19, 0x96, 0xa2, 0xde, 0x0b, 0xde, 0xb2, 0xa1, 0x13, 0x58, 0x92, 0x60,
	0x82, 0x67, 0x3d, 0x78, 0xd6, 0xf2, 0x6d, 0xbd, 0xff, 0x8a, 0x8d, 0x2d, 0x89, 0xec, 0x3e, 0x25,
	0x28, 0x85, 0x54, 0x52, 0xe5, 0xc1, 0xc9, 0xcc, 0x5b, 0x9c, 0x25, 0xa3, 0x36, 0x9c, 0xec, 0xa2,
	0xbb, 0x8b, 0x1f, 0xe6, 0x97, 0x6a, 0xc4, 0xc2, 0x04, 0xa7, 0x3d, 0x2f, 0x7e, 0x18, 0x62, 0xdc,
	0x70, 0x4b, 0xf1, 0xf3, 0xef, 0x85, 0xc7, 0xe6, 0x12, 0xdd, 0x2b, 0x6a, 0xc2, 0x7a, 0xfb, 0xe4,
	0xbe, 0x2c, 0x5f, 0x3c, 0x5c, 0x98, 0xb8, 0xf9, 0x43, 0xec, 0x7d, 0x3d, 0xed, 0xf2, 0xbf, 0x8e,
	0x5e, 0x5e, 0xb5, 0xd8, 0x07, 0x2d, 0xf9, 0x4d, 0xc4, 0x57, 0x2e, 0x7c, 0xfd, 0x69, 0x7d, 0xd2,
	0x36, 0x73, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x38, 0x11, 0xf7, 0x46, 0x03, 0x00, 0x00,
}
