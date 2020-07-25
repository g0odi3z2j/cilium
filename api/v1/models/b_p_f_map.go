// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BPFMap BPF map definition and content
//
// swagger:model BPFMap
type BPFMap struct {

	// Contents of cache
	Cache []*BPFMapEntry `json:"cache"`

	// Path to BPF map
	Path string `json:"path,omitempty"`
}

// Validate validates this b p f map
func (m *BPFMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BPFMap) validateCache(formats strfmt.Registry) error {

	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	for i := 0; i < len(m.Cache); i++ {
		if swag.IsZero(m.Cache[i]) { // not required
			continue
		}

		if m.Cache[i] != nil {
			if err := m.Cache[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cache" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BPFMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BPFMap) UnmarshalBinary(b []byte) error {
	var res BPFMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
