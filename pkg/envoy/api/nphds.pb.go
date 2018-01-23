// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/nphds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The mapping of a network policy identifier to the IP addresses of all the
// hosts on which the network policy is enforced.
// A host may be associated only with one network policy.
type NetworkPolicyHosts struct {
	// The unique identifier of the network policy enforced on the hosts.
	Policy uint64 `protobuf:"varint,1,opt,name=policy" json:"policy,omitempty"`
	// The set of IP addresses of the hosts on which the network policy is enforced.
	// Optional. May be empty.
	HostAddresses []string `protobuf:"bytes,2,rep,name=host_addresses,json=hostAddresses" json:"host_addresses,omitempty"`
}

func (m *NetworkPolicyHosts) Reset()                    { *m = NetworkPolicyHosts{} }
func (m *NetworkPolicyHosts) String() string            { return proto.CompactTextString(m) }
func (*NetworkPolicyHosts) ProtoMessage()               {}
func (*NetworkPolicyHosts) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *NetworkPolicyHosts) GetPolicy() uint64 {
	if m != nil {
		return m.Policy
	}
	return 0
}

func (m *NetworkPolicyHosts) GetHostAddresses() []string {
	if m != nil {
		return m.HostAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkPolicyHosts)(nil), "envoy.api.v2.NetworkPolicyHosts")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkPolicyHostsDiscoveryService service

type NetworkPolicyHostsDiscoveryServiceClient interface {
	StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error)
	FetchNetworkPolicyHosts(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type networkPolicyHostsDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyHostsDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyHostsDiscoveryServiceClient {
	return &networkPolicyHostsDiscoveryServiceClient{cc}
}

func (c *networkPolicyHostsDiscoveryServiceClient) StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NetworkPolicyHostsDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.NetworkPolicyHostsDiscoveryService/StreamNetworkPolicyHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient{stream}
	return x, nil
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient struct {
	grpc.ClientStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyHostsDiscoveryServiceClient) FetchNetworkPolicyHosts(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkPolicyHostsDiscoveryService service

type NetworkPolicyHostsDiscoveryServiceServer interface {
	StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error
	FetchNetworkPolicyHosts(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterNetworkPolicyHostsDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyHostsDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyHostsDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyHostsDiscoveryServiceServer).StreamNetworkPolicyHosts(&networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer{stream})
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer struct {
	grpc.ServerStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyHostsDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.NetworkPolicyHostsDiscoveryService",
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
	Metadata: "api/nphds.proto",
}

func init() { proto.RegisterFile("api/nphds.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0xc8, 0xd4,
	0xcf, 0x2b, 0xc8, 0x48, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xcd, 0x2b,
	0xcb, 0xaf, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x92, 0x12, 0x06, 0x49, 0xa7, 0x64, 0x16,
	0x27, 0xe7, 0x97, 0xa5, 0x16, 0x55, 0x42, 0x94, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4,
	0xea, 0x83, 0xe4, 0x12, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xa0, 0x06, 0x48,
	0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x5c,
	0x2e, 0x21, 0xbf, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0xec, 0x80, 0xfc, 0x9c, 0xcc, 0xe4, 0x4a, 0x8f,
	0xfc, 0xe2, 0x92, 0x62, 0x21, 0x31, 0x2e, 0xb6, 0x02, 0x30, 0x57, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x25, 0x08, 0xca, 0x13, 0xb2, 0xe7, 0xe2, 0xcb, 0xc8, 0x2f, 0x2e, 0x89, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x4e, 0x2d, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x74, 0x92, 0xd8, 0xf5, 0xf2,
	0x00, 0x33, 0xeb, 0x24, 0x46, 0x26, 0x09, 0x46, 0x10, 0x8b, 0x73, 0x12, 0x23, 0x9b, 0x12, 0x4b,
	0x11, 0x93, 0x02, 0x63, 0x10, 0x2f, 0x48, 0xbd, 0x23, 0x4c, 0xb9, 0xd1, 0x64, 0x26, 0x2e, 0x25,
	0x4c, 0xfb, 0x5c, 0x60, 0x7e, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x4a, 0xe4, 0x92,
	0x08, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xc5, 0xe2, 0x36, 0x39, 0x3d, 0xe4, 0xc0, 0xd0, 0x83, 0xeb,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x92, 0xc7, 0x29, 0x5f, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0xaa, 0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0x28, 0xd4, 0xc9, 0xc8, 0x25, 0xee, 0x96, 0x5a, 0x92,
	0x9c, 0x41, 0x0b, 0x2b, 0x74, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0xae, 0xa4, 0xa4, 0x5f, 0x66,
	0x84, 0x88, 0x29, 0xab, 0x3c, 0x88, 0x55, 0xf1, 0x90, 0xd0, 0x8c, 0x07, 0x05, 0x4d, 0xb1, 0x15,
	0xa3, 0x96, 0x13, 0x6b, 0x14, 0x73, 0x62, 0x41, 0x66, 0x12, 0x1b, 0x38, 0x4a, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x83, 0x57, 0x1f, 0x67, 0xff, 0x01, 0x00, 0x00,
}
