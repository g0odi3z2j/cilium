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

// NewPatchEndpointIDLabelsParams creates a new PatchEndpointIDLabelsParams object
// with the default values initialized.
func NewPatchEndpointIDLabelsParams() *PatchEndpointIDLabelsParams {
	var ()
	return &PatchEndpointIDLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEndpointIDLabelsParamsWithTimeout creates a new PatchEndpointIDLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEndpointIDLabelsParamsWithTimeout(timeout time.Duration) *PatchEndpointIDLabelsParams {
	var ()
	return &PatchEndpointIDLabelsParams{

		timeout: timeout,
	}
}

// NewPatchEndpointIDLabelsParamsWithContext creates a new PatchEndpointIDLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchEndpointIDLabelsParamsWithContext(ctx context.Context) *PatchEndpointIDLabelsParams {
	var ()
	return &PatchEndpointIDLabelsParams{

		Context: ctx,
	}
}

// NewPatchEndpointIDLabelsParamsWithHTTPClient creates a new PatchEndpointIDLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchEndpointIDLabelsParamsWithHTTPClient(client *http.Client) *PatchEndpointIDLabelsParams {
	var ()
	return &PatchEndpointIDLabelsParams{
		HTTPClient: client,
	}
}

/*
PatchEndpointIDLabelsParams contains all the parameters to send to the API endpoint
for the patch endpoint ID labels operation typically these are written to a http.Request
*/
type PatchEndpointIDLabelsParams struct {

	/*Configuration*/
	Configuration *models.LabelConfigurationSpec
	/*ID
	  String describing an endpoint with the format ``[prefix:]id``. If no prefix
	is specified, a prefix of ``cilium-local:`` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - container-id: Container runtime ID, e.g. container-id:22222
	  - container-name: Container name, e.g. container-name:foobar
	  - pod-name: pod name for this container if K8s is enabled, e.g. pod-name:default:foobar
	  - docker-endpoint: Docker libnetwork endpoint ID, e.g. docker-endpoint:4444


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) WithTimeout(timeout time.Duration) *PatchEndpointIDLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) WithContext(ctx context.Context) *PatchEndpointIDLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) WithHTTPClient(client *http.Client) *PatchEndpointIDLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) WithConfiguration(configuration *models.LabelConfigurationSpec) *PatchEndpointIDLabelsParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) SetConfiguration(configuration *models.LabelConfigurationSpec) {
	o.Configuration = configuration
}

// WithID adds the id to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) WithID(id string) *PatchEndpointIDLabelsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch endpoint ID labels params
func (o *PatchEndpointIDLabelsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEndpointIDLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration != nil {
		if err := r.SetBodyParam(o.Configuration); err != nil {
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
