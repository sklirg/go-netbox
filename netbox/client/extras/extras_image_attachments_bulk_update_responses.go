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

// ExtrasImageAttachmentsBulkUpdateReader is a Reader for the ExtrasImageAttachmentsBulkUpdate structure.
type ExtrasImageAttachmentsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsBulkUpdateOK creates a ExtrasImageAttachmentsBulkUpdateOK with default headers values
func NewExtrasImageAttachmentsBulkUpdateOK() *ExtrasImageAttachmentsBulkUpdateOK {
	return &ExtrasImageAttachmentsBulkUpdateOK{}
}

/*
ExtrasImageAttachmentsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsBulkUpdateOK extras image attachments bulk update o k
*/
type ExtrasImageAttachmentsBulkUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments bulk update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments bulk update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments bulk update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments bulk update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments bulk update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras image attachments bulk update o k response
func (o *ExtrasImageAttachmentsBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extrasImageAttachmentsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extrasImageAttachmentsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsBulkUpdateDefault creates a ExtrasImageAttachmentsBulkUpdateDefault with default headers values
func NewExtrasImageAttachmentsBulkUpdateDefault(code int) *ExtrasImageAttachmentsBulkUpdateDefault {
	return &ExtrasImageAttachmentsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsBulkUpdateDefault extras image attachments bulk update default
*/
type ExtrasImageAttachmentsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras image attachments bulk update default response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments bulk update default response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments bulk update default response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments bulk update default response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments bulk update default response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras image attachments bulk update default response
func (o *ExtrasImageAttachmentsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extras_image-attachments_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extras_image-attachments_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
