// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDLabelsOKCode is the HTTP code returned for type GetEndpointIDLabelsOK
const GetEndpointIDLabelsOKCode int = 200

/*GetEndpointIDLabelsOK Success

swagger:response getEndpointIdLabelsOK
*/
type GetEndpointIDLabelsOK struct {

	/*
	  In: Body
	*/
	Payload *models.LabelConfiguration `json:"body,omitempty"`
}

// NewGetEndpointIDLabelsOK creates GetEndpointIDLabelsOK with default headers values
func NewGetEndpointIDLabelsOK() *GetEndpointIDLabelsOK {

	return &GetEndpointIDLabelsOK{}
}

// WithPayload adds the payload to the get endpoint Id labels o k response
func (o *GetEndpointIDLabelsOK) WithPayload(payload *models.LabelConfiguration) *GetEndpointIDLabelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint Id labels o k response
func (o *GetEndpointIDLabelsOK) SetPayload(payload *models.LabelConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointIDLabelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEndpointIDLabelsNotFoundCode is the HTTP code returned for type GetEndpointIDLabelsNotFound
const GetEndpointIDLabelsNotFoundCode int = 404

/*GetEndpointIDLabelsNotFound Endpoint not found

swagger:response getEndpointIdLabelsNotFound
*/
type GetEndpointIDLabelsNotFound struct {
}

// NewGetEndpointIDLabelsNotFound creates GetEndpointIDLabelsNotFound with default headers values
func NewGetEndpointIDLabelsNotFound() *GetEndpointIDLabelsNotFound {

	return &GetEndpointIDLabelsNotFound{}
}

// WriteResponse to the client
func (o *GetEndpointIDLabelsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetEndpointIDLabelsTooManyRequestsCode is the HTTP code returned for type GetEndpointIDLabelsTooManyRequests
const GetEndpointIDLabelsTooManyRequestsCode int = 429

/*GetEndpointIDLabelsTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response getEndpointIdLabelsTooManyRequests
*/
type GetEndpointIDLabelsTooManyRequests struct {
}

// NewGetEndpointIDLabelsTooManyRequests creates GetEndpointIDLabelsTooManyRequests with default headers values
func NewGetEndpointIDLabelsTooManyRequests() *GetEndpointIDLabelsTooManyRequests {

	return &GetEndpointIDLabelsTooManyRequests{}
}

// WriteResponse to the client
func (o *GetEndpointIDLabelsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}
