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

// DcimFrontPortTemplatesBulkPartialUpdateReader is a Reader for the DcimFrontPortTemplatesBulkPartialUpdate structure.
type DcimFrontPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesBulkPartialUpdateOK creates a DcimFrontPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimFrontPortTemplatesBulkPartialUpdateOK() *DcimFrontPortTemplatesBulkPartialUpdateOK {
	return &DcimFrontPortTemplatesBulkPartialUpdateOK{}
}

/*
DcimFrontPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesBulkPartialUpdateOK dcim front port templates bulk partial update o k
*/
type DcimFrontPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.FrontPortTemplate
}

// IsSuccess returns true when this dcim front port templates bulk partial update o k response has a 2xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates bulk partial update o k response has a 3xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates bulk partial update o k response has a 4xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates bulk partial update o k response has a 5xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates bulk partial update o k response a status code equal to that given
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front port templates bulk partial update o k response
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesBulkPartialUpdateDefault creates a DcimFrontPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimFrontPortTemplatesBulkPartialUpdateDefault(code int) *DcimFrontPortTemplatesBulkPartialUpdateDefault {
	return &DcimFrontPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortTemplatesBulkPartialUpdateDefault dcim front port templates bulk partial update default
*/
type DcimFrontPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim front port templates bulk partial update default response has a 2xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front port templates bulk partial update default response has a 3xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front port templates bulk partial update default response has a 4xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front port templates bulk partial update default response has a 5xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front port templates bulk partial update default response a status code equal to that given
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim front port templates bulk partial update default response
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcim_front-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcim_front-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
