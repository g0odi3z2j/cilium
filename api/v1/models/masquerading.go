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

// Masquerading Status of masquerading
//
// +k8s:deepcopy-gen=true
//
// swagger:model Masquerading
type Masquerading struct {

	// Is masquerading enabled
	Enabled bool `json:"enabled,omitempty"`

	// Is BPF ip-masq-agent enabled
	IPMasqAgent bool `json:"ip-masq-agent,omitempty"`

	// mode
	// Enum: [BPF iptables]
	Mode string `json:"mode,omitempty"`

	// SnatExclusionCIDRv4 exempts SNAT from being performed on any packet sent to
	// an IPv4 address that belongs to this CIDR.
	SnatExclusionCidrV4 string `json:"snat-exclusion-cidr-v4,omitempty"`

	// SnatExclusionCIDRv6 exempts SNAT from being performed on any packet sent to
	// an IPv6 address that belongs to this CIDR.
	// For IPv6 we only do masquerading in iptables mode.
	SnatExclusionCidrV6 string `json:"snat-exclusion-cidr-v6,omitempty"`
}

// Validate validates this masquerading
func (m *Masquerading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var masqueradingTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BPF","iptables"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		masqueradingTypeModePropEnum = append(masqueradingTypeModePropEnum, v)
	}
}

const (

	// MasqueradingModeBPF captures enum value "BPF"
	MasqueradingModeBPF string = "BPF"

	// MasqueradingModeIptables captures enum value "iptables"
	MasqueradingModeIptables string = "iptables"
)

// prop value enum
func (m *Masquerading) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, masqueradingTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Masquerading) validateMode(formats strfmt.Registry) error {

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
func (m *Masquerading) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Masquerading) UnmarshalBinary(b []byte) error {
	var res Masquerading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
