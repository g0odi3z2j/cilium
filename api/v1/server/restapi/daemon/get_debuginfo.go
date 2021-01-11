// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDebuginfoHandlerFunc turns a function with the right signature into a get debuginfo handler
type GetDebuginfoHandlerFunc func(GetDebuginfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDebuginfoHandlerFunc) Handle(params GetDebuginfoParams) middleware.Responder {
	return fn(params)
}

// GetDebuginfoHandler interface for that can handle valid get debuginfo params
type GetDebuginfoHandler interface {
	Handle(GetDebuginfoParams) middleware.Responder
}

// NewGetDebuginfo creates a new http.Handler for the get debuginfo operation
func NewGetDebuginfo(ctx *middleware.Context, handler GetDebuginfoHandler) *GetDebuginfo {
	return &GetDebuginfo{Context: ctx, Handler: handler}
}

/*GetDebuginfo swagger:route GET /debuginfo daemon getDebuginfo

Retrieve information about the agent and evironment for debugging

*/
type GetDebuginfo struct {
	Context *middleware.Context
	Handler GetDebuginfoHandler
}

func (o *GetDebuginfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDebuginfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
