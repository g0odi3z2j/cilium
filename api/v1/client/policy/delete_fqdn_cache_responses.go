// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
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

// DeleteFqdnCacheReader is a Reader for the DeleteFqdnCache structure.
type DeleteFqdnCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFqdnCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFqdnCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFqdnCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFqdnCacheForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFqdnCacheOK creates a DeleteFqdnCacheOK with default headers values
func NewDeleteFqdnCacheOK() *DeleteFqdnCacheOK {
	return &DeleteFqdnCacheOK{}
}

/*
DeleteFqdnCacheOK describes a response with status code 200, with default header values.

Success
*/
type DeleteFqdnCacheOK struct {
}

// IsSuccess returns true when this delete fqdn cache o k response has a 2xx status code
func (o *DeleteFqdnCacheOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete fqdn cache o k response has a 3xx status code
func (o *DeleteFqdnCacheOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fqdn cache o k response has a 4xx status code
func (o *DeleteFqdnCacheOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fqdn cache o k response has a 5xx status code
func (o *DeleteFqdnCacheOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fqdn cache o k response a status code equal to that given
func (o *DeleteFqdnCacheOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteFqdnCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheOK ", 200)
}

func (o *DeleteFqdnCacheOK) String() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheOK ", 200)
}

func (o *DeleteFqdnCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFqdnCacheBadRequest creates a DeleteFqdnCacheBadRequest with default headers values
func NewDeleteFqdnCacheBadRequest() *DeleteFqdnCacheBadRequest {
	return &DeleteFqdnCacheBadRequest{}
}

/*
DeleteFqdnCacheBadRequest describes a response with status code 400, with default header values.

Invalid request (error parsing parameters)
*/
type DeleteFqdnCacheBadRequest struct {
	Payload models.Error
}

// IsSuccess returns true when this delete fqdn cache bad request response has a 2xx status code
func (o *DeleteFqdnCacheBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fqdn cache bad request response has a 3xx status code
func (o *DeleteFqdnCacheBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fqdn cache bad request response has a 4xx status code
func (o *DeleteFqdnCacheBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fqdn cache bad request response has a 5xx status code
func (o *DeleteFqdnCacheBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fqdn cache bad request response a status code equal to that given
func (o *DeleteFqdnCacheBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteFqdnCacheBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFqdnCacheBadRequest) String() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFqdnCacheBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteFqdnCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFqdnCacheForbidden creates a DeleteFqdnCacheForbidden with default headers values
func NewDeleteFqdnCacheForbidden() *DeleteFqdnCacheForbidden {
	return &DeleteFqdnCacheForbidden{}
}

/*
DeleteFqdnCacheForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFqdnCacheForbidden struct {
}

// IsSuccess returns true when this delete fqdn cache forbidden response has a 2xx status code
func (o *DeleteFqdnCacheForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fqdn cache forbidden response has a 3xx status code
func (o *DeleteFqdnCacheForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fqdn cache forbidden response has a 4xx status code
func (o *DeleteFqdnCacheForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fqdn cache forbidden response has a 5xx status code
func (o *DeleteFqdnCacheForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fqdn cache forbidden response a status code equal to that given
func (o *DeleteFqdnCacheForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteFqdnCacheForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheForbidden ", 403)
}

func (o *DeleteFqdnCacheForbidden) String() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheForbidden ", 403)
}

func (o *DeleteFqdnCacheForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
