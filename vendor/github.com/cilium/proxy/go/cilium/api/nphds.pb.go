// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cilium/api/nphds.proto

package cilium

import (
	context "context"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// The mapping of a network policy identifier to the IP addresses of all the
// hosts on which the network policy is enforced.
// A host may be associated only with one network policy.
type NetworkPolicyHosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the network policy enforced on the hosts.
	Policy uint64 `protobuf:"varint,1,opt,name=policy,proto3" json:"policy,omitempty"`
	// The set of IP addresses of the hosts on which the network policy is
	// enforced. Optional. May be empty.
	HostAddresses []string `protobuf:"bytes,2,rep,name=host_addresses,json=hostAddresses,proto3" json:"host_addresses,omitempty"`
}

func (x *NetworkPolicyHosts) Reset() {
	*x = NetworkPolicyHosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_nphds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkPolicyHosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkPolicyHosts) ProtoMessage() {}

func (x *NetworkPolicyHosts) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_nphds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkPolicyHosts.ProtoReflect.Descriptor instead.
func (*NetworkPolicyHosts) Descriptor() ([]byte, []int) {
	return file_cilium_api_nphds_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkPolicyHosts) GetPolicy() uint64 {
	if x != nil {
		return x.Policy
	}
	return 0
}

func (x *NetworkPolicyHosts) GetHostAddresses() []string {
	if x != nil {
		return x.HostAddresses
	}
	return nil
}

var File_cilium_api_nphds_proto protoreflect.FileDescriptor

var file_cilium_api_nphds_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x70, 0x68,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d,
	0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x0e, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x14, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x18, 0x01, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x20,
	0x01, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x32, 0xee, 0x02, 0x0a, 0x22, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0xa5, 0x01, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x21,
	0x8a, 0xa4, 0x96, 0xf3, 0x07, 0x1b, 0x0a, 0x19, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_nphds_proto_rawDescOnce sync.Once
	file_cilium_api_nphds_proto_rawDescData = file_cilium_api_nphds_proto_rawDesc
)

func file_cilium_api_nphds_proto_rawDescGZIP() []byte {
	file_cilium_api_nphds_proto_rawDescOnce.Do(func() {
		file_cilium_api_nphds_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_nphds_proto_rawDescData)
	})
	return file_cilium_api_nphds_proto_rawDescData
}

var file_cilium_api_nphds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cilium_api_nphds_proto_goTypes = []interface{}{
	(*NetworkPolicyHosts)(nil),   // 0: cilium.NetworkPolicyHosts
	(*v3.DiscoveryRequest)(nil),  // 1: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 2: envoy.service.discovery.v3.DiscoveryResponse
}
var file_cilium_api_nphds_proto_depIdxs = []int32{
	1, // 0: cilium.NetworkPolicyHostsDiscoveryService.StreamNetworkPolicyHosts:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	1, // 1: cilium.NetworkPolicyHostsDiscoveryService.FetchNetworkPolicyHosts:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	2, // 2: cilium.NetworkPolicyHostsDiscoveryService.StreamNetworkPolicyHosts:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	2, // 3: cilium.NetworkPolicyHostsDiscoveryService.FetchNetworkPolicyHosts:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cilium_api_nphds_proto_init() }
func file_cilium_api_nphds_proto_init() {
	if File_cilium_api_nphds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_nphds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkPolicyHosts); i {
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
			RawDescriptor: file_cilium_api_nphds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cilium_api_nphds_proto_goTypes,
		DependencyIndexes: file_cilium_api_nphds_proto_depIdxs,
		MessageInfos:      file_cilium_api_nphds_proto_msgTypes,
	}.Build()
	File_cilium_api_nphds_proto = out.File
	file_cilium_api_nphds_proto_rawDesc = nil
	file_cilium_api_nphds_proto_goTypes = nil
	file_cilium_api_nphds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetworkPolicyHostsDiscoveryServiceClient is the client API for NetworkPolicyHostsDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkPolicyHostsDiscoveryServiceClient interface {
	StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error)
	FetchNetworkPolicyHosts(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type networkPolicyHostsDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkPolicyHostsDiscoveryServiceClient(cc grpc.ClientConnInterface) NetworkPolicyHostsDiscoveryServiceClient {
	return &networkPolicyHostsDiscoveryServiceClient{cc}
}

func (c *networkPolicyHostsDiscoveryServiceClient) StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkPolicyHostsDiscoveryService_serviceDesc.Streams[0], "/cilium.NetworkPolicyHostsDiscoveryService/StreamNetworkPolicyHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient{stream}
	return x, nil
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient struct {
	grpc.ClientStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyHostsDiscoveryServiceClient) FetchNetworkPolicyHosts(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyHostsDiscoveryServiceServer is the server API for NetworkPolicyHostsDiscoveryService service.
type NetworkPolicyHostsDiscoveryServiceServer interface {
	StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error
	FetchNetworkPolicyHosts(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedNetworkPolicyHostsDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkPolicyHostsDiscoveryServiceServer struct {
}

func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNetworkPolicyHosts not implemented")
}
func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) FetchNetworkPolicyHosts(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchNetworkPolicyHosts not implemented")
}

func RegisterNetworkPolicyHostsDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyHostsDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyHostsDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyHostsDiscoveryServiceServer).StreamNetworkPolicyHosts(&networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer{stream})
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer struct {
	grpc.ServerStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyHostsDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cilium.NetworkPolicyHostsDiscoveryService",
	HandlerType: (*NetworkPolicyHostsDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchNetworkPolicyHosts",
			Handler:    _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkPolicyHosts",
			Handler:       _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cilium/api/nphds.proto",
}
