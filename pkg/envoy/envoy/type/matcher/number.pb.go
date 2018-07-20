// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/number.proto

package matcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_type "github.com/cilium/cilium/pkg/envoy/envoy/type"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies the way to match a double value.
type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
}

func (m *DoubleMatcher) Reset()                    { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string            { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()               {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	Range *envoy_type.DoubleRange `protobuf:"bytes,1,opt,name=range,oneof"`
}
type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}
func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *envoy_type.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DoubleMatcher_OneofMarshaler, _DoubleMatcher_OneofUnmarshaler, _DoubleMatcher_OneofSizer, []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func _DoubleMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DoubleMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *DoubleMatcher_Range:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Range); err != nil {
			return err
		}
	case *DoubleMatcher_Exact:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Exact))
	case nil:
	default:
		return fmt.Errorf("DoubleMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _DoubleMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DoubleMatcher)
	switch tag {
	case 1: // match_pattern.range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(envoy_type.DoubleRange)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &DoubleMatcher_Range{msg}
		return true, err
	case 2: // match_pattern.exact
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.MatchPattern = &DoubleMatcher_Exact{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _DoubleMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DoubleMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *DoubleMatcher_Range:
		s := proto.Size(x.Range)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DoubleMatcher_Exact:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.DoubleMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/number.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x2b, 0xd0,
	0x03, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x43, 0xd2, 0x54, 0x94, 0x98, 0x97, 0x9e, 0x0a, 0x51,
	0x2b, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94,
	0x0a, 0xb8, 0x78, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0x7d, 0x21, 0x26, 0x08, 0xe9, 0x73, 0xb1,
	0x82, 0x35, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x21, 0xd9, 0x02, 0x51, 0x19,
	0x04, 0x92, 0xf6, 0x60, 0x08, 0x82, 0xa8, 0x13, 0x12, 0xe3, 0x62, 0x4d, 0xad, 0x48, 0x4c, 0x2e,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x04, 0x89, 0x83, 0xb9, 0x4e, 0x62, 0x5c, 0xbc, 0x60, 0x57,
	0xc5, 0x17, 0x24, 0x96, 0x94, 0xa4, 0x16, 0xe5, 0x09, 0xb1, 0xee, 0x78, 0x79, 0x80, 0x99, 0xd1,
	0x89, 0x33, 0x8a, 0x1d, 0xea, 0xda, 0x24, 0x36, 0xb0, 0x1b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0xb0, 0xc9, 0x95, 0xeb, 0x00, 0x00, 0x00,
}
