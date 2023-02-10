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

	"github.com/sklirg/go-netbox/netbox/models"
)

// VirtualizationClusterGroupsBulkUpdateReader is a Reader for the VirtualizationClusterGroupsBulkUpdate structure.
type VirtualizationClusterGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsBulkUpdateOK creates a VirtualizationClusterGroupsBulkUpdateOK with default headers values
func NewVirtualizationClusterGroupsBulkUpdateOK() *VirtualizationClusterGroupsBulkUpdateOK {
	return &VirtualizationClusterGroupsBulkUpdateOK{}
}

/*
VirtualizationClusterGroupsBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsBulkUpdateOK virtualization cluster groups bulk update o k
*/
type VirtualizationClusterGroupsBulkUpdateOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups bulk update o k response has a 2xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups bulk update o k response has a 3xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups bulk update o k response has a 4xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups bulk update o k response has a 5xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups bulk update o k response a status code equal to that given
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster groups bulk update o k response
func (o *VirtualizationClusterGroupsBulkUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsBulkUpdateDefault creates a VirtualizationClusterGroupsBulkUpdateDefault with default headers values
func NewVirtualizationClusterGroupsBulkUpdateDefault(code int) *VirtualizationClusterGroupsBulkUpdateDefault {
	return &VirtualizationClusterGroupsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterGroupsBulkUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterGroupsBulkUpdateDefault virtualization cluster groups bulk update default
*/
type VirtualizationClusterGroupsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups bulk update default response has a 2xx status code
func (o *VirtualizationClusterGroupsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster groups bulk update default response has a 3xx status code
func (o *VirtualizationClusterGroupsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster groups bulk update default response has a 4xx status code
func (o *VirtualizationClusterGroupsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster groups bulk update default response has a 5xx status code
func (o *VirtualizationClusterGroupsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster groups bulk update default response a status code equal to that given
func (o *VirtualizationClusterGroupsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization cluster groups bulk update default response
func (o *VirtualizationClusterGroupsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualization_cluster-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualization_cluster-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
