// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetFqdnCacheIDReader is a Reader for the GetFqdnCacheID structure.
type GetFqdnCacheIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFqdnCacheIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFqdnCacheIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFqdnCacheIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFqdnCacheIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFqdnCacheIDOK creates a GetFqdnCacheIDOK with default headers values
func NewGetFqdnCacheIDOK() *GetFqdnCacheIDOK {
	return &GetFqdnCacheIDOK{}
}

/*GetFqdnCacheIDOK handles this case with default header values.

Success
*/
type GetFqdnCacheIDOK struct {
	Payload []*models.DNSLookup
}

func (o *GetFqdnCacheIDOK) Error() string {
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdOK  %+v", 200, o.Payload)
}

func (o *GetFqdnCacheIDOK) GetPayload() []*models.DNSLookup {
	return o.Payload
}

func (o *GetFqdnCacheIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFqdnCacheIDBadRequest creates a GetFqdnCacheIDBadRequest with default headers values
func NewGetFqdnCacheIDBadRequest() *GetFqdnCacheIDBadRequest {
	return &GetFqdnCacheIDBadRequest{}
}

/*GetFqdnCacheIDBadRequest handles this case with default header values.

Invalid request (error parsing parameters)
*/
type GetFqdnCacheIDBadRequest struct {
	Payload models.Error
}

func (o *GetFqdnCacheIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetFqdnCacheIDBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *GetFqdnCacheIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFqdnCacheIDNotFound creates a GetFqdnCacheIDNotFound with default headers values
func NewGetFqdnCacheIDNotFound() *GetFqdnCacheIDNotFound {
	return &GetFqdnCacheIDNotFound{}
}

/*GetFqdnCacheIDNotFound handles this case with default header values.

No DNS data with provided parameters found
*/
type GetFqdnCacheIDNotFound struct {
}

func (o *GetFqdnCacheIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fqdn/cache/{id}][%d] getFqdnCacheIdNotFound ", 404)
}

func (o *GetFqdnCacheIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
