// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BgpRoute Single BGP route retrieved from the RIB of underlying router
//
// swagger:model BgpRoute
type BgpRoute struct {

	// List of routing paths leading towards the prefix
	Paths []*BgpPath `json:"paths"`

	// IP prefix of the route
	Prefix string `json:"prefix,omitempty"`

	// Autonomous System Number (ASN) identifying a BGP virtual router instance
	RouterAsn int64 `json:"router-asn,omitempty"`
}

// Validate validates this bgp route
func (m *BgpRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaths(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpRoute) validatePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.Paths) { // not required
		return nil
	}

	for i := 0; i < len(m.Paths); i++ {
		if swag.IsZero(m.Paths[i]) { // not required
			continue
		}

		if m.Paths[i] != nil {
			if err := m.Paths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bgp route based on the context it is used
func (m *BgpRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpRoute) contextValidatePaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Paths); i++ {

		if m.Paths[i] != nil {
			if err := m.Paths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BgpRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpRoute) UnmarshalBinary(b []byte) error {
	var res BgpRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
