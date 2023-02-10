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

// DcimConsoleServerPortsPartialUpdateReader is a Reader for the DcimConsoleServerPortsPartialUpdate structure.
type DcimConsoleServerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsPartialUpdateOK creates a DcimConsoleServerPortsPartialUpdateOK with default headers values
func NewDcimConsoleServerPortsPartialUpdateOK() *DcimConsoleServerPortsPartialUpdateOK {
	return &DcimConsoleServerPortsPartialUpdateOK{}
}

/*
DcimConsoleServerPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsPartialUpdateOK dcim console server ports partial update o k
*/
type DcimConsoleServerPortsPartialUpdateOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports partial update o k response has a 2xx status code
func (o *DcimConsoleServerPortsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports partial update o k response has a 3xx status code
func (o *DcimConsoleServerPortsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports partial update o k response has a 4xx status code
func (o *DcimConsoleServerPortsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports partial update o k response has a 5xx status code
func (o *DcimConsoleServerPortsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports partial update o k response a status code equal to that given
func (o *DcimConsoleServerPortsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports partial update o k response
func (o *DcimConsoleServerPortsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsPartialUpdateDefault creates a DcimConsoleServerPortsPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortsPartialUpdateDefault(code int) *DcimConsoleServerPortsPartialUpdateDefault {
	return &DcimConsoleServerPortsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsPartialUpdateDefault dcim console server ports partial update default
*/
type DcimConsoleServerPortsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console server ports partial update default response has a 2xx status code
func (o *DcimConsoleServerPortsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports partial update default response has a 3xx status code
func (o *DcimConsoleServerPortsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports partial update default response has a 4xx status code
func (o *DcimConsoleServerPortsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports partial update default response has a 5xx status code
func (o *DcimConsoleServerPortsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports partial update default response a status code equal to that given
func (o *DcimConsoleServerPortsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console server ports partial update default response
func (o *DcimConsoleServerPortsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
