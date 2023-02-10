// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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

	"github.com/sklirg/go-netbox/netbox/models"
)

// NewDcimSiteGroupsBulkPartialUpdateParams creates a new DcimSiteGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSiteGroupsBulkPartialUpdateParams() *DcimSiteGroupsBulkPartialUpdateParams {
	return &DcimSiteGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSiteGroupsBulkPartialUpdateParamsWithTimeout creates a new DcimSiteGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSiteGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimSiteGroupsBulkPartialUpdateParams {
	return &DcimSiteGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSiteGroupsBulkPartialUpdateParamsWithContext creates a new DcimSiteGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimSiteGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimSiteGroupsBulkPartialUpdateParams {
	return &DcimSiteGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimSiteGroupsBulkPartialUpdateParamsWithHTTPClient creates a new DcimSiteGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSiteGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimSiteGroupsBulkPartialUpdateParams {
	return &DcimSiteGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimSiteGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim site groups bulk partial update operation.

	Typically these are written to a http.Request.
*/
type DcimSiteGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableSiteGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim site groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsBulkPartialUpdateParams) WithDefaults() *DcimSiteGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim site groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimSiteGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimSiteGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimSiteGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) WithData(data *models.WritableSiteGroup) *DcimSiteGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim site groups bulk partial update params
func (o *DcimSiteGroupsBulkPartialUpdateParams) SetData(data *models.WritableSiteGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSiteGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
