// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/type/v3/token_bucket.proto

package typev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configures a token bucket, typically used for rate limiting.
type TokenBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum tokens that the bucket can hold. This is also the number of tokens that the bucket
	// initially contains.
	MaxTokens uint32 `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// The number of tokens added to the bucket during each fill interval. If not specified, defaults
	// to a single token.
	TokensPerFill *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	// The fill interval that tokens are added to the bucket. During each fill interval
	// `tokens_per_fill` are added to the bucket. The bucket will never contain more than
	// `max_tokens` tokens.
	FillInterval *durationpb.Duration `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
}

func (x *TokenBucket) Reset() {
	*x = TokenBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_v3_token_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBucket) ProtoMessage() {}

func (x *TokenBucket) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_v3_token_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBucket.ProtoReflect.Descriptor instead.
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return file_envoy_type_v3_token_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *TokenBucket) GetMaxTokens() uint32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *TokenBucket) GetTokensPerFill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TokensPerFill
	}
	return nil
}

func (x *TokenBucket) GetFillInterval() *durationpb.Duration {
	if x != nil {
		return x.FillInterval
	}
	return nil
}

var File_envoy_type_v3_token_bucket_proto protoreflect.FileDescriptor

var file_envoy_type_v3_token_bucket_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a,
	0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x6c, 0x12, 0x4a, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a,
	0x00, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x3a,
	0x1d, 0x9a, 0xc5, 0x88, 0x1e, 0x18, 0x0a, 0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x76,
	0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_v3_token_bucket_proto_rawDescOnce sync.Once
	file_envoy_type_v3_token_bucket_proto_rawDescData = file_envoy_type_v3_token_bucket_proto_rawDesc
)

func file_envoy_type_v3_token_bucket_proto_rawDescGZIP() []byte {
	file_envoy_type_v3_token_bucket_proto_rawDescOnce.Do(func() {
		file_envoy_type_v3_token_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_v3_token_bucket_proto_rawDescData)
	})
	return file_envoy_type_v3_token_bucket_proto_rawDescData
}

var file_envoy_type_v3_token_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_type_v3_token_bucket_proto_goTypes = []interface{}{
	(*TokenBucket)(nil),            // 0: envoy.type.v3.TokenBucket
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_envoy_type_v3_token_bucket_proto_depIdxs = []int32{
	1, // 0: envoy.type.v3.TokenBucket.tokens_per_fill:type_name -> google.protobuf.UInt32Value
	2, // 1: envoy.type.v3.TokenBucket.fill_interval:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_type_v3_token_bucket_proto_init() }
func file_envoy_type_v3_token_bucket_proto_init() {
	if File_envoy_type_v3_token_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_v3_token_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBucket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_v3_token_bucket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_v3_token_bucket_proto_goTypes,
		DependencyIndexes: file_envoy_type_v3_token_bucket_proto_depIdxs,
		MessageInfos:      file_envoy_type_v3_token_bucket_proto_msgTypes,
	}.Build()
	File_envoy_type_v3_token_bucket_proto = out.File
	file_envoy_type_v3_token_bucket_proto_rawDesc = nil
	file_envoy_type_v3_token_bucket_proto_goTypes = nil
	file_envoy_type_v3_token_bucket_proto_depIdxs = nil
}
