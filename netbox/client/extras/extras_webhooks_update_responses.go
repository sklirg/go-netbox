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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// ExtrasWebhooksUpdateReader is a Reader for the ExtrasWebhooksUpdate structure.
type ExtrasWebhooksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksUpdateOK creates a ExtrasWebhooksUpdateOK with default headers values
func NewExtrasWebhooksUpdateOK() *ExtrasWebhooksUpdateOK {
	return &ExtrasWebhooksUpdateOK{}
}

/*
ExtrasWebhooksUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksUpdateOK extras webhooks update o k
*/
type ExtrasWebhooksUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks update o k response has a 2xx status code
func (o *ExtrasWebhooksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks update o k response has a 3xx status code
func (o *ExtrasWebhooksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks update o k response has a 4xx status code
func (o *ExtrasWebhooksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks update o k response has a 5xx status code
func (o *ExtrasWebhooksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks update o k response a status code equal to that given
func (o *ExtrasWebhooksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks update o k response
func (o *ExtrasWebhooksUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extrasWebhooksUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extrasWebhooksUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksUpdateDefault creates a ExtrasWebhooksUpdateDefault with default headers values
func NewExtrasWebhooksUpdateDefault(code int) *ExtrasWebhooksUpdateDefault {
	return &ExtrasWebhooksUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasWebhooksUpdateDefault describes a response with status code -1, with default header values.

ExtrasWebhooksUpdateDefault extras webhooks update default
*/
type ExtrasWebhooksUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras webhooks update default response has a 2xx status code
func (o *ExtrasWebhooksUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras webhooks update default response has a 3xx status code
func (o *ExtrasWebhooksUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras webhooks update default response has a 4xx status code
func (o *ExtrasWebhooksUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras webhooks update default response has a 5xx status code
func (o *ExtrasWebhooksUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras webhooks update default response a status code equal to that given
func (o *ExtrasWebhooksUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras webhooks update default response
func (o *ExtrasWebhooksUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extras_webhooks_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extras_webhooks_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
