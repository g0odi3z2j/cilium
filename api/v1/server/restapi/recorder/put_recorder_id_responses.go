// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PutRecorderIDOKCode is the HTTP code returned for type PutRecorderIDOK
const PutRecorderIDOKCode int = 200

/*
PutRecorderIDOK Updated

swagger:response putRecorderIdOK
*/
type PutRecorderIDOK struct {
}

// NewPutRecorderIDOK creates PutRecorderIDOK with default headers values
func NewPutRecorderIDOK() *PutRecorderIDOK {

	return &PutRecorderIDOK{}
}

// WriteResponse to the client
func (o *PutRecorderIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutRecorderIDCreatedCode is the HTTP code returned for type PutRecorderIDCreated
const PutRecorderIDCreatedCode int = 201

/*
PutRecorderIDCreated Created

swagger:response putRecorderIdCreated
*/
type PutRecorderIDCreated struct {
}

// NewPutRecorderIDCreated creates PutRecorderIDCreated with default headers values
func NewPutRecorderIDCreated() *PutRecorderIDCreated {

	return &PutRecorderIDCreated{}
}

// WriteResponse to the client
func (o *PutRecorderIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PutRecorderIDFailureCode is the HTTP code returned for type PutRecorderIDFailure
const PutRecorderIDFailureCode int = 500

/*
PutRecorderIDFailure Error while creating recorder

swagger:response putRecorderIdFailure
*/
type PutRecorderIDFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutRecorderIDFailure creates PutRecorderIDFailure with default headers values
func NewPutRecorderIDFailure() *PutRecorderIDFailure {

	return &PutRecorderIDFailure{}
}

// WithPayload adds the payload to the put recorder Id failure response
func (o *PutRecorderIDFailure) WithPayload(payload models.Error) *PutRecorderIDFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put recorder Id failure response
func (o *PutRecorderIDFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRecorderIDFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
