// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetIdentityReader is a Reader for the GetIdentity structure.
type GetIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 520:
		result := NewGetIdentityUnreachable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 521:
		result := NewGetIdentityInvalidStorageFormat()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentityOK creates a GetIdentityOK with default headers values
func NewGetIdentityOK() *GetIdentityOK {
	return &GetIdentityOK{}
}

/*GetIdentityOK handles this case with default header values.

Success
*/
type GetIdentityOK struct {
	Payload []*models.Identity
}

func (o *GetIdentityOK) Error() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityOK  %+v", 200, o.Payload)
}

func (o *GetIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityNotFound creates a GetIdentityNotFound with default headers values
func NewGetIdentityNotFound() *GetIdentityNotFound {
	return &GetIdentityNotFound{}
}

/*GetIdentityNotFound handles this case with default header values.

Identities with provided parameters not found
*/
type GetIdentityNotFound struct {
}

func (o *GetIdentityNotFound) Error() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityNotFound ", 404)
}

func (o *GetIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdentityUnreachable creates a GetIdentityUnreachable with default headers values
func NewGetIdentityUnreachable() *GetIdentityUnreachable {
	return &GetIdentityUnreachable{}
}

/*GetIdentityUnreachable handles this case with default header values.

Identity storage unreachable. Likely a network problem.
*/
type GetIdentityUnreachable struct {
	Payload models.Error
}

func (o *GetIdentityUnreachable) Error() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityUnreachable  %+v", 520, o.Payload)
}

func (o *GetIdentityUnreachable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityInvalidStorageFormat creates a GetIdentityInvalidStorageFormat with default headers values
func NewGetIdentityInvalidStorageFormat() *GetIdentityInvalidStorageFormat {
	return &GetIdentityInvalidStorageFormat{}
}

/*GetIdentityInvalidStorageFormat handles this case with default header values.

Invalid identity format in storage
*/
type GetIdentityInvalidStorageFormat struct {
	Payload models.Error
}

func (o *GetIdentityInvalidStorageFormat) Error() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityInvalidStorageFormat  %+v", 521, o.Payload)
}

func (o *GetIdentityInvalidStorageFormat) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
