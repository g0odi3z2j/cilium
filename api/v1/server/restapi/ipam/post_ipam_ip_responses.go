// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIpamIPOKCode is the HTTP code returned for type PostIpamIPOK
const PostIpamIPOKCode int = 200

/*PostIpamIPOK Success

swagger:response postIpamIpOK
*/
type PostIpamIPOK struct {
}

// NewPostIpamIPOK creates PostIpamIPOK with default headers values
func NewPostIpamIPOK() *PostIpamIPOK {

	return &PostIpamIPOK{}
}

// WriteResponse to the client
func (o *PostIpamIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostIpamIPInvalidCode is the HTTP code returned for type PostIpamIPInvalid
const PostIpamIPInvalidCode int = 400

/*PostIpamIPInvalid Invalid IP address

swagger:response postIpamIpInvalid
*/
type PostIpamIPInvalid struct {
}

// NewPostIpamIPInvalid creates PostIpamIPInvalid with default headers values
func NewPostIpamIPInvalid() *PostIpamIPInvalid {

	return &PostIpamIPInvalid{}
}

// WriteResponse to the client
func (o *PostIpamIPInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostIpamIPExistsCode is the HTTP code returned for type PostIpamIPExists
const PostIpamIPExistsCode int = 409

/*PostIpamIPExists IP already allocated

swagger:response postIpamIpExists
*/
type PostIpamIPExists struct {
}

// NewPostIpamIPExists creates PostIpamIPExists with default headers values
func NewPostIpamIPExists() *PostIpamIPExists {

	return &PostIpamIPExists{}
}

// WriteResponse to the client
func (o *PostIpamIPExists) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// PostIpamIPFailureCode is the HTTP code returned for type PostIpamIPFailure
const PostIpamIPFailureCode int = 500

/*PostIpamIPFailure IP allocation failure. Details in message.

swagger:response postIpamIpFailure
*/
type PostIpamIPFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostIpamIPFailure creates PostIpamIPFailure with default headers values
func NewPostIpamIPFailure() *PostIpamIPFailure {

	return &PostIpamIPFailure{}
}

// WithPayload adds the payload to the post ipam Ip failure response
func (o *PostIpamIPFailure) WithPayload(payload models.Error) *PostIpamIPFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ipam Ip failure response
func (o *PostIpamIPFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIpamIPFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostIpamIPDisabledCode is the HTTP code returned for type PostIpamIPDisabled
const PostIpamIPDisabledCode int = 501

/*PostIpamIPDisabled Allocation for address family disabled

swagger:response postIpamIpDisabled
*/
type PostIpamIPDisabled struct {
}

// NewPostIpamIPDisabled creates PostIpamIPDisabled with default headers values
func NewPostIpamIPDisabled() *PostIpamIPDisabled {

	return &PostIpamIPDisabled{}
}

// WriteResponse to the client
func (o *PostIpamIPDisabled) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
