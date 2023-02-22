// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.7
// source: envoy/extensions/common/ratelimit/v3/ratelimit.proto

package ratelimitv3

import (
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Defines the version of the standard to use for X-RateLimit headers.
type XRateLimitHeadersRFCVersion int32

const (
	// X-RateLimit headers disabled.
	XRateLimitHeadersRFCVersion_OFF XRateLimitHeadersRFCVersion = 0
	// Use `draft RFC Version 03 <https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html>`_ where 3 headers will be added:
	//
	//   - “X-RateLimit-Limit“ - indicates the request-quota associated to the
	//     client in the current time-window followed by the description of the
	//     quota policy. The value is returned by the maximum tokens of the token bucket.
	//   - “X-RateLimit-Remaining“ - indicates the remaining requests in the
	//     current time-window. The value is returned by the remaining tokens in the token bucket.
	//   - “X-RateLimit-Reset“ - indicates the number of seconds until reset of
	//     the current time-window. The value is returned by the remaining fill interval of the token bucket.
	XRateLimitHeadersRFCVersion_DRAFT_VERSION_03 XRateLimitHeadersRFCVersion = 1
)

// Enum value maps for XRateLimitHeadersRFCVersion.
var (
	XRateLimitHeadersRFCVersion_name = map[int32]string{
		0: "OFF",
		1: "DRAFT_VERSION_03",
	}
	XRateLimitHeadersRFCVersion_value = map[string]int32{
		"OFF":              0,
		"DRAFT_VERSION_03": 1,
	}
)

func (x XRateLimitHeadersRFCVersion) Enum() *XRateLimitHeadersRFCVersion {
	p := new(XRateLimitHeadersRFCVersion)
	*p = x
	return p
}

func (x XRateLimitHeadersRFCVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XRateLimitHeadersRFCVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes[0].Descriptor()
}

func (XRateLimitHeadersRFCVersion) Type() protoreflect.EnumType {
	return &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes[0]
}

func (x XRateLimitHeadersRFCVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XRateLimitHeadersRFCVersion.Descriptor instead.
func (XRateLimitHeadersRFCVersion) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{0}
}

type VhRateLimitsOptions int32

const (
	// Use the virtual host rate limits unless the route has a rate limit policy.
	VhRateLimitsOptions_OVERRIDE VhRateLimitsOptions = 0
	// Use the virtual host rate limits even if the route has a rate limit policy.
	VhRateLimitsOptions_INCLUDE VhRateLimitsOptions = 1
	// Ignore the virtual host rate limits even if the route does not have a rate limit policy.
	VhRateLimitsOptions_IGNORE VhRateLimitsOptions = 2
)

// Enum value maps for VhRateLimitsOptions.
var (
	VhRateLimitsOptions_name = map[int32]string{
		0: "OVERRIDE",
		1: "INCLUDE",
		2: "IGNORE",
	}
	VhRateLimitsOptions_value = map[string]int32{
		"OVERRIDE": 0,
		"INCLUDE":  1,
		"IGNORE":   2,
	}
)

func (x VhRateLimitsOptions) Enum() *VhRateLimitsOptions {
	p := new(VhRateLimitsOptions)
	*p = x
	return p
}

func (x VhRateLimitsOptions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VhRateLimitsOptions) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes[1].Descriptor()
}

func (VhRateLimitsOptions) Type() protoreflect.EnumType {
	return &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes[1]
}

func (x VhRateLimitsOptions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VhRateLimitsOptions.Descriptor instead.
func (VhRateLimitsOptions) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{1}
}

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//	["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//	["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//	["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//	["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//	["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
//
// Optionally the descriptor can contain a limit override under a "limit" key, that specifies
// the number of requests per unit to use instead of the number configured in the
// rate limiting service.
type RateLimitDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Descriptor entries.
	Entries []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// Optional rate limit override to supply to the ratelimit service.
	Limit *RateLimitDescriptor_RateLimitOverride `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RateLimitDescriptor) Reset() {
	*x = RateLimitDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescriptor) ProtoMessage() {}

func (x *RateLimitDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescriptor.ProtoReflect.Descriptor instead.
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *RateLimitDescriptor) GetLimit() *RateLimitDescriptor_RateLimitOverride {
	if x != nil {
		return x.Limit
	}
	return nil
}

type LocalRateLimitDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Descriptor entries.
	Entries []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// Token Bucket algorithm for local ratelimiting.
	TokenBucket *v3.TokenBucket `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
}

func (x *LocalRateLimitDescriptor) Reset() {
	*x = LocalRateLimitDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRateLimitDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRateLimitDescriptor) ProtoMessage() {}

func (x *LocalRateLimitDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRateLimitDescriptor.ProtoReflect.Descriptor instead.
func (*LocalRateLimitDescriptor) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{1}
}

func (x *LocalRateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *LocalRateLimitDescriptor) GetTokenBucket() *v3.TokenBucket {
	if x != nil {
		return x.TokenBucket
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RateLimitDescriptor_Entry) Reset() {
	*x = RateLimitDescriptor_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescriptor_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescriptor_Entry) ProtoMessage() {}

func (x *RateLimitDescriptor_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescriptor_Entry.ProtoReflect.Descriptor instead.
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RateLimitDescriptor_Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RateLimitDescriptor_Entry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Override rate limit to apply to this descriptor instead of the limit
// configured in the rate limit service. See :ref:`rate limit override
// <config_http_filters_rate_limit_rate_limit_override>` for more information.
type RateLimitDescriptor_RateLimitOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of requests per unit of time.
	RequestsPerUnit uint32 `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// The unit of time.
	Unit v3.RateLimitUnit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.type.v3.RateLimitUnit" json:"unit,omitempty"`
}

func (x *RateLimitDescriptor_RateLimitOverride) Reset() {
	*x = RateLimitDescriptor_RateLimitOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescriptor_RateLimitOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescriptor_RateLimitOverride) ProtoMessage() {}

func (x *RateLimitDescriptor_RateLimitOverride) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescriptor_RateLimitOverride.ProtoReflect.Descriptor instead.
func (*RateLimitDescriptor_RateLimitOverride) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RateLimitDescriptor_RateLimitOverride) GetRequestsPerUnit() uint32 {
	if x != nil {
		return x.RequestsPerUnit
	}
	return 0
}

func (x *RateLimitDescriptor_RateLimitOverride) GetUnit() v3.RateLimitUnit {
	if x != nil {
		return x.Unit
	}
	return v3.RateLimitUnit_UNKNOWN
}

var File_envoy_extensions_common_ratelimit_v3_ratelimit_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04,
	0x0a, 0x13, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x63, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x7a, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x7b, 0x0a, 0x11, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x3a, 0x31, 0x9a, 0xc5, 0x88, 0x1e, 0x2c, 0x0a, 0x2a, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x22, 0xc8, 0x01, 0x0a, 0x18, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x63, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2a, 0x3c, 0x0a, 0x1b, 0x58, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x46, 0x43, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x30, 0x33,
	0x10, 0x01, 0x2a, 0x3c, 0x0a, 0x13, 0x56, 0x68, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45,
	0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x43, 0x4c, 0x55,
	0x44, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x02,
	0x42, 0xa7, 0x01, 0x0a, 0x32, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescData = file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDesc
)

func file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescData)
	})
	return file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDescData
}

var file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_goTypes = []interface{}{
	(XRateLimitHeadersRFCVersion)(0),              // 0: envoy.extensions.common.ratelimit.v3.XRateLimitHeadersRFCVersion
	(VhRateLimitsOptions)(0),                      // 1: envoy.extensions.common.ratelimit.v3.VhRateLimitsOptions
	(*RateLimitDescriptor)(nil),                   // 2: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor
	(*LocalRateLimitDescriptor)(nil),              // 3: envoy.extensions.common.ratelimit.v3.LocalRateLimitDescriptor
	(*RateLimitDescriptor_Entry)(nil),             // 4: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.Entry
	(*RateLimitDescriptor_RateLimitOverride)(nil), // 5: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.RateLimitOverride
	(*v3.TokenBucket)(nil),                        // 6: envoy.type.v3.TokenBucket
	(v3.RateLimitUnit)(0),                         // 7: envoy.type.v3.RateLimitUnit
}
var file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.entries:type_name -> envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.Entry
	5, // 1: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.limit:type_name -> envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.RateLimitOverride
	4, // 2: envoy.extensions.common.ratelimit.v3.LocalRateLimitDescriptor.entries:type_name -> envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.Entry
	6, // 3: envoy.extensions.common.ratelimit.v3.LocalRateLimitDescriptor.token_bucket:type_name -> envoy.type.v3.TokenBucket
	7, // 4: envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.RateLimitOverride.unit:type_name -> envoy.type.v3.RateLimitUnit
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_init() }
func file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_init() {
	if File_envoy_extensions_common_ratelimit_v3_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitDescriptor); i {
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
		file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRateLimitDescriptor); i {
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
		file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitDescriptor_Entry); i {
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
		file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitDescriptor_RateLimitOverride); i {
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
			RawDescriptor: file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_ratelimit_v3_ratelimit_proto = out.File
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_rawDesc = nil
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_goTypes = nil
	file_envoy_extensions_common_ratelimit_v3_ratelimit_proto_depIdxs = nil
}
