// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetServiceIDOKCode is the HTTP code returned for type GetServiceIDOK
const GetServiceIDOKCode int = 200

/*GetServiceIDOK Success

swagger:response getServiceIdOK
*/
type GetServiceIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Service `json:"body,omitempty"`
}

// NewGetServiceIDOK creates GetServiceIDOK with default headers values
func NewGetServiceIDOK() *GetServiceIDOK {

	return &GetServiceIDOK{}
}

// WithPayload adds the payload to the get service Id o k response
func (o *GetServiceIDOK) WithPayload(payload *models.Service) *GetServiceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service Id o k response
func (o *GetServiceIDOK) SetPayload(payload *models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceIDNotFoundCode is the HTTP code returned for type GetServiceIDNotFound
const GetServiceIDNotFoundCode int = 404

/*GetServiceIDNotFound Service not found

swagger:response getServiceIdNotFound
*/
type GetServiceIDNotFound struct {
}

// NewGetServiceIDNotFound creates GetServiceIDNotFound with default headers values
func NewGetServiceIDNotFound() *GetServiceIDNotFound {

	return &GetServiceIDNotFound{}
}

// WriteResponse to the client
func (o *GetServiceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
