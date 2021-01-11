// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetServiceIDHandlerFunc turns a function with the right signature into a get service ID handler
type GetServiceIDHandlerFunc func(GetServiceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServiceIDHandlerFunc) Handle(params GetServiceIDParams) middleware.Responder {
	return fn(params)
}

// GetServiceIDHandler interface for that can handle valid get service ID params
type GetServiceIDHandler interface {
	Handle(GetServiceIDParams) middleware.Responder
}

// NewGetServiceID creates a new http.Handler for the get service ID operation
func NewGetServiceID(ctx *middleware.Context, handler GetServiceIDHandler) *GetServiceID {
	return &GetServiceID{Context: ctx, Handler: handler}
}

/*GetServiceID swagger:route GET /service/{id} service getServiceId

Retrieve configuration of a service

*/
type GetServiceID struct {
	Context *middleware.Context
	Handler GetServiceIDHandler
}

func (o *GetServiceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServiceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
