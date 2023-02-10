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

// CircuitsCircuitTypesBulkUpdateReader is a Reader for the CircuitsCircuitTypesBulkUpdate structure.
type CircuitsCircuitTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesBulkUpdateOK creates a CircuitsCircuitTypesBulkUpdateOK with default headers values
func NewCircuitsCircuitTypesBulkUpdateOK() *CircuitsCircuitTypesBulkUpdateOK {
	return &CircuitsCircuitTypesBulkUpdateOK{}
}

/*
CircuitsCircuitTypesBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesBulkUpdateOK circuits circuit types bulk update o k
*/
type CircuitsCircuitTypesBulkUpdateOK struct {
	Payload *models.CircuitType
}

// IsSuccess returns true when this circuits circuit types bulk update o k response has a 2xx status code
func (o *CircuitsCircuitTypesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types bulk update o k response has a 3xx status code
func (o *CircuitsCircuitTypesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types bulk update o k response has a 4xx status code
func (o *CircuitsCircuitTypesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types bulk update o k response has a 5xx status code
func (o *CircuitsCircuitTypesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types bulk update o k response a status code equal to that given
func (o *CircuitsCircuitTypesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit types bulk update o k response
func (o *CircuitsCircuitTypesBulkUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/][%d] circuitsCircuitTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/][%d] circuitsCircuitTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesBulkUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesBulkUpdateDefault creates a CircuitsCircuitTypesBulkUpdateDefault with default headers values
func NewCircuitsCircuitTypesBulkUpdateDefault(code int) *CircuitsCircuitTypesBulkUpdateDefault {
	return &CircuitsCircuitTypesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTypesBulkUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTypesBulkUpdateDefault circuits circuit types bulk update default
*/
type CircuitsCircuitTypesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuit types bulk update default response has a 2xx status code
func (o *CircuitsCircuitTypesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit types bulk update default response has a 3xx status code
func (o *CircuitsCircuitTypesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit types bulk update default response has a 4xx status code
func (o *CircuitsCircuitTypesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit types bulk update default response has a 5xx status code
func (o *CircuitsCircuitTypesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit types bulk update default response a status code equal to that given
func (o *CircuitsCircuitTypesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuit types bulk update default response
func (o *CircuitsCircuitTypesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTypesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/][%d] circuits_circuit-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/][%d] circuits_circuit-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
