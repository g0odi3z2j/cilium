// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetMapNameOKCode is the HTTP code returned for type GetMapNameOK
const GetMapNameOKCode int = 200

/*GetMapNameOK Success

swagger:response getMapNameOK
*/
type GetMapNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.BPFMap `json:"body,omitempty"`
}

// NewGetMapNameOK creates GetMapNameOK with default headers values
func NewGetMapNameOK() *GetMapNameOK {

	return &GetMapNameOK{}
}

// WithPayload adds the payload to the get map name o k response
func (o *GetMapNameOK) WithPayload(payload *models.BPFMap) *GetMapNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get map name o k response
func (o *GetMapNameOK) SetPayload(payload *models.BPFMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMapNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMapNameNotFoundCode is the HTTP code returned for type GetMapNameNotFound
const GetMapNameNotFoundCode int = 404

/*GetMapNameNotFound Map not found

swagger:response getMapNameNotFound
*/
type GetMapNameNotFound struct {
}

// NewGetMapNameNotFound creates GetMapNameNotFound with default headers values
func NewGetMapNameNotFound() *GetMapNameNotFound {

	return &GetMapNameNotFound{}
}

// WriteResponse to the client
func (o *GetMapNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
