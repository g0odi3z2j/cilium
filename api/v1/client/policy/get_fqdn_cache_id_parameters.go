// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFqdnCacheIDParams creates a new GetFqdnCacheIDParams object
// with the default values initialized.
func NewGetFqdnCacheIDParams() *GetFqdnCacheIDParams {
	var ()
	return &GetFqdnCacheIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFqdnCacheIDParamsWithTimeout creates a new GetFqdnCacheIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFqdnCacheIDParamsWithTimeout(timeout time.Duration) *GetFqdnCacheIDParams {
	var ()
	return &GetFqdnCacheIDParams{

		timeout: timeout,
	}
}

// NewGetFqdnCacheIDParamsWithContext creates a new GetFqdnCacheIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFqdnCacheIDParamsWithContext(ctx context.Context) *GetFqdnCacheIDParams {
	var ()
	return &GetFqdnCacheIDParams{

		Context: ctx,
	}
}

// NewGetFqdnCacheIDParamsWithHTTPClient creates a new GetFqdnCacheIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFqdnCacheIDParamsWithHTTPClient(client *http.Client) *GetFqdnCacheIDParams {
	var ()
	return &GetFqdnCacheIDParams{
		HTTPClient: client,
	}
}

/*GetFqdnCacheIDParams contains all the parameters to send to the API endpoint
for the get fqdn cache ID operation typically these are written to a http.Request
*/
type GetFqdnCacheIDParams struct {

	/*Cidr
	  A CIDR range of IPs

	*/
	Cidr *string
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
	/*Matchpattern
	  A toFQDNs compatible matchPattern expression

	*/
	Matchpattern *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithTimeout(timeout time.Duration) *GetFqdnCacheIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithContext(ctx context.Context) *GetFqdnCacheIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithHTTPClient(client *http.Client) *GetFqdnCacheIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCidr adds the cidr to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithCidr(cidr *string) *GetFqdnCacheIDParams {
	o.SetCidr(cidr)
	return o
}

// SetCidr adds the cidr to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetCidr(cidr *string) {
	o.Cidr = cidr
}

// WithID adds the id to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithID(id string) *GetFqdnCacheIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetID(id string) {
	o.ID = id
}

// WithMatchpattern adds the matchpattern to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) WithMatchpattern(matchpattern *string) *GetFqdnCacheIDParams {
	o.SetMatchpattern(matchpattern)
	return o
}

// SetMatchpattern adds the matchpattern to the get fqdn cache ID params
func (o *GetFqdnCacheIDParams) SetMatchpattern(matchpattern *string) {
	o.Matchpattern = matchpattern
}

// WriteToRequest writes these params to a swagger request
func (o *GetFqdnCacheIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cidr != nil {

		// query param cidr
		var qrCidr string
		if o.Cidr != nil {
			qrCidr = *o.Cidr
		}
		qCidr := qrCidr
		if qCidr != "" {
			if err := r.SetQueryParam("cidr", qCidr); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Matchpattern != nil {

		// query param matchpattern
		var qrMatchpattern string
		if o.Matchpattern != nil {
			qrMatchpattern = *o.Matchpattern
		}
		qMatchpattern := qrMatchpattern
		if qMatchpattern != "" {
			if err := r.SetQueryParam("matchpattern", qMatchpattern); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
