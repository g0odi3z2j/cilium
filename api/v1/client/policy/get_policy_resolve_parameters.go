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

	"github.com/cilium/cilium/api/v1/models"
)

// NewGetPolicyResolveParams creates a new GetPolicyResolveParams object
// with the default values initialized.
func NewGetPolicyResolveParams() *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyResolveParamsWithTimeout creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyResolveParamsWithTimeout(timeout time.Duration) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		timeout: timeout,
	}
}

// NewGetPolicyResolveParamsWithContext creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyResolveParamsWithContext(ctx context.Context) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		Context: ctx,
	}
}

// NewGetPolicyResolveParamsWithHTTPClient creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyResolveParamsWithHTTPClient(client *http.Client) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{
		HTTPClient: client,
	}
}

/*GetPolicyResolveParams contains all the parameters to send to the API endpoint
for the get policy resolve operation typically these are written to a http.Request
*/
type GetPolicyResolveParams struct {

	/*TraceSelector
	  Context to provide policy evaluation on

	*/
	TraceSelector *models.TraceSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy resolve params
func (o *GetPolicyResolveParams) WithTimeout(timeout time.Duration) *GetPolicyResolveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy resolve params
func (o *GetPolicyResolveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy resolve params
func (o *GetPolicyResolveParams) WithContext(ctx context.Context) *GetPolicyResolveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy resolve params
func (o *GetPolicyResolveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy resolve params
func (o *GetPolicyResolveParams) WithHTTPClient(client *http.Client) *GetPolicyResolveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy resolve params
func (o *GetPolicyResolveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTraceSelector adds the traceSelector to the get policy resolve params
func (o *GetPolicyResolveParams) WithTraceSelector(traceSelector *models.TraceSelector) *GetPolicyResolveParams {
	o.SetTraceSelector(traceSelector)
	return o
}

// SetTraceSelector adds the traceSelector to the get policy resolve params
func (o *GetPolicyResolveParams) SetTraceSelector(traceSelector *models.TraceSelector) {
	o.TraceSelector = traceSelector
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyResolveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TraceSelector != nil {
		if err := r.SetBodyParam(o.TraceSelector); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
