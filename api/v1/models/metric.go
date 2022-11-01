// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Metric Metric information
//
// swagger:model Metric
type Metric struct {

	// Labels of the metric
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the metric
	Name string `json:"name,omitempty"`

	// Value of the metric
	Value float64 `json:"value,omitempty"`
}

// Validate validates this metric
func (m *Metric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metric based on context it is used
func (m *Metric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metric) UnmarshalBinary(b []byte) error {
	var res Metric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
