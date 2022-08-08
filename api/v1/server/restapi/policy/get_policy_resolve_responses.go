// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetPolicyResolveOKCode is the HTTP code returned for type GetPolicyResolveOK
const GetPolicyResolveOKCode int = 200

/*
GetPolicyResolveOK Success

swagger:response getPolicyResolveOK
*/
type GetPolicyResolveOK struct {

	/*
	  In: Body
	*/
	Payload *models.PolicyTraceResult `json:"body,omitempty"`
}

// NewGetPolicyResolveOK creates GetPolicyResolveOK with default headers values
func NewGetPolicyResolveOK() *GetPolicyResolveOK {

	return &GetPolicyResolveOK{}
}

// WithPayload adds the payload to the get policy resolve o k response
func (o *GetPolicyResolveOK) WithPayload(payload *models.PolicyTraceResult) *GetPolicyResolveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy resolve o k response
func (o *GetPolicyResolveOK) SetPayload(payload *models.PolicyTraceResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyResolveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
