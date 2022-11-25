// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListProjectAzureAvailabilityZonesParams creates a new ListProjectAzureAvailabilityZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectAzureAvailabilityZonesParams() *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithTimeout creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a timeout on a request.
func NewListProjectAzureAvailabilityZonesParamsWithTimeout(timeout time.Duration) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		timeout: timeout,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithContext creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a context for a request.
func NewListProjectAzureAvailabilityZonesParamsWithContext(ctx context.Context) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		Context: ctx,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithHTTPClient creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectAzureAvailabilityZonesParamsWithHTTPClient(client *http.Client) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		HTTPClient: client,
	}
}

/*
ListProjectAzureAvailabilityZonesParams contains all the parameters to send to the API endpoint

	for the list project azure availability zones operation.

	Typically these are written to a http.Request.
*/
type ListProjectAzureAvailabilityZonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project azure availability zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAzureAvailabilityZonesParams) WithDefaults() *ListProjectAzureAvailabilityZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project azure availability zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAzureAvailabilityZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithTimeout(timeout time.Duration) *ListProjectAzureAvailabilityZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithContext(ctx context.Context) *ListProjectAzureAvailabilityZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithHTTPClient(client *http.Client) *ListProjectAzureAvailabilityZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectAzureAvailabilityZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
