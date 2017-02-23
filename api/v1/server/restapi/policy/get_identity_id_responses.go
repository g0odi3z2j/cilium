package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// HTTP code for type GetIdentityIDOK
const GetIdentityIDOKCode int = 200

/*GetIdentityIDOK Success

swagger:response getIdentityIdOK
*/
type GetIdentityIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Identity `json:"body,omitempty"`
}

// NewGetIdentityIDOK creates GetIdentityIDOK with default headers values
func NewGetIdentityIDOK() *GetIdentityIDOK {
	return &GetIdentityIDOK{}
}

// WithPayload adds the payload to the get identity Id o k response
func (o *GetIdentityIDOK) WithPayload(payload *models.Identity) *GetIdentityIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity Id o k response
func (o *GetIdentityIDOK) SetPayload(payload *models.Identity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type GetIdentityIDBadRequest
const GetIdentityIDBadRequestCode int = 400

/*GetIdentityIDBadRequest Invalid identity provided

swagger:response getIdentityIdBadRequest
*/
type GetIdentityIDBadRequest struct {
}

// NewGetIdentityIDBadRequest creates GetIdentityIDBadRequest with default headers values
func NewGetIdentityIDBadRequest() *GetIdentityIDBadRequest {
	return &GetIdentityIDBadRequest{}
}

// WriteResponse to the client
func (o *GetIdentityIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// HTTP code for type GetIdentityIDNotFound
const GetIdentityIDNotFoundCode int = 404

/*GetIdentityIDNotFound Identity not found

swagger:response getIdentityIdNotFound
*/
type GetIdentityIDNotFound struct {
}

// NewGetIdentityIDNotFound creates GetIdentityIDNotFound with default headers values
func NewGetIdentityIDNotFound() *GetIdentityIDNotFound {
	return &GetIdentityIDNotFound{}
}

// WriteResponse to the client
func (o *GetIdentityIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// HTTP code for type GetIdentityIDUnreachable
const GetIdentityIDUnreachableCode int = 520

/*GetIdentityIDUnreachable Identity storage unreachable. Likely a network problem.

swagger:response getIdentityIdUnreachable
*/
type GetIdentityIDUnreachable struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetIdentityIDUnreachable creates GetIdentityIDUnreachable with default headers values
func NewGetIdentityIDUnreachable() *GetIdentityIDUnreachable {
	return &GetIdentityIDUnreachable{}
}

// WithPayload adds the payload to the get identity Id unreachable response
func (o *GetIdentityIDUnreachable) WithPayload(payload models.Error) *GetIdentityIDUnreachable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity Id unreachable response
func (o *GetIdentityIDUnreachable) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityIDUnreachable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(520)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type GetIdentityIDInvalidStorageFormat
const GetIdentityIDInvalidStorageFormatCode int = 521

/*GetIdentityIDInvalidStorageFormat Invalid identity format in storage

swagger:response getIdentityIdInvalidStorageFormat
*/
type GetIdentityIDInvalidStorageFormat struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetIdentityIDInvalidStorageFormat creates GetIdentityIDInvalidStorageFormat with default headers values
func NewGetIdentityIDInvalidStorageFormat() *GetIdentityIDInvalidStorageFormat {
	return &GetIdentityIDInvalidStorageFormat{}
}

// WithPayload adds the payload to the get identity Id invalid storage format response
func (o *GetIdentityIDInvalidStorageFormat) WithPayload(payload models.Error) *GetIdentityIDInvalidStorageFormat {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity Id invalid storage format response
func (o *GetIdentityIDInvalidStorageFormat) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityIDInvalidStorageFormat) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(521)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
