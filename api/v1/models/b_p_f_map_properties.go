// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BPFMapProperties BPF map properties
//
// swagger:model BPFMapProperties
type BPFMapProperties struct {

	// Name of the BPF map
	Name string `json:"name,omitempty"`

	// Size of the BPF map
	Size int64 `json:"size,omitempty"`
}

// Validate validates this b p f map properties
func (m *BPFMapProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BPFMapProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BPFMapProperties) UnmarshalBinary(b []byte) error {
	var res BPFMapProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
