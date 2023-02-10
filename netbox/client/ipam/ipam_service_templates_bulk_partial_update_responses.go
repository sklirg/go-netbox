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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// IpamServiceTemplatesBulkPartialUpdateReader is a Reader for the IpamServiceTemplatesBulkPartialUpdate structure.
type IpamServiceTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServiceTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServiceTemplatesBulkPartialUpdateOK creates a IpamServiceTemplatesBulkPartialUpdateOK with default headers values
func NewIpamServiceTemplatesBulkPartialUpdateOK() *IpamServiceTemplatesBulkPartialUpdateOK {
	return &IpamServiceTemplatesBulkPartialUpdateOK{}
}

/*
IpamServiceTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamServiceTemplatesBulkPartialUpdateOK ipam service templates bulk partial update o k
*/
type IpamServiceTemplatesBulkPartialUpdateOK struct {
	Payload *models.ServiceTemplate
}

// IsSuccess returns true when this ipam service templates bulk partial update o k response has a 2xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates bulk partial update o k response has a 3xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates bulk partial update o k response has a 4xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates bulk partial update o k response has a 5xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates bulk partial update o k response a status code equal to that given
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam service templates bulk partial update o k response
func (o *IpamServiceTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipamServiceTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipamServiceTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) GetPayload() *models.ServiceTemplate {
	return o.Payload
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServiceTemplatesBulkPartialUpdateDefault creates a IpamServiceTemplatesBulkPartialUpdateDefault with default headers values
func NewIpamServiceTemplatesBulkPartialUpdateDefault(code int) *IpamServiceTemplatesBulkPartialUpdateDefault {
	return &IpamServiceTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServiceTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamServiceTemplatesBulkPartialUpdateDefault ipam service templates bulk partial update default
*/
type IpamServiceTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam service templates bulk partial update default response has a 2xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam service templates bulk partial update default response has a 3xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam service templates bulk partial update default response has a 4xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam service templates bulk partial update default response has a 5xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam service templates bulk partial update default response a status code equal to that given
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam service templates bulk partial update default response
func (o *IpamServiceTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServiceTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipam_service-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipam_service-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServiceTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
