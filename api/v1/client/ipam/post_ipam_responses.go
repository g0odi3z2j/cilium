// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIpamReader is a Reader for the PostIpam structure.
type PostIpamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIpamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIpamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 502:
		result := NewPostIpamFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIpamCreated creates a PostIpamCreated with default headers values
func NewPostIpamCreated() *PostIpamCreated {
	return &PostIpamCreated{}
}

/*
PostIpamCreated handles this case with default header values.

Success
*/
type PostIpamCreated struct {
	Payload *models.IPAMResponse
}

func (o *PostIpamCreated) Error() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpamCreated  %+v", 201, o.Payload)
}

func (o *PostIpamCreated) GetPayload() *models.IPAMResponse {
	return o.Payload
}

func (o *PostIpamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIpamFailure creates a PostIpamFailure with default headers values
func NewPostIpamFailure() *PostIpamFailure {
	return &PostIpamFailure{}
}

/*
PostIpamFailure handles this case with default header values.

Allocation failure
*/
type PostIpamFailure struct {
	Payload models.Error
}

func (o *PostIpamFailure) Error() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpamFailure  %+v", 502, o.Payload)
}

func (o *PostIpamFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PostIpamFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
