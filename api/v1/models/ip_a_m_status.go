// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPAMStatus Status of IP address management
//
// +k8s:deepcopy-gen=true
//
// swagger:model IPAMStatus
type IPAMStatus struct {

	// allocations
	Allocations AllocationMap `json:"allocations,omitempty"`

	// ipv4
	IPV4 []string `json:"ipv4"`

	// ipv6
	IPV6 []string `json:"ipv6"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this IP a m status
func (m *IPAMStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAMStatus) validateAllocations(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocations) { // not required
		return nil
	}

	if m.Allocations != nil {
		if err := m.Allocations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocations")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this IP a m status based on the context it is used
func (m *IPAMStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAMStatus) contextValidateAllocations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Allocations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("allocations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("allocations")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAMStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAMStatus) UnmarshalBinary(b []byte) error {
	var res IPAMStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
