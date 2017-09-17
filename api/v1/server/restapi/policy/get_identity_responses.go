// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetIdentityOKCode is the HTTP code returned for type GetIdentityOK
const GetIdentityOKCode int = 200

/*GetIdentityOK Success

swagger:response getIdentityOK
*/
type GetIdentityOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Identity `json:"body,omitempty"`
}

// NewGetIdentityOK creates GetIdentityOK with default headers values
func NewGetIdentityOK() *GetIdentityOK {
	return &GetIdentityOK{}
}

// WithPayload adds the payload to the get identity o k response
func (o *GetIdentityOK) WithPayload(payload []*models.Identity) *GetIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity o k response
func (o *GetIdentityOK) SetPayload(payload []*models.Identity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Identity, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetIdentityNotFoundCode is the HTTP code returned for type GetIdentityNotFound
const GetIdentityNotFoundCode int = 404

/*GetIdentityNotFound Identities with provided parameters not found

swagger:response getIdentityNotFound
*/
type GetIdentityNotFound struct {
}

// NewGetIdentityNotFound creates GetIdentityNotFound with default headers values
func NewGetIdentityNotFound() *GetIdentityNotFound {
	return &GetIdentityNotFound{}
}

// WriteResponse to the client
func (o *GetIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetIdentityUnreachableCode is the HTTP code returned for type GetIdentityUnreachable
const GetIdentityUnreachableCode int = 520

/*GetIdentityUnreachable Identity storage unreachable. Likely a network problem.

swagger:response getIdentityUnreachable
*/
type GetIdentityUnreachable struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetIdentityUnreachable creates GetIdentityUnreachable with default headers values
func NewGetIdentityUnreachable() *GetIdentityUnreachable {
	return &GetIdentityUnreachable{}
}

// WithPayload adds the payload to the get identity unreachable response
func (o *GetIdentityUnreachable) WithPayload(payload models.Error) *GetIdentityUnreachable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity unreachable response
func (o *GetIdentityUnreachable) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityUnreachable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(520)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetIdentityInvalidStorageFormatCode is the HTTP code returned for type GetIdentityInvalidStorageFormat
const GetIdentityInvalidStorageFormatCode int = 521

/*GetIdentityInvalidStorageFormat Invalid identity format in storage

swagger:response getIdentityInvalidStorageFormat
*/
type GetIdentityInvalidStorageFormat struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetIdentityInvalidStorageFormat creates GetIdentityInvalidStorageFormat with default headers values
func NewGetIdentityInvalidStorageFormat() *GetIdentityInvalidStorageFormat {
	return &GetIdentityInvalidStorageFormat{}
}

// WithPayload adds the payload to the get identity invalid storage format response
func (o *GetIdentityInvalidStorageFormat) WithPayload(payload models.Error) *GetIdentityInvalidStorageFormat {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity invalid storage format response
func (o *GetIdentityInvalidStorageFormat) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityInvalidStorageFormat) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(521)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
