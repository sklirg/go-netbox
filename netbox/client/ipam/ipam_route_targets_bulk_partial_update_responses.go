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

// IpamRouteTargetsBulkPartialUpdateReader is a Reader for the IpamRouteTargetsBulkPartialUpdate structure.
type IpamRouteTargetsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsBulkPartialUpdateOK creates a IpamRouteTargetsBulkPartialUpdateOK with default headers values
func NewIpamRouteTargetsBulkPartialUpdateOK() *IpamRouteTargetsBulkPartialUpdateOK {
	return &IpamRouteTargetsBulkPartialUpdateOK{}
}

/*
IpamRouteTargetsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamRouteTargetsBulkPartialUpdateOK ipam route targets bulk partial update o k
*/
type IpamRouteTargetsBulkPartialUpdateOK struct {
	Payload *models.RouteTarget
}

// IsSuccess returns true when this ipam route targets bulk partial update o k response has a 2xx status code
func (o *IpamRouteTargetsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets bulk partial update o k response has a 3xx status code
func (o *IpamRouteTargetsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets bulk partial update o k response has a 4xx status code
func (o *IpamRouteTargetsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets bulk partial update o k response has a 5xx status code
func (o *IpamRouteTargetsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets bulk partial update o k response a status code equal to that given
func (o *IpamRouteTargetsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam route targets bulk partial update o k response
func (o *IpamRouteTargetsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipamRouteTargetsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipamRouteTargetsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsBulkPartialUpdateDefault creates a IpamRouteTargetsBulkPartialUpdateDefault with default headers values
func NewIpamRouteTargetsBulkPartialUpdateDefault(code int) *IpamRouteTargetsBulkPartialUpdateDefault {
	return &IpamRouteTargetsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamRouteTargetsBulkPartialUpdateDefault ipam route targets bulk partial update default
*/
type IpamRouteTargetsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam route targets bulk partial update default response has a 2xx status code
func (o *IpamRouteTargetsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets bulk partial update default response has a 3xx status code
func (o *IpamRouteTargetsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets bulk partial update default response has a 4xx status code
func (o *IpamRouteTargetsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets bulk partial update default response has a 5xx status code
func (o *IpamRouteTargetsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets bulk partial update default response a status code equal to that given
func (o *IpamRouteTargetsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam route targets bulk partial update default response
func (o *IpamRouteTargetsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipam_route-targets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipam_route-targets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
