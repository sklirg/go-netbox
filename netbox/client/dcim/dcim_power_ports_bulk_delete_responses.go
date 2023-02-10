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
)

// DcimPowerPortsBulkDeleteReader is a Reader for the DcimPowerPortsBulkDelete structure.
type DcimPowerPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsBulkDeleteNoContent creates a DcimPowerPortsBulkDeleteNoContent with default headers values
func NewDcimPowerPortsBulkDeleteNoContent() *DcimPowerPortsBulkDeleteNoContent {
	return &DcimPowerPortsBulkDeleteNoContent{}
}

/*
DcimPowerPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortsBulkDeleteNoContent dcim power ports bulk delete no content
*/
type DcimPowerPortsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power ports bulk delete no content response has a 2xx status code
func (o *DcimPowerPortsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports bulk delete no content response has a 3xx status code
func (o *DcimPowerPortsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports bulk delete no content response has a 4xx status code
func (o *DcimPowerPortsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports bulk delete no content response has a 5xx status code
func (o *DcimPowerPortsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports bulk delete no content response a status code equal to that given
func (o *DcimPowerPortsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim power ports bulk delete no content response
func (o *DcimPowerPortsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimPowerPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/][%d] dcimPowerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerPortsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/][%d] dcimPowerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerPortsBulkDeleteDefault creates a DcimPowerPortsBulkDeleteDefault with default headers values
func NewDcimPowerPortsBulkDeleteDefault(code int) *DcimPowerPortsBulkDeleteDefault {
	return &DcimPowerPortsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimPowerPortsBulkDeleteDefault dcim power ports bulk delete default
*/
type DcimPowerPortsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power ports bulk delete default response has a 2xx status code
func (o *DcimPowerPortsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports bulk delete default response has a 3xx status code
func (o *DcimPowerPortsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports bulk delete default response has a 4xx status code
func (o *DcimPowerPortsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports bulk delete default response has a 5xx status code
func (o *DcimPowerPortsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports bulk delete default response a status code equal to that given
func (o *DcimPowerPortsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power ports bulk delete default response
func (o *DcimPowerPortsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/][%d] dcim_power-ports_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/][%d] dcim_power-ports_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
