// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
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

// HostRouting Status of host routing
//
// +k8s:deepcopy-gen=true
//
// swagger:model HostRouting
type HostRouting struct {

	// Datapath routing mode
	// Enum: [BPF Legacy]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this host routing
func (m *HostRouting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hostRoutingTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BPF","Legacy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostRoutingTypeModePropEnum = append(hostRoutingTypeModePropEnum, v)
	}
}

const (

	// HostRoutingModeBPF captures enum value "BPF"
	HostRoutingModeBPF string = "BPF"

	// HostRoutingModeLegacy captures enum value "Legacy"
	HostRoutingModeLegacy string = "Legacy"
)

// prop value enum
func (m *HostRouting) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostRoutingTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HostRouting) validateMode(formats strfmt.Registry) error {

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
func (m *HostRouting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostRouting) UnmarshalBinary(b []byte) error {
	var res HostRouting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
