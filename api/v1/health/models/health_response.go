// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

	if err := m.validateSystemLoad(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
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
