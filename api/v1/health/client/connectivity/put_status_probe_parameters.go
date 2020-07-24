// Code generated by go-swagger; DO NOT EDIT.

package connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPutStatusProbeParams creates a new PutStatusProbeParams object
// with the default values initialized.
func NewPutStatusProbeParams() *PutStatusProbeParams {

	return &PutStatusProbeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutStatusProbeParamsWithTimeout creates a new PutStatusProbeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutStatusProbeParamsWithTimeout(timeout time.Duration) *PutStatusProbeParams {

	return &PutStatusProbeParams{

		timeout: timeout,
	}
}

// NewPutStatusProbeParamsWithContext creates a new PutStatusProbeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutStatusProbeParamsWithContext(ctx context.Context) *PutStatusProbeParams {

	return &PutStatusProbeParams{

		Context: ctx,
	}
}

// NewPutStatusProbeParamsWithHTTPClient creates a new PutStatusProbeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutStatusProbeParamsWithHTTPClient(client *http.Client) *PutStatusProbeParams {

	return &PutStatusProbeParams{
		HTTPClient: client,
	}
}

/*PutStatusProbeParams contains all the parameters to send to the API endpoint
for the put status probe operation typically these are written to a http.Request
*/
type PutStatusProbeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put status probe params
func (o *PutStatusProbeParams) WithTimeout(timeout time.Duration) *PutStatusProbeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put status probe params
func (o *PutStatusProbeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put status probe params
func (o *PutStatusProbeParams) WithContext(ctx context.Context) *PutStatusProbeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put status probe params
func (o *PutStatusProbeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put status probe params
func (o *PutStatusProbeParams) WithHTTPClient(client *http.Client) *PutStatusProbeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put status probe params
func (o *PutStatusProbeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PutStatusProbeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
