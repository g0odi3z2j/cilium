// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/cluster/v3/outlier_detection.proto

package envoy_config_cluster_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
// [#next-free-field: 22]
type OutlierDetection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of consecutive 5xx responses or local origin errors that are mapped
	// to 5xx error codes before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *durationpb.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected and is
	// capped by :ref:`max_ejection_time<envoy_api_field_config.cluster.v3.OutlierDetection.max_ejection_time>`.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *durationpb.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *wrapperspb.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes)
	// before a consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *wrapperspb.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	// Determines whether to distinguish local origin failures from external errors. If set to true
	// the following configuration parameters are taken into account:
	// :ref:`consecutive_local_origin_failure<envoy_api_field_config.cluster.v3.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_config.cluster.v3.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_config.cluster.v3.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *wrapperspb.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *wrapperspb.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingLocalOriginSuccessRate *wrapperspb.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	// The failure percentage to use when determining failure percentage-based outlier detection. If
	// the failure percentage of a given host is greater than or equal to this value, it will be
	// ejected. Defaults to 85.
	FailurePercentageThreshold *wrapperspb.UInt32Value `protobuf:"bytes,16,opt,name=failure_percentage_threshold,json=failurePercentageThreshold,proto3" json:"failure_percentage_threshold,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// failure percentage statistics. This setting can be used to disable ejection or to ramp it up
	// slowly. Defaults to 0.
	//
	// [#next-major-version: setting this without setting failure_percentage_threshold should be
	// invalid in v4.]
	EnforcingFailurePercentage *wrapperspb.UInt32Value `protobuf:"bytes,17,opt,name=enforcing_failure_percentage,json=enforcingFailurePercentage,proto3" json:"enforcing_failure_percentage,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// local-origin failure percentage statistics. This setting can be used to disable ejection or to
	// ramp it up slowly. Defaults to 0.
	EnforcingFailurePercentageLocalOrigin *wrapperspb.UInt32Value `protobuf:"bytes,18,opt,name=enforcing_failure_percentage_local_origin,json=enforcingFailurePercentageLocalOrigin,proto3" json:"enforcing_failure_percentage_local_origin,omitempty"`
	// The minimum number of hosts in a cluster in order to perform failure percentage-based ejection.
	// If the total number of hosts in the cluster is less than this value, failure percentage-based
	// ejection will not be performed. Defaults to 5.
	FailurePercentageMinimumHosts *wrapperspb.UInt32Value `protobuf:"bytes,19,opt,name=failure_percentage_minimum_hosts,json=failurePercentageMinimumHosts,proto3" json:"failure_percentage_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one interval (as defined by the
	// interval duration above) to perform failure percentage-based ejection for this host. If the
	// volume is lower than this setting, failure percentage-based ejection will not be performed for
	// this host. Defaults to 50.
	FailurePercentageRequestVolume *wrapperspb.UInt32Value `protobuf:"bytes,20,opt,name=failure_percentage_request_volume,json=failurePercentageRequestVolume,proto3" json:"failure_percentage_request_volume,omitempty"`
	// The maximum time that a host is ejected for. See :ref:`base_ejection_time<envoy_api_field_config.cluster.v3.OutlierDetection.base_ejection_time>`
	// for more information.
	// Defaults to 300000ms or 300s.
	MaxEjectionTime *durationpb.Duration `protobuf:"bytes,21,opt,name=max_ejection_time,json=maxEjectionTime,proto3" json:"max_ejection_time,omitempty"`
}

func (x *OutlierDetection) Reset() {
	*x = OutlierDetection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_v3_outlier_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetection) ProtoMessage() {}

func (x *OutlierDetection) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_v3_outlier_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetection.ProtoReflect.Descriptor instead.
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_v3_outlier_detection_proto_rawDescGZIP(), []int{0}
}

func (x *OutlierDetection) GetConsecutive_5Xx() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Consecutive_5Xx
	}
	return nil
}

func (x *OutlierDetection) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *OutlierDetection) GetBaseEjectionTime() *durationpb.Duration {
	if x != nil {
		return x.BaseEjectionTime
	}
	return nil
}

func (x *OutlierDetection) GetMaxEjectionPercent() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxEjectionPercent
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutive_5Xx
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingSuccessRate() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingSuccessRate
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateMinimumHosts() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SuccessRateMinimumHosts
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateRequestVolume() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SuccessRateRequestVolume
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateStdevFactor() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SuccessRateStdevFactor
	}
	return nil
}

func (x *OutlierDetection) GetConsecutiveGatewayFailure() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ConsecutiveGatewayFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (x *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if x != nil {
		return x.SplitExternalLocalOriginErrors
	}
	return false
}

func (x *OutlierDetection) GetConsecutiveLocalOriginFailure() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageThreshold() *wrapperspb.UInt32Value {
	if x != nil {
		return x.FailurePercentageThreshold
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingFailurePercentage() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingFailurePercentage
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingFailurePercentageLocalOrigin() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EnforcingFailurePercentageLocalOrigin
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageMinimumHosts() *wrapperspb.UInt32Value {
	if x != nil {
		return x.FailurePercentageMinimumHosts
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageRequestVolume() *wrapperspb.UInt32Value {
	if x != nil {
		return x.FailurePercentageRequestVolume
	}
	return nil
}

func (x *OutlierDetection) GetMaxEjectionTime() *durationpb.Duration {
	if x != nil {
		return x.MaxEjectionTime
	}
	return nil
}

var File_envoy_config_cluster_v3_outlier_detection_proto protoreflect.FileDescriptor

var file_envoy_config_cluster_v3_outlier_detection_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x10, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x35, 0x78, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x35, 0x78,
	0x78, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x51, 0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x2a, 0x00, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x45,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x61,
	0x0a, 0x19, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x35, 0x78, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x17, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x35, 0x78,
	0x78, 0x12, 0x5b, 0x0a, 0x16, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x14, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x69, 0x6e, 0x67, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x59,
	0x0a, 0x1a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x17, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x1b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x19, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x64, 0x65, 0x76, 0x5f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x61, 0x74, 0x65, 0x53, 0x74, 0x64, 0x65, 0x76, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x5c, 0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x78, 0x0a,
	0x25, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x18, 0x64, 0x52, 0x22, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x4a, 0x0a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x65, 0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x2a, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x26, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x73,
	0x0a, 0x23, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x18, 0x64, 0x52, 0x1f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x67, 0x0a, 0x1c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64,
	0x52, 0x1a, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x67, 0x0a, 0x1c,
	0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x1a, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x7f, 0x0a, 0x29, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52,
	0x25, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x67, 0x0a,
	0x21, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x45, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x15,
	0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_cluster_v3_outlier_detection_proto_rawDescOnce sync.Once
	file_envoy_config_cluster_v3_outlier_detection_proto_rawDescData = file_envoy_config_cluster_v3_outlier_detection_proto_rawDesc
)

func file_envoy_config_cluster_v3_outlier_detection_proto_rawDescGZIP() []byte {
	file_envoy_config_cluster_v3_outlier_detection_proto_rawDescOnce.Do(func() {
		file_envoy_config_cluster_v3_outlier_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_cluster_v3_outlier_detection_proto_rawDescData)
	})
	return file_envoy_config_cluster_v3_outlier_detection_proto_rawDescData
}

var file_envoy_config_cluster_v3_outlier_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_cluster_v3_outlier_detection_proto_goTypes = []interface{}{
	(*OutlierDetection)(nil),       // 0: envoy.config.cluster.v3.OutlierDetection
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_envoy_config_cluster_v3_outlier_detection_proto_depIdxs = []int32{
	1,  // 0: envoy.config.cluster.v3.OutlierDetection.consecutive_5xx:type_name -> google.protobuf.UInt32Value
	2,  // 1: envoy.config.cluster.v3.OutlierDetection.interval:type_name -> google.protobuf.Duration
	2,  // 2: envoy.config.cluster.v3.OutlierDetection.base_ejection_time:type_name -> google.protobuf.Duration
	1,  // 3: envoy.config.cluster.v3.OutlierDetection.max_ejection_percent:type_name -> google.protobuf.UInt32Value
	1,  // 4: envoy.config.cluster.v3.OutlierDetection.enforcing_consecutive_5xx:type_name -> google.protobuf.UInt32Value
	1,  // 5: envoy.config.cluster.v3.OutlierDetection.enforcing_success_rate:type_name -> google.protobuf.UInt32Value
	1,  // 6: envoy.config.cluster.v3.OutlierDetection.success_rate_minimum_hosts:type_name -> google.protobuf.UInt32Value
	1,  // 7: envoy.config.cluster.v3.OutlierDetection.success_rate_request_volume:type_name -> google.protobuf.UInt32Value
	1,  // 8: envoy.config.cluster.v3.OutlierDetection.success_rate_stdev_factor:type_name -> google.protobuf.UInt32Value
	1,  // 9: envoy.config.cluster.v3.OutlierDetection.consecutive_gateway_failure:type_name -> google.protobuf.UInt32Value
	1,  // 10: envoy.config.cluster.v3.OutlierDetection.enforcing_consecutive_gateway_failure:type_name -> google.protobuf.UInt32Value
	1,  // 11: envoy.config.cluster.v3.OutlierDetection.consecutive_local_origin_failure:type_name -> google.protobuf.UInt32Value
	1,  // 12: envoy.config.cluster.v3.OutlierDetection.enforcing_consecutive_local_origin_failure:type_name -> google.protobuf.UInt32Value
	1,  // 13: envoy.config.cluster.v3.OutlierDetection.enforcing_local_origin_success_rate:type_name -> google.protobuf.UInt32Value
	1,  // 14: envoy.config.cluster.v3.OutlierDetection.failure_percentage_threshold:type_name -> google.protobuf.UInt32Value
	1,  // 15: envoy.config.cluster.v3.OutlierDetection.enforcing_failure_percentage:type_name -> google.protobuf.UInt32Value
	1,  // 16: envoy.config.cluster.v3.OutlierDetection.enforcing_failure_percentage_local_origin:type_name -> google.protobuf.UInt32Value
	1,  // 17: envoy.config.cluster.v3.OutlierDetection.failure_percentage_minimum_hosts:type_name -> google.protobuf.UInt32Value
	1,  // 18: envoy.config.cluster.v3.OutlierDetection.failure_percentage_request_volume:type_name -> google.protobuf.UInt32Value
	2,  // 19: envoy.config.cluster.v3.OutlierDetection.max_ejection_time:type_name -> google.protobuf.Duration
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_envoy_config_cluster_v3_outlier_detection_proto_init() }
func file_envoy_config_cluster_v3_outlier_detection_proto_init() {
	if File_envoy_config_cluster_v3_outlier_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_cluster_v3_outlier_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetection); i {
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
			RawDescriptor: file_envoy_config_cluster_v3_outlier_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_cluster_v3_outlier_detection_proto_goTypes,
		DependencyIndexes: file_envoy_config_cluster_v3_outlier_detection_proto_depIdxs,
		MessageInfos:      file_envoy_config_cluster_v3_outlier_detection_proto_msgTypes,
	}.Build()
	File_envoy_config_cluster_v3_outlier_detection_proto = out.File
	file_envoy_config_cluster_v3_outlier_detection_proto_rawDesc = nil
	file_envoy_config_cluster_v3_outlier_detection_proto_goTypes = nil
	file_envoy_config_cluster_v3_outlier_detection_proto_depIdxs = nil
}
