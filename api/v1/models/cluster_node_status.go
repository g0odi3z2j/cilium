// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNodeStatus Status of cluster
//
// swagger:model ClusterNodeStatus
type ClusterNodeStatus struct {

	// ID that should be used by the client to receive a diff from the previous request
	ClientID int64 `json:"client-id,omitempty"`

	// List of known nodes
	NodesAdded []*NodeElement `json:"nodes-added"`

	// List of known nodes
	NodesRemoved []*NodeElement `json:"nodes-removed"`

	// Name of local node (if available)
	Self string `json:"self,omitempty"`
}

// Validate validates this cluster node status
func (m *ClusterNodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodesAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesRemoved(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNodeStatus) validateNodesAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesAdded) { // not required
		return nil
	}

	for i := 0; i < len(m.NodesAdded); i++ {
		if swag.IsZero(m.NodesAdded[i]) { // not required
			continue
		}

		if m.NodesAdded[i] != nil {
			if err := m.NodesAdded[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes-added" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterNodeStatus) validateNodesRemoved(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesRemoved) { // not required
		return nil
	}

	for i := 0; i < len(m.NodesRemoved); i++ {
		if swag.IsZero(m.NodesRemoved[i]) { // not required
			continue
		}

		if m.NodesRemoved[i] != nil {
			if err := m.NodesRemoved[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes-removed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodeStatus) UnmarshalBinary(b []byte) error {
	var res ClusterNodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
