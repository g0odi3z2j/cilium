// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TraceTo trace to
// swagger:model TraceTo
type TraceTo struct {

	// List of Layer 4 port and protocol pairs which will be used in communication
	// from the source identity to the destination identity.
	//
	Dports []*Port `json:"dports"`

	// labels
	Labels Labels `json:"labels,omitempty"`
}

// Validate validates this trace to
func (m *TraceTo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceTo) validateDports(formats strfmt.Registry) error {

	if swag.IsZero(m.Dports) { // not required
		return nil
	}

	for i := 0; i < len(m.Dports); i++ {
		if swag.IsZero(m.Dports[i]) { // not required
			continue
		}

		if m.Dports[i] != nil {
			if err := m.Dports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TraceTo) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceTo) UnmarshalBinary(b []byte) error {
	var res TraceTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
