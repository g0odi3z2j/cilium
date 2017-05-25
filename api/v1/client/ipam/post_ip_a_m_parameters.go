package ipam

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

// NewPostIPAMParams creates a new PostIPAMParams object
// with the default values initialized.
func NewPostIPAMParams() *PostIPAMParams {
	var ()
	return &PostIPAMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPAMParamsWithTimeout creates a new PostIPAMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPAMParamsWithTimeout(timeout time.Duration) *PostIPAMParams {
	var ()
	return &PostIPAMParams{

		timeout: timeout,
	}
}

// NewPostIPAMParamsWithContext creates a new PostIPAMParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPAMParamsWithContext(ctx context.Context) *PostIPAMParams {
	var ()
	return &PostIPAMParams{

		Context: ctx,
	}
}

// NewPostIPAMParamsWithHTTPClient creates a new PostIPAMParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPAMParamsWithHTTPClient(client *http.Client) *PostIPAMParams {
	var ()
	return &PostIPAMParams{
		HTTPClient: client,
	}
}

/*PostIPAMParams contains all the parameters to send to the API endpoint
for the post IP a m operation typically these are written to a http.Request
*/
type PostIPAMParams struct {

	/*Family*/
	Family *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP a m params
func (o *PostIPAMParams) WithTimeout(timeout time.Duration) *PostIPAMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP a m params
func (o *PostIPAMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP a m params
func (o *PostIPAMParams) WithContext(ctx context.Context) *PostIPAMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP a m params
func (o *PostIPAMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP a m params
func (o *PostIPAMParams) WithHTTPClient(client *http.Client) *PostIPAMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP a m params
func (o *PostIPAMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFamily adds the family to the post IP a m params
func (o *PostIPAMParams) WithFamily(family *string) *PostIPAMParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the post IP a m params
func (o *PostIPAMParams) SetFamily(family *string) {
	o.Family = family
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPAMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
