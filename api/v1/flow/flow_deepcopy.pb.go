// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package flow

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Flow within kubernetes types, where deepcopy-gen is used.
func (in *Flow) DeepCopyInto(out *Flow) {
	p := proto.Clone(in).(*Flow)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow. Required by controller-gen.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Flow. Required by controller-gen.
func (in *Flow) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Layer4 within kubernetes types, where deepcopy-gen is used.
func (in *Layer4) DeepCopyInto(out *Layer4) {
	p := proto.Clone(in).(*Layer4)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer4. Required by controller-gen.
func (in *Layer4) DeepCopy() *Layer4 {
	if in == nil {
		return nil
	}
	out := new(Layer4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Layer4. Required by controller-gen.
func (in *Layer4) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Layer7 within kubernetes types, where deepcopy-gen is used.
func (in *Layer7) DeepCopyInto(out *Layer7) {
	p := proto.Clone(in).(*Layer7)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer7. Required by controller-gen.
func (in *Layer7) DeepCopy() *Layer7 {
	if in == nil {
		return nil
	}
	out := new(Layer7)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Layer7. Required by controller-gen.
func (in *Layer7) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TraceContext within kubernetes types, where deepcopy-gen is used.
func (in *TraceContext) DeepCopyInto(out *TraceContext) {
	p := proto.Clone(in).(*TraceContext)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceContext. Required by controller-gen.
func (in *TraceContext) DeepCopy() *TraceContext {
	if in == nil {
		return nil
	}
	out := new(TraceContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TraceContext. Required by controller-gen.
func (in *TraceContext) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TraceParent within kubernetes types, where deepcopy-gen is used.
func (in *TraceParent) DeepCopyInto(out *TraceParent) {
	p := proto.Clone(in).(*TraceParent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceParent. Required by controller-gen.
func (in *TraceParent) DeepCopy() *TraceParent {
	if in == nil {
		return nil
	}
	out := new(TraceParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TraceParent. Required by controller-gen.
func (in *TraceParent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Endpoint within kubernetes types, where deepcopy-gen is used.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	p := proto.Clone(in).(*Endpoint)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint. Required by controller-gen.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint. Required by controller-gen.
func (in *Endpoint) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Workload within kubernetes types, where deepcopy-gen is used.
func (in *Workload) DeepCopyInto(out *Workload) {
	p := proto.Clone(in).(*Workload)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload. Required by controller-gen.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Workload. Required by controller-gen.
func (in *Workload) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCP within kubernetes types, where deepcopy-gen is used.
func (in *TCP) DeepCopyInto(out *TCP) {
	p := proto.Clone(in).(*TCP)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCP. Required by controller-gen.
func (in *TCP) DeepCopy() *TCP {
	if in == nil {
		return nil
	}
	out := new(TCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCP. Required by controller-gen.
func (in *TCP) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using IP within kubernetes types, where deepcopy-gen is used.
func (in *IP) DeepCopyInto(out *IP) {
	p := proto.Clone(in).(*IP)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IP. Required by controller-gen.
func (in *IP) DeepCopy() *IP {
	if in == nil {
		return nil
	}
	out := new(IP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new IP. Required by controller-gen.
func (in *IP) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Ethernet within kubernetes types, where deepcopy-gen is used.
func (in *Ethernet) DeepCopyInto(out *Ethernet) {
	p := proto.Clone(in).(*Ethernet)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ethernet. Required by controller-gen.
func (in *Ethernet) DeepCopy() *Ethernet {
	if in == nil {
		return nil
	}
	out := new(Ethernet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Ethernet. Required by controller-gen.
func (in *Ethernet) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCPFlags within kubernetes types, where deepcopy-gen is used.
func (in *TCPFlags) DeepCopyInto(out *TCPFlags) {
	p := proto.Clone(in).(*TCPFlags)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPFlags. Required by controller-gen.
func (in *TCPFlags) DeepCopy() *TCPFlags {
	if in == nil {
		return nil
	}
	out := new(TCPFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCPFlags. Required by controller-gen.
func (in *TCPFlags) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using UDP within kubernetes types, where deepcopy-gen is used.
func (in *UDP) DeepCopyInto(out *UDP) {
	p := proto.Clone(in).(*UDP)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDP. Required by controller-gen.
func (in *UDP) DeepCopy() *UDP {
	if in == nil {
		return nil
	}
	out := new(UDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new UDP. Required by controller-gen.
func (in *UDP) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SCTP within kubernetes types, where deepcopy-gen is used.
func (in *SCTP) DeepCopyInto(out *SCTP) {
	p := proto.Clone(in).(*SCTP)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCTP. Required by controller-gen.
func (in *SCTP) DeepCopy() *SCTP {
	if in == nil {
		return nil
	}
	out := new(SCTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SCTP. Required by controller-gen.
func (in *SCTP) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ICMPv4 within kubernetes types, where deepcopy-gen is used.
func (in *ICMPv4) DeepCopyInto(out *ICMPv4) {
	p := proto.Clone(in).(*ICMPv4)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPv4. Required by controller-gen.
func (in *ICMPv4) DeepCopy() *ICMPv4 {
	if in == nil {
		return nil
	}
	out := new(ICMPv4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ICMPv4. Required by controller-gen.
func (in *ICMPv4) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ICMPv6 within kubernetes types, where deepcopy-gen is used.
func (in *ICMPv6) DeepCopyInto(out *ICMPv6) {
	p := proto.Clone(in).(*ICMPv6)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPv6. Required by controller-gen.
func (in *ICMPv6) DeepCopy() *ICMPv6 {
	if in == nil {
		return nil
	}
	out := new(ICMPv6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ICMPv6. Required by controller-gen.
func (in *ICMPv6) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Policy within kubernetes types, where deepcopy-gen is used.
func (in *Policy) DeepCopyInto(out *Policy) {
	p := proto.Clone(in).(*Policy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy. Required by controller-gen.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Policy. Required by controller-gen.
func (in *Policy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EventTypeFilter within kubernetes types, where deepcopy-gen is used.
func (in *EventTypeFilter) DeepCopyInto(out *EventTypeFilter) {
	p := proto.Clone(in).(*EventTypeFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeFilter. Required by controller-gen.
func (in *EventTypeFilter) DeepCopy() *EventTypeFilter {
	if in == nil {
		return nil
	}
	out := new(EventTypeFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeFilter. Required by controller-gen.
func (in *EventTypeFilter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CiliumEventType within kubernetes types, where deepcopy-gen is used.
func (in *CiliumEventType) DeepCopyInto(out *CiliumEventType) {
	p := proto.Clone(in).(*CiliumEventType)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEventType. Required by controller-gen.
func (in *CiliumEventType) DeepCopy() *CiliumEventType {
	if in == nil {
		return nil
	}
	out := new(CiliumEventType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEventType. Required by controller-gen.
func (in *CiliumEventType) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FlowFilter within kubernetes types, where deepcopy-gen is used.
func (in *FlowFilter) DeepCopyInto(out *FlowFilter) {
	p := proto.Clone(in).(*FlowFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowFilter. Required by controller-gen.
func (in *FlowFilter) DeepCopy() *FlowFilter {
	if in == nil {
		return nil
	}
	out := new(FlowFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FlowFilter. Required by controller-gen.
func (in *FlowFilter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DNS within kubernetes types, where deepcopy-gen is used.
func (in *DNS) DeepCopyInto(out *DNS) {
	p := proto.Clone(in).(*DNS)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNS. Required by controller-gen.
func (in *DNS) DeepCopy() *DNS {
	if in == nil {
		return nil
	}
	out := new(DNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DNS. Required by controller-gen.
func (in *DNS) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPHeader within kubernetes types, where deepcopy-gen is used.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	p := proto.Clone(in).(*HTTPHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader. Required by controller-gen.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader. Required by controller-gen.
func (in *HTTPHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTP within kubernetes types, where deepcopy-gen is used.
func (in *HTTP) DeepCopyInto(out *HTTP) {
	p := proto.Clone(in).(*HTTP)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP. Required by controller-gen.
func (in *HTTP) DeepCopy() *HTTP {
	if in == nil {
		return nil
	}
	out := new(HTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTP. Required by controller-gen.
func (in *HTTP) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Kafka within kubernetes types, where deepcopy-gen is used.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	p := proto.Clone(in).(*Kafka)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka. Required by controller-gen.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Kafka. Required by controller-gen.
func (in *Kafka) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Service within kubernetes types, where deepcopy-gen is used.
func (in *Service) DeepCopyInto(out *Service) {
	p := proto.Clone(in).(*Service)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LostEvent within kubernetes types, where deepcopy-gen is used.
func (in *LostEvent) DeepCopyInto(out *LostEvent) {
	p := proto.Clone(in).(*LostEvent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LostEvent. Required by controller-gen.
func (in *LostEvent) DeepCopy() *LostEvent {
	if in == nil {
		return nil
	}
	out := new(LostEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LostEvent. Required by controller-gen.
func (in *LostEvent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AgentEvent within kubernetes types, where deepcopy-gen is used.
func (in *AgentEvent) DeepCopyInto(out *AgentEvent) {
	p := proto.Clone(in).(*AgentEvent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentEvent. Required by controller-gen.
func (in *AgentEvent) DeepCopy() *AgentEvent {
	if in == nil {
		return nil
	}
	out := new(AgentEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AgentEvent. Required by controller-gen.
func (in *AgentEvent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AgentEventUnknown within kubernetes types, where deepcopy-gen is used.
func (in *AgentEventUnknown) DeepCopyInto(out *AgentEventUnknown) {
	p := proto.Clone(in).(*AgentEventUnknown)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentEventUnknown. Required by controller-gen.
func (in *AgentEventUnknown) DeepCopy() *AgentEventUnknown {
	if in == nil {
		return nil
	}
	out := new(AgentEventUnknown)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AgentEventUnknown. Required by controller-gen.
func (in *AgentEventUnknown) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TimeNotification within kubernetes types, where deepcopy-gen is used.
func (in *TimeNotification) DeepCopyInto(out *TimeNotification) {
	p := proto.Clone(in).(*TimeNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeNotification. Required by controller-gen.
func (in *TimeNotification) DeepCopy() *TimeNotification {
	if in == nil {
		return nil
	}
	out := new(TimeNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TimeNotification. Required by controller-gen.
func (in *TimeNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PolicyUpdateNotification within kubernetes types, where deepcopy-gen is used.
func (in *PolicyUpdateNotification) DeepCopyInto(out *PolicyUpdateNotification) {
	p := proto.Clone(in).(*PolicyUpdateNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyUpdateNotification. Required by controller-gen.
func (in *PolicyUpdateNotification) DeepCopy() *PolicyUpdateNotification {
	if in == nil {
		return nil
	}
	out := new(PolicyUpdateNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PolicyUpdateNotification. Required by controller-gen.
func (in *PolicyUpdateNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EndpointRegenNotification within kubernetes types, where deepcopy-gen is used.
func (in *EndpointRegenNotification) DeepCopyInto(out *EndpointRegenNotification) {
	p := proto.Clone(in).(*EndpointRegenNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointRegenNotification. Required by controller-gen.
func (in *EndpointRegenNotification) DeepCopy() *EndpointRegenNotification {
	if in == nil {
		return nil
	}
	out := new(EndpointRegenNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EndpointRegenNotification. Required by controller-gen.
func (in *EndpointRegenNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EndpointUpdateNotification within kubernetes types, where deepcopy-gen is used.
func (in *EndpointUpdateNotification) DeepCopyInto(out *EndpointUpdateNotification) {
	p := proto.Clone(in).(*EndpointUpdateNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointUpdateNotification. Required by controller-gen.
func (in *EndpointUpdateNotification) DeepCopy() *EndpointUpdateNotification {
	if in == nil {
		return nil
	}
	out := new(EndpointUpdateNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EndpointUpdateNotification. Required by controller-gen.
func (in *EndpointUpdateNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using IPCacheNotification within kubernetes types, where deepcopy-gen is used.
func (in *IPCacheNotification) DeepCopyInto(out *IPCacheNotification) {
	p := proto.Clone(in).(*IPCacheNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPCacheNotification. Required by controller-gen.
func (in *IPCacheNotification) DeepCopy() *IPCacheNotification {
	if in == nil {
		return nil
	}
	out := new(IPCacheNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new IPCacheNotification. Required by controller-gen.
func (in *IPCacheNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceUpsertNotificationAddr within kubernetes types, where deepcopy-gen is used.
func (in *ServiceUpsertNotificationAddr) DeepCopyInto(out *ServiceUpsertNotificationAddr) {
	p := proto.Clone(in).(*ServiceUpsertNotificationAddr)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceUpsertNotificationAddr. Required by controller-gen.
func (in *ServiceUpsertNotificationAddr) DeepCopy() *ServiceUpsertNotificationAddr {
	if in == nil {
		return nil
	}
	out := new(ServiceUpsertNotificationAddr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceUpsertNotificationAddr. Required by controller-gen.
func (in *ServiceUpsertNotificationAddr) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceUpsertNotification within kubernetes types, where deepcopy-gen is used.
func (in *ServiceUpsertNotification) DeepCopyInto(out *ServiceUpsertNotification) {
	p := proto.Clone(in).(*ServiceUpsertNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceUpsertNotification. Required by controller-gen.
func (in *ServiceUpsertNotification) DeepCopy() *ServiceUpsertNotification {
	if in == nil {
		return nil
	}
	out := new(ServiceUpsertNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceUpsertNotification. Required by controller-gen.
func (in *ServiceUpsertNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceDeleteNotification within kubernetes types, where deepcopy-gen is used.
func (in *ServiceDeleteNotification) DeepCopyInto(out *ServiceDeleteNotification) {
	p := proto.Clone(in).(*ServiceDeleteNotification)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDeleteNotification. Required by controller-gen.
func (in *ServiceDeleteNotification) DeepCopy() *ServiceDeleteNotification {
	if in == nil {
		return nil
	}
	out := new(ServiceDeleteNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDeleteNotification. Required by controller-gen.
func (in *ServiceDeleteNotification) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using NetworkInterface within kubernetes types, where deepcopy-gen is used.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	p := proto.Clone(in).(*NetworkInterface)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface. Required by controller-gen.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface. Required by controller-gen.
func (in *NetworkInterface) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DebugEvent within kubernetes types, where deepcopy-gen is used.
func (in *DebugEvent) DeepCopyInto(out *DebugEvent) {
	p := proto.Clone(in).(*DebugEvent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebugEvent. Required by controller-gen.
func (in *DebugEvent) DeepCopy() *DebugEvent {
	if in == nil {
		return nil
	}
	out := new(DebugEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DebugEvent. Required by controller-gen.
func (in *DebugEvent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
