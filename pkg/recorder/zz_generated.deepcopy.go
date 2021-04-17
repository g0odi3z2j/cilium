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

// Code generated by deepcopy-gen. DO NOT EDIT.

package recorder

import (
	net "net"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecInfo) DeepCopyInto(out *RecInfo) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RecorderTuple, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecInfo.
func (in *RecInfo) DeepCopy() *RecInfo {
	if in == nil {
		return nil
	}
	out := new(RecInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecMask) DeepCopyInto(out *RecMask) {
	*out = *in
	in.mask.DeepCopyInto(&out.mask)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecMask.
func (in *RecMask) DeepCopy() *RecMask {
	if in == nil {
		return nil
	}
	out := new(RecMask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecorderMask) DeepCopyInto(out *RecorderMask) {
	*out = *in
	if in.srcMask != nil {
		in, out := &in.srcMask, &out.srcMask
		*out = make(net.IPMask, len(*in))
		copy(*out, *in)
	}
	if in.dstMask != nil {
		in, out := &in.dstMask, &out.dstMask
		*out = make(net.IPMask, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecorderMask.
func (in *RecorderMask) DeepCopy() *RecorderMask {
	if in == nil {
		return nil
	}
	out := new(RecorderMask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecorderTuple) DeepCopyInto(out *RecorderTuple) {
	*out = *in
	in.SrcPrefix.DeepCopyInto(&out.SrcPrefix)
	in.DstPrefix.DeepCopyInto(&out.DstPrefix)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecorderTuple.
func (in *RecorderTuple) DeepCopy() *RecorderTuple {
	if in == nil {
		return nil
	}
	out := new(RecorderTuple)
	in.DeepCopyInto(out)
	return out
}
