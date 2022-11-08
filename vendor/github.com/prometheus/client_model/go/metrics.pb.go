// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/prometheus/client/metrics.proto

package io_prometheus_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MetricType int32

const (
	// COUNTER must use the Metric field "counter".
	MetricType_COUNTER MetricType = 0
	// GAUGE must use the Metric field "gauge".
	MetricType_GAUGE MetricType = 1
	// SUMMARY must use the Metric field "summary".
	MetricType_SUMMARY MetricType = 2
	// UNTYPED must use the Metric field "untyped".
	MetricType_UNTYPED MetricType = 3
	// HISTOGRAM must use the Metric field "histogram".
	MetricType_HISTOGRAM MetricType = 4
	// GAUGE_HISTOGRAM must use the Metric field "histogram".
	MetricType_GAUGE_HISTOGRAM MetricType = 5
)

var MetricType_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
	2: "SUMMARY",
	3: "UNTYPED",
	4: "HISTOGRAM",
	5: "GAUGE_HISTOGRAM",
}

var MetricType_value = map[string]int32{
	"COUNTER":         0,
	"GAUGE":           1,
	"SUMMARY":         2,
	"UNTYPED":         3,
	"HISTOGRAM":       4,
	"GAUGE_HISTOGRAM": 5,
}

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (x *MetricType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricType_value, data, "MetricType")
	if err != nil {
		return err
	}
	*x = MetricType(value)
	return nil
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{0}
}

type LabelPair struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}
func (*LabelPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{0}
}

func (m *LabelPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelPair.Unmarshal(m, b)
}
func (m *LabelPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelPair.Marshal(b, m, deterministic)
}
func (m *LabelPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelPair.Merge(m, src)
}
func (m *LabelPair) XXX_Size() int {
	return xxx_messageInfo_LabelPair.Size(m)
}
func (m *LabelPair) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelPair.DiscardUnknown(m)
}

var xxx_messageInfo_LabelPair proto.InternalMessageInfo

func (m *LabelPair) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Gauge struct {
	Value                *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{1}
}

func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gauge.Unmarshal(m, b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return xxx_messageInfo_Gauge.Size(m)
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Counter struct {
	Value                *float64  `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Exemplar             *Exemplar `protobuf:"bytes,2,opt,name=exemplar" json:"exemplar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{2}
}

func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Counter) GetExemplar() *Exemplar {
	if m != nil {
		return m.Exemplar
	}
	return nil
}

type Quantile struct {
	Quantile             *float64 `protobuf:"fixed64,1,opt,name=quantile" json:"quantile,omitempty"`
	Value                *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quantile) Reset()         { *m = Quantile{} }
func (m *Quantile) String() string { return proto.CompactTextString(m) }
func (*Quantile) ProtoMessage()    {}
func (*Quantile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{3}
}

func (m *Quantile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quantile.Unmarshal(m, b)
}
func (m *Quantile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quantile.Marshal(b, m, deterministic)
}
func (m *Quantile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quantile.Merge(m, src)
}
func (m *Quantile) XXX_Size() int {
	return xxx_messageInfo_Quantile.Size(m)
}
func (m *Quantile) XXX_DiscardUnknown() {
	xxx_messageInfo_Quantile.DiscardUnknown(m)
}

var xxx_messageInfo_Quantile proto.InternalMessageInfo

func (m *Quantile) GetQuantile() float64 {
	if m != nil && m.Quantile != nil {
		return *m.Quantile
	}
	return 0
}

func (m *Quantile) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Summary struct {
	SampleCount          *uint64     `protobuf:"varint,1,opt,name=sample_count,json=sampleCount" json:"sample_count,omitempty"`
	SampleSum            *float64    `protobuf:"fixed64,2,opt,name=sample_sum,json=sampleSum" json:"sample_sum,omitempty"`
	Quantile             []*Quantile `protobuf:"bytes,3,rep,name=quantile" json:"quantile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{4}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetSampleCount() uint64 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

func (m *Summary) GetSampleSum() float64 {
	if m != nil && m.SampleSum != nil {
		return *m.SampleSum
	}
	return 0
}

func (m *Summary) GetQuantile() []*Quantile {
	if m != nil {
		return m.Quantile
	}
	return nil
}

type Untyped struct {
	Value                *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Untyped) Reset()         { *m = Untyped{} }
func (m *Untyped) String() string { return proto.CompactTextString(m) }
func (*Untyped) ProtoMessage()    {}
func (*Untyped) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{5}
}

func (m *Untyped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Untyped.Unmarshal(m, b)
}
func (m *Untyped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Untyped.Marshal(b, m, deterministic)
}
func (m *Untyped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Untyped.Merge(m, src)
}
func (m *Untyped) XXX_Size() int {
	return xxx_messageInfo_Untyped.Size(m)
}
func (m *Untyped) XXX_DiscardUnknown() {
	xxx_messageInfo_Untyped.DiscardUnknown(m)
}

var xxx_messageInfo_Untyped proto.InternalMessageInfo

func (m *Untyped) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Histogram struct {
	SampleCount      *uint64  `protobuf:"varint,1,opt,name=sample_count,json=sampleCount" json:"sample_count,omitempty"`
	SampleCountFloat *float64 `protobuf:"fixed64,4,opt,name=sample_count_float,json=sampleCountFloat" json:"sample_count_float,omitempty"`
	SampleSum        *float64 `protobuf:"fixed64,2,opt,name=sample_sum,json=sampleSum" json:"sample_sum,omitempty"`
	// Buckets for the conventional histogram.
	Bucket []*Bucket `protobuf:"bytes,3,rep,name=bucket" json:"bucket,omitempty"`
	// schema defines the bucket schema. Currently, valid numbers are -4 <= n <= 8.
	// They are all for base-2 bucket schemas, where 1 is a bucket boundary in each case, and
	// then each power of two is divided into 2^n logarithmic buckets.
	// Or in other words, each bucket boundary is the previous boundary times 2^(2^-n).
	// In the future, more bucket schemas may be added using numbers < -4 or > 8.
	Schema         *int32   `protobuf:"zigzag32,5,opt,name=schema" json:"schema,omitempty"`
	ZeroThreshold  *float64 `protobuf:"fixed64,6,opt,name=zero_threshold,json=zeroThreshold" json:"zero_threshold,omitempty"`
	ZeroCount      *uint64  `protobuf:"varint,7,opt,name=zero_count,json=zeroCount" json:"zero_count,omitempty"`
	ZeroCountFloat *float64 `protobuf:"fixed64,8,opt,name=zero_count_float,json=zeroCountFloat" json:"zero_count_float,omitempty"`
	// Negative buckets for the native histogram.
	NegativeSpan []*BucketSpan `protobuf:"bytes,9,rep,name=negative_span,json=negativeSpan" json:"negative_span,omitempty"`
	// Use either "negative_delta" or "negative_count", the former for
	// regular histograms with integer counts, the latter for float
	// histograms.
	NegativeDelta []int64   `protobuf:"zigzag64,10,rep,name=negative_delta,json=negativeDelta" json:"negative_delta,omitempty"`
	NegativeCount []float64 `protobuf:"fixed64,11,rep,name=negative_count,json=negativeCount" json:"negative_count,omitempty"`
	// Positive buckets for the native histogram.
	PositiveSpan []*BucketSpan `protobuf:"bytes,12,rep,name=positive_span,json=positiveSpan" json:"positive_span,omitempty"`
	// Use either "positive_delta" or "positive_count", the former for
	// regular histograms with integer counts, the latter for float
	// histograms.
	PositiveDelta        []int64   `protobuf:"zigzag64,13,rep,name=positive_delta,json=positiveDelta" json:"positive_delta,omitempty"`
	PositiveCount        []float64 `protobuf:"fixed64,14,rep,name=positive_count,json=positiveCount" json:"positive_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Histogram) Reset()         { *m = Histogram{} }
func (m *Histogram) String() string { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()    {}
func (*Histogram) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{6}
}

func (m *Histogram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Histogram.Unmarshal(m, b)
}
func (m *Histogram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Histogram.Marshal(b, m, deterministic)
}
func (m *Histogram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Histogram.Merge(m, src)
}
func (m *Histogram) XXX_Size() int {
	return xxx_messageInfo_Histogram.Size(m)
}
func (m *Histogram) XXX_DiscardUnknown() {
	xxx_messageInfo_Histogram.DiscardUnknown(m)
}

var xxx_messageInfo_Histogram proto.InternalMessageInfo

func (m *Histogram) GetSampleCount() uint64 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

func (m *Histogram) GetSampleCountFloat() float64 {
	if m != nil && m.SampleCountFloat != nil {
		return *m.SampleCountFloat
	}
	return 0
}

func (m *Histogram) GetSampleSum() float64 {
	if m != nil && m.SampleSum != nil {
		return *m.SampleSum
	}
	return 0
}

func (m *Histogram) GetBucket() []*Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *Histogram) GetSchema() int32 {
	if m != nil && m.Schema != nil {
		return *m.Schema
	}
	return 0
}

func (m *Histogram) GetZeroThreshold() float64 {
	if m != nil && m.ZeroThreshold != nil {
		return *m.ZeroThreshold
	}
	return 0
}

func (m *Histogram) GetZeroCount() uint64 {
	if m != nil && m.ZeroCount != nil {
		return *m.ZeroCount
	}
	return 0
}

func (m *Histogram) GetZeroCountFloat() float64 {
	if m != nil && m.ZeroCountFloat != nil {
		return *m.ZeroCountFloat
	}
	return 0
}

func (m *Histogram) GetNegativeSpan() []*BucketSpan {
	if m != nil {
		return m.NegativeSpan
	}
	return nil
}

func (m *Histogram) GetNegativeDelta() []int64 {
	if m != nil {
		return m.NegativeDelta
	}
	return nil
}

func (m *Histogram) GetNegativeCount() []float64 {
	if m != nil {
		return m.NegativeCount
	}
	return nil
}

func (m *Histogram) GetPositiveSpan() []*BucketSpan {
	if m != nil {
		return m.PositiveSpan
	}
	return nil
}

func (m *Histogram) GetPositiveDelta() []int64 {
	if m != nil {
		return m.PositiveDelta
	}
	return nil
}

func (m *Histogram) GetPositiveCount() []float64 {
	if m != nil {
		return m.PositiveCount
	}
	return nil
}

// A Bucket of a conventional histogram, each of which is treated as
// an individual counter-like time series by Prometheus.
type Bucket struct {
	CumulativeCount      *uint64   `protobuf:"varint,1,opt,name=cumulative_count,json=cumulativeCount" json:"cumulative_count,omitempty"`
	CumulativeCountFloat *float64  `protobuf:"fixed64,4,opt,name=cumulative_count_float,json=cumulativeCountFloat" json:"cumulative_count_float,omitempty"`
	UpperBound           *float64  `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound" json:"upper_bound,omitempty"`
	Exemplar             *Exemplar `protobuf:"bytes,3,opt,name=exemplar" json:"exemplar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{7}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetCumulativeCount() uint64 {
	if m != nil && m.CumulativeCount != nil {
		return *m.CumulativeCount
	}
	return 0
}

func (m *Bucket) GetCumulativeCountFloat() float64 {
	if m != nil && m.CumulativeCountFloat != nil {
		return *m.CumulativeCountFloat
	}
	return 0
}

func (m *Bucket) GetUpperBound() float64 {
	if m != nil && m.UpperBound != nil {
		return *m.UpperBound
	}
	return 0
}

func (m *Bucket) GetExemplar() *Exemplar {
	if m != nil {
		return m.Exemplar
	}
	return nil
}

// A BucketSpan defines a number of consecutive buckets in a native
// histogram with their offset. Logically, it would be more
// straightforward to include the bucket counts in the Span. However,
// the protobuf representation is more compact in the way the data is
// structured here (with all the buckets in a single array separate
// from the Spans).
type BucketSpan struct {
	Offset               *int32   `protobuf:"zigzag32,1,opt,name=offset" json:"offset,omitempty"`
	Length               *uint32  `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketSpan) Reset()         { *m = BucketSpan{} }
func (m *BucketSpan) String() string { return proto.CompactTextString(m) }
func (*BucketSpan) ProtoMessage()    {}
func (*BucketSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{8}
}

func (m *BucketSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketSpan.Unmarshal(m, b)
}
func (m *BucketSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketSpan.Marshal(b, m, deterministic)
}
func (m *BucketSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketSpan.Merge(m, src)
}
func (m *BucketSpan) XXX_Size() int {
	return xxx_messageInfo_BucketSpan.Size(m)
}
func (m *BucketSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketSpan.DiscardUnknown(m)
}

var xxx_messageInfo_BucketSpan proto.InternalMessageInfo

func (m *BucketSpan) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *BucketSpan) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type Exemplar struct {
	Label                []*LabelPair         `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Value                *float64             `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Exemplar) Reset()         { *m = Exemplar{} }
func (m *Exemplar) String() string { return proto.CompactTextString(m) }
func (*Exemplar) ProtoMessage()    {}
func (*Exemplar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{9}
}

func (m *Exemplar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exemplar.Unmarshal(m, b)
}
func (m *Exemplar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exemplar.Marshal(b, m, deterministic)
}
func (m *Exemplar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exemplar.Merge(m, src)
}
func (m *Exemplar) XXX_Size() int {
	return xxx_messageInfo_Exemplar.Size(m)
}
func (m *Exemplar) XXX_DiscardUnknown() {
	xxx_messageInfo_Exemplar.DiscardUnknown(m)
}

var xxx_messageInfo_Exemplar proto.InternalMessageInfo

func (m *Exemplar) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Exemplar) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Exemplar) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Metric struct {
	Label                []*LabelPair `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Gauge                *Gauge       `protobuf:"bytes,2,opt,name=gauge" json:"gauge,omitempty"`
	Counter              *Counter     `protobuf:"bytes,3,opt,name=counter" json:"counter,omitempty"`
	Summary              *Summary     `protobuf:"bytes,4,opt,name=summary" json:"summary,omitempty"`
	Untyped              *Untyped     `protobuf:"bytes,5,opt,name=untyped" json:"untyped,omitempty"`
	Histogram            *Histogram   `protobuf:"bytes,7,opt,name=histogram" json:"histogram,omitempty"`
	TimestampMs          *int64       `protobuf:"varint,6,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{10}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Metric) GetGauge() *Gauge {
	if m != nil {
		return m.Gauge
	}
	return nil
}

func (m *Metric) GetCounter() *Counter {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *Metric) GetSummary() *Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Metric) GetUntyped() *Untyped {
	if m != nil {
		return m.Untyped
	}
	return nil
}

func (m *Metric) GetHistogram() *Histogram {
	if m != nil {
		return m.Histogram
	}
	return nil
}

func (m *Metric) GetTimestampMs() int64 {
	if m != nil && m.TimestampMs != nil {
		return *m.TimestampMs
	}
	return 0
}

type MetricFamily struct {
	Name                 *string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Help                 *string     `protobuf:"bytes,2,opt,name=help" json:"help,omitempty"`
	Type                 *MetricType `protobuf:"varint,3,opt,name=type,enum=io.prometheus.client.MetricType" json:"type,omitempty"`
	Metric               []*Metric   `protobuf:"bytes,4,rep,name=metric" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MetricFamily) Reset()         { *m = MetricFamily{} }
func (m *MetricFamily) String() string { return proto.CompactTextString(m) }
func (*MetricFamily) ProtoMessage()    {}
func (*MetricFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e5ddb18987a258, []int{11}
}

func (m *MetricFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricFamily.Unmarshal(m, b)
}
func (m *MetricFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricFamily.Marshal(b, m, deterministic)
}
func (m *MetricFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricFamily.Merge(m, src)
}
func (m *MetricFamily) XXX_Size() int {
	return xxx_messageInfo_MetricFamily.Size(m)
}
func (m *MetricFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricFamily.DiscardUnknown(m)
}

var xxx_messageInfo_MetricFamily proto.InternalMessageInfo

func (m *MetricFamily) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MetricFamily) GetHelp() string {
	if m != nil && m.Help != nil {
		return *m.Help
	}
	return ""
}

func (m *MetricFamily) GetType() MetricType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MetricType_COUNTER
}

func (m *MetricFamily) GetMetric() []*Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.prometheus.client.MetricType", MetricType_name, MetricType_value)
	proto.RegisterType((*LabelPair)(nil), "io.prometheus.client.LabelPair")
	proto.RegisterType((*Gauge)(nil), "io.prometheus.client.Gauge")
	proto.RegisterType((*Counter)(nil), "io.prometheus.client.Counter")
	proto.RegisterType((*Quantile)(nil), "io.prometheus.client.Quantile")
	proto.RegisterType((*Summary)(nil), "io.prometheus.client.Summary")
	proto.RegisterType((*Untyped)(nil), "io.prometheus.client.Untyped")
	proto.RegisterType((*Histogram)(nil), "io.prometheus.client.Histogram")
	proto.RegisterType((*Bucket)(nil), "io.prometheus.client.Bucket")
	proto.RegisterType((*BucketSpan)(nil), "io.prometheus.client.BucketSpan")
	proto.RegisterType((*Exemplar)(nil), "io.prometheus.client.Exemplar")
	proto.RegisterType((*Metric)(nil), "io.prometheus.client.Metric")
	proto.RegisterType((*MetricFamily)(nil), "io.prometheus.client.MetricFamily")
}

func init() {
	proto.RegisterFile("io/prometheus/client/metrics.proto", fileDescriptor_d1e5ddb18987a258)
}

var fileDescriptor_d1e5ddb18987a258 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x18, 0xc5, 0x9b, 0x5f, 0x7f, 0xd9, 0x6c, 0xd3, 0x61, 0x55, 0x59, 0x0b, 0xcb, 0x06, 0x4b, 0x48,
	0x0b, 0x42, 0x8e, 0x40, 0x5b, 0x81, 0x0a, 0x5c, 0xec, 0xb6, 0xe9, 0x16, 0x89, 0xb4, 0x65, 0x92,
	0x5c, 0x14, 0x2e, 0xac, 0x49, 0x32, 0xeb, 0x58, 0x78, 0x3c, 0xc6, 0x1e, 0x57, 0x2c, 0x2f, 0xc0,
	0x35, 0xaf, 0xc0, 0xc3, 0xf0, 0x22, 0x3c, 0x08, 0x68, 0xfe, 0xec, 0xdd, 0xe2, 0x94, 0xd2, 0x3b,
	0x7f, 0x67, 0xce, 0xf7, 0xcd, 0x39, 0xe3, 0xc9, 0x71, 0xc0, 0x8f, 0xf9, 0x24, 0xcb, 0x39, 0xa3,
	0x62, 0x4b, 0xcb, 0x62, 0xb2, 0x4e, 0x62, 0x9a, 0x8a, 0x09, 0xa3, 0x22, 0x8f, 0xd7, 0x45, 0x90,
	0xe5, 0x5c, 0x70, 0x74, 0x18, 0xf3, 0xa0, 0xe6, 0x04, 0x9a, 0x73, 0x74, 0x12, 0x71, 0x1e, 0x25,
	0x74, 0xa2, 0x38, 0xab, 0xf2, 0x6a, 0x22, 0x62, 0x46, 0x0b, 0x41, 0x58, 0xa6, 0xdb, 0xfc, 0xfb,
	0xe0, 0x7e, 0x47, 0x56, 0x34, 0x79, 0x4e, 0xe2, 0x1c, 0x21, 0x68, 0xa7, 0x84, 0x51, 0xcf, 0x19,
	0x3b, 0xa7, 0x2e, 0x56, 0xcf, 0xe8, 0x10, 0x3a, 0x2f, 0x49, 0x52, 0x52, 0x6f, 0x4f, 0x81, 0xba,
	0xf0, 0x8f, 0xa1, 0x73, 0x49, 0xca, 0xe8, 0xc6, 0xb2, 0xec, 0x71, 0xec, 0xf2, 0x8f, 0xd0, 0x7b,
	0xc8, 0xcb, 0x54, 0xd0, 0xbc, 0x99, 0x80, 0x1e, 0x40, 0x9f, 0xfe, 0x42, 0x59, 0x96, 0x90, 0x5c,
	0x0d, 0x1e, 0x7c, 0xfe, 0x41, 0xd0, 0x64, 0x20, 0x98, 0x1a, 0x16, 0xae, 0xf8, 0xfe, 0xd7, 0xd0,
	0xff, 0xbe, 0x24, 0xa9, 0x88, 0x13, 0x8a, 0x8e, 0xa0, 0xff, 0xb3, 0x79, 0x36, 0x1b, 0x54, 0xf5,
	0x6d, 0xe5, 0x95, 0xb4, 0xdf, 0x1c, 0xe8, 0xcd, 0x4b, 0xc6, 0x48, 0x7e, 0x8d, 0x3e, 0x84, 0xfd,
	0x82, 0xb0, 0x2c, 0xa1, 0xe1, 0x5a, 0xaa, 0x55, 0x13, 0xda, 0x78, 0xa0, 0x31, 0x65, 0x00, 0x1d,
	0x03, 0x18, 0x4a, 0x51, 0x32, 0x33, 0xc9, 0xd5, 0xc8, 0xbc, 0x64, 0xd2, 0x47, 0xb5, 0x7f, 0x6b,
	0xdc, 0xda, 0xed, 0xc3, 0x2a, 0xae, 0xf5, 0xf9, 0x27, 0xd0, 0x5b, 0xa6, 0xe2, 0x3a, 0xa3, 0x9b,
	0x1d, 0xa7, 0xf8, 0x57, 0x1b, 0xdc, 0x27, 0x71, 0x21, 0x78, 0x94, 0x13, 0xf6, 0x26, 0x62, 0x3f,
	0x05, 0x74, 0x93, 0x12, 0x5e, 0x25, 0x9c, 0x08, 0xaf, 0xad, 0x66, 0x8e, 0x6e, 0x10, 0x1f, 0x4b,
	0xfc, 0xbf, 0xac, 0x9d, 0x41, 0x77, 0x55, 0xae, 0x7f, 0xa2, 0xc2, 0x18, 0x7b, 0xbf, 0xd9, 0xd8,
	0x85, 0xe2, 0x60, 0xc3, 0x45, 0xf7, 0xa0, 0x5b, 0xac, 0xb7, 0x94, 0x11, 0xaf, 0x33, 0x76, 0x4e,
	0xef, 0x62, 0x53, 0xa1, 0x8f, 0xe0, 0xe0, 0x57, 0x9a, 0xf3, 0x50, 0x6c, 0x73, 0x5a, 0x6c, 0x79,
	0xb2, 0xf1, 0xba, 0x6a, 0xc3, 0xa1, 0x44, 0x17, 0x16, 0x94, 0x9a, 0x14, 0x4d, 0x5b, 0xec, 0x29,
	0x8b, 0xae, 0x44, 0xb4, 0xc1, 0x53, 0x18, 0xd5, 0xcb, 0xc6, 0x5e, 0x5f, 0xcd, 0x39, 0xa8, 0x48,
	0xda, 0xdc, 0x14, 0x86, 0x29, 0x8d, 0x88, 0x88, 0x5f, 0xd2, 0xb0, 0xc8, 0x48, 0xea, 0xb9, 0xca,
	0xc4, 0xf8, 0x75, 0x26, 0xe6, 0x19, 0x49, 0xf1, 0xbe, 0x6d, 0x93, 0x95, 0x94, 0x5d, 0x8d, 0xd9,
	0xd0, 0x44, 0x10, 0x0f, 0xc6, 0xad, 0x53, 0x84, 0xab, 0xe1, 0x8f, 0x24, 0x78, 0x8b, 0xa6, 0xa5,
	0x0f, 0xc6, 0x2d, 0xe9, 0xce, 0xa2, 0x5a, 0xfe, 0x14, 0x86, 0x19, 0x2f, 0xe2, 0x5a, 0xd4, 0xfe,
	0x9b, 0x8a, 0xb2, 0x6d, 0x56, 0x54, 0x35, 0x46, 0x8b, 0x1a, 0x6a, 0x51, 0x16, 0xad, 0x44, 0x55,
	0x34, 0x2d, 0xea, 0x40, 0x8b, 0xb2, 0xa8, 0x12, 0xe5, 0xff, 0xe9, 0x40, 0x57, 0x6f, 0x85, 0x3e,
	0x86, 0xd1, 0xba, 0x64, 0x65, 0x72, 0xd3, 0x88, 0xbe, 0x66, 0x77, 0x6a, 0x5c, 0x5b, 0x39, 0x83,
	0x7b, 0xaf, 0x52, 0x6f, 0x5d, 0xb7, 0xc3, 0x57, 0x1a, 0xf4, 0x5b, 0x39, 0x81, 0x41, 0x99, 0x65,
	0x34, 0x0f, 0x57, 0xbc, 0x4c, 0x37, 0xe6, 0xce, 0x81, 0x82, 0x2e, 0x24, 0x72, 0x2b, 0x17, 0x5a,
	0xff, 0x3b, 0x17, 0xa0, 0x3e, 0x32, 0x79, 0x11, 0xf9, 0xd5, 0x55, 0x41, 0xb5, 0x83, 0xbb, 0xd8,
	0x54, 0x12, 0x4f, 0x68, 0x1a, 0x89, 0xad, 0xda, 0x7d, 0x88, 0x4d, 0xe5, 0xff, 0xee, 0x40, 0xdf,
	0x0e, 0x45, 0xf7, 0xa1, 0x93, 0xc8, 0x54, 0xf4, 0x1c, 0xf5, 0x82, 0x4e, 0x9a, 0x35, 0x54, 0xc1,
	0x89, 0x35, 0xbb, 0x39, 0x71, 0xd0, 0x97, 0xe0, 0x56, 0xa9, 0x6b, 0x4c, 0x1d, 0x05, 0x3a, 0x97,
	0x03, 0x9b, 0xcb, 0xc1, 0xc2, 0x32, 0x70, 0x4d, 0xf6, 0xff, 0xde, 0x83, 0xee, 0x4c, 0xa5, 0xfc,
	0xdb, 0x2a, 0xfa, 0x0c, 0x3a, 0x91, 0xcc, 0x69, 0x13, 0xb2, 0xef, 0x35, 0xb7, 0xa9, 0x28, 0xc7,
	0x9a, 0x89, 0xbe, 0x80, 0xde, 0x5a, 0x67, 0xb7, 0x11, 0x7b, 0xdc, 0xdc, 0x64, 0x02, 0x1e, 0x5b,
	0xb6, 0x6c, 0x2c, 0x74, 0xb0, 0xaa, 0x3b, 0xb0, 0xb3, 0xd1, 0xa4, 0x2f, 0xb6, 0x6c, 0xd9, 0x58,
	0xea, 0x20, 0x54, 0xa1, 0xb1, 0xb3, 0xd1, 0xa4, 0x25, 0xb6, 0x6c, 0xf4, 0x0d, 0xb8, 0x5b, 0x9b,
	0x8f, 0x2a, 0x2c, 0x76, 0x1e, 0x4c, 0x15, 0xa3, 0xb8, 0xee, 0x90, 0x89, 0x5a, 0x9d, 0x75, 0xc8,
	0x0a, 0x95, 0x48, 0x2d, 0x3c, 0xa8, 0xb0, 0x59, 0xe1, 0xff, 0xe1, 0xc0, 0xbe, 0x7e, 0x03, 0x8f,
	0x09, 0x8b, 0x93, 0xeb, 0xc6, 0x4f, 0x24, 0x82, 0xf6, 0x96, 0x26, 0x99, 0xf9, 0x42, 0xaa, 0x67,
	0x74, 0x06, 0x6d, 0xa9, 0x51, 0x1d, 0xe1, 0xc1, 0xae, 0x5f, 0xb8, 0x9e, 0xbc, 0xb8, 0xce, 0x28,
	0x56, 0x6c, 0x99, 0xb9, 0xfa, 0xab, 0xee, 0xb5, 0x5f, 0x97, 0xb9, 0xba, 0x0f, 0x1b, 0xee, 0x27,
	0x2b, 0x80, 0x7a, 0x12, 0x1a, 0x40, 0xef, 0xe1, 0xb3, 0xe5, 0xd3, 0xc5, 0x14, 0x8f, 0xde, 0x41,
	0x2e, 0x74, 0x2e, 0xcf, 0x97, 0x97, 0xd3, 0x91, 0x23, 0xf1, 0xf9, 0x72, 0x36, 0x3b, 0xc7, 0x2f,
	0x46, 0x7b, 0xb2, 0x58, 0x3e, 0x5d, 0xbc, 0x78, 0x3e, 0x7d, 0x34, 0x6a, 0xa1, 0x21, 0xb8, 0x4f,
	0xbe, 0x9d, 0x2f, 0x9e, 0x5d, 0xe2, 0xf3, 0xd9, 0xa8, 0x8d, 0xde, 0x85, 0x3b, 0xaa, 0x27, 0xac,
	0xc1, 0xce, 0x05, 0x86, 0xc6, 0x3f, 0x18, 0x3f, 0x3c, 0x88, 0x62, 0xb1, 0x2d, 0x57, 0xc1, 0x9a,
	0xb3, 0x7f, 0xff, 0x45, 0x09, 0x19, 0xdf, 0xd0, 0x64, 0x12, 0xf1, 0xaf, 0x62, 0x1e, 0xd6, 0xab,
	0xa1, 0x5e, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x16, 0x77, 0x81, 0x98, 0xd7, 0x08, 0x00, 0x00,
}
