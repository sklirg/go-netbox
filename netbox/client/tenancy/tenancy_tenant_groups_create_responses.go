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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyTenantGroupsCreateReader is a Reader for the TenancyTenantGroupsCreate structure.
type TenancyTenantGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsCreateCreated creates a TenancyTenantGroupsCreateCreated with default headers values
func NewTenancyTenantGroupsCreateCreated() *TenancyTenantGroupsCreateCreated {
	return &TenancyTenantGroupsCreateCreated{}
}

/*
TenancyTenantGroupsCreateCreated describes a response with status code 201, with default header values.

TenancyTenantGroupsCreateCreated tenancy tenant groups create created
*/
type TenancyTenantGroupsCreateCreated struct {
	Payload *models.TenantGroup
}

// IsSuccess returns true when this tenancy tenant groups create created response has a 2xx status code
func (o *TenancyTenantGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenant groups create created response has a 3xx status code
func (o *TenancyTenantGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups create created response has a 4xx status code
func (o *TenancyTenantGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenant groups create created response has a 5xx status code
func (o *TenancyTenantGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups create created response a status code equal to that given
func (o *TenancyTenantGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy tenant groups create created response
func (o *TenancyTenantGroupsCreateCreated) Code() int {
	return 201
}

func (o *TenancyTenantGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancyTenantGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancyTenantGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantGroupsCreateCreated) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantGroupsCreateDefault creates a TenancyTenantGroupsCreateDefault with default headers values
func NewTenancyTenantGroupsCreateDefault(code int) *TenancyTenantGroupsCreateDefault {
	return &TenancyTenantGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantGroupsCreateDefault describes a response with status code -1, with default header values.

TenancyTenantGroupsCreateDefault tenancy tenant groups create default
*/
type TenancyTenantGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy tenant groups create default response has a 2xx status code
func (o *TenancyTenantGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenant groups create default response has a 3xx status code
func (o *TenancyTenantGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenant groups create default response has a 4xx status code
func (o *TenancyTenantGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenant groups create default response has a 5xx status code
func (o *TenancyTenantGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenant groups create default response a status code equal to that given
func (o *TenancyTenantGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy tenant groups create default response
func (o *TenancyTenantGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancy_tenant-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancy_tenant-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
