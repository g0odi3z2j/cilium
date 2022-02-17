//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2alpha1

import (
	models "github.com/cilium/cilium/api/v1/models"
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPLoadBalancerIPPool) DeepCopyInto(out *CiliumBGPLoadBalancerIPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPLoadBalancerIPPool.
func (in *CiliumBGPLoadBalancerIPPool) DeepCopy() *CiliumBGPLoadBalancerIPPool {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPLoadBalancerIPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumBGPLoadBalancerIPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPLoadBalancerIPPoolList) DeepCopyInto(out *CiliumBGPLoadBalancerIPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumBGPLoadBalancerIPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPLoadBalancerIPPoolList.
func (in *CiliumBGPLoadBalancerIPPoolList) DeepCopy() *CiliumBGPLoadBalancerIPPoolList {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPLoadBalancerIPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumBGPLoadBalancerIPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPLoadBalancerIPPoolSpec) DeepCopyInto(out *CiliumBGPLoadBalancerIPPoolSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.LBSelector != nil {
		in, out := &in.LBSelector, &out.LBSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPLoadBalancerIPPoolSpec.
func (in *CiliumBGPLoadBalancerIPPoolSpec) DeepCopy() *CiliumBGPLoadBalancerIPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPLoadBalancerIPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPNeighbor) DeepCopyInto(out *CiliumBGPNeighbor) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPNeighbor.
func (in *CiliumBGPNeighbor) DeepCopy() *CiliumBGPNeighbor {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPNeighbor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPPeeringPolicy) DeepCopyInto(out *CiliumBGPPeeringPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPPeeringPolicy.
func (in *CiliumBGPPeeringPolicy) DeepCopy() *CiliumBGPPeeringPolicy {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPPeeringPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumBGPPeeringPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPPeeringPolicyList) DeepCopyInto(out *CiliumBGPPeeringPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumBGPPeeringPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPPeeringPolicyList.
func (in *CiliumBGPPeeringPolicyList) DeepCopy() *CiliumBGPPeeringPolicyList {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPPeeringPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumBGPPeeringPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPPeeringPolicySpec) DeepCopyInto(out *CiliumBGPPeeringPolicySpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualRouters != nil {
		in, out := &in.VirtualRouters, &out.VirtualRouters
		*out = make([]CiliumBGPVirtualRouter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPPeeringPolicySpec.
func (in *CiliumBGPPeeringPolicySpec) DeepCopy() *CiliumBGPPeeringPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPPeeringPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumBGPVirtualRouter) DeepCopyInto(out *CiliumBGPVirtualRouter) {
	*out = *in
	if in.Neighbors != nil {
		in, out := &in.Neighbors, &out.Neighbors
		*out = make([]CiliumBGPNeighbor, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumBGPVirtualRouter.
func (in *CiliumBGPVirtualRouter) DeepCopy() *CiliumBGPVirtualRouter {
	if in == nil {
		return nil
	}
	out := new(CiliumBGPVirtualRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumClusterwideEnvoyConfig) DeepCopyInto(out *CiliumClusterwideEnvoyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumClusterwideEnvoyConfig.
func (in *CiliumClusterwideEnvoyConfig) DeepCopy() *CiliumClusterwideEnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(CiliumClusterwideEnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumClusterwideEnvoyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumClusterwideEnvoyConfigList) DeepCopyInto(out *CiliumClusterwideEnvoyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumClusterwideEnvoyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumClusterwideEnvoyConfigList.
func (in *CiliumClusterwideEnvoyConfigList) DeepCopy() *CiliumClusterwideEnvoyConfigList {
	if in == nil {
		return nil
	}
	out := new(CiliumClusterwideEnvoyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumClusterwideEnvoyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEgressNATPolicy) DeepCopyInto(out *CiliumEgressNATPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEgressNATPolicy.
func (in *CiliumEgressNATPolicy) DeepCopy() *CiliumEgressNATPolicy {
	if in == nil {
		return nil
	}
	out := new(CiliumEgressNATPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEgressNATPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEgressNATPolicyList) DeepCopyInto(out *CiliumEgressNATPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumEgressNATPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEgressNATPolicyList.
func (in *CiliumEgressNATPolicyList) DeepCopy() *CiliumEgressNATPolicyList {
	if in == nil {
		return nil
	}
	out := new(CiliumEgressNATPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEgressNATPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEgressNATPolicySpec) DeepCopyInto(out *CiliumEgressNATPolicySpec) {
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DestinationCIDRs != nil {
		in, out := &in.DestinationCIDRs, &out.DestinationCIDRs
		*out = make([]IPv4CIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEgressNATPolicySpec.
func (in *CiliumEgressNATPolicySpec) DeepCopy() *CiliumEgressNATPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CiliumEgressNATPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEndpointSlice) DeepCopyInto(out *CiliumEndpointSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]CoreCiliumEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEndpointSlice.
func (in *CiliumEndpointSlice) DeepCopy() *CiliumEndpointSlice {
	if in == nil {
		return nil
	}
	out := new(CiliumEndpointSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEndpointSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEndpointSliceList) DeepCopyInto(out *CiliumEndpointSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumEndpointSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEndpointSliceList.
func (in *CiliumEndpointSliceList) DeepCopy() *CiliumEndpointSliceList {
	if in == nil {
		return nil
	}
	out := new(CiliumEndpointSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEndpointSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEnvoyConfig) DeepCopyInto(out *CiliumEnvoyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEnvoyConfig.
func (in *CiliumEnvoyConfig) DeepCopy() *CiliumEnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(CiliumEnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEnvoyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEnvoyConfigList) DeepCopyInto(out *CiliumEnvoyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumEnvoyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEnvoyConfigList.
func (in *CiliumEnvoyConfigList) DeepCopy() *CiliumEnvoyConfigList {
	if in == nil {
		return nil
	}
	out := new(CiliumEnvoyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEnvoyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEnvoyConfigSpec) DeepCopyInto(out *CiliumEnvoyConfigSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]*ServiceListener, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceListener)
				**out = **in
			}
		}
	}
	if in.BackendServices != nil {
		in, out := &in.BackendServices, &out.BackendServices
		*out = make([]*Service, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Service)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]XDSResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEnvoyConfigSpec.
func (in *CiliumEnvoyConfigSpec) DeepCopy() *CiliumEnvoyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CiliumEnvoyConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreCiliumEndpoint) DeepCopyInto(out *CoreCiliumEndpoint) {
	*out = *in
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(v2.EndpointNetworking)
		(*in).DeepCopyInto(*out)
	}
	out.Encryption = in.Encryption
	if in.NamedPorts != nil {
		in, out := &in.NamedPorts, &out.NamedPorts
		*out = make(models.NamedPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(models.Port)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreCiliumEndpoint.
func (in *CoreCiliumEndpoint) DeepCopy() *CoreCiliumEndpoint {
	if in == nil {
		return nil
	}
	out := new(CoreCiliumEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceListener) DeepCopyInto(out *ServiceListener) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceListener.
func (in *ServiceListener) DeepCopy() *ServiceListener {
	if in == nil {
		return nil
	}
	out := new(ServiceListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XDSResource.
func (in *XDSResource) DeepCopy() *XDSResource {
	if in == nil {
		return nil
	}
	out := new(XDSResource)
	in.DeepCopyInto(out)
	return out
}
