// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cilium/cilium/api/v1/models"
)

// NewPatchPrefilterParams creates a new PatchPrefilterParams object
// no default values defined in spec.
func NewPatchPrefilterParams() PatchPrefilterParams {

	return PatchPrefilterParams{}
}

// PatchPrefilterParams contains all the bound params for the patch prefilter operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchPrefilter
type PatchPrefilterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*List of CIDR ranges for filter table
	  Required: true
	  In: body
	*/
	PrefilterSpec *models.PrefilterSpec
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchPrefilterParams() beforehand.
func (o *PatchPrefilterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PrefilterSpec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("prefilterSpec", "body", ""))
			} else {
				res = append(res, errors.NewParseError("prefilterSpec", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.PrefilterSpec = &body
			}
		}
	} else {
		res = append(res, errors.Required("prefilterSpec", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
