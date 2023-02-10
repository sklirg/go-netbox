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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimManufacturersListReader is a Reader for the DcimManufacturersList structure.
type DcimManufacturersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersListOK creates a DcimManufacturersListOK with default headers values
func NewDcimManufacturersListOK() *DcimManufacturersListOK {
	return &DcimManufacturersListOK{}
}

/*
DcimManufacturersListOK describes a response with status code 200, with default header values.

DcimManufacturersListOK dcim manufacturers list o k
*/
type DcimManufacturersListOK struct {
	Payload *DcimManufacturersListOKBody
}

// IsSuccess returns true when this dcim manufacturers list o k response has a 2xx status code
func (o *DcimManufacturersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers list o k response has a 3xx status code
func (o *DcimManufacturersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers list o k response has a 4xx status code
func (o *DcimManufacturersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers list o k response has a 5xx status code
func (o *DcimManufacturersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers list o k response a status code equal to that given
func (o *DcimManufacturersListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim manufacturers list o k response
func (o *DcimManufacturersListOK) Code() int {
	return 200
}

func (o *DcimManufacturersListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/][%d] dcimManufacturersListOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersListOK) String() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/][%d] dcimManufacturersListOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersListOK) GetPayload() *DcimManufacturersListOKBody {
	return o.Payload
}

func (o *DcimManufacturersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimManufacturersListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersListDefault creates a DcimManufacturersListDefault with default headers values
func NewDcimManufacturersListDefault(code int) *DcimManufacturersListDefault {
	return &DcimManufacturersListDefault{
		_statusCode: code,
	}
}

/*
DcimManufacturersListDefault describes a response with status code -1, with default header values.

DcimManufacturersListDefault dcim manufacturers list default
*/
type DcimManufacturersListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim manufacturers list default response has a 2xx status code
func (o *DcimManufacturersListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim manufacturers list default response has a 3xx status code
func (o *DcimManufacturersListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim manufacturers list default response has a 4xx status code
func (o *DcimManufacturersListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim manufacturers list default response has a 5xx status code
func (o *DcimManufacturersListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim manufacturers list default response a status code equal to that given
func (o *DcimManufacturersListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim manufacturers list default response
func (o *DcimManufacturersListDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/][%d] dcim_manufacturers_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/][%d] dcim_manufacturers_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimManufacturersListOKBody dcim manufacturers list o k body
swagger:model DcimManufacturersListOKBody
*/
type DcimManufacturersListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Manufacturer `json:"results"`
}

// Validate validates this dcim manufacturers list o k body
func (o *DcimManufacturersListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimManufacturersListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimManufacturersListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimManufacturersListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimManufacturersListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimManufacturersListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimManufacturersListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimManufacturersListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimManufacturersListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimManufacturersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimManufacturersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim manufacturers list o k body based on the context it is used
func (o *DcimManufacturersListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimManufacturersListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimManufacturersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimManufacturersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimManufacturersListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimManufacturersListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimManufacturersListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
