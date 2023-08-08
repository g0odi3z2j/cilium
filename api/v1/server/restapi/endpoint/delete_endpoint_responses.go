// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteEndpointOKCode is the HTTP code returned for type DeleteEndpointOK
const DeleteEndpointOKCode int = 200

/*
DeleteEndpointOK Success

swagger:response deleteEndpointOK
*/
type DeleteEndpointOK struct {
}

// NewDeleteEndpointOK creates DeleteEndpointOK with default headers values
func NewDeleteEndpointOK() *DeleteEndpointOK {

	return &DeleteEndpointOK{}
}

// WriteResponse to the client
func (o *DeleteEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteEndpointErrorsCode is the HTTP code returned for type DeleteEndpointErrors
const DeleteEndpointErrorsCode int = 206

/*
DeleteEndpointErrors Deleted with a number of errors encountered

swagger:response deleteEndpointErrors
*/
type DeleteEndpointErrors struct {

	/*
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewDeleteEndpointErrors creates DeleteEndpointErrors with default headers values
func NewDeleteEndpointErrors() *DeleteEndpointErrors {

	return &DeleteEndpointErrors{}
}

// WithPayload adds the payload to the delete endpoint errors response
func (o *DeleteEndpointErrors) WithPayload(payload int64) *DeleteEndpointErrors {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete endpoint errors response
func (o *DeleteEndpointErrors) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEndpointErrors) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(206)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteEndpointInvalidCode is the HTTP code returned for type DeleteEndpointInvalid
const DeleteEndpointInvalidCode int = 400

/*
DeleteEndpointInvalid Invalid endpoint delete request

swagger:response deleteEndpointInvalid
*/
type DeleteEndpointInvalid struct {
}

// NewDeleteEndpointInvalid creates DeleteEndpointInvalid with default headers values
func NewDeleteEndpointInvalid() *DeleteEndpointInvalid {

	return &DeleteEndpointInvalid{}
}

// WriteResponse to the client
func (o *DeleteEndpointInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteEndpointNotFoundCode is the HTTP code returned for type DeleteEndpointNotFound
const DeleteEndpointNotFoundCode int = 404

/*
DeleteEndpointNotFound No endpoints with provided parameters found

swagger:response deleteEndpointNotFound
*/
type DeleteEndpointNotFound struct {
}

// NewDeleteEndpointNotFound creates DeleteEndpointNotFound with default headers values
func NewDeleteEndpointNotFound() *DeleteEndpointNotFound {

	return &DeleteEndpointNotFound{}
}

// WriteResponse to the client
func (o *DeleteEndpointNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteEndpointTooManyRequestsCode is the HTTP code returned for type DeleteEndpointTooManyRequests
const DeleteEndpointTooManyRequestsCode int = 429

/*
DeleteEndpointTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response deleteEndpointTooManyRequests
*/
type DeleteEndpointTooManyRequests struct {
}

// NewDeleteEndpointTooManyRequests creates DeleteEndpointTooManyRequests with default headers values
func NewDeleteEndpointTooManyRequests() *DeleteEndpointTooManyRequests {

	return &DeleteEndpointTooManyRequests{}
}

// WriteResponse to the client
func (o *DeleteEndpointTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}
