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

// DcimFrontPortTemplatesPartialUpdateReader is a Reader for the DcimFrontPortTemplatesPartialUpdate structure.
type DcimFrontPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesPartialUpdateOK creates a DcimFrontPortTemplatesPartialUpdateOK with default headers values
func NewDcimFrontPortTemplatesPartialUpdateOK() *DcimFrontPortTemplatesPartialUpdateOK {
	return &DcimFrontPortTemplatesPartialUpdateOK{}
}

/*
DcimFrontPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesPartialUpdateOK dcim front port templates partial update o k
*/
type DcimFrontPortTemplatesPartialUpdateOK struct {
	Payload *models.FrontPortTemplate
}

// IsSuccess returns true when this dcim front port templates partial update o k response has a 2xx status code
func (o *DcimFrontPortTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates partial update o k response has a 3xx status code
func (o *DcimFrontPortTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates partial update o k response has a 4xx status code
func (o *DcimFrontPortTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates partial update o k response has a 5xx status code
func (o *DcimFrontPortTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates partial update o k response a status code equal to that given
func (o *DcimFrontPortTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front port templates partial update o k response
func (o *DcimFrontPortTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesPartialUpdateDefault creates a DcimFrontPortTemplatesPartialUpdateDefault with default headers values
func NewDcimFrontPortTemplatesPartialUpdateDefault(code int) *DcimFrontPortTemplatesPartialUpdateDefault {
	return &DcimFrontPortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortTemplatesPartialUpdateDefault dcim front port templates partial update default
*/
type DcimFrontPortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim front port templates partial update default response has a 2xx status code
func (o *DcimFrontPortTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front port templates partial update default response has a 3xx status code
func (o *DcimFrontPortTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front port templates partial update default response has a 4xx status code
func (o *DcimFrontPortTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front port templates partial update default response has a 5xx status code
func (o *DcimFrontPortTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front port templates partial update default response a status code equal to that given
func (o *DcimFrontPortTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim front port templates partial update default response
func (o *DcimFrontPortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
