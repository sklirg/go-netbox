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

// DcimPowerPortTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerPortTemplatesBulkPartialUpdate structure.
type DcimPowerPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateOK creates a DcimPowerPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesBulkPartialUpdateOK() *DcimPowerPortTemplatesBulkPartialUpdateOK {
	return &DcimPowerPortTemplatesBulkPartialUpdateOK{}
}

/*
DcimPowerPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesBulkPartialUpdateOK dcim power port templates bulk partial update o k
*/
type DcimPowerPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates bulk partial update o k response has a 2xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates bulk partial update o k response has a 3xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates bulk partial update o k response has a 4xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates bulk partial update o k response has a 5xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates bulk partial update o k response a status code equal to that given
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates bulk partial update o k response
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesBulkPartialUpdateDefault creates a DcimPowerPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimPowerPortTemplatesBulkPartialUpdateDefault(code int) *DcimPowerPortTemplatesBulkPartialUpdateDefault {
	return &DcimPowerPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesBulkPartialUpdateDefault dcim power port templates bulk partial update default
*/
type DcimPowerPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates bulk partial update default response has a 2xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates bulk partial update default response has a 3xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates bulk partial update default response has a 4xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates bulk partial update default response has a 5xx status code
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates bulk partial update default response a status code equal to that given
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power port templates bulk partial update default response
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcim_power-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcim_power-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
