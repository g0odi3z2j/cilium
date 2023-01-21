//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package tunnel

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelIP) DeepCopyInto(out *TunnelIP) {
	*out = *in
	in.IP.DeepCopyInto(&out.IP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelIP.
func (in *TunnelIP) DeepCopy() *TunnelIP {
	if in == nil {
		return nil
	}
	out := new(TunnelIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelKey) DeepCopyInto(out *TunnelKey) {
	*out = *in
	in.TunnelIP.DeepCopyInto(&out.TunnelIP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelKey.
func (in *TunnelKey) DeepCopy() *TunnelKey {
	if in == nil {
		return nil
	}
	out := new(TunnelKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TunnelKey) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelValue) DeepCopyInto(out *TunnelValue) {
	*out = *in
	in.TunnelIP.DeepCopyInto(&out.TunnelIP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelValue.
func (in *TunnelValue) DeepCopy() *TunnelValue {
	if in == nil {
		return nil
	}
	out := new(TunnelValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *TunnelValue) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
