// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
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

// NameManager Internal state about DNS names in relation to policy subsystem
//
// swagger:model NameManager
type NameManager struct {

	// Names to poll for DNS Poller
	DNSPollNames []string `json:"DNSPollNames"`

	// Mapping of FQDNSelectors to corresponding regular expressions
	FQDNPolicySelectors []*SelectorEntry `json:"FQDNPolicySelectors"`
}

// Validate validates this name manager
func (m *NameManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFQDNPolicySelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameManager) validateFQDNPolicySelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.FQDNPolicySelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.FQDNPolicySelectors); i++ {
		if swag.IsZero(m.FQDNPolicySelectors[i]) { // not required
			continue
		}

		if m.FQDNPolicySelectors[i] != nil {
			if err := m.FQDNPolicySelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FQDNPolicySelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameManager) UnmarshalBinary(b []byte) error {
	var res NameManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
