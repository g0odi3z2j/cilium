// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DatapathMode Datapath mode
//
// swagger:model DatapathMode
type DatapathMode string

const (

	// DatapathModeVeth captures enum value "veth"
	DatapathModeVeth DatapathMode = "veth"

	// DatapathModeIpvlan captures enum value "ipvlan"
	DatapathModeIpvlan DatapathMode = "ipvlan"
)

// for schema
var datapathModeEnum []interface{}

func init() {
	var res []DatapathMode
	if err := json.Unmarshal([]byte(`["veth","ipvlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datapathModeEnum = append(datapathModeEnum, v)
	}
}

func (m DatapathMode) validateDatapathModeEnum(path, location string, value DatapathMode) error {
	if err := validate.EnumCase(path, location, value, datapathModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datapath mode
func (m DatapathMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatapathModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
