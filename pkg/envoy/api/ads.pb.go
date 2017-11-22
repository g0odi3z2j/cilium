// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/ads.proto

package envoy_api_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AggregatedDiscoveryService service

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
	stream, err := grpc.NewClientStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AggregatedDiscoveryService service

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
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.AggregatedDiscoveryService",
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
	Metadata: "api/ads.proto",
}

func init() { proto.RegisterFile("api/ads.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0xb1, 0xaa, 0xc2, 0x30,
	0x14, 0xc6, 0xf1, 0xdb, 0xe5, 0x0e, 0x45, 0x1d, 0xba, 0x19, 0x44, 0xc1, 0xc9, 0x29, 0x95, 0xfa,
	0x04, 0x82, 0x4f, 0xd0, 0x3e, 0xc1, 0x69, 0xfb, 0x11, 0x02, 0x9a, 0x13, 0x73, 0xd2, 0x40, 0x37,
	0x1f, 0x5d, 0xac, 0xa2, 0x2e, 0xce, 0xff, 0xf3, 0xe3, 0x7c, 0xf9, 0x9c, 0xbc, 0x2d, 0xa9, 0x17,
	0xed, 0x03, 0x47, 0x2e, 0x66, 0x70, 0x89, 0x47, 0x4d, 0xde, 0xea, 0x54, 0xa9, 0xc5, 0x23, 0xb6,
	0x24, 0x78, 0x56, 0xb5, 0x32, 0xcc, 0xe6, 0x8c, 0x72, 0x32, 0xce, 0x71, 0xa4, 0x68, 0xd9, 0xbd,
	0x6c, 0x75, 0xcb, 0x72, 0x75, 0x34, 0x26, 0xc0, 0x50, 0x44, 0x7f, 0xb2, 0xd2, 0x71, 0x42, 0x18,
	0x1b, 0x84, 0x64, 0x3b, 0x14, 0x6d, 0xbe, 0x6c, 0x62, 0x00, 0x5d, 0x3e, 0x37, 0x35, 0x84, 0x87,
	0xd0, 0x41, 0x8a, 0xb5, 0xfe, 0x7e, 0xac, 0xdf, 0xb8, 0xc6, 0x75, 0x80, 0x44, 0xb5, 0xf9, 0xd9,
	0xc5, 0xb3, 0x13, 0x6c, 0xff, 0x76, 0xd9, 0x3e, 0x6b, 0xff, 0xa7, 0x25, 0x87, 0x7b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xab, 0xc1, 0x9b, 0xd6, 0x00, 0x00, 0x00,
}
