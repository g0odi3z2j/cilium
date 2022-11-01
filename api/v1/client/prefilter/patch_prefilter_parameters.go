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

	"github.com/cilium/cilium/api/v1/models"
)

// NewPatchPrefilterParams creates a new PatchPrefilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPrefilterParams() *PatchPrefilterParams {
	return &PatchPrefilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPrefilterParamsWithTimeout creates a new PatchPrefilterParams object
// with the ability to set a timeout on a request.
func NewPatchPrefilterParamsWithTimeout(timeout time.Duration) *PatchPrefilterParams {
	return &PatchPrefilterParams{
		timeout: timeout,
	}
}

// NewPatchPrefilterParamsWithContext creates a new PatchPrefilterParams object
// with the ability to set a context for a request.
func NewPatchPrefilterParamsWithContext(ctx context.Context) *PatchPrefilterParams {
	return &PatchPrefilterParams{
		Context: ctx,
	}
}

// NewPatchPrefilterParamsWithHTTPClient creates a new PatchPrefilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPrefilterParamsWithHTTPClient(client *http.Client) *PatchPrefilterParams {
	return &PatchPrefilterParams{
		HTTPClient: client,
	}
}

/*
PatchPrefilterParams contains all the parameters to send to the API endpoint

	for the patch prefilter operation.

	Typically these are written to a http.Request.
*/
type PatchPrefilterParams struct {

	/* PrefilterSpec.

	   List of CIDR ranges for filter table
	*/
	PrefilterSpec *models.PrefilterSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch prefilter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPrefilterParams) WithDefaults() *PatchPrefilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch prefilter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPrefilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch prefilter params
func (o *PatchPrefilterParams) WithTimeout(timeout time.Duration) *PatchPrefilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch prefilter params
func (o *PatchPrefilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch prefilter params
func (o *PatchPrefilterParams) WithContext(ctx context.Context) *PatchPrefilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch prefilter params
func (o *PatchPrefilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch prefilter params
func (o *PatchPrefilterParams) WithHTTPClient(client *http.Client) *PatchPrefilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch prefilter params
func (o *PatchPrefilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrefilterSpec adds the prefilterSpec to the patch prefilter params
func (o *PatchPrefilterParams) WithPrefilterSpec(prefilterSpec *models.PrefilterSpec) *PatchPrefilterParams {
	o.SetPrefilterSpec(prefilterSpec)
	return o
}

// SetPrefilterSpec adds the prefilterSpec to the patch prefilter params
func (o *PatchPrefilterParams) SetPrefilterSpec(prefilterSpec *models.PrefilterSpec) {
	o.PrefilterSpec = prefilterSpec
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPrefilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PrefilterSpec != nil {
		if err := r.SetBodyParam(o.PrefilterSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
