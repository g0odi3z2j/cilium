// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/ads.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"

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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type AdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdsDummy) Reset()         { *m = AdsDummy{} }
func (m *AdsDummy) String() string { return proto.CompactTextString(m) }
func (*AdsDummy) ProtoMessage()    {}
func (*AdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_ads_0bf0bd91837d02c3, []int{0}
}
func (m *AdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdsDummy.Unmarshal(m, b)
}
func (m *AdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdsDummy.Marshal(b, m, deterministic)
}
func (dst *AdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdsDummy.Merge(dst, src)
}
func (m *AdsDummy) XXX_Size() int {
	return xxx_messageInfo_AdsDummy.Size(m)
}
func (m *AdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_AdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_AdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v2.AdsDummy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedDiscoveryServiceClient(cc *grpc.ClientConn) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/ads.proto", fileDescriptor_ads_0bf0bd91837d02c3)
}

var fileDescriptor_ads_0bf0bd91837d02c3 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0xb1, 0xae, 0x82, 0x30,
	0x18, 0x05, 0xe0, 0xdb, 0x3b, 0xa8, 0xe9, 0xc8, 0x66, 0x63, 0x34, 0x21, 0x0e, 0x4e, 0xad, 0xa9,
	0x4f, 0x80, 0xe1, 0x09, 0x60, 0x73, 0x2b, 0xf4, 0x0f, 0xe9, 0x00, 0xad, 0xfd, 0x4b, 0x13, 0x36,
	0x47, 0x1f, 0xdb, 0x08, 0x0a, 0x93, 0xf3, 0xf9, 0xce, 0xc9, 0xa1, 0x47, 0xe8, 0xa2, 0x1d, 0x04,
	0x82, 0x8f, 0xa6, 0x06, 0xa1, 0x0d, 0xd6, 0x36, 0x82, 0x1f, 0x44, 0x94, 0x42, 0x69, 0xe4, 0xce,
	0xdb, 0x60, 0x13, 0x36, 0x2a, 0xfe, 0x51, 0x7c, 0x56, 0x3c, 0x4a, 0xb6, 0x9b, 0x16, 0x94, 0x33,
	0xef, 0xce, 0x12, 0x8d, 0xcd, 0x94, 0xd2, 0x4d, 0xa6, 0x31, 0xef, 0xdb, 0x76, 0x90, 0x0f, 0x42,
	0x59, 0xd6, 0x34, 0x1e, 0x1a, 0x15, 0x40, 0xe7, 0x5f, 0x59, 0x4e, 0xab, 0x49, 0x45, 0xb7, 0x65,
	0xf0, 0xa0, 0xda, 0xc5, 0x14, 0x80, 0xb6, 0xf7, 0x35, 0x60, 0xb2, 0xe7, 0xd3, 0x05, 0xe5, 0x0c,
	0x8f, 0x92, 0xcf, 0xe5, 0x02, 0xee, 0x3d, 0x60, 0x60, 0x87, 0x9f, 0x39, 0x3a, 0xdb, 0x21, 0xa4,
	0x7f, 0x27, 0x72, 0x26, 0xd7, 0xf5, 0xed, 0x3f, 0xca, 0x27, 0x21, 0xd5, 0x6a, 0xbc, 0x77, 0x79,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x44, 0x0f, 0x59, 0xd2, 0x00, 0x01, 0x00, 0x00,
}
