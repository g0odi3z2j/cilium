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

// PatchEndpointIDHandlerFunc turns a function with the right signature into a patch endpoint ID handler
type PatchEndpointIDHandlerFunc func(PatchEndpointIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchEndpointIDHandlerFunc) Handle(params PatchEndpointIDParams) middleware.Responder {
	return fn(params)
}

// PatchEndpointIDHandler interface for that can handle valid patch endpoint ID params
type PatchEndpointIDHandler interface {
	Handle(PatchEndpointIDParams) middleware.Responder
}

// NewPatchEndpointID creates a new http.Handler for the patch endpoint ID operation
func NewPatchEndpointID(ctx *middleware.Context, handler PatchEndpointIDHandler) *PatchEndpointID {
	return &PatchEndpointID{Context: ctx, Handler: handler}
}

/*PatchEndpointID swagger:route PATCH /endpoint/{id} endpoint patchEndpointId

Modify existing endpoint

Applies the endpoint change request to an existing endpoint


*/
type PatchEndpointID struct {
	Context *middleware.Context
	Handler PatchEndpointIDHandler
}

func (o *PatchEndpointID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchEndpointIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
