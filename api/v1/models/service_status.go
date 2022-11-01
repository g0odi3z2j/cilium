// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceStatus Configuration of a service
//
// swagger:model ServiceStatus
type ServiceStatus struct {

	// realized
	Realized *ServiceSpec `json:"realized,omitempty"`
}

// Validate validates this service status
func (m *ServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceStatus) validateRealized(formats strfmt.Registry) error {
	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {
		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service status based on the context it is used
func (m *ServiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRealized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceStatus) contextValidateRealized(ctx context.Context, formats strfmt.Registry) error {

	if m.Realized != nil {
		if err := m.Realized.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceStatus) UnmarshalBinary(b []byte) error {
	var res ServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
