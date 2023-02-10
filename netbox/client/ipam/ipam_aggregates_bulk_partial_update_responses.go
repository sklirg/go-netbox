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

// IpamAggregatesBulkPartialUpdateReader is a Reader for the IpamAggregatesBulkPartialUpdate structure.
type IpamAggregatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAggregatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesBulkPartialUpdateOK creates a IpamAggregatesBulkPartialUpdateOK with default headers values
func NewIpamAggregatesBulkPartialUpdateOK() *IpamAggregatesBulkPartialUpdateOK {
	return &IpamAggregatesBulkPartialUpdateOK{}
}

/*
IpamAggregatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesBulkPartialUpdateOK ipam aggregates bulk partial update o k
*/
type IpamAggregatesBulkPartialUpdateOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates bulk partial update o k response has a 2xx status code
func (o *IpamAggregatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates bulk partial update o k response has a 3xx status code
func (o *IpamAggregatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates bulk partial update o k response has a 4xx status code
func (o *IpamAggregatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates bulk partial update o k response has a 5xx status code
func (o *IpamAggregatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates bulk partial update o k response a status code equal to that given
func (o *IpamAggregatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam aggregates bulk partial update o k response
func (o *IpamAggregatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamAggregatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/][%d] ipamAggregatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/][%d] ipamAggregatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesBulkPartialUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesBulkPartialUpdateDefault creates a IpamAggregatesBulkPartialUpdateDefault with default headers values
func NewIpamAggregatesBulkPartialUpdateDefault(code int) *IpamAggregatesBulkPartialUpdateDefault {
	return &IpamAggregatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamAggregatesBulkPartialUpdateDefault ipam aggregates bulk partial update default
*/
type IpamAggregatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam aggregates bulk partial update default response has a 2xx status code
func (o *IpamAggregatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates bulk partial update default response has a 3xx status code
func (o *IpamAggregatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates bulk partial update default response has a 4xx status code
func (o *IpamAggregatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates bulk partial update default response has a 5xx status code
func (o *IpamAggregatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates bulk partial update default response a status code equal to that given
func (o *IpamAggregatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam aggregates bulk partial update default response
func (o *IpamAggregatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamAggregatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/][%d] ipam_aggregates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/][%d] ipam_aggregates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
