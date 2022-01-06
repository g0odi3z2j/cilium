// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutStatusProbeHandlerFunc turns a function with the right signature into a put status probe handler
type PutStatusProbeHandlerFunc func(PutStatusProbeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutStatusProbeHandlerFunc) Handle(params PutStatusProbeParams) middleware.Responder {
	return fn(params)
}

// PutStatusProbeHandler interface for that can handle valid put status probe params
type PutStatusProbeHandler interface {
	Handle(PutStatusProbeParams) middleware.Responder
}

// NewPutStatusProbe creates a new http.Handler for the put status probe operation
func NewPutStatusProbe(ctx *middleware.Context, handler PutStatusProbeHandler) *PutStatusProbe {
	return &PutStatusProbe{Context: ctx, Handler: handler}
}

/*PutStatusProbe swagger:route PUT /status/probe connectivity putStatusProbe

Run synchronous connectivity probe to determine status of the Cilium cluster

Runs a synchronous probe to all other cilium-health instances and
returns the connectivity status.


*/
type PutStatusProbe struct {
	Context *middleware.Context
	Handler PutStatusProbeHandler
}

func (o *PutStatusProbe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutStatusProbeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
