// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPrefilterHandlerFunc turns a function with the right signature into a get prefilter handler
type GetPrefilterHandlerFunc func(GetPrefilterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPrefilterHandlerFunc) Handle(params GetPrefilterParams) middleware.Responder {
	return fn(params)
}

// GetPrefilterHandler interface for that can handle valid get prefilter params
type GetPrefilterHandler interface {
	Handle(GetPrefilterParams) middleware.Responder
}

// NewGetPrefilter creates a new http.Handler for the get prefilter operation
func NewGetPrefilter(ctx *middleware.Context, handler GetPrefilterHandler) *GetPrefilter {
	return &GetPrefilter{Context: ctx, Handler: handler}
}

/*
	GetPrefilter swagger:route GET /prefilter prefilter getPrefilter

Retrieve list of CIDRs
*/
type GetPrefilter struct {
	Context *middleware.Context
	Handler GetPrefilterHandler
}

func (o *GetPrefilter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPrefilterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
