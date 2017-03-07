package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetPolicyOKCode is the HTTP code returned for type GetPolicyOK
const GetPolicyOKCode int = 200

/*GetPolicyOK Success

swagger:response getPolicyOK
*/
type GetPolicyOK struct {

	/*
	  In: Body
	*/
	Payload models.PolicyTree `json:"body,omitempty"`
}

// NewGetPolicyOK creates GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {
	return &GetPolicyOK{}
}

// WithPayload adds the payload to the get policy o k response
func (o *GetPolicyOK) WithPayload(payload models.PolicyTree) *GetPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy o k response
func (o *GetPolicyOK) SetPayload(payload models.PolicyTree) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
