// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/percent.proto

package envoy_type

import (
	fmt "fmt"
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

// Fraction percentages support several fixed denominator values.
type FractionalPercent_DenominatorType int32

const (
	// 100.
	//
	// **Example**: 1/100 = 1%.
	FractionalPercent_HUNDRED FractionalPercent_DenominatorType = 0
	// 10,000.
	//
	// **Example**: 1/10000 = 0.01%.
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	// 1,000,000.
	//
	// **Example**: 1/1000000 = 0.0001%.
	FractionalPercent_MILLION FractionalPercent_DenominatorType = 2
)

var FractionalPercent_DenominatorType_name = map[int32]string{
	0: "HUNDRED",
	1: "TEN_THOUSAND",
	2: "MILLION",
}

var FractionalPercent_DenominatorType_value = map[string]int32{
	"HUNDRED":      0,
	"TEN_THOUSAND": 1,
	"MILLION":      2,
}

func (x FractionalPercent_DenominatorType) String() string {
	return proto.EnumName(FractionalPercent_DenominatorType_name, int32(x))
}

func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89401f90eb07307e, []int{1, 0}
}

// Identifies a percentage, in the range [0.0, 100.0].
type Percent struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Percent) Reset()         { *m = Percent{} }
func (m *Percent) String() string { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()    {}
func (*Percent) Descriptor() ([]byte, []int) {
	return fileDescriptor_89401f90eb07307e, []int{0}
}

func (m *Percent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Percent.Unmarshal(m, b)
}
func (m *Percent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Percent.Marshal(b, m, deterministic)
}
func (m *Percent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Percent.Merge(m, src)
}
func (m *Percent) XXX_Size() int {
	return xxx_messageInfo_Percent.Size(m)
}
func (m *Percent) XXX_DiscardUnknown() {
	xxx_messageInfo_Percent.DiscardUnknown(m)
}

var xxx_messageInfo_Percent proto.InternalMessageInfo

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// A fractional percentage is used in cases in which for performance reasons performing floating
// point to integer conversions during randomness calculations is undesirable. The message includes
// both a numerator and denominator that together determine the final fractional value.
//
// * **Example**: 1/100 = 1%.
// * **Example**: 3/10000 = 0.03%.
type FractionalPercent struct {
	// Specifies the numerator. Defaults to 0.
	Numerator uint32 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// Specifies the denominator. If the denominator specified is less than the numerator, the final
	// fractional percentage is capped at 1 (100%).
	Denominator          FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FractionalPercent) Reset()         { *m = FractionalPercent{} }
func (m *FractionalPercent) String() string { return proto.CompactTextString(m) }
func (*FractionalPercent) ProtoMessage()    {}
func (*FractionalPercent) Descriptor() ([]byte, []int) {
	return fileDescriptor_89401f90eb07307e, []int{1}
}

func (m *FractionalPercent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FractionalPercent.Unmarshal(m, b)
}
func (m *FractionalPercent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FractionalPercent.Marshal(b, m, deterministic)
}
func (m *FractionalPercent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPercent.Merge(m, src)
}
func (m *FractionalPercent) XXX_Size() int {
	return xxx_messageInfo_FractionalPercent.Size(m)
}
func (m *FractionalPercent) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPercent.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPercent proto.InternalMessageInfo

func (m *FractionalPercent) GetNumerator() uint32 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if m != nil {
		return m.Denominator
	}
	return FractionalPercent_HUNDRED
}

func init() {
	proto.RegisterEnum("envoy.type.FractionalPercent_DenominatorType", FractionalPercent_DenominatorType_name, FractionalPercent_DenominatorType_value)
	proto.RegisterType((*Percent)(nil), "envoy.type.Percent")
	proto.RegisterType((*FractionalPercent)(nil), "envoy.type.FractionalPercent")
}

func init() { proto.RegisterFile("envoy/type/percent.proto", fileDescriptor_89401f90eb07307e) }

var fileDescriptor_89401f90eb07307e = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0xcb, 0xe8, 0x81, 0x64, 0xa4, 0xc4, 0xcb, 0x12, 0x73,
	0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x22, 0x25, 0x0b, 0x2e, 0xf6, 0x00, 0x88,
	0x2e, 0x21, 0x5d, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x46,
	0x27, 0xf1, 0x5f, 0x4e, 0x22, 0x42, 0x42, 0x92, 0x0c, 0x60, 0x10, 0xe9, 0xa0, 0xc9, 0x00, 0x05,
	0x41, 0x10, 0x55, 0x4a, 0xa7, 0x19, 0xb9, 0x04, 0xdd, 0x8a, 0x12, 0x93, 0x4b, 0x32, 0xf3, 0xf3,
	0x12, 0x73, 0x60, 0x86, 0xc8, 0x70, 0x71, 0xe6, 0x95, 0xe6, 0xa6, 0x16, 0x25, 0x96, 0xe4, 0x17,
	0x81, 0x0d, 0xe2, 0x0d, 0x42, 0x08, 0x08, 0x45, 0x72, 0x71, 0xa7, 0xa4, 0xe6, 0xe5, 0xe7, 0x66,
	0xe6, 0x81, 0xe5, 0x99, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x74, 0xf5, 0x10, 0x0e, 0xd5, 0xc3, 0x30,
	0x51, 0xcf, 0x05, 0xa1, 0x21, 0xa4, 0xb2, 0x20, 0xd5, 0x89, 0xe3, 0x97, 0x13, 0x6b, 0x13, 0x23,
	0x93, 0x00, 0x63, 0x10, 0xb2, 0x59, 0x4a, 0xb6, 0x5c, 0xfc, 0x68, 0x2a, 0x85, 0xb8, 0xb9, 0xd8,
	0x3d, 0x42, 0xfd, 0x5c, 0x82, 0x5c, 0x5d, 0x04, 0x18, 0x84, 0x04, 0xb8, 0x78, 0x42, 0x5c, 0xfd,
	0xe2, 0x43, 0x3c, 0xfc, 0x43, 0x83, 0x1d, 0xfd, 0x5c, 0x04, 0x18, 0x41, 0xd2, 0xbe, 0x9e, 0x3e,
	0x3e, 0x9e, 0xfe, 0x7e, 0x02, 0x4c, 0x4e, 0x5a, 0x5c, 0x12, 0x99, 0xf9, 0x10, 0x87, 0x14, 0x14,
	0xe5, 0x57, 0x54, 0x22, 0xb9, 0xc9, 0x89, 0x07, 0xea, 0x94, 0x00, 0x50, 0x88, 0x05, 0x30, 0x26,
	0xb1, 0x81, 0x83, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xca, 0xbe, 0xff, 0x0b, 0x7b, 0x01,
	0x00, 0x00,
}
