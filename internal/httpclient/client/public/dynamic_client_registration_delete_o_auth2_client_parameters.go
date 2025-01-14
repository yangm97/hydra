// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewDynamicClientRegistrationDeleteOAuth2ClientParams creates a new DynamicClientRegistrationDeleteOAuth2ClientParams object
// with the default values initialized.
func NewDynamicClientRegistrationDeleteOAuth2ClientParams() *DynamicClientRegistrationDeleteOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationDeleteOAuth2ClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithTimeout creates a new DynamicClientRegistrationDeleteOAuth2ClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithTimeout(timeout time.Duration) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationDeleteOAuth2ClientParams{

		timeout: timeout,
	}
}

// NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithContext creates a new DynamicClientRegistrationDeleteOAuth2ClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithContext(ctx context.Context) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationDeleteOAuth2ClientParams{

		Context: ctx,
	}
}

// NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithHTTPClient creates a new DynamicClientRegistrationDeleteOAuth2ClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDynamicClientRegistrationDeleteOAuth2ClientParamsWithHTTPClient(client *http.Client) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationDeleteOAuth2ClientParams{
		HTTPClient: client,
	}
}

/*DynamicClientRegistrationDeleteOAuth2ClientParams contains all the parameters to send to the API endpoint
for the dynamic client registration delete o auth2 client operation typically these are written to a http.Request
*/
type DynamicClientRegistrationDeleteOAuth2ClientParams struct {

	/*ID
	  The id of the OAuth 2.0 Client.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) WithTimeout(timeout time.Duration) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) WithContext(ctx context.Context) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) WithHTTPClient(client *http.Client) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) WithID(id string) *DynamicClientRegistrationDeleteOAuth2ClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dynamic client registration delete o auth2 client params
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DynamicClientRegistrationDeleteOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
