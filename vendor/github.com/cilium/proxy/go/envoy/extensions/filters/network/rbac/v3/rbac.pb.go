// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/filters/network/rbac/v3/rbac.proto

package rbacv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/rbac/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v31 "github.com/cncf/xds/go/xds/type/matcher/v3"
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

type RBAC_EnforcementType int32

const (
	// Apply RBAC policies when the first byte of data arrives on the connection.
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	// Continuously apply RBAC policies as data arrives. Use this mode when
	// using RBAC with message oriented protocols such as Mongo, MySQL, Kafka,
	// etc. when the protocol decoders emit dynamic metadata such as the
	// resources being accessed and the operations on the resources.
	RBAC_CONTINUOUS RBAC_EnforcementType = 1
)

// Enum value maps for RBAC_EnforcementType.
var (
	RBAC_EnforcementType_name = map[int32]string{
		0: "ONE_TIME_ON_FIRST_BYTE",
		1: "CONTINUOUS",
	}
	RBAC_EnforcementType_value = map[string]int32{
		"ONE_TIME_ON_FIRST_BYTE": 0,
		"CONTINUOUS":             1,
	}
)

func (x RBAC_EnforcementType) Enum() *RBAC_EnforcementType {
	p := new(RBAC_EnforcementType)
	*p = x
	return p
}

func (x RBAC_EnforcementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RBAC_EnforcementType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_rbac_v3_rbac_proto_enumTypes[0].Descriptor()
}

func (RBAC_EnforcementType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_rbac_v3_rbac_proto_enumTypes[0]
}

func (x RBAC_EnforcementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RBAC_EnforcementType.Descriptor instead.
func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescGZIP(), []int{0, 0}
}

// RBAC network filter config.
//
// Header should not be used in rules/shadow_rules in RBAC network filter as
// this information is only available in :ref:`RBAC http filter <config_http_filters_rbac>`.
// [#next-free-field: 8]
type RBAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	// If present and empty, DENY.
	// If both rules and matcher are configured, rules will be ignored.
	Rules *v3.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// The match tree to use when resolving RBAC action for incoming connections. Connections do
	// not match any matcher will be denied.
	// If absent, no enforcing RBAC matcher will be applied.
	// If present and empty, deny all connections.
	Matcher *v31.Matcher `protobuf:"bytes,6,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// Shadow rules are not enforced by the filter but will emit stats and logs
	// and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	// If both shadow rules and shadow matcher are configured, shadow rules will be ignored.
	ShadowRules *v3.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	// The match tree to use for emitting stats and logs which can be used for rule testing for
	// incoming connections.
	// If absent, no shadow matcher will be applied.
	ShadowMatcher *v31.Matcher `protobuf:"bytes,7,opt,name=shadow_matcher,json=shadowMatcher,proto3" json:"shadow_matcher,omitempty"`
	// If specified, shadow rules will emit stats with the given prefix.
	// This is useful to distinguish the stat when there are more than 1 RBAC filter configured with
	// shadow rules.
	ShadowRulesStatPrefix string `protobuf:"bytes,5,opt,name=shadow_rules_stat_prefix,json=shadowRulesStatPrefix,proto3" json:"shadow_rules_stat_prefix,omitempty"`
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// RBAC enforcement strategy. By default RBAC will be enforced only once
	// when the first byte of data arrives from the downstream. When used in
	// conjunction with filters that emit dynamic metadata after decoding
	// every payload (e.g., Mongo, MySQL, Kafka) set the enforcement type to
	// CONTINUOUS to enforce RBAC policies on every message boundary.
	EnforcementType RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.extensions.filters.network.rbac.v3.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
}

func (x *RBAC) Reset() {
	*x = RBAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_rbac_v3_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBAC) ProtoMessage() {}

func (x *RBAC) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_rbac_v3_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RBAC.ProtoReflect.Descriptor instead.
func (*RBAC) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *RBAC) GetRules() *v3.RBAC {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RBAC) GetMatcher() *v31.Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *RBAC) GetShadowRules() *v3.RBAC {
	if x != nil {
		return x.ShadowRules
	}
	return nil
}

func (x *RBAC) GetShadowMatcher() *v31.Matcher {
	if x != nil {
		return x.ShadowMatcher
	}
	return nil
}

func (x *RBAC) GetShadowRulesStatPrefix() string {
	if x != nil {
		return x.ShadowRulesStatPrefix
	}
	return ""
}

func (x *RBAC) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if x != nil {
		return x.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

var File_envoy_extensions_filters_network_rbac_v3_rbac_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x62,
	0x61, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb4, 0x05, 0x0a, 0x04, 0x52, 0x42, 0x41, 0x43, 0x12, 0x49, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x52,
	0x42, 0x41, 0x43, 0x42, 0x17, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x11, 0x12, 0x0f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x1f, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x11, 0x12, 0x0f, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0xd2, 0xc6, 0xa4, 0xe1, 0x06,
	0x02, 0x08, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0c,
	0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x42, 0x1e,
	0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x18, 0x12, 0x16, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0b,
	0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x6b, 0x0a, 0x0e, 0x73,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x42, 0x26, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x18, 0x12, 0x16, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x72, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x69, 0x0a, 0x10, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x42, 0x41, 0x43, 0x2e, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x4e, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x42,
	0x59, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55,
	0x4f, 0x55, 0x53, 0x10, 0x01, 0x3a, 0x2f, 0x9a, 0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x42, 0xa5, 0x01, 0x0a, 0x36, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76,
	0x33, 0x42, 0x09, 0x52, 0x62, 0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x33, 0x3b,
	0x72, 0x62, 0x61, 0x63, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescData = file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDesc
)

func file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDescData
}

var file_envoy_extensions_filters_network_rbac_v3_rbac_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_network_rbac_v3_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_rbac_v3_rbac_proto_goTypes = []interface{}{
	(RBAC_EnforcementType)(0), // 0: envoy.extensions.filters.network.rbac.v3.RBAC.EnforcementType
	(*RBAC)(nil),              // 1: envoy.extensions.filters.network.rbac.v3.RBAC
	(*v3.RBAC)(nil),           // 2: envoy.config.rbac.v3.RBAC
	(*v31.Matcher)(nil),       // 3: xds.type.matcher.v3.Matcher
}
var file_envoy_extensions_filters_network_rbac_v3_rbac_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.network.rbac.v3.RBAC.rules:type_name -> envoy.config.rbac.v3.RBAC
	3, // 1: envoy.extensions.filters.network.rbac.v3.RBAC.matcher:type_name -> xds.type.matcher.v3.Matcher
	2, // 2: envoy.extensions.filters.network.rbac.v3.RBAC.shadow_rules:type_name -> envoy.config.rbac.v3.RBAC
	3, // 3: envoy.extensions.filters.network.rbac.v3.RBAC.shadow_matcher:type_name -> xds.type.matcher.v3.Matcher
	0, // 4: envoy.extensions.filters.network.rbac.v3.RBAC.enforcement_type:type_name -> envoy.extensions.filters.network.rbac.v3.RBAC.EnforcementType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_rbac_v3_rbac_proto_init() }
func file_envoy_extensions_filters_network_rbac_v3_rbac_proto_init() {
	if File_envoy_extensions_filters_network_rbac_v3_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_rbac_v3_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RBAC); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_rbac_v3_rbac_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_rbac_v3_rbac_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_rbac_v3_rbac_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_rbac_v3_rbac_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_rbac_v3_rbac_proto = out.File
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_rawDesc = nil
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_goTypes = nil
	file_envoy_extensions_filters_network_rbac_v3_rbac_proto_depIdxs = nil
}
