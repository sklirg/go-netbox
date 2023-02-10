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

// DcimRearPortsBulkUpdateReader is a Reader for the DcimRearPortsBulkUpdate structure.
type DcimRearPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsBulkUpdateOK creates a DcimRearPortsBulkUpdateOK with default headers values
func NewDcimRearPortsBulkUpdateOK() *DcimRearPortsBulkUpdateOK {
	return &DcimRearPortsBulkUpdateOK{}
}

/*
DcimRearPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsBulkUpdateOK dcim rear ports bulk update o k
*/
type DcimRearPortsBulkUpdateOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports bulk update o k response has a 2xx status code
func (o *DcimRearPortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports bulk update o k response has a 3xx status code
func (o *DcimRearPortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports bulk update o k response has a 4xx status code
func (o *DcimRearPortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports bulk update o k response has a 5xx status code
func (o *DcimRearPortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports bulk update o k response a status code equal to that given
func (o *DcimRearPortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rear ports bulk update o k response
func (o *DcimRearPortsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimRearPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/][%d] dcimRearPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/][%d] dcimRearPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsBulkUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsBulkUpdateDefault creates a DcimRearPortsBulkUpdateDefault with default headers values
func NewDcimRearPortsBulkUpdateDefault(code int) *DcimRearPortsBulkUpdateDefault {
	return &DcimRearPortsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimRearPortsBulkUpdateDefault dcim rear ports bulk update default
*/
type DcimRearPortsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear ports bulk update default response has a 2xx status code
func (o *DcimRearPortsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear ports bulk update default response has a 3xx status code
func (o *DcimRearPortsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear ports bulk update default response has a 4xx status code
func (o *DcimRearPortsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear ports bulk update default response has a 5xx status code
func (o *DcimRearPortsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear ports bulk update default response a status code equal to that given
func (o *DcimRearPortsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear ports bulk update default response
func (o *DcimRearPortsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/][%d] dcim_rear-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/][%d] dcim_rear-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
