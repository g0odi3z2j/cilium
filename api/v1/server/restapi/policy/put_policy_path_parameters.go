package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutPolicyPathParams creates a new PutPolicyPathParams object
// with the default values initialized.
func NewPutPolicyPathParams() PutPolicyPathParams {
	var ()
	return PutPolicyPathParams{}
}

// PutPolicyPathParams contains all the bound params for the put policy path operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutPolicyPath
type PutPolicyPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Path to policy node
	  Required: true
	  In: path
	*/
	Path string
	/*Policy tree or subtree
	  Required: true
	  In: body
	*/
	Policy *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutPolicyPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

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

func (o *PutPolicyPathParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Path = raw

	return nil
}
