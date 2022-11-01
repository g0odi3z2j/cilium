// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PatchPrefilterReader is a Reader for the PatchPrefilter structure.
type PatchPrefilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPrefilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPrefilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 461:
		result := NewPatchPrefilterInvalidCIDR()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchPrefilterFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPrefilterOK creates a PatchPrefilterOK with default headers values
func NewPatchPrefilterOK() *PatchPrefilterOK {
	return &PatchPrefilterOK{}
}

/*
PatchPrefilterOK describes a response with status code 200, with default header values.

Updated
*/
type PatchPrefilterOK struct {
	Payload *models.Prefilter
}

// IsSuccess returns true when this patch prefilter o k response has a 2xx status code
func (o *PatchPrefilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch prefilter o k response has a 3xx status code
func (o *PatchPrefilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch prefilter o k response has a 4xx status code
func (o *PatchPrefilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch prefilter o k response has a 5xx status code
func (o *PatchPrefilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch prefilter o k response a status code equal to that given
func (o *PatchPrefilterOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchPrefilterOK) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterOK  %+v", 200, o.Payload)
}

func (o *PatchPrefilterOK) String() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterOK  %+v", 200, o.Payload)
}

func (o *PatchPrefilterOK) GetPayload() *models.Prefilter {
	return o.Payload
}

func (o *PatchPrefilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPrefilterInvalidCIDR creates a PatchPrefilterInvalidCIDR with default headers values
func NewPatchPrefilterInvalidCIDR() *PatchPrefilterInvalidCIDR {
	return &PatchPrefilterInvalidCIDR{}
}

/*
PatchPrefilterInvalidCIDR describes a response with status code 461, with default header values.

Invalid CIDR prefix
*/
type PatchPrefilterInvalidCIDR struct {
	Payload models.Error
}

// IsSuccess returns true when this patch prefilter invalid c Id r response has a 2xx status code
func (o *PatchPrefilterInvalidCIDR) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch prefilter invalid c Id r response has a 3xx status code
func (o *PatchPrefilterInvalidCIDR) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch prefilter invalid c Id r response has a 4xx status code
func (o *PatchPrefilterInvalidCIDR) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch prefilter invalid c Id r response has a 5xx status code
func (o *PatchPrefilterInvalidCIDR) IsServerError() bool {
	return false
}

// IsCode returns true when this patch prefilter invalid c Id r response a status code equal to that given
func (o *PatchPrefilterInvalidCIDR) IsCode(code int) bool {
	return code == 461
}

func (o *PatchPrefilterInvalidCIDR) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterInvalidCIdR  %+v", 461, o.Payload)
}

func (o *PatchPrefilterInvalidCIDR) String() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterInvalidCIdR  %+v", 461, o.Payload)
}

func (o *PatchPrefilterInvalidCIDR) GetPayload() models.Error {
	return o.Payload
}

func (o *PatchPrefilterInvalidCIDR) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPrefilterFailure creates a PatchPrefilterFailure with default headers values
func NewPatchPrefilterFailure() *PatchPrefilterFailure {
	return &PatchPrefilterFailure{}
}

/*
PatchPrefilterFailure describes a response with status code 500, with default header values.

Prefilter update failed
*/
type PatchPrefilterFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this patch prefilter failure response has a 2xx status code
func (o *PatchPrefilterFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch prefilter failure response has a 3xx status code
func (o *PatchPrefilterFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch prefilter failure response has a 4xx status code
func (o *PatchPrefilterFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch prefilter failure response has a 5xx status code
func (o *PatchPrefilterFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this patch prefilter failure response a status code equal to that given
func (o *PatchPrefilterFailure) IsCode(code int) bool {
	return code == 500
}

func (o *PatchPrefilterFailure) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterFailure  %+v", 500, o.Payload)
}

func (o *PatchPrefilterFailure) String() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterFailure  %+v", 500, o.Payload)
}

func (o *PatchPrefilterFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PatchPrefilterFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
