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
)

// NewGetEndpointIDLabelsParams creates a new GetEndpointIDLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEndpointIDLabelsParams() *GetEndpointIDLabelsParams {
	return &GetEndpointIDLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointIDLabelsParamsWithTimeout creates a new GetEndpointIDLabelsParams object
// with the ability to set a timeout on a request.
func NewGetEndpointIDLabelsParamsWithTimeout(timeout time.Duration) *GetEndpointIDLabelsParams {
	return &GetEndpointIDLabelsParams{
		timeout: timeout,
	}
}

// NewGetEndpointIDLabelsParamsWithContext creates a new GetEndpointIDLabelsParams object
// with the ability to set a context for a request.
func NewGetEndpointIDLabelsParamsWithContext(ctx context.Context) *GetEndpointIDLabelsParams {
	return &GetEndpointIDLabelsParams{
		Context: ctx,
	}
}

// NewGetEndpointIDLabelsParamsWithHTTPClient creates a new GetEndpointIDLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEndpointIDLabelsParamsWithHTTPClient(client *http.Client) *GetEndpointIDLabelsParams {
	return &GetEndpointIDLabelsParams{
		HTTPClient: client,
	}
}

/*
GetEndpointIDLabelsParams contains all the parameters to send to the API endpoint

	for the get endpoint ID labels operation.

	Typically these are written to a http.Request.
*/
type GetEndpointIDLabelsParams struct {

	/* ID.

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

// WithDefaults hydrates default values in the get endpoint ID labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointIDLabelsParams) WithDefaults() *GetEndpointIDLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get endpoint ID labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointIDLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) WithTimeout(timeout time.Duration) *GetEndpointIDLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) WithContext(ctx context.Context) *GetEndpointIDLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) WithHTTPClient(client *http.Client) *GetEndpointIDLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) WithID(id string) *GetEndpointIDLabelsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get endpoint ID labels params
func (o *GetEndpointIDLabelsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointIDLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
