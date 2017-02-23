package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StatusResponse Health and status information of daemon
// swagger:model StatusResponse
type StatusResponse struct {

	// Status of Cilium daemon
	Cilium *Status `json:"cilium,omitempty"`

	// Status of local container runtime
	ContainerRuntime *Status `json:"container-runtime,omitempty"`

	// Status of IP address management
	IPAM *IPAMStatus `json:"ipam,omitempty"`

	// Status of Kubernetes integration
	Kubernetes *Status `json:"kubernetes,omitempty"`

	// Status of key/value datastore
	Kvstore *Status `json:"kvstore,omitempty"`
}

// Validate validates this status response
func (m *StatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCilium(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerRuntime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAM(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKvstore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusResponse) validateCilium(formats strfmt.Registry) error {

	if swag.IsZero(m.Cilium) { // not required
		return nil
	}

	if m.Cilium != nil {

		if err := m.Cilium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateContainerRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerRuntime) { // not required
		return nil
	}

	if m.ContainerRuntime != nil {

		if err := m.ContainerRuntime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container-runtime")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIPAM(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAM) { // not required
		return nil
	}

	if m.IPAM != nil {

		if err := m.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipam")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {

		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKvstore(formats strfmt.Registry) error {

	if swag.IsZero(m.Kvstore) { // not required
		return nil
	}

	if m.Kvstore != nil {

		if err := m.Kvstore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstore")
			}
			return err
		}
	}

	return nil
}
