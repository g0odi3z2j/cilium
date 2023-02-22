// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.7
// source: envoy/config/cluster/v3/circuit_breaker.proto

package clusterv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v31 "github.com/cilium/proxy/go/envoy/type/v3"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If multiple :ref:`Thresholds<envoy_v3_api_msg_config.cluster.v3.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`, the default values
	// are used.
	Thresholds []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	// Optional per-host limits which apply to each individual host in a cluster.
	//
	// .. note::
	//
	//	currently only the :ref:`max_connections
	//	<envoy_v3_api_field_config.cluster.v3.CircuitBreakers.Thresholds.max_connections>` field is supported for per-host limits.
	//
	// If multiple per-host :ref:`Thresholds<envoy_v3_api_msg_config.cluster.v3.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`,
	// the first one in the list is used. If no per-host Thresholds are defined for a given
	// :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`,
	// the cluster will not have per-host limits.
	PerHostThresholds []*CircuitBreakers_Thresholds `protobuf:"bytes,2,rep,name=per_host_thresholds,json=perHostThresholds,proto3" json:"per_host_thresholds,omitempty"`
}

func (x *CircuitBreakers) Reset() {
	*x = CircuitBreakers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircuitBreakers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircuitBreakers) ProtoMessage() {}

func (x *CircuitBreakers) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircuitBreakers.ProtoReflect.Descriptor instead.
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescGZIP(), []int{0}
}

func (x *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if x != nil {
		return x.Thresholds
	}
	return nil
}

func (x *CircuitBreakers) GetPerHostThresholds() []*CircuitBreakers_Thresholds {
	if x != nil {
		return x.PerHostThresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`.
// [#next-free-field: 9]
type CircuitBreakers_Thresholds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The :ref:`RoutingPriority<envoy_v3_api_enum_config.core.v3.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	Priority v3.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v3.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	// This limit is applied as a connection limit for non-HTTP traffic.
	MaxPendingRequests *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	// This limit does not apply to non-HTTP traffic.
	MaxRequests *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Specifies a limit on concurrent retries in relation to the number of active requests. This
	// parameter is optional.
	//
	// .. note::
	//
	//	If this field is set, the retry budget will override any configured retry circuit
	//	breaker.
	RetryBudget *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	//
	// .. note::
	//
	//	If a retry budget is used in lieu of the max_retries circuit breaker,
	//	the remaining retry resources remaining will not be tracked.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
}

func (x *CircuitBreakers_Thresholds) Reset() {
	*x = CircuitBreakers_Thresholds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircuitBreakers_Thresholds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircuitBreakers_Thresholds) ProtoMessage() {}

func (x *CircuitBreakers_Thresholds) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircuitBreakers_Thresholds.ProtoReflect.Descriptor instead.
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CircuitBreakers_Thresholds) GetPriority() v3.RoutingPriority {
	if x != nil {
		return x.Priority
	}
	return v3.RoutingPriority_DEFAULT
}

func (x *CircuitBreakers_Thresholds) GetMaxConnections() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxConnections
	}
	return nil
}

func (x *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxPendingRequests
	}
	return nil
}

func (x *CircuitBreakers_Thresholds) GetMaxRequests() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxRequests
	}
	return nil
}

func (x *CircuitBreakers_Thresholds) GetMaxRetries() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxRetries
	}
	return nil
}

func (x *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if x != nil {
		return x.RetryBudget
	}
	return nil
}

func (x *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if x != nil {
		return x.TrackRemaining
	}
	return false
}

func (x *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the limit on concurrent retries as a percentage of the sum of active requests and
	// active pending requests. For example, if there are 100 active requests and the
	// budget_percent is set to 25, there may be 25 active retries.
	//
	// This parameter is optional. Defaults to 20%.
	BudgetPercent *v31.Percent `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	// Specifies the minimum retry concurrency allowed for the retry budget. The limit on the
	// number of active retries may never go below this number.
	//
	// This parameter is optional. Defaults to 3.
	MinRetryConcurrency *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
}

func (x *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*x = CircuitBreakers_Thresholds_RetryBudget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircuitBreakers_Thresholds_RetryBudget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage() {}

func (x *CircuitBreakers_Thresholds_RetryBudget) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircuitBreakers_Thresholds_RetryBudget.ProtoReflect.Descriptor instead.
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v31.Percent {
	if x != nil {
		return x.BudgetPercent
	}
	return nil
}

func (x *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MinRetryConcurrency
	}
	return nil
}

var File_envoy_config_cluster_v3_circuit_breaker_proto protoreflect.FileDescriptor

var file_envoy_config_cluster_v3_circuit_breaker_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x08, 0x0a, 0x0f, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x0a,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x63, 0x0a, 0x13, 0x70, 0x65,
	0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x11, 0x70, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x1a,
	0xea, 0x06, 0x0a, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x4b,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0f, 0x6d,
	0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12,
	0x6d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x62, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x4e, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x1a,
	0xe2, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x3d, 0x0a, 0x0e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52,
	0x0d, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x50,
	0x0a, 0x15, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x3a, 0x42, 0x9a, 0xc5, 0x88, 0x1e, 0x3d, 0x0a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x69,
	0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x3a, 0x36, 0x9a, 0xc5, 0x88, 0x1e, 0x31, 0x0a, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x3a, 0x2b, 0x9a, 0xc5,
	0x88, 0x1e, 0x26, 0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x42, 0x90, 0x01, 0x0a, 0x25, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x42, 0x13, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescOnce sync.Once
	file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescData = file_envoy_config_cluster_v3_circuit_breaker_proto_rawDesc
)

func file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescGZIP() []byte {
	file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescOnce.Do(func() {
		file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescData)
	})
	return file_envoy_config_cluster_v3_circuit_breaker_proto_rawDescData
}

var file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_config_cluster_v3_circuit_breaker_proto_goTypes = []interface{}{
	(*CircuitBreakers)(nil),                        // 0: envoy.config.cluster.v3.CircuitBreakers
	(*CircuitBreakers_Thresholds)(nil),             // 1: envoy.config.cluster.v3.CircuitBreakers.Thresholds
	(*CircuitBreakers_Thresholds_RetryBudget)(nil), // 2: envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget
	(v3.RoutingPriority)(0),                        // 3: envoy.config.core.v3.RoutingPriority
	(*wrapperspb.UInt32Value)(nil),                 // 4: google.protobuf.UInt32Value
	(*v31.Percent)(nil),                            // 5: envoy.type.v3.Percent
}
var file_envoy_config_cluster_v3_circuit_breaker_proto_depIdxs = []int32{
	1,  // 0: envoy.config.cluster.v3.CircuitBreakers.thresholds:type_name -> envoy.config.cluster.v3.CircuitBreakers.Thresholds
	1,  // 1: envoy.config.cluster.v3.CircuitBreakers.per_host_thresholds:type_name -> envoy.config.cluster.v3.CircuitBreakers.Thresholds
	3,  // 2: envoy.config.cluster.v3.CircuitBreakers.Thresholds.priority:type_name -> envoy.config.core.v3.RoutingPriority
	4,  // 3: envoy.config.cluster.v3.CircuitBreakers.Thresholds.max_connections:type_name -> google.protobuf.UInt32Value
	4,  // 4: envoy.config.cluster.v3.CircuitBreakers.Thresholds.max_pending_requests:type_name -> google.protobuf.UInt32Value
	4,  // 5: envoy.config.cluster.v3.CircuitBreakers.Thresholds.max_requests:type_name -> google.protobuf.UInt32Value
	4,  // 6: envoy.config.cluster.v3.CircuitBreakers.Thresholds.max_retries:type_name -> google.protobuf.UInt32Value
	2,  // 7: envoy.config.cluster.v3.CircuitBreakers.Thresholds.retry_budget:type_name -> envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget
	4,  // 8: envoy.config.cluster.v3.CircuitBreakers.Thresholds.max_connection_pools:type_name -> google.protobuf.UInt32Value
	5,  // 9: envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget.budget_percent:type_name -> envoy.type.v3.Percent
	4,  // 10: envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget.min_retry_concurrency:type_name -> google.protobuf.UInt32Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_envoy_config_cluster_v3_circuit_breaker_proto_init() }
func file_envoy_config_cluster_v3_circuit_breaker_proto_init() {
	if File_envoy_config_cluster_v3_circuit_breaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircuitBreakers); i {
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
		file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircuitBreakers_Thresholds); i {
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
		file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircuitBreakers_Thresholds_RetryBudget); i {
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
			RawDescriptor: file_envoy_config_cluster_v3_circuit_breaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_cluster_v3_circuit_breaker_proto_goTypes,
		DependencyIndexes: file_envoy_config_cluster_v3_circuit_breaker_proto_depIdxs,
		MessageInfos:      file_envoy_config_cluster_v3_circuit_breaker_proto_msgTypes,
	}.Build()
	File_envoy_config_cluster_v3_circuit_breaker_proto = out.File
	file_envoy_config_cluster_v3_circuit_breaker_proto_rawDesc = nil
	file_envoy_config_cluster_v3_circuit_breaker_proto_goTypes = nil
	file_envoy_config_cluster_v3_circuit_breaker_proto_depIdxs = nil
}
