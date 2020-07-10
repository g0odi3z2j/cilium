// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/event_service_config.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// [#not-implemented-hide:]
// Configuration of the event reporting service endpoint.
type EventServiceConfig struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*EventServiceConfig_GrpcService
	ConfigSourceSpecifier isEventServiceConfig_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                                   `json:"-"`
	XXX_unrecognized      []byte                                     `json:"-"`
	XXX_sizecache         int32                                      `json:"-"`
}

func (m *EventServiceConfig) Reset()         { *m = EventServiceConfig{} }
func (m *EventServiceConfig) String() string { return proto.CompactTextString(m) }
func (*EventServiceConfig) ProtoMessage()    {}
func (*EventServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac21a54a741b487, []int{0}
}

func (m *EventServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventServiceConfig.Unmarshal(m, b)
}
func (m *EventServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventServiceConfig.Marshal(b, m, deterministic)
}
func (m *EventServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventServiceConfig.Merge(m, src)
}
func (m *EventServiceConfig) XXX_Size() int {
	return xxx_messageInfo_EventServiceConfig.Size(m)
}
func (m *EventServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EventServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EventServiceConfig proto.InternalMessageInfo

type isEventServiceConfig_ConfigSourceSpecifier interface {
	isEventServiceConfig_ConfigSourceSpecifier()
}

type EventServiceConfig_GrpcService struct {
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*EventServiceConfig_GrpcService) isEventServiceConfig_ConfigSourceSpecifier() {}

func (m *EventServiceConfig) GetConfigSourceSpecifier() isEventServiceConfig_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *EventServiceConfig) GetGrpcService() *GrpcService {
	if x, ok := m.GetConfigSourceSpecifier().(*EventServiceConfig_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventServiceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventServiceConfig_GrpcService)(nil),
	}
}

func init() {
	proto.RegisterType((*EventServiceConfig)(nil), "envoy.config.core.v3.EventServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/event_service_config.proto", fileDescriptor_8ac21a54a741b487)
}

var fileDescriptor_8ac21a54a741b487 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xff, 0xcc, 0x0f, 0xb3, 0xc8, 0xb8, 0x90, 0x22, 0x54, 0x0a, 0x0e, 0xce, 0x20, 0x2a,
	0x08, 0x09, 0x4c, 0x77, 0x2e, 0x2b, 0x3a, 0x2e, 0x07, 0x7d, 0x80, 0x12, 0x33, 0x77, 0x4a, 0x40,
	0x72, 0x43, 0x92, 0x06, 0x67, 0xe7, 0xd2, 0x67, 0xf0, 0x0d, 0x7c, 0x05, 0xf7, 0x82, 0x5b, 0xdf,
	0xc6, 0xa5, 0xb4, 0xe9, 0xa0, 0xd0, 0xee, 0xda, 0x1c, 0xbe, 0x7b, 0xbe, 0x43, 0x39, 0xe8, 0x80,
	0x5b, 0x2e, 0x51, 0x6f, 0x54, 0xc5, 0x25, 0x5a, 0xe0, 0x21, 0xe7, 0x10, 0x40, 0xfb, 0xd2, 0x81,
	0x0d, 0x4a, 0x42, 0x19, 0x43, 0x66, 0x2c, 0x7a, 0x4c, 0x0e, 0x5a, 0x80, 0x75, 0x6f, 0x0d, 0xc0,
	0x42, 0x9e, 0x9d, 0x0d, 0x9e, 0xa9, 0xac, 0x91, 0xbb, 0x2b, 0x11, 0xcf, 0x8e, 0xea, 0xb5, 0x11,
	0x5c, 0x68, 0x8d, 0x5e, 0x78, 0x85, 0xda, 0x71, 0xe7, 0x85, 0xaf, 0x5d, 0x17, 0xcf, 0x7a, 0x71,
	0x00, 0xeb, 0x14, 0x6a, 0xa5, 0x3b, 0x81, 0x2c, 0x0d, 0xe2, 0x51, 0xad, 0x85, 0x07, 0xbe, 0xfb,
	0x88, 0xc1, 0xfc, 0x8d, 0xd0, 0xe4, 0xba, 0x11, 0xbf, 0x8f, 0x8d, 0x57, 0xad, 0x4c, 0x72, 0x43,
	0xf7, 0xfe, 0x7a, 0x1c, 0x92, 0x63, 0x72, 0x3e, 0x59, 0xcc, 0xd8, 0xd0, 0x0e, 0xb6, 0xb4, 0x46,
	0x76, 0xf8, 0xed, 0xbf, 0xbb, 0x49, 0xf5, 0xfb, 0x7b, 0x79, 0xf1, 0xfa, 0xf1, 0x32, 0x3d, 0xa5,
	0x27, 0x91, 0x13, 0x46, 0xb1, 0xb0, 0x88, 0x5c, 0xbf, 0xb4, 0x98, 0xd2, 0x34, 0x5e, 0x2e, 0x1d,
	0xd6, 0x56, 0x42, 0xe9, 0x0c, 0x48, 0xb5, 0x51, 0x60, 0x93, 0xff, 0xdf, 0x05, 0x29, 0x96, 0xef,
	0xcf, 0x9f, 0x5f, 0xe3, 0xd1, 0xfe, 0x88, 0xce, 0x15, 0x46, 0x15, 0x63, 0xf1, 0x69, 0x3b, 0x68,
	0x55, 0xa4, 0xfd, 0x86, 0x55, 0x33, 0x79, 0x45, 0x1e, 0xc6, 0xed, 0xf6, 0xfc, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x8b, 0xa4, 0x85, 0x92, 0xc8, 0x01, 0x00, 0x00,
}
