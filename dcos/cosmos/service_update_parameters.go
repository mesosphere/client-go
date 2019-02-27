// Code generated by go-swagger; DO NOT EDIT.

package cosmos

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

	types "github.com/dcos/client-go/dcos/cosmos/types"
)

// NewServiceUpdateParams creates a new ServiceUpdateParams object
// with the default values initialized.
func NewServiceUpdateParams() *ServiceUpdateParams {
	var ()
	return &ServiceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceUpdateParamsWithTimeout creates a new ServiceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceUpdateParamsWithTimeout(timeout time.Duration) *ServiceUpdateParams {
	var ()
	return &ServiceUpdateParams{

		timeout: timeout,
	}
}

// NewServiceUpdateParamsWithContext creates a new ServiceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceUpdateParamsWithContext(ctx context.Context) *ServiceUpdateParams {
	var ()
	return &ServiceUpdateParams{

		Context: ctx,
	}
}

// NewServiceUpdateParamsWithHTTPClient creates a new ServiceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceUpdateParamsWithHTTPClient(client *http.Client) *ServiceUpdateParams {
	var ()
	return &ServiceUpdateParams{
		HTTPClient: client,
	}
}

/*ServiceUpdateParams contains all the parameters to send to the API endpoint
for the service update operation typically these are written to a http.Request
*/
type ServiceUpdateParams struct {

	/*Body*/
	Body *types.ServiceUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service update params
func (o *ServiceUpdateParams) WithTimeout(timeout time.Duration) *ServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service update params
func (o *ServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service update params
func (o *ServiceUpdateParams) WithContext(ctx context.Context) *ServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service update params
func (o *ServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service update params
func (o *ServiceUpdateParams) WithHTTPClient(client *http.Client) *ServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service update params
func (o *ServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service update params
func (o *ServiceUpdateParams) WithBody(body *types.ServiceUpdateRequest) *ServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service update params
func (o *ServiceUpdateParams) SetBody(body *types.ServiceUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
