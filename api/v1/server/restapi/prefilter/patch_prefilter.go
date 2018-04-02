// Code generated by go-swagger; DO NOT EDIT.

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchPrefilterHandlerFunc turns a function with the right signature into a patch prefilter handler
type PatchPrefilterHandlerFunc func(PatchPrefilterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPrefilterHandlerFunc) Handle(params PatchPrefilterParams) middleware.Responder {
	return fn(params)
}

// PatchPrefilterHandler interface for that can handle valid patch prefilter params
type PatchPrefilterHandler interface {
	Handle(PatchPrefilterParams) middleware.Responder
}

// NewPatchPrefilter creates a new http.Handler for the patch prefilter operation
func NewPatchPrefilter(ctx *middleware.Context, handler PatchPrefilterHandler) *PatchPrefilter {
	return &PatchPrefilter{Context: ctx, Handler: handler}
}

/*PatchPrefilter swagger:route PATCH /prefilter prefilter patchPrefilter

Update list of CIDRs

*/
type PatchPrefilter struct {
	Context *middleware.Context
	Handler PatchPrefilterHandler
}

func (o *PatchPrefilter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchPrefilterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
