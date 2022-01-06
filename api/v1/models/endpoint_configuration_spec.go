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

// EndpointConfigurationSpec An endpoint's configuration
//
// swagger:model EndpointConfigurationSpec
type EndpointConfigurationSpec struct {

	// the endpoint's labels
	LabelConfiguration *LabelConfigurationSpec `json:"label-configuration,omitempty"`

	// Changeable configuration
	Options ConfigurationMap `json:"options,omitempty"`
}

// Validate validates this endpoint configuration spec
func (m *EndpointConfigurationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointConfigurationSpec) validateLabelConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.LabelConfiguration) { // not required
		return nil
	}

	if m.LabelConfiguration != nil {
		if err := m.LabelConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label-configuration")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointConfigurationSpec) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if err := m.Options.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("options")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointConfigurationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointConfigurationSpec) UnmarshalBinary(b []byte) error {
	var res EndpointConfigurationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
