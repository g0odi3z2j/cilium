// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewGetMetricsParams creates a new GetMetricsParams object
// with the default values initialized.
func NewGetMetricsParams() *GetMetricsParams {

	return &GetMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricsParamsWithTimeout creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetricsParamsWithTimeout(timeout time.Duration) *GetMetricsParams {

	return &GetMetricsParams{

		timeout: timeout,
	}
}

// NewGetMetricsParamsWithContext creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetricsParamsWithContext(ctx context.Context) *GetMetricsParams {

	return &GetMetricsParams{

		Context: ctx,
	}
}

// NewGetMetricsParamsWithHTTPClient creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetricsParamsWithHTTPClient(client *http.Client) *GetMetricsParams {

	return &GetMetricsParams{
		HTTPClient: client,
	}
}

/*GetMetricsParams contains all the parameters to send to the API endpoint
for the get metrics operation typically these are written to a http.Request
*/
type GetMetricsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) WithTimeout(timeout time.Duration) *GetMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metrics params
func (o *GetMetricsParams) WithContext(ctx context.Context) *GetMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metrics params
func (o *GetMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) WithHTTPClient(client *http.Client) *GetMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
