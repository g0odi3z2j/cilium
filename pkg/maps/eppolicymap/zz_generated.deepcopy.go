//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package eppolicymap

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EPPolicyValue) DeepCopyInto(out *EPPolicyValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EPPolicyValue.
func (in *EPPolicyValue) DeepCopy() *EPPolicyValue {
	if in == nil {
		return nil
	}
	out := new(EPPolicyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *EPPolicyValue) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointKey) DeepCopyInto(out *EndpointKey) {
	*out = *in
	in.EndpointKey.DeepCopyInto(&out.EndpointKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointKey.
func (in *EndpointKey) DeepCopy() *EndpointKey {
	if in == nil {
		return nil
	}
	out := new(EndpointKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *EndpointKey) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
