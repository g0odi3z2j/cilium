// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// GetServiceIDReader is a Reader for the GetServiceID structure.
type GetServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceIDOK creates a GetServiceIDOK with default headers values
func NewGetServiceIDOK() *GetServiceIDOK {
	return &GetServiceIDOK{}
}

/*GetServiceIDOK handles this case with default header values.

Success
*/
type GetServiceIDOK struct {
	Payload *models.Service
}

func (o *GetServiceIDOK) Error() string {
	return fmt.Sprintf("[GET /service/{id}][%d] getServiceIdOK  %+v", 200, o.Payload)
}

func (o *GetServiceIDOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *GetServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceIDNotFound creates a GetServiceIDNotFound with default headers values
func NewGetServiceIDNotFound() *GetServiceIDNotFound {
	return &GetServiceIDNotFound{}
}

/*GetServiceIDNotFound handles this case with default header values.

Service not found
*/
type GetServiceIDNotFound struct {
}

func (o *GetServiceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /service/{id}][%d] getServiceIdNotFound ", 404)
}

func (o *GetServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
