// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteIpamIPHandlerFunc turns a function with the right signature into a delete ipam IP handler
type DeleteIpamIPHandlerFunc func(DeleteIpamIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteIpamIPHandlerFunc) Handle(params DeleteIpamIPParams) middleware.Responder {
	return fn(params)
}

// DeleteIpamIPHandler interface for that can handle valid delete ipam IP params
type DeleteIpamIPHandler interface {
	Handle(DeleteIpamIPParams) middleware.Responder
}

// NewDeleteIpamIP creates a new http.Handler for the delete ipam IP operation
func NewDeleteIpamIP(ctx *middleware.Context, handler DeleteIpamIPHandler) *DeleteIpamIP {
	return &DeleteIpamIP{Context: ctx, Handler: handler}
}

/*DeleteIpamIP swagger:route DELETE /ipam/{ip} ipam deleteIpamIp

Release an allocated IP address

*/
type DeleteIpamIP struct {
	Context *middleware.Context
	Handler DeleteIpamIPHandler
}

func (o *DeleteIpamIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteIpamIPParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
