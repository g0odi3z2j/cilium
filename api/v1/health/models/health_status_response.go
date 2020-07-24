// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthStatusResponse Connectivity status to other daemons
//
// swagger:model HealthStatusResponse
type HealthStatusResponse struct {

	// Description of the local node
	Local *SelfStatus `json:"local,omitempty"`

	// Connectivity status to each other node
	Nodes []*NodeStatus `json:"nodes"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this health status response
func (m *HealthStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthStatusResponse) validateLocal(formats strfmt.Registry) error {

	if swag.IsZero(m.Local) { // not required
		return nil
	}

	if m.Local != nil {
		if err := m.Local.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local")
			}
			return err
		}
	}

	return nil
}

func (m *HealthStatusResponse) validateNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthStatusResponse) UnmarshalBinary(b []byte) error {
	var res HealthStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
