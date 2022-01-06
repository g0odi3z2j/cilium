// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecorderMaskStatus Configuration of a recorder mask
//
// swagger:model RecorderMaskStatus
type RecorderMaskStatus struct {

	// realized
	Realized *RecorderMaskSpec `json:"realized,omitempty"`
}

// Validate validates this recorder mask status
func (m *RecorderMaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecorderMaskStatus) validateRealized(formats strfmt.Registry) error {

	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {
		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecorderMaskStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecorderMaskStatus) UnmarshalBinary(b []byte) error {
	var res RecorderMaskStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
