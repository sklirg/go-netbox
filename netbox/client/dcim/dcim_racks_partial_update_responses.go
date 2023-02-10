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

// DcimRacksPartialUpdateReader is a Reader for the DcimRacksPartialUpdate structure.
type DcimRacksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksPartialUpdateOK creates a DcimRacksPartialUpdateOK with default headers values
func NewDcimRacksPartialUpdateOK() *DcimRacksPartialUpdateOK {
	return &DcimRacksPartialUpdateOK{}
}

/*
DcimRacksPartialUpdateOK describes a response with status code 200, with default header values.

DcimRacksPartialUpdateOK dcim racks partial update o k
*/
type DcimRacksPartialUpdateOK struct {
	Payload *models.Rack
}

// IsSuccess returns true when this dcim racks partial update o k response has a 2xx status code
func (o *DcimRacksPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks partial update o k response has a 3xx status code
func (o *DcimRacksPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks partial update o k response has a 4xx status code
func (o *DcimRacksPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks partial update o k response has a 5xx status code
func (o *DcimRacksPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks partial update o k response a status code equal to that given
func (o *DcimRacksPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim racks partial update o k response
func (o *DcimRacksPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimRacksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/{id}/][%d] dcimRacksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRacksPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/racks/{id}/][%d] dcimRacksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRacksPartialUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksPartialUpdateDefault creates a DcimRacksPartialUpdateDefault with default headers values
func NewDcimRacksPartialUpdateDefault(code int) *DcimRacksPartialUpdateDefault {
	return &DcimRacksPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRacksPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRacksPartialUpdateDefault dcim racks partial update default
*/
type DcimRacksPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim racks partial update default response has a 2xx status code
func (o *DcimRacksPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks partial update default response has a 3xx status code
func (o *DcimRacksPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks partial update default response has a 4xx status code
func (o *DcimRacksPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks partial update default response has a 5xx status code
func (o *DcimRacksPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks partial update default response a status code equal to that given
func (o *DcimRacksPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim racks partial update default response
func (o *DcimRacksPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRacksPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/{id}/][%d] dcim_racks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/racks/{id}/][%d] dcim_racks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
