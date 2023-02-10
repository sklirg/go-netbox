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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersTokensBulkPartialUpdateReader is a Reader for the UsersTokensBulkPartialUpdate structure.
type UsersTokensBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensBulkPartialUpdateOK creates a UsersTokensBulkPartialUpdateOK with default headers values
func NewUsersTokensBulkPartialUpdateOK() *UsersTokensBulkPartialUpdateOK {
	return &UsersTokensBulkPartialUpdateOK{}
}

/*
UsersTokensBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersTokensBulkPartialUpdateOK users tokens bulk partial update o k
*/
type UsersTokensBulkPartialUpdateOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens bulk partial update o k response has a 2xx status code
func (o *UsersTokensBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens bulk partial update o k response has a 3xx status code
func (o *UsersTokensBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens bulk partial update o k response has a 4xx status code
func (o *UsersTokensBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens bulk partial update o k response has a 5xx status code
func (o *UsersTokensBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens bulk partial update o k response a status code equal to that given
func (o *UsersTokensBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users tokens bulk partial update o k response
func (o *UsersTokensBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *UsersTokensBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/tokens/][%d] usersTokensBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /users/tokens/][%d] usersTokensBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensBulkPartialUpdateOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensBulkPartialUpdateDefault creates a UsersTokensBulkPartialUpdateDefault with default headers values
func NewUsersTokensBulkPartialUpdateDefault(code int) *UsersTokensBulkPartialUpdateDefault {
	return &UsersTokensBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersTokensBulkPartialUpdateDefault describes a response with status code -1, with default header values.

UsersTokensBulkPartialUpdateDefault users tokens bulk partial update default
*/
type UsersTokensBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users tokens bulk partial update default response has a 2xx status code
func (o *UsersTokensBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens bulk partial update default response has a 3xx status code
func (o *UsersTokensBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens bulk partial update default response has a 4xx status code
func (o *UsersTokensBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens bulk partial update default response has a 5xx status code
func (o *UsersTokensBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens bulk partial update default response a status code equal to that given
func (o *UsersTokensBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users tokens bulk partial update default response
func (o *UsersTokensBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/tokens/][%d] users_tokens_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /users/tokens/][%d] users_tokens_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
