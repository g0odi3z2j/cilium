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

// PutPolicyReader is a Reader for the PutPolicy structure.
type PutPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutPolicyInvalidPolicy()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 460:
		result := NewPutPolicyInvalidPath()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutPolicyFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutPolicyOK creates a PutPolicyOK with default headers values
func NewPutPolicyOK() *PutPolicyOK {
	return &PutPolicyOK{}
}

/*
PutPolicyOK describes a response with status code 200, with default header values.

Success
*/
type PutPolicyOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this put policy o k response has a 2xx status code
func (o *PutPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put policy o k response has a 3xx status code
func (o *PutPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put policy o k response has a 4xx status code
func (o *PutPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put policy o k response has a 5xx status code
func (o *PutPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put policy o k response a status code equal to that given
func (o *PutPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyOK  %+v", 200, o.Payload)
}

func (o *PutPolicyOK) String() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyOK  %+v", 200, o.Payload)
}

func (o *PutPolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *PutPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyInvalidPolicy creates a PutPolicyInvalidPolicy with default headers values
func NewPutPolicyInvalidPolicy() *PutPolicyInvalidPolicy {
	return &PutPolicyInvalidPolicy{}
}

/*
PutPolicyInvalidPolicy describes a response with status code 400, with default header values.

Invalid policy
*/
type PutPolicyInvalidPolicy struct {
	Payload models.Error
}

// IsSuccess returns true when this put policy invalid policy response has a 2xx status code
func (o *PutPolicyInvalidPolicy) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put policy invalid policy response has a 3xx status code
func (o *PutPolicyInvalidPolicy) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put policy invalid policy response has a 4xx status code
func (o *PutPolicyInvalidPolicy) IsClientError() bool {
	return true
}

// IsServerError returns true when this put policy invalid policy response has a 5xx status code
func (o *PutPolicyInvalidPolicy) IsServerError() bool {
	return false
}

// IsCode returns true when this put policy invalid policy response a status code equal to that given
func (o *PutPolicyInvalidPolicy) IsCode(code int) bool {
	return code == 400
}

func (o *PutPolicyInvalidPolicy) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPolicy  %+v", 400, o.Payload)
}

func (o *PutPolicyInvalidPolicy) String() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPolicy  %+v", 400, o.Payload)
}

func (o *PutPolicyInvalidPolicy) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyInvalidPolicy) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyInvalidPath creates a PutPolicyInvalidPath with default headers values
func NewPutPolicyInvalidPath() *PutPolicyInvalidPath {
	return &PutPolicyInvalidPath{}
}

/*
PutPolicyInvalidPath describes a response with status code 460, with default header values.

Invalid path
*/
type PutPolicyInvalidPath struct {
	Payload models.Error
}

// IsSuccess returns true when this put policy invalid path response has a 2xx status code
func (o *PutPolicyInvalidPath) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put policy invalid path response has a 3xx status code
func (o *PutPolicyInvalidPath) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put policy invalid path response has a 4xx status code
func (o *PutPolicyInvalidPath) IsClientError() bool {
	return true
}

// IsServerError returns true when this put policy invalid path response has a 5xx status code
func (o *PutPolicyInvalidPath) IsServerError() bool {
	return false
}

// IsCode returns true when this put policy invalid path response a status code equal to that given
func (o *PutPolicyInvalidPath) IsCode(code int) bool {
	return code == 460
}

func (o *PutPolicyInvalidPath) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPath  %+v", 460, o.Payload)
}

func (o *PutPolicyInvalidPath) String() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPath  %+v", 460, o.Payload)
}

func (o *PutPolicyInvalidPath) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyInvalidPath) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyFailure creates a PutPolicyFailure with default headers values
func NewPutPolicyFailure() *PutPolicyFailure {
	return &PutPolicyFailure{}
}

/*
PutPolicyFailure describes a response with status code 500, with default header values.

Policy import failed
*/
type PutPolicyFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this put policy failure response has a 2xx status code
func (o *PutPolicyFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put policy failure response has a 3xx status code
func (o *PutPolicyFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put policy failure response has a 4xx status code
func (o *PutPolicyFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this put policy failure response has a 5xx status code
func (o *PutPolicyFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this put policy failure response a status code equal to that given
func (o *PutPolicyFailure) IsCode(code int) bool {
	return code == 500
}

func (o *PutPolicyFailure) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyFailure  %+v", 500, o.Payload)
}

func (o *PutPolicyFailure) String() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyFailure  %+v", 500, o.Payload)
}

func (o *PutPolicyFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
