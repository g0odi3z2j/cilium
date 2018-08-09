// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// A load report Envoy sends to the management server.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsRequest struct {
	// Node identifier for Envoy instance.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of load stats to report.
	ClusterStats         []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_34d443be72931312, []int{0}
}
func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (dst *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(dst, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

// The management server sends envoy a LoadStatsResponse with all clusters it
// is interested in learning load stats about.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsResponse struct {
	// Clusters to report stats for.
	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The minimum interval of time to collect stats over. This is only a minimum for two reasons:
	// 1. There may be some delay from when the timer fires until stats sampling occurs.
	// 2. For clusters that were already feature in the previous *LoadStatsResponse*, any traffic
	//    that is observed in between the corresponding previous *LoadStatsRequest* and this
	//    *LoadStatsResponse* will also be accumulated and billed to the cluster. This avoids a period
	//    of inobservability that might otherwise exists between the messages. New clusters are not
	//    subject to this consideration.
	LoadReportingInterval *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_34d443be72931312, []int{1}
}
func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (dst *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(dst, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v2.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v2.LoadStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v2/lrs.proto", fileDescriptor_lrs_34d443be72931312)
}

var fileDescriptor_lrs_34d443be72931312 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4e, 0xe3, 0x30,
	0x1c, 0xc5, 0xc7, 0xe9, 0x7c, 0x74, 0xdc, 0x19, 0xcd, 0x4c, 0x34, 0xa3, 0x66, 0x0a, 0x42, 0x55,
	0x11, 0x10, 0x09, 0x61, 0xa3, 0x70, 0x83, 0xc2, 0x02, 0xa4, 0x0a, 0x89, 0x74, 0xc7, 0xa6, 0x72,
	0x93, 0x3f, 0x95, 0xa5, 0x60, 0x07, 0xdb, 0xb1, 0xc4, 0x0d, 0x60, 0xd3, 0x05, 0xc7, 0x61, 0xc5,
	0x75, 0xb8, 0x05, 0x4a, 0x9c, 0x94, 0x94, 0x05, 0x62, 0x17, 0xeb, 0xfd, 0x5e, 0xfc, 0xde, 0x33,
	0xde, 0x01, 0x61, 0xe5, 0x2d, 0xd5, 0xa0, 0x2c, 0x4f, 0x80, 0x66, 0x92, 0xa5, 0x33, 0x6d, 0x98,
	0xd1, 0xd4, 0x46, 0x34, 0x53, 0x9a, 0xe4, 0x4a, 0x1a, 0xe9, 0x6f, 0x54, 0x18, 0xa9, 0x31, 0xf2,
	0x8a, 0x11, 0x1b, 0x0d, 0x36, 0xdd, 0x3f, 0x58, 0xce, 0x4b, 0x53, 0x22, 0x15, 0xd0, 0x39, 0xd3,
	0xe0, 0xac, 0x83, 0xbd, 0x35, 0x15, 0x44, 0x9a, 0x4b, 0x2e, 0x8c, 0xbb, 0x49, 0x41, 0x2e, 0x95,
	0xa9, 0xc1, 0xad, 0x85, 0x94, 0x8b, 0x0c, 0x68, 0x75, 0x9a, 0x17, 0x57, 0x34, 0x2d, 0x14, 0x33,
	0x5c, 0x8a, 0x5a, 0xef, 0x5b, 0x96, 0xf1, 0x94, 0x19, 0xa0, 0xcd, 0x87, 0x13, 0x46, 0xf7, 0x08,
	0xff, 0x9e, 0x48, 0x96, 0x4e, 0xcb, 0x40, 0x31, 0xdc, 0x14, 0xa0, 0x8d, 0xbf, 0x8f, 0x3f, 0x0b,
	0x99, 0x42, 0x80, 0x86, 0x28, 0xec, 0x45, 0x7d, 0xe2, 0x0a, 0xb0, 0x9c, 0x13, 0x1b, 0x91, 0x32,
	0x23, 0x39, 0x97, 0x29, 0xc4, 0x15, 0xe4, 0x9f, 0xe2, 0x9f, 0x49, 0x56, 0x68, 0x03, 0xca, 0xb5,
	0x0a, 0xbc, 0x61, 0x27, 0xec, 0x45, 0xdb, 0xeb, 0xae, 0x26, 0x3b, 0x39, 0x76, 0xac, 0xbb, 0xef,
	0x47, 0xd2, 0x3a, 0x8d, 0x96, 0x08, 0xff, 0x69, 0x65, 0xd1, 0xb9, 0x14, 0x1a, 0xfc, 0x5d, 0xdc,
	0xad, 0x29, 0x1d, 0xa0, 0x61, 0x27, 0xfc, 0x3e, 0xc6, 0x8f, 0xcf, 0x4f, 0x9d, 0x2f, 0x0f, 0xc8,
	0xeb, 0xa2, 0x78, 0xa5, 0xf9, 0x17, 0xb8, 0xdf, 0xda, 0x85, 0x8b, 0xc5, 0x8c, 0x0b, 0x03, 0xca,
	0xb2, 0x2c, 0xf0, 0xaa, 0x1e, 0xff, 0x89, 0x1b, 0x89, 0x34, 0x23, 0x91, 0x93, 0x7a, 0xa4, 0xf8,
	0x5f, 0xe9, 0x8c, 0x1b, 0xe3, 0x59, 0xed, 0x8b, 0x96, 0x08, 0xff, 0x9d, 0xb4, 0x95, 0xa9, 0x7b,
	0x43, 0xdf, 0xe2, 0x5f, 0x53, 0xa3, 0x80, 0x5d, 0xaf, 0xe2, 0xfa, 0x07, 0xe4, 0x9d, 0x67, 0x26,
	0x6f, 0x27, 0x1e, 0x90, 0x8f, 0xe2, 0x6e, 0x85, 0xd1, 0xa7, 0x10, 0x1d, 0xa2, 0xf1, 0xb7, 0x4b,
	0xcf, 0x46, 0x77, 0x08, 0xcd, 0xbf, 0x56, 0x1d, 0x8e, 0x5e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x59, 0x2c, 0xb2, 0x83, 0x02, 0x00, 0x00,
}
