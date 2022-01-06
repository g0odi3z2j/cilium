// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumLocalRedirectPolicies implements CiliumLocalRedirectPolicyInterface
type FakeCiliumLocalRedirectPolicies struct {
	Fake *FakeCiliumV2
	ns   string
}

var ciliumlocalredirectpoliciesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumlocalredirectpolicies"}

var ciliumlocalredirectpoliciesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumLocalRedirectPolicy"}

// Get takes name of the ciliumLocalRedirectPolicy, and returns the corresponding ciliumLocalRedirectPolicy object, and an error if there is any.
func (c *FakeCiliumLocalRedirectPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumLocalRedirectPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ciliumlocalredirectpoliciesResource, c.ns, name), &v2.CiliumLocalRedirectPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumLocalRedirectPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumLocalRedirectPolicies that match those selectors.
func (c *FakeCiliumLocalRedirectPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumLocalRedirectPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ciliumlocalredirectpoliciesResource, ciliumlocalredirectpoliciesKind, c.ns, opts), &v2.CiliumLocalRedirectPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumLocalRedirectPolicyList{ListMeta: obj.(*v2.CiliumLocalRedirectPolicyList).ListMeta}
	for _, item := range obj.(*v2.CiliumLocalRedirectPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumLocalRedirectPolicies.
func (c *FakeCiliumLocalRedirectPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ciliumlocalredirectpoliciesResource, c.ns, opts))

}

// Create takes the representation of a ciliumLocalRedirectPolicy and creates it.  Returns the server's representation of the ciliumLocalRedirectPolicy, and an error, if there is any.
func (c *FakeCiliumLocalRedirectPolicies) Create(ctx context.Context, ciliumLocalRedirectPolicy *v2.CiliumLocalRedirectPolicy, opts v1.CreateOptions) (result *v2.CiliumLocalRedirectPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ciliumlocalredirectpoliciesResource, c.ns, ciliumLocalRedirectPolicy), &v2.CiliumLocalRedirectPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumLocalRedirectPolicy), err
}

// Update takes the representation of a ciliumLocalRedirectPolicy and updates it. Returns the server's representation of the ciliumLocalRedirectPolicy, and an error, if there is any.
func (c *FakeCiliumLocalRedirectPolicies) Update(ctx context.Context, ciliumLocalRedirectPolicy *v2.CiliumLocalRedirectPolicy, opts v1.UpdateOptions) (result *v2.CiliumLocalRedirectPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ciliumlocalredirectpoliciesResource, c.ns, ciliumLocalRedirectPolicy), &v2.CiliumLocalRedirectPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumLocalRedirectPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiliumLocalRedirectPolicies) UpdateStatus(ctx context.Context, ciliumLocalRedirectPolicy *v2.CiliumLocalRedirectPolicy, opts v1.UpdateOptions) (*v2.CiliumLocalRedirectPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ciliumlocalredirectpoliciesResource, "status", c.ns, ciliumLocalRedirectPolicy), &v2.CiliumLocalRedirectPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumLocalRedirectPolicy), err
}

// Delete takes name of the ciliumLocalRedirectPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumLocalRedirectPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ciliumlocalredirectpoliciesResource, c.ns, name, opts), &v2.CiliumLocalRedirectPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumLocalRedirectPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ciliumlocalredirectpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumLocalRedirectPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumLocalRedirectPolicy.
func (c *FakeCiliumLocalRedirectPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumLocalRedirectPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ciliumlocalredirectpoliciesResource, c.ns, name, pt, data, subresources...), &v2.CiliumLocalRedirectPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumLocalRedirectPolicy), err
}
