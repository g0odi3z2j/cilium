// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool     `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute" json:"match_incoming_request_route,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_5b78334701011728, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (dst *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(dst, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,oneof"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GrpcJsonTranscoder_OneofMarshaler, _GrpcJsonTranscoder_OneofUnmarshaler, _GrpcJsonTranscoder_OneofSizer, []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

func _GrpcJsonTranscoder_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ProtoDescriptorBin)
	case nil:
	default:
		return fmt.Errorf("GrpcJsonTranscoder.DescriptorSet has unexpected type %T", x)
	}
	return nil
}

func _GrpcJsonTranscoder_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GrpcJsonTranscoder)
	switch tag {
	case 1: // descriptor_set.proto_descriptor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{x}
		return true, err
	case 4: // descriptor_set.proto_descriptor_bin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{x}
		return true, err
	default:
		return false, nil
	}
}

func _GrpcJsonTranscoder_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptor)))
		n += len(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptorBin)))
		n += len(x.ProtoDescriptorBin)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_5b78334701011728, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (dst *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(dst, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_transcoder_5b78334701011728)
}

var fileDescriptor_transcoder_5b78334701011728 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x6a, 0xd4, 0x4c,
	0x18, 0xc6, 0x9b, 0xfd, 0xf3, 0x91, 0xce, 0xb7, 0xad, 0x32, 0x88, 0x1b, 0x83, 0x42, 0x10, 0x2c,
	0x01, 0x21, 0x81, 0x78, 0x20, 0xe8, 0x81, 0x74, 0x51, 0xdb, 0x7a, 0xa0, 0xcb, 0x20, 0x08, 0x9e,
	0x0c, 0xd3, 0xe4, 0xdd, 0xee, 0x40, 0x32, 0x33, 0xce, 0xcc, 0xa6, 0xf4, 0x36, 0xbc, 0x1a, 0xf1,
	0xa8, 0xb7, 0xe1, 0x25, 0xf4, 0xd0, 0x3b, 0x90, 0x99, 0x71, 0x97, 0x58, 0x4f, 0x3c, 0xcb, 0xe4,
	0xf9, 0x3d, 0xcf, 0xcb, 0xfb, 0xbc, 0xe8, 0x39, 0x88, 0x5e, 0x5e, 0x95, 0xb5, 0x14, 0x2b, 0x7e,
	0x51, 0xae, 0x78, 0x6b, 0x41, 0x97, 0x6b, 0x6b, 0x55, 0x69, 0x35, 0x13, 0xa6, 0x96, 0x0d, 0xe8,
	0xb2, 0xaf, 0x06, 0xaf, 0x42, 0x69, 0x69, 0x25, 0x3e, 0xf2, 0xc6, 0x22, 0x18, 0x8b, 0x60, 0x2c,
	0x9c, 0xb1, 0x18, 0xa0, 0x7d, 0x95, 0xce, 0x7b, 0xd6, 0xf2, 0x86, 0x59, 0x28, 0xb7, 0x1f, 0x21,
	0xe0, 0xf1, 0x8f, 0x09, 0xc2, 0x27, 0x5a, 0xd5, 0xef, 0x8c, 0x14, 0x1f, 0x77, 0x16, 0xfc, 0x14,
	0xdd, 0xf5, 0x3a, 0x6d, 0xc0, 0xd4, 0x9a, 0x2b, 0x2b, 0x75, 0x12, 0x65, 0x51, 0xbe, 0x7f, 0xba,
	0x47, 0xee, 0x78, 0xe5, 0xf5, 0x4e, 0xc0, 0x15, 0xba, 0x77, 0x1b, 0xa6, 0xe7, 0x5c, 0x24, 0x93,
	0x2c, 0xca, 0x67, 0xa7, 0x7b, 0x04, 0xdf, 0x32, 0x2c, 0xb8, 0xc0, 0x47, 0x28, 0x36, 0xa0, 0x7b,
	0x5e, 0x83, 0x49, 0x46, 0xd9, 0x38, 0xdf, 0x5f, 0xa0, 0xef, 0x37, 0xd7, 0xe3, 0xe9, 0xd7, 0x68,
	0x14, 0x47, 0x64, 0xa7, 0xe1, 0x16, 0x1d, 0x28, 0xcd, 0x85, 0xa5, 0x52, 0x59, 0x2e, 0x85, 0x49,
	0xc6, 0x59, 0x94, 0xff, 0x5f, 0x9d, 0x14, 0xff, 0xb6, 0x78, 0xf1, 0xf7, 0x6e, 0xc5, 0xd2, 0xe5,
	0x7d, 0x08, 0x71, 0x64, 0xa6, 0x06, 0x2f, 0xfc, 0x0a, 0x3d, 0xec, 0x98, 0xad, 0xd7, 0x94, 0x8b,
	0x5a, 0x76, 0x5c, 0x5c, 0x50, 0x0d, 0x5f, 0x36, 0x60, 0x2c, 0xd5, 0x72, 0x63, 0x21, 0x99, 0x66,
	0x51, 0x1e, 0x93, 0x07, 0x9e, 0x39, 0xfb, 0x8d, 0x90, 0x40, 0x10, 0x07, 0xa4, 0x3f, 0x23, 0x34,
	0x1b, 0xe6, 0xe3, 0x27, 0xe8, 0x90, 0x35, 0x0d, 0xbd, 0x5c, 0x73, 0x0b, 0x46, 0xb1, 0x1a, 0x7c,
	0x8d, 0x31, 0x39, 0x60, 0x4d, 0xf3, 0x69, 0xf7, 0x13, 0x1f, 0xa3, 0x47, 0xac, 0xbd, 0x64, 0x57,
	0x86, 0x86, 0x6d, 0x95, 0xe6, 0x1d, 0xb7, 0xbc, 0x07, 0xba, 0xe2, 0xd0, 0x36, 0xae, 0x23, 0xe7,
	0x4a, 0x03, 0xe4, 0x27, 0x2c, 0xb7, 0xc8, 0x5b, 0x4f, 0xe0, 0x17, 0x28, 0xfd, 0x23, 0x02, 0xc4,
	0xa6, 0x33, 0x94, 0x19, 0xca, 0x85, 0x0d, 0xb5, 0xc5, 0xe4, 0xfe, 0xc0, 0xff, 0xc6, 0xe9, 0xc7,
	0xe6, 0x4c, 0x58, 0x83, 0x5f, 0xa2, 0x54, 0x69, 0x70, 0xa5, 0x03, 0x0d, 0xa7, 0xf4, 0x63, 0xa9,
	0x60, 0x1d, 0x18, 0x7f, 0xc7, 0x98, 0xcc, 0xb7, 0xc4, 0xd2, 0x01, 0x7e, 0xe8, 0x7b, 0x27, 0x2f,
	0xe6, 0xe8, 0x70, 0x70, 0x78, 0x03, 0x16, 0x4f, 0xbf, 0xdd, 0x5c, 0x8f, 0xa3, 0xc5, 0xe4, 0xf3,
	0xa8, 0xaf, 0xce, 0xff, 0xf3, 0x81, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xa4, 0xd9,
	0x9e, 0xe4, 0x02, 0x00, 0x00,
}
