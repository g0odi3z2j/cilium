// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetHelloHandlerFunc turns a function with the right signature into a get hello handler
type GetHelloHandlerFunc func(GetHelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHelloHandlerFunc) Handle(params GetHelloParams) middleware.Responder {
	return fn(params)
}

// GetHelloHandler interface for that can handle valid get hello params
type GetHelloHandler interface {
	Handle(GetHelloParams) middleware.Responder
}

// NewGetHello creates a new http.Handler for the get hello operation
func NewGetHello(ctx *middleware.Context, handler GetHelloHandler) *GetHello {
	return &GetHello{Context: ctx, Handler: handler}
}

/*GetHello swagger:route GET /hello getHello

Say hello to cilium-health

Returns a successful status code if this cilium-health instance is
reachable.


*/
type GetHello struct {
	Context *middleware.Context
	Handler GetHelloHandler
}

func (o *GetHello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHelloParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
