//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by main. DO NOT EDIT.

package v2alpha1

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumClusterwideEnvoyConfig) DeepEqual(other *CiliumClusterwideEnvoyConfig) bool {
	if other == nil {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumEgressNATPolicy) DeepEqual(other *CiliumEgressNATPolicy) bool {
	if other == nil {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumEgressNATPolicySpec) DeepEqual(other *CiliumEgressNATPolicySpec) bool {
	if other == nil {
		return false
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

	if ((in.DestinationCIDRs != nil) && (other.DestinationCIDRs != nil)) || ((in.DestinationCIDRs == nil) != (other.DestinationCIDRs == nil)) {
		in, other := &in.DestinationCIDRs, &other.DestinationCIDRs
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

	if in.EgressSourceIP != other.EgressSourceIP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumEndpointSlice) DeepEqual(other *CiliumEndpointSlice) bool {
	if other == nil {
		return false
	}

	if in.Namespace != other.Namespace {
		return false
	}
	if ((in.Endpoints != nil) && (other.Endpoints != nil)) || ((in.Endpoints == nil) != (other.Endpoints == nil)) {
		in, other := &in.Endpoints, &other.Endpoints
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
func (in *CiliumEnvoyConfig) DeepEqual(other *CiliumEnvoyConfig) bool {
	if other == nil {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumEnvoyConfigSpec) DeepEqual(other *CiliumEnvoyConfigSpec) bool {
	if other == nil {
		return false
	}

	if ((in.Services != nil) && (other.Services != nil)) || ((in.Services == nil) != (other.Services == nil)) {
		in, other := &in.Services, &other.Services
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual((*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.BackendServices != nil) && (other.BackendServices != nil)) || ((in.BackendServices == nil) != (other.BackendServices == nil)) {
		in, other := &in.BackendServices, &other.BackendServices
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual((*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Resources != nil) && (other.Resources != nil)) || ((in.Resources == nil) != (other.Resources == nil)) {
		in, other := &in.Resources, &other.Resources
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
func (in *CoreCiliumEndpoint) DeepEqual(other *CoreCiliumEndpoint) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.IdentityID != other.IdentityID {
		return false
	}
	if (in.Networking == nil) != (other.Networking == nil) {
		return false
	} else if in.Networking != nil {
		if !in.Networking.DeepEqual(other.Networking) {
			return false
		}
	}

	if in.Encryption != other.Encryption {
		return false
	}

	if ((in.NamedPorts != nil) && (other.NamedPorts != nil)) || ((in.NamedPorts == nil) != (other.NamedPorts == nil)) {
		in, other := &in.NamedPorts, &other.NamedPorts
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual((*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EgressRule) DeepEqual(other *EgressRule) bool {
	if other == nil {
		return false
	}

	if (in.NamespaceSelector == nil) != (other.NamespaceSelector == nil) {
		return false
	} else if in.NamespaceSelector != nil {
		if !in.NamespaceSelector.DeepEqual(other.NamespaceSelector) {
			return false
		}
	}

	if (in.PodSelector == nil) != (other.PodSelector == nil) {
		return false
	} else if in.PodSelector != nil {
		if !in.PodSelector.DeepEqual(other.PodSelector) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Service) DeepEqual(other *Service) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Namespace != other.Namespace {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceListener) DeepEqual(other *ServiceListener) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Namespace != other.Namespace {
		return false
	}
	if in.Listener != other.Listener {
		return false
	}

	return true
}
