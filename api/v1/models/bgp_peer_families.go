// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BgpPeerFamilies BGP AFI SAFI state of the peer
//
// +k8s:deepcopy-gen=true
//
// swagger:model BgpPeerFamilies
type BgpPeerFamilies struct {

	// Number of routes accepted from the peer of this address family
	Accepted int64 `json:"accepted,omitempty"`

	// Number of routes advertised of this address family to the peer
	Advertised int64 `json:"advertised,omitempty"`

	// BGP address family indicator
	Afi string `json:"afi,omitempty"`

	// Number of routes received from the peer of this address family
	Received int64 `json:"received,omitempty"`

	// BGP subsequent address family indicator
	Safi string `json:"safi,omitempty"`
}

// Validate validates this bgp peer families
func (m *BgpPeerFamilies) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bgp peer families based on context it is used
func (m *BgpPeerFamilies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BgpPeerFamilies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpPeerFamilies) UnmarshalBinary(b []byte) error {
	var res BgpPeerFamilies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
