// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// NewPutServiceIDParams creates a new PutServiceIDParams object
// with the default values initialized.
func NewPutServiceIDParams() *PutServiceIDParams {
	var ()
	return &PutServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutServiceIDParamsWithTimeout creates a new PutServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutServiceIDParamsWithTimeout(timeout time.Duration) *PutServiceIDParams {
	var ()
	return &PutServiceIDParams{

		timeout: timeout,
	}
}

// NewPutServiceIDParamsWithContext creates a new PutServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutServiceIDParamsWithContext(ctx context.Context) *PutServiceIDParams {
	var ()
	return &PutServiceIDParams{

		Context: ctx,
	}
}

// NewPutServiceIDParamsWithHTTPClient creates a new PutServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutServiceIDParamsWithHTTPClient(client *http.Client) *PutServiceIDParams {
	var ()
	return &PutServiceIDParams{
		HTTPClient: client,
	}
}

/*PutServiceIDParams contains all the parameters to send to the API endpoint
for the put service ID operation typically these are written to a http.Request
*/
type PutServiceIDParams struct {

	/*Config
	  Service configuration

	*/
	Config *models.ServiceSpec
	/*ID
	  ID of service

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put service ID params
func (o *PutServiceIDParams) WithTimeout(timeout time.Duration) *PutServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put service ID params
func (o *PutServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put service ID params
func (o *PutServiceIDParams) WithContext(ctx context.Context) *PutServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put service ID params
func (o *PutServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put service ID params
func (o *PutServiceIDParams) WithHTTPClient(client *http.Client) *PutServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put service ID params
func (o *PutServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put service ID params
func (o *PutServiceIDParams) WithConfig(config *models.ServiceSpec) *PutServiceIDParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put service ID params
func (o *PutServiceIDParams) SetConfig(config *models.ServiceSpec) {
	o.Config = config
}

// WithID adds the id to the put service ID params
func (o *PutServiceIDParams) WithID(id int64) *PutServiceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put service ID params
func (o *PutServiceIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
