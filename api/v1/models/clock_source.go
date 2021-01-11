// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
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

// ClockSource Status of BPF clock source
//
// +k8s:deepcopy-gen=true
//
// swagger:model ClockSource
type ClockSource struct {

	// Kernel Hz
	Hertz int64 `json:"hertz,omitempty"`

	// Datapath clock source
	// Enum: [ktime jiffies]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this clock source
func (m *ClockSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clockSourceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ktime","jiffies"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clockSourceTypeModePropEnum = append(clockSourceTypeModePropEnum, v)
	}
}

const (

	// ClockSourceModeKtime captures enum value "ktime"
	ClockSourceModeKtime string = "ktime"

	// ClockSourceModeJiffies captures enum value "jiffies"
	ClockSourceModeJiffies string = "jiffies"
)

// prop value enum
func (m *ClockSource) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clockSourceTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClockSource) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClockSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClockSource) UnmarshalBinary(b []byte) error {
	var res ClockSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
