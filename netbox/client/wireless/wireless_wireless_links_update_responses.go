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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sklirg/go-netbox/netbox/models"
)

// WirelessWirelessLinksUpdateReader is a Reader for the WirelessWirelessLinksUpdate structure.
type WirelessWirelessLinksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksUpdateOK creates a WirelessWirelessLinksUpdateOK with default headers values
func NewWirelessWirelessLinksUpdateOK() *WirelessWirelessLinksUpdateOK {
	return &WirelessWirelessLinksUpdateOK{}
}

/*
WirelessWirelessLinksUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLinksUpdateOK wireless wireless links update o k
*/
type WirelessWirelessLinksUpdateOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links update o k response has a 2xx status code
func (o *WirelessWirelessLinksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links update o k response has a 3xx status code
func (o *WirelessWirelessLinksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links update o k response has a 4xx status code
func (o *WirelessWirelessLinksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links update o k response has a 5xx status code
func (o *WirelessWirelessLinksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links update o k response a status code equal to that given
func (o *WirelessWirelessLinksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links update o k response
func (o *WirelessWirelessLinksUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksUpdateOK) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksUpdateOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLinksUpdateDefault creates a WirelessWirelessLinksUpdateDefault with default headers values
func NewWirelessWirelessLinksUpdateDefault(code int) *WirelessWirelessLinksUpdateDefault {
	return &WirelessWirelessLinksUpdateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksUpdateDefault wireless wireless links update default
*/
type WirelessWirelessLinksUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless links update default response has a 2xx status code
func (o *WirelessWirelessLinksUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links update default response has a 3xx status code
func (o *WirelessWirelessLinksUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links update default response has a 4xx status code
func (o *WirelessWirelessLinksUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links update default response has a 5xx status code
func (o *WirelessWirelessLinksUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links update default response a status code equal to that given
func (o *WirelessWirelessLinksUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless links update default response
func (o *WirelessWirelessLinksUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLinksUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wireless_wireless-links_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wireless_wireless-links_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
