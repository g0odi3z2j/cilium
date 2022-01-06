//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by main. DO NOT EDIT.

package v2

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AddressPair) DeepEqual(other *AddressPair) bool {
	if other == nil {
		return false
	}

	if in.IPV4 != other.IPV4 {
		return false
	}
	if in.IPV6 != other.IPV6 {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AddressPairList) DeepEqual(other *AddressPairList) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AllowedIdentityList) DeepEqual(other *AllowedIdentityList) bool {
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

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumClusterwideNetworkPolicy) deepEqual(other *CiliumClusterwideNetworkPolicy) bool {
	if other == nil {
		return false
	}

	if (in.Spec == nil) != (other.Spec == nil) {
		return false
	} else if in.Spec != nil {
		if !in.Spec.DeepEqual(other.Spec) {
			return false
		}
	}

	if ((in.Specs != nil) && (other.Specs != nil)) || ((in.Specs == nil) != (other.Specs == nil)) {
		in, other := &in.Specs, &other.Specs
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumEndpoint) DeepEqual(other *CiliumEndpoint) bool {
	if other == nil {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumExternalWorkload) DeepEqual(other *CiliumExternalWorkload) bool {
	if other == nil {
		return false
	}

	if in.Spec != other.Spec {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumExternalWorkloadSpec) DeepEqual(other *CiliumExternalWorkloadSpec) bool {
	if other == nil {
		return false
	}

	if in.IPv4AllocCIDR != other.IPv4AllocCIDR {
		return false
	}
	if in.IPv6AllocCIDR != other.IPv6AllocCIDR {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumExternalWorkloadStatus) DeepEqual(other *CiliumExternalWorkloadStatus) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.IP != other.IP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumIdentity) DeepEqual(other *CiliumIdentity) bool {
	if other == nil {
		return false
	}

	if ((in.SecurityLabels != nil) && (other.SecurityLabels != nil)) || ((in.SecurityLabels == nil) != (other.SecurityLabels == nil)) {
		in, other := &in.SecurityLabels, &other.SecurityLabels
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumLocalRedirectPolicy) DeepEqual(other *CiliumLocalRedirectPolicy) bool {
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
func (in *CiliumLocalRedirectPolicySpec) DeepEqual(other *CiliumLocalRedirectPolicySpec) bool {
	if other == nil {
		return false
	}

	if !in.RedirectFrontend.DeepEqual(&other.RedirectFrontend) {
		return false
	}

	if !in.RedirectBackend.DeepEqual(&other.RedirectBackend) {
		return false
	}

	if in.Description != other.Description {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumLocalRedirectPolicyStatus) DeepEqual(other *CiliumLocalRedirectPolicyStatus) bool {
	if other == nil {
		return false
	}

	if in.OK != other.OK {
		return false
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumNetworkPolicy) deepEqual(other *CiliumNetworkPolicy) bool {
	if other == nil {
		return false
	}

	if (in.Spec == nil) != (other.Spec == nil) {
		return false
	} else if in.Spec != nil {
		if !in.Spec.DeepEqual(other.Spec) {
			return false
		}
	}

	if ((in.Specs != nil) && (other.Specs != nil)) || ((in.Specs == nil) != (other.Specs == nil)) {
		in, other := &in.Specs, &other.Specs
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumNetworkPolicyNodeStatus) DeepEqual(other *CiliumNetworkPolicyNodeStatus) bool {
	if other == nil {
		return false
	}

	if in.OK != other.OK {
		return false
	}
	if in.Error != other.Error {
		return false
	}
	if !in.LastUpdated.DeepEqual(&other.LastUpdated) {
		return false
	}

	if in.Revision != other.Revision {
		return false
	}
	if in.Enforcing != other.Enforcing {
		return false
	}
	if ((in.Annotations != nil) && (other.Annotations != nil)) || ((in.Annotations == nil) != (other.Annotations == nil)) {
		in, other := &in.Annotations, &other.Annotations
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumNetworkPolicyStatus) DeepEqual(other *CiliumNetworkPolicyStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Nodes != nil) && (other.Nodes != nil)) || ((in.Nodes == nil) != (other.Nodes == nil)) {
		in, other := &in.Nodes, &other.Nodes
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.DerivativePolicies != nil) && (other.DerivativePolicies != nil)) || ((in.DerivativePolicies == nil) != (other.DerivativePolicies == nil)) {
		in, other := &in.DerivativePolicies, &other.DerivativePolicies
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CiliumNode) DeepEqual(other *CiliumNode) bool {
	if other == nil {
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
func (in *ControllerList) DeepEqual(other *ControllerList) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ControllerStatus) DeepEqual(other *ControllerStatus) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if (in.Configuration == nil) != (other.Configuration == nil) {
		return false
	} else if in.Configuration != nil {
		if !in.Configuration.DeepEqual(other.Configuration) {
			return false
		}
	}

	if in.Status != other.Status {
		return false
	}

	if in.UUID != other.UUID {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ControllerStatusStatus) DeepEqual(other *ControllerStatusStatus) bool {
	if other == nil {
		return false
	}

	if in.ConsecutiveFailureCount != other.ConsecutiveFailureCount {
		return false
	}
	if in.FailureCount != other.FailureCount {
		return false
	}
	if in.LastFailureMsg != other.LastFailureMsg {
		return false
	}
	if in.LastFailureTimestamp != other.LastFailureTimestamp {
		return false
	}
	if in.LastSuccessTimestamp != other.LastSuccessTimestamp {
		return false
	}
	if in.SuccessCount != other.SuccessCount {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *DenyIdentityList) DeepEqual(other *DenyIdentityList) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EncryptionSpec) DeepEqual(other *EncryptionSpec) bool {
	if other == nil {
		return false
	}

	if in.Key != other.Key {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointIdentity) DeepEqual(other *EndpointIdentity) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if ((in.Labels != nil) && (other.Labels != nil)) || ((in.Labels == nil) != (other.Labels == nil)) {
		in, other := &in.Labels, &other.Labels
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
func (in *EndpointNetworking) DeepEqual(other *EndpointNetworking) bool {
	if other == nil {
		return false
	}

	if ((in.Addressing != nil) && (other.Addressing != nil)) || ((in.Addressing == nil) != (other.Addressing == nil)) {
		in, other := &in.Addressing, &other.Addressing
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if in.NodeIP != other.NodeIP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointPolicy) DeepEqual(other *EndpointPolicy) bool {
	if other == nil {
		return false
	}

	if (in.Ingress == nil) != (other.Ingress == nil) {
		return false
	} else if in.Ingress != nil {
		if !in.Ingress.DeepEqual(other.Ingress) {
			return false
		}
	}

	if (in.Egress == nil) != (other.Egress == nil) {
		return false
	} else if in.Egress != nil {
		if !in.Egress.DeepEqual(other.Egress) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointPolicyDirection) DeepEqual(other *EndpointPolicyDirection) bool {
	if other == nil {
		return false
	}

	if in.Enforcing != other.Enforcing {
		return false
	}
	if ((in.Allowed != nil) && (other.Allowed != nil)) || ((in.Allowed == nil) != (other.Allowed == nil)) {
		in, other := &in.Allowed, &other.Allowed
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.Denied != nil) && (other.Denied != nil)) || ((in.Denied == nil) != (other.Denied == nil)) {
		in, other := &in.Denied, &other.Denied
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.Removing != nil) && (other.Removing != nil)) || ((in.Removing == nil) != (other.Removing == nil)) {
		in, other := &in.Removing, &other.Removing
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.Adding != nil) && (other.Adding != nil)) || ((in.Adding == nil) != (other.Adding == nil)) {
		in, other := &in.Adding, &other.Adding
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointStatus) DeepEqual(other *EndpointStatus) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if ((in.Controllers != nil) && (other.Controllers != nil)) || ((in.Controllers == nil) != (other.Controllers == nil)) {
		in, other := &in.Controllers, &other.Controllers
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if (in.ExternalIdentifiers == nil) != (other.ExternalIdentifiers == nil) {
		return false
	} else if in.ExternalIdentifiers != nil {
		if !in.ExternalIdentifiers.DeepEqual(other.ExternalIdentifiers) {
			return false
		}
	}

	if (in.Health == nil) != (other.Health == nil) {
		return false
	} else if in.Health != nil {
		if !in.Health.DeepEqual(other.Health) {
			return false
		}
	}

	if (in.Identity == nil) != (other.Identity == nil) {
		return false
	} else if in.Identity != nil {
		if !in.Identity.DeepEqual(other.Identity) {
			return false
		}
	}

	if ((in.Log != nil) && (other.Log != nil)) || ((in.Log == nil) != (other.Log == nil)) {
		in, other := &in.Log, &other.Log
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

	if (in.Policy == nil) != (other.Policy == nil) {
		return false
	} else if in.Policy != nil {
		if !in.Policy.DeepEqual(other.Policy) {
			return false
		}
	}

	if (in.VisibilityPolicyStatus == nil) != (other.VisibilityPolicyStatus == nil) {
		return false
	} else if in.VisibilityPolicyStatus != nil {
		if *in.VisibilityPolicyStatus != *other.VisibilityPolicyStatus {
			return false
		}
	}

	if in.State != other.State {
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
func (in *Frontend) DeepEqual(other *Frontend) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}
	if ((in.ToPorts != nil) && (other.ToPorts != nil)) || ((in.ToPorts == nil) != (other.ToPorts == nil)) {
		in, other := &in.ToPorts, &other.ToPorts
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
func (in *HealthAddressingSpec) DeepEqual(other *HealthAddressingSpec) bool {
	if other == nil {
		return false
	}

	if in.IPv4 != other.IPv4 {
		return false
	}
	if in.IPv6 != other.IPv6 {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IdentityList) DeepEqual(other *IdentityList) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IdentityTuple) DeepEqual(other *IdentityTuple) bool {
	if other == nil {
		return false
	}

	if in.Identity != other.Identity {
		return false
	}
	if ((in.IdentityLabels != nil) && (other.IdentityLabels != nil)) || ((in.IdentityLabels == nil) != (other.IdentityLabels == nil)) {
		in, other := &in.IdentityLabels, &other.IdentityLabels
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if in.DestPort != other.DestPort {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeAddress) DeepEqual(other *NodeAddress) bool {
	if other == nil {
		return false
	}

	if in.Type != other.Type {
		return false
	}
	if in.IP != other.IP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeSpec) DeepEqual(other *NodeSpec) bool {
	if other == nil {
		return false
	}

	if in.InstanceID != other.InstanceID {
		return false
	}
	if ((in.Addresses != nil) && (other.Addresses != nil)) || ((in.Addresses == nil) != (other.Addresses == nil)) {
		in, other := &in.Addresses, &other.Addresses
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

	if in.HealthAddressing != other.HealthAddressing {
		return false
	}

	if in.Encryption != other.Encryption {
		return false
	}

	if !in.ENI.DeepEqual(&other.ENI) {
		return false
	}

	if in.Azure != other.Azure {
		return false
	}

	if !in.AlibabaCloud.DeepEqual(&other.AlibabaCloud) {
		return false
	}

	if !in.IPAM.DeepEqual(&other.IPAM) {
		return false
	}

	if in.NodeIdentity != other.NodeIdentity {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeStatus) DeepEqual(other *NodeStatus) bool {
	if other == nil {
		return false
	}

	if !in.ENI.DeepEqual(&other.ENI) {
		return false
	}

	if !in.Azure.DeepEqual(&other.Azure) {
		return false
	}

	if !in.IPAM.DeepEqual(&other.IPAM) {
		return false
	}

	if !in.AlibabaCloud.DeepEqual(&other.AlibabaCloud) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortInfo) DeepEqual(other *PortInfo) bool {
	if other == nil {
		return false
	}

	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}
	if in.Name != other.Name {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *RedirectBackend) DeepEqual(other *RedirectBackend) bool {
	if other == nil {
		return false
	}

	if !in.LocalEndpointSelector.DeepEqual(&other.LocalEndpointSelector) {
		return false
	}

	if ((in.ToPorts != nil) && (other.ToPorts != nil)) || ((in.ToPorts == nil) != (other.ToPorts == nil)) {
		in, other := &in.ToPorts, &other.ToPorts
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
func (in *RedirectFrontend) DeepEqual(other *RedirectFrontend) bool {
	if other == nil {
		return false
	}

	if (in.AddressMatcher == nil) != (other.AddressMatcher == nil) {
		return false
	} else if in.AddressMatcher != nil {
		if !in.AddressMatcher.DeepEqual(other.AddressMatcher) {
			return false
		}
	}

	if (in.ServiceMatcher == nil) != (other.ServiceMatcher == nil) {
		return false
	} else if in.ServiceMatcher != nil {
		if !in.ServiceMatcher.DeepEqual(other.ServiceMatcher) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceInfo) DeepEqual(other *ServiceInfo) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Namespace != other.Namespace {
		return false
	}
	if ((in.ToPorts != nil) && (other.ToPorts != nil)) || ((in.ToPorts == nil) != (other.ToPorts == nil)) {
		in, other := &in.ToPorts, &other.ToPorts
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
