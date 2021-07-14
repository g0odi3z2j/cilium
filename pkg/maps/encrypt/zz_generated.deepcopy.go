// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package encrypt

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptKey) DeepCopyInto(out *EncryptKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptKey.
func (in *EncryptKey) DeepCopy() *EncryptKey {
	if in == nil {
		return nil
	}
	out := new(EncryptKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *EncryptKey) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptValue) DeepCopyInto(out *EncryptValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptValue.
func (in *EncryptValue) DeepCopy() *EncryptValue {
	if in == nil {
		return nil
	}
	out := new(EncryptValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *EncryptValue) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
