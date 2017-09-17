// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cilium/cilium/api/v1/models"
)

// NewDeletePolicyParams creates a new DeletePolicyParams object
// with the default values initialized.
func NewDeletePolicyParams() DeletePolicyParams {
	var ()
	return DeletePolicyParams{}
}

// DeletePolicyParams contains all the bound params for the delete policy operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeletePolicy
type DeletePolicyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  In: body
	*/
	Labels models.Labels
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeletePolicyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Labels
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("labels", "body", "", err))
		} else {

			if len(res) == 0 {
				o.Labels = body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
