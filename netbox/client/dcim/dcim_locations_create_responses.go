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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// DcimLocationsCreateReader is a Reader for the DcimLocationsCreate structure.
type DcimLocationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimLocationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsCreateCreated creates a DcimLocationsCreateCreated with default headers values
func NewDcimLocationsCreateCreated() *DcimLocationsCreateCreated {
	return &DcimLocationsCreateCreated{}
}

/*
DcimLocationsCreateCreated describes a response with status code 201, with default header values.

DcimLocationsCreateCreated dcim locations create created
*/
type DcimLocationsCreateCreated struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations create created response has a 2xx status code
func (o *DcimLocationsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations create created response has a 3xx status code
func (o *DcimLocationsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations create created response has a 4xx status code
func (o *DcimLocationsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations create created response has a 5xx status code
func (o *DcimLocationsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations create created response a status code equal to that given
func (o *DcimLocationsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim locations create created response
func (o *DcimLocationsCreateCreated) Code() int {
	return 201
}

func (o *DcimLocationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/locations/][%d] dcimLocationsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimLocationsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/locations/][%d] dcimLocationsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimLocationsCreateCreated) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsCreateDefault creates a DcimLocationsCreateDefault with default headers values
func NewDcimLocationsCreateDefault(code int) *DcimLocationsCreateDefault {
	return &DcimLocationsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsCreateDefault describes a response with status code -1, with default header values.

DcimLocationsCreateDefault dcim locations create default
*/
type DcimLocationsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim locations create default response has a 2xx status code
func (o *DcimLocationsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations create default response has a 3xx status code
func (o *DcimLocationsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations create default response has a 4xx status code
func (o *DcimLocationsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations create default response has a 5xx status code
func (o *DcimLocationsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations create default response a status code equal to that given
func (o *DcimLocationsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim locations create default response
func (o *DcimLocationsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimLocationsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/locations/][%d] dcim_locations_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/locations/][%d] dcim_locations_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
