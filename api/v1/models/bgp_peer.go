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
	"github.com/go-openapi/validate"
)

// BgpPeer State of a BGP Peer
//
// +k8s:deepcopy-gen=true
//
// swagger:model BgpPeer
type BgpPeer struct {

	// Applied initial value for the BGP HoldTimer (RFC 4271, Section 4.2) in seconds.
	// The applied value holds the value that is in effect on the current BGP session.
	//
	AppliedHoldTimeSeconds int64 `json:"applied-hold-time-seconds,omitempty"`

	// Applied initial value for the BGP KeepaliveTimer (RFC 4271, Section 8) in seconds.
	// The applied value holds the value that is in effect on the current BGP session.
	//
	AppliedKeepAliveTimeSeconds int64 `json:"applied-keep-alive-time-seconds,omitempty"`

	// Configured initial value for the BGP HoldTimer (RFC 4271, Section 4.2) in seconds.
	// The configured value will be used for negotiation with the peer during the BGP session establishment.
	//
	ConfiguredHoldTimeSeconds int64 `json:"configured-hold-time-seconds,omitempty"`

	// Configured initial value for the BGP KeepaliveTimer (RFC 4271, Section 8) in seconds.
	// The applied value may be different than the configured value, as it depends on the negotiated hold time interval.
	//
	ConfiguredKeepAliveTimeSeconds int64 `json:"configured-keep-alive-time-seconds,omitempty"`

	// Initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8) in seconds
	ConnectRetryTimeSeconds int64 `json:"connect-retry-time-seconds,omitempty"`

	// Time To Live (TTL) value used in BGP packets sent to the eBGP neighbor.
	// 1 implies that eBGP multi-hop feature is disabled (only a single hop is allowed).
	//
	EbgpMultihopTTL int64 `json:"ebgp-multihop-ttl,omitempty"`

	// BGP peer address family state
	Families []*BgpPeerFamilies `json:"families"`

	// Graceful restart capability
	GracefulRestart *BgpGracefulRestart `json:"graceful-restart,omitempty"`

	// Local AS Number
	LocalAsn int64 `json:"local-asn,omitempty"`

	// IP Address of peer
	PeerAddress string `json:"peer-address,omitempty"`

	// Peer AS Number
	PeerAsn int64 `json:"peer-asn,omitempty"`

	// TCP port number of peer
	// Maximum: 65535
	// Minimum: 1
	PeerPort int64 `json:"peer-port,omitempty"`

	// BGP peer operational state as described here
	// https://www.rfc-editor.org/rfc/rfc4271#section-8.2.2
	//
	SessionState string `json:"session-state,omitempty"`

	// BGP peer connection uptime in nano seconds.
	UptimeNanoseconds int64 `json:"uptime-nanoseconds,omitempty"`
}

// Validate validates this bgp peer
func (m *BgpPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFamilies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGracefulRestart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPeer) validateFamilies(formats strfmt.Registry) error {
	if swag.IsZero(m.Families) { // not required
		return nil
	}

	for i := 0; i < len(m.Families); i++ {
		if swag.IsZero(m.Families[i]) { // not required
			continue
		}

		if m.Families[i] != nil {
			if err := m.Families[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("families" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("families" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BgpPeer) validateGracefulRestart(formats strfmt.Registry) error {
	if swag.IsZero(m.GracefulRestart) { // not required
		return nil
	}

	if m.GracefulRestart != nil {
		if err := m.GracefulRestart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("graceful-restart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("graceful-restart")
			}
			return err
		}
	}

	return nil
}

func (m *BgpPeer) validatePeerPort(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("peer-port", "body", m.PeerPort, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("peer-port", "body", m.PeerPort, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bgp peer based on the context it is used
func (m *BgpPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFamilies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGracefulRestart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPeer) contextValidateFamilies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Families); i++ {

		if m.Families[i] != nil {
			if err := m.Families[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("families" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("families" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BgpPeer) contextValidateGracefulRestart(ctx context.Context, formats strfmt.Registry) error {

	if m.GracefulRestart != nil {
		if err := m.GracefulRestart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("graceful-restart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("graceful-restart")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BgpPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpPeer) UnmarshalBinary(b []byte) error {
	var res BgpPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
