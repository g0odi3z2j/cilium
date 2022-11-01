// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

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

// NewGetMapNameParams creates a new GetMapNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMapNameParams() *GetMapNameParams {
	return &GetMapNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMapNameParamsWithTimeout creates a new GetMapNameParams object
// with the ability to set a timeout on a request.
func NewGetMapNameParamsWithTimeout(timeout time.Duration) *GetMapNameParams {
	return &GetMapNameParams{
		timeout: timeout,
	}
}

// NewGetMapNameParamsWithContext creates a new GetMapNameParams object
// with the ability to set a context for a request.
func NewGetMapNameParamsWithContext(ctx context.Context) *GetMapNameParams {
	return &GetMapNameParams{
		Context: ctx,
	}
}

// NewGetMapNameParamsWithHTTPClient creates a new GetMapNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMapNameParamsWithHTTPClient(client *http.Client) *GetMapNameParams {
	return &GetMapNameParams{
		HTTPClient: client,
	}
}

/*
GetMapNameParams contains all the parameters to send to the API endpoint

	for the get map name operation.

	Typically these are written to a http.Request.
*/
type GetMapNameParams struct {

	/* Name.

	   Name of map
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get map name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMapNameParams) WithDefaults() *GetMapNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get map name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMapNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get map name params
func (o *GetMapNameParams) WithTimeout(timeout time.Duration) *GetMapNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get map name params
func (o *GetMapNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get map name params
func (o *GetMapNameParams) WithContext(ctx context.Context) *GetMapNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get map name params
func (o *GetMapNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get map name params
func (o *GetMapNameParams) WithHTTPClient(client *http.Client) *GetMapNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get map name params
func (o *GetMapNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get map name params
func (o *GetMapNameParams) WithName(name string) *GetMapNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get map name params
func (o *GetMapNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetMapNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
