// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceSpec Configuration of a service
//
// swagger:model ServiceSpec
type ServiceSpec struct {

	// List of backend addresses
	BackendAddresses []*BackendAddress `json:"backend-addresses"`

	// flags
	Flags *ServiceSpecFlags `json:"flags,omitempty"`

	// Frontend address
	// Required: true
	FrontendAddress *FrontendAddress `json:"frontend-address"`

	// Unique identification
	ID int64 `json:"id,omitempty"`
}

// Validate validates this service spec
func (m *ServiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceSpec) validateBackendAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.BackendAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.BackendAddresses); i++ {
		if swag.IsZero(m.BackendAddresses[i]) { // not required
			continue
		}

		if m.BackendAddresses[i] != nil {
			if err := m.BackendAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backend-addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceSpec) validateFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	if m.Flags != nil {
		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceSpec) validateFrontendAddress(formats strfmt.Registry) error {

	if err := validate.Required("frontend-address", "body", m.FrontendAddress); err != nil {
		return err
	}

	if m.FrontendAddress != nil {
		if err := m.FrontendAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontend-address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSpec) UnmarshalBinary(b []byte) error {
	var res ServiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceSpecFlags Optional service configuration flags
//
// swagger:model ServiceSpecFlags
type ServiceSpecFlags struct {

	// Service health check node port
	HealthCheckNodePort uint16 `json:"healthCheckNodePort,omitempty"`

	// Service name  (e.g. Kubernetes service name)
	Name string `json:"name,omitempty"`

	// Service namespace  (e.g. Kubernetes namespace)
	Namespace string `json:"namespace,omitempty"`

	// Service traffic policy
	// Enum: [Cluster Local]
	TrafficPolicy string `json:"trafficPolicy,omitempty"`

	// Service type
	// Enum: [ClusterIP NodePort ExternalIPs HostPort LoadBalancer]
	Type string `json:"type,omitempty"`
}

// Validate validates this service spec flags
func (m *ServiceSpecFlags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrafficPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceSpecFlagsTypeTrafficPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cluster","Local"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceSpecFlagsTypeTrafficPolicyPropEnum = append(serviceSpecFlagsTypeTrafficPolicyPropEnum, v)
	}
}

const (

	// ServiceSpecFlagsTrafficPolicyCluster captures enum value "Cluster"
	ServiceSpecFlagsTrafficPolicyCluster string = "Cluster"

	// ServiceSpecFlagsTrafficPolicyLocal captures enum value "Local"
	ServiceSpecFlagsTrafficPolicyLocal string = "Local"
)

// prop value enum
func (m *ServiceSpecFlags) validateTrafficPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceSpecFlagsTypeTrafficPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceSpecFlags) validateTrafficPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.TrafficPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrafficPolicyEnum("flags"+"."+"trafficPolicy", "body", m.TrafficPolicy); err != nil {
		return err
	}

	return nil
}

var serviceSpecFlagsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ClusterIP","NodePort","ExternalIPs","HostPort","LoadBalancer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceSpecFlagsTypeTypePropEnum = append(serviceSpecFlagsTypeTypePropEnum, v)
	}
}

const (

	// ServiceSpecFlagsTypeClusterIP captures enum value "ClusterIP"
	ServiceSpecFlagsTypeClusterIP string = "ClusterIP"

	// ServiceSpecFlagsTypeNodePort captures enum value "NodePort"
	ServiceSpecFlagsTypeNodePort string = "NodePort"

	// ServiceSpecFlagsTypeExternalIPs captures enum value "ExternalIPs"
	ServiceSpecFlagsTypeExternalIPs string = "ExternalIPs"

	// ServiceSpecFlagsTypeHostPort captures enum value "HostPort"
	ServiceSpecFlagsTypeHostPort string = "HostPort"

	// ServiceSpecFlagsTypeLoadBalancer captures enum value "LoadBalancer"
	ServiceSpecFlagsTypeLoadBalancer string = "LoadBalancer"
)

// prop value enum
func (m *ServiceSpecFlags) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceSpecFlagsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceSpecFlags) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("flags"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSpecFlags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSpecFlags) UnmarshalBinary(b []byte) error {
	var res ServiceSpecFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
