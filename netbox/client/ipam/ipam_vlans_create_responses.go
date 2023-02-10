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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamVlansCreateReader is a Reader for the IpamVlansCreate structure.
type IpamVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansCreateCreated creates a IpamVlansCreateCreated with default headers values
func NewIpamVlansCreateCreated() *IpamVlansCreateCreated {
	return &IpamVlansCreateCreated{}
}

/*
IpamVlansCreateCreated describes a response with status code 201, with default header values.

IpamVlansCreateCreated ipam vlans create created
*/
type IpamVlansCreateCreated struct {
	Payload *models.VLAN
}

// IsSuccess returns true when this ipam vlans create created response has a 2xx status code
func (o *IpamVlansCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlans create created response has a 3xx status code
func (o *IpamVlansCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlans create created response has a 4xx status code
func (o *IpamVlansCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlans create created response has a 5xx status code
func (o *IpamVlansCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlans create created response a status code equal to that given
func (o *IpamVlansCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam vlans create created response
func (o *IpamVlansCreateCreated) Code() int {
	return 201
}

func (o *IpamVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlansCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlansCreateCreated) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlansCreateDefault creates a IpamVlansCreateDefault with default headers values
func NewIpamVlansCreateDefault(code int) *IpamVlansCreateDefault {
	return &IpamVlansCreateDefault{
		_statusCode: code,
	}
}

/*
IpamVlansCreateDefault describes a response with status code -1, with default header values.

IpamVlansCreateDefault ipam vlans create default
*/
type IpamVlansCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vlans create default response has a 2xx status code
func (o *IpamVlansCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlans create default response has a 3xx status code
func (o *IpamVlansCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlans create default response has a 4xx status code
func (o *IpamVlansCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlans create default response has a 5xx status code
func (o *IpamVlansCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlans create default response a status code equal to that given
func (o *IpamVlansCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vlans create default response
func (o *IpamVlansCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipam_vlans_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlansCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipam_vlans_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlansCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
