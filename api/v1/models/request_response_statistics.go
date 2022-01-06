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

// RequestResponseStatistics Statistics of a proxy redirect
//
// +k8s:deepcopy-gen=true
//
// swagger:model RequestResponseStatistics
type RequestResponseStatistics struct {

	// requests
	Requests *MessageForwardingStatistics `json:"requests,omitempty"`

	// responses
	Responses *MessageForwardingStatistics `json:"responses,omitempty"`
}

// Validate validates this request response statistics
func (m *RequestResponseStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestResponseStatistics) validateRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	if m.Requests != nil {
		if err := m.Requests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requests")
			}
			return err
		}
	}

	return nil
}

func (m *RequestResponseStatistics) validateResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	if m.Responses != nil {
		if err := m.Responses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responses")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestResponseStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestResponseStatistics) UnmarshalBinary(b []byte) error {
	var res RequestResponseStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
