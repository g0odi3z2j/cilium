// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetClusterNodesReader is a Reader for the GetClusterNodes structure.
type GetClusterNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterNodesOK creates a GetClusterNodesOK with default headers values
func NewGetClusterNodesOK() *GetClusterNodesOK {
	return &GetClusterNodesOK{}
}

/*
GetClusterNodesOK describes a response with status code 200, with default header values.

Success
*/
type GetClusterNodesOK struct {
	Payload *models.ClusterNodeStatus
}

// IsSuccess returns true when this get cluster nodes o k response has a 2xx status code
func (o *GetClusterNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster nodes o k response has a 3xx status code
func (o *GetClusterNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster nodes o k response has a 4xx status code
func (o *GetClusterNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster nodes o k response has a 5xx status code
func (o *GetClusterNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster nodes o k response a status code equal to that given
func (o *GetClusterNodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterNodesOK) Error() string {
	return fmt.Sprintf("[GET /cluster/nodes][%d] getClusterNodesOK  %+v", 200, o.Payload)
}

func (o *GetClusterNodesOK) String() string {
	return fmt.Sprintf("[GET /cluster/nodes][%d] getClusterNodesOK  %+v", 200, o.Payload)
}

func (o *GetClusterNodesOK) GetPayload() *models.ClusterNodeStatus {
	return o.Payload
}

func (o *GetClusterNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNodeStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
