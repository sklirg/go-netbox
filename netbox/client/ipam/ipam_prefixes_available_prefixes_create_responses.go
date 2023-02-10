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

// IpamPrefixesAvailablePrefixesCreateReader is a Reader for the IpamPrefixesAvailablePrefixesCreate structure.
type IpamPrefixesAvailablePrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailablePrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailablePrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesAvailablePrefixesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesAvailablePrefixesCreateCreated creates a IpamPrefixesAvailablePrefixesCreateCreated with default headers values
func NewIpamPrefixesAvailablePrefixesCreateCreated() *IpamPrefixesAvailablePrefixesCreateCreated {
	return &IpamPrefixesAvailablePrefixesCreateCreated{}
}

/*
IpamPrefixesAvailablePrefixesCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesAvailablePrefixesCreateCreated ipam prefixes available prefixes create created
*/
type IpamPrefixesAvailablePrefixesCreateCreated struct {
	Payload []*models.Prefix
}

// IsSuccess returns true when this ipam prefixes available prefixes create created response has a 2xx status code
func (o *IpamPrefixesAvailablePrefixesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes available prefixes create created response has a 3xx status code
func (o *IpamPrefixesAvailablePrefixesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes available prefixes create created response has a 4xx status code
func (o *IpamPrefixesAvailablePrefixesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes available prefixes create created response has a 5xx status code
func (o *IpamPrefixesAvailablePrefixesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes available prefixes create created response a status code equal to that given
func (o *IpamPrefixesAvailablePrefixesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam prefixes available prefixes create created response
func (o *IpamPrefixesAvailablePrefixesCreateCreated) Code() int {
	return 201
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) GetPayload() []*models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesAvailablePrefixesCreateDefault creates a IpamPrefixesAvailablePrefixesCreateDefault with default headers values
func NewIpamPrefixesAvailablePrefixesCreateDefault(code int) *IpamPrefixesAvailablePrefixesCreateDefault {
	return &IpamPrefixesAvailablePrefixesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesAvailablePrefixesCreateDefault describes a response with status code -1, with default header values.

IpamPrefixesAvailablePrefixesCreateDefault ipam prefixes available prefixes create default
*/
type IpamPrefixesAvailablePrefixesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes available prefixes create default response has a 2xx status code
func (o *IpamPrefixesAvailablePrefixesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes available prefixes create default response has a 3xx status code
func (o *IpamPrefixesAvailablePrefixesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes available prefixes create default response has a 4xx status code
func (o *IpamPrefixesAvailablePrefixesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes available prefixes create default response has a 5xx status code
func (o *IpamPrefixesAvailablePrefixesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes available prefixes create default response a status code equal to that given
func (o *IpamPrefixesAvailablePrefixesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam prefixes available prefixes create default response
func (o *IpamPrefixesAvailablePrefixesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesAvailablePrefixesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipam_prefixes_available-prefixes_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailablePrefixesCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipam_prefixes_available-prefixes_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailablePrefixesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
