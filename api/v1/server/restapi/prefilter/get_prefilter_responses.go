// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetPrefilterOKCode is the HTTP code returned for type GetPrefilterOK
const GetPrefilterOKCode int = 200

/*GetPrefilterOK Success

swagger:response getPrefilterOK
*/
type GetPrefilterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Prefilter `json:"body,omitempty"`
}

// NewGetPrefilterOK creates GetPrefilterOK with default headers values
func NewGetPrefilterOK() *GetPrefilterOK {

	return &GetPrefilterOK{}
}

// WithPayload adds the payload to the get prefilter o k response
func (o *GetPrefilterOK) WithPayload(payload *models.Prefilter) *GetPrefilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get prefilter o k response
func (o *GetPrefilterOK) SetPayload(payload *models.Prefilter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPrefilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPrefilterFailureCode is the HTTP code returned for type GetPrefilterFailure
const GetPrefilterFailureCode int = 500

/*GetPrefilterFailure Prefilter get failed

swagger:response getPrefilterFailure
*/
type GetPrefilterFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetPrefilterFailure creates GetPrefilterFailure with default headers values
func NewGetPrefilterFailure() *GetPrefilterFailure {

	return &GetPrefilterFailure{}
}

// WithPayload adds the payload to the get prefilter failure response
func (o *GetPrefilterFailure) WithPayload(payload models.Error) *GetPrefilterFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get prefilter failure response
func (o *GetPrefilterFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPrefilterFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
