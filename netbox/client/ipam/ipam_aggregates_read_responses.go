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

// IpamAggregatesReadReader is a Reader for the IpamAggregatesRead structure.
type IpamAggregatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAggregatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesReadOK creates a IpamAggregatesReadOK with default headers values
func NewIpamAggregatesReadOK() *IpamAggregatesReadOK {
	return &IpamAggregatesReadOK{}
}

/*
IpamAggregatesReadOK describes a response with status code 200, with default header values.

IpamAggregatesReadOK ipam aggregates read o k
*/
type IpamAggregatesReadOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates read o k response has a 2xx status code
func (o *IpamAggregatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates read o k response has a 3xx status code
func (o *IpamAggregatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates read o k response has a 4xx status code
func (o *IpamAggregatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates read o k response has a 5xx status code
func (o *IpamAggregatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates read o k response a status code equal to that given
func (o *IpamAggregatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam aggregates read o k response
func (o *IpamAggregatesReadOK) Code() int {
	return 200
}

func (o *IpamAggregatesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipamAggregatesReadOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipamAggregatesReadOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesReadOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesReadDefault creates a IpamAggregatesReadDefault with default headers values
func NewIpamAggregatesReadDefault(code int) *IpamAggregatesReadDefault {
	return &IpamAggregatesReadDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesReadDefault describes a response with status code -1, with default header values.

IpamAggregatesReadDefault ipam aggregates read default
*/
type IpamAggregatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam aggregates read default response has a 2xx status code
func (o *IpamAggregatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates read default response has a 3xx status code
func (o *IpamAggregatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates read default response has a 4xx status code
func (o *IpamAggregatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates read default response has a 5xx status code
func (o *IpamAggregatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates read default response a status code equal to that given
func (o *IpamAggregatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam aggregates read default response
func (o *IpamAggregatesReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamAggregatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipam_aggregates_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipam_aggregates_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
