// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumEgressGatewayPoliciesGetter has a method to return a CiliumEgressGatewayPolicyInterface.
// A group's client should implement this interface.
type CiliumEgressGatewayPoliciesGetter interface {
	CiliumEgressGatewayPolicies() CiliumEgressGatewayPolicyInterface
}

// CiliumEgressGatewayPolicyInterface has methods to work with CiliumEgressGatewayPolicy resources.
type CiliumEgressGatewayPolicyInterface interface {
	Create(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.CreateOptions) (*v2.CiliumEgressGatewayPolicy, error)
	Update(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.UpdateOptions) (*v2.CiliumEgressGatewayPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.CiliumEgressGatewayPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.CiliumEgressGatewayPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressGatewayPolicy, err error)
	CiliumEgressGatewayPolicyExpansion
}

// ciliumEgressGatewayPolicies implements CiliumEgressGatewayPolicyInterface
type ciliumEgressGatewayPolicies struct {
	client rest.Interface
}

// newCiliumEgressGatewayPolicies returns a CiliumEgressGatewayPolicies
func newCiliumEgressGatewayPolicies(c *CiliumV2Client) *ciliumEgressGatewayPolicies {
	return &ciliumEgressGatewayPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumEgressGatewayPolicy, and returns the corresponding ciliumEgressGatewayPolicy object, and an error if there is any.
func (c *ciliumEgressGatewayPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	result = &v2.CiliumEgressGatewayPolicy{}
	err = c.client.Get().
		Resource("ciliumegressgatewaypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumEgressGatewayPolicies that match those selectors.
func (c *ciliumEgressGatewayPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumEgressGatewayPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumEgressGatewayPolicyList{}
	err = c.client.Get().
		Resource("ciliumegressgatewaypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumEgressGatewayPolicies.
func (c *ciliumEgressGatewayPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumegressgatewaypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumEgressGatewayPolicy and creates it.  Returns the server's representation of the ciliumEgressGatewayPolicy, and an error, if there is any.
func (c *ciliumEgressGatewayPolicies) Create(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.CreateOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	result = &v2.CiliumEgressGatewayPolicy{}
	err = c.client.Post().
		Resource("ciliumegressgatewaypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressGatewayPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumEgressGatewayPolicy and updates it. Returns the server's representation of the ciliumEgressGatewayPolicy, and an error, if there is any.
func (c *ciliumEgressGatewayPolicies) Update(ctx context.Context, ciliumEgressGatewayPolicy *v2.CiliumEgressGatewayPolicy, opts v1.UpdateOptions) (result *v2.CiliumEgressGatewayPolicy, err error) {
	result = &v2.CiliumEgressGatewayPolicy{}
	err = c.client.Put().
		Resource("ciliumegressgatewaypolicies").
		Name(ciliumEgressGatewayPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressGatewayPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumEgressGatewayPolicy and deletes it. Returns an error if one occurs.
func (c *ciliumEgressGatewayPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumegressgatewaypolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumEgressGatewayPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumegressgatewaypolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumEgressGatewayPolicy.
func (c *ciliumEgressGatewayPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressGatewayPolicy, err error) {
	result = &v2.CiliumEgressGatewayPolicy{}
	err = c.client.Patch(pt).
		Resource("ciliumegressgatewaypolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
