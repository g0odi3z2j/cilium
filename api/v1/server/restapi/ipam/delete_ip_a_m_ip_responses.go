// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// DeleteIPAMIPOKCode is the HTTP code returned for type DeleteIPAMIPOK
const DeleteIPAMIPOKCode int = 200

/*DeleteIPAMIPOK Success

swagger:response deleteIpAMIpOK
*/
type DeleteIPAMIPOK struct {
}

// NewDeleteIPAMIPOK creates DeleteIPAMIPOK with default headers values
func NewDeleteIPAMIPOK() *DeleteIPAMIPOK {
	return &DeleteIPAMIPOK{}
}

// WriteResponse to the client
func (o *DeleteIPAMIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteIPAMIPInvalidCode is the HTTP code returned for type DeleteIPAMIPInvalid
const DeleteIPAMIPInvalidCode int = 400

/*DeleteIPAMIPInvalid Invalid IP address

swagger:response deleteIpAMIpInvalid
*/
type DeleteIPAMIPInvalid struct {
}

// NewDeleteIPAMIPInvalid creates DeleteIPAMIPInvalid with default headers values
func NewDeleteIPAMIPInvalid() *DeleteIPAMIPInvalid {
	return &DeleteIPAMIPInvalid{}
}

// WriteResponse to the client
func (o *DeleteIPAMIPInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// DeleteIPAMIPNotFoundCode is the HTTP code returned for type DeleteIPAMIPNotFound
const DeleteIPAMIPNotFoundCode int = 404

/*DeleteIPAMIPNotFound IP address not found

swagger:response deleteIpAMIpNotFound
*/
type DeleteIPAMIPNotFound struct {
}

// NewDeleteIPAMIPNotFound creates DeleteIPAMIPNotFound with default headers values
func NewDeleteIPAMIPNotFound() *DeleteIPAMIPNotFound {
	return &DeleteIPAMIPNotFound{}
}

// WriteResponse to the client
func (o *DeleteIPAMIPNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// DeleteIPAMIPFailureCode is the HTTP code returned for type DeleteIPAMIPFailure
const DeleteIPAMIPFailureCode int = 500

/*DeleteIPAMIPFailure Address release failure

swagger:response deleteIpAMIpFailure
*/
type DeleteIPAMIPFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteIPAMIPFailure creates DeleteIPAMIPFailure with default headers values
func NewDeleteIPAMIPFailure() *DeleteIPAMIPFailure {
	return &DeleteIPAMIPFailure{}
}

// WithPayload adds the payload to the delete Ip a m Ip failure response
func (o *DeleteIPAMIPFailure) WithPayload(payload models.Error) *DeleteIPAMIPFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Ip a m Ip failure response
func (o *DeleteIPAMIPFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIPAMIPFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DeleteIPAMIPDisabledCode is the HTTP code returned for type DeleteIPAMIPDisabled
const DeleteIPAMIPDisabledCode int = 501

/*DeleteIPAMIPDisabled Allocation for address family disabled

swagger:response deleteIpAMIpDisabled
*/
type DeleteIPAMIPDisabled struct {
}

// NewDeleteIPAMIPDisabled creates DeleteIPAMIPDisabled with default headers values
func NewDeleteIPAMIPDisabled() *DeleteIPAMIPDisabled {
	return &DeleteIPAMIPDisabled{}
}

// WriteResponse to the client
func (o *DeleteIPAMIPDisabled) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
