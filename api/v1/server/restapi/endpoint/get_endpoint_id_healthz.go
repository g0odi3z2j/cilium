// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEndpointIDHealthzHandlerFunc turns a function with the right signature into a get endpoint ID healthz handler
type GetEndpointIDHealthzHandlerFunc func(GetEndpointIDHealthzParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointIDHealthzHandlerFunc) Handle(params GetEndpointIDHealthzParams) middleware.Responder {
	return fn(params)
}

// GetEndpointIDHealthzHandler interface for that can handle valid get endpoint ID healthz params
type GetEndpointIDHealthzHandler interface {
	Handle(GetEndpointIDHealthzParams) middleware.Responder
}

// NewGetEndpointIDHealthz creates a new http.Handler for the get endpoint ID healthz operation
func NewGetEndpointIDHealthz(ctx *middleware.Context, handler GetEndpointIDHealthzHandler) *GetEndpointIDHealthz {
	return &GetEndpointIDHealthz{Context: ctx, Handler: handler}
}

/*GetEndpointIDHealthz swagger:route GET /endpoint/{id}/healthz endpoint getEndpointIdHealthz

Retrieves the status logs associated with this endpoint.

*/
type GetEndpointIDHealthz struct {
	Context *middleware.Context
	Handler GetEndpointIDHealthzHandler
}

func (o *GetEndpointIDHealthz) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEndpointIDHealthzParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
