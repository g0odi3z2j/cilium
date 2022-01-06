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

// GetEndpointIDLogReader is a Reader for the GetEndpointIDLog structure.
type GetEndpointIDLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointIDLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointIDLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointIDLogInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointIDLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEndpointIDLogTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointIDLogOK creates a GetEndpointIDLogOK with default headers values
func NewGetEndpointIDLogOK() *GetEndpointIDLogOK {
	return &GetEndpointIDLogOK{}
}

/*GetEndpointIDLogOK handles this case with default header values.

Success
*/
type GetEndpointIDLogOK struct {
	Payload models.EndpointStatusLog
}

func (o *GetEndpointIDLogOK) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/log][%d] getEndpointIdLogOK  %+v", 200, o.Payload)
}

func (o *GetEndpointIDLogOK) GetPayload() models.EndpointStatusLog {
	return o.Payload
}

func (o *GetEndpointIDLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDLogInvalid creates a GetEndpointIDLogInvalid with default headers values
func NewGetEndpointIDLogInvalid() *GetEndpointIDLogInvalid {
	return &GetEndpointIDLogInvalid{}
}

/*GetEndpointIDLogInvalid handles this case with default header values.

Invalid identity provided
*/
type GetEndpointIDLogInvalid struct {
}

func (o *GetEndpointIDLogInvalid) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/log][%d] getEndpointIdLogInvalid ", 400)
}

func (o *GetEndpointIDLogInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDLogNotFound creates a GetEndpointIDLogNotFound with default headers values
func NewGetEndpointIDLogNotFound() *GetEndpointIDLogNotFound {
	return &GetEndpointIDLogNotFound{}
}

/*GetEndpointIDLogNotFound handles this case with default header values.

Endpoint not found
*/
type GetEndpointIDLogNotFound struct {
}

func (o *GetEndpointIDLogNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/log][%d] getEndpointIdLogNotFound ", 404)
}

func (o *GetEndpointIDLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDLogTooManyRequests creates a GetEndpointIDLogTooManyRequests with default headers values
func NewGetEndpointIDLogTooManyRequests() *GetEndpointIDLogTooManyRequests {
	return &GetEndpointIDLogTooManyRequests{}
}

/*GetEndpointIDLogTooManyRequests handles this case with default header values.

Rate-limiting too many requests in the given time frame
*/
type GetEndpointIDLogTooManyRequests struct {
}

func (o *GetEndpointIDLogTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/log][%d] getEndpointIdLogTooManyRequests ", 429)
}

func (o *GetEndpointIDLogTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
