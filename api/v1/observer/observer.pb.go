// Copyright 2019 Authors of Hubble
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
// source: observer/observer.proto

package observer

import (
	flow "github.com/cilium/cilium/api/v1/flow"
	relay "github.com/cilium/cilium/api/v1/relay"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Symbols defined in public import of flow/flow.proto.

type FlowType = flow.FlowType

const FlowType_UNKNOWN_TYPE = flow.FlowType_UNKNOWN_TYPE
const FlowType_L3_L4 = flow.FlowType_L3_L4
const FlowType_L7 = flow.FlowType_L7

var FlowType_name = flow.FlowType_name
var FlowType_value = flow.FlowType_value

type TraceObservationPoint = flow.TraceObservationPoint

const TraceObservationPoint_UNKNOWN_POINT = flow.TraceObservationPoint_UNKNOWN_POINT
const TraceObservationPoint_TO_PROXY = flow.TraceObservationPoint_TO_PROXY
const TraceObservationPoint_TO_HOST = flow.TraceObservationPoint_TO_HOST
const TraceObservationPoint_TO_STACK = flow.TraceObservationPoint_TO_STACK
const TraceObservationPoint_TO_OVERLAY = flow.TraceObservationPoint_TO_OVERLAY
const TraceObservationPoint_TO_ENDPOINT = flow.TraceObservationPoint_TO_ENDPOINT
const TraceObservationPoint_FROM_ENDPOINT = flow.TraceObservationPoint_FROM_ENDPOINT
const TraceObservationPoint_FROM_PROXY = flow.TraceObservationPoint_FROM_PROXY
const TraceObservationPoint_FROM_HOST = flow.TraceObservationPoint_FROM_HOST
const TraceObservationPoint_FROM_STACK = flow.TraceObservationPoint_FROM_STACK
const TraceObservationPoint_FROM_OVERLAY = flow.TraceObservationPoint_FROM_OVERLAY
const TraceObservationPoint_FROM_NETWORK = flow.TraceObservationPoint_FROM_NETWORK
const TraceObservationPoint_TO_NETWORK = flow.TraceObservationPoint_TO_NETWORK

var TraceObservationPoint_name = flow.TraceObservationPoint_name
var TraceObservationPoint_value = flow.TraceObservationPoint_value

type L7FlowType = flow.L7FlowType

const L7FlowType_UNKNOWN_L7_TYPE = flow.L7FlowType_UNKNOWN_L7_TYPE
const L7FlowType_REQUEST = flow.L7FlowType_REQUEST
const L7FlowType_RESPONSE = flow.L7FlowType_RESPONSE
const L7FlowType_SAMPLE = flow.L7FlowType_SAMPLE

var L7FlowType_name = flow.L7FlowType_name
var L7FlowType_value = flow.L7FlowType_value

type IPVersion = flow.IPVersion

const IPVersion_IP_NOT_USED = flow.IPVersion_IP_NOT_USED
const IPVersion_IPv4 = flow.IPVersion_IPv4
const IPVersion_IPv6 = flow.IPVersion_IPv6

var IPVersion_name = flow.IPVersion_name
var IPVersion_value = flow.IPVersion_value

type Verdict = flow.Verdict

const Verdict_VERDICT_UNKNOWN = flow.Verdict_VERDICT_UNKNOWN
const Verdict_FORWARDED = flow.Verdict_FORWARDED
const Verdict_DROPPED = flow.Verdict_DROPPED
const Verdict_ERROR = flow.Verdict_ERROR

var Verdict_name = flow.Verdict_name
var Verdict_value = flow.Verdict_value

type TrafficDirection = flow.TrafficDirection

const TrafficDirection_TRAFFIC_DIRECTION_UNKNOWN = flow.TrafficDirection_TRAFFIC_DIRECTION_UNKNOWN
const TrafficDirection_INGRESS = flow.TrafficDirection_INGRESS
const TrafficDirection_EGRESS = flow.TrafficDirection_EGRESS

var TrafficDirection_name = flow.TrafficDirection_name
var TrafficDirection_value = flow.TrafficDirection_value

type EventType = flow.EventType

const EventType_UNKNOWN = flow.EventType_UNKNOWN
const EventType_EventSample = flow.EventType_EventSample
const EventType_RecordLost = flow.EventType_RecordLost

var EventType_name = flow.EventType_name
var EventType_value = flow.EventType_value

type LostEventSource = flow.LostEventSource

const LostEventSource_UNKNOWN_LOST_EVENT_SOURCE = flow.LostEventSource_UNKNOWN_LOST_EVENT_SOURCE
const LostEventSource_PERF_EVENT_RING_BUFFER = flow.LostEventSource_PERF_EVENT_RING_BUFFER
const LostEventSource_OBSERVER_EVENTS_QUEUE = flow.LostEventSource_OBSERVER_EVENTS_QUEUE

var LostEventSource_name = flow.LostEventSource_name
var LostEventSource_value = flow.LostEventSource_value

type Flow = flow.Flow
type Layer4 = flow.Layer4
type Layer4_TCP = flow.Layer4_TCP
type Layer4_UDP = flow.Layer4_UDP
type Layer4_ICMPv4 = flow.Layer4_ICMPv4
type Layer4_ICMPv6 = flow.Layer4_ICMPv6
type Layer7 = flow.Layer7
type Layer7_Dns = flow.Layer7_Dns
type Layer7_Http = flow.Layer7_Http
type Layer7_Kafka = flow.Layer7_Kafka
type Endpoint = flow.Endpoint
type TCP = flow.TCP
type IP = flow.IP
type Ethernet = flow.Ethernet
type TCPFlags = flow.TCPFlags
type UDP = flow.UDP
type ICMPv4 = flow.ICMPv4
type ICMPv6 = flow.ICMPv6
type EventTypeFilter = flow.EventTypeFilter
type CiliumEventType = flow.CiliumEventType
type FlowFilter = flow.FlowFilter
type Payload = flow.Payload
type DNS = flow.DNS
type HTTPHeader = flow.HTTPHeader
type HTTP = flow.HTTP
type Kafka = flow.Kafka
type Service = flow.Service
type LostEvent = flow.LostEvent

type ServerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerStatusRequest) Reset() {
	*x = ServerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_observer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusRequest) ProtoMessage() {}

func (x *ServerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_observer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusRequest.ProtoReflect.Descriptor instead.
func (*ServerStatusRequest) Descriptor() ([]byte, []int) {
	return file_observer_observer_proto_rawDescGZIP(), []int{0}
}

type ServerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of currently captured flows
	// In a multi-node context, this is the cumulative count of all captured
	// flows.
	NumFlows uint64 `protobuf:"varint,1,opt,name=num_flows,json=numFlows,proto3" json:"num_flows,omitempty"`
	// maximum capacity of the ring buffer
	// In a multi-node context, this is the aggregation of all ring buffers
	// capacities.
	MaxFlows uint64 `protobuf:"varint,2,opt,name=max_flows,json=maxFlows,proto3" json:"max_flows,omitempty"`
	// total amount of flows observed since the observer was started
	// In a multi-node context, this is the aggregation of all flows that have
	// been seen.
	SeenFlows uint64 `protobuf:"varint,3,opt,name=seen_flows,json=seenFlows,proto3" json:"seen_flows,omitempty"`
	// uptime of this observer instance in nanoseconds
	// In a multi-node context, this field corresponds to the uptime of the
	// longest living instance.
	UptimeNs uint64 `protobuf:"varint,4,opt,name=uptime_ns,json=uptimeNs,proto3" json:"uptime_ns,omitempty"`
	// number of nodes for which a connection is established
	NumConnectedNodes *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=num_connected_nodes,json=numConnectedNodes,proto3" json:"num_connected_nodes,omitempty"`
	// number of nodes for which a connection cannot be established
	NumUnavailableNodes *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=num_unavailable_nodes,json=numUnavailableNodes,proto3" json:"num_unavailable_nodes,omitempty"`
	// list of nodes that are unavailable
	// This list may not be exhaustive.
	UnavailableNodes []string `protobuf:"bytes,7,rep,name=unavailable_nodes,json=unavailableNodes,proto3" json:"unavailable_nodes,omitempty"`
}

func (x *ServerStatusResponse) Reset() {
	*x = ServerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_observer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusResponse) ProtoMessage() {}

func (x *ServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_observer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusResponse.ProtoReflect.Descriptor instead.
func (*ServerStatusResponse) Descriptor() ([]byte, []int) {
	return file_observer_observer_proto_rawDescGZIP(), []int{1}
}

func (x *ServerStatusResponse) GetNumFlows() uint64 {
	if x != nil {
		return x.NumFlows
	}
	return 0
}

func (x *ServerStatusResponse) GetMaxFlows() uint64 {
	if x != nil {
		return x.MaxFlows
	}
	return 0
}

func (x *ServerStatusResponse) GetSeenFlows() uint64 {
	if x != nil {
		return x.SeenFlows
	}
	return 0
}

func (x *ServerStatusResponse) GetUptimeNs() uint64 {
	if x != nil {
		return x.UptimeNs
	}
	return 0
}

func (x *ServerStatusResponse) GetNumConnectedNodes() *wrappers.UInt32Value {
	if x != nil {
		return x.NumConnectedNodes
	}
	return nil
}

func (x *ServerStatusResponse) GetNumUnavailableNodes() *wrappers.UInt32Value {
	if x != nil {
		return x.NumUnavailableNodes
	}
	return nil
}

func (x *ServerStatusResponse) GetUnavailableNodes() []string {
	if x != nil {
		return x.UnavailableNodes
	}
	return nil
}

type GetFlowsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of flows that should be returned. Incompatible with `since/until`.
	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// follow sets when the server should continue to stream flows after
	// printing the last N flows.
	Follow bool `protobuf:"varint,3,opt,name=follow,proto3" json:"follow,omitempty"`
	// blacklist defines a list of filters which have to match for a flow to be
	// excluded from the result.
	// If multiple blacklist filters are specified, only one of them has to
	// match for a flow to be excluded.
	Blacklist []*flow.FlowFilter `protobuf:"bytes,5,rep,name=blacklist,proto3" json:"blacklist,omitempty"`
	// whitelist defines a list of filters which have to match for a flow to be
	// included in the result.
	// If multiple whitelist filters are specified, only one of them has to
	// match for a flow to be included.
	// The whitelist and blacklist can both be specified. In such cases, the
	// set of the returned flows is the set difference `whitelist - blacklist`.
	// In other words, the result will contain all flows matched by the
	// whitelist that are not also simultaneously matched by the blacklist.
	Whitelist []*flow.FlowFilter `protobuf:"bytes,6,rep,name=whitelist,proto3" json:"whitelist,omitempty"`
	// Since this time for returned flows. Incompatible with `number`.
	Since *timestamp.Timestamp `protobuf:"bytes,7,opt,name=since,proto3" json:"since,omitempty"`
	// Until this time for returned flows. Incompatible with `number`.
	Until *timestamp.Timestamp `protobuf:"bytes,8,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *GetFlowsRequest) Reset() {
	*x = GetFlowsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_observer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlowsRequest) ProtoMessage() {}

func (x *GetFlowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_observer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlowsRequest.ProtoReflect.Descriptor instead.
func (*GetFlowsRequest) Descriptor() ([]byte, []int) {
	return file_observer_observer_proto_rawDescGZIP(), []int{2}
}

func (x *GetFlowsRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetFlowsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

func (x *GetFlowsRequest) GetBlacklist() []*flow.FlowFilter {
	if x != nil {
		return x.Blacklist
	}
	return nil
}

func (x *GetFlowsRequest) GetWhitelist() []*flow.FlowFilter {
	if x != nil {
		return x.Whitelist
	}
	return nil
}

func (x *GetFlowsRequest) GetSince() *timestamp.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *GetFlowsRequest) GetUntil() *timestamp.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

// GetFlowsResponse contains either a flow or a protocol message.
type GetFlowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResponseTypes:
	//	*GetFlowsResponse_Flow
	//	*GetFlowsResponse_NodeStatus
	//	*GetFlowsResponse_LostEvents
	ResponseTypes isGetFlowsResponse_ResponseTypes `protobuf_oneof:"response_types"`
	// Name of the node where this event was observed.
	NodeName string `protobuf:"bytes,1000,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Timestamp at which this event was observed.
	Time *timestamp.Timestamp `protobuf:"bytes,1001,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetFlowsResponse) Reset() {
	*x = GetFlowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_observer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlowsResponse) ProtoMessage() {}

func (x *GetFlowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_observer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlowsResponse.ProtoReflect.Descriptor instead.
func (*GetFlowsResponse) Descriptor() ([]byte, []int) {
	return file_observer_observer_proto_rawDescGZIP(), []int{3}
}

func (m *GetFlowsResponse) GetResponseTypes() isGetFlowsResponse_ResponseTypes {
	if m != nil {
		return m.ResponseTypes
	}
	return nil
}

func (x *GetFlowsResponse) GetFlow() *flow.Flow {
	if x, ok := x.GetResponseTypes().(*GetFlowsResponse_Flow); ok {
		return x.Flow
	}
	return nil
}

func (x *GetFlowsResponse) GetNodeStatus() *relay.NodeStatusEvent {
	if x, ok := x.GetResponseTypes().(*GetFlowsResponse_NodeStatus); ok {
		return x.NodeStatus
	}
	return nil
}

func (x *GetFlowsResponse) GetLostEvents() *flow.LostEvent {
	if x, ok := x.GetResponseTypes().(*GetFlowsResponse_LostEvents); ok {
		return x.LostEvents
	}
	return nil
}

func (x *GetFlowsResponse) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *GetFlowsResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type isGetFlowsResponse_ResponseTypes interface {
	isGetFlowsResponse_ResponseTypes()
}

type GetFlowsResponse_Flow struct {
	Flow *flow.Flow `protobuf:"bytes,1,opt,name=flow,proto3,oneof"`
}

type GetFlowsResponse_NodeStatus struct {
	// node_status informs clients about the state of the nodes
	// participating in this particular GetFlows request.
	NodeStatus *relay.NodeStatusEvent `protobuf:"bytes,2,opt,name=node_status,json=nodeStatus,proto3,oneof"`
}

type GetFlowsResponse_LostEvents struct {
	// lost_events informs clients about events which got dropped due to
	// a Hubble component being unavailable
	LostEvents *flow.LostEvent `protobuf:"bytes,3,opt,name=lost_events,json=lostEvents,proto3,oneof"`
}

func (*GetFlowsResponse_Flow) isGetFlowsResponse_ResponseTypes() {}

func (*GetFlowsResponse_NodeStatus) isGetFlowsResponse_ResponseTypes() {}

func (*GetFlowsResponse_LostEvents) isGetFlowsResponse_ResponseTypes() {}

var File_observer_observer_proto protoreflect.FileDescriptor

var file_observer_observer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xd9, 0x02, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d,
	0x46, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x46, 0x6c, 0x6f,
	0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x65, 0x6e, 0x46, 0x6c, 0x6f, 0x77,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12, 0x4c,
	0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x15,
	0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6e, 0x75, 0x6d, 0x55, 0x6e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2b,
	0x0a, 0x11, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x75, 0x6e, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x85, 0x02, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x2e, 0x0a, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x22, 0x84, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x39, 0x0a, 0x0b, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0b, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x4c, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6c,
	0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x32, 0xa2, 0x01, 0x0a, 0x08, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6c,
	0x6f, 0x77, 0x73, 0x12, 0x19, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69,
	0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_observer_observer_proto_rawDescOnce sync.Once
	file_observer_observer_proto_rawDescData = file_observer_observer_proto_rawDesc
)

func file_observer_observer_proto_rawDescGZIP() []byte {
	file_observer_observer_proto_rawDescOnce.Do(func() {
		file_observer_observer_proto_rawDescData = protoimpl.X.CompressGZIP(file_observer_observer_proto_rawDescData)
	})
	return file_observer_observer_proto_rawDescData
}

var file_observer_observer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_observer_observer_proto_goTypes = []interface{}{
	(*ServerStatusRequest)(nil),   // 0: observer.ServerStatusRequest
	(*ServerStatusResponse)(nil),  // 1: observer.ServerStatusResponse
	(*GetFlowsRequest)(nil),       // 2: observer.GetFlowsRequest
	(*GetFlowsResponse)(nil),      // 3: observer.GetFlowsResponse
	(*wrappers.UInt32Value)(nil),  // 4: google.protobuf.UInt32Value
	(*flow.FlowFilter)(nil),       // 5: flow.FlowFilter
	(*timestamp.Timestamp)(nil),   // 6: google.protobuf.Timestamp
	(*flow.Flow)(nil),             // 7: flow.Flow
	(*relay.NodeStatusEvent)(nil), // 8: relay.NodeStatusEvent
	(*flow.LostEvent)(nil),        // 9: flow.LostEvent
}
var file_observer_observer_proto_depIdxs = []int32{
	4,  // 0: observer.ServerStatusResponse.num_connected_nodes:type_name -> google.protobuf.UInt32Value
	4,  // 1: observer.ServerStatusResponse.num_unavailable_nodes:type_name -> google.protobuf.UInt32Value
	5,  // 2: observer.GetFlowsRequest.blacklist:type_name -> flow.FlowFilter
	5,  // 3: observer.GetFlowsRequest.whitelist:type_name -> flow.FlowFilter
	6,  // 4: observer.GetFlowsRequest.since:type_name -> google.protobuf.Timestamp
	6,  // 5: observer.GetFlowsRequest.until:type_name -> google.protobuf.Timestamp
	7,  // 6: observer.GetFlowsResponse.flow:type_name -> flow.Flow
	8,  // 7: observer.GetFlowsResponse.node_status:type_name -> relay.NodeStatusEvent
	9,  // 8: observer.GetFlowsResponse.lost_events:type_name -> flow.LostEvent
	6,  // 9: observer.GetFlowsResponse.time:type_name -> google.protobuf.Timestamp
	2,  // 10: observer.Observer.GetFlows:input_type -> observer.GetFlowsRequest
	0,  // 11: observer.Observer.ServerStatus:input_type -> observer.ServerStatusRequest
	3,  // 12: observer.Observer.GetFlows:output_type -> observer.GetFlowsResponse
	1,  // 13: observer.Observer.ServerStatus:output_type -> observer.ServerStatusResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_observer_observer_proto_init() }
func file_observer_observer_proto_init() {
	if File_observer_observer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_observer_observer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusRequest); i {
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
		file_observer_observer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusResponse); i {
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
		file_observer_observer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlowsRequest); i {
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
		file_observer_observer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlowsResponse); i {
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
	file_observer_observer_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetFlowsResponse_Flow)(nil),
		(*GetFlowsResponse_NodeStatus)(nil),
		(*GetFlowsResponse_LostEvents)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_observer_observer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_observer_observer_proto_goTypes,
		DependencyIndexes: file_observer_observer_proto_depIdxs,
		MessageInfos:      file_observer_observer_proto_msgTypes,
	}.Build()
	File_observer_observer_proto = out.File
	file_observer_observer_proto_rawDesc = nil
	file_observer_observer_proto_goTypes = nil
	file_observer_observer_proto_depIdxs = nil
}
