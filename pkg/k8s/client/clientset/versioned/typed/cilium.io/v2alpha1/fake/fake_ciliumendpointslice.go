// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumEndpointSlices implements CiliumEndpointSliceInterface
type FakeCiliumEndpointSlices struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumendpointslicesResource = v2alpha1.SchemeGroupVersion.WithResource("ciliumendpointslices")

var ciliumendpointslicesKind = v2alpha1.SchemeGroupVersion.WithKind("CiliumEndpointSlice")

// Get takes name of the ciliumEndpointSlice, and returns the corresponding ciliumEndpointSlice object, and an error if there is any.
func (c *FakeCiliumEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumendpointslicesResource, name), &v2alpha1.CiliumEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointSlice), err
}

// List takes label and field selectors, and returns the list of CiliumEndpointSlices that match those selectors.
func (c *FakeCiliumEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumEndpointSliceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumendpointslicesResource, ciliumendpointslicesKind, opts), &v2alpha1.CiliumEndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumEndpointSliceList{ListMeta: obj.(*v2alpha1.CiliumEndpointSliceList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumEndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEndpointSlices.
func (c *FakeCiliumEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumendpointslicesResource, opts))
}

// Create takes the representation of a ciliumEndpointSlice and creates it.  Returns the server's representation of the ciliumEndpointSlice, and an error, if there is any.
func (c *FakeCiliumEndpointSlices) Create(ctx context.Context, ciliumEndpointSlice *v2alpha1.CiliumEndpointSlice, opts v1.CreateOptions) (result *v2alpha1.CiliumEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumendpointslicesResource, ciliumEndpointSlice), &v2alpha1.CiliumEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointSlice), err
}

// Update takes the representation of a ciliumEndpointSlice and updates it. Returns the server's representation of the ciliumEndpointSlice, and an error, if there is any.
func (c *FakeCiliumEndpointSlices) Update(ctx context.Context, ciliumEndpointSlice *v2alpha1.CiliumEndpointSlice, opts v1.UpdateOptions) (result *v2alpha1.CiliumEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumendpointslicesResource, ciliumEndpointSlice), &v2alpha1.CiliumEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointSlice), err
}

// Delete takes name of the ciliumEndpointSlice and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumendpointslicesResource, name, opts), &v2alpha1.CiliumEndpointSlice{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumendpointslicesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumEndpointSliceList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEndpointSlice.
func (c *FakeCiliumEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumendpointslicesResource, name, pt, data, subresources...), &v2alpha1.CiliumEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumEndpointSlice), err
}
