// +build !ignore_autogenerated

// Copyright 2017-2019 Authors of Cilium
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package tunnel

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelEndpoint) DeepCopyInto(out *TunnelEndpoint) {
	*out = *in
	in.EndpointKey.DeepCopyInto(&out.EndpointKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelEndpoint.
func (in *TunnelEndpoint) DeepCopy() *TunnelEndpoint {
	if in == nil {
		return nil
	}
	out := new(TunnelEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TunnelEndpoint) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *TunnelEndpoint) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
