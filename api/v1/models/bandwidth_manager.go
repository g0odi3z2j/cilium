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

// BandwidthManager Status of bandwidth manager
//
// +k8s:deepcopy-gen=true
//
// swagger:model BandwidthManager
type BandwidthManager struct {

	// congestion control
	// Enum: [cubic bbr]
	CongestionControl string `json:"congestionControl,omitempty"`

	// devices
	Devices []string `json:"devices"`

	// Is bandwidth manager enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this bandwidth manager
func (m *BandwidthManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCongestionControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bandwidthManagerTypeCongestionControlPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cubic","bbr"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bandwidthManagerTypeCongestionControlPropEnum = append(bandwidthManagerTypeCongestionControlPropEnum, v)
	}
}

const (

	// BandwidthManagerCongestionControlCubic captures enum value "cubic"
	BandwidthManagerCongestionControlCubic string = "cubic"

	// BandwidthManagerCongestionControlBbr captures enum value "bbr"
	BandwidthManagerCongestionControlBbr string = "bbr"
)

// prop value enum
func (m *BandwidthManager) validateCongestionControlEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bandwidthManagerTypeCongestionControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BandwidthManager) validateCongestionControl(formats strfmt.Registry) error {
	if swag.IsZero(m.CongestionControl) { // not required
		return nil
	}

	// value enum
	if err := m.validateCongestionControlEnum("congestionControl", "body", m.CongestionControl); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bandwidth manager based on context it is used
func (m *BandwidthManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BandwidthManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BandwidthManager) UnmarshalBinary(b []byte) error {
	var res BandwidthManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
