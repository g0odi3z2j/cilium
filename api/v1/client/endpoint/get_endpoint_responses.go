// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
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

/*GetEndpointOK handles this case with default header values.

Success
*/
type GetEndpointOK struct {
	Payload []*models.Endpoint
}

func (o *GetEndpointOK) Error() string {
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

/*GetEndpointNotFound handles this case with default header values.

Endpoints with provided parameters not found
*/
type GetEndpointNotFound struct {
}

func (o *GetEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointNotFound ", 404)
}

func (o *GetEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointTooManyRequests creates a GetEndpointTooManyRequests with default headers values
func NewGetEndpointTooManyRequests() *GetEndpointTooManyRequests {
	return &GetEndpointTooManyRequests{}
}

/*GetEndpointTooManyRequests handles this case with default header values.

Rate-limiting too many requests in the given time frame
*/
type GetEndpointTooManyRequests struct {
}

func (o *GetEndpointTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointTooManyRequests ", 429)
}

func (o *GetEndpointTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
