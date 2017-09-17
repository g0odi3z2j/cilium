// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewPutPolicyParams creates a new PutPolicyParams object
// with the default values initialized.
func NewPutPolicyParams() PutPolicyParams {
	var ()
	return PutPolicyParams{}
}

// PutPolicyParams contains all the bound params for the put policy operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutPolicy
type PutPolicyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Policy rules
	  Required: true
	  In: body
	*/
	Policy *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutPolicyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("policy", "body"))
			} else {
				res = append(res, errors.NewParseError("policy", "body", "", err))
			}

		} else {

			if len(res) == 0 {
				o.Policy = &body
			}
		}

	} else {
		res = append(res, errors.Required("policy", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
