// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BgpPath Single BGP routing Path containing BGP Network Layer Reachability Information (NLRI) and path attributes
//
// swagger:model BgpPath
type BgpPath struct {

	// Age of the path (time since its creation) in nanoseconds
	AgeNanoseconds int64 `json:"age-nanoseconds,omitempty"`

	// True value flags the best path towards the destination prefix
	Best bool `json:"best,omitempty"`

	// Address Family Indicator (AFI) and Subsequent Address Family Indicator (SAFI) of the path
	Family *BgpFamily `json:"family,omitempty"`

	// Network Layer Reachability Information of the path
	Nlri *BgpNlri `json:"nlri,omitempty"`

	// List of BGP path attributes specific for the path
	PathAttributes []*BgpPathAttribute `json:"path-attributes"`

	// True value marks the path as stale
	Stale bool `json:"stale,omitempty"`
}

// Validate validates this bgp path
func (m *BgpPath) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNlri(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPath) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Family) { // not required
		return nil
	}

	if m.Family != nil {
		if err := m.Family.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *BgpPath) validateNlri(formats strfmt.Registry) error {
	if swag.IsZero(m.Nlri) { // not required
		return nil
	}

	if m.Nlri != nil {
		if err := m.Nlri.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nlri")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nlri")
			}
			return err
		}
	}

	return nil
}

func (m *BgpPath) validatePathAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.PathAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.PathAttributes); i++ {
		if swag.IsZero(m.PathAttributes[i]) { // not required
			continue
		}

		if m.PathAttributes[i] != nil {
			if err := m.PathAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("path-attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("path-attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bgp path based on the context it is used
func (m *BgpPath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNlri(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePathAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPath) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if m.Family != nil {
		if err := m.Family.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *BgpPath) contextValidateNlri(ctx context.Context, formats strfmt.Registry) error {

	if m.Nlri != nil {
		if err := m.Nlri.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nlri")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nlri")
			}
			return err
		}
	}

	return nil
}

func (m *BgpPath) contextValidatePathAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PathAttributes); i++ {

		if m.PathAttributes[i] != nil {
			if err := m.PathAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("path-attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("path-attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BgpPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpPath) UnmarshalBinary(b []byte) error {
	var res BgpPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
