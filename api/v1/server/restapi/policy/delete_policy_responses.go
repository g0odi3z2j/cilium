// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/cilium/cilium/api/v1/models"
)

// DeletePolicyOKCode is the HTTP code returned for type DeletePolicyOK
const DeletePolicyOKCode int = 200

/*DeletePolicyOK Success

swagger:response deletePolicyOK
*/
type DeletePolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Policy `json:"body,omitempty"`
}

// NewDeletePolicyOK creates DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {

	return &DeletePolicyOK{}
}

// WithPayload adds the payload to the delete policy o k response
func (o *DeletePolicyOK) WithPayload(payload *models.Policy) *DeletePolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy o k response
func (o *DeletePolicyOK) SetPayload(payload *models.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePolicyInvalidCode is the HTTP code returned for type DeletePolicyInvalid
const DeletePolicyInvalidCode int = 400

/*DeletePolicyInvalid Invalid request

swagger:response deletePolicyInvalid
*/
type DeletePolicyInvalid struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyInvalid creates DeletePolicyInvalid with default headers values
func NewDeletePolicyInvalid() *DeletePolicyInvalid {

	return &DeletePolicyInvalid{}
}

// WithPayload adds the payload to the delete policy invalid response
func (o *DeletePolicyInvalid) WithPayload(payload models.Error) *DeletePolicyInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy invalid response
func (o *DeletePolicyInvalid) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DeletePolicyNotFoundCode is the HTTP code returned for type DeletePolicyNotFound
const DeletePolicyNotFoundCode int = 404

/*DeletePolicyNotFound Policy not found

swagger:response deletePolicyNotFound
*/
type DeletePolicyNotFound struct {
}

// NewDeletePolicyNotFound creates DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {

	return &DeletePolicyNotFound{}
}

// WriteResponse to the client
func (o *DeletePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeletePolicyFailureCode is the HTTP code returned for type DeletePolicyFailure
const DeletePolicyFailureCode int = 500

/*DeletePolicyFailure Error while deleting policy

swagger:response deletePolicyFailure
*/
type DeletePolicyFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyFailure creates DeletePolicyFailure with default headers values
func NewDeletePolicyFailure() *DeletePolicyFailure {

	return &DeletePolicyFailure{}
}

// WithPayload adds the payload to the delete policy failure response
func (o *DeletePolicyFailure) WithPayload(payload models.Error) *DeletePolicyFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy failure response
func (o *DeletePolicyFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
