// Code generated by go-swagger; DO NOT EDIT.

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

/*PutStatusProbeOK Success

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

// PutStatusProbeFailedCode is the HTTP code returned for type PutStatusProbeFailed
const PutStatusProbeFailedCode int = 500

/*PutStatusProbeFailed Internal error occurred while conducting connectivity probe

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
