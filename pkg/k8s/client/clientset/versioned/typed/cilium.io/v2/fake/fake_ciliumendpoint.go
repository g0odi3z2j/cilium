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

// FakeCiliumEndpoints implements CiliumEndpointInterface
type FakeCiliumEndpoints struct {
	Fake *FakeCiliumV2
	ns   string
}

var ciliumendpointsResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumendpoints"}

var ciliumendpointsKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumEndpoint"}

// Get takes name of the ciliumEndpoint, and returns the corresponding ciliumEndpoint object, and an error if there is any.
func (c *FakeCiliumEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ciliumendpointsResource, c.ns, name), &v2.CiliumEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEndpoint), err
}

// List takes label and field selectors, and returns the list of CiliumEndpoints that match those selectors.
func (c *FakeCiliumEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ciliumendpointsResource, ciliumendpointsKind, c.ns, opts), &v2.CiliumEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumEndpointList{ListMeta: obj.(*v2.CiliumEndpointList).ListMeta}
	for _, item := range obj.(*v2.CiliumEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEndpoints.
func (c *FakeCiliumEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ciliumendpointsResource, c.ns, opts))

}

// Create takes the representation of a ciliumEndpoint and creates it.  Returns the server's representation of the ciliumEndpoint, and an error, if there is any.
func (c *FakeCiliumEndpoints) Create(ctx context.Context, ciliumEndpoint *v2.CiliumEndpoint, opts v1.CreateOptions) (result *v2.CiliumEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ciliumendpointsResource, c.ns, ciliumEndpoint), &v2.CiliumEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEndpoint), err
}

// Update takes the representation of a ciliumEndpoint and updates it. Returns the server's representation of the ciliumEndpoint, and an error, if there is any.
func (c *FakeCiliumEndpoints) Update(ctx context.Context, ciliumEndpoint *v2.CiliumEndpoint, opts v1.UpdateOptions) (result *v2.CiliumEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ciliumendpointsResource, c.ns, ciliumEndpoint), &v2.CiliumEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiliumEndpoints) UpdateStatus(ctx context.Context, ciliumEndpoint *v2.CiliumEndpoint, opts v1.UpdateOptions) (*v2.CiliumEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ciliumendpointsResource, "status", c.ns, ciliumEndpoint), &v2.CiliumEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEndpoint), err
}

// Delete takes name of the ciliumEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ciliumendpointsResource, c.ns, name, opts), &v2.CiliumEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ciliumendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumEndpointList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEndpoint.
func (c *FakeCiliumEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ciliumendpointsResource, c.ns, name, pt, data, subresources...), &v2.CiliumEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEndpoint), err
}
