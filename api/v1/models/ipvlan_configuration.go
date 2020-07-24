// Code generated by go-swagger; DO NOT EDIT.

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

// IpvlanConfiguration Setup for datapath when operating in ipvlan mode.
//
// swagger:model IpvlanConfiguration
type IpvlanConfiguration struct {

	// Workload facing ipvlan master device ifindex.
	MasterDeviceIndex int64 `json:"masterDeviceIndex,omitempty"`

	// Mode in which ipvlan setup operates.
	// Enum: [L3 L3S]
	OperationMode string `json:"operationMode,omitempty"`
}

// Validate validates this ipvlan configuration
func (m *IpvlanConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ipvlanConfigurationTypeOperationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L3","L3S"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipvlanConfigurationTypeOperationModePropEnum = append(ipvlanConfigurationTypeOperationModePropEnum, v)
	}
}

const (

	// IpvlanConfigurationOperationModeL3 captures enum value "L3"
	IpvlanConfigurationOperationModeL3 string = "L3"

	// IpvlanConfigurationOperationModeL3S captures enum value "L3S"
	IpvlanConfigurationOperationModeL3S string = "L3S"
)

// prop value enum
func (m *IpvlanConfiguration) validateOperationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipvlanConfigurationTypeOperationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IpvlanConfiguration) validateOperationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationModeEnum("operationMode", "body", m.OperationMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpvlanConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpvlanConfiguration) UnmarshalBinary(b []byte) error {
	var res IpvlanConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
