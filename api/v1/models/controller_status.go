// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ControllerStatus Status of a controller
// swagger:model ControllerStatus

type ControllerStatus struct {

	// configuration
	Configuration *ControllerStatusConfiguration `json:"configuration,omitempty"`

	// Name of controller
	Name string `json:"name,omitempty"`

	// status
	Status *ControllerStatusStatus `json:"status,omitempty"`

	// UUID of controller
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

/* polymorph ControllerStatus configuration false */

/* polymorph ControllerStatus name false */

/* polymorph ControllerStatus status false */

/* polymorph ControllerStatus uuid false */

// Validate validates this controller status
func (m *ControllerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerStatus) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {

		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ControllerStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatus) UnmarshalBinary(b []byte) error {
	var res ControllerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ControllerStatusConfiguration Configuration of controller
// swagger:model ControllerStatusConfiguration

type ControllerStatusConfiguration struct {

	// Retry on error
	ErrorRetry bool `json:"error-retry,omitempty"`

	// Base error retry back-off time
	ErrorRetryBase strfmt.Duration `json:"error-retry-base,omitempty"`

	// Regular synchronization interval
	Interval strfmt.Duration `json:"interval,omitempty"`
}

/* polymorph ControllerStatusConfiguration error-retry false */

/* polymorph ControllerStatusConfiguration error-retry-base false */

/* polymorph ControllerStatusConfiguration interval false */

// Validate validates this controller status configuration
func (m *ControllerStatusConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatusConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatusConfiguration) UnmarshalBinary(b []byte) error {
	var res ControllerStatusConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ControllerStatusStatus Current status of controller
// swagger:model ControllerStatusStatus

type ControllerStatusStatus struct {

	// Number of consecutive errors since last success
	ConsecutiveFailureCount int64 `json:"consecutive-failure-count,omitempty"`

	// Total number of failed runs
	FailureCount int64 `json:"failure-count,omitempty"`

	// Error message of last failed run
	LastFailureMsg string `json:"last-failure-msg,omitempty"`

	// Timestamp of last error
	LastFailureTimestamp strfmt.DateTime `json:"last-failure-timestamp,omitempty"`

	// Timestsamp of last success
	LastSuccessTimestamp strfmt.DateTime `json:"last-success-timestamp,omitempty"`

	// Total number of successful runs
	SuccessCount int64 `json:"success-count,omitempty"`
}

/* polymorph ControllerStatusStatus consecutive-failure-count false */

/* polymorph ControllerStatusStatus failure-count false */

/* polymorph ControllerStatusStatus last-failure-msg false */

/* polymorph ControllerStatusStatus last-failure-timestamp false */

/* polymorph ControllerStatusStatus last-success-timestamp false */

/* polymorph ControllerStatusStatus success-count false */

// Validate validates this controller status status
func (m *ControllerStatusStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatusStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatusStatus) UnmarshalBinary(b []byte) error {
	var res ControllerStatusStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
