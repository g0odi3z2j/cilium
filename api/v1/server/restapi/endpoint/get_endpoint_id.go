// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEndpointIDHandlerFunc turns a function with the right signature into a get endpoint ID handler
type GetEndpointIDHandlerFunc func(GetEndpointIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointIDHandlerFunc) Handle(params GetEndpointIDParams) middleware.Responder {
	return fn(params)
}

// GetEndpointIDHandler interface for that can handle valid get endpoint ID params
type GetEndpointIDHandler interface {
	Handle(GetEndpointIDParams) middleware.Responder
}

// NewGetEndpointID creates a new http.Handler for the get endpoint ID operation
func NewGetEndpointID(ctx *middleware.Context, handler GetEndpointIDHandler) *GetEndpointID {
	return &GetEndpointID{Context: ctx, Handler: handler}
}

/*GetEndpointID swagger:route GET /endpoint/{id} endpoint getEndpointId

Get endpoint by endpoint ID

Returns endpoint information


*/
type GetEndpointID struct {
	Context *middleware.Context
	Handler GetEndpointIDHandler
}

func (o *GetEndpointID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEndpointIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
