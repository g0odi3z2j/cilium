// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEndpointIDConfigHandlerFunc turns a function with the right signature into a get endpoint ID config handler
type GetEndpointIDConfigHandlerFunc func(GetEndpointIDConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointIDConfigHandlerFunc) Handle(params GetEndpointIDConfigParams) middleware.Responder {
	return fn(params)
}

// GetEndpointIDConfigHandler interface for that can handle valid get endpoint ID config params
type GetEndpointIDConfigHandler interface {
	Handle(GetEndpointIDConfigParams) middleware.Responder
}

// NewGetEndpointIDConfig creates a new http.Handler for the get endpoint ID config operation
func NewGetEndpointIDConfig(ctx *middleware.Context, handler GetEndpointIDConfigHandler) *GetEndpointIDConfig {
	return &GetEndpointIDConfig{Context: ctx, Handler: handler}
}

/*GetEndpointIDConfig swagger:route GET /endpoint/{id}/config endpoint getEndpointIdConfig

Retrieve endpoint configuration

Retrieves the configuration of the specified endpoint.


*/
type GetEndpointIDConfig struct {
	Context *middleware.Context
	Handler GetEndpointIDConfigHandler
}

func (o *GetEndpointIDConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEndpointIDConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
