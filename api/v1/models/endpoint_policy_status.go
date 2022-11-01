// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointPolicyStatus Policy information of an endpoint
//
// swagger:model EndpointPolicyStatus
type EndpointPolicyStatus struct {

	// The policy revision currently enforced in the proxy for this endpoint
	ProxyPolicyRevision int64 `json:"proxy-policy-revision,omitempty"`

	// Statistics of the proxy redirects configured for this endpoint
	ProxyStatistics []*ProxyStatistics `json:"proxy-statistics"`

	// The policy in the datapath for this endpoint
	Realized *EndpointPolicy `json:"realized,omitempty"`

	// The policy that should apply to this endpoint
	Spec *EndpointPolicy `json:"spec,omitempty"`
}

// Validate validates this endpoint policy status
func (m *EndpointPolicyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxyStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointPolicyStatus) validateProxyStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyStatistics) { // not required
		return nil
	}

	for i := 0; i < len(m.ProxyStatistics); i++ {
		if swag.IsZero(m.ProxyStatistics[i]) { // not required
			continue
		}

		if m.ProxyStatistics[i] != nil {
			if err := m.ProxyStatistics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxy-statistics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxy-statistics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EndpointPolicyStatus) validateRealized(formats strfmt.Registry) error {
	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {
		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointPolicyStatus) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoint policy status based on the context it is used
func (m *EndpointPolicyStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProxyStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRealized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointPolicyStatus) contextValidateProxyStatistics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProxyStatistics); i++ {

		if m.ProxyStatistics[i] != nil {
			if err := m.ProxyStatistics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxy-statistics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxy-statistics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EndpointPolicyStatus) contextValidateRealized(ctx context.Context, formats strfmt.Registry) error {

	if m.Realized != nil {
		if err := m.Realized.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointPolicyStatus) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointPolicyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointPolicyStatus) UnmarshalBinary(b []byte) error {
	var res EndpointPolicyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
