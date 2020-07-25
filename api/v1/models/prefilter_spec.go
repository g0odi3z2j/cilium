// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrefilterSpec CIDR ranges implemented in the Prefilter
//
// swagger:model PrefilterSpec
type PrefilterSpec struct {

	// deny
	Deny []string `json:"deny"`

	// revision
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this prefilter spec
func (m *PrefilterSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrefilterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrefilterSpec) UnmarshalBinary(b []byte) error {
	var res PrefilterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
