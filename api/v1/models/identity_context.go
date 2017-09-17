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

// IdentityContext Context describing a pair of source and destination identity
// swagger:model IdentityContext

type IdentityContext struct {

	// List of Layer 4 port and protocol pairs which will be used in communication
	// from the source identity to the destination identity.
	//
	Dports []*Port `json:"dports"`

	// from
	From Labels `json:"from"`

	// to
	To Labels `json:"to"`

	// Enable verbose tracing.
	//
	Verbose bool `json:"verbose,omitempty"`
}

/* polymorph IdentityContext dports false */

/* polymorph IdentityContext from false */

/* polymorph IdentityContext to false */

/* polymorph IdentityContext verbose false */

// Validate validates this identity context
func (m *IdentityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDports(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityContext) validateDports(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IdentityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityContext) UnmarshalBinary(b []byte) error {
	var res IdentityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
