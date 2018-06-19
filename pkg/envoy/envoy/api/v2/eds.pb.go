// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Each route from RDS will map to a single cluster or traffic split across
// clusters using weights expressed in the RDS WeightedCluster.
//
// With EDS, each cluster is treated independently from a LB perspective, with
// LB taking place between the Localities within a cluster and at a finer
// granularity between the hosts within a locality. For a given cluster, the
// effective weight of a host is its load_balancing_weight multiplied by the
// load_balancing_weight of its Locality.
type ClusterLoadAssignment struct {
	// Name of the cluster. This will be the :ref:`service_name
	// <envoy_api_field_Cluster.EdsClusterConfig.service_name>` value if specified
	// in the cluster :ref:`EdsClusterConfig
	// <envoy_api_msg_Cluster.EdsClusterConfig>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// List of endpoints to load balance to.
	Endpoints []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Load balancing policy settings.
	Policy               *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_67a9695caaec9f00, []int{0}
}
func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(dst, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Load balancing policy settings.
type ClusterLoadAssignment_Policy struct {
	// Percentage of traffic (0-100) that should be dropped. This
	// action allows protection of upstream hosts should they unable to
	// recover from an outage or should they be unable to autoscale and hence
	// overall incoming traffic volume need to be trimmed to protect them.
	// [#v2-api-diff: This is known as maintenance mode in v1.]
	DropOverload         float64  `protobuf:"fixed64,1,opt,name=drop_overload,json=dropOverload,proto3" json:"drop_overload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_67a9695caaec9f00, []int{0, 0}
}
func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(dst, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverload() float64 {
	if m != nil {
		return m.DropOverload
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_eds_67a9695caaec9f00) }

var fileDescriptor_eds_67a9695caaec9f00 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x8b, 0xd4, 0x40,
	0x18, 0xce, 0xe4, 0x96, 0x95, 0x9d, 0x5b, 0x15, 0x86, 0xd3, 0xcb, 0x85, 0xe5, 0x2e, 0x04, 0x8b,
	0xf5, 0x90, 0x44, 0x62, 0x77, 0x85, 0x68, 0xfc, 0xc0, 0x62, 0x39, 0x25, 0xd7, 0x68, 0x75, 0xcc,
	0x25, 0x2f, 0x71, 0x20, 0x99, 0x37, 0x66, 0x66, 0x03, 0x69, 0xad, 0xec, 0xfd, 0x09, 0x36, 0xfe,
	0x06, 0x2b, 0x4b, 0x7b, 0x7b, 0x0b, 0xb1, 0x11, 0xff, 0x84, 0xe4, 0x53, 0x83, 0xda, 0x39, 0xd5,
	0xcb, 0xfb, 0x7c, 0xe4, 0x79, 0x26, 0x43, 0xaf, 0x83, 0xac, 0xb0, 0xf6, 0x79, 0x21, 0xfc, 0x2a,
	0xf0, 0x21, 0x51, 0x5e, 0x51, 0xa2, 0x46, 0xb6, 0x6c, 0xf7, 0x1e, 0x2f, 0x84, 0x57, 0x05, 0xf6,
	0x6a, 0xc2, 0x4a, 0x84, 0x8a, 0xb1, 0x82, 0xb2, 0xee, 0xb8, 0xf6, 0x8d, 0xa9, 0x87, 0x4c, 0x0a,
	0x14, 0x52, 0x8f, 0x43, 0xcf, 0x5a, 0xa5, 0x88, 0x69, 0x06, 0x2d, 0x8d, 0x4b, 0x89, 0x9a, 0x6b,
	0x81, 0xb2, 0xff, 0x9e, 0xbd, 0x5f, 0xf1, 0x4c, 0x24, 0x5c, 0x83, 0x3f, 0x0c, 0x3d, 0xb0, 0x97,
	0x62, 0x8a, 0xed, 0xe8, 0x37, 0x53, 0xb7, 0x75, 0xdf, 0x99, 0xf4, 0xda, 0x83, 0x6c, 0xab, 0x34,
	0x94, 0x1b, 0xe4, 0xc9, 0x7d, 0xa5, 0x44, 0x2a, 0x73, 0x90, 0x9a, 0xdd, 0xa2, 0xcb, 0xb8, 0x03,
	0xce, 0x25, 0xcf, 0xc1, 0x22, 0x0e, 0x59, 0x2f, 0xc2, 0xc5, 0x87, 0xef, 0x1f, 0x77, 0x66, 0xa5,
	0xe9, 0x90, 0x68, 0xb7, 0x87, 0x4f, 0x79, 0x0e, 0xec, 0x94, 0x2e, 0x86, 0x98, 0xca, 0x32, 0x9d,
	0x9d, 0xf5, 0x6e, 0x70, 0xec, 0xfd, 0x5e, 0xdd, 0x1b, 0x5b, 0x6c, 0x30, 0xe6, 0x99, 0xd0, 0xf5,
	0xe6, 0xe2, 0xd1, 0xa0, 0x08, 0x67, 0x9f, 0xbe, 0x1c, 0x19, 0xd1, 0x2f, 0x0b, 0x16, 0xd2, 0x79,
	0x81, 0x99, 0x88, 0x6b, 0x6b, 0xe6, 0x90, 0x3f, 0xcd, 0xfe, 0x1a, 0xd9, 0x7b, 0xd6, 0x2a, 0xa2,
	0x5e, 0x69, 0x3f, 0xa1, 0xf3, 0x6e, 0xc3, 0xee, 0xd2, 0xcb, 0x49, 0x89, 0xc5, 0x79, 0x73, 0xd9,
	0x19, 0xf2, 0xa4, 0x2d, 0x43, 0xc2, 0x83, 0xa6, 0xcc, 0x1e, 0x63, 0x07, 0x46, 0x7b, 0x5e, 0xdc,
	0xbb, 0x69, 0xf4, 0x27, 0x5a, 0x36, 0xfc, 0xa7, 0x3d, 0x3d, 0xf8, 0x41, 0xa8, 0x35, 0x84, 0x7d,
	0x38, 0xfc, 0xb4, 0x33, 0x28, 0x2b, 0x11, 0x03, 0x7b, 0x4e, 0xaf, 0x9e, 0xe9, 0x12, 0x78, 0x3e,
	0xd6, 0x61, 0x87, 0xd3, 0xb4, 0xa3, 0x24, 0x82, 0x57, 0x5b, 0x50, 0xda, 0x3e, 0xfa, 0x27, 0xae,
	0x0a, 0x94, 0x0a, 0x5c, 0x63, 0x4d, 0x6e, 0x13, 0xb6, 0xa5, 0x57, 0x1e, 0x83, 0x8e, 0x5f, 0xfe,
	0x47, 0x63, 0xf7, 0xf5, 0xe7, 0x6f, 0x6f, 0xcd, 0x95, 0xbb, 0x3f, 0x79, 0x7f, 0x27, 0xe3, 0xc5,
	0x9f, 0x90, 0xe3, 0xf0, 0xd2, 0xfb, 0xaf, 0x87, 0xe4, 0x0d, 0x21, 0x17, 0xf3, 0xf6, 0x8d, 0xdc,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xc0, 0xab, 0x3d, 0xdc, 0x02, 0x00, 0x00,
}
