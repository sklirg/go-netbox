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

// DcimInventoryItemTemplatesUpdateReader is a Reader for the DcimInventoryItemTemplatesUpdate structure.
type DcimInventoryItemTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesUpdateOK creates a DcimInventoryItemTemplatesUpdateOK with default headers values
func NewDcimInventoryItemTemplatesUpdateOK() *DcimInventoryItemTemplatesUpdateOK {
	return &DcimInventoryItemTemplatesUpdateOK{}
}

/*
DcimInventoryItemTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesUpdateOK dcim inventory item templates update o k
*/
type DcimInventoryItemTemplatesUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim inventory item templates update o k response
func (o *DcimInventoryItemTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimInventoryItemTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesUpdateDefault creates a DcimInventoryItemTemplatesUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesUpdateDefault(code int) *DcimInventoryItemTemplatesUpdateDefault {
	return &DcimInventoryItemTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemTemplatesUpdateDefault dcim inventory item templates update default
*/
type DcimInventoryItemTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item templates update default response has a 2xx status code
func (o *DcimInventoryItemTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates update default response has a 3xx status code
func (o *DcimInventoryItemTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates update default response has a 4xx status code
func (o *DcimInventoryItemTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates update default response has a 5xx status code
func (o *DcimInventoryItemTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates update default response a status code equal to that given
func (o *DcimInventoryItemTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim inventory item templates update default response
func (o *DcimInventoryItemTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
