// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchEndpointIDConfigHandlerFunc turns a function with the right signature into a patch endpoint ID config handler
type PatchEndpointIDConfigHandlerFunc func(PatchEndpointIDConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchEndpointIDConfigHandlerFunc) Handle(params PatchEndpointIDConfigParams) middleware.Responder {
	return fn(params)
}

// PatchEndpointIDConfigHandler interface for that can handle valid patch endpoint ID config params
type PatchEndpointIDConfigHandler interface {
	Handle(PatchEndpointIDConfigParams) middleware.Responder
}

// NewPatchEndpointIDConfig creates a new http.Handler for the patch endpoint ID config operation
func NewPatchEndpointIDConfig(ctx *middleware.Context, handler PatchEndpointIDConfigHandler) *PatchEndpointIDConfig {
	return &PatchEndpointIDConfig{Context: ctx, Handler: handler}
}

/*PatchEndpointIDConfig swagger:route PATCH /endpoint/{id}/config endpoint patchEndpointIdConfig

Modify mutable endpoint configuration

Update the configuration of an existing endpoint and regenerates &
recompiles the corresponding programs automatically.


*/
type PatchEndpointIDConfig struct {
	Context *middleware.Context
	Handler PatchEndpointIDConfigHandler
}

func (o *PatchEndpointIDConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchEndpointIDConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
