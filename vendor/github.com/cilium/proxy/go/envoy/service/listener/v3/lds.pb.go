// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/listener/v3/lds.proto

package envoy_service_listener_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type LdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdsDummy) Reset()         { *m = LdsDummy{} }
func (m *LdsDummy) String() string { return proto.CompactTextString(m) }
func (*LdsDummy) ProtoMessage()    {}
func (*LdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4f7cdedc179239, []int{0}
}

func (m *LdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdsDummy.Unmarshal(m, b)
}
func (m *LdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdsDummy.Marshal(b, m, deterministic)
}
func (m *LdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdsDummy.Merge(m, src)
}
func (m *LdsDummy) XXX_Size() int {
	return xxx_messageInfo_LdsDummy.Size(m)
}
func (m *LdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_LdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_LdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LdsDummy)(nil), "envoy.service.listener.v3.LdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/listener/v3/lds.proto", fileDescriptor_2f4f7cdedc179239)
}

var fileDescriptor_2f4f7cdedc179239 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0xca, 0xd3, 0x40,
	0x14, 0x85, 0x4d, 0x7f, 0xa8, 0x32, 0x8b, 0x2a, 0x01, 0xa9, 0x0d, 0xb5, 0xd8, 0x76, 0x61, 0x2d,
	0x3a, 0xa3, 0xcd, 0xae, 0x2b, 0x29, 0xc5, 0x55, 0x17, 0xc5, 0x3e, 0xc1, 0x34, 0xb9, 0x8d, 0x03,
	0xe9, 0xcc, 0x38, 0x33, 0x19, 0xcd, 0x42, 0x10, 0x57, 0xc5, 0xad, 0xa0, 0xe0, 0xda, 0xa7, 0x70,
	0x2f, 0xb8, 0x15, 0x7d, 0x04, 0xdf, 0xc0, 0x17, 0x90, 0x26, 0x99, 0xb6, 0x51, 0x2a, 0xba, 0x71,
	0x97, 0xe1, 0x7c, 0xf7, 0xdc, 0xcb, 0xe1, 0x04, 0x0d, 0x81, 0x5b, 0x91, 0x13, 0x0d, 0xca, 0xb2,
	0x08, 0x48, 0xca, 0xb4, 0x01, 0x0e, 0x8a, 0xd8, 0x90, 0xa4, 0xb1, 0xc6, 0x52, 0x09, 0x23, 0xfc,
	0x4e, 0x01, 0xe1, 0x0a, 0xc2, 0x0e, 0xc2, 0x36, 0x0c, 0xc6, 0xf5, 0xf9, 0x98, 0xe9, 0x48, 0x58,
	0x50, 0xf9, 0xde, 0xe0, 0xf0, 0x28, 0x6d, 0x82, 0x6e, 0x22, 0x44, 0x92, 0x02, 0xa1, 0x92, 0x11,
	0xca, 0xb9, 0x30, 0xd4, 0x30, 0xc1, 0xab, 0x25, 0x41, 0xaf, 0x52, 0x8b, 0xd7, 0x3a, 0xdb, 0x90,
	0x38, 0x53, 0x05, 0x70, 0x4e, 0x7f, 0xa6, 0xa8, 0x94, 0xa0, 0xdc, 0xfc, 0xad, 0xf2, 0x92, 0x13,
	0x63, 0xa2, 0x40, 0x8b, 0x4c, 0x45, 0x50, 0x11, 0x37, 0xb3, 0x58, 0xd2, 0x1a, 0xa0, 0x0d, 0x35,
	0x99, 0x33, 0xe8, 0xff, 0x26, 0x5b, 0x50, 0x9a, 0x09, 0xce, 0x78, 0x52, 0x21, 0x6d, 0x4b, 0x53,
	0x16, 0x53, 0x03, 0xc4, 0x7d, 0x94, 0xc2, 0x60, 0x84, 0xae, 0x2c, 0x62, 0x3d, 0xcf, 0xb6, 0xdb,
	0x7c, 0xda, 0x7d, 0xff, 0x69, 0xd7, 0x6b, 0xa3, 0xeb, 0x65, 0x68, 0x54, 0x32, 0x6c, 0x27, 0xd8,
	0xa9, 0x93, 0x6f, 0x17, 0xe8, 0xc6, 0xa2, 0x0a, 0x70, 0xee, 0x02, 0x5a, 0x95, 0xf1, 0xf9, 0x2f,
	0x50, 0x6b, 0x0e, 0xa9, 0xa1, 0x0e, 0xd0, 0xfe, 0x03, 0x5c, 0xcf, 0xfe, 0x98, 0xa9, 0x0d, 0x71,
	0xc1, 0x1e, 0x4c, 0x1e, 0xc3, 0xd3, 0x0c, 0xb4, 0x09, 0x26, 0xff, 0x32, 0xa2, 0xa5, 0xe0, 0x1a,
	0x06, 0x97, 0x46, 0xde, 0x7d, 0xcf, 0x37, 0xe8, 0xea, 0xca, 0x28, 0xa0, 0xdb, 0xe3, 0xfe, 0xbb,
	0x7f, 0x34, 0xfb, 0x75, 0xf5, 0xbd, 0xbf, 0xa4, 0x6b, 0x5b, 0xdf, 0x79, 0xa8, 0xf5, 0x08, 0x4c,
	0xf4, 0xe4, 0x3f, 0x6d, 0x1d, 0xbd, 0xfa, 0xfa, 0xfd, 0x4d, 0xa3, 0x33, 0x68, 0xd7, 0x3a, 0x3a,
	0x75, 0xad, 0xd6, 0x85, 0x7c, 0x31, 0xf5, 0xc6, 0xc1, 0x9d, 0xd7, 0x1f, 0xde, 0xfe, 0xb8, 0x3c,
	0x44, 0xfd, 0xd2, 0x3f, 0x12, 0x7c, 0xc3, 0x92, 0xd3, 0xfa, 0x63, 0x77, 0xf2, 0xec, 0xe1, 0xc7,
	0x97, 0x9f, 0xbf, 0x34, 0x1b, 0xd7, 0x1a, 0xe8, 0x36, 0x13, 0xe5, 0x3d, 0x52, 0x89, 0xe7, 0x39,
	0x3e, 0xfb, 0xeb, 0xcc, 0xf6, 0x8d, 0x59, 0xee, 0xdb, 0xb3, 0xf4, 0x76, 0x9e, 0xb7, 0x6e, 0x16,
	0x4d, 0x0a, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x89, 0x77, 0x18, 0x92, 0x03, 0x00, 0x00,
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
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.service.listener.v3.ListenerDiscoveryService/DeltaListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[1], "/envoy.service.listener.v3.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.listener.v3.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (*UnimplementedListenerDiscoveryServiceServer) DeltaListeners(srv ListenerDiscoveryService_DeltaListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) StreamListeners(srv ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) FetchListeners(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.listener.v3.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.listener.v3.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/listener/v3/lds.proto",
}
