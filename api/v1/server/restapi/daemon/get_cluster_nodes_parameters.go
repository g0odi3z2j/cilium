// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterNodesParams creates a new GetClusterNodesParams object
// no default values defined in spec.
func NewGetClusterNodesParams() GetClusterNodesParams {

	return GetClusterNodesParams{}
}

// GetClusterNodesParams contains all the bound params for the get cluster nodes operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetClusterNodes
type GetClusterNodesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Client UUID should be used when the client wants to request
	a diff of nodes added and / or removed since the last time
	that client has made a request.

	  In: header
	*/
	ClientID *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetClusterNodesParams() beforehand.
func (o *GetClusterNodesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindClientID(r.Header[http.CanonicalHeaderKey("client-id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClientID binds and validates parameter ClientID from header.
func (o *GetClusterNodesParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("client-id", "header", "int64", raw)
	}
	o.ClientID = &value

	return nil
}
