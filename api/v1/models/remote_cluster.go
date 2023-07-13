// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoteCluster Status of remote cluster
//
// +k8s:deepcopy-gen=true
//
// swagger:model RemoteCluster
type RemoteCluster struct {

	// Cluster configuration exposed by the remote cluster
	Config *RemoteClusterConfig `json:"config,omitempty"`

	// Indicates whether the connection to the remote kvstore is established
	Connected bool `json:"connected,omitempty"`

	// Time of last failure that occurred while attempting to reach the cluster
	// Format: date-time
	LastFailure strfmt.DateTime `json:"last-failure,omitempty"`

	// Name of the cluster
	Name string `json:"name,omitempty"`

	// Number of endpoints in the cluster
	NumEndpoints int64 `json:"num-endpoints,omitempty"`

	// Number of failures reaching the cluster
	NumFailures int64 `json:"num-failures,omitempty"`

	// Number of identities in the cluster
	NumIdentities int64 `json:"num-identities,omitempty"`

	// Number of nodes in the cluster
	NumNodes int64 `json:"num-nodes,omitempty"`

	// Number of services in the cluster
	NumSharedServices int64 `json:"num-shared-services,omitempty"`

	// Indicates readiness of the remote cluster
	Ready bool `json:"ready,omitempty"`

	// Status of the control plane
	Status string `json:"status,omitempty"`

	// Synchronization status about each resource type
	Synced *RemoteClusterSynced `json:"synced,omitempty"`
}

// Validate validates this remote cluster
func (m *RemoteCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynced(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteCluster) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteCluster) validateLastFailure(formats strfmt.Registry) error {
	if swag.IsZero(m.LastFailure) { // not required
		return nil
	}

	if err := validate.FormatOf("last-failure", "body", "date-time", m.LastFailure.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RemoteCluster) validateSynced(formats strfmt.Registry) error {
	if swag.IsZero(m.Synced) { // not required
		return nil
	}

	if m.Synced != nil {
		if err := m.Synced.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synced")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("synced")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this remote cluster based on the context it is used
func (m *RemoteCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSynced(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteCluster) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteCluster) contextValidateSynced(ctx context.Context, formats strfmt.Registry) error {

	if m.Synced != nil {
		if err := m.Synced.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synced")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("synced")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteCluster) UnmarshalBinary(b []byte) error {
	var res RemoteCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
