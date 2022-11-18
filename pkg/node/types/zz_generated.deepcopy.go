//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package types

import (
	net "net"

	cidr "github.com/cilium/cilium/pkg/cidr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]Address, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPv4AllocCIDR != nil {
		in, out := &in.IPv4AllocCIDR, &out.IPv4AllocCIDR
		*out = (*in).DeepCopy()
	}
	if in.IPv4SecondaryAllocCIDRs != nil {
		in, out := &in.IPv4SecondaryAllocCIDRs, &out.IPv4SecondaryAllocCIDRs
		*out = make([]*cidr.CIDR, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IPv6AllocCIDR != nil {
		in, out := &in.IPv6AllocCIDR, &out.IPv6AllocCIDR
		*out = (*in).DeepCopy()
	}
	if in.IPv6SecondaryAllocCIDRs != nil {
		in, out := &in.IPv6SecondaryAllocCIDRs, &out.IPv6SecondaryAllocCIDRs
		*out = make([]*cidr.CIDR, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IPv4HealthIP != nil {
		in, out := &in.IPv4HealthIP, &out.IPv4HealthIP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	if in.IPv6HealthIP != nil {
		in, out := &in.IPv6HealthIP, &out.IPv6HealthIP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	if in.IPv4IngressIP != nil {
		in, out := &in.IPv4IngressIP, &out.IPv4IngressIP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	if in.IPv6IngressIP != nil {
		in, out := &in.IPv6IngressIP, &out.IPv6IngressIP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterNode) DeepCopyInto(out *RegisterNode) {
	*out = *in
	in.Node.DeepCopyInto(&out.Node)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterNode.
func (in *RegisterNode) DeepCopy() *RegisterNode {
	if in == nil {
		return nil
	}
	out := new(RegisterNode)
	in.DeepCopyInto(out)
	return out
}
