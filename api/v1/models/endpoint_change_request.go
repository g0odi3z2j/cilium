package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndpointChangeRequest Structure which contains the mutable elements of an Endpoint.
//
// swagger:model EndpointChangeRequest
type EndpointChangeRequest struct {

	// addressing
	Addressing *EndpointAddressing `json:"addressing,omitempty"`

	// ID assigned by container runtime
	ContainerID string `json:"container-id,omitempty"`

	// Docker endpoint ID
	DockerEndpointID string `json:"docker-endpoint-id,omitempty"`

	// Docker network ID
	DockerNetworkID string `json:"docker-network-id,omitempty"`

	// MAC address
	HostMac string `json:"host-mac,omitempty"`

	// Local endpoint ID
	ID int64 `json:"id,omitempty"`

	// Index of network device
	InterfaceIndex int64 `json:"interface-index,omitempty"`

	// Name of network device
	InterfaceName string `json:"interface-name,omitempty"`

	// MAC address
	Mac string `json:"mac,omitempty"`

	// Whether policy enforcement is enabled or not
	PolicyEnabled bool `json:"policy-enabled,omitempty"`

	// Current state of endpoint
	// Required: true
	State EndpointState `json:"state"`
}

// Validate validates this endpoint change request
func (m *EndpointChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointChangeRequest) validateAddressing(formats strfmt.Registry) error {

	if swag.IsZero(m.Addressing) { // not required
		return nil
	}

	if m.Addressing != nil {

		if err := m.Addressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointChangeRequest) UnmarshalBinary(b []byte) error {
	var res EndpointChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
