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

// PatchPrefilterOKCode is the HTTP code returned for type PatchPrefilterOK
const PatchPrefilterOKCode int = 200

/*PatchPrefilterOK Updated

swagger:response patchPrefilterOK
*/
type PatchPrefilterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Prefilter `json:"body,omitempty"`
}

// NewPatchPrefilterOK creates PatchPrefilterOK with default headers values
func NewPatchPrefilterOK() *PatchPrefilterOK {

	return &PatchPrefilterOK{}
}

// WithPayload adds the payload to the patch prefilter o k response
func (o *PatchPrefilterOK) WithPayload(payload *models.Prefilter) *PatchPrefilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch prefilter o k response
func (o *PatchPrefilterOK) SetPayload(payload *models.Prefilter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPrefilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPrefilterInvalidCIDRCode is the HTTP code returned for type PatchPrefilterInvalidCIDR
const PatchPrefilterInvalidCIDRCode int = 461

/*PatchPrefilterInvalidCIDR Invalid CIDR prefix

swagger:response patchPrefilterInvalidCIdR
*/
type PatchPrefilterInvalidCIDR struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchPrefilterInvalidCIDR creates PatchPrefilterInvalidCIDR with default headers values
func NewPatchPrefilterInvalidCIDR() *PatchPrefilterInvalidCIDR {

	return &PatchPrefilterInvalidCIDR{}
}

// WithPayload adds the payload to the patch prefilter invalid c Id r response
func (o *PatchPrefilterInvalidCIDR) WithPayload(payload models.Error) *PatchPrefilterInvalidCIDR {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch prefilter invalid c Id r response
func (o *PatchPrefilterInvalidCIDR) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPrefilterInvalidCIDR) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(461)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchPrefilterFailureCode is the HTTP code returned for type PatchPrefilterFailure
const PatchPrefilterFailureCode int = 500

/*PatchPrefilterFailure Prefilter update failed

swagger:response patchPrefilterFailure
*/
type PatchPrefilterFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchPrefilterFailure creates PatchPrefilterFailure with default headers values
func NewPatchPrefilterFailure() *PatchPrefilterFailure {

	return &PatchPrefilterFailure{}
}

// WithPayload adds the payload to the patch prefilter failure response
func (o *PatchPrefilterFailure) WithPayload(payload models.Error) *PatchPrefilterFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch prefilter failure response
func (o *PatchPrefilterFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPrefilterFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
