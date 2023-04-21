// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkingv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressClasses implements IngressClassInterface
type FakeIngressClasses struct {
	Fake *FakeNetworkingV1
}

var ingressclassesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingressclasses"}

var ingressclassesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1", Kind: "IngressClass"}

// Get takes name of the ingressClass, and returns the corresponding ingressClass object, and an error if there is any.
func (c *FakeIngressClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ingressclassesResource, name), &networkingv1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IngressClass), err
}

// List takes label and field selectors, and returns the list of IngressClasses that match those selectors.
func (c *FakeIngressClasses) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.IngressClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ingressclassesResource, ingressclassesKind, opts), &networkingv1.IngressClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.IngressClassList{ListMeta: obj.(*networkingv1.IngressClassList).ListMeta}
	for _, item := range obj.(*networkingv1.IngressClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressClasses.
func (c *FakeIngressClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ingressclassesResource, opts))
}

// Create takes the representation of a ingressClass and creates it.  Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Create(ctx context.Context, ingressClass *networkingv1.IngressClass, opts v1.CreateOptions) (result *networkingv1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ingressclassesResource, ingressClass), &networkingv1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IngressClass), err
}

// Update takes the representation of a ingressClass and updates it. Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Update(ctx context.Context, ingressClass *networkingv1.IngressClass, opts v1.UpdateOptions) (result *networkingv1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ingressclassesResource, ingressClass), &networkingv1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IngressClass), err
}

// Delete takes name of the ingressClass and deletes it. Returns an error if one occurs.
func (c *FakeIngressClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ingressclassesResource, name, opts), &networkingv1.IngressClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ingressclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.IngressClassList{})
	return err
}

// Patch applies the patch and returns the patched ingressClass.
func (c *FakeIngressClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ingressclassesResource, name, pt, data, subresources...), &networkingv1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IngressClass), err
}
