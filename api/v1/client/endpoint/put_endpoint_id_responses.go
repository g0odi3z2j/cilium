// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PutEndpointIDReader is a Reader for the PutEndpointID structure.
type PutEndpointIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEndpointIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutEndpointIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutEndpointIDInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutEndpointIDExists()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutEndpointIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutEndpointIDFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutEndpointIDCreated creates a PutEndpointIDCreated with default headers values
func NewPutEndpointIDCreated() *PutEndpointIDCreated {
	return &PutEndpointIDCreated{}
}

/*
PutEndpointIDCreated handles this case with default header values.

Created
*/
type PutEndpointIDCreated struct {
}

func (o *PutEndpointIDCreated) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdCreated ", 201)
}

func (o *PutEndpointIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDInvalid creates a PutEndpointIDInvalid with default headers values
func NewPutEndpointIDInvalid() *PutEndpointIDInvalid {
	return &PutEndpointIDInvalid{}
}

/*
PutEndpointIDInvalid handles this case with default header values.

Invalid endpoint in request
*/
type PutEndpointIDInvalid struct {
	Payload models.Error
}

func (o *PutEndpointIDInvalid) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdInvalid  %+v", 400, o.Payload)
}

func (o *PutEndpointIDInvalid) GetPayload() models.Error {
	return o.Payload
}

func (o *PutEndpointIDInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEndpointIDExists creates a PutEndpointIDExists with default headers values
func NewPutEndpointIDExists() *PutEndpointIDExists {
	return &PutEndpointIDExists{}
}

/*
PutEndpointIDExists handles this case with default header values.

Endpoint already exists
*/
type PutEndpointIDExists struct {
}

func (o *PutEndpointIDExists) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdExists ", 409)
}

func (o *PutEndpointIDExists) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDTooManyRequests creates a PutEndpointIDTooManyRequests with default headers values
func NewPutEndpointIDTooManyRequests() *PutEndpointIDTooManyRequests {
	return &PutEndpointIDTooManyRequests{}
}

/*
PutEndpointIDTooManyRequests handles this case with default header values.

Rate-limiting too many requests in the given time frame
*/
type PutEndpointIDTooManyRequests struct {
}

func (o *PutEndpointIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdTooManyRequests ", 429)
}

func (o *PutEndpointIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDFailed creates a PutEndpointIDFailed with default headers values
func NewPutEndpointIDFailed() *PutEndpointIDFailed {
	return &PutEndpointIDFailed{}
}

/*
PutEndpointIDFailed handles this case with default header values.

Endpoint creation failed
*/
type PutEndpointIDFailed struct {
	Payload models.Error
}

func (o *PutEndpointIDFailed) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdFailed  %+v", 500, o.Payload)
}

func (o *PutEndpointIDFailed) GetPayload() models.Error {
	return o.Payload
}

func (o *PutEndpointIDFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
