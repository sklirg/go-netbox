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

// DcimPowerOutletsBulkUpdateReader is a Reader for the DcimPowerOutletsBulkUpdate structure.
type DcimPowerOutletsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsBulkUpdateOK creates a DcimPowerOutletsBulkUpdateOK with default headers values
func NewDcimPowerOutletsBulkUpdateOK() *DcimPowerOutletsBulkUpdateOK {
	return &DcimPowerOutletsBulkUpdateOK{}
}

/*
DcimPowerOutletsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkUpdateOK dcim power outlets bulk update o k
*/
type DcimPowerOutletsBulkUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets bulk update o k response has a 2xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets bulk update o k response has a 3xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets bulk update o k response has a 4xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets bulk update o k response has a 5xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets bulk update o k response a status code equal to that given
func (o *DcimPowerOutletsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets bulk update o k response
func (o *DcimPowerOutletsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcimPowerOutletsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcimPowerOutletsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsBulkUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsBulkUpdateDefault creates a DcimPowerOutletsBulkUpdateDefault with default headers values
func NewDcimPowerOutletsBulkUpdateDefault(code int) *DcimPowerOutletsBulkUpdateDefault {
	return &DcimPowerOutletsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletsBulkUpdateDefault dcim power outlets bulk update default
*/
type DcimPowerOutletsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets bulk update default response has a 2xx status code
func (o *DcimPowerOutletsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets bulk update default response has a 3xx status code
func (o *DcimPowerOutletsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets bulk update default response has a 4xx status code
func (o *DcimPowerOutletsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets bulk update default response has a 5xx status code
func (o *DcimPowerOutletsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets bulk update default response a status code equal to that given
func (o *DcimPowerOutletsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlets bulk update default response
func (o *DcimPowerOutletsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcim_power-outlets_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcim_power-outlets_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
