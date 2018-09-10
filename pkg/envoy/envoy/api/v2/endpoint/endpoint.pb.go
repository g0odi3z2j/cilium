// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

/*
Package endpoint is a generated protocol buffer package.

It is generated from these files:
	envoy/api/v2/endpoint/endpoint.proto
	envoy/api/v2/endpoint/load_report.proto

It has these top-level messages:
	Endpoint
	LbEndpoint
	LocalityLbEndpoints
	UpstreamLocalityStats
	UpstreamEndpointStats
	EndpointLoadMetricStats
	ClusterStats
*/
package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import envoy_api_v2_core2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	Address *envoy_api_v2_core1.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	//
	// .. attention::
	//
	//   This takes into effect only for upstream clusters with
	//   :ref:`active health checking <arch_overview_health_checking>` enabled.
	HealthCheckConfig *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig" json:"health_check_config,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Endpoint) GetAddress() *envoy_api_v2_core1.Address {
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
	PortValue uint32 `protobuf:"varint,1,opt,name=port_value,json=portValue" json:"port_value,omitempty"`
}

func (m *Endpoint_HealthCheckConfig) Reset()                    { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string            { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()               {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	// Upstream host identifier
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	// Optional health status when known and supplied by EDS server.
	HealthStatus envoy_api_v2_core2.HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	// The endpoint metadata specifies values that may be used by the load
	// balancer to select endpoints in a cluster for a given request. The filter
	// name should be specified as *envoy.lb*. An example boolean key-value pair
	// is *canary*, providing the optional canary status of the upstream host.
	// This may be matched against in a route's ForwardAction metadata_match field
	// to subset the endpoints considered in cluster load balancing.
	Metadata *envoy_api_v2_core.Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// The optional load balancing weight of the upstream host, in the range 1 -
	// 128. Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight *google_protobuf1.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
}

func (m *LbEndpoint) Reset()                    { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string            { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()               {}
func (*LbEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetHealthStatus() envoy_api_v2_core2.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return envoy_api_v2_core2.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *envoy_api_v2_core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies location of where the upstream hosts run.
	Locality *envoy_api_v2_core.Locality `protobuf:"bytes,1,opt,name=locality" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight - range 1-128. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities  at the same priority level to produce the effective percentage
	// of traffic for the locality.
	//
	// Locality weights are only considered when :ref:`locality weighted load
	// balancing <arch_overview_load_balancing_locality_weighted_lb>` is
	// configured. These weights are ignored otherwise. If no weights are
	// specificed when locality weighted load balancing is enabled, the cluster is
	// assumed to have a weight of 1.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight *google_protobuf1.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
	// Optional: the priority for this LocalityLbEndpoints. If unspecified this will
	// default to the highest priority (0).
	//
	// Under usual circumstances, Envoy will only select endpoints for the highest
	// priority (0). In the event all endpoints for a particular priority are
	// unavailable/unhealthy, Envoy will fail over to selecting endpoints for the
	// next highest priority group.
	//
	// Priorities should range from 0 (highest) to N (lowest) without skipping.
	Priority uint32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
}

func (m *LocalityLbEndpoints) Reset()                    { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string            { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()               {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LocalityLbEndpoints) GetLocality() *envoy_api_v2_core.Locality {
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

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *google_protobuf1.UInt32Value {
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

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}

func init() { proto.RegisterFile("envoy/api/v2/endpoint/endpoint.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x26, 0x4d, 0xbb, 0x15, 0xb7, 0x9b, 0xb4, 0x94, 0x89, 0xa8, 0x4c, 0x2b, 0x54, 0x13, 0xaa,
	0x7a, 0x70, 0x44, 0x87, 0xc4, 0x81, 0x03, 0x22, 0x03, 0x09, 0xd0, 0xb8, 0x18, 0x01, 0x12, 0x07,
	0x22, 0x27, 0xf1, 0x12, 0x0b, 0x13, 0x47, 0x89, 0x9b, 0x69, 0xb7, 0xfd, 0x16, 0x4e, 0xfb, 0x0d,
	0x9c, 0x38, 0xf2, 0x2b, 0x38, 0x20, 0x2e, 0xfc, 0x8a, 0x21, 0x3b, 0x76, 0xd6, 0xad, 0x9d, 0xb8,
	0xec, 0xf6, 0xec, 0xf7, 0x7d, 0xef, 0x7d, 0xdf, 0xa7, 0x07, 0xf6, 0x48, 0x56, 0xf1, 0x13, 0x0f,
	0xe7, 0xd4, 0xab, 0x66, 0x1e, 0xc9, 0xe2, 0x9c, 0xd3, 0x4c, 0x34, 0x05, 0xcc, 0x0b, 0x2e, 0xb8,
	0xb3, 0xad, 0x50, 0x10, 0xe7, 0x14, 0x56, 0x33, 0x68, 0x9a, 0xc3, 0xd1, 0x25, 0x72, 0xc4, 0x0b,
	0xe2, 0xe1, 0x38, 0x2e, 0x48, 0x59, 0xd6, 0xbc, 0xe1, 0xce, 0x32, 0x20, 0xc4, 0x25, 0xd1, 0xdd,
	0xbd, 0xe5, 0x6e, 0x4a, 0x30, 0x13, 0x69, 0x10, 0xa5, 0x24, 0xfa, 0xa2, 0x51, 0xbb, 0x09, 0xe7,
	0x09, 0x23, 0x9e, 0x7a, 0x85, 0xf3, 0x23, 0xef, 0xb8, 0xc0, 0x79, 0x4e, 0x0a, 0xb3, 0xe3, 0x6e,
	0x85, 0x19, 0x8d, 0xb1, 0x20, 0x9e, 0x29, 0x74, 0xe3, 0x4e, 0xc2, 0x13, 0xae, 0x4a, 0x4f, 0x56,
	0xf5, 0xef, 0xf8, 0x8f, 0x05, 0xba, 0x2f, 0xb5, 0x01, 0xe7, 0x31, 0x58, 0xd7, 0x82, 0x5d, 0xeb,
	0xbe, 0x35, 0xe9, 0xcd, 0x86, 0xf0, 0x92, 0x53, 0xa9, 0x09, 0x3e, 0xaf, 0x11, 0xc8, 0x40, 0x1d,
	0x0c, 0x06, 0x8b, 0x3a, 0x83, 0x88, 0x67, 0x47, 0x34, 0x71, 0x5b, 0x6a, 0xc2, 0x23, 0xb8, 0x32,
	0x2b, 0x68, 0x76, 0xc2, 0x57, 0x8a, 0x7a, 0x20, 0x99, 0x07, 0x8a, 0x88, 0xb6, 0xd2, 0xab, 0x5f,
	0xc3, 0x67, 0x60, 0x6b, 0x09, 0xe7, 0x4c, 0x01, 0xc8, 0x79, 0x21, 0x82, 0x0a, 0xb3, 0x39, 0x51,
	0x82, 0x37, 0xfc, 0xde, 0xf7, 0xbf, 0x3f, 0xec, 0xb5, 0x69, 0xdb, 0x3d, 0x3f, 0xb7, 0xd1, 0x6d,
	0xd9, 0xfe, 0x20, 0xbb, 0xe3, 0xb3, 0x16, 0x00, 0x87, 0x61, 0x63, 0xf4, 0x29, 0xe8, 0x1a, 0x25,
	0xda, 0xe9, 0xe8, 0x3f, 0x3a, 0x51, 0x43, 0x70, 0x5e, 0x80, 0x0d, 0xed, 0xb7, 0x14, 0x58, 0xcc,
	0x4b, 0xe5, 0x74, 0xf3, 0xea, 0x04, 0x95, 0x55, 0x2d, 0xfa, 0x9d, 0x82, 0xa1, 0x7e, 0xba, 0xf0,
	0x72, 0x9e, 0x80, 0xee, 0x57, 0x22, 0x70, 0x8c, 0x05, 0x76, 0x6d, 0x25, 0xe1, 0xde, 0x8a, 0x01,
	0x6f, 0x35, 0x04, 0x35, 0x60, 0xe7, 0x33, 0xd8, 0x66, 0x1c, 0xc7, 0x41, 0x88, 0x19, 0xce, 0x22,
	0x9a, 0x25, 0xc1, 0x31, 0xa1, 0x49, 0x2a, 0xdc, 0xb6, 0x9a, 0xb2, 0x03, 0xeb, 0x03, 0x81, 0xe6,
	0x40, 0xe0, 0xfb, 0xd7, 0x99, 0xd8, 0x9f, 0xa9, 0x1c, 0xfc, 0xbe, 0xcc, 0x67, 0x7d, 0xda, 0x71,
	0x4f, 0xad, 0x89, 0x85, 0x06, 0x72, 0x90, 0x6f, 0xe6, 0x7c, 0x54, 0x63, 0xc6, 0xdf, 0x5a, 0x60,
	0x70, 0xc8, 0x23, 0xcc, 0xa8, 0x38, 0xb9, 0x88, 0x4c, 0x09, 0x66, 0xfa, 0x5b, 0x67, 0xb6, 0x4a,
	0xb0, 0x61, 0xa2, 0x06, 0xec, 0xbc, 0x01, 0x7d, 0x16, 0x06, 0x26, 0x3e, 0x19, 0x97, 0x3d, 0xe9,
	0xcd, 0x1e, 0x5c, 0x13, 0xf8, 0xc5, 0x4a, 0xbf, 0xfd, 0xf3, 0xd7, 0xe8, 0x16, 0xea, 0xb1, 0x05,
	0x11, 0xd7, 0x9a, 0xb7, 0x6f, 0xc4, 0xbc, 0xf3, 0x10, 0x74, 0xf3, 0x82, 0xf2, 0x42, 0x9a, 0xec,
	0xa8, 0x8b, 0x02, 0x92, 0xd4, 0x99, 0xda, 0xee, 0xa9, 0x85, 0x9a, 0x9e, 0xbf, 0x79, 0xf6, 0x7b,
	0xd7, 0xfa, 0xd4, 0xdc, 0x44, 0xb8, 0xa6, 0x16, 0xee, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x65, 0xf1, 0xe6, 0x40, 0x04, 0x00, 0x00,
}
