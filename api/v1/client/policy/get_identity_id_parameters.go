// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

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

// NewGetIdentityIDParams creates a new GetIdentityIDParams object
// with the default values initialized.
func NewGetIdentityIDParams() *GetIdentityIDParams {
	var ()
	return &GetIdentityIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityIDParamsWithTimeout creates a new GetIdentityIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityIDParamsWithTimeout(timeout time.Duration) *GetIdentityIDParams {
	var ()
	return &GetIdentityIDParams{

		timeout: timeout,
	}
}

// NewGetIdentityIDParamsWithContext creates a new GetIdentityIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityIDParamsWithContext(ctx context.Context) *GetIdentityIDParams {
	var ()
	return &GetIdentityIDParams{

		Context: ctx,
	}
}

// NewGetIdentityIDParamsWithHTTPClient creates a new GetIdentityIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityIDParamsWithHTTPClient(client *http.Client) *GetIdentityIDParams {
	var ()
	return &GetIdentityIDParams{
		HTTPClient: client,
	}
}

/*GetIdentityIDParams contains all the parameters to send to the API endpoint
for the get identity ID operation typically these are written to a http.Request
*/
type GetIdentityIDParams struct {

	/*ID
	  Cluster wide unique identifier of a security identity.


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identity ID params
func (o *GetIdentityIDParams) WithTimeout(timeout time.Duration) *GetIdentityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity ID params
func (o *GetIdentityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity ID params
func (o *GetIdentityIDParams) WithContext(ctx context.Context) *GetIdentityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity ID params
func (o *GetIdentityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity ID params
func (o *GetIdentityIDParams) WithHTTPClient(client *http.Client) *GetIdentityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity ID params
func (o *GetIdentityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get identity ID params
func (o *GetIdentityIDParams) WithID(id string) *GetIdentityIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get identity ID params
func (o *GetIdentityIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
