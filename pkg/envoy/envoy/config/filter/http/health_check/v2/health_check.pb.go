// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

package v2

import (
	fmt "fmt"
	route "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/route"
	_type "github.com/cilium/cilium/pkg/envoy/envoy/type"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *wrappers.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *duration.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*_type.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_75439d7b4d98e201, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*_type.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
	proto.RegisterMapType((map[string]*_type.Percent)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptor_75439d7b4d98e201)
}

var fileDescriptor_75439d7b4d98e201 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8a, 0xd4, 0x30,
	0x18, 0xa6, 0xed, 0x8c, 0xee, 0x66, 0x0e, 0x8e, 0x55, 0xb0, 0x0e, 0xe2, 0xee, 0x7a, 0x1a, 0x2f,
	0x09, 0xd4, 0x8b, 0xb8, 0xe0, 0x61, 0x56, 0x61, 0x11, 0x06, 0xa4, 0x2c, 0x0a, 0x5e, 0x4a, 0x36,
	0xfd, 0xa7, 0x09, 0xdb, 0x49, 0x42, 0x9a, 0x56, 0xfa, 0x0a, 0x3e, 0x81, 0x27, 0x1f, 0xc4, 0x93,
	0x6f, 0xe2, 0xd9, 0x07, 0xf0, 0x2e, 0x49, 0x5a, 0x74, 0x59, 0x50, 0x2f, 0xe5, 0x6f, 0xbe, 0xff,
	0xfb, 0xfe, 0x2f, 0xdf, 0x1f, 0x74, 0x0a, 0xb2, 0x57, 0x03, 0x61, 0x4a, 0xee, 0x44, 0x4d, 0x76,
	0xa2, 0xb1, 0x60, 0x08, 0xb7, 0x56, 0x13, 0x0e, 0xb4, 0xb1, 0xbc, 0x64, 0x1c, 0xd8, 0x15, 0xe9,
	0xf3, 0x6b, 0xff, 0x58, 0x1b, 0x65, 0x55, 0xba, 0xf6, 0x64, 0x1c, 0xc8, 0x38, 0x90, 0xb1, 0x23,
	0xe3, 0x6b, 0xcd, 0x7d, 0xbe, 0x7a, 0x5c, 0x2b, 0x55, 0x37, 0x40, 0x3c, 0xef, 0xb2, 0xdb, 0x91,
	0xaa, 0x33, 0xd4, 0x0a, 0x25, 0x83, 0xd2, 0x4d, 0xfc, 0xa3, 0xa1, 0x5a, 0x83, 0x69, 0x27, 0x3c,
	0xd8, 0xa4, 0x5a, 0x38, 0x2b, 0x46, 0x75, 0x16, 0xc2, 0x77, 0xc4, 0xb3, 0x80, 0xdb, 0x41, 0x03,
	0xd1, 0x60, 0x18, 0x48, 0x3b, 0x22, 0x0f, 0x7a, 0xda, 0x88, 0x8a, 0x5a, 0x20, 0x53, 0x31, 0x02,
	0xf7, 0x6b, 0x55, 0x2b, 0x5f, 0x12, 0x57, 0x85, 0xd3, 0x27, 0x3f, 0x13, 0xb4, 0x38, 0xf7, 0xe6,
	0xcf, 0x9c, 0xf7, 0xb4, 0x40, 0x77, 0x35, 0x6d, 0xdb, 0xd2, 0x72, 0xa3, 0xba, 0x9a, 0x97, 0x7b,
	0x55, 0x41, 0x16, 0x1d, 0x47, 0xeb, 0x45, 0xbe, 0xc2, 0xc1, 0x34, 0x9e, 0x4c, 0xe3, 0x8d, 0x52,
	0xcd, 0x3b, 0xda, 0x74, 0xb0, 0x41, 0x5f, 0x7f, 0x7c, 0x4b, 0xe6, 0x9f, 0xa2, 0x78, 0x19, 0x15,
	0x77, 0x9c, 0xc0, 0x45, 0xe0, 0x6f, 0x55, 0x05, 0xe9, 0x4b, 0x84, 0x18, 0x65, 0x1c, 0x4a, 0x2b,
	0xf6, 0x90, 0x25, 0x5e, 0xec, 0xe1, 0x0d, 0xb1, 0x57, 0x63, 0x42, 0x9b, 0xd9, 0xe7, 0xef, 0x47,
	0x51, 0x71, 0xe8, 0x29, 0x17, 0x62, 0x0f, 0xe9, 0x97, 0x08, 0x1d, 0xb1, 0xa6, 0x6b, 0x2d, 0x98,
	0x72, 0x2f, 0x64, 0x19, 0xc2, 0x1e, 0xca, 0xf1, 0xe2, 0xb4, 0x86, 0x36, 0x9b, 0x1d, 0x27, 0xeb,
	0x45, 0xfe, 0x1e, 0xff, 0xef, 0x86, 0xf0, 0x1f, 0x97, 0xc6, 0x67, 0x41, 0x7c, 0x2b, 0x64, 0x38,
	0x1d, 0xde, 0xfe, 0x56, 0x7e, 0x2d, 0xad, 0x19, 0x8a, 0x47, 0xec, 0x2f, 0x2d, 0xe9, 0x29, 0xba,
	0xcd, 0x81, 0x56, 0x60, 0xda, 0x6c, 0xee, 0x7d, 0x9c, 0x8c, 0x3e, 0xa8, 0x16, 0x6e, 0x56, 0xd8,
	0xdc, 0xb9, 0x6f, 0xd9, 0x52, 0xcb, 0x38, 0x98, 0x62, 0x62, 0xac, 0x2a, 0x74, 0xf2, 0xcf, 0xf9,
	0xe9, 0x12, 0x25, 0x57, 0x30, 0xf8, 0x45, 0x1c, 0x16, 0xae, 0x4c, 0x9f, 0xa2, 0x79, 0xef, 0xa2,
	0xcf, 0x62, 0x9f, 0xe7, 0xbd, 0x71, 0xa2, 0x7b, 0x11, 0x78, 0xa4, 0x17, 0xa1, 0xe3, 0x45, 0xfc,
	0x3c, 0x7a, 0x33, 0x3b, 0x88, 0x97, 0x49, 0x71, 0x00, 0xb2, 0xd2, 0x4a, 0x48, 0xbb, 0x99, 0x7d,
	0x88, 0xfb, 0xfc, 0xf2, 0x96, 0x4f, 0xff, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x2f,
	0x6d, 0x53, 0x16, 0x03, 0x00, 0x00,
}
