// Copyright 2021 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: recorder/recorder.proto

package recorder

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Protocol is a one of the supported protocols for packet capture
type Protocol int32

const (
	Protocol_PROTOCOL_ANY Protocol = 0
	Protocol_PROTOCOL_TCP Protocol = 6
	Protocol_PROTOCOL_UDP Protocol = 17
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0:  "PROTOCOL_ANY",
		6:  "PROTOCOL_TCP",
		17: "PROTOCOL_UDP",
	}
	Protocol_value = map[string]int32{
		"PROTOCOL_ANY": 0,
		"PROTOCOL_TCP": 6,
		"PROTOCOL_UDP": 17,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_recorder_recorder_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_recorder_recorder_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{0}
}

type RecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestType:
	//	*RecordRequest_Start
	//	*RecordRequest_Stop
	RequestType isRecordRequest_RequestType `protobuf_oneof:"request_type"`
}

func (x *RecordRequest) Reset() {
	*x = RecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRequest) ProtoMessage() {}

func (x *RecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRequest.ProtoReflect.Descriptor instead.
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{0}
}

func (m *RecordRequest) GetRequestType() isRecordRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (x *RecordRequest) GetStart() *StartRecording {
	if x, ok := x.GetRequestType().(*RecordRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *RecordRequest) GetStop() *StopRecording {
	if x, ok := x.GetRequestType().(*RecordRequest_Stop); ok {
		return x.Stop
	}
	return nil
}

type isRecordRequest_RequestType interface {
	isRecordRequest_RequestType()
}

type RecordRequest_Start struct {
	// start starts a new recording with the given parameters.
	Start *StartRecording `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type RecordRequest_Stop struct {
	// stop stops the running recording.
	Stop *StopRecording `protobuf:"bytes,2,opt,name=stop,proto3,oneof"`
}

func (*RecordRequest_Start) isRecordRequest_RequestType() {}

func (*RecordRequest_Stop) isRecordRequest_RequestType() {}

type StartRecording struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filesink configures the outfile of this recording
	// Future alternative sink configurations may be added as a
	// backwards-compatible change by moving this field into a oneof.
	Filesink *FileSinkConfiguration `protobuf:"bytes,1,opt,name=filesink,proto3" json:"filesink,omitempty"`
	// include list for this recording. Packets matching any of the provided
	// filters will be recorded.
	Include []*Filter `protobuf:"bytes,2,rep,name=include,proto3" json:"include,omitempty"`
	// max_capture_length specifies the maximum packet length.
	// Full packet length will be captured if absent/zero.
	MaxCaptureLength uint32 `protobuf:"varint,3,opt,name=max_capture_length,json=maxCaptureLength,proto3" json:"max_capture_length,omitempty"`
}

func (x *StartRecording) Reset() {
	*x = StartRecording{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRecording) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecording) ProtoMessage() {}

func (x *StartRecording) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecording.ProtoReflect.Descriptor instead.
func (*StartRecording) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{1}
}

func (x *StartRecording) GetFilesink() *FileSinkConfiguration {
	if x != nil {
		return x.Filesink
	}
	return nil
}

func (x *StartRecording) GetInclude() []*Filter {
	if x != nil {
		return x.Include
	}
	return nil
}

func (x *StartRecording) GetMaxCaptureLength() uint32 {
	if x != nil {
		return x.MaxCaptureLength
	}
	return 0
}

// FileSinkConfiguration configures the file output. Possible future additions
// might be the selection of the output volume. The initial implementation will
// only support a single volume which is configured as a cilium-agent CLI flag.
type FileSinkConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// file_prefix is an optional prefix for the file name.
	// Defaults to `hubble` if empty. Must match the following regex if not
	// empty: ^[a-z][a-z0-9]{0,19}$
	// The generated filename will be of format
	//  <file_prefix>_<unixtime>_<unique_random>_<node_name>.pcap
	FilePrefix string `protobuf:"bytes,1,opt,name=file_prefix,json=filePrefix,proto3" json:"file_prefix,omitempty"`
}

func (x *FileSinkConfiguration) Reset() {
	*x = FileSinkConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSinkConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSinkConfiguration) ProtoMessage() {}

func (x *FileSinkConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSinkConfiguration.ProtoReflect.Descriptor instead.
func (*FileSinkConfiguration) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{2}
}

func (x *FileSinkConfiguration) GetFilePrefix() string {
	if x != nil {
		return x.FilePrefix
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source_cidr. Must not be empty.
	// Set to 0.0.0.0/0 to match any IPv4 source address (::/0 for IPv6).
	SourceCidr string `protobuf:"bytes,1,opt,name=source_cidr,json=sourceCidr,proto3" json:"source_cidr,omitempty"`
	// source_port. Matches any source port if empty.
	SourcePort uint32 `protobuf:"varint,2,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// destination_cidr. Must not be empty.
	// Set to 0.0.0.0/0 to match any IPv4 destination address (::/0 for IPv6).
	DestinationCidr string `protobuf:"bytes,3,opt,name=destination_cidr,json=destinationCidr,proto3" json:"destination_cidr,omitempty"`
	// destination_port. Matches any destination port if empty.
	DestinationPort uint32 `protobuf:"varint,4,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// protocol. Matches any protocol if empty.
	Protocol Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=recorder.Protocol" json:"protocol,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetSourceCidr() string {
	if x != nil {
		return x.SourceCidr
	}
	return ""
}

func (x *Filter) GetSourcePort() uint32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *Filter) GetDestinationCidr() string {
	if x != nil {
		return x.DestinationCidr
	}
	return ""
}

func (x *Filter) GetDestinationPort() uint32 {
	if x != nil {
		return x.DestinationPort
	}
	return 0
}

func (x *Filter) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_PROTOCOL_ANY
}

type StopRecording struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRecording) Reset() {
	*x = StopRecording{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRecording) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRecording) ProtoMessage() {}

func (x *StopRecording) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRecording.ProtoReflect.Descriptor instead.
func (*StopRecording) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{4}
}

type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the node where this recording is happening
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// time at which this event was observed on the above node
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// Note: In this initial design, any fatal error will be returned as
	// gRPC errors and are not part of the regular response type.
	// It is a forward-compatible change to introduce additional more
	// granular or structured error responses here.
	//
	// Types that are assignable to ResponseType:
	//	*RecordResponse_Running
	//	*RecordResponse_Stopped
	ResponseType isRecordResponse_ResponseType `protobuf_oneof:"response_type"`
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{5}
}

func (x *RecordResponse) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *RecordResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (m *RecordResponse) GetResponseType() isRecordResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (x *RecordResponse) GetRunning() *RecordingRunningResponse {
	if x, ok := x.GetResponseType().(*RecordResponse_Running); ok {
		return x.Running
	}
	return nil
}

func (x *RecordResponse) GetStopped() *RecordingStoppedResponse {
	if x, ok := x.GetResponseType().(*RecordResponse_Stopped); ok {
		return x.Stopped
	}
	return nil
}

type isRecordResponse_ResponseType interface {
	isRecordResponse_ResponseType()
}

type RecordResponse_Running struct {
	// running means that the recording is capturing packets. This is
	// emitted in regular intervals
	Running *RecordingRunningResponse `protobuf:"bytes,3,opt,name=running,proto3,oneof"`
}

type RecordResponse_Stopped struct {
	// stopped means the recording has stopped
	Stopped *RecordingStoppedResponse `protobuf:"bytes,4,opt,name=stopped,proto3,oneof"`
}

func (*RecordResponse_Running) isRecordResponse_ResponseType() {}

func (*RecordResponse_Stopped) isRecordResponse_ResponseType() {}

type RecordingStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bytes_captured is the total amount of bytes captured in the recording
	BytesCaptured uint64 `protobuf:"varint,1,opt,name=bytes_captured,json=bytesCaptured,proto3" json:"bytes_captured,omitempty"`
	// packets_captured is the total amount of packets captured the recording
	PacketsCaptured uint64 `protobuf:"varint,2,opt,name=packets_captured,json=packetsCaptured,proto3" json:"packets_captured,omitempty"`
	// packets_lost is the total amount of packets matching the filter during
	// the recording, but never written to the sink because it was overloaded.
	PacketsLost uint64 `protobuf:"varint,3,opt,name=packets_lost,json=packetsLost,proto3" json:"packets_lost,omitempty"`
	// bytes_lost is the total amount of bytes matching the filter during
	// the recording, but never written to the sink because it was overloaded.
	BytesLost uint64 `protobuf:"varint,4,opt,name=bytes_lost,json=bytesLost,proto3" json:"bytes_lost,omitempty"`
}

func (x *RecordingStatistics) Reset() {
	*x = RecordingStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingStatistics) ProtoMessage() {}

func (x *RecordingStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingStatistics.ProtoReflect.Descriptor instead.
func (*RecordingStatistics) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{6}
}

func (x *RecordingStatistics) GetBytesCaptured() uint64 {
	if x != nil {
		return x.BytesCaptured
	}
	return 0
}

func (x *RecordingStatistics) GetPacketsCaptured() uint64 {
	if x != nil {
		return x.PacketsCaptured
	}
	return 0
}

func (x *RecordingStatistics) GetPacketsLost() uint64 {
	if x != nil {
		return x.PacketsLost
	}
	return 0
}

func (x *RecordingStatistics) GetBytesLost() uint64 {
	if x != nil {
		return x.BytesLost
	}
	return 0
}

type RecordingRunningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stats for the running recording
	Stats *RecordingStatistics `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *RecordingRunningResponse) Reset() {
	*x = RecordingRunningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingRunningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingRunningResponse) ProtoMessage() {}

func (x *RecordingRunningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingRunningResponse.ProtoReflect.Descriptor instead.
func (*RecordingRunningResponse) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{7}
}

func (x *RecordingRunningResponse) GetStats() *RecordingStatistics {
	if x != nil {
		return x.Stats
	}
	return nil
}

type RecordingStoppedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stats for the recording
	Stats *RecordingStatistics `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	// filesink contains the path to the captured file
	Filesink *FileSinkResult `protobuf:"bytes,2,opt,name=filesink,proto3" json:"filesink,omitempty"`
}

func (x *RecordingStoppedResponse) Reset() {
	*x = RecordingStoppedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingStoppedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingStoppedResponse) ProtoMessage() {}

func (x *RecordingStoppedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingStoppedResponse.ProtoReflect.Descriptor instead.
func (*RecordingStoppedResponse) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{8}
}

func (x *RecordingStoppedResponse) GetStats() *RecordingStatistics {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *RecordingStoppedResponse) GetFilesink() *FileSinkResult {
	if x != nil {
		return x.Filesink
	}
	return nil
}

type FileSinkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// file_path is the absolute path to the captured pcap file
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *FileSinkResult) Reset() {
	*x = FileSinkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_recorder_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSinkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSinkResult) ProtoMessage() {}

func (x *FileSinkResult) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_recorder_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSinkResult.ProtoReflect.Descriptor instead.
func (*FileSinkResult) Descriptor() ([]byte, []int) {
	return file_recorder_recorder_proto_rawDescGZIP(), []int{9}
}

func (x *FileSinkResult) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_recorder_recorder_proto protoreflect.FileDescriptor

var file_recorder_recorder_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x69, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x22, 0x38, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xd0, 0x01, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x69, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x69, 0x64, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2e,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x0f,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0xee, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x3e, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x42,
	0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xa9, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x18,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x85, 0x01,
	0x0a, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x70, 0x70,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x34, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x69, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x2a, 0x40, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x41, 0x4e, 0x59,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x54,
	0x43, 0x50, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c,
	0x5f, 0x55, 0x44, 0x50, 0x10, 0x11, 0x32, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recorder_recorder_proto_rawDescOnce sync.Once
	file_recorder_recorder_proto_rawDescData = file_recorder_recorder_proto_rawDesc
)

func file_recorder_recorder_proto_rawDescGZIP() []byte {
	file_recorder_recorder_proto_rawDescOnce.Do(func() {
		file_recorder_recorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_recorder_recorder_proto_rawDescData)
	})
	return file_recorder_recorder_proto_rawDescData
}

var file_recorder_recorder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_recorder_recorder_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_recorder_recorder_proto_goTypes = []interface{}{
	(Protocol)(0),                    // 0: recorder.Protocol
	(*RecordRequest)(nil),            // 1: recorder.RecordRequest
	(*StartRecording)(nil),           // 2: recorder.StartRecording
	(*FileSinkConfiguration)(nil),    // 3: recorder.FileSinkConfiguration
	(*Filter)(nil),                   // 4: recorder.Filter
	(*StopRecording)(nil),            // 5: recorder.StopRecording
	(*RecordResponse)(nil),           // 6: recorder.RecordResponse
	(*RecordingStatistics)(nil),      // 7: recorder.RecordingStatistics
	(*RecordingRunningResponse)(nil), // 8: recorder.RecordingRunningResponse
	(*RecordingStoppedResponse)(nil), // 9: recorder.RecordingStoppedResponse
	(*FileSinkResult)(nil),           // 10: recorder.FileSinkResult
	(*timestamp.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_recorder_recorder_proto_depIdxs = []int32{
	2,  // 0: recorder.RecordRequest.start:type_name -> recorder.StartRecording
	5,  // 1: recorder.RecordRequest.stop:type_name -> recorder.StopRecording
	3,  // 2: recorder.StartRecording.filesink:type_name -> recorder.FileSinkConfiguration
	4,  // 3: recorder.StartRecording.include:type_name -> recorder.Filter
	0,  // 4: recorder.Filter.protocol:type_name -> recorder.Protocol
	11, // 5: recorder.RecordResponse.time:type_name -> google.protobuf.Timestamp
	8,  // 6: recorder.RecordResponse.running:type_name -> recorder.RecordingRunningResponse
	9,  // 7: recorder.RecordResponse.stopped:type_name -> recorder.RecordingStoppedResponse
	7,  // 8: recorder.RecordingRunningResponse.stats:type_name -> recorder.RecordingStatistics
	7,  // 9: recorder.RecordingStoppedResponse.stats:type_name -> recorder.RecordingStatistics
	10, // 10: recorder.RecordingStoppedResponse.filesink:type_name -> recorder.FileSinkResult
	1,  // 11: recorder.Recorder.Record:input_type -> recorder.RecordRequest
	6,  // 12: recorder.Recorder.Record:output_type -> recorder.RecordResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_recorder_recorder_proto_init() }
func file_recorder_recorder_proto_init() {
	if File_recorder_recorder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recorder_recorder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRequest); i {
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
		file_recorder_recorder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRecording); i {
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
		file_recorder_recorder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSinkConfiguration); i {
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
		file_recorder_recorder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_recorder_recorder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRecording); i {
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
		file_recorder_recorder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
		file_recorder_recorder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingStatistics); i {
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
		file_recorder_recorder_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingRunningResponse); i {
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
		file_recorder_recorder_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingStoppedResponse); i {
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
		file_recorder_recorder_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSinkResult); i {
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
	file_recorder_recorder_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RecordRequest_Start)(nil),
		(*RecordRequest_Stop)(nil),
	}
	file_recorder_recorder_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RecordResponse_Running)(nil),
		(*RecordResponse_Stopped)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recorder_recorder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recorder_recorder_proto_goTypes,
		DependencyIndexes: file_recorder_recorder_proto_depIdxs,
		EnumInfos:         file_recorder_recorder_proto_enumTypes,
		MessageInfos:      file_recorder_recorder_proto_msgTypes,
	}.Build()
	File_recorder_recorder_proto = out.File
	file_recorder_recorder_proto_rawDesc = nil
	file_recorder_recorder_proto_goTypes = nil
	file_recorder_recorder_proto_depIdxs = nil
}
