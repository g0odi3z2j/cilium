// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRecorderMasksHandlerFunc turns a function with the right signature into a get recorder masks handler
type GetRecorderMasksHandlerFunc func(GetRecorderMasksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecorderMasksHandlerFunc) Handle(params GetRecorderMasksParams) middleware.Responder {
	return fn(params)
}

// GetRecorderMasksHandler interface for that can handle valid get recorder masks params
type GetRecorderMasksHandler interface {
	Handle(GetRecorderMasksParams) middleware.Responder
}

// NewGetRecorderMasks creates a new http.Handler for the get recorder masks operation
func NewGetRecorderMasks(ctx *middleware.Context, handler GetRecorderMasksHandler) *GetRecorderMasks {
	return &GetRecorderMasks{Context: ctx, Handler: handler}
}

/*
GetRecorderMasks swagger:route GET /recorder/masks recorder getRecorderMasks

Retrieve list of all recorder masks
*/
type GetRecorderMasks struct {
	Context *middleware.Context
	Handler GetRecorderMasksHandler
}

func (o *GetRecorderMasks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRecorderMasksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
