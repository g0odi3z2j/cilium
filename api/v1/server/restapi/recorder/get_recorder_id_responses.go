// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetRecorderIDOKCode is the HTTP code returned for type GetRecorderIDOK
const GetRecorderIDOKCode int = 200

/*GetRecorderIDOK Success

swagger:response getRecorderIdOK
*/
type GetRecorderIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Recorder `json:"body,omitempty"`
}

// NewGetRecorderIDOK creates GetRecorderIDOK with default headers values
func NewGetRecorderIDOK() *GetRecorderIDOK {

	return &GetRecorderIDOK{}
}

// WithPayload adds the payload to the get recorder Id o k response
func (o *GetRecorderIDOK) WithPayload(payload *models.Recorder) *GetRecorderIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recorder Id o k response
func (o *GetRecorderIDOK) SetPayload(payload *models.Recorder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecorderIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRecorderIDNotFoundCode is the HTTP code returned for type GetRecorderIDNotFound
const GetRecorderIDNotFoundCode int = 404

/*GetRecorderIDNotFound Recorder not found

swagger:response getRecorderIdNotFound
*/
type GetRecorderIDNotFound struct {
}

// NewGetRecorderIDNotFound creates GetRecorderIDNotFound with default headers values
func NewGetRecorderIDNotFound() *GetRecorderIDNotFound {

	return &GetRecorderIDNotFound{}
}

// WriteResponse to the client
func (o *GetRecorderIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
