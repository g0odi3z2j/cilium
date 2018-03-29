// Code generated by go-swagger; DO NOT EDIT.

package daemon

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

// NewPatchConfigParams creates a new PatchConfigParams object
// with the default values initialized.
func NewPatchConfigParams() *PatchConfigParams {
	var ()
	return &PatchConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConfigParamsWithTimeout creates a new PatchConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConfigParamsWithTimeout(timeout time.Duration) *PatchConfigParams {
	var ()
	return &PatchConfigParams{

		timeout: timeout,
	}
}

// NewPatchConfigParamsWithContext creates a new PatchConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConfigParamsWithContext(ctx context.Context) *PatchConfigParams {
	var ()
	return &PatchConfigParams{

		Context: ctx,
	}
}

// NewPatchConfigParamsWithHTTPClient creates a new PatchConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConfigParamsWithHTTPClient(client *http.Client) *PatchConfigParams {
	var ()
	return &PatchConfigParams{
		HTTPClient: client,
	}
}

/*PatchConfigParams contains all the parameters to send to the API endpoint
for the patch config operation typically these are written to a http.Request
*/
type PatchConfigParams struct {

	/*Configuration*/
	Configuration *models.DaemonConfigurationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch config params
func (o *PatchConfigParams) WithTimeout(timeout time.Duration) *PatchConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch config params
func (o *PatchConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch config params
func (o *PatchConfigParams) WithContext(ctx context.Context) *PatchConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch config params
func (o *PatchConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch config params
func (o *PatchConfigParams) WithHTTPClient(client *http.Client) *PatchConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch config params
func (o *PatchConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the patch config params
func (o *PatchConfigParams) WithConfiguration(configuration *models.DaemonConfigurationSpec) *PatchConfigParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the patch config params
func (o *PatchConfigParams) SetConfiguration(configuration *models.DaemonConfigurationSpec) {
	o.Configuration = configuration
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration == nil {
		o.Configuration = new(models.DaemonConfigurationSpec)
	}

	if err := r.SetBodyParam(o.Configuration); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
