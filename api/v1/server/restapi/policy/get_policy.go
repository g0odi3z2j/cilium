// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPolicyHandlerFunc turns a function with the right signature into a get policy handler
type GetPolicyHandlerFunc func(GetPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPolicyHandlerFunc) Handle(params GetPolicyParams) middleware.Responder {
	return fn(params)
}

// GetPolicyHandler interface for that can handle valid get policy params
type GetPolicyHandler interface {
	Handle(GetPolicyParams) middleware.Responder
}

// NewGetPolicy creates a new http.Handler for the get policy operation
func NewGetPolicy(ctx *middleware.Context, handler GetPolicyHandler) *GetPolicy {
	return &GetPolicy{Context: ctx, Handler: handler}
}

/*
GetPolicy swagger:route GET /policy policy getPolicy

# Retrieve entire policy tree

Returns the entire policy tree with all children.
*/
type GetPolicy struct {
	Context *middleware.Context
	Handler GetPolicyHandler
}

func (o *GetPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
