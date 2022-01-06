// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
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

// NewDeleteFqdnCacheParams creates a new DeleteFqdnCacheParams object
// with the default values initialized.
func NewDeleteFqdnCacheParams() *DeleteFqdnCacheParams {
	var ()
	return &DeleteFqdnCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFqdnCacheParamsWithTimeout creates a new DeleteFqdnCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFqdnCacheParamsWithTimeout(timeout time.Duration) *DeleteFqdnCacheParams {
	var ()
	return &DeleteFqdnCacheParams{

		timeout: timeout,
	}
}

// NewDeleteFqdnCacheParamsWithContext creates a new DeleteFqdnCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFqdnCacheParamsWithContext(ctx context.Context) *DeleteFqdnCacheParams {
	var ()
	return &DeleteFqdnCacheParams{

		Context: ctx,
	}
}

// NewDeleteFqdnCacheParamsWithHTTPClient creates a new DeleteFqdnCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFqdnCacheParamsWithHTTPClient(client *http.Client) *DeleteFqdnCacheParams {
	var ()
	return &DeleteFqdnCacheParams{
		HTTPClient: client,
	}
}

/*DeleteFqdnCacheParams contains all the parameters to send to the API endpoint
for the delete fqdn cache operation typically these are written to a http.Request
*/
type DeleteFqdnCacheParams struct {

	/*Matchpattern
	  A toFQDNs compatible matchPattern expression

	*/
	Matchpattern *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) WithTimeout(timeout time.Duration) *DeleteFqdnCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) WithContext(ctx context.Context) *DeleteFqdnCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) WithHTTPClient(client *http.Client) *DeleteFqdnCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMatchpattern adds the matchpattern to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) WithMatchpattern(matchpattern *string) *DeleteFqdnCacheParams {
	o.SetMatchpattern(matchpattern)
	return o
}

// SetMatchpattern adds the matchpattern to the delete fqdn cache params
func (o *DeleteFqdnCacheParams) SetMatchpattern(matchpattern *string) {
	o.Matchpattern = matchpattern
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFqdnCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Matchpattern != nil {

		// query param matchpattern
		var qrMatchpattern string
		if o.Matchpattern != nil {
			qrMatchpattern = *o.Matchpattern
		}
		qMatchpattern := qrMatchpattern
		if qMatchpattern != "" {
			if err := r.SetQueryParam("matchpattern", qMatchpattern); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
