package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// EndpointState State of endpoint
// swagger:model EndpointState
type EndpointState string

const (
	// EndpointStateCreating captures enum value "creating"
	EndpointStateCreating EndpointState = "creating"
	// EndpointStateDisconnected captures enum value "disconnected"
	EndpointStateDisconnected EndpointState = "disconnected"
	// EndpointStateWaitingForIdentity captures enum value "waiting-for-identity"
	EndpointStateWaitingForIdentity EndpointState = "waiting-for-identity"
	// EndpointStateNotReady captures enum value "not-ready"
	EndpointStateNotReady EndpointState = "not-ready"
	// EndpointStateReady captures enum value "ready"
	EndpointStateReady EndpointState = "ready"
)

// for schema
var endpointStateEnum []interface{}

func init() {
	var res []EndpointState
	if err := json.Unmarshal([]byte(`["creating","disconnected","waiting-for-identity","not-ready","ready"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointStateEnum = append(endpointStateEnum, v)
	}
}

func (m EndpointState) validateEndpointStateEnum(path, location string, value EndpointState) error {
	if err := validate.Enum(path, location, value, endpointStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this endpoint state
func (m EndpointState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEndpointStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
