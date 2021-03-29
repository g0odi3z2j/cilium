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
func (in *Endpoint) DeepEqual(other *Endpoint) bool {
	if other == nil {
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
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if !in.Conditions.DeepEqual(&other.Conditions) {
		return false
	}

	if ((in.DeprecatedTopology != nil) && (other.DeprecatedTopology != nil)) || ((in.DeprecatedTopology == nil) != (other.DeprecatedTopology == nil)) {
		in, other := &in.DeprecatedTopology, &other.DeprecatedTopology
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

	if (in.NodeName == nil) != (other.NodeName == nil) {
		return false
	} else if in.NodeName != nil {
		if *in.NodeName != *other.NodeName {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointConditions) DeepEqual(other *EndpointConditions) bool {
	if other == nil {
		return false
	}

	if (in.Ready == nil) != (other.Ready == nil) {
		return false
	} else if in.Ready != nil {
		if *in.Ready != *other.Ready {
			return false
		}
	}

	if (in.Serving == nil) != (other.Serving == nil) {
		return false
	} else if in.Serving != nil {
		if *in.Serving != *other.Serving {
			return false
		}
	}

	if (in.Terminating == nil) != (other.Terminating == nil) {
		return false
	} else if in.Terminating != nil {
		if *in.Terminating != *other.Terminating {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointPort) DeepEqual(other *EndpointPort) bool {
	if other == nil {
		return false
	}

	if (in.Name == nil) != (other.Name == nil) {
		return false
	} else if in.Name != nil {
		if *in.Name != *other.Name {
			return false
		}
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
		if *in.Port != *other.Port {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointSlice) DeepEqual(other *EndpointSlice) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if in.AddressType != other.AddressType {
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
func (in *EndpointSliceList) DeepEqual(other *EndpointSliceList) bool {
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
