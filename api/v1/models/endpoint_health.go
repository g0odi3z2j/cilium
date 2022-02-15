// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointHealth Health of the endpoint
//
// +deepequal-gen=true
//
// swagger:model EndpointHealth
type EndpointHealth struct {

	// bpf
	Bpf EndpointHealthStatus `json:"bpf,omitempty"`

	// Is this endpoint reachable
	Connected bool `json:"connected,omitempty"`

	// overall health
	OverallHealth EndpointHealthStatus `json:"overallHealth,omitempty"`

	// policy
	Policy EndpointHealthStatus `json:"policy,omitempty"`
}

// Validate validates this endpoint health
func (m *EndpointHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBpf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverallHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointHealth) validateBpf(formats strfmt.Registry) error {

	if swag.IsZero(m.Bpf) { // not required
		return nil
	}

	if err := m.Bpf.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bpf")
		}
		return err
	}

	return nil
}

func (m *EndpointHealth) validateOverallHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallHealth) { // not required
		return nil
	}

	if err := m.OverallHealth.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("overallHealth")
		}
		return err
	}

	return nil
}

func (m *EndpointHealth) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if err := m.Policy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointHealth) UnmarshalBinary(b []byte) error {
	var res EndpointHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
