// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewPostIpamIPParams creates a new PostIpamIPParams object
//
// There are no default values defined in the spec.
func NewPostIpamIPParams() PostIpamIPParams {

	return PostIpamIPParams{}
}

// PostIpamIPParams contains all the bound params for the post ipam IP operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostIpamIP
type PostIpamIPParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*IP address
	  Required: true
	  In: path
	*/
	IP string
	/*
	  In: query
	*/
	Owner *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostIpamIPParams() beforehand.
func (o *PostIpamIPParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rIP, rhkIP, _ := route.Params.GetOK("ip")
	if err := o.bindIP(rIP, rhkIP, route.Formats); err != nil {
		res = append(res, err)
	}

	qOwner, qhkOwner, _ := qs.GetOK("owner")
	if err := o.bindOwner(qOwner, qhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIP binds and validates parameter IP from path.
func (o *PostIpamIPParams) bindIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IP = raw

	return nil
}

// bindOwner binds and validates parameter Owner from query.
func (o *PostIpamIPParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Owner = &raw

	return nil
}
