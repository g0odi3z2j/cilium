// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetHealthzReader is a Reader for the GetHealthz structure.
type GetHealthzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthzOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHealthzInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetHealthzNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHealthzOK creates a GetHealthzOK with default headers values
func NewGetHealthzOK() *GetHealthzOK {
	return &GetHealthzOK{}
}

/*
GetHealthzOK describes a response with status code 200, with default header values.

Cilium operator is healthy
*/
type GetHealthzOK struct {
	Payload string
}

// IsSuccess returns true when this get healthz o k response has a 2xx status code
func (o *GetHealthzOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get healthz o k response has a 3xx status code
func (o *GetHealthzOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get healthz o k response has a 4xx status code
func (o *GetHealthzOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get healthz o k response has a 5xx status code
func (o *GetHealthzOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get healthz o k response a status code equal to that given
func (o *GetHealthzOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHealthzOK) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzOK  %+v", 200, o.Payload)
}

func (o *GetHealthzOK) String() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzOK  %+v", 200, o.Payload)
}

func (o *GetHealthzOK) GetPayload() string {
	return o.Payload
}

func (o *GetHealthzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthzInternalServerError creates a GetHealthzInternalServerError with default headers values
func NewGetHealthzInternalServerError() *GetHealthzInternalServerError {
	return &GetHealthzInternalServerError{}
}

/*
GetHealthzInternalServerError describes a response with status code 500, with default header values.

Cilium operator is not healthy
*/
type GetHealthzInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get healthz internal server error response has a 2xx status code
func (o *GetHealthzInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get healthz internal server error response has a 3xx status code
func (o *GetHealthzInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get healthz internal server error response has a 4xx status code
func (o *GetHealthzInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get healthz internal server error response has a 5xx status code
func (o *GetHealthzInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get healthz internal server error response a status code equal to that given
func (o *GetHealthzInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHealthzInternalServerError) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthzInternalServerError) String() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthzInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetHealthzInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthzNotImplemented creates a GetHealthzNotImplemented with default headers values
func NewGetHealthzNotImplemented() *GetHealthzNotImplemented {
	return &GetHealthzNotImplemented{}
}

/*
GetHealthzNotImplemented describes a response with status code 501, with default header values.

Cilium operator health status not available
*/
type GetHealthzNotImplemented struct {
	Payload string
}

// IsSuccess returns true when this get healthz not implemented response has a 2xx status code
func (o *GetHealthzNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get healthz not implemented response has a 3xx status code
func (o *GetHealthzNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get healthz not implemented response has a 4xx status code
func (o *GetHealthzNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this get healthz not implemented response has a 5xx status code
func (o *GetHealthzNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this get healthz not implemented response a status code equal to that given
func (o *GetHealthzNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *GetHealthzNotImplemented) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzNotImplemented  %+v", 501, o.Payload)
}

func (o *GetHealthzNotImplemented) String() string {
	return fmt.Sprintf("[GET /healthz][%d] getHealthzNotImplemented  %+v", 501, o.Payload)
}

func (o *GetHealthzNotImplemented) GetPayload() string {
	return o.Payload
}

func (o *GetHealthzNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
