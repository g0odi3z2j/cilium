// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.12
// source: envoy/service/cluster/v3/cds.proto

package clusterv3

import (
	context "context"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type CdsDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CdsDummy) Reset() {
	*x = CdsDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_cluster_v3_cds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CdsDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CdsDummy) ProtoMessage() {}

func (x *CdsDummy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_cluster_v3_cds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CdsDummy.ProtoReflect.Descriptor instead.
func (*CdsDummy) Descriptor() ([]byte, []int) {
	return file_envoy_service_cluster_v3_cds_proto_rawDescGZIP(), []int{0}
}

var File_envoy_service_cluster_v3_cds_proto protoreflect.FileDescriptor

var file_envoy_service_cluster_v3_cds_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x2a,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x08,
	0x43, 0x64, 0x73, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x3a, 0x1c, 0x9a, 0xc5, 0x88, 0x1e, 0x17, 0x0a,
	0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x64,
	0x73, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0xcd, 0x03, 0x0a, 0x17, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x95, 0x01, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x16, 0x2f, 0x76,
	0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x1a, 0x27, 0x8a,
	0xa4, 0x96, 0xf3, 0x07, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x8a, 0x01, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x42, 0x08, 0x43, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x76, 0x33, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_cluster_v3_cds_proto_rawDescOnce sync.Once
	file_envoy_service_cluster_v3_cds_proto_rawDescData = file_envoy_service_cluster_v3_cds_proto_rawDesc
)

func file_envoy_service_cluster_v3_cds_proto_rawDescGZIP() []byte {
	file_envoy_service_cluster_v3_cds_proto_rawDescOnce.Do(func() {
		file_envoy_service_cluster_v3_cds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_cluster_v3_cds_proto_rawDescData)
	})
	return file_envoy_service_cluster_v3_cds_proto_rawDescData
}

var file_envoy_service_cluster_v3_cds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_service_cluster_v3_cds_proto_goTypes = []interface{}{
	(*CdsDummy)(nil),                  // 0: envoy.service.cluster.v3.CdsDummy
	(*v3.DiscoveryRequest)(nil),       // 1: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DeltaDiscoveryRequest)(nil),  // 2: envoy.service.discovery.v3.DeltaDiscoveryRequest
	(*v3.DiscoveryResponse)(nil),      // 3: envoy.service.discovery.v3.DiscoveryResponse
	(*v3.DeltaDiscoveryResponse)(nil), // 4: envoy.service.discovery.v3.DeltaDiscoveryResponse
}
var file_envoy_service_cluster_v3_cds_proto_depIdxs = []int32{
	1, // 0: envoy.service.cluster.v3.ClusterDiscoveryService.StreamClusters:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	2, // 1: envoy.service.cluster.v3.ClusterDiscoveryService.DeltaClusters:input_type -> envoy.service.discovery.v3.DeltaDiscoveryRequest
	1, // 2: envoy.service.cluster.v3.ClusterDiscoveryService.FetchClusters:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	3, // 3: envoy.service.cluster.v3.ClusterDiscoveryService.StreamClusters:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	4, // 4: envoy.service.cluster.v3.ClusterDiscoveryService.DeltaClusters:output_type -> envoy.service.discovery.v3.DeltaDiscoveryResponse
	3, // 5: envoy.service.cluster.v3.ClusterDiscoveryService.FetchClusters:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_service_cluster_v3_cds_proto_init() }
func file_envoy_service_cluster_v3_cds_proto_init() {
	if File_envoy_service_cluster_v3_cds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_cluster_v3_cds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CdsDummy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_cluster_v3_cds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_cluster_v3_cds_proto_goTypes,
		DependencyIndexes: file_envoy_service_cluster_v3_cds_proto_depIdxs,
		MessageInfos:      file_envoy_service_cluster_v3_cds_proto_msgTypes,
	}.Build()
	File_envoy_service_cluster_v3_cds_proto = out.File
	file_envoy_service_cluster_v3_cds_proto_rawDesc = nil
	file_envoy_service_cluster_v3_cds_proto_goTypes = nil
	file_envoy_service_cluster_v3_cds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error)
	DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error)
	FetchClusters(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterDiscoveryServiceClient(cc grpc.ClientConnInterface) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[0], "/envoy.service.cluster.v3.ClusterDiscoveryService/StreamClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceStreamClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_StreamClustersClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceStreamClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceStreamClustersClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[1], "/envoy.service.cluster.v3.ClusterDiscoveryService/DeltaClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceDeltaClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_DeltaClustersClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceDeltaClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) FetchClusters(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.cluster.v3.ClusterDiscoveryService/FetchClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterDiscoveryServiceServer is the server API for ClusterDiscoveryService service.
type ClusterDiscoveryServiceServer interface {
	StreamClusters(ClusterDiscoveryService_StreamClustersServer) error
	DeltaClusters(ClusterDiscoveryService_DeltaClustersServer) error
	FetchClusters(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedClusterDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterDiscoveryServiceServer struct {
}

func (*UnimplementedClusterDiscoveryServiceServer) StreamClusters(ClusterDiscoveryService_StreamClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) DeltaClusters(ClusterDiscoveryService_DeltaClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) FetchClusters(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClusters not implemented")
}

func RegisterClusterDiscoveryServiceServer(s *grpc.Server, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&_ClusterDiscoveryService_serviceDesc, srv)
}

func _ClusterDiscoveryService_StreamClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).StreamClusters(&clusterDiscoveryServiceStreamClustersServer{stream})
}

type ClusterDiscoveryService_StreamClustersServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceStreamClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceStreamClustersServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_DeltaClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).DeltaClusters(&clusterDiscoveryServiceDeltaClustersServer{stream})
}

type ClusterDiscoveryService_DeltaClustersServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceDeltaClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_FetchClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.cluster.v3.ClusterDiscoveryService/FetchClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.cluster.v3.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClusters",
			Handler:    _ClusterDiscoveryService_FetchClusters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClusters",
			Handler:       _ClusterDiscoveryService_StreamClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaClusters",
			Handler:       _ClusterDiscoveryService_DeltaClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/cluster/v3/cds.proto",
}
