// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
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

// PatchEndpointIDLabelsReader is a Reader for the PatchEndpointIDLabels structure.
type PatchEndpointIDLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEndpointIDLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEndpointIDLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPatchEndpointIDLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEndpointIDLabelsUpdateFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchEndpointIDLabelsOK creates a PatchEndpointIDLabelsOK with default headers values
func NewPatchEndpointIDLabelsOK() *PatchEndpointIDLabelsOK {
	return &PatchEndpointIDLabelsOK{}
}

/*PatchEndpointIDLabelsOK handles this case with default header values.

Success
*/
type PatchEndpointIDLabelsOK struct {
}

func (o *PatchEndpointIDLabelsOK) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/labels][%d] patchEndpointIdLabelsOK ", 200)
}

func (o *PatchEndpointIDLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEndpointIDLabelsNotFound creates a PatchEndpointIDLabelsNotFound with default headers values
func NewPatchEndpointIDLabelsNotFound() *PatchEndpointIDLabelsNotFound {
	return &PatchEndpointIDLabelsNotFound{}
}

/*PatchEndpointIDLabelsNotFound handles this case with default header values.

Endpoint not found
*/
type PatchEndpointIDLabelsNotFound struct {
}

func (o *PatchEndpointIDLabelsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/labels][%d] patchEndpointIdLabelsNotFound ", 404)
}

func (o *PatchEndpointIDLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEndpointIDLabelsUpdateFailed creates a PatchEndpointIDLabelsUpdateFailed with default headers values
func NewPatchEndpointIDLabelsUpdateFailed() *PatchEndpointIDLabelsUpdateFailed {
	return &PatchEndpointIDLabelsUpdateFailed{}
}

/*PatchEndpointIDLabelsUpdateFailed handles this case with default header values.

Error while updating labels
*/
type PatchEndpointIDLabelsUpdateFailed struct {
	Payload models.Error
}

func (o *PatchEndpointIDLabelsUpdateFailed) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/labels][%d] patchEndpointIdLabelsUpdateFailed  %+v", 500, o.Payload)
}

func (o *PatchEndpointIDLabelsUpdateFailed) GetPayload() models.Error {
	return o.Payload
}

func (o *PatchEndpointIDLabelsUpdateFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
