// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LabelConfigurationStatus Labels and label configuration of an endpoint
//
// swagger:model LabelConfigurationStatus
type LabelConfigurationStatus struct {

	// All labels derived from the orchestration system
	Derived Labels `json:"derived,omitempty"`

	// Labels derived from orchestration system which have been disabled.
	Disabled Labels `json:"disabled,omitempty"`

	// The current configuration
	Realized *LabelConfigurationSpec `json:"realized,omitempty"`

	// Labels derived from orchestration system that are used in computing a security identity
	SecurityRelevant Labels `json:"security-relevant,omitempty"`
}

// Validate validates this label configuration status
func (m *LabelConfigurationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDerived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityRelevant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelConfigurationStatus) validateDerived(formats strfmt.Registry) error {

	if swag.IsZero(m.Derived) { // not required
		return nil
	}

	if err := m.Derived.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("derived")
		}
		return err
	}

	return nil
}

func (m *LabelConfigurationStatus) validateDisabled(formats strfmt.Registry) error {

	if swag.IsZero(m.Disabled) { // not required
		return nil
	}

	if err := m.Disabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("disabled")
		}
		return err
	}

	return nil
}

func (m *LabelConfigurationStatus) validateRealized(formats strfmt.Registry) error {

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

func (m *LabelConfigurationStatus) validateSecurityRelevant(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityRelevant) { // not required
		return nil
	}

	if err := m.SecurityRelevant.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("security-relevant")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LabelConfigurationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelConfigurationStatus) UnmarshalBinary(b []byte) error {
	var res LabelConfigurationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
