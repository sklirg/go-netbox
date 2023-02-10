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

package ipam

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

// IpamRirsListReader is a Reader for the IpamRirsList structure.
type IpamRirsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsListOK creates a IpamRirsListOK with default headers values
func NewIpamRirsListOK() *IpamRirsListOK {
	return &IpamRirsListOK{}
}

/*
IpamRirsListOK describes a response with status code 200, with default header values.

IpamRirsListOK ipam rirs list o k
*/
type IpamRirsListOK struct {
	Payload *IpamRirsListOKBody
}

// IsSuccess returns true when this ipam rirs list o k response has a 2xx status code
func (o *IpamRirsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs list o k response has a 3xx status code
func (o *IpamRirsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs list o k response has a 4xx status code
func (o *IpamRirsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs list o k response has a 5xx status code
func (o *IpamRirsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs list o k response a status code equal to that given
func (o *IpamRirsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam rirs list o k response
func (o *IpamRirsListOK) Code() int {
	return 200
}

func (o *IpamRirsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/rirs/][%d] ipamRirsListOK  %+v", 200, o.Payload)
}

func (o *IpamRirsListOK) String() string {
	return fmt.Sprintf("[GET /ipam/rirs/][%d] ipamRirsListOK  %+v", 200, o.Payload)
}

func (o *IpamRirsListOK) GetPayload() *IpamRirsListOKBody {
	return o.Payload
}

func (o *IpamRirsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamRirsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRirsListDefault creates a IpamRirsListDefault with default headers values
func NewIpamRirsListDefault(code int) *IpamRirsListDefault {
	return &IpamRirsListDefault{
		_statusCode: code,
	}
}

/*
IpamRirsListDefault describes a response with status code -1, with default header values.

IpamRirsListDefault ipam rirs list default
*/
type IpamRirsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam rirs list default response has a 2xx status code
func (o *IpamRirsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs list default response has a 3xx status code
func (o *IpamRirsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs list default response has a 4xx status code
func (o *IpamRirsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs list default response has a 5xx status code
func (o *IpamRirsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs list default response a status code equal to that given
func (o *IpamRirsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam rirs list default response
func (o *IpamRirsListDefault) Code() int {
	return o._statusCode
}

func (o *IpamRirsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/rirs/][%d] ipam_rirs_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsListDefault) String() string {
	return fmt.Sprintf("[GET /ipam/rirs/][%d] ipam_rirs_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IpamRirsListOKBody ipam rirs list o k body
swagger:model IpamRirsListOKBody
*/
type IpamRirsListOKBody struct {

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
	Results []*models.RIR `json:"results"`
}

// Validate validates this ipam rirs list o k body
func (o *IpamRirsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *IpamRirsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamRirsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamRirsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamRirsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamRirsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamRirsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamRirsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamRirsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamRirsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamRirsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam rirs list o k body based on the context it is used
func (o *IpamRirsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamRirsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamRirsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamRirsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamRirsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamRirsListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamRirsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
