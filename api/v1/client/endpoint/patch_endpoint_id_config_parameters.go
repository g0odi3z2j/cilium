// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// NewPatchEndpointIDConfigParams creates a new PatchEndpointIDConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEndpointIDConfigParams() *PatchEndpointIDConfigParams {
	return &PatchEndpointIDConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEndpointIDConfigParamsWithTimeout creates a new PatchEndpointIDConfigParams object
// with the ability to set a timeout on a request.
func NewPatchEndpointIDConfigParamsWithTimeout(timeout time.Duration) *PatchEndpointIDConfigParams {
	return &PatchEndpointIDConfigParams{
		timeout: timeout,
	}
}

// NewPatchEndpointIDConfigParamsWithContext creates a new PatchEndpointIDConfigParams object
// with the ability to set a context for a request.
func NewPatchEndpointIDConfigParamsWithContext(ctx context.Context) *PatchEndpointIDConfigParams {
	return &PatchEndpointIDConfigParams{
		Context: ctx,
	}
}

// NewPatchEndpointIDConfigParamsWithHTTPClient creates a new PatchEndpointIDConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEndpointIDConfigParamsWithHTTPClient(client *http.Client) *PatchEndpointIDConfigParams {
	return &PatchEndpointIDConfigParams{
		HTTPClient: client,
	}
}

/*
PatchEndpointIDConfigParams contains all the parameters to send to the API endpoint

	for the patch endpoint ID config operation.

	Typically these are written to a http.Request.
*/
type PatchEndpointIDConfigParams struct {

	// EndpointConfiguration.
	EndpointConfiguration *models.EndpointConfigurationSpec

	/* ID.

	     String describing an endpoint with the format ``[prefix:]id``. If no prefix
	is specified, a prefix of ``cilium-local:`` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - cni-attachment-id: CNI attachment ID, e.g. cni-attachment-id:22222:eth0
	  - container-id: Container runtime ID, e.g. container-id:22222 (deprecated, may not be unique)
	  - container-name: Container name, e.g. container-name:foobar (deprecated, may not be unique)
	  - pod-name: pod name for this container if K8s is enabled, e.g. pod-name:default:foobar (deprecated, may not be unique)
	  - cep-name: cep name for this container if K8s is enabled, e.g. pod-name:default:foobar-net1
	  - docker-endpoint: Docker libnetwork endpoint ID, e.g. docker-endpoint:4444

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch endpoint ID config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEndpointIDConfigParams) WithDefaults() *PatchEndpointIDConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch endpoint ID config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEndpointIDConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) WithTimeout(timeout time.Duration) *PatchEndpointIDConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) WithContext(ctx context.Context) *PatchEndpointIDConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) WithHTTPClient(client *http.Client) *PatchEndpointIDConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointConfiguration adds the endpointConfiguration to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) WithEndpointConfiguration(endpointConfiguration *models.EndpointConfigurationSpec) *PatchEndpointIDConfigParams {
	o.SetEndpointConfiguration(endpointConfiguration)
	return o
}

// SetEndpointConfiguration adds the endpointConfiguration to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) SetEndpointConfiguration(endpointConfiguration *models.EndpointConfigurationSpec) {
	o.EndpointConfiguration = endpointConfiguration
}

// WithID adds the id to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) WithID(id string) *PatchEndpointIDConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch endpoint ID config params
func (o *PatchEndpointIDConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEndpointIDConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.EndpointConfiguration != nil {
		if err := r.SetBodyParam(o.EndpointConfiguration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
