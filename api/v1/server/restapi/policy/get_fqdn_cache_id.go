// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFqdnCacheIDHandlerFunc turns a function with the right signature into a get fqdn cache ID handler
type GetFqdnCacheIDHandlerFunc func(GetFqdnCacheIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFqdnCacheIDHandlerFunc) Handle(params GetFqdnCacheIDParams) middleware.Responder {
	return fn(params)
}

// GetFqdnCacheIDHandler interface for that can handle valid get fqdn cache ID params
type GetFqdnCacheIDHandler interface {
	Handle(GetFqdnCacheIDParams) middleware.Responder
}

// NewGetFqdnCacheID creates a new http.Handler for the get fqdn cache ID operation
func NewGetFqdnCacheID(ctx *middleware.Context, handler GetFqdnCacheIDHandler) *GetFqdnCacheID {
	return &GetFqdnCacheID{Context: ctx, Handler: handler}
}

/*GetFqdnCacheID swagger:route GET /fqdn/cache/{id} policy getFqdnCacheId

Retrieves the list of DNS lookups intercepted from an endpoint.

Retrieves the list of DNS lookups intercepted from endpoints,
optionally filtered by endpoint id, DNS name, or CIDR IP range.


*/
type GetFqdnCacheID struct {
	Context *middleware.Context
	Handler GetFqdnCacheIDHandler
}

func (o *GetFqdnCacheID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFqdnCacheIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
