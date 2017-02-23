package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetEndpointHandlerFunc turns a function with the right signature into a get endpoint handler
type GetEndpointHandlerFunc func(GetEndpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointHandlerFunc) Handle(params GetEndpointParams) middleware.Responder {
	return fn(params)
}

// GetEndpointHandler interface for that can handle valid get endpoint params
type GetEndpointHandler interface {
	Handle(GetEndpointParams) middleware.Responder
}

// NewGetEndpoint creates a new http.Handler for the get endpoint operation
func NewGetEndpoint(ctx *middleware.Context, handler GetEndpointHandler) *GetEndpoint {
	return &GetEndpoint{Context: ctx, Handler: handler}
}

/*GetEndpoint swagger:route GET /endpoint endpoint getEndpoint

Get list of all endpoints

Returns an array of all local endpoints.


*/
type GetEndpoint struct {
	Context *middleware.Context
	Handler GetEndpointHandler
}

func (o *GetEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetEndpointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
