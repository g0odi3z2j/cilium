// Code generated by go-swagger; DO NOT EDIT.

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

	"github.com/cilium/cilium/api/v1/models"
)

// NewPutEndpointIDParams creates a new PutEndpointIDParams object
// with the default values initialized.
func NewPutEndpointIDParams() *PutEndpointIDParams {
	var ()
	return &PutEndpointIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEndpointIDParamsWithTimeout creates a new PutEndpointIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEndpointIDParamsWithTimeout(timeout time.Duration) *PutEndpointIDParams {
	var ()
	return &PutEndpointIDParams{

		timeout: timeout,
	}
}

// NewPutEndpointIDParamsWithContext creates a new PutEndpointIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEndpointIDParamsWithContext(ctx context.Context) *PutEndpointIDParams {
	var ()
	return &PutEndpointIDParams{

		Context: ctx,
	}
}

// NewPutEndpointIDParamsWithHTTPClient creates a new PutEndpointIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEndpointIDParamsWithHTTPClient(client *http.Client) *PutEndpointIDParams {
	var ()
	return &PutEndpointIDParams{
		HTTPClient: client,
	}
}

/*PutEndpointIDParams contains all the parameters to send to the API endpoint
for the put endpoint ID operation typically these are written to a http.Request
*/
type PutEndpointIDParams struct {

	/*Endpoint*/
	Endpoint *models.EndpointChangeRequest
	/*ID
	  String describing an endpoint with the format ``[prefix:]id``. If no prefix
	is specified, a prefix of ``cilium-local:`` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - container-id: Container runtime ID, e.g. container-id:22222
	  - container-name: Container name, e.g. container-name:foobar
	  - pod-name: pod name for this container if K8s is enabled, e.g. pod-name:default:foobar
	  - docker-endpoint: Docker libnetwork endpoint ID, e.g. docker-endpoint:4444


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put endpoint ID params
func (o *PutEndpointIDParams) WithTimeout(timeout time.Duration) *PutEndpointIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put endpoint ID params
func (o *PutEndpointIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put endpoint ID params
func (o *PutEndpointIDParams) WithContext(ctx context.Context) *PutEndpointIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put endpoint ID params
func (o *PutEndpointIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put endpoint ID params
func (o *PutEndpointIDParams) WithHTTPClient(client *http.Client) *PutEndpointIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put endpoint ID params
func (o *PutEndpointIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpoint adds the endpoint to the put endpoint ID params
func (o *PutEndpointIDParams) WithEndpoint(endpoint *models.EndpointChangeRequest) *PutEndpointIDParams {
	o.SetEndpoint(endpoint)
	return o
}

// SetEndpoint adds the endpoint to the put endpoint ID params
func (o *PutEndpointIDParams) SetEndpoint(endpoint *models.EndpointChangeRequest) {
	o.Endpoint = endpoint
}

// WithID adds the id to the put endpoint ID params
func (o *PutEndpointIDParams) WithID(id string) *PutEndpointIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put endpoint ID params
func (o *PutEndpointIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutEndpointIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Endpoint == nil {
		o.Endpoint = new(models.EndpointChangeRequest)
	}

	if err := r.SetBodyParam(o.Endpoint); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
