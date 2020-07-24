// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDReader is a Reader for the GetEndpointID structure.
type GetEndpointIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointIDInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointIDOK creates a GetEndpointIDOK with default headers values
func NewGetEndpointIDOK() *GetEndpointIDOK {
	return &GetEndpointIDOK{}
}

/*GetEndpointIDOK handles this case with default header values.

Success
*/
type GetEndpointIDOK struct {
	Payload *models.Endpoint
}

func (o *GetEndpointIDOK) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}][%d] getEndpointIdOK  %+v", 200, o.Payload)
}

func (o *GetEndpointIDOK) GetPayload() *models.Endpoint {
	return o.Payload
}

func (o *GetEndpointIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Endpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDInvalid creates a GetEndpointIDInvalid with default headers values
func NewGetEndpointIDInvalid() *GetEndpointIDInvalid {
	return &GetEndpointIDInvalid{}
}

/*GetEndpointIDInvalid handles this case with default header values.

Invalid endpoint ID format for specified type
*/
type GetEndpointIDInvalid struct {
	Payload models.Error
}

func (o *GetEndpointIDInvalid) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}][%d] getEndpointIdInvalid  %+v", 400, o.Payload)
}

func (o *GetEndpointIDInvalid) GetPayload() models.Error {
	return o.Payload
}

func (o *GetEndpointIDInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDNotFound creates a GetEndpointIDNotFound with default headers values
func NewGetEndpointIDNotFound() *GetEndpointIDNotFound {
	return &GetEndpointIDNotFound{}
}

/*GetEndpointIDNotFound handles this case with default header values.

Endpoint not found
*/
type GetEndpointIDNotFound struct {
}

func (o *GetEndpointIDNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}][%d] getEndpointIdNotFound ", 404)
}

func (o *GetEndpointIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
