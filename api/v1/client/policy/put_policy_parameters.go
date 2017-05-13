package policy

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

// NewPutPolicyParams creates a new PutPolicyParams object
// with the default values initialized.
func NewPutPolicyParams() *PutPolicyParams {
	var ()
	return &PutPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPolicyParamsWithTimeout creates a new PutPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPolicyParamsWithTimeout(timeout time.Duration) *PutPolicyParams {
	var ()
	return &PutPolicyParams{

		timeout: timeout,
	}
}

// NewPutPolicyParamsWithContext creates a new PutPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPolicyParamsWithContext(ctx context.Context) *PutPolicyParams {
	var ()
	return &PutPolicyParams{

		Context: ctx,
	}
}

// NewPutPolicyParamsWithHTTPClient creates a new PutPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPolicyParamsWithHTTPClient(client *http.Client) *PutPolicyParams {
	var ()
	return &PutPolicyParams{
		HTTPClient: client,
	}
}

/*PutPolicyParams contains all the parameters to send to the API endpoint
for the put policy operation typically these are written to a http.Request
*/
type PutPolicyParams struct {

	/*Policy
	  Policy rules

	*/
	Policy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put policy params
func (o *PutPolicyParams) WithTimeout(timeout time.Duration) *PutPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put policy params
func (o *PutPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put policy params
func (o *PutPolicyParams) WithContext(ctx context.Context) *PutPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put policy params
func (o *PutPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put policy params
func (o *PutPolicyParams) WithHTTPClient(client *http.Client) *PutPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put policy params
func (o *PutPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the put policy params
func (o *PutPolicyParams) WithPolicy(policy *string) *PutPolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the put policy params
func (o *PutPolicyParams) SetPolicy(policy *string) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PutPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Policy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
