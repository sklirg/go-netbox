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

// TenancyTenantsCreateReader is a Reader for the TenancyTenantsCreate structure.
type TenancyTenantsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsCreateCreated creates a TenancyTenantsCreateCreated with default headers values
func NewTenancyTenantsCreateCreated() *TenancyTenantsCreateCreated {
	return &TenancyTenantsCreateCreated{}
}

/*
TenancyTenantsCreateCreated describes a response with status code 201, with default header values.

TenancyTenantsCreateCreated tenancy tenants create created
*/
type TenancyTenantsCreateCreated struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this tenancy tenants create created response has a 2xx status code
func (o *TenancyTenantsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants create created response has a 3xx status code
func (o *TenancyTenantsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants create created response has a 4xx status code
func (o *TenancyTenantsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants create created response has a 5xx status code
func (o *TenancyTenantsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants create created response a status code equal to that given
func (o *TenancyTenantsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy tenants create created response
func (o *TenancyTenantsCreateCreated) Code() int {
	return 201
}

func (o *TenancyTenantsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancyTenantsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantsCreateCreated) String() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancyTenantsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantsCreateCreated) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsCreateDefault creates a TenancyTenantsCreateDefault with default headers values
func NewTenancyTenantsCreateDefault(code int) *TenancyTenantsCreateDefault {
	return &TenancyTenantsCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantsCreateDefault describes a response with status code -1, with default header values.

TenancyTenantsCreateDefault tenancy tenants create default
*/
type TenancyTenantsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy tenants create default response has a 2xx status code
func (o *TenancyTenantsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenants create default response has a 3xx status code
func (o *TenancyTenantsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenants create default response has a 4xx status code
func (o *TenancyTenantsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenants create default response has a 5xx status code
func (o *TenancyTenantsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenants create default response a status code equal to that given
func (o *TenancyTenantsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy tenants create default response
func (o *TenancyTenantsCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancy_tenants_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsCreateDefault) String() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancy_tenants_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
