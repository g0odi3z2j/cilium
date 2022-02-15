// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetHealthzOKCode is the HTTP code returned for type GetHealthzOK
const GetHealthzOKCode int = 200

/*GetHealthzOK Success

swagger:response getHealthzOK
*/
type GetHealthzOK struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetHealthzOK creates GetHealthzOK with default headers values
func NewGetHealthzOK() *GetHealthzOK {

	return &GetHealthzOK{}
}

// WithPayload adds the payload to the get healthz o k response
func (o *GetHealthzOK) WithPayload(payload *models.StatusResponse) *GetHealthzOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthz o k response
func (o *GetHealthzOK) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
