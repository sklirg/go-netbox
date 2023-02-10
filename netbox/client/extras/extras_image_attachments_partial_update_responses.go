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

// ExtrasImageAttachmentsPartialUpdateReader is a Reader for the ExtrasImageAttachmentsPartialUpdate structure.
type ExtrasImageAttachmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsPartialUpdateOK creates a ExtrasImageAttachmentsPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsPartialUpdateOK() *ExtrasImageAttachmentsPartialUpdateOK {
	return &ExtrasImageAttachmentsPartialUpdateOK{}
}

/*
ExtrasImageAttachmentsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsPartialUpdateOK extras image attachments partial update o k
*/
type ExtrasImageAttachmentsPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments partial update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments partial update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments partial update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments partial update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments partial update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras image attachments partial update o k response
func (o *ExtrasImageAttachmentsPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsPartialUpdateDefault creates a ExtrasImageAttachmentsPartialUpdateDefault with default headers values
func NewExtrasImageAttachmentsPartialUpdateDefault(code int) *ExtrasImageAttachmentsPartialUpdateDefault {
	return &ExtrasImageAttachmentsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsPartialUpdateDefault extras image attachments partial update default
*/
type ExtrasImageAttachmentsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras image attachments partial update default response has a 2xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments partial update default response has a 3xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments partial update default response has a 4xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments partial update default response has a 5xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments partial update default response a status code equal to that given
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras image attachments partial update default response
func (o *ExtrasImageAttachmentsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extras_image-attachments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extras_image-attachments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
