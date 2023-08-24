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

// FakeCiliumBGPClusterConfigs implements CiliumBGPClusterConfigInterface
type FakeCiliumBGPClusterConfigs struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumbgpclusterconfigsResource = v2alpha1.SchemeGroupVersion.WithResource("ciliumbgpclusterconfigs")

var ciliumbgpclusterconfigsKind = v2alpha1.SchemeGroupVersion.WithKind("CiliumBGPClusterConfig")

// Get takes name of the ciliumBGPClusterConfig, and returns the corresponding ciliumBGPClusterConfig object, and an error if there is any.
func (c *FakeCiliumBGPClusterConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumBGPClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumbgpclusterconfigsResource, name), &v2alpha1.CiliumBGPClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPClusterConfig), err
}

// List takes label and field selectors, and returns the list of CiliumBGPClusterConfigs that match those selectors.
func (c *FakeCiliumBGPClusterConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumBGPClusterConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumbgpclusterconfigsResource, ciliumbgpclusterconfigsKind, opts), &v2alpha1.CiliumBGPClusterConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumBGPClusterConfigList{ListMeta: obj.(*v2alpha1.CiliumBGPClusterConfigList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumBGPClusterConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumBGPClusterConfigs.
func (c *FakeCiliumBGPClusterConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumbgpclusterconfigsResource, opts))
}

// Create takes the representation of a ciliumBGPClusterConfig and creates it.  Returns the server's representation of the ciliumBGPClusterConfig, and an error, if there is any.
func (c *FakeCiliumBGPClusterConfigs) Create(ctx context.Context, ciliumBGPClusterConfig *v2alpha1.CiliumBGPClusterConfig, opts v1.CreateOptions) (result *v2alpha1.CiliumBGPClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumbgpclusterconfigsResource, ciliumBGPClusterConfig), &v2alpha1.CiliumBGPClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPClusterConfig), err
}

// Update takes the representation of a ciliumBGPClusterConfig and updates it. Returns the server's representation of the ciliumBGPClusterConfig, and an error, if there is any.
func (c *FakeCiliumBGPClusterConfigs) Update(ctx context.Context, ciliumBGPClusterConfig *v2alpha1.CiliumBGPClusterConfig, opts v1.UpdateOptions) (result *v2alpha1.CiliumBGPClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumbgpclusterconfigsResource, ciliumBGPClusterConfig), &v2alpha1.CiliumBGPClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPClusterConfig), err
}

// Delete takes name of the ciliumBGPClusterConfig and deletes it. Returns an error if one occurs.
func (c *FakeCiliumBGPClusterConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumbgpclusterconfigsResource, name, opts), &v2alpha1.CiliumBGPClusterConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumBGPClusterConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumbgpclusterconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumBGPClusterConfigList{})
	return err
}

// Patch applies the patch and returns the patched ciliumBGPClusterConfig.
func (c *FakeCiliumBGPClusterConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumbgpclusterconfigsResource, name, pt, data, subresources...), &v2alpha1.CiliumBGPClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPClusterConfig), err
}
