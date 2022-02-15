// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

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

// NewDeleteIpamIPParams creates a new DeleteIpamIPParams object
// with the default values initialized.
func NewDeleteIpamIPParams() *DeleteIpamIPParams {
	var ()
	return &DeleteIpamIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIpamIPParamsWithTimeout creates a new DeleteIpamIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIpamIPParamsWithTimeout(timeout time.Duration) *DeleteIpamIPParams {
	var ()
	return &DeleteIpamIPParams{

		timeout: timeout,
	}
}

// NewDeleteIpamIPParamsWithContext creates a new DeleteIpamIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIpamIPParamsWithContext(ctx context.Context) *DeleteIpamIPParams {
	var ()
	return &DeleteIpamIPParams{

		Context: ctx,
	}
}

// NewDeleteIpamIPParamsWithHTTPClient creates a new DeleteIpamIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIpamIPParamsWithHTTPClient(client *http.Client) *DeleteIpamIPParams {
	var ()
	return &DeleteIpamIPParams{
		HTTPClient: client,
	}
}

/*DeleteIpamIPParams contains all the parameters to send to the API endpoint
for the delete ipam IP operation typically these are written to a http.Request
*/
type DeleteIpamIPParams struct {

	/*IP
	  IP address or owner name

	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ipam IP params
func (o *DeleteIpamIPParams) WithTimeout(timeout time.Duration) *DeleteIpamIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ipam IP params
func (o *DeleteIpamIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ipam IP params
func (o *DeleteIpamIPParams) WithContext(ctx context.Context) *DeleteIpamIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ipam IP params
func (o *DeleteIpamIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ipam IP params
func (o *DeleteIpamIPParams) WithHTTPClient(client *http.Client) *DeleteIpamIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ipam IP params
func (o *DeleteIpamIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the delete ipam IP params
func (o *DeleteIpamIPParams) WithIP(ip string) *DeleteIpamIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the delete ipam IP params
func (o *DeleteIpamIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIpamIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
