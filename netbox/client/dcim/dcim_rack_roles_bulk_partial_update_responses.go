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

// DcimRackRolesBulkPartialUpdateReader is a Reader for the DcimRackRolesBulkPartialUpdate structure.
type DcimRackRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesBulkPartialUpdateOK creates a DcimRackRolesBulkPartialUpdateOK with default headers values
func NewDcimRackRolesBulkPartialUpdateOK() *DcimRackRolesBulkPartialUpdateOK {
	return &DcimRackRolesBulkPartialUpdateOK{}
}

/*
DcimRackRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesBulkPartialUpdateOK dcim rack roles bulk partial update o k
*/
type DcimRackRolesBulkPartialUpdateOK struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles bulk partial update o k response has a 2xx status code
func (o *DcimRackRolesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles bulk partial update o k response has a 3xx status code
func (o *DcimRackRolesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles bulk partial update o k response has a 4xx status code
func (o *DcimRackRolesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles bulk partial update o k response has a 5xx status code
func (o *DcimRackRolesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles bulk partial update o k response a status code equal to that given
func (o *DcimRackRolesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack roles bulk partial update o k response
func (o *DcimRackRolesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimRackRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/][%d] dcimRackRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/][%d] dcimRackRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesBulkPartialUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesBulkPartialUpdateDefault creates a DcimRackRolesBulkPartialUpdateDefault with default headers values
func NewDcimRackRolesBulkPartialUpdateDefault(code int) *DcimRackRolesBulkPartialUpdateDefault {
	return &DcimRackRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRackRolesBulkPartialUpdateDefault dcim rack roles bulk partial update default
*/
type DcimRackRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles bulk partial update default response has a 2xx status code
func (o *DcimRackRolesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles bulk partial update default response has a 3xx status code
func (o *DcimRackRolesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles bulk partial update default response has a 4xx status code
func (o *DcimRackRolesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles bulk partial update default response has a 5xx status code
func (o *DcimRackRolesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles bulk partial update default response a status code equal to that given
func (o *DcimRackRolesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rack roles bulk partial update default response
func (o *DcimRackRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/][%d] dcim_rack-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/][%d] dcim_rack-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
