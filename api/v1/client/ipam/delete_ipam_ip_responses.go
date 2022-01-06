// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// DeleteIpamIPReader is a Reader for the DeleteIpamIP structure.
type DeleteIpamIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIpamIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIpamIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIpamIPInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIpamIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIpamIPFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteIpamIPDisabled()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIpamIPOK creates a DeleteIpamIPOK with default headers values
func NewDeleteIpamIPOK() *DeleteIpamIPOK {
	return &DeleteIpamIPOK{}
}

/*DeleteIpamIPOK handles this case with default header values.

Success
*/
type DeleteIpamIPOK struct {
}

func (o *DeleteIpamIPOK) Error() string {
	return fmt.Sprintf("[DELETE /ipam/{ip}][%d] deleteIpamIpOK ", 200)
}

func (o *DeleteIpamIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIpamIPInvalid creates a DeleteIpamIPInvalid with default headers values
func NewDeleteIpamIPInvalid() *DeleteIpamIPInvalid {
	return &DeleteIpamIPInvalid{}
}

/*DeleteIpamIPInvalid handles this case with default header values.

Invalid IP address
*/
type DeleteIpamIPInvalid struct {
}

func (o *DeleteIpamIPInvalid) Error() string {
	return fmt.Sprintf("[DELETE /ipam/{ip}][%d] deleteIpamIpInvalid ", 400)
}

func (o *DeleteIpamIPInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIpamIPNotFound creates a DeleteIpamIPNotFound with default headers values
func NewDeleteIpamIPNotFound() *DeleteIpamIPNotFound {
	return &DeleteIpamIPNotFound{}
}

/*DeleteIpamIPNotFound handles this case with default header values.

IP address not found
*/
type DeleteIpamIPNotFound struct {
}

func (o *DeleteIpamIPNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ipam/{ip}][%d] deleteIpamIpNotFound ", 404)
}

func (o *DeleteIpamIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIpamIPFailure creates a DeleteIpamIPFailure with default headers values
func NewDeleteIpamIPFailure() *DeleteIpamIPFailure {
	return &DeleteIpamIPFailure{}
}

/*DeleteIpamIPFailure handles this case with default header values.

Address release failure
*/
type DeleteIpamIPFailure struct {
	Payload models.Error
}

func (o *DeleteIpamIPFailure) Error() string {
	return fmt.Sprintf("[DELETE /ipam/{ip}][%d] deleteIpamIpFailure  %+v", 500, o.Payload)
}

func (o *DeleteIpamIPFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteIpamIPFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIpamIPDisabled creates a DeleteIpamIPDisabled with default headers values
func NewDeleteIpamIPDisabled() *DeleteIpamIPDisabled {
	return &DeleteIpamIPDisabled{}
}

/*DeleteIpamIPDisabled handles this case with default header values.

Allocation for address family disabled
*/
type DeleteIpamIPDisabled struct {
}

func (o *DeleteIpamIPDisabled) Error() string {
	return fmt.Sprintf("[DELETE /ipam/{ip}][%d] deleteIpamIpDisabled ", 501)
}

func (o *DeleteIpamIPDisabled) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
