// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/socket_option.proto

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

type SocketOption_SocketState int32

const (
	// Socket options are applied after socket creation but before binding the socket to a port
	SocketOption_STATE_PREBIND SocketOption_SocketState = 0
	// Socket options are applied after binding the socket to a port but before calling listen()
	SocketOption_STATE_BOUND SocketOption_SocketState = 1
	// Socket options are applied after calling listen()
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

var SocketOption_SocketState_name = map[int32]string{
	0: "STATE_PREBIND",
	1: "STATE_BOUND",
	2: "STATE_LISTENING",
}

var SocketOption_SocketState_value = map[string]int32{
	"STATE_PREBIND":   0,
	"STATE_BOUND":     1,
	"STATE_LISTENING": 2,
}

func (x SocketOption_SocketState) String() string {
	return proto.EnumName(SocketOption_SocketState_name, int32(x))
}

func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e4a49540b4573ea, []int{0, 0}
}

// Generic socket option message. This would be used to set socket options that
// might not exist in upstream kernels or precompiled Envoy binaries.
// [#next-free-field: 7]
type SocketOption struct {
	// An optional name to give this socket option for debugging, etc.
	// Uniqueness is not required and no special meaning is assumed.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Corresponding to the level value passed to setsockopt, such as IPPROTO_TCP
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// The numeric name as passed to setsockopt
	Name int64 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value isSocketOption_Value `protobuf_oneof:"value"`
	// The state in which the option will be applied. When used in BindConfig
	// STATE_PREBIND is currently the only valid value.
	State                SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=envoy.config.core.v3.SocketOption_SocketState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SocketOption) Reset()         { *m = SocketOption{} }
func (m *SocketOption) String() string { return proto.CompactTextString(m) }
func (*SocketOption) ProtoMessage()    {}
func (*SocketOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4a49540b4573ea, []int{0}
}

func (m *SocketOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketOption.Unmarshal(m, b)
}
func (m *SocketOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketOption.Marshal(b, m, deterministic)
}
func (m *SocketOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketOption.Merge(m, src)
}
func (m *SocketOption) XXX_Size() int {
	return xxx_messageInfo_SocketOption.Size(m)
}
func (m *SocketOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketOption.DiscardUnknown(m)
}

var xxx_messageInfo_SocketOption proto.InternalMessageInfo

func (m *SocketOption) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SocketOption) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SocketOption) GetName() int64 {
	if m != nil {
		return m.Name
	}
	return 0
}

type isSocketOption_Value interface {
	isSocketOption_Value()
}

type SocketOption_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}

type SocketOption_BufValue struct {
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}

func (*SocketOption_BufValue) isSocketOption_Value() {}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SocketOption) GetIntValue() int64 {
	if x, ok := m.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *SocketOption) GetBufValue() []byte {
	if x, ok := m.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (m *SocketOption) GetState() SocketOption_SocketState {
	if m != nil {
		return m.State
	}
	return SocketOption_STATE_PREBIND
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketOption) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.SocketOption_SocketState", SocketOption_SocketState_name, SocketOption_SocketState_value)
	proto.RegisterType((*SocketOption)(nil), "envoy.config.core.v3.SocketOption")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/socket_option.proto", fileDescriptor_8e4a49540b4573ea)
}

var fileDescriptor_8e4a49540b4573ea = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x8a, 0x13, 0x31,
	0x1c, 0xc6, 0x37, 0xd3, 0x9d, 0xb2, 0x9b, 0x56, 0x77, 0x36, 0x2e, 0x38, 0x2c, 0xec, 0x32, 0x16,
	0x84, 0x39, 0x65, 0x60, 0x7b, 0xf3, 0xd6, 0xd0, 0xaa, 0x05, 0x99, 0x96, 0x99, 0xea, 0xb5, 0xa4,
	0xd3, 0xb4, 0x04, 0xc7, 0x64, 0x98, 0xc9, 0x04, 0x7b, 0x13, 0x4f, 0x3e, 0x83, 0x8f, 0xe2, 0x5d,
	0xf0, 0xea, 0x43, 0xf8, 0x0e, 0xe2, 0x49, 0x92, 0x54, 0x18, 0xb0, 0xb7, 0xfc, 0xbf, 0xdf, 0x97,
	0xf0, 0x7d, 0xf9, 0xc3, 0x98, 0x09, 0x2d, 0x0f, 0x49, 0x21, 0xc5, 0x8e, 0xef, 0x93, 0x42, 0xd6,
	0x2c, 0xd1, 0xe3, 0xa4, 0x91, 0xc5, 0x7b, 0xa6, 0xd6, 0xb2, 0x52, 0x5c, 0x0a, 0x5c, 0xd5, 0x52,
	0x49, 0x74, 0x63, 0x9d, 0xd8, 0x39, 0xb1, 0x71, 0x62, 0x3d, 0xbe, 0xbd, 0x6b, 0xb7, 0x15, 0x4d,
	0xa8, 0x10, 0x52, 0x51, 0x63, 0x6e, 0x92, 0x46, 0x51, 0xd5, 0x36, 0xee, 0xd2, 0xed, 0xb3, 0xff,
	0xb0, 0x66, 0x75, 0xc3, 0xa5, 0xe0, 0x62, 0x7f, 0xb4, 0x3c, 0xd5, 0xb4, 0xe4, 0x5b, 0xaa, 0x58,
	0xf2, 0xef, 0xe0, 0xc0, 0xe8, 0x97, 0x07, 0x87, 0xb9, 0x0d, 0xb2, 0xb0, 0x39, 0x50, 0x04, 0x07,
	0x5b, 0xd6, 0x14, 0x35, 0xb7, 0x63, 0x08, 0x22, 0x10, 0x5f, 0x66, 0x5d, 0x09, 0xdd, 0x40, 0xbf,
	0x64, 0x9a, 0x95, 0xa1, 0x17, 0x81, 0xb8, 0x97, 0xb9, 0x01, 0x21, 0x78, 0x2e, 0xe8, 0x07, 0x16,
	0xf6, 0xac, 0x68, 0xcf, 0xe8, 0x0e, 0x5e, 0x72, 0xa1, 0xd6, 0x9a, 0x96, 0x2d, 0x0b, 0xcf, 0x0d,
	0x78, 0x7d, 0x96, 0x5d, 0x70, 0xa1, 0xde, 0x19, 0xc5, 0xe0, 0x4d, 0xbb, 0x3b, 0x62, 0x3f, 0x02,
	0xf1, 0xd0, 0xe0, 0x4d, 0xbb, 0x73, 0x38, 0x85, 0xbe, 0xa9, 0xc9, 0xc2, 0x7e, 0x04, 0xe2, 0xc7,
	0x0f, 0x18, 0x9f, 0xfa, 0x1b, 0xdc, 0x0d, 0x7f, 0x1c, 0x72, 0x73, 0x8b, 0x5c, 0xfc, 0x21, 0xfe,
	0x67, 0xe0, 0x05, 0x20, 0x73, 0xcf, 0x8c, 0x5e, 0xc2, 0x41, 0x87, 0xa3, 0x6b, 0xf8, 0x28, 0x5f,
	0x4d, 0x56, 0xb3, 0xf5, 0x32, 0x9b, 0x91, 0x79, 0x3a, 0x0d, 0xce, 0xd0, 0x15, 0x1c, 0x38, 0x89,
	0x2c, 0xde, 0xa6, 0xd3, 0x00, 0xa0, 0x27, 0xf0, 0xca, 0x09, 0x6f, 0xe6, 0xf9, 0x6a, 0x96, 0xce,
	0xd3, 0x57, 0x81, 0xf7, 0xe2, 0xf9, 0xd7, 0xef, 0x5f, 0xee, 0x23, 0x78, 0xef, 0xe2, 0xd0, 0x8a,
	0x63, 0xfd, 0xe0, 0xe2, 0x74, 0xb3, 0x90, 0x21, 0xf4, 0x6d, 0x33, 0xd4, 0xfb, 0x4d, 0x00, 0x99,
	0x7c, 0xfb, 0xf4, 0xe3, 0x67, 0xdf, 0x0b, 0x3c, 0x38, 0xe2, 0xd2, 0x35, 0xa9, 0x6a, 0xf9, 0xf1,
	0x70, 0xb2, 0x14, 0xb9, 0xee, 0xbe, 0xb4, 0x34, 0x8b, 0x5a, 0x82, 0x4d, 0xdf, 0x6e, 0x6c, 0xfc,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x42, 0x4e, 0x17, 0x40, 0x4e, 0x02, 0x00, 0x00,
}
