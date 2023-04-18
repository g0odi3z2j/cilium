// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeNetworkingV1
	ns   string
}

var networkpoliciesResource = v1.SchemeGroupVersion.WithResource("networkpolicies")

var networkpoliciesKind = v1.SchemeGroupVersion.WithKind("NetworkPolicy")

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkpoliciesResource, c.ns, name), &v1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *FakeNetworkPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkpoliciesResource, networkpoliciesKind, c.ns, opts), &v1.NetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NetworkPolicyList{ListMeta: obj.(*v1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Create(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.CreateOptions) (result *v1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkpoliciesResource, c.ns, networkPolicy), &v1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Update(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (result *v1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkpoliciesResource, c.ns, networkPolicy), &v1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPolicies) UpdateStatus(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (*v1.NetworkPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkpoliciesResource, "status", c.ns, networkPolicy), &v1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkpoliciesResource, c.ns, name, opts), &v1.NetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *FakeNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkpoliciesResource, c.ns, name, pt, data, subresources...), &v1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPolicy), err
}
