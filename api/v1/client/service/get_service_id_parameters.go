// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServiceIDParams creates a new GetServiceIDParams object
// with the default values initialized.
func NewGetServiceIDParams() *GetServiceIDParams {
	var ()
	return &GetServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceIDParamsWithTimeout creates a new GetServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceIDParamsWithTimeout(timeout time.Duration) *GetServiceIDParams {
	var ()
	return &GetServiceIDParams{

		timeout: timeout,
	}
}

// NewGetServiceIDParamsWithContext creates a new GetServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceIDParamsWithContext(ctx context.Context) *GetServiceIDParams {
	var ()
	return &GetServiceIDParams{

		Context: ctx,
	}
}

// NewGetServiceIDParamsWithHTTPClient creates a new GetServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceIDParamsWithHTTPClient(client *http.Client) *GetServiceIDParams {
	var ()
	return &GetServiceIDParams{
		HTTPClient: client,
	}
}

/*GetServiceIDParams contains all the parameters to send to the API endpoint
for the get service ID operation typically these are written to a http.Request
*/
type GetServiceIDParams struct {

	/*ID
	  ID of service

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service ID params
func (o *GetServiceIDParams) WithTimeout(timeout time.Duration) *GetServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service ID params
func (o *GetServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service ID params
func (o *GetServiceIDParams) WithContext(ctx context.Context) *GetServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service ID params
func (o *GetServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service ID params
func (o *GetServiceIDParams) WithHTTPClient(client *http.Client) *GetServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service ID params
func (o *GetServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get service ID params
func (o *GetServiceIDParams) WithID(id int64) *GetServiceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get service ID params
func (o *GetServiceIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
