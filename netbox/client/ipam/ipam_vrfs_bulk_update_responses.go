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

// IpamVrfsBulkUpdateReader is a Reader for the IpamVrfsBulkUpdate structure.
type IpamVrfsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsBulkUpdateOK creates a IpamVrfsBulkUpdateOK with default headers values
func NewIpamVrfsBulkUpdateOK() *IpamVrfsBulkUpdateOK {
	return &IpamVrfsBulkUpdateOK{}
}

/*
IpamVrfsBulkUpdateOK describes a response with status code 200, with default header values.

IpamVrfsBulkUpdateOK ipam vrfs bulk update o k
*/
type IpamVrfsBulkUpdateOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs bulk update o k response has a 2xx status code
func (o *IpamVrfsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs bulk update o k response has a 3xx status code
func (o *IpamVrfsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs bulk update o k response has a 4xx status code
func (o *IpamVrfsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs bulk update o k response has a 5xx status code
func (o *IpamVrfsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs bulk update o k response a status code equal to that given
func (o *IpamVrfsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vrfs bulk update o k response
func (o *IpamVrfsBulkUpdateOK) Code() int {
	return 200
}

func (o *IpamVrfsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/][%d] ipamVrfsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/][%d] ipamVrfsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsBulkUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsBulkUpdateDefault creates a IpamVrfsBulkUpdateDefault with default headers values
func NewIpamVrfsBulkUpdateDefault(code int) *IpamVrfsBulkUpdateDefault {
	return &IpamVrfsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsBulkUpdateDefault describes a response with status code -1, with default header values.

IpamVrfsBulkUpdateDefault ipam vrfs bulk update default
*/
type IpamVrfsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vrfs bulk update default response has a 2xx status code
func (o *IpamVrfsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs bulk update default response has a 3xx status code
func (o *IpamVrfsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs bulk update default response has a 4xx status code
func (o *IpamVrfsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs bulk update default response has a 5xx status code
func (o *IpamVrfsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs bulk update default response a status code equal to that given
func (o *IpamVrfsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vrfs bulk update default response
func (o *IpamVrfsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/][%d] ipam_vrfs_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/][%d] ipam_vrfs_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
