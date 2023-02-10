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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// CircuitsCircuitTypesReadReader is a Reader for the CircuitsCircuitTypesRead structure.
type CircuitsCircuitTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesReadOK creates a CircuitsCircuitTypesReadOK with default headers values
func NewCircuitsCircuitTypesReadOK() *CircuitsCircuitTypesReadOK {
	return &CircuitsCircuitTypesReadOK{}
}

/*
CircuitsCircuitTypesReadOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesReadOK circuits circuit types read o k
*/
type CircuitsCircuitTypesReadOK struct {
	Payload *models.CircuitType
}

// IsSuccess returns true when this circuits circuit types read o k response has a 2xx status code
func (o *CircuitsCircuitTypesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types read o k response has a 3xx status code
func (o *CircuitsCircuitTypesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types read o k response has a 4xx status code
func (o *CircuitsCircuitTypesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types read o k response has a 5xx status code
func (o *CircuitsCircuitTypesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types read o k response a status code equal to that given
func (o *CircuitsCircuitTypesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit types read o k response
func (o *CircuitsCircuitTypesReadOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesReadOK) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesReadOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesReadDefault creates a CircuitsCircuitTypesReadDefault with default headers values
func NewCircuitsCircuitTypesReadDefault(code int) *CircuitsCircuitTypesReadDefault {
	return &CircuitsCircuitTypesReadDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTypesReadDefault describes a response with status code -1, with default header values.

CircuitsCircuitTypesReadDefault circuits circuit types read default
*/
type CircuitsCircuitTypesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuit types read default response has a 2xx status code
func (o *CircuitsCircuitTypesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit types read default response has a 3xx status code
func (o *CircuitsCircuitTypesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit types read default response has a 4xx status code
func (o *CircuitsCircuitTypesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit types read default response has a 5xx status code
func (o *CircuitsCircuitTypesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit types read default response a status code equal to that given
func (o *CircuitsCircuitTypesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuit types read default response
func (o *CircuitsCircuitTypesReadDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTypesReadDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/{id}/][%d] circuits_circuit-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesReadDefault) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/{id}/][%d] circuits_circuit-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
