// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/range.proto

package envoy_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	// end of the range (exclusive)
	End int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *Int64Range) Reset()                    { *m = Int64Range{} }
func (m *Int64Range) String() string            { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()               {}
func (*Int64Range) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.Int64Range")
}

func init() { proto.RegisterFile("envoy/type/range.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x8b, 0xeb, 0x81, 0xc4, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0xc2, 0xfa, 0x20, 0x16, 0x44, 0x85, 0x92, 0x09, 0x17, 0x97, 0x67, 0x5e, 0x89, 0x99, 0x49, 0x10,
	0x48, 0x97, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x49, 0x62, 0x51, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x73, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x9a, 0x97, 0x22, 0xc1, 0x04, 0x16, 0x03, 0x31,
	0x9d, 0x04, 0x56, 0x3c, 0x92, 0x63, 0x8c, 0x82, 0x98, 0x1e, 0x0f, 0x32, 0x3d, 0x89, 0x0d, 0x6c,
	0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x58, 0x8e, 0xb2, 0x8a, 0x00, 0x00, 0x00,
}
