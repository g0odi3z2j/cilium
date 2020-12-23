// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHealthzHandlerFunc turns a function with the right signature into a get healthz handler
type GetHealthzHandlerFunc func(GetHealthzParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHealthzHandlerFunc) Handle(params GetHealthzParams) middleware.Responder {
	return fn(params)
}

// GetHealthzHandler interface for that can handle valid get healthz params
type GetHealthzHandler interface {
	Handle(GetHealthzParams) middleware.Responder
}

// NewGetHealthz creates a new http.Handler for the get healthz operation
func NewGetHealthz(ctx *middleware.Context, handler GetHealthzHandler) *GetHealthz {
	return &GetHealthz{Context: ctx, Handler: handler}
}

/*GetHealthz swagger:route GET /healthz operator getHealthz

Get health of Cilium operator

This path will return the status of cilium operator instance.

*/
type GetHealthz struct {
	Context *middleware.Context
	Handler GetHealthzHandler
}

func (o *GetHealthz) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHealthzParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
