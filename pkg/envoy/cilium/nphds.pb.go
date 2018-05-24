// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/nphds.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
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

// The mapping of a network policy identifier to the IP addresses of all the
// hosts on which the network policy is enforced.
// A host may be associated only with one network policy.
type NetworkPolicyHosts struct {
	// The unique identifier of the network policy enforced on the hosts.
	Policy uint64 `protobuf:"varint,1,opt,name=policy" json:"policy,omitempty"`
	// The set of IP addresses of the hosts on which the network policy is enforced.
	// Optional. May be empty.
	HostAddresses        []string `protobuf:"bytes,2,rep,name=host_addresses,json=hostAddresses" json:"host_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPolicyHosts) Reset()         { *m = NetworkPolicyHosts{} }
func (m *NetworkPolicyHosts) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicyHosts) ProtoMessage()    {}
func (*NetworkPolicyHosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_nphds_9150389911c00044, []int{0}
}
func (m *NetworkPolicyHosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPolicyHosts.Unmarshal(m, b)
}
func (m *NetworkPolicyHosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPolicyHosts.Marshal(b, m, deterministic)
}
func (dst *NetworkPolicyHosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicyHosts.Merge(dst, src)
}
func (m *NetworkPolicyHosts) XXX_Size() int {
	return xxx_messageInfo_NetworkPolicyHosts.Size(m)
}
func (m *NetworkPolicyHosts) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicyHosts.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicyHosts proto.InternalMessageInfo

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
	proto.RegisterType((*NetworkPolicyHosts)(nil), "cilium.NetworkPolicyHosts")
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
	FetchNetworkPolicyHosts(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type networkPolicyHostsDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyHostsDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyHostsDiscoveryServiceClient {
	return &networkPolicyHostsDiscoveryServiceClient{cc}
}

func (c *networkPolicyHostsDiscoveryServiceClient) StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NetworkPolicyHostsDiscoveryService_serviceDesc.Streams[0], c.cc, "/cilium.NetworkPolicyHostsDiscoveryService/StreamNetworkPolicyHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient{stream}
	return x, nil
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient struct {
	grpc.ClientStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyHostsDiscoveryServiceClient) FetchNetworkPolicyHosts(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := grpc.Invoke(ctx, "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkPolicyHostsDiscoveryService service

type NetworkPolicyHostsDiscoveryServiceServer interface {
	StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error
	FetchNetworkPolicyHosts(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterNetworkPolicyHostsDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyHostsDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyHostsDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyHostsDiscoveryServiceServer).StreamNetworkPolicyHosts(&networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer{stream})
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer struct {
	grpc.ServerStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
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
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, req.(*v2.DiscoveryRequest))
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
	Metadata: "cilium/nphds.proto",
}

func init() { proto.RegisterFile("cilium/nphds.proto", fileDescriptor_nphds_9150389911c00044) }

var fileDescriptor_nphds_9150389911c00044 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x31, 0x4f, 0xf3, 0x30,
	0x14, 0xfc, 0x9c, 0xf6, 0x8b, 0xa8, 0x25, 0x18, 0x3c, 0xd0, 0x28, 0xaa, 0xa0, 0xca, 0x42, 0x85,
	0x44, 0x82, 0xca, 0xd6, 0x05, 0x51, 0x21, 0xc4, 0x84, 0x50, 0xbb, 0xb1, 0x54, 0x26, 0x79, 0x6a,
	0x2d, 0x52, 0xbf, 0xe0, 0xe7, 0x1a, 0x75, 0x65, 0x63, 0x6d, 0xff, 0x0d, 0x13, 0xff, 0x81, 0xbf,
	0xc0, 0xc2, 0xbf, 0x40, 0x4d, 0x5a, 0x24, 0x04, 0x8c, 0x6c, 0xef, 0x7c, 0x67, 0x9f, 0xdf, 0x1d,
	0x17, 0xa9, 0xca, 0xd5, 0x6c, 0x9a, 0xe8, 0x62, 0x92, 0x51, 0x5c, 0x18, 0xb4, 0x28, 0xfc, 0xea,
	0x2c, 0x6c, 0x81, 0x76, 0x38, 0x4f, 0x64, 0xa1, 0x12, 0xd7, 0x4d, 0x32, 0x45, 0x29, 0x3a, 0x30,
	0xf3, 0x4a, 0x15, 0xb6, 0xc6, 0x88, 0xe3, 0x1c, 0x4a, 0x5a, 0x6a, 0x8d, 0x56, 0x5a, 0x85, 0x7a,
	0xfd, 0x46, 0xd8, 0x74, 0x32, 0x57, 0x99, 0xb4, 0x90, 0x6c, 0x86, 0x8a, 0x88, 0xa6, 0x5c, 0x5c,
	0x81, 0x7d, 0x40, 0x73, 0x77, 0x8d, 0xb9, 0x4a, 0xe7, 0x97, 0x48, 0x96, 0xc4, 0x2e, 0xf7, 0x8b,
	0x12, 0x06, 0xac, 0xcd, 0x3a, 0xf5, 0xc1, 0x1a, 0x89, 0x53, 0xbe, 0x33, 0x41, 0xb2, 0x23, 0x99,
	0x65, 0x06, 0x88, 0x80, 0x02, 0xaf, 0x5d, 0xeb, 0x34, 0xfa, 0xc1, 0xf3, 0xfb, 0x4b, 0xed, 0xff,
	0x82, 0x79, 0x01, 0x5b, 0x4d, 0x8d, 0x05, 0xf3, 0xa3, 0xba, 0xf1, 0xda, 0x6c, 0xb0, 0xbd, 0xd2,
	0x9f, 0x6d, 0xe4, 0xdd, 0xa5, 0xc7, 0xa3, 0xef, 0x7e, 0xe7, 0x9b, 0x5d, 0x86, 0x60, 0x9c, 0x4a,
	0x41, 0x48, 0x1e, 0x0c, 0xad, 0x01, 0x39, 0xfd, 0xe1, 0x6f, 0x7b, 0x71, 0x99, 0x43, 0x2c, 0x0b,
	0x15, 0xbb, 0x6e, 0xfc, 0x79, 0x77, 0x00, 0xf7, 0x33, 0x20, 0x1b, 0xee, 0xff, 0xca, 0x53, 0x81,
	0x9a, 0x20, 0xfa, 0xd7, 0x61, 0xc7, 0x4c, 0x3c, 0x31, 0xde, 0xbc, 0x00, 0x9b, 0x4e, 0xfe, 0xc2,
	0xe2, 0xe8, 0xf1, 0xf5, 0x6d, 0xe9, 0x1d, 0x44, 0xd1, 0x97, 0xa6, 0x7a, 0xba, 0xb2, 0x1a, 0x55,
	0x69, 0x8e, 0x56, 0xd1, 0x50, 0x8f, 0x1d, 0xf6, 0xb7, 0x6e, 0xd6, 0x1d, 0xdf, 0xfa, 0x65, 0x2b,
	0x27, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0x1c, 0x18, 0xe9, 0x08, 0x02, 0x00, 0x00,
}
