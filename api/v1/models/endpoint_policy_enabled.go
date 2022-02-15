// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EndpointPolicyEnabled Whether policy enforcement is enabled (ingress, egress, both or none)
//
// swagger:model EndpointPolicyEnabled
type EndpointPolicyEnabled string

const (

	// EndpointPolicyEnabledNone captures enum value "none"
	EndpointPolicyEnabledNone EndpointPolicyEnabled = "none"

	// EndpointPolicyEnabledIngress captures enum value "ingress"
	EndpointPolicyEnabledIngress EndpointPolicyEnabled = "ingress"

	// EndpointPolicyEnabledEgress captures enum value "egress"
	EndpointPolicyEnabledEgress EndpointPolicyEnabled = "egress"

	// EndpointPolicyEnabledBoth captures enum value "both"
	EndpointPolicyEnabledBoth EndpointPolicyEnabled = "both"

	// EndpointPolicyEnabledAuditIngress captures enum value "audit-ingress"
	EndpointPolicyEnabledAuditIngress EndpointPolicyEnabled = "audit-ingress"

	// EndpointPolicyEnabledAuditEgress captures enum value "audit-egress"
	EndpointPolicyEnabledAuditEgress EndpointPolicyEnabled = "audit-egress"

	// EndpointPolicyEnabledAuditBoth captures enum value "audit-both"
	EndpointPolicyEnabledAuditBoth EndpointPolicyEnabled = "audit-both"
)

// for schema
var endpointPolicyEnabledEnum []interface{}

func init() {
	var res []EndpointPolicyEnabled
	if err := json.Unmarshal([]byte(`["none","ingress","egress","both","audit-ingress","audit-egress","audit-both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointPolicyEnabledEnum = append(endpointPolicyEnabledEnum, v)
	}
}

func (m EndpointPolicyEnabled) validateEndpointPolicyEnabledEnum(path, location string, value EndpointPolicyEnabled) error {
	if err := validate.EnumCase(path, location, value, endpointPolicyEnabledEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this endpoint policy enabled
func (m EndpointPolicyEnabled) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEndpointPolicyEnabledEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
