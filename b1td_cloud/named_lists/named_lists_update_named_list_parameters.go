// Code generated by go-swagger; DO NOT EDIT.

package named_lists

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
	"github.com/go-openapi/swag"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// NewNamedListsUpdateNamedListParams creates a new NamedListsUpdateNamedListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamedListsUpdateNamedListParams() *NamedListsUpdateNamedListParams {
	return &NamedListsUpdateNamedListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamedListsUpdateNamedListParamsWithTimeout creates a new NamedListsUpdateNamedListParams object
// with the ability to set a timeout on a request.
func NewNamedListsUpdateNamedListParamsWithTimeout(timeout time.Duration) *NamedListsUpdateNamedListParams {
	return &NamedListsUpdateNamedListParams{
		timeout: timeout,
	}
}

// NewNamedListsUpdateNamedListParamsWithContext creates a new NamedListsUpdateNamedListParams object
// with the ability to set a context for a request.
func NewNamedListsUpdateNamedListParamsWithContext(ctx context.Context) *NamedListsUpdateNamedListParams {
	return &NamedListsUpdateNamedListParams{
		Context: ctx,
	}
}

// NewNamedListsUpdateNamedListParamsWithHTTPClient creates a new NamedListsUpdateNamedListParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamedListsUpdateNamedListParamsWithHTTPClient(client *http.Client) *NamedListsUpdateNamedListParams {
	return &NamedListsUpdateNamedListParams{
		HTTPClient: client,
	}
}

/*
NamedListsUpdateNamedListParams contains all the parameters to send to the API endpoint

	for the named lists update named list operation.

	Typically these are written to a http.Request.
*/
type NamedListsUpdateNamedListParams struct {

	/* Body.

	   The Named List object.
	*/
	Body *models.AtcfwNamedList

	/* ID.

	   The Named List object identifier.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the named lists update named list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamedListsUpdateNamedListParams) WithDefaults() *NamedListsUpdateNamedListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the named lists update named list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamedListsUpdateNamedListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) WithTimeout(timeout time.Duration) *NamedListsUpdateNamedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) WithContext(ctx context.Context) *NamedListsUpdateNamedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) WithHTTPClient(client *http.Client) *NamedListsUpdateNamedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) WithBody(body *models.AtcfwNamedList) *NamedListsUpdateNamedListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) SetBody(body *models.AtcfwNamedList) {
	o.Body = body
}

// WithID adds the id to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) WithID(id int32) *NamedListsUpdateNamedListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the named lists update named list params
func (o *NamedListsUpdateNamedListParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NamedListsUpdateNamedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
