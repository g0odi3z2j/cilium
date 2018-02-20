// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package cluster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
type OutlierDetection struct {
	// The number of consecutive 5xx responses before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *google_protobuf3.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *google_protobuf3.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10%.
	MaxEjectionPercent *google_protobuf1.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *google_protobuf1.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *google_protobuf1.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *google_protobuf1.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *google_protobuf1.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *google_protobuf1.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status or
	// connection errors that are mapped to one of those status codes) before a
	// consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *google_protobuf1.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *google_protobuf1.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure" json:"enforcing_consecutive_gateway_failure,omitempty"`
}

func (m *OutlierDetection) Reset()                    { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string            { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()               {}
func (*OutlierDetection) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *OutlierDetection) GetConsecutive_5Xx() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *google_protobuf3.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *google_protobuf3.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() { proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xdd, 0x76, 0xfb, 0x35, 0x05, 0x2d, 0xc3, 0xd2, 0x9d, 0x6d, 0xa5, 0x48, 0x41, 0x90,
	0x22, 0x09, 0x6c, 0xe9, 0x03, 0xb8, 0xb6, 0x55, 0x2f, 0x44, 0xc9, 0x6a, 0x45, 0x54, 0x86, 0xd9,
	0xc9, 0xd9, 0x38, 0x92, 0x64, 0xe2, 0x7c, 0xa4, 0xa9, 0x4f, 0x24, 0x3e, 0x82, 0x57, 0x3e, 0x8b,
	0x77, 0x3e, 0x83, 0x37, 0x92, 0x4c, 0x76, 0x37, 0xdb, 0x16, 0xdc, 0xbd, 0x1b, 0x32, 0xe7, 0xf7,
	0xfb, 0x9f, 0x9c, 0x39, 0xe8, 0x31, 0xa4, 0xb9, 0xbc, 0xf2, 0x59, 0x26, 0xfc, 0xbc, 0xef, 0xf3,
	0xd8, 0x6a, 0x03, 0xca, 0x97, 0xd6, 0xc4, 0x02, 0x14, 0x0d, 0xc1, 0x00, 0x37, 0x42, 0xa6, 0x5e,
	0xa6, 0xa4, 0x91, 0xb8, 0x53, 0x55, 0x7b, 0x2c, 0x13, 0x5e, 0xde, 0xf7, 0xea, 0xea, 0xbd, 0x83,
	0x48, 0xca, 0x28, 0x06, 0xbf, 0xaa, 0x19, 0xd9, 0xb1, 0x1f, 0x5a, 0xc5, 0x66, 0xd4, 0xcd, 0xfb,
	0x4b, 0xc5, 0xb2, 0x0c, 0x94, 0xae, 0xef, 0xbb, 0x39, 0x8b, 0x45, 0xc8, 0x0c, 0xf8, 0x93, 0x43,
	0x7d, 0xd1, 0x89, 0x64, 0x24, 0xab, 0xa3, 0x5f, 0x9e, 0xdc, 0xd7, 0xc3, 0xbf, 0x1b, 0x68, 0xe7,
	0x95, 0x6b, 0xf0, 0x74, 0xd2, 0x1f, 0x3e, 0x43, 0xf7, 0xb8, 0x4c, 0x35, 0x70, 0x6b, 0x44, 0x0e,
	0xf4, 0xa4, 0x28, 0x48, 0xeb, 0x41, 0xeb, 0xd1, 0x76, 0xff, 0xbe, 0xe7, 0xd2, 0xbd, 0x49, 0xba,
	0xf7, 0xf6, 0x45, 0x6a, 0x8e, 0xfb, 0x17, 0x2c, 0xb6, 0x10, 0xdc, 0x6d, 0x40, 0x27, 0x45, 0x81,
	0x9f, 0xa0, 0x4d, 0x91, 0x1a, 0x50, 0x39, 0x8b, 0xc9, 0x4a, 0xc5, 0xf7, 0x6e, 0xf0, 0xa7, 0xf5,
	0xdf, 0x0d, 0xd0, 0xcf, 0x3f, 0xbf, 0x56, 0xd7, 0x7e, 0xb4, 0x56, 0x8e, 0xee, 0x04, 0x53, 0x0c,
	0x0f, 0x11, 0x1e, 0x31, 0x0d, 0x14, 0xbe, 0xb8, 0xd6, 0xa8, 0x11, 0x09, 0x90, 0xd5, 0x65, 0x64,
	0x3b, 0xa5, 0xe0, 0xac, 0xe6, 0xdf, 0x88, 0x04, 0xf0, 0x7b, 0xd4, 0x49, 0x58, 0x31, 0x73, 0x66,
	0xa0, 0x38, 0xa4, 0x86, 0xb4, 0xff, 0xff, 0x8f, 0x83, 0xad, 0xd2, 0xdc, 0x3e, 0x5a, 0x21, 0x61,
	0x80, 0x13, 0x56, 0x4c, 0xbc, 0xaf, 0x9d, 0x02, 0x73, 0xd4, 0x83, 0x74, 0x2c, 0x15, 0x17, 0x69,
	0x44, 0xaf, 0xcf, 0x70, 0x6d, 0x39, 0x7f, 0x77, 0x6a, 0x7a, 0x3a, 0x3f, 0xd7, 0x4f, 0x68, 0x77,
	0x16, 0xa2, 0x2d, 0xe7, 0xa0, 0x35, 0x55, 0xcc, 0x00, 0x59, 0x5f, 0x2e, 0xa1, 0x33, 0xd5, 0x0c,
	0x9d, 0x25, 0x60, 0xa6, 0x1c, 0xcf, 0x5e, 0x53, 0x4a, 0x13, 0x91, 0x8a, 0xc4, 0x26, 0xf4, 0xb3,
	0xd4, 0x46, 0x93, 0x8d, 0x05, 0x16, 0xa1, 0xab, 0x67, 0xba, 0x97, 0x8e, 0x7e, 0x5e, 0xc2, 0xf8,
	0x03, 0xda, 0x9f, 0x53, 0x2b, 0xf8, 0x6a, 0x41, 0x1b, 0x9a, 0xcb, 0xd8, 0x26, 0x40, 0x36, 0x17,
	0x70, 0x93, 0x86, 0x3b, 0x70, 0xf8, 0x45, 0x45, 0xe3, 0x77, 0xa8, 0x37, 0x27, 0xd7, 0x26, 0x84,
	0x9c, 0x8e, 0x19, 0x37, 0x52, 0x91, 0xad, 0x05, 0xd4, 0xbb, 0x0d, 0xf5, 0xb0, 0x84, 0xcf, 0x2b,
	0x16, 0x7f, 0x44, 0xfb, 0xcd, 0xa7, 0x8c, 0x98, 0x81, 0x4b, 0x76, 0x45, 0xc7, 0x4c, 0xc4, 0x56,
	0x01, 0x41, 0x0b, 0xa8, 0x7b, 0x0d, 0xc1, 0x33, 0xc7, 0x9f, 0x3b, 0x1c, 0x7f, 0x43, 0x0f, 0x6f,
	0x5f, 0x99, 0xeb, 0x39, 0xdb, 0xcb, 0x3d, 0xee, 0xe1, 0x6d, 0xeb, 0x33, 0x9f, 0x3d, 0x68, 0x7f,
	0xff, 0x7d, 0xd0, 0x1a, 0xad, 0x57, 0xaa, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x27,
	0x3d, 0xba, 0xbf, 0x04, 0x00, 0x00,
}
