// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/health/models"
)

// PutStatusProbeOKCode is the HTTP code returned for type PutStatusProbeOK
const PutStatusProbeOKCode int = 200

/*
PutStatusProbeOK Success

swagger:response putStatusProbeOK
*/
type PutStatusProbeOK struct {

	/*
	  In: Body
	*/
	Payload *models.HealthStatusResponse `json:"body,omitempty"`
}

// NewPutStatusProbeOK creates PutStatusProbeOK with default headers values
func NewPutStatusProbeOK() *PutStatusProbeOK {

	return &PutStatusProbeOK{}
}

// WithPayload adds the payload to the put status probe o k response
func (o *PutStatusProbeOK) WithPayload(payload *models.HealthStatusResponse) *PutStatusProbeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put status probe o k response
func (o *PutStatusProbeOK) SetPayload(payload *models.HealthStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutStatusProbeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutStatusProbeForbiddenCode is the HTTP code returned for type PutStatusProbeForbidden
const PutStatusProbeForbiddenCode int = 403

/*
PutStatusProbeForbidden Forbidden

swagger:response putStatusProbeForbidden
*/
type PutStatusProbeForbidden struct {
}

// NewPutStatusProbeForbidden creates PutStatusProbeForbidden with default headers values
func NewPutStatusProbeForbidden() *PutStatusProbeForbidden {

	return &PutStatusProbeForbidden{}
}

// WriteResponse to the client
func (o *PutStatusProbeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PutStatusProbeFailedCode is the HTTP code returned for type PutStatusProbeFailed
const PutStatusProbeFailedCode int = 500

/*
PutStatusProbeFailed Internal error occurred while conducting connectivity probe

swagger:response putStatusProbeFailed
*/
type PutStatusProbeFailed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutStatusProbeFailed creates PutStatusProbeFailed with default headers values
func NewPutStatusProbeFailed() *PutStatusProbeFailed {

	return &PutStatusProbeFailed{}
}

// WithPayload adds the payload to the put status probe failed response
func (o *PutStatusProbeFailed) WithPayload(payload models.Error) *PutStatusProbeFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put status probe failed response
func (o *PutStatusProbeFailed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutStatusProbeFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
