// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.0
// source: trf.proto

package bpftests

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SuiteResult_TestResult_TestStatus int32

const (
	// Unable to execute test, not indicative of the test itself, usually an unexpected error in the framework
	// like out of memory conditions.
	SuiteResult_TestResult_ERROR SuiteResult_TestResult_TestStatus = 0
	SuiteResult_TestResult_PASS  SuiteResult_TestResult_TestStatus = 1
	SuiteResult_TestResult_FAIL  SuiteResult_TestResult_TestStatus = 2
	// The test was skipped, for example because the feature under test is disabled or unavailable with the
	// current settings or on this platform.
	SuiteResult_TestResult_SKIP SuiteResult_TestResult_TestStatus = 3
)

// Enum value maps for SuiteResult_TestResult_TestStatus.
var (
	SuiteResult_TestResult_TestStatus_name = map[int32]string{
		0: "ERROR",
		1: "PASS",
		2: "FAIL",
		3: "SKIP",
	}
	SuiteResult_TestResult_TestStatus_value = map[string]int32{
		"ERROR": 0,
		"PASS":  1,
		"FAIL":  2,
		"SKIP":  3,
	}
)

func (x SuiteResult_TestResult_TestStatus) Enum() *SuiteResult_TestResult_TestStatus {
	p := new(SuiteResult_TestResult_TestStatus)
	*p = x
	return p
}

func (x SuiteResult_TestResult_TestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SuiteResult_TestResult_TestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_trf_proto_enumTypes[0].Descriptor()
}

func (SuiteResult_TestResult_TestStatus) Type() protoreflect.EnumType {
	return &file_trf_proto_enumTypes[0]
}

func (x SuiteResult_TestResult_TestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SuiteResult_TestResult_TestStatus.Descriptor instead.
func (SuiteResult_TestResult_TestStatus) EnumDescriptor() ([]byte, []int) {
	return file_trf_proto_rawDescGZIP(), []int{0, 0, 0}
}

// TRF (Test Result Format)
type SuiteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results  []*SuiteResult_TestResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	SuiteLog []*Log                    `protobuf:"bytes,2,rep,name=suite_log,json=suiteLog,proto3" json:"suite_log,omitempty"`
}

func (x *SuiteResult) Reset() {
	*x = SuiteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuiteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuiteResult) ProtoMessage() {}

func (x *SuiteResult) ProtoReflect() protoreflect.Message {
	mi := &file_trf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuiteResult.ProtoReflect.Descriptor instead.
func (*SuiteResult) Descriptor() ([]byte, []int) {
	return file_trf_proto_rawDescGZIP(), []int{0}
}

func (x *SuiteResult) GetResults() []*SuiteResult_TestResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SuiteResult) GetSuiteLog() []*Log {
	if x != nil {
		return x.SuiteLog
	}
	return nil
}

// bpf_trace_printk style logging
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fmt  string   `protobuf:"bytes,1,opt,name=fmt,proto3" json:"fmt,omitempty"`
	Args []uint64 `protobuf:"fixed64,2,rep,packed,name=args,proto3" json:"args,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_trf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_trf_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetFmt() string {
	if x != nil {
		return x.Fmt
	}
	return ""
}

func (x *Log) GetArgs() []uint64 {
	if x != nil {
		return x.Args
	}
	return nil
}

type SuiteResult_TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  SuiteResult_TestResult_TestStatus `protobuf:"varint,2,opt,name=status,proto3,enum=SuiteResult_TestResult_TestStatus" json:"status,omitempty"`
	TestLog []*Log                            `protobuf:"bytes,3,rep,name=test_log,json=testLog,proto3" json:"test_log,omitempty"`
}

func (x *SuiteResult_TestResult) Reset() {
	*x = SuiteResult_TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuiteResult_TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuiteResult_TestResult) ProtoMessage() {}

func (x *SuiteResult_TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_trf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuiteResult_TestResult.ProtoReflect.Descriptor instead.
func (*SuiteResult_TestResult) Descriptor() ([]byte, []int) {
	return file_trf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SuiteResult_TestResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuiteResult_TestResult) GetStatus() SuiteResult_TestResult_TestStatus {
	if x != nil {
		return x.Status
	}
	return SuiteResult_TestResult_ERROR
}

func (x *SuiteResult_TestResult) GetTestLog() []*Log {
	if x != nil {
		return x.TestLog
	}
	return nil
}

var File_trf_proto protoreflect.FileDescriptor

var file_trf_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x72, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x0b,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x09, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x73, 0x75, 0x69, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x1a, 0xb4, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x22, 0x35, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x03, 0x22, 0x2b, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x6d, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6d,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x06, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x62, 0x70, 0x66, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trf_proto_rawDescOnce sync.Once
	file_trf_proto_rawDescData = file_trf_proto_rawDesc
)

func file_trf_proto_rawDescGZIP() []byte {
	file_trf_proto_rawDescOnce.Do(func() {
		file_trf_proto_rawDescData = protoimpl.X.CompressGZIP(file_trf_proto_rawDescData)
	})
	return file_trf_proto_rawDescData
}

var file_trf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trf_proto_goTypes = []interface{}{
	(SuiteResult_TestResult_TestStatus)(0), // 0: SuiteResult.TestResult.TestStatus
	(*SuiteResult)(nil),                    // 1: SuiteResult
	(*Log)(nil),                            // 2: Log
	(*SuiteResult_TestResult)(nil),         // 3: SuiteResult.TestResult
}
var file_trf_proto_depIdxs = []int32{
	3, // 0: SuiteResult.results:type_name -> SuiteResult.TestResult
	2, // 1: SuiteResult.suite_log:type_name -> Log
	0, // 2: SuiteResult.TestResult.status:type_name -> SuiteResult.TestResult.TestStatus
	2, // 3: SuiteResult.TestResult.test_log:type_name -> Log
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_trf_proto_init() }
func file_trf_proto_init() {
	if File_trf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuiteResult); i {
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
		file_trf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_trf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuiteResult_TestResult); i {
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
			RawDescriptor: file_trf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trf_proto_goTypes,
		DependencyIndexes: file_trf_proto_depIdxs,
		EnumInfos:         file_trf_proto_enumTypes,
		MessageInfos:      file_trf_proto_msgTypes,
	}.Build()
	File_trf_proto = out.File
	file_trf_proto_rawDesc = nil
	file_trf_proto_goTypes = nil
	file_trf_proto_depIdxs = nil
}
