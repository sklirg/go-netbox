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

// WirelessWirelessLinksReadReader is a Reader for the WirelessWirelessLinksRead structure.
type WirelessWirelessLinksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksReadOK creates a WirelessWirelessLinksReadOK with default headers values
func NewWirelessWirelessLinksReadOK() *WirelessWirelessLinksReadOK {
	return &WirelessWirelessLinksReadOK{}
}

/*
WirelessWirelessLinksReadOK describes a response with status code 200, with default header values.

WirelessWirelessLinksReadOK wireless wireless links read o k
*/
type WirelessWirelessLinksReadOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links read o k response has a 2xx status code
func (o *WirelessWirelessLinksReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links read o k response has a 3xx status code
func (o *WirelessWirelessLinksReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links read o k response has a 4xx status code
func (o *WirelessWirelessLinksReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links read o k response has a 5xx status code
func (o *WirelessWirelessLinksReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links read o k response a status code equal to that given
func (o *WirelessWirelessLinksReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links read o k response
func (o *WirelessWirelessLinksReadOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksReadOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksReadOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksReadOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLinksReadDefault creates a WirelessWirelessLinksReadDefault with default headers values
func NewWirelessWirelessLinksReadDefault(code int) *WirelessWirelessLinksReadDefault {
	return &WirelessWirelessLinksReadDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksReadDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksReadDefault wireless wireless links read default
*/
type WirelessWirelessLinksReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless links read default response has a 2xx status code
func (o *WirelessWirelessLinksReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links read default response has a 3xx status code
func (o *WirelessWirelessLinksReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links read default response has a 4xx status code
func (o *WirelessWirelessLinksReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links read default response has a 5xx status code
func (o *WirelessWirelessLinksReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links read default response a status code equal to that given
func (o *WirelessWirelessLinksReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless links read default response
func (o *WirelessWirelessLinksReadDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLinksReadDefault) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wireless_wireless-links_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksReadDefault) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wireless_wireless-links_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
