// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumClusterwideNetworkPolicies implements CiliumClusterwideNetworkPolicyInterface
type FakeCiliumClusterwideNetworkPolicies struct {
	Fake *FakeCiliumV2
}

var ciliumclusterwidenetworkpoliciesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumclusterwidenetworkpolicies"}

var ciliumclusterwidenetworkpoliciesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumClusterwideNetworkPolicy"}

// Get takes name of the ciliumClusterwideNetworkPolicy, and returns the corresponding ciliumClusterwideNetworkPolicy object, and an error if there is any.
func (c *FakeCiliumClusterwideNetworkPolicies) Get(name string, options v1.GetOptions) (result *v2.CiliumClusterwideNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumclusterwidenetworkpoliciesResource, name), &v2.CiliumClusterwideNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumClusterwideNetworkPolicies that match those selectors.
func (c *FakeCiliumClusterwideNetworkPolicies) List(opts v1.ListOptions) (result *v2.CiliumClusterwideNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumclusterwidenetworkpoliciesResource, ciliumclusterwidenetworkpoliciesKind, opts), &v2.CiliumClusterwideNetworkPolicyList{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicyList), err
}

// Watch returns a watch.Interface that watches the requested ciliumClusterwideNetworkPolicies.
func (c *FakeCiliumClusterwideNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumclusterwidenetworkpoliciesResource, opts))
}

// Create takes the representation of a ciliumClusterwideNetworkPolicy and creates it.  Returns the server's representation of the ciliumClusterwideNetworkPolicy, and an error, if there is any.
func (c *FakeCiliumClusterwideNetworkPolicies) Create(ciliumClusterwideNetworkPolicy *v2.CiliumClusterwideNetworkPolicy) (result *v2.CiliumClusterwideNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumclusterwidenetworkpoliciesResource, ciliumClusterwideNetworkPolicy), &v2.CiliumClusterwideNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicy), err
}

// Update takes the representation of a ciliumClusterwideNetworkPolicy and updates it. Returns the server's representation of the ciliumClusterwideNetworkPolicy, and an error, if there is any.
func (c *FakeCiliumClusterwideNetworkPolicies) Update(ciliumClusterwideNetworkPolicy *v2.CiliumClusterwideNetworkPolicy) (result *v2.CiliumClusterwideNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumclusterwidenetworkpoliciesResource, ciliumClusterwideNetworkPolicy), &v2.CiliumClusterwideNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicy), err
}

// Delete takes name of the ciliumClusterwideNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumClusterwideNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ciliumclusterwidenetworkpoliciesResource, name), &v2.CiliumClusterwideNetworkPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumClusterwideNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumclusterwidenetworkpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v2.CiliumClusterwideNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumClusterwideNetworkPolicy.
func (c *FakeCiliumClusterwideNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.CiliumClusterwideNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumclusterwidenetworkpoliciesResource, name, pt, data, subresources...), &v2.CiliumClusterwideNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicy), err
}
