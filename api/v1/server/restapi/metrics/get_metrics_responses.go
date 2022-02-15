// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetMetricsOKCode is the HTTP code returned for type GetMetricsOK
const GetMetricsOKCode int = 200

/*GetMetricsOK Success

swagger:response getMetricsOK
*/
type GetMetricsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Metric `json:"body,omitempty"`
}

// NewGetMetricsOK creates GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {

	return &GetMetricsOK{}
}

// WithPayload adds the payload to the get metrics o k response
func (o *GetMetricsOK) WithPayload(payload []*models.Metric) *GetMetricsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics o k response
func (o *GetMetricsOK) SetPayload(payload []*models.Metric) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Metric, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMetricsInternalServerErrorCode is the HTTP code returned for type GetMetricsInternalServerError
const GetMetricsInternalServerErrorCode int = 500

/*GetMetricsInternalServerError Metrics cannot be retrieved

swagger:response getMetricsInternalServerError
*/
type GetMetricsInternalServerError struct {
}

// NewGetMetricsInternalServerError creates GetMetricsInternalServerError with default headers values
func NewGetMetricsInternalServerError() *GetMetricsInternalServerError {

	return &GetMetricsInternalServerError{}
}

// WriteResponse to the client
func (o *GetMetricsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
