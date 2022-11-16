//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package v1

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *HTTPIngressPath) DeepEqual(other *HTTPIngressPath) bool {
	if other == nil {
		return false
	}

	if in.Path != other.Path {
		return false
	}
	if (in.PathType == nil) != (other.PathType == nil) {
		return false
	} else if in.PathType != nil {
		if *in.PathType != *other.PathType {
			return false
		}
	}

	if !in.Backend.DeepEqual(&other.Backend) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *HTTPIngressRuleValue) DeepEqual(other *HTTPIngressRuleValue) bool {
	if other == nil {
		return false
	}

	if ((in.Paths != nil) && (other.Paths != nil)) || ((in.Paths == nil) != (other.Paths == nil)) {
		in, other := &in.Paths, &other.Paths
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IPBlock) DeepEqual(other *IPBlock) bool {
	if other == nil {
		return false
	}

	if in.CIDR != other.CIDR {
		return false
	}
	if ((in.Except != nil) && (other.Except != nil)) || ((in.Except == nil) != (other.Except == nil)) {
		in, other := &in.Except, &other.Except
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Ingress) DeepEqual(other *Ingress) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressBackend) DeepEqual(other *IngressBackend) bool {
	if other == nil {
		return false
	}

	if (in.Service == nil) != (other.Service == nil) {
		return false
	} else if in.Service != nil {
		if !in.Service.DeepEqual(other.Service) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressList) DeepEqual(other *IngressList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressLoadBalancerIngress) DeepEqual(other *IngressLoadBalancerIngress) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}
	if in.Hostname != other.Hostname {
		return false
	}
	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressLoadBalancerStatus) DeepEqual(other *IngressLoadBalancerStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Ingress != nil) && (other.Ingress != nil)) || ((in.Ingress == nil) != (other.Ingress == nil)) {
		in, other := &in.Ingress, &other.Ingress
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressPortStatus) DeepEqual(other *IngressPortStatus) bool {
	if other == nil {
		return false
	}

	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}
	if (in.Error == nil) != (other.Error == nil) {
		return false
	} else if in.Error != nil {
		if *in.Error != *other.Error {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressRule) DeepEqual(other *IngressRule) bool {
	if other == nil {
		return false
	}

	if in.Host != other.Host {
		return false
	}
	if !in.IngressRuleValue.DeepEqual(&other.IngressRuleValue) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressRuleValue) DeepEqual(other *IngressRuleValue) bool {
	if other == nil {
		return false
	}

	if (in.HTTP == nil) != (other.HTTP == nil) {
		return false
	} else if in.HTTP != nil {
		if !in.HTTP.DeepEqual(other.HTTP) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressServiceBackend) DeepEqual(other *IngressServiceBackend) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Port != other.Port {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressSpec) DeepEqual(other *IngressSpec) bool {
	if other == nil {
		return false
	}

	if (in.IngressClassName == nil) != (other.IngressClassName == nil) {
		return false
	} else if in.IngressClassName != nil {
		if *in.IngressClassName != *other.IngressClassName {
			return false
		}
	}

	if (in.DefaultBackend == nil) != (other.DefaultBackend == nil) {
		return false
	} else if in.DefaultBackend != nil {
		if !in.DefaultBackend.DeepEqual(other.DefaultBackend) {
			return false
		}
	}

	if ((in.TLS != nil) && (other.TLS != nil)) || ((in.TLS == nil) != (other.TLS == nil)) {
		in, other := &in.TLS, &other.TLS
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Rules != nil) && (other.Rules != nil)) || ((in.Rules == nil) != (other.Rules == nil)) {
		in, other := &in.Rules, &other.Rules
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressStatus) DeepEqual(other *IngressStatus) bool {
	if other == nil {
		return false
	}

	if !in.LoadBalancer.DeepEqual(&other.LoadBalancer) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressTLS) DeepEqual(other *IngressTLS) bool {
	if other == nil {
		return false
	}

	if ((in.Hosts != nil) && (other.Hosts != nil)) || ((in.Hosts == nil) != (other.Hosts == nil)) {
		in, other := &in.Hosts, &other.Hosts
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if in.SecretName != other.SecretName {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicy) DeepEqual(other *NetworkPolicy) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicyEgressRule) DeepEqual(other *NetworkPolicyEgressRule) bool {
	if other == nil {
		return false
	}

	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.To != nil) && (other.To != nil)) || ((in.To == nil) != (other.To == nil)) {
		in, other := &in.To, &other.To
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicyIngressRule) DeepEqual(other *NetworkPolicyIngressRule) bool {
	if other == nil {
		return false
	}

	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.From != nil) && (other.From != nil)) || ((in.From == nil) != (other.From == nil)) {
		in, other := &in.From, &other.From
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicyList) DeepEqual(other *NetworkPolicyList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicyPeer) DeepEqual(other *NetworkPolicyPeer) bool {
	if other == nil {
		return false
	}

	if (in.PodSelector == nil) != (other.PodSelector == nil) {
		return false
	} else if in.PodSelector != nil {
		if !in.PodSelector.DeepEqual(other.PodSelector) {
			return false
		}
	}

	if (in.NamespaceSelector == nil) != (other.NamespaceSelector == nil) {
		return false
	} else if in.NamespaceSelector != nil {
		if !in.NamespaceSelector.DeepEqual(other.NamespaceSelector) {
			return false
		}
	}

	if (in.IPBlock == nil) != (other.IPBlock == nil) {
		return false
	} else if in.IPBlock != nil {
		if !in.IPBlock.DeepEqual(other.IPBlock) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicyPort) DeepEqual(other *NetworkPolicyPort) bool {
	if other == nil {
		return false
	}

	if (in.Protocol == nil) != (other.Protocol == nil) {
		return false
	} else if in.Protocol != nil {
		if *in.Protocol != *other.Protocol {
			return false
		}
	}

	if (in.Port == nil) != (other.Port == nil) {
		return false
	} else if in.Port != nil {
		if !in.Port.DeepEqual(other.Port) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NetworkPolicySpec) DeepEqual(other *NetworkPolicySpec) bool {
	if other == nil {
		return false
	}

	if !in.PodSelector.DeepEqual(&other.PodSelector) {
		return false
	}

	if ((in.Ingress != nil) && (other.Ingress != nil)) || ((in.Ingress == nil) != (other.Ingress == nil)) {
		in, other := &in.Ingress, &other.Ingress
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Egress != nil) && (other.Egress != nil)) || ((in.Egress == nil) != (other.Egress == nil)) {
		in, other := &in.Egress, &other.Egress
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.PolicyTypes != nil) && (other.PolicyTypes != nil)) || ((in.PolicyTypes == nil) != (other.PolicyTypes == nil)) {
		in, other := &in.PolicyTypes, &other.PolicyTypes
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceBackendPort) DeepEqual(other *ServiceBackendPort) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Number != other.Number {
		return false
	}

	return true
}
