//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSubnet) DeepCopyInto(out *AwsSubnet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSubnet.
func (in *AwsSubnet) DeepCopy() *AwsSubnet {
	if in == nil {
		return nil
	}
	out := new(AwsSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsVPC) DeepCopyInto(out *AwsVPC) {
	*out = *in
	if in.CIDRs != nil {
		in, out := &in.CIDRs, &out.CIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsVPC.
func (in *AwsVPC) DeepCopy() *AwsVPC {
	if in == nil {
		return nil
	}
	out := new(AwsVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENI) DeepCopyInto(out *ENI) {
	*out = *in
	out.Subnet = in.Subnet
	in.VPC.DeepCopyInto(&out.VPC)
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENI.
func (in *ENI) DeepCopy() *ENI {
	if in == nil {
		return nil
	}
	out := new(ENI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENISpec) DeepCopyInto(out *ENISpec) {
	*out = *in
	if in.FirstInterfaceIndex != nil {
		in, out := &in.FirstInterfaceIndex, &out.FirstInterfaceIndex
		*out = new(int)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupTags != nil {
		in, out := &in.SecurityGroupTags, &out.SecurityGroupTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetTags != nil {
		in, out := &in.SubnetTags, &out.SubnetTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENISpec.
func (in *ENISpec) DeepCopy() *ENISpec {
	if in == nil {
		return nil
	}
	out := new(ENISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENIStatus) DeepCopyInto(out *ENIStatus) {
	*out = *in
	if in.ENIs != nil {
		in, out := &in.ENIs, &out.ENIs
		*out = make(map[string]ENI, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENIStatus.
func (in *ENIStatus) DeepCopy() *ENIStatus {
	if in == nil {
		return nil
	}
	out := new(ENIStatus)
	in.DeepCopyInto(out)
	return out
}
