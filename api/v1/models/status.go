// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Status Status of an individual component
//
// swagger:model Status
type Status struct {

	// Human readable status/error/warning message
	Msg string `json:"msg,omitempty"`

	// State the component is in
	// Enum: [Ok Warning Failure Disabled]
	State string `json:"state,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ok","Warning","Failure","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusTypeStatePropEnum = append(statusTypeStatePropEnum, v)
	}
}

const (

	// StatusStateOk captures enum value "Ok"
	StatusStateOk string = "Ok"

	// StatusStateWarning captures enum value "Warning"
	StatusStateWarning string = "Warning"

	// StatusStateFailure captures enum value "Failure"
	StatusStateFailure string = "Failure"

	// StatusStateDisabled captures enum value "Disabled"
	StatusStateDisabled string = "Disabled"
)

// prop value enum
func (m *Status) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Status) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
