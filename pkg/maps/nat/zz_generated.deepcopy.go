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

package nat

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatEntry4) DeepCopyInto(out *NatEntry4) {
	*out = *in
	in.Addr.DeepCopyInto(&out.Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatEntry4.
func (in *NatEntry4) DeepCopy() *NatEntry4 {
	if in == nil {
		return nil
	}
	out := new(NatEntry4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *NatEntry4) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatEntry6) DeepCopyInto(out *NatEntry6) {
	*out = *in
	in.Addr.DeepCopyInto(&out.Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatEntry6.
func (in *NatEntry6) DeepCopy() *NatEntry6 {
	if in == nil {
		return nil
	}
	out := new(NatEntry6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *NatEntry6) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatKey4) DeepCopyInto(out *NatKey4) {
	*out = *in
	in.TupleKey4Global.DeepCopyInto(&out.TupleKey4Global)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatKey4.
func (in *NatKey4) DeepCopy() *NatKey4 {
	if in == nil {
		return nil
	}
	out := new(NatKey4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *NatKey4) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatKey6) DeepCopyInto(out *NatKey6) {
	*out = *in
	in.TupleKey6Global.DeepCopyInto(&out.TupleKey6Global)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatKey6.
func (in *NatKey6) DeepCopy() *NatKey6 {
	if in == nil {
		return nil
	}
	out := new(NatKey6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *NatKey6) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
