// +build !ignore_autogenerated

// Copyright 2017-2021 Authors of Cilium
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

// Code generated by main. DO NOT EDIT.

package v1

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
