// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The external authorization HTTP service configuration.
type HttpService struct {
	// Sets the HTTP server URI which the authorization requests must be sent to.
	ServerUri *core.HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri" json:"server_uri,omitempty"`
	// Sets an optional prefix to the value of authorization request header `path`.
	PathPrefix           string   `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix" json:"path_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_ext_authz_8cebb836c8eae313, []int{0}
}
func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (dst *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(dst, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *core.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

type ExtAuthz struct {
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services isExtAuthz_Services `protobuf_oneof:"services"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. If set to true then in case of failure to get a
	// response back from the authorization service or getting a response that
	// is NOT denied then traffic will be permitted.
	// Defaults to false.
	FailureModeAllow     bool     `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow" json:"failure_mode_allow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_ext_authz_8cebb836c8eae313, []int{1}
}
func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (dst *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(dst, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,oneof"`
}
type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}
func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtAuthz_OneofMarshaler, _ExtAuthz_OneofUnmarshaler, _ExtAuthz_OneofSizer, []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

func _ExtAuthz_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtAuthz)
	// services
	switch x := m.Services.(type) {
	case *ExtAuthz_GrpcService:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GrpcService); err != nil {
			return err
		}
	case *ExtAuthz_HttpService:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpService); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExtAuthz.Services has unexpected type %T", x)
	}
	return nil
}

func _ExtAuthz_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtAuthz)
	switch tag {
	case 1: // services.grpc_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.GrpcService)
		err := b.DecodeMessage(msg)
		m.Services = &ExtAuthz_GrpcService{msg}
		return true, err
	case 3: // services.http_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpService)
		err := b.DecodeMessage(msg)
		m.Services = &ExtAuthz_HttpService{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExtAuthz_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtAuthz)
	// services
	switch x := m.Services.(type) {
	case *ExtAuthz_GrpcService:
		s := proto.Size(x.GrpcService)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtAuthz_HttpService:
		s := proto.Size(x.HttpService)
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
	proto.RegisterType((*HttpService)(nil), "envoy.config.filter.http.ext_authz.v2alpha.HttpService")
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v2alpha.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto", fileDescriptor_ext_authz_8cebb836c8eae313)
}

var fileDescriptor_ext_authz_8cebb836c8eae313 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x5d, 0x05, 0x6d, 0x67, 0x7b, 0x90, 0x9c, 0x4a, 0x0f, 0x5a, 0x8a, 0x87, 0x22, 0x92,
	0xc0, 0x7a, 0x10, 0xbd, 0xb5, 0x22, 0xf6, 0x22, 0xc8, 0x4a, 0x2f, 0x22, 0x84, 0xb8, 0xcd, 0x76,
	0x03, 0x6b, 0x13, 0xa6, 0xe9, 0x5a, 0xfd, 0xc2, 0x7e, 0x0d, 0xc9, 0x1f, 0xe9, 0x82, 0x1e, 0x3c,
	0x66, 0xf2, 0xde, 0x6f, 0xde, 0x1b, 0xb8, 0x91, 0xab, 0x46, 0x7f, 0xb0, 0x42, 0xaf, 0x4a, 0xb5,
	0x64, 0xa5, 0xaa, 0xad, 0x44, 0x56, 0x59, 0x6b, 0x98, 0xdc, 0x5a, 0x2e, 0x36, 0xb6, 0xfa, 0x64,
	0x4d, 0x26, 0x6a, 0x53, 0x89, 0xdd, 0x84, 0x1a, 0xd4, 0x56, 0x93, 0x73, 0xef, 0xa5, 0xc1, 0x4b,
	0x83, 0x97, 0x3a, 0x2f, 0xdd, 0x29, 0xa3, 0x77, 0x70, 0x16, 0xf6, 0x08, 0xa3, 0x58, 0x93, 0xb1,
	0x42, 0xa3, 0x64, 0x4b, 0x34, 0x05, 0x5f, 0x4b, 0x6c, 0x54, 0x21, 0x03, 0x71, 0x30, 0xfc, 0xad,
	0x72, 0x3c, 0xbe, 0x41, 0x15, 0x14, 0x23, 0x05, 0xe9, 0xcc, 0x5a, 0xf3, 0x14, 0x6c, 0xe4, 0x1a,
	0xc0, 0x11, 0x24, 0x3a, 0x49, 0x3f, 0x19, 0x26, 0xe3, 0x34, 0x1b, 0xd0, 0x90, 0x4b, 0x18, 0x45,
	0x9b, 0x8c, 0x3a, 0x0a, 0x75, 0x9e, 0x39, 0xaa, 0xbc, 0x1b, 0xd4, 0x73, 0x54, 0xe4, 0x14, 0x52,
	0x23, 0x6c, 0xc5, 0x0d, 0xca, 0x52, 0x6d, 0xfb, 0xfb, 0xc3, 0x64, 0xdc, 0xcd, 0xc1, 0x8d, 0x1e,
	0xfd, 0x64, 0xf4, 0x95, 0x40, 0xe7, 0x6e, 0x6b, 0x27, 0xae, 0x07, 0xb9, 0x85, 0x5e, 0x3b, 0x6f,
	0x5c, 0x75, 0xf2, 0xc7, 0xaa, 0x7b, 0x34, 0x45, 0x8c, 0x37, 0xdb, 0xcb, 0xd3, 0xe5, 0xee, 0x49,
	0x5e, 0xa0, 0xe7, 0xeb, 0xfc, 0x40, 0x0e, 0x3c, 0xe4, 0x8a, 0xfe, 0xff, 0x8e, 0xb4, 0x55, 0xde,
	0xd1, 0xab, 0xd6, 0x2d, 0x2e, 0x80, 0x94, 0x42, 0xd5, 0x1b, 0x94, 0xfc, 0x4d, 0x2f, 0x24, 0x17,
	0x75, 0xad, 0xdf, 0x7d, 0xaf, 0x4e, 0x7e, 0x1c, 0x7f, 0x1e, 0xf4, 0x42, 0x4e, 0xdc, 0x7c, 0x0a,
	0xd0, 0x89, 0x31, 0xd6, 0xd3, 0xee, 0xf3, 0x51, 0xe4, 0xbf, 0x1e, 0xfa, 0x33, 0x5f, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x5f, 0x58, 0x08, 0x18, 0x02, 0x00, 0x00,
}
