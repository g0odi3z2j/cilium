// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

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

// FakeCiliumNetworkPolicies implements CiliumNetworkPolicyInterface
type FakeCiliumNetworkPolicies struct {
	Fake *FakeCiliumV2
	ns   string
}

var ciliumnetworkpoliciesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumnetworkpolicies"}

var ciliumnetworkpoliciesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumNetworkPolicy"}

// Get takes name of the ciliumNetworkPolicy, and returns the corresponding ciliumNetworkPolicy object, and an error if there is any.
func (c *FakeCiliumNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ciliumnetworkpoliciesResource, c.ns, name), &v2.CiliumNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumNetworkPolicies that match those selectors.
func (c *FakeCiliumNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ciliumnetworkpoliciesResource, ciliumnetworkpoliciesKind, c.ns, opts), &v2.CiliumNetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumNetworkPolicyList{ListMeta: obj.(*v2.CiliumNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v2.CiliumNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumNetworkPolicies.
func (c *FakeCiliumNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ciliumnetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a ciliumNetworkPolicy and creates it.  Returns the server's representation of the ciliumNetworkPolicy, and an error, if there is any.
func (c *FakeCiliumNetworkPolicies) Create(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.CreateOptions) (result *v2.CiliumNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ciliumnetworkpoliciesResource, c.ns, ciliumNetworkPolicy), &v2.CiliumNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumNetworkPolicy), err
}

// Update takes the representation of a ciliumNetworkPolicy and updates it. Returns the server's representation of the ciliumNetworkPolicy, and an error, if there is any.
func (c *FakeCiliumNetworkPolicies) Update(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (result *v2.CiliumNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ciliumnetworkpoliciesResource, c.ns, ciliumNetworkPolicy), &v2.CiliumNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumNetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiliumNetworkPolicies) UpdateStatus(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (*v2.CiliumNetworkPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ciliumnetworkpoliciesResource, "status", c.ns, ciliumNetworkPolicy), &v2.CiliumNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumNetworkPolicy), err
}

// Delete takes name of the ciliumNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ciliumnetworkpoliciesResource, c.ns, name, opts), &v2.CiliumNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ciliumnetworkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumNetworkPolicy.
func (c *FakeCiliumNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ciliumnetworkpoliciesResource, c.ns, name, pt, data, subresources...), &v2.CiliumNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumNetworkPolicy), err
}
