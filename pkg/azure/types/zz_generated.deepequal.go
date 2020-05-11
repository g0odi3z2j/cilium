// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
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

package types

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AzureAddress) DeepEqual(other *AzureAddress) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}
	if in.Subnet != other.Subnet {
		return false
	}
	if in.State != other.State {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AzureInterface) DeepEqual(other *AzureInterface) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.Name != other.Name {
		return false
	}
	if in.MAC != other.MAC {
		return false
	}
	if in.State != other.State {
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

	if in.SecurityGroup != other.SecurityGroup {
		return false
	}
	if in.GatewayIP != other.GatewayIP {
		return false
	}
	if in.vmssName != other.vmssName {
		return false
	}
	if in.vmID != other.vmID {
		return false
	}
	if in.resourceGroup != other.resourceGroup {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AzureSpec) DeepEqual(other *AzureSpec) bool {
	if other == nil {
		return false
	}

	if in.InterfaceName != other.InterfaceName {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AzureStatus) DeepEqual(other *AzureStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Interfaces != nil) && (other.Interfaces != nil)) || ((in.Interfaces == nil) != (other.Interfaces == nil)) {
		in, other := &in.Interfaces, &other.Interfaces
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
