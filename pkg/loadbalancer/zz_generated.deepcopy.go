//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package loadbalancer

import (
	net "net"

	cidr "github.com/cilium/cilium/pkg/cidr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	in.L3n4Addr.DeepCopyInto(&out.L3n4Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3n4Addr) DeepCopyInto(out *L3n4Addr) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	out.L4Addr = in.L4Addr
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3n4Addr.
func (in *L3n4Addr) DeepCopy() *L3n4Addr {
	if in == nil {
		return nil
	}
	out := new(L3n4Addr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3n4AddrID) DeepCopyInto(out *L3n4AddrID) {
	*out = *in
	in.L3n4Addr.DeepCopyInto(&out.L3n4Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3n4AddrID.
func (in *L3n4AddrID) DeepCopy() *L3n4AddrID {
	if in == nil {
		return nil
	}
	out := new(L3n4AddrID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4Addr) DeepCopyInto(out *L4Addr) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4Addr.
func (in *L4Addr) DeepCopy() *L4Addr {
	if in == nil {
		return nil
	}
	out := new(L4Addr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SVC) DeepCopyInto(out *SVC) {
	*out = *in
	in.Frontend.DeepCopyInto(&out.Frontend)
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]*cidr.CIDR, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SVC.
func (in *SVC) DeepCopy() *SVC {
	if in == nil {
		return nil
	}
	out := new(SVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SvcFlagParam) DeepCopyInto(out *SvcFlagParam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SvcFlagParam.
func (in *SvcFlagParam) DeepCopy() *SvcFlagParam {
	if in == nil {
		return nil
	}
	out := new(SvcFlagParam)
	in.DeepCopyInto(out)
	return out
}
