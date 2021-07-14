// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumEgressNATPolicies implements CiliumEgressNATPolicyInterface
type FakeCiliumEgressNATPolicies struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumegressnatpoliciesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumegressnatpolicies"}

var ciliumegressnatpoliciesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2alpha1", Kind: "CiliumEgressNATPolicy"}

// Get takes name of the ciliumEgressNATPolicy, and returns the corresponding ciliumEgressNATPolicy object, and an error if there is any.
func (c *FakeCiliumEgressNATPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumegressnatpoliciesResource, name), &v2alpha1.CiliumEgressNATPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEgressNATPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumEgressNATPolicies that match those selectors.
func (c *FakeCiliumEgressNATPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumEgressNATPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumegressnatpoliciesResource, ciliumegressnatpoliciesKind, opts), &v2alpha1.CiliumEgressNATPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumEgressNATPolicyList{ListMeta: obj.(*v2alpha1.CiliumEgressNATPolicyList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumEgressNATPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEgressNATPolicies.
func (c *FakeCiliumEgressNATPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumegressnatpoliciesResource, opts))
}

// Create takes the representation of a ciliumEgressNATPolicy and creates it.  Returns the server's representation of the ciliumEgressNATPolicy, and an error, if there is any.
func (c *FakeCiliumEgressNATPolicies) Create(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.CreateOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumegressnatpoliciesResource, ciliumEgressNATPolicy), &v2alpha1.CiliumEgressNATPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEgressNATPolicy), err
}

// Update takes the representation of a ciliumEgressNATPolicy and updates it. Returns the server's representation of the ciliumEgressNATPolicy, and an error, if there is any.
func (c *FakeCiliumEgressNATPolicies) Update(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.UpdateOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumegressnatpoliciesResource, ciliumEgressNATPolicy), &v2alpha1.CiliumEgressNATPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEgressNATPolicy), err
}

// Delete takes name of the ciliumEgressNATPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEgressNATPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ciliumegressnatpoliciesResource, name), &v2alpha1.CiliumEgressNATPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEgressNATPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumegressnatpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumEgressNATPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEgressNATPolicy.
func (c *FakeCiliumEgressNATPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumegressnatpoliciesResource, name, pt, data, subresources...), &v2alpha1.CiliumEgressNATPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEgressNATPolicy), err
}
