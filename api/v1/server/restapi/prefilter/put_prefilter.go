// Code generated by go-swagger; DO NOT EDIT.

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutPrefilterHandlerFunc turns a function with the right signature into a put prefilter handler
type PutPrefilterHandlerFunc func(PutPrefilterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPrefilterHandlerFunc) Handle(params PutPrefilterParams) middleware.Responder {
	return fn(params)
}

// PutPrefilterHandler interface for that can handle valid put prefilter params
type PutPrefilterHandler interface {
	Handle(PutPrefilterParams) middleware.Responder
}

// NewPutPrefilter creates a new http.Handler for the put prefilter operation
func NewPutPrefilter(ctx *middleware.Context, handler PutPrefilterHandler) *PutPrefilter {
	return &PutPrefilter{Context: ctx, Handler: handler}
}

/*PutPrefilter swagger:route PUT /prefilter prefilter putPrefilter

Update list of CIDRs

*/
type PutPrefilter struct {
	Context *middleware.Context
	Handler PutPrefilterHandler
}

func (o *PutPrefilter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPrefilterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
