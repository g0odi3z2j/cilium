// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
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

Success
*/
type GetHealthzOK struct {
	Payload *models.StatusResponse
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

func (o *GetHealthzOK) GetPayload() *models.StatusResponse {
	return o.Payload
}

func (o *GetHealthzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
