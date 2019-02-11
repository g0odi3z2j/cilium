// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BPFMapEntry BPF map cache entry"
// swagger:model BPFMapEntry
type BPFMapEntry struct {

	// Desired action to be performed
	// Enum: [ok insert delete]
	DesiredAction string `json:"desired-action,omitempty"`

	// Key of map entry
	Key string `json:"key,omitempty"`

	// Last error seen while performing desired action
	LastError string `json:"last-error,omitempty"`

	// Value of map entry
	Value string `json:"value,omitempty"`
}

// Validate validates this b p f map entry
func (m *BPFMapEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bPFMapEntryTypeDesiredActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","insert","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bPFMapEntryTypeDesiredActionPropEnum = append(bPFMapEntryTypeDesiredActionPropEnum, v)
	}
}

const (

	// BPFMapEntryDesiredActionOk captures enum value "ok"
	BPFMapEntryDesiredActionOk string = "ok"

	// BPFMapEntryDesiredActionInsert captures enum value "insert"
	BPFMapEntryDesiredActionInsert string = "insert"

	// BPFMapEntryDesiredActionDelete captures enum value "delete"
	BPFMapEntryDesiredActionDelete string = "delete"
)

// prop value enum
func (m *BPFMapEntry) validateDesiredActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bPFMapEntryTypeDesiredActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BPFMapEntry) validateDesiredAction(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDesiredActionEnum("desired-action", "body", m.DesiredAction); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BPFMapEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BPFMapEntry) UnmarshalBinary(b []byte) error {
	var res BPFMapEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
