// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIdentityIDHandlerFunc turns a function with the right signature into a get identity ID handler
type GetIdentityIDHandlerFunc func(GetIdentityIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIdentityIDHandlerFunc) Handle(params GetIdentityIDParams) middleware.Responder {
	return fn(params)
}

// GetIdentityIDHandler interface for that can handle valid get identity ID params
type GetIdentityIDHandler interface {
	Handle(GetIdentityIDParams) middleware.Responder
}

// NewGetIdentityID creates a new http.Handler for the get identity ID operation
func NewGetIdentityID(ctx *middleware.Context, handler GetIdentityIDHandler) *GetIdentityID {
	return &GetIdentityID{Context: ctx, Handler: handler}
}

/*GetIdentityID swagger:route GET /identity/{id} policy getIdentityId

Retrieve identity

*/
type GetIdentityID struct {
	Context *middleware.Context
	Handler GetIdentityIDHandler
}

func (o *GetIdentityID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIdentityIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
