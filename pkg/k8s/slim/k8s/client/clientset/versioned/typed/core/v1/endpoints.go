// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/core/v1"
	scheme "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EndpointsGetter has a method to return a EndpointsInterface.
// A group's client should implement this interface.
type EndpointsGetter interface {
	Endpoints(namespace string) EndpointsInterface
}

// EndpointsInterface has methods to work with Endpoints resources.
type EndpointsInterface interface {
	Create(ctx context.Context, endpoints *v1.Endpoints, opts metav1.CreateOptions) (*v1.Endpoints, error)
	Update(ctx context.Context, endpoints *v1.Endpoints, opts metav1.UpdateOptions) (*v1.Endpoints, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Endpoints, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EndpointsList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoints, err error)
	EndpointsExpansion
}

// endpoints implements EndpointsInterface
type endpoints struct {
	client rest.Interface
	ns     string
}

// newEndpoints returns a Endpoints
func newEndpoints(c *CoreV1Client, namespace string) *endpoints {
	return &endpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *endpoints) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *endpoints) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EndpointsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EndpointsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *endpoints) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Create(ctx context.Context, endpoints *v1.Endpoints, opts metav1.CreateOptions) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(endpoints).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Update(ctx context.Context, endpoints *v1.Endpoints, opts metav1.UpdateOptions) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpoints").
		Name(endpoints.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(endpoints).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *endpoints) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpoints) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched endpoints.
func (c *endpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
