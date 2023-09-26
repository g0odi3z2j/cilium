// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/filters/http/gcp_authn/v3/gcp_authn.proto

package gcp_authnv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Filter configuration.
type GcpAuthnFilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP URI to fetch tokens from GCE Metadata Server(https://cloud.google.com/compute/docs/metadata/overview).
	// The URL format is "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=[AUDIENCE]"
	HttpUri *v3.HttpUri `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	// Retry policy for fetching tokens. This field is optional.
	RetryPolicy *v3.RetryPolicy `protobuf:"bytes,2,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
	// Token cache configuration. This field is optional.
	CacheConfig *TokenCacheConfig `protobuf:"bytes,3,opt,name=cache_config,json=cacheConfig,proto3" json:"cache_config,omitempty"`
	// Request header location to extract the token. By default (i.e. if this field is not specified), the token
	// is extracted to the Authorization HTTP header, in the format "Authorization: Bearer <token>".
	TokenHeader *TokenHeader `protobuf:"bytes,4,opt,name=token_header,json=tokenHeader,proto3" json:"token_header,omitempty"`
}

func (x *GcpAuthnFilterConfig) Reset() {
	*x = GcpAuthnFilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpAuthnFilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpAuthnFilterConfig) ProtoMessage() {}

func (x *GcpAuthnFilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpAuthnFilterConfig.ProtoReflect.Descriptor instead.
func (*GcpAuthnFilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{0}
}

func (x *GcpAuthnFilterConfig) GetHttpUri() *v3.HttpUri {
	if x != nil {
		return x.HttpUri
	}
	return nil
}

func (x *GcpAuthnFilterConfig) GetRetryPolicy() *v3.RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

func (x *GcpAuthnFilterConfig) GetCacheConfig() *TokenCacheConfig {
	if x != nil {
		return x.CacheConfig
	}
	return nil
}

func (x *GcpAuthnFilterConfig) GetTokenHeader() *TokenHeader {
	if x != nil {
		return x.TokenHeader
	}
	return nil
}

// Audience is the URL of the receiving service that performs token authentication.
// It will be provided to the filter through cluster's typed_filter_metadata.
type Audience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Audience) Reset() {
	*x = Audience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audience) ProtoMessage() {}

func (x *Audience) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audience.ProtoReflect.Descriptor instead.
func (*Audience) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{1}
}

func (x *Audience) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Token Cache configuration.
type TokenCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of cache entries. The maximum number of entries is INT64_MAX as it is constrained by underlying cache implementation.
	// Default value 0 (i.e., proto3 defaults) disables the cache by default. Other default values will enable the cache.
	CacheSize *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=cache_size,json=cacheSize,proto3" json:"cache_size,omitempty"`
}

func (x *TokenCacheConfig) Reset() {
	*x = TokenCacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCacheConfig) ProtoMessage() {}

func (x *TokenCacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCacheConfig.ProtoReflect.Descriptor instead.
func (*TokenCacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{2}
}

func (x *TokenCacheConfig) GetCacheSize() *wrapperspb.UInt64Value {
	if x != nil {
		return x.CacheSize
	}
	return nil
}

type TokenHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP header's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The header's prefix. The format is "value_prefix<token>"
	// For example, for "Authorization: Bearer <token>", value_prefix="Bearer " with a space at the
	// end.
	ValuePrefix string `protobuf:"bytes,2,opt,name=value_prefix,json=valuePrefix,proto3" json:"value_prefix,omitempty"`
}

func (x *TokenHeader) Reset() {
	*x = TokenHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenHeader) ProtoMessage() {}

func (x *TokenHeader) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenHeader.ProtoReflect.Descriptor instead.
func (*TokenHeader) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{3}
}

func (x *TokenHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenHeader) GetValuePrefix() string {
	if x != nil {
		return x.ValuePrefix
	}
	return ""
}

var File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x63, 0x70,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x14, 0x47, 0x63, 0x70, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x42, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x68, 0x74, 0x74, 0x70,
	0x55, 0x72, 0x69, 0x12, 0x44, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5f, 0x0a, 0x0c, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x0c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x60, 0x0a,
	0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4c, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x32, 0x0a, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x7f, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x60, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42,
	0x0a, 0x72, 0x08, 0x10, 0x01, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01,
	0x02, 0xc8, 0x01, 0x00, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x42, 0xb2, 0x01, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0d,
	0x47, 0x63, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f,
	0x76, 0x33, 0x3b, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData = file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc
)

func file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData
}

var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes = []interface{}{
	(*GcpAuthnFilterConfig)(nil),   // 0: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig
	(*Audience)(nil),               // 1: envoy.extensions.filters.http.gcp_authn.v3.Audience
	(*TokenCacheConfig)(nil),       // 2: envoy.extensions.filters.http.gcp_authn.v3.TokenCacheConfig
	(*TokenHeader)(nil),            // 3: envoy.extensions.filters.http.gcp_authn.v3.TokenHeader
	(*v3.HttpUri)(nil),             // 4: envoy.config.core.v3.HttpUri
	(*v3.RetryPolicy)(nil),         // 5: envoy.config.core.v3.RetryPolicy
	(*wrapperspb.UInt64Value)(nil), // 6: google.protobuf.UInt64Value
}
var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.http_uri:type_name -> envoy.config.core.v3.HttpUri
	5, // 1: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.retry_policy:type_name -> envoy.config.core.v3.RetryPolicy
	2, // 2: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.cache_config:type_name -> envoy.extensions.filters.http.gcp_authn.v3.TokenCacheConfig
	3, // 3: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.token_header:type_name -> envoy.extensions.filters.http.gcp_authn.v3.TokenHeader
	6, // 4: envoy.extensions.filters.http.gcp_authn.v3.TokenCacheConfig.cache_size:type_name -> google.protobuf.UInt64Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_init() }
func file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_init() {
	if File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpAuthnFilterConfig); i {
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
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audience); i {
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
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenCacheConfig); i {
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
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenHeader); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto = out.File
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc = nil
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes = nil
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs = nil
}
