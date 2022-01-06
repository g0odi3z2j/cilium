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

// EndpointStatus Connectivity status to host cilium-health endpoints via different paths
//
//
// swagger:model EndpointStatus
type EndpointStatus struct {

	// primary address
	PrimaryAddress *PathStatus `json:"primary-address,omitempty"`

	// secondary addresses
	SecondaryAddresses []*PathStatus `json:"secondary-addresses"`
}

// Validate validates this endpoint status
func (m *EndpointStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimaryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointStatus) validatePrimaryAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryAddress) { // not required
		return nil
	}

	if m.PrimaryAddress != nil {
		if err := m.PrimaryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary-address")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateSecondaryAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryAddresses); i++ {
		if swag.IsZero(m.SecondaryAddresses[i]) { // not required
			continue
		}

		if m.SecondaryAddresses[i] != nil {
			if err := m.SecondaryAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondary-addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointStatus) UnmarshalBinary(b []byte) error {
	var res EndpointStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
