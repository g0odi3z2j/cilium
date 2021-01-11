// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// ConfigurationMap Map of configuration key/value pairs.
//
//
// swagger:model ConfigurationMap
type ConfigurationMap map[string]string

// Validate validates this configuration map
func (m ConfigurationMap) Validate(formats strfmt.Registry) error {
	return nil
}
