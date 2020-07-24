// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoadResponse System load on node
//
// swagger:model LoadResponse
type LoadResponse struct {

	// Load average over the past 15 minutes
	Last15min string `json:"last15min,omitempty"`

	// Load average over the past minute
	Last1min string `json:"last1min,omitempty"`

	// Load average over the past 5 minutes
	Last5min string `json:"last5min,omitempty"`
}

// Validate validates this load response
func (m *LoadResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadResponse) UnmarshalBinary(b []byte) error {
	var res LoadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
