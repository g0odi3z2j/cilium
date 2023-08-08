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

// NewDeleteEndpointParams creates a new DeleteEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEndpointParams() *DeleteEndpointParams {
	return &DeleteEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEndpointParamsWithTimeout creates a new DeleteEndpointParams object
// with the ability to set a timeout on a request.
func NewDeleteEndpointParamsWithTimeout(timeout time.Duration) *DeleteEndpointParams {
	return &DeleteEndpointParams{
		timeout: timeout,
	}
}

// NewDeleteEndpointParamsWithContext creates a new DeleteEndpointParams object
// with the ability to set a context for a request.
func NewDeleteEndpointParamsWithContext(ctx context.Context) *DeleteEndpointParams {
	return &DeleteEndpointParams{
		Context: ctx,
	}
}

// NewDeleteEndpointParamsWithHTTPClient creates a new DeleteEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEndpointParamsWithHTTPClient(client *http.Client) *DeleteEndpointParams {
	return &DeleteEndpointParams{
		HTTPClient: client,
	}
}

/*
DeleteEndpointParams contains all the parameters to send to the API endpoint

	for the delete endpoint operation.

	Typically these are written to a http.Request.
*/
type DeleteEndpointParams struct {

	// Endpoint.
	Endpoint *models.EndpointBatchDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointParams) WithDefaults() *DeleteEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete endpoint params
func (o *DeleteEndpointParams) WithTimeout(timeout time.Duration) *DeleteEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete endpoint params
func (o *DeleteEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete endpoint params
func (o *DeleteEndpointParams) WithContext(ctx context.Context) *DeleteEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete endpoint params
func (o *DeleteEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete endpoint params
func (o *DeleteEndpointParams) WithHTTPClient(client *http.Client) *DeleteEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete endpoint params
func (o *DeleteEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpoint adds the endpoint to the delete endpoint params
func (o *DeleteEndpointParams) WithEndpoint(endpoint *models.EndpointBatchDeleteRequest) *DeleteEndpointParams {
	o.SetEndpoint(endpoint)
	return o
}

// SetEndpoint adds the endpoint to the delete endpoint params
func (o *DeleteEndpointParams) SetEndpoint(endpoint *models.EndpointBatchDeleteRequest) {
	o.Endpoint = endpoint
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Endpoint != nil {
		if err := r.SetBodyParam(o.Endpoint); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
