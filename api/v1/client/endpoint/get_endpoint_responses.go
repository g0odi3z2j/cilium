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

// GetEndpointReader is a Reader for the GetEndpoint structure.
type GetEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEndpointTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointOK creates a GetEndpointOK with default headers values
func NewGetEndpointOK() *GetEndpointOK {
	return &GetEndpointOK{}
}

/*
GetEndpointOK describes a response with status code 200, with default header values.

Success
*/
type GetEndpointOK struct {
	Payload []*models.Endpoint
}

// IsSuccess returns true when this get endpoint o k response has a 2xx status code
func (o *GetEndpointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint o k response has a 3xx status code
func (o *GetEndpointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint o k response has a 4xx status code
func (o *GetEndpointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint o k response has a 5xx status code
func (o *GetEndpointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint o k response a status code equal to that given
func (o *GetEndpointOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEndpointOK) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEndpointOK) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEndpointOK) GetPayload() []*models.Endpoint {
	return o.Payload
}

func (o *GetEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointNotFound creates a GetEndpointNotFound with default headers values
func NewGetEndpointNotFound() *GetEndpointNotFound {
	return &GetEndpointNotFound{}
}

/*
GetEndpointNotFound describes a response with status code 404, with default header values.

Endpoints with provided parameters not found
*/
type GetEndpointNotFound struct {
}

// IsSuccess returns true when this get endpoint not found response has a 2xx status code
func (o *GetEndpointNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint not found response has a 3xx status code
func (o *GetEndpointNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint not found response has a 4xx status code
func (o *GetEndpointNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint not found response has a 5xx status code
func (o *GetEndpointNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint not found response a status code equal to that given
func (o *GetEndpointNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointNotFound ", 404)
}

func (o *GetEndpointNotFound) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointNotFound ", 404)
}

func (o *GetEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointTooManyRequests creates a GetEndpointTooManyRequests with default headers values
func NewGetEndpointTooManyRequests() *GetEndpointTooManyRequests {
	return &GetEndpointTooManyRequests{}
}

/*
GetEndpointTooManyRequests describes a response with status code 429, with default header values.

Rate-limiting too many requests in the given time frame
*/
type GetEndpointTooManyRequests struct {
}

// IsSuccess returns true when this get endpoint too many requests response has a 2xx status code
func (o *GetEndpointTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint too many requests response has a 3xx status code
func (o *GetEndpointTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint too many requests response has a 4xx status code
func (o *GetEndpointTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint too many requests response has a 5xx status code
func (o *GetEndpointTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint too many requests response a status code equal to that given
func (o *GetEndpointTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetEndpointTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointTooManyRequests ", 429)
}

func (o *GetEndpointTooManyRequests) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointTooManyRequests ", 429)
}

func (o *GetEndpointTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
