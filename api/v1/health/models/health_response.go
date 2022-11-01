// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	ciliumModels "github.com/cilium/cilium/api/v1/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthResponse Health and status information of local node
//
// swagger:model HealthResponse
type HealthResponse struct {

	// Status of Cilium daemon
	Cilium ciliumModels.StatusResponse `json:"cilium,omitempty"`

	// System load on node
	SystemLoad *LoadResponse `json:"system-load,omitempty"`

	// Uptime of cilium-health instance
	Uptime string `json:"uptime,omitempty"`
}

// Validate validates this health response
func (m *HealthResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCilium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemLoad(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthResponse) validateCilium(formats strfmt.Registry) error {
	if swag.IsZero(m.Cilium) { // not required
		return nil
	}

	if err := m.Cilium.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cilium")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cilium")
		}
		return err
	}

	return nil
}

func (m *HealthResponse) validateSystemLoad(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemLoad) { // not required
		return nil
	}

	if m.SystemLoad != nil {
		if err := m.SystemLoad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system-load")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system-load")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this health response based on the context it is used
func (m *HealthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCilium(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemLoad(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthResponse) contextValidateCilium(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Cilium.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cilium")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cilium")
		}
		return err
	}

	return nil
}

func (m *HealthResponse) contextValidateSystemLoad(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemLoad != nil {
		if err := m.SystemLoad.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system-load")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system-load")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthResponse) UnmarshalBinary(b []byte) error {
	var res HealthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
