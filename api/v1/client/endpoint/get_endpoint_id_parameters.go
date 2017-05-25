package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEndpointIDParams creates a new GetEndpointIDParams object
// with the default values initialized.
func NewGetEndpointIDParams() *GetEndpointIDParams {
	var ()
	return &GetEndpointIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointIDParamsWithTimeout creates a new GetEndpointIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndpointIDParamsWithTimeout(timeout time.Duration) *GetEndpointIDParams {
	var ()
	return &GetEndpointIDParams{

		timeout: timeout,
	}
}

// NewGetEndpointIDParamsWithContext creates a new GetEndpointIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndpointIDParamsWithContext(ctx context.Context) *GetEndpointIDParams {
	var ()
	return &GetEndpointIDParams{

		Context: ctx,
	}
}

// NewGetEndpointIDParamsWithHTTPClient creates a new GetEndpointIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndpointIDParamsWithHTTPClient(client *http.Client) *GetEndpointIDParams {
	var ()
	return &GetEndpointIDParams{
		HTTPClient: client,
	}
}

/*GetEndpointIDParams contains all the parameters to send to the API endpoint
for the get endpoint ID operation typically these are written to a http.Request
*/
type GetEndpointIDParams struct {

	/*ID
	  String describing an endpoint with the format `[prefix:]id`. If no prefix
	is specified, a prefix of `cilium-local:` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - container-id: Container runtime ID, e.g. container-id:22222
	  - docker-net-endpoint: Docker libnetwork endpoint ID, e.g. docker-net-endpoint:4444


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endpoint ID params
func (o *GetEndpointIDParams) WithTimeout(timeout time.Duration) *GetEndpointIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint ID params
func (o *GetEndpointIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint ID params
func (o *GetEndpointIDParams) WithContext(ctx context.Context) *GetEndpointIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint ID params
func (o *GetEndpointIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint ID params
func (o *GetEndpointIDParams) WithHTTPClient(client *http.Client) *GetEndpointIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint ID params
func (o *GetEndpointIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get endpoint ID params
func (o *GetEndpointIDParams) WithID(id string) *GetEndpointIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get endpoint ID params
func (o *GetEndpointIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
