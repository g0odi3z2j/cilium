package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteServiceIDHandlerFunc turns a function with the right signature into a delete service ID handler
type DeleteServiceIDHandlerFunc func(DeleteServiceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteServiceIDHandlerFunc) Handle(params DeleteServiceIDParams) middleware.Responder {
	return fn(params)
}

// DeleteServiceIDHandler interface for that can handle valid delete service ID params
type DeleteServiceIDHandler interface {
	Handle(DeleteServiceIDParams) middleware.Responder
}

// NewDeleteServiceID creates a new http.Handler for the delete service ID operation
func NewDeleteServiceID(ctx *middleware.Context, handler DeleteServiceIDHandler) *DeleteServiceID {
	return &DeleteServiceID{Context: ctx, Handler: handler}
}

/*DeleteServiceID swagger:route DELETE /service/{id} service deleteServiceId

Delete a service

*/
type DeleteServiceID struct {
	Context *middleware.Context
	Handler DeleteServiceIDHandler
}

func (o *DeleteServiceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteServiceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
