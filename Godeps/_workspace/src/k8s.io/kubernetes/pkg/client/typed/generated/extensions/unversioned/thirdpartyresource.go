/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	api "github.com/noironetworks/cilium-net/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	watch "k8s.io/kubernetes/pkg/watch"
)

// ThirdPartyResourceNamespacer has methods to work with ThirdPartyResource resources in a namespace
type ThirdPartyResourceNamespacer interface {
	ThirdPartyResources(namespace string) ThirdPartyResourceInterface
}

// ThirdPartyResourceInterface has methods to work with ThirdPartyResource resources.
type ThirdPartyResourceInterface interface {
	Create(*extensions.ThirdPartyResource) (*extensions.ThirdPartyResource, error)
	Update(*extensions.ThirdPartyResource) (*extensions.ThirdPartyResource, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*extensions.ThirdPartyResource, error)
	List(opts api.ListOptions) (*extensions.ThirdPartyResourceList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
}

// thirdPartyResources implements ThirdPartyResourceInterface
type thirdPartyResources struct {
	client *ExtensionsClient
	ns     string
}

// newThirdPartyResources returns a ThirdPartyResources
func newThirdPartyResources(c *ExtensionsClient, namespace string) *thirdPartyResources {
	return &thirdPartyResources{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a thirdPartyResource and creates it.  Returns the server's representation of the thirdPartyResource, and an error, if there is any.
func (c *thirdPartyResources) Create(thirdPartyResource *extensions.ThirdPartyResource) (result *extensions.ThirdPartyResource, err error) {
	result = &extensions.ThirdPartyResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("thirdPartyResources").
		Body(thirdPartyResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a thirdPartyResource and updates it. Returns the server's representation of the thirdPartyResource, and an error, if there is any.
func (c *thirdPartyResources) Update(thirdPartyResource *extensions.ThirdPartyResource) (result *extensions.ThirdPartyResource, err error) {
	result = &extensions.ThirdPartyResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("thirdPartyResources").
		Name(thirdPartyResource.Name).
		Body(thirdPartyResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the thirdPartyResource and deletes it. Returns an error if one occurs.
func (c *thirdPartyResources) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("thirdPartyResources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *thirdPartyResources) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		NamespaceIfScoped(c.ns, len(c.ns) > 0).
		Resource("thirdPartyResources").
		VersionedParams(&listOptions, api.Scheme).
		Body(options).
		Do().
		Error()
}

// Get takes name of the thirdPartyResource, and returns the corresponding thirdPartyResource object, and an error if there is any.
func (c *thirdPartyResources) Get(name string) (result *extensions.ThirdPartyResource, err error) {
	result = &extensions.ThirdPartyResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("thirdPartyResources").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ThirdPartyResources that match those selectors.
func (c *thirdPartyResources) List(opts api.ListOptions) (result *extensions.ThirdPartyResourceList, err error) {
	result = &extensions.ThirdPartyResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("thirdPartyResources").
		VersionedParams(&opts, api.Scheme).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested thirdPartyResources.
func (c *thirdPartyResources) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("thirdPartyResources").
		VersionedParams(&opts, api.Scheme).
		Watch()
}
