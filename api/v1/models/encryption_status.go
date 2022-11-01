// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EncryptionStatus Status of transparent encryption
//
// +k8s:deepcopy-gen=true
//
// swagger:model EncryptionStatus
type EncryptionStatus struct {

	// mode
	// Enum: [Disabled IPsec Wireguard]
	Mode string `json:"mode,omitempty"`

	// Human readable status/error/warning message
	Msg string `json:"msg,omitempty"`

	// Status of the Wireguard agent
	Wireguard *WireguardStatus `json:"wireguard,omitempty"`
}

// Validate validates this encryption status
func (m *EncryptionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWireguard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var encryptionStatusTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","IPsec","Wireguard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		encryptionStatusTypeModePropEnum = append(encryptionStatusTypeModePropEnum, v)
	}
}

const (

	// EncryptionStatusModeDisabled captures enum value "Disabled"
	EncryptionStatusModeDisabled string = "Disabled"

	// EncryptionStatusModeIPsec captures enum value "IPsec"
	EncryptionStatusModeIPsec string = "IPsec"

	// EncryptionStatusModeWireguard captures enum value "Wireguard"
	EncryptionStatusModeWireguard string = "Wireguard"
)

// prop value enum
func (m *EncryptionStatus) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, encryptionStatusTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EncryptionStatus) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionStatus) validateWireguard(formats strfmt.Registry) error {
	if swag.IsZero(m.Wireguard) { // not required
		return nil
	}

	if m.Wireguard != nil {
		if err := m.Wireguard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireguard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wireguard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this encryption status based on the context it is used
func (m *EncryptionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWireguard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionStatus) contextValidateWireguard(ctx context.Context, formats strfmt.Registry) error {

	if m.Wireguard != nil {
		if err := m.Wireguard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireguard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wireguard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionStatus) UnmarshalBinary(b []byte) error {
	var res EncryptionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
