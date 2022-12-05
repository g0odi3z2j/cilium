/*
Copyright 2018 The Kubernetes Authors.

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

package client

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"

	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// WarningHandlerOptions are options for configuring a
// warning handler for the client which is responsible
// for surfacing API Server warnings.
type WarningHandlerOptions struct {
	// SuppressWarnings decides if the warnings from the
	// API server are suppressed or surfaced in the client.
	SuppressWarnings bool
	// AllowDuplicateLogs does not deduplicate the to-be
	// logged surfaced warnings messages. See
	// log.WarningHandlerOptions for considerations
	// regarding deduplication
	AllowDuplicateLogs bool
}

// Options are creation options for a Client.
type Options struct {
	// Scheme, if provided, will be used to map go structs to GroupVersionKinds
	Scheme *runtime.Scheme

	// Mapper, if provided, will be used to map GroupVersionKinds to Resources
	Mapper meta.RESTMapper

	// Opts is used to configure the warning handler responsible for
	// surfacing and handling warnings messages sent by the API server.
	Opts WarningHandlerOptions
}

// New returns a new Client using the provided config and Options.
// The returned client reads *and* writes directly from the server
// (it doesn't use object caches).  It understands how to work with
// normal types (both custom resources and aggregated/built-in resources),
// as well as unstructured types.
//
// In the case of normal types, the scheme will be used to look up the
// corresponding group, version, and kind for the given type.  In the
// case of unstructured types, the group, version, and kind will be extracted
// from the corresponding fields on the object.
func New(config *rest.Config, options Options) (Client, error) {
	return newClient(config, options)
}

func newClient(config *rest.Config, options Options) (*client, error) {
	if config == nil {
		return nil, fmt.Errorf("must provide non-nil rest.Config to client.New")
	}

	if !options.Opts.SuppressWarnings {
		// surface warnings
		logger := log.Log.WithName("KubeAPIWarningLogger")
		// Set a WarningHandler, the default WarningHandler
		// is log.KubeAPIWarningLogger with deduplication enabled.
		// See log.KubeAPIWarningLoggerOptions for considerations
		// regarding deduplication.
		config = rest.CopyConfig(config)
		config.WarningHandler = log.NewKubeAPIWarningLogger(
			logger,
			log.KubeAPIWarningLoggerOptions{
				Deduplicate: !options.Opts.AllowDuplicateLogs,
			},
		)
	}

	// Init a scheme if none provided
	if options.Scheme == nil {
		options.Scheme = scheme.Scheme
	}

	// Init a Mapper if none provided
	if options.Mapper == nil {
		var err error
		options.Mapper, err = apiutil.NewDynamicRESTMapper(config)
		if err != nil {
			return nil, err
		}
	}

	clientcache := &clientCache{
		config: config,
		scheme: options.Scheme,
		mapper: options.Mapper,
		codecs: serializer.NewCodecFactory(options.Scheme),

		structuredResourceByType:   make(map[schema.GroupVersionKind]*resourceMeta),
		unstructuredResourceByType: make(map[schema.GroupVersionKind]*resourceMeta),
	}

	rawMetaClient, err := metadata.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to construct metadata-only client for use as part of client: %w", err)
	}

	c := &client{
		typedClient: typedClient{
			cache:      clientcache,
			paramCodec: runtime.NewParameterCodec(options.Scheme),
		},
		unstructuredClient: unstructuredClient{
			cache:      clientcache,
			paramCodec: noConversionParamCodec{},
		},
		metadataClient: metadataClient{
			client:     rawMetaClient,
			restMapper: options.Mapper,
		},
		scheme: options.Scheme,
		mapper: options.Mapper,
	}

	return c, nil
}

var _ Client = &client{}

// client is a client.Client that reads and writes directly from/to an API server.  It lazily initializes
// new clients at the time they are used, and caches the client.
type client struct {
	typedClient        typedClient
	unstructuredClient unstructuredClient
	metadataClient     metadataClient
	scheme             *runtime.Scheme
	mapper             meta.RESTMapper
}

// resetGroupVersionKind is a helper function to restore and preserve GroupVersionKind on an object.
func (c *client) resetGroupVersionKind(obj runtime.Object, gvk schema.GroupVersionKind) {
	if gvk != schema.EmptyObjectKind.GroupVersionKind() {
		if v, ok := obj.(schema.ObjectKind); ok {
			v.SetGroupVersionKind(gvk)
		}
	}
}

// Scheme returns the scheme this client is using.
func (c *client) Scheme() *runtime.Scheme {
	return c.scheme
}

// RESTMapper returns the scheme this client is using.
func (c *client) RESTMapper() meta.RESTMapper {
	return c.mapper
}

// Create implements client.Client.
func (c *client) Create(ctx context.Context, obj Object, opts ...CreateOption) error {
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.Create(ctx, obj, opts...)
	case *metav1.PartialObjectMetadata:
		return fmt.Errorf("cannot create using only metadata")
	default:
		return c.typedClient.Create(ctx, obj, opts...)
	}
}

// Update implements client.Client.
func (c *client) Update(ctx context.Context, obj Object, opts ...UpdateOption) error {
	defer c.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.Update(ctx, obj, opts...)
	case *metav1.PartialObjectMetadata:
		return fmt.Errorf("cannot update using only metadata -- did you mean to patch?")
	default:
		return c.typedClient.Update(ctx, obj, opts...)
	}
}

// Delete implements client.Client.
func (c *client) Delete(ctx context.Context, obj Object, opts ...DeleteOption) error {
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.Delete(ctx, obj, opts...)
	case *metav1.PartialObjectMetadata:
		return c.metadataClient.Delete(ctx, obj, opts...)
	default:
		return c.typedClient.Delete(ctx, obj, opts...)
	}
}

// DeleteAllOf implements client.Client.
func (c *client) DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error {
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.DeleteAllOf(ctx, obj, opts...)
	case *metav1.PartialObjectMetadata:
		return c.metadataClient.DeleteAllOf(ctx, obj, opts...)
	default:
		return c.typedClient.DeleteAllOf(ctx, obj, opts...)
	}
}

// Patch implements client.Client.
func (c *client) Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
	defer c.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.Patch(ctx, obj, patch, opts...)
	case *metav1.PartialObjectMetadata:
		return c.metadataClient.Patch(ctx, obj, patch, opts...)
	default:
		return c.typedClient.Patch(ctx, obj, patch, opts...)
	}
}

// Get implements client.Client.
func (c *client) Get(ctx context.Context, key ObjectKey, obj Object, opts ...GetOption) error {
	switch obj.(type) {
	case *unstructured.Unstructured:
		return c.unstructuredClient.Get(ctx, key, obj, opts...)
	case *metav1.PartialObjectMetadata:
		// Metadata only object should always preserve the GVK coming in from the caller.
		defer c.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
		return c.metadataClient.Get(ctx, key, obj, opts...)
	default:
		return c.typedClient.Get(ctx, key, obj, opts...)
	}
}

// List implements client.Client.
func (c *client) List(ctx context.Context, obj ObjectList, opts ...ListOption) error {
	switch x := obj.(type) {
	case *unstructured.UnstructuredList:
		return c.unstructuredClient.List(ctx, obj, opts...)
	case *metav1.PartialObjectMetadataList:
		// Metadata only object should always preserve the GVK.
		gvk := obj.GetObjectKind().GroupVersionKind()
		defer c.resetGroupVersionKind(obj, gvk)

		// Call the list client.
		if err := c.metadataClient.List(ctx, obj, opts...); err != nil {
			return err
		}

		// Restore the GVK for each item in the list.
		itemGVK := schema.GroupVersionKind{
			Group:   gvk.Group,
			Version: gvk.Version,
			// TODO: this is producing unsafe guesses that don't actually work,
			// but it matches ~99% of the cases out there.
			Kind: strings.TrimSuffix(gvk.Kind, "List"),
		}
		for i := range x.Items {
			item := &x.Items[i]
			item.SetGroupVersionKind(itemGVK)
		}

		return nil
	default:
		return c.typedClient.List(ctx, obj, opts...)
	}
}

// Status implements client.StatusClient.
func (c *client) Status() SubResourceWriter {
	return c.SubResource("status")
}

func (c *client) SubResource(subResource string) SubResourceWriter {
	return &subResourceWriter{client: c, subResource: subResource}
}

// subResourceWriter is client.SubResourceWriter that writes to subresources.
type subResourceWriter struct {
	client      *client
	subResource string
}

// ensure subResourceWriter implements client.SubResourceWriter.
var _ SubResourceWriter = &subResourceWriter{}

// SubResourceUpdateOptions holds all the possible configuration
// for a subresource update request.
type SubResourceUpdateOptions struct {
	UpdateOptions
	SubResourceBody Object
}

// ApplyToSubResourceUpdate updates the configuration on the given create options
func (uo *SubResourceUpdateOptions) ApplyToSubResourceUpdate(o *SubResourceUpdateOptions) {
	uo.UpdateOptions.ApplyToUpdate(&o.UpdateOptions)
	if uo.SubResourceBody != nil {
		o.SubResourceBody = uo.SubResourceBody
	}
}

// ApplyOptions applies the given options.
func (uo *SubResourceUpdateOptions) ApplyOptions(opts []SubResourceUpdateOption) *SubResourceUpdateOptions {
	for _, o := range opts {
		o.ApplyToSubResourceUpdate(uo)
	}

	return uo
}

// SubResourceUpdateAndPatchOption is an option that can be used for either
// a subresource update or patch request.
type SubResourceUpdateAndPatchOption interface {
	SubResourceUpdateOption
	SubResourcePatchOption
}

// WithSubResourceBody returns an option that uses the given body
// for a subresource Update or Patch operation.
func WithSubResourceBody(body Object) SubResourceUpdateAndPatchOption {
	return &withSubresourceBody{body: body}
}

type withSubresourceBody struct {
	body Object
}

func (wsr *withSubresourceBody) ApplyToSubResourceUpdate(o *SubResourceUpdateOptions) {
	o.SubResourceBody = wsr.body
}

func (wsr *withSubresourceBody) ApplyToSubResourcePatch(o *SubResourcePatchOptions) {
	o.SubResourceBody = wsr.body
}

// SubResourceCreateOptions are all the possible configurations for a subresource
// create request.
type SubResourceCreateOptions struct {
	CreateOptions
}

// ApplyOptions applies the given options.
func (co *SubResourceCreateOptions) ApplyOptions(opts []SubResourceCreateOption) *SubResourceCreateOptions {
	for _, o := range opts {
		o.ApplyToSubResourceCreate(co)
	}

	return co
}

// ApplyToSubresourceCreate applies the the configuration on the given create options.
func (co *SubResourceCreateOptions) ApplyToSubresourceCreate(o *SubResourceCreateOptions) {
	co.CreateOptions.ApplyToCreate(&co.CreateOptions)
}

// SubResourcePatchOptions holds all possible configurations for a subresource patch
// request.
type SubResourcePatchOptions struct {
	PatchOptions
	SubResourceBody Object
}

// ApplyOptions applies the given options.
func (po *SubResourcePatchOptions) ApplyOptions(opts []SubResourcePatchOption) *SubResourcePatchOptions {
	for _, o := range opts {
		o.ApplyToSubResourcePatch(po)
	}

	return po
}

// ApplyToSubResourcePatch applies the configuration on the given patch options.
func (po *SubResourcePatchOptions) ApplyToSubResourcePatch(o *SubResourcePatchOptions) {
	po.PatchOptions.ApplyToPatch(&o.PatchOptions)
	if po.SubResourceBody != nil {
		o.SubResourceBody = po.SubResourceBody
	}
}

func (sw *subResourceWriter) Create(ctx context.Context, obj Object, subResource Object, opts ...SubResourceCreateOption) error {
	defer sw.client.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
	defer sw.client.resetGroupVersionKind(subResource, subResource.GetObjectKind().GroupVersionKind())

	switch obj.(type) {
	case *unstructured.Unstructured:
		return sw.client.unstructuredClient.CreateSubResource(ctx, obj, subResource, sw.subResource, opts...)
	case *metav1.PartialObjectMetadata:
		return fmt.Errorf("cannot update status using only metadata -- did you mean to patch?")
	default:
		return sw.client.typedClient.CreateSubResource(ctx, obj, subResource, sw.subResource, opts...)
	}
}

// Update implements client.SubResourceWriter.
func (sw *subResourceWriter) Update(ctx context.Context, obj Object, opts ...SubResourceUpdateOption) error {
	defer sw.client.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
	switch obj.(type) {
	case *unstructured.Unstructured:
		return sw.client.unstructuredClient.UpdateSubResource(ctx, obj, sw.subResource, opts...)
	case *metav1.PartialObjectMetadata:
		return fmt.Errorf("cannot update status using only metadata -- did you mean to patch?")
	default:
		return sw.client.typedClient.UpdateSubResource(ctx, obj, sw.subResource, opts...)
	}
}

// Patch implements client.SubResourceWriter.
func (sw *subResourceWriter) Patch(ctx context.Context, obj Object, patch Patch, opts ...SubResourcePatchOption) error {
	defer sw.client.resetGroupVersionKind(obj, obj.GetObjectKind().GroupVersionKind())
	switch obj.(type) {
	case *unstructured.Unstructured:
		return sw.client.unstructuredClient.PatchSubResource(ctx, obj, sw.subResource, patch, opts...)
	case *metav1.PartialObjectMetadata:
		return sw.client.metadataClient.PatchSubResource(ctx, obj, sw.subResource, patch, opts...)
	default:
		return sw.client.typedClient.PatchSubResource(ctx, obj, sw.subResource, patch, opts...)
	}
}
