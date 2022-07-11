//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package api

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AWSGroup) DeepEqual(other *AWSGroup) bool {
	if other == nil {
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

	if ((in.SecurityGroupsIds != nil) && (other.SecurityGroupsIds != nil)) || ((in.SecurityGroupsIds == nil) != (other.SecurityGroupsIds == nil)) {
		in, other := &in.SecurityGroupsIds, &other.SecurityGroupsIds
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

	if ((in.SecurityGroupsNames != nil) && (other.SecurityGroupsNames != nil)) || ((in.SecurityGroupsNames == nil) != (other.SecurityGroupsNames == nil)) {
		in, other := &in.SecurityGroupsNames, &other.SecurityGroupsNames
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

	if in.Region != other.Region {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CIDRRule) DeepEqual(other *CIDRRule) bool {
	if other == nil {
		return false
	}

	if in.Cidr != other.Cidr {
		return false
	}
	if ((in.ExceptCIDRs != nil) && (other.ExceptCIDRs != nil)) || ((in.ExceptCIDRs == nil) != (other.ExceptCIDRs == nil)) {
		in, other := &in.ExceptCIDRs, &other.ExceptCIDRs
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

	if in.Generated != other.Generated {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *CIDRRuleSlice) DeepEqual(other *CIDRRuleSlice) bool {
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
func (in *CIDRSlice) DeepEqual(other *CIDRSlice) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EgressCommonRule) DeepEqual(other *EgressCommonRule) bool {
	if other == nil {
		return false
	}

	if ((in.ToEndpoints != nil) && (other.ToEndpoints != nil)) || ((in.ToEndpoints == nil) != (other.ToEndpoints == nil)) {
		in, other := &in.ToEndpoints, &other.ToEndpoints
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

	if ((in.ToRequires != nil) && (other.ToRequires != nil)) || ((in.ToRequires == nil) != (other.ToRequires == nil)) {
		in, other := &in.ToRequires, &other.ToRequires
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

	if ((in.ToCIDR != nil) && (other.ToCIDR != nil)) || ((in.ToCIDR == nil) != (other.ToCIDR == nil)) {
		in, other := &in.ToCIDR, &other.ToCIDR
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.ToCIDRSet != nil) && (other.ToCIDRSet != nil)) || ((in.ToCIDRSet == nil) != (other.ToCIDRSet == nil)) {
		in, other := &in.ToCIDRSet, &other.ToCIDRSet
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.ToEntities != nil) && (other.ToEntities != nil)) || ((in.ToEntities == nil) != (other.ToEntities == nil)) {
		in, other := &in.ToEntities, &other.ToEntities
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

	if ((in.ToServices != nil) && (other.ToServices != nil)) || ((in.ToServices == nil) != (other.ToServices == nil)) {
		in, other := &in.ToServices, &other.ToServices
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

	if ((in.ToGroups != nil) && (other.ToGroups != nil)) || ((in.ToGroups == nil) != (other.ToGroups == nil)) {
		in, other := &in.ToGroups, &other.ToGroups
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

	if ((in.aggregatedSelectors != nil) && (other.aggregatedSelectors != nil)) || ((in.aggregatedSelectors == nil) != (other.aggregatedSelectors == nil)) {
		in, other := &in.aggregatedSelectors, &other.aggregatedSelectors
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
func (in *EgressDenyRule) DeepEqual(other *EgressDenyRule) bool {
	if other == nil {
		return false
	}

	if !in.EgressCommonRule.DeepEqual(&other.EgressCommonRule) {
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

	if ((in.ICMPs != nil) && (other.ICMPs != nil)) || ((in.ICMPs == nil) != (other.ICMPs == nil)) {
		in, other := &in.ICMPs, &other.ICMPs
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
func (in *EgressRule) DeepEqual(other *EgressRule) bool {
	if other == nil {
		return false
	}

	if !in.EgressCommonRule.DeepEqual(&other.EgressCommonRule) {
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

	if ((in.ToFQDNs != nil) && (other.ToFQDNs != nil)) || ((in.ToFQDNs == nil) != (other.ToFQDNs == nil)) {
		in, other := &in.ToFQDNs, &other.ToFQDNs
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

	if ((in.ICMPs != nil) && (other.ICMPs != nil)) || ((in.ICMPs == nil) != (other.ICMPs == nil)) {
		in, other := &in.ICMPs, &other.ICMPs
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
func (in *EndpointSelector) DeepEqual(other *EndpointSelector) bool {
	if other == nil {
		return false
	}

	if (in.LabelSelector == nil) != (other.LabelSelector == nil) {
		return false
	} else if in.LabelSelector != nil {
		if !in.LabelSelector.DeepEqual(other.LabelSelector) {
			return false
		}
	}

	if (in.requirements == nil) != (other.requirements == nil) {
		return false
	} else if in.requirements != nil {
		if !in.requirements.DeepEqual(other.requirements) {
			return false
		}
	}

	if in.cachedLabelSelectorString != other.cachedLabelSelectorString {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointSelectorSlice) DeepEqual(other *EndpointSelectorSlice) bool {
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
func (in *EntitySlice) DeepEqual(other *EntitySlice) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EnvoyConfig) DeepEqual(other *EnvoyConfig) bool {
	if other == nil {
		return false
	}

	if in.Kind != other.Kind {
		return false
	}
	if in.Name != other.Name {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *FQDNSelector) DeepEqual(other *FQDNSelector) bool {
	if other == nil {
		return false
	}

	if in.MatchName != other.MatchName {
		return false
	}
	if in.MatchPattern != other.MatchPattern {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *FQDNSelectorSlice) DeepEqual(other *FQDNSelectorSlice) bool {
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
func (in *HeaderMatch) DeepEqual(other *HeaderMatch) bool {
	if other == nil {
		return false
	}

	if in.Mismatch != other.Mismatch {
		return false
	}
	if in.Name != other.Name {
		return false
	}
	if (in.Secret == nil) != (other.Secret == nil) {
		return false
	} else if in.Secret != nil {
		if !in.Secret.DeepEqual(other.Secret) {
			return false
		}
	}

	if in.Value != other.Value {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ICMPField) DeepEqual(other *ICMPField) bool {
	if other == nil {
		return false
	}

	if in.Family != other.Family {
		return false
	}
	if in.Type != other.Type {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ICMPRule) DeepEqual(other *ICMPRule) bool {
	if other == nil {
		return false
	}

	if ((in.Fields != nil) && (other.Fields != nil)) || ((in.Fields == nil) != (other.Fields == nil)) {
		in, other := &in.Fields, &other.Fields
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
func (in *ICMPRules) DeepEqual(other *ICMPRules) bool {
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
func (in *IngressCommonRule) DeepEqual(other *IngressCommonRule) bool {
	if other == nil {
		return false
	}

	if ((in.FromEndpoints != nil) && (other.FromEndpoints != nil)) || ((in.FromEndpoints == nil) != (other.FromEndpoints == nil)) {
		in, other := &in.FromEndpoints, &other.FromEndpoints
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

	if ((in.FromRequires != nil) && (other.FromRequires != nil)) || ((in.FromRequires == nil) != (other.FromRequires == nil)) {
		in, other := &in.FromRequires, &other.FromRequires
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

	if ((in.FromCIDR != nil) && (other.FromCIDR != nil)) || ((in.FromCIDR == nil) != (other.FromCIDR == nil)) {
		in, other := &in.FromCIDR, &other.FromCIDR
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.FromCIDRSet != nil) && (other.FromCIDRSet != nil)) || ((in.FromCIDRSet == nil) != (other.FromCIDRSet == nil)) {
		in, other := &in.FromCIDRSet, &other.FromCIDRSet
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.FromEntities != nil) && (other.FromEntities != nil)) || ((in.FromEntities == nil) != (other.FromEntities == nil)) {
		in, other := &in.FromEntities, &other.FromEntities
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.aggregatedSelectors != nil) && (other.aggregatedSelectors != nil)) || ((in.aggregatedSelectors == nil) != (other.aggregatedSelectors == nil)) {
		in, other := &in.aggregatedSelectors, &other.aggregatedSelectors
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IngressDenyRule) DeepEqual(other *IngressDenyRule) bool {
	if other == nil {
		return false
	}

	if !in.IngressCommonRule.DeepEqual(&other.IngressCommonRule) {
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

	if ((in.ICMPs != nil) && (other.ICMPs != nil)) || ((in.ICMPs == nil) != (other.ICMPs == nil)) {
		in, other := &in.ICMPs, &other.ICMPs
		if other == nil || !in.DeepEqual(other) {
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

	if !in.IngressCommonRule.DeepEqual(&other.IngressCommonRule) {
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

	if ((in.ICMPs != nil) && (other.ICMPs != nil)) || ((in.ICMPs == nil) != (other.ICMPs == nil)) {
		in, other := &in.ICMPs, &other.ICMPs
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *K8sServiceNamespace) DeepEqual(other *K8sServiceNamespace) bool {
	if other == nil {
		return false
	}

	if in.ServiceName != other.ServiceName {
		return false
	}
	if in.Namespace != other.Namespace {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *K8sServiceSelectorNamespace) DeepEqual(other *K8sServiceSelectorNamespace) bool {
	if other == nil {
		return false
	}

	if !in.Selector.DeepEqual(&other.Selector) {
		return false
	}

	if in.Namespace != other.Namespace {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *L7Rules) DeepEqual(other *L7Rules) bool {
	if other == nil {
		return false
	}

	if ((in.HTTP != nil) && (other.HTTP != nil)) || ((in.HTTP == nil) != (other.HTTP == nil)) {
		in, other := &in.HTTP, &other.HTTP
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

	if ((in.Kafka != nil) && (other.Kafka != nil)) || ((in.Kafka == nil) != (other.Kafka == nil)) {
		in, other := &in.Kafka, &other.Kafka
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

	if ((in.DNS != nil) && (other.DNS != nil)) || ((in.DNS == nil) != (other.DNS == nil)) {
		in, other := &in.DNS, &other.DNS
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

	if in.L7Proto != other.L7Proto {
		return false
	}
	if ((in.L7 != nil) && (other.L7 != nil)) || ((in.L7 == nil) != (other.L7 == nil)) {
		in, other := &in.L7, &other.L7
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
func (in *Listener) DeepEqual(other *Listener) bool {
	if other == nil {
		return false
	}

	if (in.EnvoyConfig == nil) != (other.EnvoyConfig == nil) {
		return false
	} else if in.EnvoyConfig != nil {
		if !in.EnvoyConfig.DeepEqual(other.EnvoyConfig) {
			return false
		}
	}

	if in.Name != other.Name {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortDenyRule) DeepEqual(other *PortDenyRule) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortDenyRules) DeepEqual(other *PortDenyRules) bool {
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
func (in *PortProtocol) DeepEqual(other *PortProtocol) bool {
	if other == nil {
		return false
	}

	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRule) DeepEqual(other *PortRule) bool {
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

	if (in.TerminatingTLS == nil) != (other.TerminatingTLS == nil) {
		return false
	} else if in.TerminatingTLS != nil {
		if !in.TerminatingTLS.DeepEqual(other.TerminatingTLS) {
			return false
		}
	}

	if (in.OriginatingTLS == nil) != (other.OriginatingTLS == nil) {
		return false
	} else if in.OriginatingTLS != nil {
		if !in.OriginatingTLS.DeepEqual(other.OriginatingTLS) {
			return false
		}
	}

	if ((in.ServerNames != nil) && (other.ServerNames != nil)) || ((in.ServerNames == nil) != (other.ServerNames == nil)) {
		in, other := &in.ServerNames, &other.ServerNames
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

	if (in.Listener == nil) != (other.Listener == nil) {
		return false
	} else if in.Listener != nil {
		if !in.Listener.DeepEqual(other.Listener) {
			return false
		}
	}

	if (in.Rules == nil) != (other.Rules == nil) {
		return false
	} else if in.Rules != nil {
		if !in.Rules.DeepEqual(other.Rules) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRuleDNS) DeepEqual(other *PortRuleDNS) bool {
	if other == nil {
		return false
	}

	if in.MatchName != other.MatchName {
		return false
	}
	if in.MatchPattern != other.MatchPattern {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRuleHTTP) DeepEqual(other *PortRuleHTTP) bool {
	if other == nil {
		return false
	}

	if in.Path != other.Path {
		return false
	}
	if in.Method != other.Method {
		return false
	}
	if in.Host != other.Host {
		return false
	}
	if ((in.Headers != nil) && (other.Headers != nil)) || ((in.Headers == nil) != (other.Headers == nil)) {
		in, other := &in.Headers, &other.Headers
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

	if ((in.HeaderMatches != nil) && (other.HeaderMatches != nil)) || ((in.HeaderMatches == nil) != (other.HeaderMatches == nil)) {
		in, other := &in.HeaderMatches, &other.HeaderMatches
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
func (in *PortRuleL7) DeepEqual(other *PortRuleL7) bool {
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

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRules) DeepEqual(other *PortRules) bool {
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
func (in *Rule) deepEqual(other *Rule) bool {
	if other == nil {
		return false
	}

	if !in.EndpointSelector.DeepEqual(&other.EndpointSelector) {
		return false
	}

	if !in.NodeSelector.DeepEqual(&other.NodeSelector) {
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

	if ((in.IngressDeny != nil) && (other.IngressDeny != nil)) || ((in.IngressDeny == nil) != (other.IngressDeny == nil)) {
		in, other := &in.IngressDeny, &other.IngressDeny
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

	if ((in.EgressDeny != nil) && (other.EgressDeny != nil)) || ((in.EgressDeny == nil) != (other.EgressDeny == nil)) {
		in, other := &in.EgressDeny, &other.EgressDeny
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

	if ((in.Labels != nil) && (other.Labels != nil)) || ((in.Labels == nil) != (other.Labels == nil)) {
		in, other := &in.Labels, &other.Labels
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

	if in.Description != other.Description {
		return false
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Rules) deepEqual(other *Rules) bool {
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
func (in *Secret) DeepEqual(other *Secret) bool {
	if other == nil {
		return false
	}

	if in.Namespace != other.Namespace {
		return false
	}
	if in.Name != other.Name {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Service) DeepEqual(other *Service) bool {
	if other == nil {
		return false
	}

	if (in.K8sServiceSelector == nil) != (other.K8sServiceSelector == nil) {
		return false
	} else if in.K8sServiceSelector != nil {
		if !in.K8sServiceSelector.DeepEqual(other.K8sServiceSelector) {
			return false
		}
	}

	if (in.K8sService == nil) != (other.K8sService == nil) {
		return false
	} else if in.K8sService != nil {
		if !in.K8sService.DeepEqual(other.K8sService) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceSelector) DeepEqual(other *ServiceSelector) bool {
	if other == nil {
		return false
	}

	if (in.LabelSelector == nil) != (other.LabelSelector == nil) {
		return false
	} else if in.LabelSelector != nil {
		if !in.LabelSelector.DeepEqual(other.LabelSelector) {
			return false
		}
	}

	if (in.requirements == nil) != (other.requirements == nil) {
		return false
	} else if in.requirements != nil {
		if !in.requirements.DeepEqual(other.requirements) {
			return false
		}
	}

	if in.cachedLabelSelectorString != other.cachedLabelSelectorString {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *TLSContext) DeepEqual(other *TLSContext) bool {
	if other == nil {
		return false
	}

	if (in.Secret == nil) != (other.Secret == nil) {
		return false
	} else if in.Secret != nil {
		if !in.Secret.DeepEqual(other.Secret) {
			return false
		}
	}

	if in.TrustedCA != other.TrustedCA {
		return false
	}
	if in.Certificate != other.Certificate {
		return false
	}
	if in.PrivateKey != other.PrivateKey {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ToGroups) DeepEqual(other *ToGroups) bool {
	if other == nil {
		return false
	}

	if (in.AWS == nil) != (other.AWS == nil) {
		return false
	} else if in.AWS != nil {
		if !in.AWS.DeepEqual(other.AWS) {
			return false
		}
	}

	return true
}
