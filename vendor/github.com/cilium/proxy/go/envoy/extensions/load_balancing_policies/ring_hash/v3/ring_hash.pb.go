// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/load_balancing_policies/ring_hash/v3/ring_hash.proto

package ring_hashv3

import (
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/extensions/load_balancing_policies/common/v3"
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

// The hash function used to hash hosts onto the ketama ring.
type RingHash_HashFunction int32

const (
	// Currently defaults to XX_HASH.
	RingHash_DEFAULT_HASH RingHash_HashFunction = 0
	// Use `xxHash <https://github.com/Cyan4973/xxHash>`_.
	RingHash_XX_HASH RingHash_HashFunction = 1
	// Use `MurmurHash2 <https://sites.google.com/site/murmurhash/>`_, this is compatible with
	// std:hash<string> in GNU libstdc++ 3.4.20 or above. This is typically the case when compiled
	// on Linux and not macOS.
	RingHash_MURMUR_HASH_2 RingHash_HashFunction = 2
)

// Enum value maps for RingHash_HashFunction.
var (
	RingHash_HashFunction_name = map[int32]string{
		0: "DEFAULT_HASH",
		1: "XX_HASH",
		2: "MURMUR_HASH_2",
	}
	RingHash_HashFunction_value = map[string]int32{
		"DEFAULT_HASH":  0,
		"XX_HASH":       1,
		"MURMUR_HASH_2": 2,
	}
)

func (x RingHash_HashFunction) Enum() *RingHash_HashFunction {
	p := new(RingHash_HashFunction)
	*p = x
	return p
}

func (x RingHash_HashFunction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RingHash_HashFunction) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_enumTypes[0].Descriptor()
}

func (RingHash_HashFunction) Type() protoreflect.EnumType {
	return &file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_enumTypes[0]
}

func (x RingHash_HashFunction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RingHash_HashFunction.Descriptor instead.
func (RingHash_HashFunction) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescGZIP(), []int{0, 0}
}

// This configuration allows the built-in RING_HASH LB policy to be configured via the LB policy
// extension point. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` for more information.
// [#next-free-field: 8]
type RingHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash function used to hash hosts onto the ketama ring. The value defaults to
	// :ref:`XX_HASH<envoy_v3_api_enum_value_config.cluster.v3.Cluster.RingHashLbConfig.HashFunction.XX_HASH>`.
	HashFunction RingHash_HashFunction `protobuf:"varint,1,opt,name=hash_function,json=hashFunction,proto3,enum=envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash_HashFunction" json:"hash_function,omitempty"`
	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each
	// provided host) the better the request distribution will reflect the desired weights. Defaults
	// to 1024 entries, and limited to 8M entries. See also
	// :ref:`maximum_ring_size<envoy_v3_api_field_config.cluster.v3.Cluster.RingHashLbConfig.maximum_ring_size>`.
	MinimumRingSize *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered
	// to further constrain resource use. See also
	// :ref:`minimum_ring_size<envoy_v3_api_field_config.cluster.v3.Cluster.RingHashLbConfig.minimum_ring_size>`.
	MaximumRingSize *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
	// If set to `true`, the cluster will use hostname instead of the resolved
	// address as the key to consistently hash to an upstream host. Only valid for StrictDNS clusters with hostnames which resolve to a single IP address.
	//
	// ..note::
	//
	//	This is deprecated and please use :ref:`consistent_hashing_lb_config
	//	<envoy_v3_api_field_extensions.load_balancing_policies.ring_hash.v3.RingHash.consistent_hashing_lb_config>` instead.
	//
	// Deprecated: Do not use.
	UseHostnameForHashing bool `protobuf:"varint,4,opt,name=use_hostname_for_hashing,json=useHostnameForHashing,proto3" json:"use_hostname_for_hashing,omitempty"`
	// Configures percentage of average cluster load to bound per upstream host. For example, with a value of 150
	// no upstream host will get a load more than 1.5 times the average load of all the hosts in the cluster.
	// If not specified, the load is not bounded for any upstream host. Typical value for this parameter is between 120 and 200.
	// Minimum is 100.
	//
	// This is implemented based on the method described in the paper https://arxiv.org/abs/1608.01350. For the specified
	// `hash_balance_factor`, requests to any upstream host are capped at `hash_balance_factor/100` times the average number of requests
	// across the cluster. When a request arrives for an upstream host that is currently serving at its max capacity, linear probing
	// is used to identify an eligible host. Further, the linear probe is implemented using a random jump in hosts ring/table to identify
	// the eligible host (this technique is as described in the paper https://arxiv.org/abs/1908.08762 - the random jump avoids the
	// cascading overflow effect when choosing the next host in the ring/table).
	//
	// If weights are specified on the hosts, they are respected.
	//
	// This is an O(N) algorithm, unlike other load balancers. Using a lower `hash_balance_factor` results in more hosts
	// being probed, so use a higher value if you require better performance.
	//
	// ..note::
	//
	//	This is deprecated and please use :ref:`consistent_hashing_lb_config
	//	<envoy_v3_api_field_extensions.load_balancing_policies.ring_hash.v3.RingHash.consistent_hashing_lb_config>` instead.
	//
	// Deprecated: Do not use.
	HashBalanceFactor *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=hash_balance_factor,json=hashBalanceFactor,proto3" json:"hash_balance_factor,omitempty"`
	// Common configuration for hashing-based load balancing policies.
	ConsistentHashingLbConfig *v3.ConsistentHashingLbConfig `protobuf:"bytes,6,opt,name=consistent_hashing_lb_config,json=consistentHashingLbConfig,proto3" json:"consistent_hashing_lb_config,omitempty"`
	// Enable locality weighted load balancing for ring hash lb explicitly.
	LocalityWeightedLbConfig *v3.LocalityLbConfig_LocalityWeightedLbConfig `protobuf:"bytes,7,opt,name=locality_weighted_lb_config,json=localityWeightedLbConfig,proto3" json:"locality_weighted_lb_config,omitempty"`
}

func (x *RingHash) Reset() {
	*x = RingHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingHash) ProtoMessage() {}

func (x *RingHash) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingHash.ProtoReflect.Descriptor instead.
func (*RingHash) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescGZIP(), []int{0}
}

func (x *RingHash) GetHashFunction() RingHash_HashFunction {
	if x != nil {
		return x.HashFunction
	}
	return RingHash_DEFAULT_HASH
}

func (x *RingHash) GetMinimumRingSize() *wrapperspb.UInt64Value {
	if x != nil {
		return x.MinimumRingSize
	}
	return nil
}

func (x *RingHash) GetMaximumRingSize() *wrapperspb.UInt64Value {
	if x != nil {
		return x.MaximumRingSize
	}
	return nil
}

// Deprecated: Do not use.
func (x *RingHash) GetUseHostnameForHashing() bool {
	if x != nil {
		return x.UseHostnameForHashing
	}
	return false
}

// Deprecated: Do not use.
func (x *RingHash) GetHashBalanceFactor() *wrapperspb.UInt32Value {
	if x != nil {
		return x.HashBalanceFactor
	}
	return nil
}

func (x *RingHash) GetConsistentHashingLbConfig() *v3.ConsistentHashingLbConfig {
	if x != nil {
		return x.ConsistentHashingLbConfig
	}
	return nil
}

func (x *RingHash) GetLocalityWeightedLbConfig() *v3.LocalityLbConfig_LocalityWeightedLbConfig {
	if x != nil {
		return x.LocalityWeightedLbConfig
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x1a, 0x3f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x06, 0x0a,
	0x08, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x7b, 0x0a, 0x0d, 0x68, 0x61, 0x73,
	0x68, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73,
	0x68, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x32, 0x05, 0x18, 0x80, 0x80, 0x80, 0x04, 0x52, 0x0f, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x54, 0x0a, 0x11,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x32, 0x05, 0x18, 0x80, 0x80, 0x80,
	0x04, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x44, 0x0a, 0x18, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0x18, 0x01, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e,
	0x30, 0x52, 0x15, 0x75, 0x73, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x6f,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x60, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x68,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x12, 0x18, 0x01, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x64, 0x92, 0xc7,
	0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x52, 0x11, 0x68, 0x61, 0x73, 0x68, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x8e, 0x01, 0x0a, 0x1c, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x4d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x19, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x9c, 0x01, 0x0a, 0x1b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x5d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x64, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x40, 0x0a, 0x0c, 0x48, 0x61,
	0x73, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x58, 0x58, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x55, 0x52,
	0x4d, 0x55, 0x52, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x32, 0x10, 0x02, 0x42, 0xc8, 0x01, 0x0a,
	0x43, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x2f, 0x76, 0x33, 0x3b, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x76, 0x33, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescData = file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_goTypes = []interface{}{
	(RingHash_HashFunction)(0),                           // 0: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.HashFunction
	(*RingHash)(nil),                                     // 1: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash
	(*wrapperspb.UInt64Value)(nil),                       // 2: google.protobuf.UInt64Value
	(*wrapperspb.UInt32Value)(nil),                       // 3: google.protobuf.UInt32Value
	(*v3.ConsistentHashingLbConfig)(nil),                 // 4: envoy.extensions.load_balancing_policies.common.v3.ConsistentHashingLbConfig
	(*v3.LocalityLbConfig_LocalityWeightedLbConfig)(nil), // 5: envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig.LocalityWeightedLbConfig
}
var file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.hash_function:type_name -> envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.HashFunction
	2, // 1: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.minimum_ring_size:type_name -> google.protobuf.UInt64Value
	2, // 2: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.maximum_ring_size:type_name -> google.protobuf.UInt64Value
	3, // 3: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.hash_balance_factor:type_name -> google.protobuf.UInt32Value
	4, // 4: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.consistent_hashing_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.ConsistentHashingLbConfig
	5, // 5: envoy.extensions.load_balancing_policies.ring_hash.v3.RingHash.locality_weighted_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig.LocalityWeightedLbConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_init() }
func file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_init() {
	if File_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingHash); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto = out.File
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_ring_hash_v3_ring_hash_proto_depIdxs = nil
}
