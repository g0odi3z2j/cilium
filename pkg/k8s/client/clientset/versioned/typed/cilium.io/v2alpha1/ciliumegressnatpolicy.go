// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumEgressNATPoliciesGetter has a method to return a CiliumEgressNATPolicyInterface.
// A group's client should implement this interface.
type CiliumEgressNATPoliciesGetter interface {
	CiliumEgressNATPolicies() CiliumEgressNATPolicyInterface
}

// CiliumEgressNATPolicyInterface has methods to work with CiliumEgressNATPolicy resources.
type CiliumEgressNATPolicyInterface interface {
	Create(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.CreateOptions) (*v2alpha1.CiliumEgressNATPolicy, error)
	Update(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.UpdateOptions) (*v2alpha1.CiliumEgressNATPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumEgressNATPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumEgressNATPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEgressNATPolicy, err error)
	CiliumEgressNATPolicyExpansion
}

// ciliumEgressNATPolicies implements CiliumEgressNATPolicyInterface
type ciliumEgressNATPolicies struct {
	client rest.Interface
}

// newCiliumEgressNATPolicies returns a CiliumEgressNATPolicies
func newCiliumEgressNATPolicies(c *CiliumV2alpha1Client) *ciliumEgressNATPolicies {
	return &ciliumEgressNATPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumEgressNATPolicy, and returns the corresponding ciliumEgressNATPolicy object, and an error if there is any.
func (c *ciliumEgressNATPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	result = &v2alpha1.CiliumEgressNATPolicy{}
	err = c.client.Get().
		Resource("ciliumegressnatpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumEgressNATPolicies that match those selectors.
func (c *ciliumEgressNATPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumEgressNATPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.CiliumEgressNATPolicyList{}
	err = c.client.Get().
		Resource("ciliumegressnatpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumEgressNATPolicies.
func (c *ciliumEgressNATPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumegressnatpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumEgressNATPolicy and creates it.  Returns the server's representation of the ciliumEgressNATPolicy, and an error, if there is any.
func (c *ciliumEgressNATPolicies) Create(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.CreateOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	result = &v2alpha1.CiliumEgressNATPolicy{}
	err = c.client.Post().
		Resource("ciliumegressnatpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressNATPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumEgressNATPolicy and updates it. Returns the server's representation of the ciliumEgressNATPolicy, and an error, if there is any.
func (c *ciliumEgressNATPolicies) Update(ctx context.Context, ciliumEgressNATPolicy *v2alpha1.CiliumEgressNATPolicy, opts v1.UpdateOptions) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	result = &v2alpha1.CiliumEgressNATPolicy{}
	err = c.client.Put().
		Resource("ciliumegressnatpolicies").
		Name(ciliumEgressNATPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressNATPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumEgressNATPolicy and deletes it. Returns an error if one occurs.
func (c *ciliumEgressNATPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumegressnatpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumEgressNATPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumegressnatpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumEgressNATPolicy.
func (c *ciliumEgressNATPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEgressNATPolicy, err error) {
	result = &v2alpha1.CiliumEgressNATPolicy{}
	err = c.client.Patch(pt).
		Resource("ciliumegressnatpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
