package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPolicyResolveHandlerFunc turns a function with the right signature into a get policy resolve handler
type GetPolicyResolveHandlerFunc func(GetPolicyResolveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPolicyResolveHandlerFunc) Handle(params GetPolicyResolveParams) middleware.Responder {
	return fn(params)
}

// GetPolicyResolveHandler interface for that can handle valid get policy resolve params
type GetPolicyResolveHandler interface {
	Handle(GetPolicyResolveParams) middleware.Responder
}

// NewGetPolicyResolve creates a new http.Handler for the get policy resolve operation
func NewGetPolicyResolve(ctx *middleware.Context, handler GetPolicyResolveHandler) *GetPolicyResolve {
	return &GetPolicyResolve{Context: ctx, Handler: handler}
}

/*GetPolicyResolve swagger:route GET /policy/resolve policy getPolicyResolve

Resolve policy for an identity context

*/
type GetPolicyResolve struct {
	Context *middleware.Context
	Handler GetPolicyResolveHandler
}

func (o *GetPolicyResolve) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetPolicyResolveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
