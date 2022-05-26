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

// NewGetFqdnCacheParams creates a new GetFqdnCacheParams object
// with the default values initialized.
func NewGetFqdnCacheParams() *GetFqdnCacheParams {
	var ()
	return &GetFqdnCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFqdnCacheParamsWithTimeout creates a new GetFqdnCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFqdnCacheParamsWithTimeout(timeout time.Duration) *GetFqdnCacheParams {
	var ()
	return &GetFqdnCacheParams{

		timeout: timeout,
	}
}

// NewGetFqdnCacheParamsWithContext creates a new GetFqdnCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFqdnCacheParamsWithContext(ctx context.Context) *GetFqdnCacheParams {
	var ()
	return &GetFqdnCacheParams{

		Context: ctx,
	}
}

// NewGetFqdnCacheParamsWithHTTPClient creates a new GetFqdnCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFqdnCacheParamsWithHTTPClient(client *http.Client) *GetFqdnCacheParams {
	var ()
	return &GetFqdnCacheParams{
		HTTPClient: client,
	}
}

/*GetFqdnCacheParams contains all the parameters to send to the API endpoint
for the get fqdn cache operation typically these are written to a http.Request
*/
type GetFqdnCacheParams struct {

	/*Cidr
	  A CIDR range of IPs

	*/
	Cidr *string
	/*Matchpattern
	  A toFQDNs compatible matchPattern expression

	*/
	Matchpattern *string
	/*Source
	  Source from which FQDN entries come from

	*/
	Source *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fqdn cache params
func (o *GetFqdnCacheParams) WithTimeout(timeout time.Duration) *GetFqdnCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fqdn cache params
func (o *GetFqdnCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fqdn cache params
func (o *GetFqdnCacheParams) WithContext(ctx context.Context) *GetFqdnCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fqdn cache params
func (o *GetFqdnCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fqdn cache params
func (o *GetFqdnCacheParams) WithHTTPClient(client *http.Client) *GetFqdnCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fqdn cache params
func (o *GetFqdnCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCidr adds the cidr to the get fqdn cache params
func (o *GetFqdnCacheParams) WithCidr(cidr *string) *GetFqdnCacheParams {
	o.SetCidr(cidr)
	return o
}

// SetCidr adds the cidr to the get fqdn cache params
func (o *GetFqdnCacheParams) SetCidr(cidr *string) {
	o.Cidr = cidr
}

// WithMatchpattern adds the matchpattern to the get fqdn cache params
func (o *GetFqdnCacheParams) WithMatchpattern(matchpattern *string) *GetFqdnCacheParams {
	o.SetMatchpattern(matchpattern)
	return o
}

// SetMatchpattern adds the matchpattern to the get fqdn cache params
func (o *GetFqdnCacheParams) SetMatchpattern(matchpattern *string) {
	o.Matchpattern = matchpattern
}

// WithSource adds the source to the get fqdn cache params
func (o *GetFqdnCacheParams) WithSource(source *string) *GetFqdnCacheParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the get fqdn cache params
func (o *GetFqdnCacheParams) SetSource(source *string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *GetFqdnCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cidr != nil {

		// query param cidr
		var qrCidr string
		if o.Cidr != nil {
			qrCidr = *o.Cidr
		}
		qCidr := qrCidr
		if qCidr != "" {
			if err := r.SetQueryParam("cidr", qCidr); err != nil {
				return err
			}
		}

	}

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

	if o.Source != nil {

		// query param source
		var qrSource string
		if o.Source != nil {
			qrSource = *o.Source
		}
		qSource := qrSource
		if qSource != "" {
			if err := r.SetQueryParam("source", qSource); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
