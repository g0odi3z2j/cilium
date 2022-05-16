// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/type/metadata/v3/metadata.proto

package metadatav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// MetadataKey provides a general interface using `key` and `path` to retrieve value from
// :ref:`Metadata <envoy_v3_api_msg_config.core.v3.Metadata>`.
//
// For example, for the following Metadata:
//
// .. code-block:: yaml
//
//	filter_metadata:
//	  envoy.xxx:
//	    prop:
//	      foo: bar
//	      xyz:
//	        hello: envoy
//
// The following MetadataKey will retrieve a string value "bar" from the Metadata.
//
// .. code-block:: yaml
//
//	key: envoy.xxx
//	path:
//	- key: prop
//	- key: foo
type MetadataKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key name of Metadata to retrieve the Struct from the metadata.
	// Typically, it represents a builtin subsystem or custom extension.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The path to retrieve the Value from the Struct. It can be a prefix or a full path,
	// e.g. “[prop, xyz]“ for a struct or “[prop, foo]“ for a string in the example,
	// which depends on the particular scenario.
	//
	// Note: Due to that only the key type segment is supported, the path can not specify a list
	// unless the list is the last segment.
	Path []*MetadataKey_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *MetadataKey) Reset() {
	*x = MetadataKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKey) ProtoMessage() {}

func (x *MetadataKey) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKey.ProtoReflect.Descriptor instead.
func (*MetadataKey) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MetadataKey) GetPath() []*MetadataKey_PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

// Describes what kind of metadata.
type MetadataKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*MetadataKind_Request_
	//	*MetadataKind_Route_
	//	*MetadataKind_Cluster_
	//	*MetadataKind_Host_
	Kind isMetadataKind_Kind `protobuf_oneof:"kind"`
}

func (x *MetadataKind) Reset() {
	*x = MetadataKind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKind) ProtoMessage() {}

func (x *MetadataKind) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKind.ProtoReflect.Descriptor instead.
func (*MetadataKind) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{1}
}

func (m *MetadataKind) GetKind() isMetadataKind_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *MetadataKind) GetRequest() *MetadataKind_Request {
	if x, ok := x.GetKind().(*MetadataKind_Request_); ok {
		return x.Request
	}
	return nil
}

func (x *MetadataKind) GetRoute() *MetadataKind_Route {
	if x, ok := x.GetKind().(*MetadataKind_Route_); ok {
		return x.Route
	}
	return nil
}

func (x *MetadataKind) GetCluster() *MetadataKind_Cluster {
	if x, ok := x.GetKind().(*MetadataKind_Cluster_); ok {
		return x.Cluster
	}
	return nil
}

func (x *MetadataKind) GetHost() *MetadataKind_Host {
	if x, ok := x.GetKind().(*MetadataKind_Host_); ok {
		return x.Host
	}
	return nil
}

type isMetadataKind_Kind interface {
	isMetadataKind_Kind()
}

type MetadataKind_Request_ struct {
	// Request kind of metadata.
	Request *MetadataKind_Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type MetadataKind_Route_ struct {
	// Route kind of metadata.
	Route *MetadataKind_Route `protobuf:"bytes,2,opt,name=route,proto3,oneof"`
}

type MetadataKind_Cluster_ struct {
	// Cluster kind of metadata.
	Cluster *MetadataKind_Cluster `protobuf:"bytes,3,opt,name=cluster,proto3,oneof"`
}

type MetadataKind_Host_ struct {
	// Host kind of metadata.
	Host *MetadataKind_Host `protobuf:"bytes,4,opt,name=host,proto3,oneof"`
}

func (*MetadataKind_Request_) isMetadataKind_Kind() {}

func (*MetadataKind_Route_) isMetadataKind_Kind() {}

func (*MetadataKind_Cluster_) isMetadataKind_Kind() {}

func (*MetadataKind_Host_) isMetadataKind_Kind() {}

// Specifies the segment in a path to retrieve value from Metadata.
// Currently it is only supported to specify the key, i.e. field name, as one segment of a path.
type MetadataKey_PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Segment:
	//
	//	*MetadataKey_PathSegment_Key
	Segment isMetadataKey_PathSegment_Segment `protobuf_oneof:"segment"`
}

func (x *MetadataKey_PathSegment) Reset() {
	*x = MetadataKey_PathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKey_PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKey_PathSegment) ProtoMessage() {}

func (x *MetadataKey_PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKey_PathSegment.ProtoReflect.Descriptor instead.
func (*MetadataKey_PathSegment) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (m *MetadataKey_PathSegment) GetSegment() isMetadataKey_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *MetadataKey_PathSegment) GetKey() string {
	if x, ok := x.GetSegment().(*MetadataKey_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

type isMetadataKey_PathSegment_Segment interface {
	isMetadataKey_PathSegment_Segment()
}

type MetadataKey_PathSegment_Key struct {
	// If specified, use the key to retrieve the value in a Struct.
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataKey_PathSegment_Key) isMetadataKey_PathSegment_Segment() {}

// Represents dynamic metadata associated with the request.
type MetadataKind_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataKind_Request) Reset() {
	*x = MetadataKind_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKind_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKind_Request) ProtoMessage() {}

func (x *MetadataKind_Request) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKind_Request.ProtoReflect.Descriptor instead.
func (*MetadataKind_Request) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{1, 0}
}

// Represents metadata from :ref:`the route<envoy_v3_api_field_config.route.v3.Route.metadata>`.
type MetadataKind_Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataKind_Route) Reset() {
	*x = MetadataKind_Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKind_Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKind_Route) ProtoMessage() {}

func (x *MetadataKind_Route) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKind_Route.ProtoReflect.Descriptor instead.
func (*MetadataKind_Route) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{1, 1}
}

// Represents metadata from :ref:`the upstream cluster<envoy_v3_api_field_config.cluster.v3.Cluster.metadata>`.
type MetadataKind_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataKind_Cluster) Reset() {
	*x = MetadataKind_Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKind_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKind_Cluster) ProtoMessage() {}

func (x *MetadataKind_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKind_Cluster.ProtoReflect.Descriptor instead.
func (*MetadataKind_Cluster) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{1, 2}
}

// Represents metadata from :ref:`the upstream
// host<envoy_v3_api_field_config.endpoint.v3.LbEndpoint.metadata>`.
type MetadataKind_Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataKind_Host) Reset() {
	*x = MetadataKind_Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataKind_Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataKind_Host) ProtoMessage() {}

func (x *MetadataKind_Host) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_metadata_v3_metadata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataKind_Host.ProtoReflect.Descriptor instead.
func (*MetadataKind_Host) Descriptor() ([]byte, []int) {
	return file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP(), []int{1, 3}
}

var File_envoy_type_metadata_v3_metadata_proto protoreflect.FileDescriptor

var file_envoy_type_metadata_v3_metadata_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4d, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x1a, 0x71, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x3a, 0x35, 0x9a, 0xc5, 0x88, 0x1e, 0x30, 0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x3a, 0x29, 0x9a, 0xc5, 0x88, 0x1e, 0x24, 0x0a, 0x22,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x65, 0x79, 0x22, 0xd2, 0x04, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69,
	0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x3d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x69, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x05, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x3a, 0x30, 0x9a, 0xc5, 0x88, 0x1e, 0x2b, 0x0a, 0x29, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x3d, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x32,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x37, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x3a, 0x2f, 0x9a,
	0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x3a, 0x2a,
	0x9a, 0xc5, 0x88, 0x1e, 0x25, 0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x0b, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x89, 0x01, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33,
	0x42, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33,
	0x3b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_metadata_v3_metadata_proto_rawDescOnce sync.Once
	file_envoy_type_metadata_v3_metadata_proto_rawDescData = file_envoy_type_metadata_v3_metadata_proto_rawDesc
)

func file_envoy_type_metadata_v3_metadata_proto_rawDescGZIP() []byte {
	file_envoy_type_metadata_v3_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_type_metadata_v3_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_metadata_v3_metadata_proto_rawDescData)
	})
	return file_envoy_type_metadata_v3_metadata_proto_rawDescData
}

var file_envoy_type_metadata_v3_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_envoy_type_metadata_v3_metadata_proto_goTypes = []interface{}{
	(*MetadataKey)(nil),             // 0: envoy.type.metadata.v3.MetadataKey
	(*MetadataKind)(nil),            // 1: envoy.type.metadata.v3.MetadataKind
	(*MetadataKey_PathSegment)(nil), // 2: envoy.type.metadata.v3.MetadataKey.PathSegment
	(*MetadataKind_Request)(nil),    // 3: envoy.type.metadata.v3.MetadataKind.Request
	(*MetadataKind_Route)(nil),      // 4: envoy.type.metadata.v3.MetadataKind.Route
	(*MetadataKind_Cluster)(nil),    // 5: envoy.type.metadata.v3.MetadataKind.Cluster
	(*MetadataKind_Host)(nil),       // 6: envoy.type.metadata.v3.MetadataKind.Host
}
var file_envoy_type_metadata_v3_metadata_proto_depIdxs = []int32{
	2, // 0: envoy.type.metadata.v3.MetadataKey.path:type_name -> envoy.type.metadata.v3.MetadataKey.PathSegment
	3, // 1: envoy.type.metadata.v3.MetadataKind.request:type_name -> envoy.type.metadata.v3.MetadataKind.Request
	4, // 2: envoy.type.metadata.v3.MetadataKind.route:type_name -> envoy.type.metadata.v3.MetadataKind.Route
	5, // 3: envoy.type.metadata.v3.MetadataKind.cluster:type_name -> envoy.type.metadata.v3.MetadataKind.Cluster
	6, // 4: envoy.type.metadata.v3.MetadataKind.host:type_name -> envoy.type.metadata.v3.MetadataKind.Host
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_type_metadata_v3_metadata_proto_init() }
func file_envoy_type_metadata_v3_metadata_proto_init() {
	if File_envoy_type_metadata_v3_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKey); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKind); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKey_PathSegment); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKind_Request); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKind_Route); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKind_Cluster); i {
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
		file_envoy_type_metadata_v3_metadata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataKind_Host); i {
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
	file_envoy_type_metadata_v3_metadata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetadataKind_Request_)(nil),
		(*MetadataKind_Route_)(nil),
		(*MetadataKind_Cluster_)(nil),
		(*MetadataKind_Host_)(nil),
	}
	file_envoy_type_metadata_v3_metadata_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MetadataKey_PathSegment_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_metadata_v3_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_metadata_v3_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_type_metadata_v3_metadata_proto_depIdxs,
		MessageInfos:      file_envoy_type_metadata_v3_metadata_proto_msgTypes,
	}.Build()
	File_envoy_type_metadata_v3_metadata_proto = out.File
	file_envoy_type_metadata_v3_metadata_proto_rawDesc = nil
	file_envoy_type_metadata_v3_metadata_proto_goTypes = nil
	file_envoy_type_metadata_v3_metadata_proto_depIdxs = nil
}
