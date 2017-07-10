package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrontendAddress Layer 4 address
// swagger:model FrontendAddress
type FrontendAddress struct {

	// Layer 3 address
	IP string `json:"ip,omitempty"`

	// Layer 4 port number
	Port uint16 `json:"port,omitempty"`

	// Layer 4 protocol
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this frontend address
func (m *FrontendAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var frontendAddressTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendAddressTypeProtocolPropEnum = append(frontendAddressTypeProtocolPropEnum, v)
	}
}

const (
	// FrontendAddressProtocolTCP captures enum value "tcp"
	FrontendAddressProtocolTCP string = "tcp"
	// FrontendAddressProtocolUDP captures enum value "udp"
	FrontendAddressProtocolUDP string = "udp"
	// FrontendAddressProtocolAny captures enum value "any"
	FrontendAddressProtocolAny string = "any"
)

// prop value enum
func (m *FrontendAddress) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontendAddressTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FrontendAddress) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}
