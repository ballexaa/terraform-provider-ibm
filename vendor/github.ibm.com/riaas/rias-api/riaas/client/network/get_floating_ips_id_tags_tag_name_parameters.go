// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetFloatingIpsIDTagsTagNameParams creates a new GetFloatingIpsIDTagsTagNameParams object
// with the default values initialized.
func NewGetFloatingIpsIDTagsTagNameParams() *GetFloatingIpsIDTagsTagNameParams {
	var ()
	return &GetFloatingIpsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFloatingIpsIDTagsTagNameParamsWithTimeout creates a new GetFloatingIpsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFloatingIpsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *GetFloatingIpsIDTagsTagNameParams {
	var ()
	return &GetFloatingIpsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewGetFloatingIpsIDTagsTagNameParamsWithContext creates a new GetFloatingIpsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFloatingIpsIDTagsTagNameParamsWithContext(ctx context.Context) *GetFloatingIpsIDTagsTagNameParams {
	var ()
	return &GetFloatingIpsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewGetFloatingIpsIDTagsTagNameParamsWithHTTPClient creates a new GetFloatingIpsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFloatingIpsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *GetFloatingIpsIDTagsTagNameParams {
	var ()
	return &GetFloatingIpsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*GetFloatingIpsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the get floating ips ID tags tag name operation typically these are written to a http.Request
*/
type GetFloatingIpsIDTagsTagNameParams struct {

	/*ID
	  The floating IP address

	*/
	ID string
	/*TagName
	  The name of the tag

	*/
	TagName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *GetFloatingIpsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithContext(ctx context.Context) *GetFloatingIpsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *GetFloatingIpsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithID(id string) *GetFloatingIpsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithTagName(tagName string) *GetFloatingIpsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) WithVersion(version string) *GetFloatingIpsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get floating ips ID tags tag name params
func (o *GetFloatingIpsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetFloatingIpsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}