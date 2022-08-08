// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

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
	"github.com/go-openapi/swag"
)

// NewDeleteServiceIDParams creates a new DeleteServiceIDParams object
// with the default values initialized.
func NewDeleteServiceIDParams() *DeleteServiceIDParams {
	var ()
	return &DeleteServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceIDParamsWithTimeout creates a new DeleteServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServiceIDParamsWithTimeout(timeout time.Duration) *DeleteServiceIDParams {
	var ()
	return &DeleteServiceIDParams{

		timeout: timeout,
	}
}

// NewDeleteServiceIDParamsWithContext creates a new DeleteServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServiceIDParamsWithContext(ctx context.Context) *DeleteServiceIDParams {
	var ()
	return &DeleteServiceIDParams{

		Context: ctx,
	}
}

// NewDeleteServiceIDParamsWithHTTPClient creates a new DeleteServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServiceIDParamsWithHTTPClient(client *http.Client) *DeleteServiceIDParams {
	var ()
	return &DeleteServiceIDParams{
		HTTPClient: client,
	}
}

/*
DeleteServiceIDParams contains all the parameters to send to the API endpoint
for the delete service ID operation typically these are written to a http.Request
*/
type DeleteServiceIDParams struct {

	/*ID
	  ID of service

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service ID params
func (o *DeleteServiceIDParams) WithTimeout(timeout time.Duration) *DeleteServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service ID params
func (o *DeleteServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service ID params
func (o *DeleteServiceIDParams) WithContext(ctx context.Context) *DeleteServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service ID params
func (o *DeleteServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service ID params
func (o *DeleteServiceIDParams) WithHTTPClient(client *http.Client) *DeleteServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service ID params
func (o *DeleteServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete service ID params
func (o *DeleteServiceIDParams) WithID(id int64) *DeleteServiceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete service ID params
func (o *DeleteServiceIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
