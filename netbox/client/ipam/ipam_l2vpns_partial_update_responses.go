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

// IpamL2vpnsPartialUpdateReader is a Reader for the IpamL2vpnsPartialUpdate structure.
type IpamL2vpnsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnsPartialUpdateOK creates a IpamL2vpnsPartialUpdateOK with default headers values
func NewIpamL2vpnsPartialUpdateOK() *IpamL2vpnsPartialUpdateOK {
	return &IpamL2vpnsPartialUpdateOK{}
}

/*
IpamL2vpnsPartialUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnsPartialUpdateOK ipam l2vpns partial update o k
*/
type IpamL2vpnsPartialUpdateOK struct {
	Payload *models.L2VPN
}

// IsSuccess returns true when this ipam l2vpns partial update o k response has a 2xx status code
func (o *IpamL2vpnsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpns partial update o k response has a 3xx status code
func (o *IpamL2vpnsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns partial update o k response has a 4xx status code
func (o *IpamL2vpnsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpns partial update o k response has a 5xx status code
func (o *IpamL2vpnsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns partial update o k response a status code equal to that given
func (o *IpamL2vpnsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam l2vpns partial update o k response
func (o *IpamL2vpnsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamL2vpnsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpns/{id}/][%d] ipamL2vpnsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpns/{id}/][%d] ipamL2vpnsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsPartialUpdateOK) GetPayload() *models.L2VPN {
	return o.Payload
}

func (o *IpamL2vpnsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsPartialUpdateDefault creates a IpamL2vpnsPartialUpdateDefault with default headers values
func NewIpamL2vpnsPartialUpdateDefault(code int) *IpamL2vpnsPartialUpdateDefault {
	return &IpamL2vpnsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamL2vpnsPartialUpdateDefault ipam l2vpns partial update default
*/
type IpamL2vpnsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam l2vpns partial update default response has a 2xx status code
func (o *IpamL2vpnsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpns partial update default response has a 3xx status code
func (o *IpamL2vpnsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpns partial update default response has a 4xx status code
func (o *IpamL2vpnsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpns partial update default response has a 5xx status code
func (o *IpamL2vpnsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpns partial update default response a status code equal to that given
func (o *IpamL2vpnsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam l2vpns partial update default response
func (o *IpamL2vpnsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamL2vpnsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpns/{id}/][%d] ipam_l2vpns_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpns/{id}/][%d] ipam_l2vpns_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
