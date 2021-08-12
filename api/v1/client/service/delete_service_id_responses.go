// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
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

// DeleteServiceIDReader is a Reader for the DeleteServiceID structure.
type DeleteServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteServiceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceIDFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceIDOK creates a DeleteServiceIDOK with default headers values
func NewDeleteServiceIDOK() *DeleteServiceIDOK {
	return &DeleteServiceIDOK{}
}

/*
DeleteServiceIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteServiceIDOK struct {
}

// IsSuccess returns true when this delete service Id o k response has a 2xx status code
func (o *DeleteServiceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service Id o k response has a 3xx status code
func (o *DeleteServiceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service Id o k response has a 4xx status code
func (o *DeleteServiceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service Id o k response has a 5xx status code
func (o *DeleteServiceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service Id o k response a status code equal to that given
func (o *DeleteServiceIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteServiceIDOK) Error() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdOK ", 200)
}

func (o *DeleteServiceIDOK) String() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdOK ", 200)
}

func (o *DeleteServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceIDForbidden creates a DeleteServiceIDForbidden with default headers values
func NewDeleteServiceIDForbidden() *DeleteServiceIDForbidden {
	return &DeleteServiceIDForbidden{}
}

/*
DeleteServiceIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteServiceIDForbidden struct {
}

// IsSuccess returns true when this delete service Id forbidden response has a 2xx status code
func (o *DeleteServiceIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service Id forbidden response has a 3xx status code
func (o *DeleteServiceIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service Id forbidden response has a 4xx status code
func (o *DeleteServiceIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service Id forbidden response has a 5xx status code
func (o *DeleteServiceIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service Id forbidden response a status code equal to that given
func (o *DeleteServiceIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteServiceIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdForbidden ", 403)
}

func (o *DeleteServiceIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdForbidden ", 403)
}

func (o *DeleteServiceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceIDNotFound creates a DeleteServiceIDNotFound with default headers values
func NewDeleteServiceIDNotFound() *DeleteServiceIDNotFound {
	return &DeleteServiceIDNotFound{}
}

/*
DeleteServiceIDNotFound describes a response with status code 404, with default header values.

Service not found
*/
type DeleteServiceIDNotFound struct {
}

// IsSuccess returns true when this delete service Id not found response has a 2xx status code
func (o *DeleteServiceIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service Id not found response has a 3xx status code
func (o *DeleteServiceIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service Id not found response has a 4xx status code
func (o *DeleteServiceIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service Id not found response has a 5xx status code
func (o *DeleteServiceIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service Id not found response a status code equal to that given
func (o *DeleteServiceIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteServiceIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdNotFound ", 404)
}

func (o *DeleteServiceIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdNotFound ", 404)
}

func (o *DeleteServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceIDFailure creates a DeleteServiceIDFailure with default headers values
func NewDeleteServiceIDFailure() *DeleteServiceIDFailure {
	return &DeleteServiceIDFailure{}
}

/*
DeleteServiceIDFailure describes a response with status code 500, with default header values.

Service deletion failed
*/
type DeleteServiceIDFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this delete service Id failure response has a 2xx status code
func (o *DeleteServiceIDFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service Id failure response has a 3xx status code
func (o *DeleteServiceIDFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service Id failure response has a 4xx status code
func (o *DeleteServiceIDFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service Id failure response has a 5xx status code
func (o *DeleteServiceIDFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this delete service Id failure response a status code equal to that given
func (o *DeleteServiceIDFailure) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteServiceIDFailure) Error() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdFailure  %+v", 500, o.Payload)
}

func (o *DeleteServiceIDFailure) String() string {
	return fmt.Sprintf("[DELETE /service/{id}][%d] deleteServiceIdFailure  %+v", 500, o.Payload)
}

func (o *DeleteServiceIDFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteServiceIDFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
