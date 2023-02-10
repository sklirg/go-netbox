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

// DcimPowerPortTemplatesUpdateReader is a Reader for the DcimPowerPortTemplatesUpdate structure.
type DcimPowerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesUpdateOK creates a DcimPowerPortTemplatesUpdateOK with default headers values
func NewDcimPowerPortTemplatesUpdateOK() *DcimPowerPortTemplatesUpdateOK {
	return &DcimPowerPortTemplatesUpdateOK{}
}

/*
DcimPowerPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesUpdateOK dcim power port templates update o k
*/
type DcimPowerPortTemplatesUpdateOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates update o k response has a 2xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates update o k response has a 3xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates update o k response has a 4xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates update o k response has a 5xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates update o k response a status code equal to that given
func (o *DcimPowerPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates update o k response
func (o *DcimPowerPortTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesUpdateDefault creates a DcimPowerPortTemplatesUpdateDefault with default headers values
func NewDcimPowerPortTemplatesUpdateDefault(code int) *DcimPowerPortTemplatesUpdateDefault {
	return &DcimPowerPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesUpdateDefault dcim power port templates update default
*/
type DcimPowerPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates update default response has a 2xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates update default response has a 3xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates update default response has a 4xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates update default response has a 5xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates update default response a status code equal to that given
func (o *DcimPowerPortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power port templates update default response
func (o *DcimPowerPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
