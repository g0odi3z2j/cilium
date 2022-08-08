// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

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

// NewGetPrefilterParams creates a new GetPrefilterParams object
// with the default values initialized.
func NewGetPrefilterParams() *GetPrefilterParams {

	return &GetPrefilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrefilterParamsWithTimeout creates a new GetPrefilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrefilterParamsWithTimeout(timeout time.Duration) *GetPrefilterParams {

	return &GetPrefilterParams{

		timeout: timeout,
	}
}

// NewGetPrefilterParamsWithContext creates a new GetPrefilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrefilterParamsWithContext(ctx context.Context) *GetPrefilterParams {

	return &GetPrefilterParams{

		Context: ctx,
	}
}

// NewGetPrefilterParamsWithHTTPClient creates a new GetPrefilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrefilterParamsWithHTTPClient(client *http.Client) *GetPrefilterParams {

	return &GetPrefilterParams{
		HTTPClient: client,
	}
}

/*
GetPrefilterParams contains all the parameters to send to the API endpoint
for the get prefilter operation typically these are written to a http.Request
*/
type GetPrefilterParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get prefilter params
func (o *GetPrefilterParams) WithTimeout(timeout time.Duration) *GetPrefilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get prefilter params
func (o *GetPrefilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get prefilter params
func (o *GetPrefilterParams) WithContext(ctx context.Context) *GetPrefilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get prefilter params
func (o *GetPrefilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get prefilter params
func (o *GetPrefilterParams) WithHTTPClient(client *http.Client) *GetPrefilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get prefilter params
func (o *GetPrefilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrefilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
