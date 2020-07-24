// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KVstoreConfiguration Configuration used for the kvstore
//
// swagger:model KVstoreConfiguration
type KVstoreConfiguration struct {

	// Configuration options
	Options map[string]string `json:"options,omitempty"`

	// Type of kvstore
	Type string `json:"type,omitempty"`
}

// Validate validates this k vstore configuration
func (m *KVstoreConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KVstoreConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KVstoreConfiguration) UnmarshalBinary(b []byte) error {
	var res KVstoreConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
