// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Endpoint Endpoint
// swagger:model Endpoint

type Endpoint struct {

	// addressing
	Addressing *EndpointAddressing `json:"addressing,omitempty"`

	// ID assigned by container runtime
	ContainerID string `json:"container-id,omitempty"`

	// Name assigned to container
	ContainerName string `json:"container-name,omitempty"`

	// Docker endpoint ID
	DockerEndpointID string `json:"docker-endpoint-id,omitempty"`

	// Docker network ID
	DockerNetworkID string `json:"docker-network-id,omitempty"`

	// Health of the endpoint
	Health *EndpointHealth `json:"health,omitempty"`

	// MAC address
	HostMac string `json:"host-mac,omitempty"`

	// Local endpoint ID
	ID int64 `json:"id,omitempty"`

	// Security identity
	Identity *Identity `json:"identity,omitempty"`

	// Index of network device
	InterfaceIndex int64 `json:"interface-index,omitempty"`

	// Name of network device
	InterfaceName string `json:"interface-name,omitempty"`

	// Labels describing the identity
	Labels *LabelConfiguration `json:"labels,omitempty"`

	// MAC address
	Mac string `json:"mac,omitempty"`

	// K8s pod for this endpoint
	PodName string `json:"pod-name,omitempty"`

	// Policy information of endpoint
	Policy *EndpointPolicy `json:"policy,omitempty"`

	// Whether policy enforcement is enabled or not
	// Required: true
	PolicyEnabled *bool `json:"policy-enabled"`

	// The policy revision this endpoint is running on
	PolicyRevision int64 `json:"policy-revision,omitempty"`

	// Current state of endpoint
	// Required: true
	State EndpointState `json:"state"`

	// Most recent status log. See endpoint/{id}/log for the complete log.
	Status EndpointStatusLog `json:"status"`
}

/* polymorph Endpoint addressing false */

/* polymorph Endpoint container-id false */

/* polymorph Endpoint container-name false */

/* polymorph Endpoint docker-endpoint-id false */

/* polymorph Endpoint docker-network-id false */

/* polymorph Endpoint health false */

/* polymorph Endpoint host-mac false */

/* polymorph Endpoint id false */

/* polymorph Endpoint identity false */

/* polymorph Endpoint interface-index false */

/* polymorph Endpoint interface-name false */

/* polymorph Endpoint labels false */

/* polymorph Endpoint mac false */

/* polymorph Endpoint pod-name false */

/* polymorph Endpoint policy false */

/* polymorph Endpoint policy-enabled false */

/* polymorph Endpoint policy-revision false */

/* polymorph Endpoint state false */

/* polymorph Endpoint status false */

// Validate validates this endpoint
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicyEnabled(formats); err != nil {
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

func (m *Endpoint) validateAddressing(formats strfmt.Registry) error {

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

func (m *Endpoint) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {

		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *Endpoint) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {

		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *Endpoint) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {

		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *Endpoint) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {

		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *Endpoint) validatePolicyEnabled(formats strfmt.Registry) error {

	if err := validate.Required("policy-enabled", "body", m.PolicyEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Endpoint) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
