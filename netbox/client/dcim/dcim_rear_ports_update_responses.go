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

// DcimRearPortsUpdateReader is a Reader for the DcimRearPortsUpdate structure.
type DcimRearPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsUpdateOK creates a DcimRearPortsUpdateOK with default headers values
func NewDcimRearPortsUpdateOK() *DcimRearPortsUpdateOK {
	return &DcimRearPortsUpdateOK{}
}

/*
DcimRearPortsUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsUpdateOK dcim rear ports update o k
*/
type DcimRearPortsUpdateOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports update o k response has a 2xx status code
func (o *DcimRearPortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports update o k response has a 3xx status code
func (o *DcimRearPortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports update o k response has a 4xx status code
func (o *DcimRearPortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports update o k response has a 5xx status code
func (o *DcimRearPortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports update o k response a status code equal to that given
func (o *DcimRearPortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rear ports update o k response
func (o *DcimRearPortsUpdateOK) Code() int {
	return 200
}

func (o *DcimRearPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcimRearPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcimRearPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsUpdateDefault creates a DcimRearPortsUpdateDefault with default headers values
func NewDcimRearPortsUpdateDefault(code int) *DcimRearPortsUpdateDefault {
	return &DcimRearPortsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortsUpdateDefault describes a response with status code -1, with default header values.

DcimRearPortsUpdateDefault dcim rear ports update default
*/
type DcimRearPortsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear ports update default response has a 2xx status code
func (o *DcimRearPortsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear ports update default response has a 3xx status code
func (o *DcimRearPortsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear ports update default response has a 4xx status code
func (o *DcimRearPortsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear ports update default response has a 5xx status code
func (o *DcimRearPortsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear ports update default response a status code equal to that given
func (o *DcimRearPortsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear ports update default response
func (o *DcimRearPortsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcim_rear-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcim_rear-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
