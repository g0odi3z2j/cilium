/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	networkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// IPAddressesGetter has a method to return a IPAddressInterface.
// A group's client should implement this interface.
type IPAddressesGetter interface {
	IPAddresses() IPAddressInterface
}

// IPAddressInterface has methods to work with IPAddress resources.
type IPAddressInterface interface {
	Create(ctx context.Context, iPAddress *v1alpha1.IPAddress, opts v1.CreateOptions) (*v1alpha1.IPAddress, error)
	Update(ctx context.Context, iPAddress *v1alpha1.IPAddress, opts v1.UpdateOptions) (*v1alpha1.IPAddress, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.IPAddress, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.IPAddressList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPAddress, err error)
	Apply(ctx context.Context, iPAddress *networkingv1alpha1.IPAddressApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.IPAddress, err error)
	IPAddressExpansion
}

// iPAddresses implements IPAddressInterface
type iPAddresses struct {
	client rest.Interface
}

// newIPAddresses returns a IPAddresses
func newIPAddresses(c *NetworkingV1alpha1Client) *iPAddresses {
	return &iPAddresses{
		client: c.RESTClient(),
	}
}

// Get takes name of the iPAddress, and returns the corresponding iPAddress object, and an error if there is any.
func (c *iPAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IPAddress, err error) {
	result = &v1alpha1.IPAddress{}
	err = c.client.Get().
		Resource("ipaddresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPAddresses that match those selectors.
func (c *iPAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPAddressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IPAddressList{}
	err = c.client.Get().
		Resource("ipaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPAddresses.
func (c *iPAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ipaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPAddress and creates it.  Returns the server's representation of the iPAddress, and an error, if there is any.
func (c *iPAddresses) Create(ctx context.Context, iPAddress *v1alpha1.IPAddress, opts v1.CreateOptions) (result *v1alpha1.IPAddress, err error) {
	result = &v1alpha1.IPAddress{}
	err = c.client.Post().
		Resource("ipaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAddress).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPAddress and updates it. Returns the server's representation of the iPAddress, and an error, if there is any.
func (c *iPAddresses) Update(ctx context.Context, iPAddress *v1alpha1.IPAddress, opts v1.UpdateOptions) (result *v1alpha1.IPAddress, err error) {
	result = &v1alpha1.IPAddress{}
	err = c.client.Put().
		Resource("ipaddresses").
		Name(iPAddress.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAddress).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPAddress and deletes it. Returns an error if one occurs.
func (c *iPAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ipaddresses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ipaddresses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPAddress.
func (c *iPAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPAddress, err error) {
	result = &v1alpha1.IPAddress{}
	err = c.client.Patch(pt).
		Resource("ipaddresses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied iPAddress.
func (c *iPAddresses) Apply(ctx context.Context, iPAddress *networkingv1alpha1.IPAddressApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.IPAddress, err error) {
	if iPAddress == nil {
		return nil, fmt.Errorf("iPAddress provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(iPAddress)
	if err != nil {
		return nil, err
	}
	name := iPAddress.Name
	if name == nil {
		return nil, fmt.Errorf("iPAddress.Name must be provided to Apply")
	}
	result = &v1alpha1.IPAddress{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("ipaddresses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
