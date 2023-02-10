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

// DcimPowerPortTemplatesPartialUpdateReader is a Reader for the DcimPowerPortTemplatesPartialUpdate structure.
type DcimPowerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesPartialUpdateOK creates a DcimPowerPortTemplatesPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesPartialUpdateOK() *DcimPowerPortTemplatesPartialUpdateOK {
	return &DcimPowerPortTemplatesPartialUpdateOK{}
}

/*
DcimPowerPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesPartialUpdateOK dcim power port templates partial update o k
*/
type DcimPowerPortTemplatesPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates partial update o k response has a 2xx status code
func (o *DcimPowerPortTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates partial update o k response has a 3xx status code
func (o *DcimPowerPortTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates partial update o k response has a 4xx status code
func (o *DcimPowerPortTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates partial update o k response has a 5xx status code
func (o *DcimPowerPortTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates partial update o k response a status code equal to that given
func (o *DcimPowerPortTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates partial update o k response
func (o *DcimPowerPortTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesPartialUpdateDefault creates a DcimPowerPortTemplatesPartialUpdateDefault with default headers values
func NewDcimPowerPortTemplatesPartialUpdateDefault(code int) *DcimPowerPortTemplatesPartialUpdateDefault {
	return &DcimPowerPortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesPartialUpdateDefault dcim power port templates partial update default
*/
type DcimPowerPortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates partial update default response has a 2xx status code
func (o *DcimPowerPortTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates partial update default response has a 3xx status code
func (o *DcimPowerPortTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates partial update default response has a 4xx status code
func (o *DcimPowerPortTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates partial update default response has a 5xx status code
func (o *DcimPowerPortTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates partial update default response a status code equal to that given
func (o *DcimPowerPortTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power port templates partial update default response
func (o *DcimPowerPortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
