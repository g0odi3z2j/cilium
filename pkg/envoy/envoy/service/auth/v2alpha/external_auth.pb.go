// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _type "github.com/cilium/cilium/pkg/envoy/envoy/type"
import _ "github.com/lyft/protoc-gen-validate/validate"
import status "google.golang.org/genproto/googleapis/rpc/status"

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

type CheckRequest struct {
	// The request attributes.
	Attributes           *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_external_auth_bdb349b8028b1219, []int{0}
}
func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (dst *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(dst, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// HTTP attributes for a denied response.
type DeniedHttpResponse struct {
	// This field allows the authorization service to send a HTTP response status
	// code to the downstream client other than 403 (Forbidden).
	Status *_type.HttpStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the the downstream client.
	Headers []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// This field allows the authorization service to send a response body data
	// to the the downstream client.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeniedHttpResponse) Reset()         { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()    {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_external_auth_bdb349b8028b1219, []int{1}
}
func (m *DeniedHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeniedHttpResponse.Unmarshal(m, b)
}
func (m *DeniedHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeniedHttpResponse.Marshal(b, m, deterministic)
}
func (dst *DeniedHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeniedHttpResponse.Merge(dst, src)
}
func (m *DeniedHttpResponse) XXX_Size() int {
	return xxx_messageInfo_DeniedHttpResponse.Size(m)
}
func (m *DeniedHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeniedHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeniedHttpResponse proto.InternalMessageInfo

func (m *DeniedHttpResponse) GetStatus() *_type.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// HTTP attributes for an ok response.
type OkHttpResponse struct {
	// HTTP entity headers in addition to the original request headers. This allows the authorization
	// service to append, to add or to override headers from the original request before
	// dispatching it to the upstream. By setting `append` field to `true` in the `HeaderValueOption`,
	// the filter will append the correspondent header value to the matched request header. Note that
	// by Leaving `append` as false, the filter will either add a new header, or override an existing
	// one if there is a match.
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OkHttpResponse) Reset()         { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()    {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_external_auth_bdb349b8028b1219, []int{2}
}
func (m *OkHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkHttpResponse.Unmarshal(m, b)
}
func (m *OkHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkHttpResponse.Marshal(b, m, deterministic)
}
func (dst *OkHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkHttpResponse.Merge(dst, src)
}
func (m *OkHttpResponse) XXX_Size() int {
	return xxx_messageInfo_OkHttpResponse.Size(m)
}
func (m *OkHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkHttpResponse proto.InternalMessageInfo

func (m *OkHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Intended for gRPC and Network Authorization servers `only`.
type CheckResponse struct {
	// Status `OK` allows the request. Any other status indicates the request should be denied.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// An message that contains HTTP response attributes. This message is
	// used when the authorization service needs to send custom responses to the
	// downstream client or, to modify/add request headers being dispatched to the upstream.
	//
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse         isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_external_auth_bdb349b8028b1219, []int{3}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (dst *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(dst, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof"`
}

type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}

func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse() {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CheckResponse_OneofMarshaler, _CheckResponse_OneofUnmarshaler, _CheckResponse_OneofSizer, []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func _CheckResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CheckResponse)
	// http_response
	switch x := m.HttpResponse.(type) {
	case *CheckResponse_DeniedResponse:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeniedResponse); err != nil {
			return err
		}
	case *CheckResponse_OkResponse:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OkResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CheckResponse.HttpResponse has unexpected type %T", x)
	}
	return nil
}

func _CheckResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CheckResponse)
	switch tag {
	case 2: // http_response.denied_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeniedHttpResponse)
		err := b.DecodeMessage(msg)
		m.HttpResponse = &CheckResponse_DeniedResponse{msg}
		return true, err
	case 3: // http_response.ok_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OkHttpResponse)
		err := b.DecodeMessage(msg)
		m.HttpResponse = &CheckResponse_OkResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CheckResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CheckResponse)
	// http_response
	switch x := m.HttpResponse.(type) {
	case *CheckResponse_DeniedResponse:
		s := proto.Size(x.DeniedResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CheckResponse_OkResponse:
		s := proto.Size(x.OkResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v2alpha.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v2alpha.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v2alpha.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v2alpha.CheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptor_external_auth_bdb349b8028b1219)
}

var fileDescriptor_external_auth_bdb349b8028b1219 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xef, 0x34, 0xb6, 0xc5, 0x97, 0xa6, 0x95, 0x39, 0xd8, 0x10, 0x3c, 0x84, 0xe0, 0x21, 0x16,
	0x99, 0x81, 0xf5, 0xe6, 0x41, 0x68, 0xeb, 0x21, 0x07, 0xa5, 0xb2, 0x82, 0xa0, 0x14, 0xc2, 0x64,
	0xf7, 0xd1, 0x5d, 0xb2, 0xee, 0x8c, 0x33, 0x6f, 0x97, 0xc6, 0x4f, 0x20, 0x7e, 0x0e, 0x3f, 0x85,
	0x27, 0xbf, 0x8e, 0x5f, 0xc0, 0xb3, 0xec, 0xec, 0x6c, 0x6c, 0x14, 0x83, 0xe0, 0x6d, 0xd9, 0xdf,
	0x9f, 0xf7, 0x7e, 0xbf, 0x79, 0x20, 0xb0, 0xac, 0xf5, 0x4a, 0x3a, 0xb4, 0x75, 0x9e, 0xa0, 0x54,
	0x15, 0x65, 0xb2, 0x8e, 0x54, 0x61, 0x32, 0x25, 0xf1, 0x86, 0xd0, 0x96, 0xaa, 0x98, 0x37, 0x7f,
	0x85, 0xb1, 0x9a, 0x34, 0x1f, 0x79, 0xbe, 0x08, 0x7c, 0xe1, 0x91, 0xc0, 0x1f, 0x3d, 0x68, 0xbd,
	0x94, 0xc9, 0x65, 0x1d, 0xc9, 0x44, 0x5b, 0x94, 0x0b, 0xe5, 0xb0, 0x55, 0x76, 0x28, 0xad, 0x0c,
	0xca, 0x8c, 0xc8, 0xcc, 0x1d, 0x29, 0xaa, 0x5c, 0x40, 0xa3, 0x2d, 0x7b, 0x28, 0x22, 0x9b, 0x2f,
	0x2a, 0xc2, 0x79, 0xa2, 0x4b, 0xc2, 0x1b, 0x0a, 0x9a, 0x93, 0x6b, 0xad, 0xaf, 0x0b, 0x94, 0xd6,
	0x24, 0x72, 0xc3, 0xec, 0xa4, 0x56, 0x45, 0x9e, 0x2a, 0x42, 0xd9, 0x7d, 0xb4, 0xc0, 0xe4, 0x0a,
	0x0e, 0x2f, 0x32, 0x4c, 0x96, 0x31, 0x7e, 0xa8, 0xd0, 0x11, 0x7f, 0x01, 0xb0, 0x36, 0x77, 0x43,
	0x36, 0x66, 0xd3, 0x7e, 0xf4, 0x58, 0xfc, 0x3d, 0xa2, 0x38, 0xeb, 0xd8, 0x17, 0xed, 0x26, 0xf1,
	0x2d, 0xfd, 0xe4, 0x0b, 0x03, 0xfe, 0x1c, 0xcb, 0x1c, 0xd3, 0x19, 0x91, 0x89, 0xd1, 0x19, 0x5d,
	0x3a, 0xe4, 0x4f, 0x61, 0xbf, 0xdd, 0x2e, 0x0c, 0xb8, 0x1f, 0x06, 0x34, 0x4d, 0x88, 0x86, 0xf9,
	0xda, 0xa3, 0xe7, 0xf0, 0xf5, 0xfb, 0xb7, 0xde, 0xde, 0x67, 0xb6, 0x7b, 0x8f, 0xc5, 0x41, 0xc1,
	0x9f, 0xc1, 0x41, 0x86, 0x2a, 0x45, 0xeb, 0x86, 0xbb, 0xe3, 0xde, 0xb4, 0x1f, 0x3d, 0x0c, 0x62,
	0x65, 0x72, 0x51, 0x47, 0xa2, 0x29, 0x59, 0xcc, 0x3c, 0xe3, 0x8d, 0x2a, 0x2a, 0xbc, 0x34, 0x94,
	0xeb, 0x32, 0xee, 0x44, 0x9c, 0xc3, 0x9d, 0x85, 0x4e, 0x57, 0xc3, 0xde, 0x98, 0x4d, 0xef, 0xc6,
	0xfe, 0x7b, 0xf2, 0x0a, 0x8e, 0x2e, 0x97, 0x1b, 0x1b, 0xfe, 0xe7, 0x94, 0xc9, 0x0f, 0x06, 0x83,
	0xd0, 0x6b, 0x70, 0x3c, 0xfd, 0x2d, 0x33, 0x17, 0xed, 0x5b, 0x09, 0x6b, 0x12, 0xd1, 0xe6, 0x5d,
	0x67, 0x7c, 0x0b, 0xc7, 0xa9, 0x6f, 0x6d, 0x6e, 0x83, 0x7c, 0xb8, 0xeb, 0x45, 0x62, 0xdb, 0x4b,
	0xfc, 0x59, 0xf4, 0x6c, 0x27, 0x3e, 0x6a, 0x8d, 0xd6, 0x6b, 0xbc, 0x84, 0xbe, 0x5e, 0xfe, 0xb2,
	0xed, 0x79, 0xdb, 0xd3, 0x6d, 0xb6, 0x9b, 0xcd, 0xcc, 0x76, 0x62, 0xd0, 0xeb, 0x54, 0xe7, 0xc7,
	0x30, 0xf0, 0x97, 0xdb, 0x19, 0x46, 0xef, 0x61, 0x70, 0x56, 0x51, 0xa6, 0x6d, 0xfe, 0x51, 0x35,
	0x95, 0xf0, 0x2b, 0xd8, 0xf3, 0x45, 0xf0, 0xe9, 0xb6, 0x21, 0xb7, 0x6f, 0x70, 0xf4, 0xe8, 0x1f,
	0x98, 0x61, 0xfe, 0xe1, 0xbb, 0x83, 0x00, 0x7c, 0x62, 0x6c, 0xb1, 0xef, 0x6f, 0xfa, 0xc9, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xde, 0x56, 0xee, 0xc3, 0x03, 0x00, 0x00,
}
