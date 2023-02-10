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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationClusterTypesUpdateReader is a Reader for the VirtualizationClusterTypesUpdate structure.
type VirtualizationClusterTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesUpdateOK creates a VirtualizationClusterTypesUpdateOK with default headers values
func NewVirtualizationClusterTypesUpdateOK() *VirtualizationClusterTypesUpdateOK {
	return &VirtualizationClusterTypesUpdateOK{}
}

/*
VirtualizationClusterTypesUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesUpdateOK virtualization cluster types update o k
*/
type VirtualizationClusterTypesUpdateOK struct {
	Payload *models.ClusterType
}

// IsSuccess returns true when this virtualization cluster types update o k response has a 2xx status code
func (o *VirtualizationClusterTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster types update o k response has a 3xx status code
func (o *VirtualizationClusterTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster types update o k response has a 4xx status code
func (o *VirtualizationClusterTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster types update o k response has a 5xx status code
func (o *VirtualizationClusterTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster types update o k response a status code equal to that given
func (o *VirtualizationClusterTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster types update o k response
func (o *VirtualizationClusterTypesUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClusterTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesUpdateDefault creates a VirtualizationClusterTypesUpdateDefault with default headers values
func NewVirtualizationClusterTypesUpdateDefault(code int) *VirtualizationClusterTypesUpdateDefault {
	return &VirtualizationClusterTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterTypesUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesUpdateDefault virtualization cluster types update default
*/
type VirtualizationClusterTypesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster types update default response has a 2xx status code
func (o *VirtualizationClusterTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster types update default response has a 3xx status code
func (o *VirtualizationClusterTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster types update default response has a 4xx status code
func (o *VirtualizationClusterTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster types update default response has a 5xx status code
func (o *VirtualizationClusterTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster types update default response a status code equal to that given
func (o *VirtualizationClusterTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization cluster types update default response
func (o *VirtualizationClusterTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
