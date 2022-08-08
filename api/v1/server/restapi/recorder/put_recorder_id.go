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

// PutRecorderIDHandlerFunc turns a function with the right signature into a put recorder ID handler
type PutRecorderIDHandlerFunc func(PutRecorderIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutRecorderIDHandlerFunc) Handle(params PutRecorderIDParams) middleware.Responder {
	return fn(params)
}

// PutRecorderIDHandler interface for that can handle valid put recorder ID params
type PutRecorderIDHandler interface {
	Handle(PutRecorderIDParams) middleware.Responder
}

// NewPutRecorderID creates a new http.Handler for the put recorder ID operation
func NewPutRecorderID(ctx *middleware.Context, handler PutRecorderIDHandler) *PutRecorderID {
	return &PutRecorderID{Context: ctx, Handler: handler}
}

/*
PutRecorderID swagger:route PUT /recorder/{id} recorder putRecorderId

Create or update recorder
*/
type PutRecorderID struct {
	Context *middleware.Context
	Handler PutRecorderIDHandler
}

func (o *PutRecorderID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutRecorderIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
