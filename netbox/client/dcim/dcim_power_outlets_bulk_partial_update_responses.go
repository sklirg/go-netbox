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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimPowerOutletsBulkPartialUpdateReader is a Reader for the DcimPowerOutletsBulkPartialUpdate structure.
type DcimPowerOutletsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsBulkPartialUpdateOK creates a DcimPowerOutletsBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletsBulkPartialUpdateOK() *DcimPowerOutletsBulkPartialUpdateOK {
	return &DcimPowerOutletsBulkPartialUpdateOK{}
}

/*
DcimPowerOutletsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkPartialUpdateOK dcim power outlets bulk partial update o k
*/
type DcimPowerOutletsBulkPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets bulk partial update o k response has a 2xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets bulk partial update o k response has a 3xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets bulk partial update o k response has a 4xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets bulk partial update o k response has a 5xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets bulk partial update o k response a status code equal to that given
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets bulk partial update o k response
func (o *DcimPowerOutletsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcimPowerOutletsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcimPowerOutletsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsBulkPartialUpdateDefault creates a DcimPowerOutletsBulkPartialUpdateDefault with default headers values
func NewDcimPowerOutletsBulkPartialUpdateDefault(code int) *DcimPowerOutletsBulkPartialUpdateDefault {
	return &DcimPowerOutletsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletsBulkPartialUpdateDefault dcim power outlets bulk partial update default
*/
type DcimPowerOutletsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets bulk partial update default response has a 2xx status code
func (o *DcimPowerOutletsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets bulk partial update default response has a 3xx status code
func (o *DcimPowerOutletsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets bulk partial update default response has a 4xx status code
func (o *DcimPowerOutletsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets bulk partial update default response has a 5xx status code
func (o *DcimPowerOutletsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets bulk partial update default response a status code equal to that given
func (o *DcimPowerOutletsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlets bulk partial update default response
func (o *DcimPowerOutletsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcim_power-outlets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcim_power-outlets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
