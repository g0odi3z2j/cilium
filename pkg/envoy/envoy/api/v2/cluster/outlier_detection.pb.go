// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package cluster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
type OutlierDetection struct {
	// The number of consecutive 5xx responses before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *duration.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status or
	// connection errors that are mapped to one of those status codes) before a
	// consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}              `json:"-"`
	XXX_unrecognized                   []byte                `json:"-"`
	XXX_sizecache                      int32                 `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_outlier_detection_35470c6804ebd24d, []int{0}
}
func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetection.Unmarshal(m, b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
}
func (dst *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(dst, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_OutlierDetection.Size(m)
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *duration.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor_outlier_detection_35470c6804ebd24d)
}

var fileDescriptor_outlier_detection_35470c6804ebd24d = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x4d, 0x7a, 0x9e, 0x82, 0x96, 0x21, 0x24, 0x93, 0x56, 0x8b, 0x14, 0x04, 0x29, 0xb2,
	0x0b, 0x29, 0x7d, 0x80, 0xa6, 0x4d, 0xd5, 0x0b, 0x0f, 0x24, 0x1a, 0x11, 0x95, 0x61, 0x32, 0xf9,
	0xb2, 0x8e, 0xec, 0xee, 0xac, 0x73, 0xd8, 0x6e, 0x7d, 0x22, 0xe9, 0x23, 0x78, 0xe5, 0xb3, 0x78,
	0xe7, 0x33, 0x78, 0x23, 0xbb, 0x93, 0xc3, 0xa6, 0x2d, 0x34, 0xb9, 0x1b, 0x76, 0xe6, 0xf7, 0xfb,
	0x7f, 0xfb, 0xcd, 0xc7, 0xa0, 0x67, 0x10, 0xa7, 0xf2, 0xd2, 0x67, 0x89, 0xf0, 0xd3, 0x96, 0xcf,
	0x43, 0xab, 0x0d, 0x28, 0x5f, 0x5a, 0x13, 0x0a, 0x50, 0x74, 0x08, 0x06, 0xb8, 0x11, 0x32, 0xf6,
	0x12, 0x25, 0x8d, 0xc4, 0xb5, 0xe2, 0xb4, 0xc7, 0x12, 0xe1, 0xa5, 0x2d, 0x6f, 0x7c, 0x7a, 0x77,
	0x3f, 0x90, 0x32, 0x08, 0xc1, 0x2f, 0xce, 0x0c, 0xec, 0xc8, 0x1f, 0x5a, 0xc5, 0x66, 0xd4, 0xcd,
	0xfd, 0x0b, 0xc5, 0x92, 0x04, 0x94, 0x1e, 0xef, 0x37, 0x52, 0x16, 0x8a, 0x21, 0x33, 0xe0, 0x4f,
	0x16, 0xe3, 0x8d, 0x5a, 0x20, 0x03, 0x59, 0x2c, 0xfd, 0x7c, 0xe5, 0xbe, 0x1e, 0xfc, 0xdb, 0x40,
	0x3b, 0x6f, 0x5c, 0x81, 0x67, 0x93, 0xfa, 0x70, 0x07, 0x3d, 0xe0, 0x32, 0xd6, 0xc0, 0xad, 0x11,
	0x29, 0xd0, 0xe3, 0x2c, 0x23, 0x95, 0xc7, 0x95, 0xa7, 0xdb, 0xad, 0x87, 0x9e, 0x4b, 0xf7, 0x26,
	0xe9, 0xde, 0xfb, 0x97, 0xb1, 0x39, 0x6a, 0xf5, 0x59, 0x68, 0xa1, 0x7b, 0xbf, 0x04, 0x1d, 0x67,
	0x19, 0x3e, 0x41, 0x9b, 0x22, 0x36, 0xa0, 0x52, 0x16, 0x92, 0x6a, 0xc1, 0x37, 0x6f, 0xf0, 0x67,
	0xe3, 0xbf, 0x6b, 0xa3, 0x5f, 0x7f, 0x7f, 0xaf, 0xac, 0x5d, 0x55, 0xaa, 0x87, 0xf7, 0xba, 0x53,
	0x0c, 0xf7, 0x10, 0x1e, 0x30, 0x0d, 0x14, 0xbe, 0xb9, 0xd2, 0xa8, 0x11, 0x11, 0x90, 0x95, 0x65,
	0x64, 0x3b, 0xb9, 0xa0, 0x33, 0xe6, 0xdf, 0x89, 0x08, 0xf0, 0x47, 0x54, 0x8b, 0x58, 0x36, 0x73,
	0x26, 0xa0, 0x38, 0xc4, 0x86, 0xac, 0xde, 0xfd, 0x8f, 0xed, 0xad, 0xdc, 0xbc, 0x7a, 0x58, 0x25,
	0xc3, 0x2e, 0x8e, 0x58, 0x36, 0xf1, 0xbe, 0x75, 0x0a, 0xcc, 0x51, 0x13, 0xe2, 0x91, 0x54, 0x5c,
	0xc4, 0x01, 0xbd, 0xde, 0xc3, 0xb5, 0xe5, 0xfc, 0x8d, 0xa9, 0xe9, 0x74, 0xbe, 0xaf, 0x5f, 0x50,
	0x7d, 0x16, 0xa2, 0x2d, 0xe7, 0xa0, 0x35, 0x55, 0xcc, 0x00, 0x59, 0x5f, 0x2e, 0xa1, 0x36, 0xd5,
	0xf4, 0x9c, 0xa5, 0xcb, 0x4c, 0xde, 0x9e, 0xdd, 0xb2, 0x94, 0x46, 0x22, 0x16, 0x91, 0x8d, 0xe8,
	0x57, 0xa9, 0x8d, 0x26, 0x1b, 0x0b, 0x0c, 0x42, 0x43, 0xcf, 0x74, 0xaf, 0x1c, 0xfd, 0x22, 0x87,
	0xf1, 0x27, 0xb4, 0x37, 0xa7, 0x56, 0xf0, 0xdd, 0x82, 0x36, 0x34, 0x95, 0xa1, 0x8d, 0x80, 0x6c,
	0x2e, 0xe0, 0x26, 0x25, 0x77, 0xd7, 0xe1, 0xfd, 0x82, 0xc6, 0x1f, 0x50, 0x73, 0x4e, 0xae, 0xcd,
	0x10, 0x52, 0x3a, 0x62, 0xdc, 0x48, 0x45, 0xb6, 0x16, 0x50, 0xd7, 0x4b, 0xea, 0x5e, 0x0e, 0x9f,
	0x17, 0x2c, 0xfe, 0x8c, 0xf6, 0xca, 0x57, 0x19, 0x30, 0x03, 0x17, 0xec, 0x92, 0x8e, 0x98, 0x08,
	0xad, 0x02, 0x82, 0x16, 0x50, 0x37, 0x4b, 0x82, 0xe7, 0x8e, 0x3f, 0x77, 0x38, 0xfe, 0x81, 0x9e,
	0xdc, 0x3e, 0x32, 0xd7, 0x73, 0xb6, 0x97, 0xbb, 0xdc, 0x83, 0xdb, 0xc6, 0x67, 0x3e, 0xbb, 0xfd,
	0xe8, 0xe7, 0x9f, 0xfd, 0xca, 0x55, 0xb5, 0xde, 0x29, 0x5e, 0xa2, 0x93, 0x44, 0x78, 0xfd, 0x96,
	0x77, 0xea, 0x5e, 0xa2, 0xd7, 0xbd, 0xc1, 0x7a, 0x91, 0x71, 0xf4, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0xdf, 0x24, 0x7d, 0xd8, 0x04, 0x00, 0x00,
}
