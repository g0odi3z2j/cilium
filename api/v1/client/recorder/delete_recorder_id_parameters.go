// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

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

// NewDeleteRecorderIDParams creates a new DeleteRecorderIDParams object
// with the default values initialized.
func NewDeleteRecorderIDParams() *DeleteRecorderIDParams {
	var ()
	return &DeleteRecorderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecorderIDParamsWithTimeout creates a new DeleteRecorderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecorderIDParamsWithTimeout(timeout time.Duration) *DeleteRecorderIDParams {
	var ()
	return &DeleteRecorderIDParams{

		timeout: timeout,
	}
}

// NewDeleteRecorderIDParamsWithContext creates a new DeleteRecorderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecorderIDParamsWithContext(ctx context.Context) *DeleteRecorderIDParams {
	var ()
	return &DeleteRecorderIDParams{

		Context: ctx,
	}
}

// NewDeleteRecorderIDParamsWithHTTPClient creates a new DeleteRecorderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecorderIDParamsWithHTTPClient(client *http.Client) *DeleteRecorderIDParams {
	var ()
	return &DeleteRecorderIDParams{
		HTTPClient: client,
	}
}

/*DeleteRecorderIDParams contains all the parameters to send to the API endpoint
for the delete recorder ID operation typically these are written to a http.Request
*/
type DeleteRecorderIDParams struct {

	/*ID
	  ID of recorder

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recorder ID params
func (o *DeleteRecorderIDParams) WithTimeout(timeout time.Duration) *DeleteRecorderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recorder ID params
func (o *DeleteRecorderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recorder ID params
func (o *DeleteRecorderIDParams) WithContext(ctx context.Context) *DeleteRecorderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recorder ID params
func (o *DeleteRecorderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recorder ID params
func (o *DeleteRecorderIDParams) WithHTTPClient(client *http.Client) *DeleteRecorderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recorder ID params
func (o *DeleteRecorderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete recorder ID params
func (o *DeleteRecorderIDParams) WithID(id int64) *DeleteRecorderIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete recorder ID params
func (o *DeleteRecorderIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecorderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
