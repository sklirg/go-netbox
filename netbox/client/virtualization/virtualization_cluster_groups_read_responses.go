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

// VirtualizationClusterGroupsReadReader is a Reader for the VirtualizationClusterGroupsRead structure.
type VirtualizationClusterGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsReadOK creates a VirtualizationClusterGroupsReadOK with default headers values
func NewVirtualizationClusterGroupsReadOK() *VirtualizationClusterGroupsReadOK {
	return &VirtualizationClusterGroupsReadOK{}
}

/*
VirtualizationClusterGroupsReadOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsReadOK virtualization cluster groups read o k
*/
type VirtualizationClusterGroupsReadOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups read o k response has a 2xx status code
func (o *VirtualizationClusterGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups read o k response has a 3xx status code
func (o *VirtualizationClusterGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups read o k response has a 4xx status code
func (o *VirtualizationClusterGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups read o k response has a 5xx status code
func (o *VirtualizationClusterGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups read o k response a status code equal to that given
func (o *VirtualizationClusterGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster groups read o k response
func (o *VirtualizationClusterGroupsReadOK) Code() int {
	return 200
}

func (o *VirtualizationClusterGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsReadOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsReadDefault creates a VirtualizationClusterGroupsReadDefault with default headers values
func NewVirtualizationClusterGroupsReadDefault(code int) *VirtualizationClusterGroupsReadDefault {
	return &VirtualizationClusterGroupsReadDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterGroupsReadDefault describes a response with status code -1, with default header values.

VirtualizationClusterGroupsReadDefault virtualization cluster groups read default
*/
type VirtualizationClusterGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups read default response has a 2xx status code
func (o *VirtualizationClusterGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster groups read default response has a 3xx status code
func (o *VirtualizationClusterGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster groups read default response has a 4xx status code
func (o *VirtualizationClusterGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster groups read default response has a 5xx status code
func (o *VirtualizationClusterGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster groups read default response a status code equal to that given
func (o *VirtualizationClusterGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization cluster groups read default response
func (o *VirtualizationClusterGroupsReadDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
