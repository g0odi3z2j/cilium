// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

package endpoint

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Upstream host identifier.
type Endpoint struct {
	// The upstream host address.
	//
	// .. attention::
	//
	//   The form of host address depends on the given cluster type. For STATIC or EDS,
	//   it is expected to be a direct IP address (or something resolvable by the
	//   specified :ref:`resolver <envoy_api_field_core.SocketAddress.resolver_name>`
	//   in the Address). For LOGICAL or STRICT DNS, it is expected to be hostname,
	//   and will be resolved via DNS.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	//
	// .. attention::
	//
	//   This takes into effect only for upstream clusters with
	//   :ref:`active health checking <arch_overview_health_checking>` enabled.
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2a3e4ee06324f, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

// The optional health check configuration.
type Endpoint_HealthCheckConfig struct {
	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2a3e4ee06324f, []int{0, 0}
}

func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(m, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	// Upstream host identifier or a named reference.
	//
	// Types that are valid to be assigned to HostIdentifier:
	//	*LbEndpoint_Endpoint
	//	*LbEndpoint_EndpointName
	HostIdentifier isLbEndpoint_HostIdentifier `protobuf_oneof:"host_identifier"`
	// Optional health status when known and supplied by EDS server.
	HealthStatus core.HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	// The endpoint metadata specifies values that may be used by the load
	// balancer to select endpoints in a cluster for a given request. The filter
	// name should be specified as *envoy.lb*. An example boolean key-value pair
	// is *canary*, providing the optional canary status of the upstream host.
	// This may be matched against in a route's
	// :ref:`RouteAction <envoy_api_msg_route.RouteAction>` metadata_match field
	// to subset the endpoints considered in cluster load balancing.
	Metadata *core.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The optional load balancing weight of the upstream host; at least 1.
	// Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2a3e4ee06324f, []int{1}
}

func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

type isLbEndpoint_HostIdentifier interface {
	isLbEndpoint_HostIdentifier()
}

type LbEndpoint_Endpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3,oneof"`
}

type LbEndpoint_EndpointName struct {
	EndpointName string `protobuf:"bytes,5,opt,name=endpoint_name,json=endpointName,proto3,oneof"`
}

func (*LbEndpoint_Endpoint) isLbEndpoint_HostIdentifier() {}

func (*LbEndpoint_EndpointName) isLbEndpoint_HostIdentifier() {}

func (m *LbEndpoint) GetHostIdentifier() isLbEndpoint_HostIdentifier {
	if m != nil {
		return m.HostIdentifier
	}
	return nil
}

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_Endpoint); ok {
		return x.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetEndpointName() string {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_EndpointName); ok {
		return x.EndpointName
	}
	return ""
}

func (m *LbEndpoint) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LbEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LbEndpoint_Endpoint)(nil),
		(*LbEndpoint_EndpointName)(nil),
	}
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies location of where the upstream hosts run.
	Locality *core.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight; at least 1. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities  at the same priority level to produce the effective percentage
	// of traffic for the locality.
	//
	// Locality weights are only considered when :ref:`locality weighted load
	// balancing <arch_overview_load_balancing_locality_weighted_lb>` is
	// configured. These weights are ignored otherwise. If no weights are
	// specified when locality weighted load balancing is enabled, the locality is
	// assigned no load.
	LoadBalancingWeight *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	// Optional: the priority for this LocalityLbEndpoints. If unspecified this will
	// default to the highest priority (0).
	//
	// Under usual circumstances, Envoy will only select endpoints for the highest
	// priority (0). In the event all endpoints for a particular priority are
	// unavailable/unhealthy, Envoy will fail over to selecting endpoints for the
	// next highest priority group.
	//
	// Priorities should range from 0 (highest) to N (lowest) without skipping.
	Priority uint32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	// Optional: Per locality proximity value which indicates how close this
	// locality is from the source locality. This value only provides ordering
	// information (lower the value, closer it is to the source locality).
	// This will be consumed by load balancing schemes that need proximity order
	// to determine where to route the requests.
	// [#not-implemented-hide:]
	Proximity            *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=proximity,proto3" json:"proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2a3e4ee06324f, []int{2}
}

func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LocalityLbEndpoints) GetProximity() *wrappers.UInt32Value {
	if m != nil {
		return m.Proximity
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/endpoint.proto", fileDescriptor_a9d2a3e4ee06324f)
}

var fileDescriptor_a9d2a3e4ee06324f = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0xad, 0xe3, 0xb4, 0x2f, 0x99, 0x34, 0x0f, 0xd5, 0xa1, 0xc2, 0x0a, 0x55, 0x53, 0x42, 0x41,
	0x51, 0x16, 0x63, 0x91, 0x22, 0x21, 0x21, 0x21, 0x84, 0x5b, 0xa4, 0x80, 0x0a, 0xaa, 0x06, 0x01,
	0x12, 0x2c, 0xac, 0xb1, 0x3d, 0xb1, 0x47, 0x38, 0x1e, 0xcb, 0x9e, 0xa4, 0x64, 0xc7, 0xe7, 0xf0,
	0x0d, 0xac, 0x58, 0xf2, 0x0f, 0x48, 0x48, 0x20, 0x36, 0x7c, 0x45, 0xd1, 0x8c, 0x3d, 0x4e, 0xda,
	0xa4, 0xea, 0x82, 0xdd, 0xf5, 0xdc, 0x73, 0xee, 0x3d, 0xe7, 0xf8, 0x82, 0x7d, 0x12, 0x4f, 0xd9,
	0xcc, 0xc2, 0x09, 0xb5, 0xa6, 0x03, 0x8b, 0xc4, 0x7e, 0xc2, 0x68, 0xcc, 0xcb, 0x02, 0x26, 0x29,
	0xe3, 0xcc, 0xd8, 0x96, 0x28, 0x88, 0x13, 0x0a, 0xa7, 0x03, 0xa8, 0x9a, 0xed, 0xce, 0x39, 0xb2,
	0xc7, 0x52, 0x62, 0x61, 0xdf, 0x4f, 0x49, 0x96, 0xe5, 0xbc, 0xf6, 0xce, 0x32, 0xc0, 0xc5, 0x19,
	0x29, 0xba, 0xfb, 0xcb, 0xdd, 0x90, 0xe0, 0x88, 0x87, 0x8e, 0x17, 0x12, 0xef, 0x43, 0x81, 0xda,
	0x0d, 0x18, 0x0b, 0x22, 0x62, 0xc9, 0x2f, 0x77, 0x32, 0xb2, 0x4e, 0x53, 0x9c, 0x24, 0x24, 0x55,
	0x3b, 0x6e, 0x4c, 0x71, 0x44, 0x7d, 0xcc, 0x89, 0xa5, 0x8a, 0xa2, 0x71, 0x3d, 0x60, 0x01, 0x93,
	0xa5, 0x25, 0xaa, 0xfc, 0xb5, 0xfb, 0x5b, 0x03, 0xb5, 0xa7, 0x85, 0x01, 0xe3, 0x3e, 0xf8, 0xaf,
	0x10, 0x6c, 0x6a, 0x7b, 0x5a, 0xaf, 0x31, 0x68, 0xc3, 0x73, 0x4e, 0x85, 0x26, 0xf8, 0x24, 0x47,
	0x20, 0x05, 0x35, 0x30, 0x68, 0x2d, 0xea, 0x74, 0x3c, 0x16, 0x8f, 0x68, 0x60, 0x56, 0xe4, 0x84,
	0x7b, 0x70, 0x65, 0x56, 0x50, 0xed, 0x84, 0x43, 0x49, 0x3d, 0x14, 0xcc, 0x43, 0x49, 0x44, 0x5b,
	0xe1, 0xc5, 0xa7, 0xf6, 0x63, 0xb0, 0xb5, 0x84, 0x33, 0xfa, 0x00, 0x24, 0x2c, 0xe5, 0xce, 0x14,
	0x47, 0x13, 0x22, 0x05, 0x37, 0xed, 0xc6, 0x97, 0x3f, 0x5f, 0xf5, 0x8d, 0x7e, 0xd5, 0x3c, 0x3b,
	0xd3, 0x51, 0x5d, 0xb4, 0xdf, 0x88, 0x6e, 0xf7, 0x67, 0x05, 0x80, 0x63, 0xb7, 0x34, 0xfa, 0x08,
	0xd4, 0x94, 0x92, 0xc2, 0x69, 0xe7, 0x0a, 0x9d, 0xc3, 0x35, 0x54, 0x52, 0x8c, 0x3b, 0xa0, 0xa9,
	0x6a, 0x27, 0xc6, 0x63, 0x62, 0xae, 0xef, 0x69, 0xbd, 0xfa, 0x70, 0x0d, 0x6d, 0xaa, 0xe7, 0x97,
	0x78, 0x4c, 0x8c, 0x23, 0xd0, 0x2c, 0x82, 0xc9, 0x38, 0xe6, 0x93, 0x4c, 0x46, 0xf2, 0xff, 0xc5,
	0x55, 0x32, 0xd4, 0xdc, 0xdd, 0x2b, 0x09, 0x43, 0x9b, 0xe1, 0xc2, 0x97, 0xf1, 0x00, 0xd4, 0xc6,
	0x84, 0x63, 0x1f, 0x73, 0x6c, 0xea, 0x52, 0xeb, 0xcd, 0x15, 0x03, 0x5e, 0x14, 0x10, 0x54, 0x82,
	0x8d, 0xf7, 0x60, 0x3b, 0x62, 0xd8, 0x77, 0x5c, 0x1c, 0xe1, 0xd8, 0xa3, 0x71, 0xe0, 0x9c, 0x12,
	0x1a, 0x84, 0xdc, 0xac, 0xca, 0x29, 0x3b, 0x30, 0xbf, 0x24, 0xa8, 0x2e, 0x09, 0xbe, 0x7e, 0x16,
	0xf3, 0x83, 0x81, 0x0c, 0xcc, 0xae, 0x8b, 0x20, 0xab, 0xfd, 0x4a, 0x4f, 0x43, 0x2d, 0x31, 0xc5,
	0x56, 0x43, 0xde, 0xca, 0x19, 0xf6, 0x16, 0xb8, 0x16, 0xb2, 0x8c, 0x3b, 0xd4, 0x27, 0x31, 0xa7,
	0x23, 0x4a, 0xd2, 0xee, 0xf7, 0x0a, 0x68, 0x1d, 0x33, 0x0f, 0x47, 0x94, 0xcf, 0xe6, 0x59, 0x4b,
	0x03, 0x51, 0xf1, 0x5c, 0x84, 0xbd, 0xca, 0x80, 0x62, 0xa2, 0x12, 0x6c, 0x3c, 0x07, 0x9b, 0x91,
	0xeb, 0xa8, 0x48, 0x45, 0x7c, 0x7a, 0xaf, 0x31, 0xb8, 0x75, 0xc9, 0x9f, 0x9a, 0xaf, 0xb4, 0xab,
	0xdf, 0x7e, 0x74, 0xd6, 0x50, 0x23, 0x5a, 0x10, 0x71, 0x69, 0x18, 0xfa, 0xbf, 0x87, 0x61, 0xdc,
	0x05, 0xb5, 0x24, 0xa5, 0x2c, 0x15, 0x0e, 0xd7, 0xe5, 0x1d, 0x02, 0xc1, 0x58, 0xef, 0xeb, 0xe6,
	0x27, 0x0d, 0x95, 0x3d, 0xe3, 0x21, 0xa8, 0x27, 0x29, 0xfb, 0x48, 0xc7, 0x02, 0xb8, 0x71, 0xf5,
	0x62, 0x34, 0x87, 0xdb, 0x47, 0x9f, 0x7f, 0xed, 0x6a, 0xe0, 0x36, 0x65, 0xb9, 0x7d, 0xd1, 0x98,
	0xad, 0x4e, 0xc2, 0x6e, 0x2a, 0xdb, 0x27, 0x62, 0xe6, 0x89, 0xf6, 0xae, 0xbc, 0x5c, 0x77, 0x43,
	0xae, 0x39, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xc0, 0xa8, 0xdc, 0xe8, 0x04, 0x00, 0x00,
}
