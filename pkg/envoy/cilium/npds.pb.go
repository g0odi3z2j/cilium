// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/npds.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import route "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/route"
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

// A network policy that is enforced by a filter on the network flows to/from
// associated hosts.
type NetworkPolicy struct {
	// The unique identifier of the network policy.
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The policy identifier associated with the network policy. Corresponds to
	// remote_policies entries in PortNetworkPolicyRule.
	// Required.
	Policy uint64 `protobuf:"varint,2,opt,name=policy,proto3" json:"policy,omitempty"`
	// The part of the policy to be enforced at ingress by the filter, as a set
	// of per-port network policies, one per destination L4 port.
	// Every PortNetworkPolicy element in this set has a unique port / protocol
	// combination.
	// Optional. If empty, all flows in this direction are denied.
	IngressPerPortPolicies []*PortNetworkPolicy `protobuf:"bytes,3,rep,name=ingress_per_port_policies,json=ingressPerPortPolicies,proto3" json:"ingress_per_port_policies,omitempty"`
	// The part of the policy to be enforced at egress by the filter, as a set
	// of per-port network policies, one per destination L4 port.
	// Every PortNetworkPolicy element in this set has a unique port / protocol
	// combination.
	// Optional. If empty, all flows in this direction are denied.
	EgressPerPortPolicies []*PortNetworkPolicy `protobuf:"bytes,4,rep,name=egress_per_port_policies,json=egressPerPortPolicies,proto3" json:"egress_per_port_policies,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *NetworkPolicy) Reset()         { *m = NetworkPolicy{} }
func (m *NetworkPolicy) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicy) ProtoMessage()    {}
func (*NetworkPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_npds_edb9353f6ee20ddf, []int{0}
}
func (m *NetworkPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPolicy.Unmarshal(m, b)
}
func (m *NetworkPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPolicy.Marshal(b, m, deterministic)
}
func (dst *NetworkPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicy.Merge(dst, src)
}
func (m *NetworkPolicy) XXX_Size() int {
	return xxx_messageInfo_NetworkPolicy.Size(m)
}
func (m *NetworkPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicy proto.InternalMessageInfo

func (m *NetworkPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkPolicy) GetPolicy() uint64 {
	if m != nil {
		return m.Policy
	}
	return 0
}

func (m *NetworkPolicy) GetIngressPerPortPolicies() []*PortNetworkPolicy {
	if m != nil {
		return m.IngressPerPortPolicies
	}
	return nil
}

func (m *NetworkPolicy) GetEgressPerPortPolicies() []*PortNetworkPolicy {
	if m != nil {
		return m.EgressPerPortPolicies
	}
	return nil
}

// A network policy to whitelist flows to a specific destination L4 port,
// as a conjunction of predicates on L3/L4/L7 flows.
// If all the predicates of a policy match a flow, the flow is whitelisted.
type PortNetworkPolicy struct {
	// The flows' destination L4 port number, as an unsigned 16-bit integer.
	// If 0, all destination L4 port numbers are matched by this predicate.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// The flows' L4 transport protocol.
	// Required.
	Protocol core.SocketAddress_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.api.v2.core.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The network policy rules to be enforced on the flows to the port.
	// Optional. A flow is matched by this predicate if either the set of
	// rules is empty or any of the rules matches it.
	Rules                []*PortNetworkPolicyRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PortNetworkPolicy) Reset()         { *m = PortNetworkPolicy{} }
func (m *PortNetworkPolicy) String() string { return proto.CompactTextString(m) }
func (*PortNetworkPolicy) ProtoMessage()    {}
func (*PortNetworkPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_npds_edb9353f6ee20ddf, []int{1}
}
func (m *PortNetworkPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortNetworkPolicy.Unmarshal(m, b)
}
func (m *PortNetworkPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortNetworkPolicy.Marshal(b, m, deterministic)
}
func (dst *PortNetworkPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortNetworkPolicy.Merge(dst, src)
}
func (m *PortNetworkPolicy) XXX_Size() int {
	return xxx_messageInfo_PortNetworkPolicy.Size(m)
}
func (m *PortNetworkPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_PortNetworkPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_PortNetworkPolicy proto.InternalMessageInfo

func (m *PortNetworkPolicy) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PortNetworkPolicy) GetProtocol() core.SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return core.SocketAddress_TCP
}

func (m *PortNetworkPolicy) GetRules() []*PortNetworkPolicyRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A network policy rule, as a conjunction of predicates on L3/L7 flows.
// If all the predicates of a rule match a flow, the flow is matched by the
// rule.
type PortNetworkPolicyRule struct {
	// The set of identifiers of policies of remote hosts.
	// A flow is matched by this predicate if the identifier of the policy
	// applied on the flow's remote host is contained in this set.
	// Optional. If not specified, any remote host is matched by this predicate.
	RemotePolicies []uint64 `protobuf:"varint,1,rep,packed,name=remote_policies,json=remotePolicies,proto3" json:"remote_policies,omitempty"`
	// Optional. If not specified, any L7 request is matched by this predicate.
	//
	// Types that are valid to be assigned to L7Rules:
	//	*PortNetworkPolicyRule_HttpRules
	L7Rules              isPortNetworkPolicyRule_L7Rules `protobuf_oneof:"l7_rules"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *PortNetworkPolicyRule) Reset()         { *m = PortNetworkPolicyRule{} }
func (m *PortNetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*PortNetworkPolicyRule) ProtoMessage()    {}
func (*PortNetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_npds_edb9353f6ee20ddf, []int{2}
}
func (m *PortNetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortNetworkPolicyRule.Unmarshal(m, b)
}
func (m *PortNetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortNetworkPolicyRule.Marshal(b, m, deterministic)
}
func (dst *PortNetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortNetworkPolicyRule.Merge(dst, src)
}
func (m *PortNetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_PortNetworkPolicyRule.Size(m)
}
func (m *PortNetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_PortNetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_PortNetworkPolicyRule proto.InternalMessageInfo

type isPortNetworkPolicyRule_L7Rules interface {
	isPortNetworkPolicyRule_L7Rules()
}

type PortNetworkPolicyRule_HttpRules struct {
	HttpRules *HttpNetworkPolicyRules `protobuf:"bytes,100,opt,name=http_rules,json=httpRules,proto3,oneof"`
}

func (*PortNetworkPolicyRule_HttpRules) isPortNetworkPolicyRule_L7Rules() {}

func (m *PortNetworkPolicyRule) GetL7Rules() isPortNetworkPolicyRule_L7Rules {
	if m != nil {
		return m.L7Rules
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetRemotePolicies() []uint64 {
	if m != nil {
		return m.RemotePolicies
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetHttpRules() *HttpNetworkPolicyRules {
	if x, ok := m.GetL7Rules().(*PortNetworkPolicyRule_HttpRules); ok {
		return x.HttpRules
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PortNetworkPolicyRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PortNetworkPolicyRule_OneofMarshaler, _PortNetworkPolicyRule_OneofUnmarshaler, _PortNetworkPolicyRule_OneofSizer, []interface{}{
		(*PortNetworkPolicyRule_HttpRules)(nil),
	}
}

func _PortNetworkPolicyRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PortNetworkPolicyRule)
	// l7_rules
	switch x := m.L7Rules.(type) {
	case *PortNetworkPolicyRule_HttpRules:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpRules); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PortNetworkPolicyRule.L7Rules has unexpected type %T", x)
	}
	return nil
}

func _PortNetworkPolicyRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PortNetworkPolicyRule)
	switch tag {
	case 100: // l7_rules.http_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpNetworkPolicyRules)
		err := b.DecodeMessage(msg)
		m.L7Rules = &PortNetworkPolicyRule_HttpRules{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PortNetworkPolicyRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PortNetworkPolicyRule)
	// l7_rules
	switch x := m.L7Rules.(type) {
	case *PortNetworkPolicyRule_HttpRules:
		s := proto.Size(x.HttpRules)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A set of network policy rules that match HTTP requests.
type HttpNetworkPolicyRules struct {
	// The set of HTTP network policy rules.
	// An HTTP request is matched if any of its rules matches the request.
	// Required and may not be empty.
	HttpRules            []*HttpNetworkPolicyRule `protobuf:"bytes,1,rep,name=http_rules,json=httpRules,proto3" json:"http_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HttpNetworkPolicyRules) Reset()         { *m = HttpNetworkPolicyRules{} }
func (m *HttpNetworkPolicyRules) String() string { return proto.CompactTextString(m) }
func (*HttpNetworkPolicyRules) ProtoMessage()    {}
func (*HttpNetworkPolicyRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_npds_edb9353f6ee20ddf, []int{3}
}
func (m *HttpNetworkPolicyRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNetworkPolicyRules.Unmarshal(m, b)
}
func (m *HttpNetworkPolicyRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNetworkPolicyRules.Marshal(b, m, deterministic)
}
func (dst *HttpNetworkPolicyRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNetworkPolicyRules.Merge(dst, src)
}
func (m *HttpNetworkPolicyRules) XXX_Size() int {
	return xxx_messageInfo_HttpNetworkPolicyRules.Size(m)
}
func (m *HttpNetworkPolicyRules) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNetworkPolicyRules.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNetworkPolicyRules proto.InternalMessageInfo

func (m *HttpNetworkPolicyRules) GetHttpRules() []*HttpNetworkPolicyRule {
	if m != nil {
		return m.HttpRules
	}
	return nil
}

// An HTTP network policy rule, as a conjunction of predicates on HTTP requests.
// If all the predicates of a rule match an HTTP request, the request is allowed. Otherwise, it is
// denied.
type HttpNetworkPolicyRule struct {
	// A set of matchers on the HTTP request's headers' names and values.
	// If all the matchers in this set match an HTTP request, the request is allowed by this rule.
	// Otherwise, it is denied.
	//
	// Some special header names are:
	//
	// * *:uri*: The HTTP request's URI.
	// * *:method*: The HTTP request's method.
	// * *:authority*: Also maps to the HTTP 1.1 *Host* header.
	//
	// Optional. If empty, matches any HTTP request.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HttpNetworkPolicyRule) Reset()         { *m = HttpNetworkPolicyRule{} }
func (m *HttpNetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*HttpNetworkPolicyRule) ProtoMessage()    {}
func (*HttpNetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_npds_edb9353f6ee20ddf, []int{4}
}
func (m *HttpNetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNetworkPolicyRule.Unmarshal(m, b)
}
func (m *HttpNetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNetworkPolicyRule.Marshal(b, m, deterministic)
}
func (dst *HttpNetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNetworkPolicyRule.Merge(dst, src)
}
func (m *HttpNetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_HttpNetworkPolicyRule.Size(m)
}
func (m *HttpNetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNetworkPolicyRule proto.InternalMessageInfo

func (m *HttpNetworkPolicyRule) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkPolicy)(nil), "cilium.NetworkPolicy")
	proto.RegisterType((*PortNetworkPolicy)(nil), "cilium.PortNetworkPolicy")
	proto.RegisterType((*PortNetworkPolicyRule)(nil), "cilium.PortNetworkPolicyRule")
	proto.RegisterType((*HttpNetworkPolicyRules)(nil), "cilium.HttpNetworkPolicyRules")
	proto.RegisterType((*HttpNetworkPolicyRule)(nil), "cilium.HttpNetworkPolicyRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkPolicyDiscoveryServiceClient is the client API for NetworkPolicyDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkPolicyDiscoveryServiceClient interface {
	StreamNetworkPolicies(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient, error)
	FetchNetworkPolicies(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type networkPolicyDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyDiscoveryServiceClient {
	return &networkPolicyDiscoveryServiceClient{cc}
}

func (c *networkPolicyDiscoveryServiceClient) StreamNetworkPolicies(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkPolicyDiscoveryService_serviceDesc.Streams[0], "/cilium.NetworkPolicyDiscoveryService/StreamNetworkPolicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyDiscoveryServiceStreamNetworkPoliciesClient{stream}
	return x, nil
}

type NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyDiscoveryServiceStreamNetworkPoliciesClient struct {
	grpc.ClientStream
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyDiscoveryServiceClient) FetchNetworkPolicies(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/cilium.NetworkPolicyDiscoveryService/FetchNetworkPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyDiscoveryServiceServer is the server API for NetworkPolicyDiscoveryService service.
type NetworkPolicyDiscoveryServiceServer interface {
	StreamNetworkPolicies(NetworkPolicyDiscoveryService_StreamNetworkPoliciesServer) error
	FetchNetworkPolicies(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterNetworkPolicyDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyDiscoveryService_StreamNetworkPolicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyDiscoveryServiceServer).StreamNetworkPolicies(&networkPolicyDiscoveryServiceStreamNetworkPoliciesServer{stream})
}

type NetworkPolicyDiscoveryService_StreamNetworkPoliciesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyDiscoveryServiceStreamNetworkPoliciesServer struct {
	grpc.ServerStream
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyDiscoveryService_FetchNetworkPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyDiscoveryServiceServer).FetchNetworkPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cilium.NetworkPolicyDiscoveryService/FetchNetworkPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyDiscoveryServiceServer).FetchNetworkPolicies(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cilium.NetworkPolicyDiscoveryService",
	HandlerType: (*NetworkPolicyDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchNetworkPolicies",
			Handler:    _NetworkPolicyDiscoveryService_FetchNetworkPolicies_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkPolicies",
			Handler:       _NetworkPolicyDiscoveryService_StreamNetworkPolicies_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cilium/npds.proto",
}

func init() { proto.RegisterFile("cilium/npds.proto", fileDescriptor_npds_edb9353f6ee20ddf) }

var fileDescriptor_npds_edb9353f6ee20ddf = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xbf, 0x6e, 0xd3, 0x40,
	0x18, 0xef, 0x25, 0x6e, 0x68, 0xbf, 0xaa, 0x45, 0x3d, 0x91, 0xe0, 0x46, 0x34, 0x09, 0x66, 0x71,
	0x3b, 0xd8, 0xc8, 0x19, 0x90, 0xca, 0x80, 0x88, 0x00, 0x65, 0x01, 0x45, 0x4e, 0x67, 0x22, 0xd7,
	0xfe, 0x94, 0x9c, 0xea, 0xf8, 0xcc, 0xf9, 0x62, 0x94, 0xb5, 0xe2, 0x09, 0x60, 0xe2, 0x2d, 0x98,
	0x99, 0x78, 0x07, 0x5e, 0x01, 0x06, 0x9e, 0xa2, 0xc8, 0x67, 0x3b, 0xad, 0xd5, 0x44, 0x2c, 0x2c,
	0xd6, 0x9d, 0x7f, 0x7f, 0xbe, 0x7f, 0x77, 0x07, 0x87, 0x3e, 0x0b, 0xd9, 0x62, 0x6e, 0x47, 0x71,
	0x90, 0x58, 0xb1, 0xe0, 0x92, 0xd3, 0x46, 0xfe, 0xab, 0xdd, 0xc5, 0x28, 0xe5, 0x4b, 0xdb, 0x8b,
	0x99, 0x9d, 0x3a, 0xb6, 0xcf, 0x05, 0xda, 0x5e, 0x10, 0x08, 0x4c, 0x0a, 0x62, 0xfb, 0x51, 0x85,
	0x10, 0xb0, 0xc4, 0xe7, 0x29, 0x8a, 0x65, 0x81, 0x76, 0x2a, 0xa8, 0xe0, 0x0b, 0x89, 0xf9, 0xb7,
	0x54, 0x4f, 0x39, 0x9f, 0x86, 0xa8, 0x08, 0x5e, 0x14, 0x71, 0xe9, 0x49, 0xc6, 0xa3, 0xd2, 0xfb,
	0x61, 0xea, 0x85, 0x2c, 0xf0, 0x24, 0xda, 0xe5, 0x22, 0x07, 0x8c, 0xdf, 0x04, 0xf6, 0xdf, 0xa1,
	0xfc, 0xc8, 0xc5, 0xe5, 0x88, 0x87, 0xcc, 0x5f, 0x52, 0x0a, 0x5a, 0xe4, 0xcd, 0x51, 0x27, 0x3d,
	0x62, 0xee, 0xba, 0x6a, 0x4d, 0x5b, 0xd0, 0x88, 0x15, 0xaa, 0xd7, 0x7a, 0xc4, 0xd4, 0xdc, 0x62,
	0x47, 0xcf, 0xe1, 0x88, 0x45, 0xd3, 0xac, 0x86, 0x49, 0x8c, 0x62, 0x12, 0x73, 0x21, 0x27, 0x0a,
	0x62, 0x98, 0xe8, 0xf5, 0x5e, 0xdd, 0xdc, 0x73, 0x8e, 0xac, 0xbc, 0x7e, 0x6b, 0xc4, 0x85, 0xac,
	0x44, 0x72, 0x5b, 0x85, 0x76, 0x84, 0x22, 0x03, 0x47, 0x85, 0x90, 0xba, 0xa0, 0xe3, 0x26, 0x53,
	0xed, 0x5f, 0xa6, 0x4d, 0x5c, 0xe7, 0x69, 0x7c, 0x23, 0x70, 0x78, 0x87, 0x4c, 0xbb, 0xa0, 0x65,
	0xf6, 0xaa, 0xd6, 0xfd, 0xc1, 0xde, 0xf7, 0x3f, 0x3f, 0xea, 0x8d, 0x53, 0x4d, 0xbf, 0xbe, 0xae,
	0xbb, 0x0a, 0xa0, 0xaf, 0x61, 0x47, 0xf5, 0xc9, 0xe7, 0xa1, 0x2a, 0xfd, 0xc0, 0x39, 0xb1, 0xd4,
	0x20, 0x2c, 0x2f, 0x66, 0x56, 0xea, 0x58, 0xd9, 0x1c, 0xad, 0x31, 0xf7, 0x2f, 0x51, 0xbe, 0x2c,
	0xa6, 0x39, 0x2a, 0x04, 0xee, 0x4a, 0x4a, 0xfb, 0xb0, 0x2d, 0x16, 0xe1, 0xaa, 0x27, 0xc7, 0x9b,
	0xd3, 0x5f, 0x84, 0xe8, 0xe6, 0x5c, 0xe3, 0x2b, 0x81, 0xe6, 0x5a, 0x02, 0xed, 0xc3, 0x7d, 0x81,
	0x73, 0x2e, 0xf1, 0xa6, 0x2f, 0xa4, 0x57, 0x37, 0xb5, 0x01, 0x64, 0x15, 0x6c, 0x7f, 0x26, 0x35,
	0x9d, 0xb8, 0x07, 0x39, 0x65, 0xd5, 0xd5, 0x17, 0x00, 0x33, 0x29, 0xe3, 0x49, 0x9e, 0x48, 0xd0,
	0x23, 0xe6, 0x9e, 0xd3, 0x29, 0x13, 0x19, 0x4a, 0x19, 0xdf, 0x89, 0x93, 0x0c, 0xb7, 0xdc, 0xdd,
	0x4c, 0xa3, 0x36, 0x03, 0x80, 0x9d, 0xf0, 0x59, 0x2e, 0x37, 0x2e, 0xa0, 0xb5, 0x5e, 0x42, 0x87,
	0x95, 0x30, 0xa4, 0x5a, 0xef, 0x5a, 0xcd, 0x4d, 0xd6, 0x3b, 0xe4, 0x56, 0x3c, 0xe3, 0x1c, 0x9a,
	0x6b, 0xf9, 0xf4, 0x39, 0xdc, 0x9b, 0xa1, 0x17, 0xa0, 0x28, 0xfd, 0x1f, 0x57, 0x67, 0x92, 0x5f,
	0x8b, 0xa1, 0xa2, 0xbc, 0xf5, 0xa4, 0x3f, 0x43, 0xe1, 0x96, 0x0a, 0xe7, 0x53, 0x0d, 0x8e, 0x2b,
	0x96, 0xaf, 0xca, 0x8b, 0x36, 0x46, 0x91, 0x32, 0x1f, 0xe9, 0x7b, 0x68, 0x8e, 0xa5, 0x40, 0x6f,
	0x7e, 0x9b, 0x96, 0x75, 0xb0, 0x53, 0x0d, 0xb3, 0x12, 0xba, 0xf8, 0x61, 0x81, 0x89, 0x6c, 0x77,
	0x37, 0xe2, 0x49, 0xcc, 0xa3, 0x04, 0x8d, 0x2d, 0x93, 0x3c, 0x25, 0xf4, 0x8a, 0xc0, 0x83, 0x37,
	0x28, 0xfd, 0xd9, 0x7f, 0xf7, 0x3f, 0xb9, 0xfa, 0xf9, 0xeb, 0x4b, 0xed, 0x89, 0xd1, 0xa9, 0x3c,
	0x20, 0x67, 0x51, 0x1e, 0x67, 0x75, 0x58, 0xce, 0xc8, 0xe9, 0x45, 0x43, 0x9d, 0xcd, 0xfe, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xa9, 0xbf, 0xe5, 0xb1, 0x04, 0x00, 0x00,
}
