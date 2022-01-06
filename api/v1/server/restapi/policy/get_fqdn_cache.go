// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFqdnCacheHandlerFunc turns a function with the right signature into a get fqdn cache handler
type GetFqdnCacheHandlerFunc func(GetFqdnCacheParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFqdnCacheHandlerFunc) Handle(params GetFqdnCacheParams) middleware.Responder {
	return fn(params)
}

// GetFqdnCacheHandler interface for that can handle valid get fqdn cache params
type GetFqdnCacheHandler interface {
	Handle(GetFqdnCacheParams) middleware.Responder
}

// NewGetFqdnCache creates a new http.Handler for the get fqdn cache operation
func NewGetFqdnCache(ctx *middleware.Context, handler GetFqdnCacheHandler) *GetFqdnCache {
	return &GetFqdnCache{Context: ctx, Handler: handler}
}

/*GetFqdnCache swagger:route GET /fqdn/cache policy getFqdnCache

Retrieves the list of DNS lookups intercepted from all endpoints.

Retrieves the list of DNS lookups intercepted from endpoints,
optionally filtered by endpoint id, DNS name, or CIDR IP range.


*/
type GetFqdnCache struct {
	Context *middleware.Context
	Handler GetFqdnCacheHandler
}

func (o *GetFqdnCache) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFqdnCacheParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
