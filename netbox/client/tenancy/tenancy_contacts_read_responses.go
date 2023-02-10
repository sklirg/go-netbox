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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// TenancyContactsReadReader is a Reader for the TenancyContactsRead structure.
type TenancyContactsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsReadOK creates a TenancyContactsReadOK with default headers values
func NewTenancyContactsReadOK() *TenancyContactsReadOK {
	return &TenancyContactsReadOK{}
}

/*
TenancyContactsReadOK describes a response with status code 200, with default header values.

TenancyContactsReadOK tenancy contacts read o k
*/
type TenancyContactsReadOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts read o k response has a 2xx status code
func (o *TenancyContactsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts read o k response has a 3xx status code
func (o *TenancyContactsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts read o k response has a 4xx status code
func (o *TenancyContactsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts read o k response has a 5xx status code
func (o *TenancyContactsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts read o k response a status code equal to that given
func (o *TenancyContactsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contacts read o k response
func (o *TenancyContactsReadOK) Code() int {
	return 200
}

func (o *TenancyContactsReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancyContactsReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsReadOK) String() string {
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancyContactsReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsReadOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsReadDefault creates a TenancyContactsReadDefault with default headers values
func NewTenancyContactsReadDefault(code int) *TenancyContactsReadDefault {
	return &TenancyContactsReadDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsReadDefault describes a response with status code -1, with default header values.

TenancyContactsReadDefault tenancy contacts read default
*/
type TenancyContactsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contacts read default response has a 2xx status code
func (o *TenancyContactsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts read default response has a 3xx status code
func (o *TenancyContactsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts read default response has a 4xx status code
func (o *TenancyContactsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts read default response has a 5xx status code
func (o *TenancyContactsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts read default response a status code equal to that given
func (o *TenancyContactsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contacts read default response
func (o *TenancyContactsReadDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsReadDefault) Error() string {
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancy_contacts_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsReadDefault) String() string {
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancy_contacts_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
