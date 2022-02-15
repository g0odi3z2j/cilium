// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetPolicyResolveReader is a Reader for the GetPolicyResolve structure.
type GetPolicyResolveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyResolveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyResolveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicyResolveOK creates a GetPolicyResolveOK with default headers values
func NewGetPolicyResolveOK() *GetPolicyResolveOK {
	return &GetPolicyResolveOK{}
}

/*GetPolicyResolveOK handles this case with default header values.

Success
*/
type GetPolicyResolveOK struct {
	Payload *models.PolicyTraceResult
}

func (o *GetPolicyResolveOK) Error() string {
	return fmt.Sprintf("[GET /policy/resolve][%d] getPolicyResolveOK  %+v", 200, o.Payload)
}

func (o *GetPolicyResolveOK) GetPayload() *models.PolicyTraceResult {
	return o.Payload
}

func (o *GetPolicyResolveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTraceResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
