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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// ExtrasCustomFieldsCreateReader is a Reader for the ExtrasCustomFieldsCreate structure.
type ExtrasCustomFieldsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomFieldsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsCreateCreated creates a ExtrasCustomFieldsCreateCreated with default headers values
func NewExtrasCustomFieldsCreateCreated() *ExtrasCustomFieldsCreateCreated {
	return &ExtrasCustomFieldsCreateCreated{}
}

/*
ExtrasCustomFieldsCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomFieldsCreateCreated extras custom fields create created
*/
type ExtrasCustomFieldsCreateCreated struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this extras custom fields create created response has a 2xx status code
func (o *ExtrasCustomFieldsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields create created response has a 3xx status code
func (o *ExtrasCustomFieldsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields create created response has a 4xx status code
func (o *ExtrasCustomFieldsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields create created response has a 5xx status code
func (o *ExtrasCustomFieldsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields create created response a status code equal to that given
func (o *ExtrasCustomFieldsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the extras custom fields create created response
func (o *ExtrasCustomFieldsCreateCreated) Code() int {
	return 201
}

func (o *ExtrasCustomFieldsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extrasCustomFieldsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasCustomFieldsCreateCreated) String() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extrasCustomFieldsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasCustomFieldsCreateCreated) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsCreateDefault creates a ExtrasCustomFieldsCreateDefault with default headers values
func NewExtrasCustomFieldsCreateDefault(code int) *ExtrasCustomFieldsCreateDefault {
	return &ExtrasCustomFieldsCreateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsCreateDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldsCreateDefault extras custom fields create default
*/
type ExtrasCustomFieldsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras custom fields create default response has a 2xx status code
func (o *ExtrasCustomFieldsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields create default response has a 3xx status code
func (o *ExtrasCustomFieldsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields create default response has a 4xx status code
func (o *ExtrasCustomFieldsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields create default response has a 5xx status code
func (o *ExtrasCustomFieldsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields create default response a status code equal to that given
func (o *ExtrasCustomFieldsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras custom fields create default response
func (o *ExtrasCustomFieldsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extras_custom-fields_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsCreateDefault) String() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extras_custom-fields_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
