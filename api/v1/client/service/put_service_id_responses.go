// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PutServiceIDReader is a Reader for the PutServiceID structure.
type PutServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPutServiceIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 460:
		result := NewPutServiceIDInvalidFrontend()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 461:
		result := NewPutServiceIDInvalidBackend()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutServiceIDFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutServiceIDOK creates a PutServiceIDOK with default headers values
func NewPutServiceIDOK() *PutServiceIDOK {
	return &PutServiceIDOK{}
}

/*PutServiceIDOK handles this case with default header values.

Updated
*/
type PutServiceIDOK struct {
}

func (o *PutServiceIDOK) Error() string {
	return fmt.Sprintf("[PUT /service/{id}][%d] putServiceIdOK ", 200)
}

func (o *PutServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutServiceIDCreated creates a PutServiceIDCreated with default headers values
func NewPutServiceIDCreated() *PutServiceIDCreated {
	return &PutServiceIDCreated{}
}

/*PutServiceIDCreated handles this case with default header values.

Created
*/
type PutServiceIDCreated struct {
}

func (o *PutServiceIDCreated) Error() string {
	return fmt.Sprintf("[PUT /service/{id}][%d] putServiceIdCreated ", 201)
}

func (o *PutServiceIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutServiceIDInvalidFrontend creates a PutServiceIDInvalidFrontend with default headers values
func NewPutServiceIDInvalidFrontend() *PutServiceIDInvalidFrontend {
	return &PutServiceIDInvalidFrontend{}
}

/*PutServiceIDInvalidFrontend handles this case with default header values.

Invalid frontend in service configuration
*/
type PutServiceIDInvalidFrontend struct {
	Payload models.Error
}

func (o *PutServiceIDInvalidFrontend) Error() string {
	return fmt.Sprintf("[PUT /service/{id}][%d] putServiceIdInvalidFrontend  %+v", 460, o.Payload)
}

func (o *PutServiceIDInvalidFrontend) GetPayload() models.Error {
	return o.Payload
}

func (o *PutServiceIDInvalidFrontend) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutServiceIDInvalidBackend creates a PutServiceIDInvalidBackend with default headers values
func NewPutServiceIDInvalidBackend() *PutServiceIDInvalidBackend {
	return &PutServiceIDInvalidBackend{}
}

/*PutServiceIDInvalidBackend handles this case with default header values.

Invalid backend in service configuration
*/
type PutServiceIDInvalidBackend struct {
	Payload models.Error
}

func (o *PutServiceIDInvalidBackend) Error() string {
	return fmt.Sprintf("[PUT /service/{id}][%d] putServiceIdInvalidBackend  %+v", 461, o.Payload)
}

func (o *PutServiceIDInvalidBackend) GetPayload() models.Error {
	return o.Payload
}

func (o *PutServiceIDInvalidBackend) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutServiceIDFailure creates a PutServiceIDFailure with default headers values
func NewPutServiceIDFailure() *PutServiceIDFailure {
	return &PutServiceIDFailure{}
}

/*PutServiceIDFailure handles this case with default header values.

Error while creating service
*/
type PutServiceIDFailure struct {
	Payload models.Error
}

func (o *PutServiceIDFailure) Error() string {
	return fmt.Sprintf("[PUT /service/{id}][%d] putServiceIdFailure  %+v", 500, o.Payload)
}

func (o *PutServiceIDFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PutServiceIDFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
