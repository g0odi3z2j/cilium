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

// GetIdentityIDReader is a Reader for the GetIdentityID structure.
type GetIdentityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdentityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetIdentityIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetIdentityIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 520:
		result := NewGetIdentityIDUnreachable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 521:
		result := NewGetIdentityIDInvalidStorageFormat()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentityIDOK creates a GetIdentityIDOK with default headers values
func NewGetIdentityIDOK() *GetIdentityIDOK {
	return &GetIdentityIDOK{}
}

/*GetIdentityIDOK handles this case with default header values.

Success
*/
type GetIdentityIDOK struct {
	Payload *models.Identity
}

func (o *GetIdentityIDOK) Error() string {
	return fmt.Sprintf("[GET /identity/{id}][%d] getIdentityIdOK  %+v", 200, o.Payload)
}

func (o *GetIdentityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Identity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityIDBadRequest creates a GetIdentityIDBadRequest with default headers values
func NewGetIdentityIDBadRequest() *GetIdentityIDBadRequest {
	return &GetIdentityIDBadRequest{}
}

/*GetIdentityIDBadRequest handles this case with default header values.

Invalid identity provided
*/
type GetIdentityIDBadRequest struct {
}

func (o *GetIdentityIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /identity/{id}][%d] getIdentityIdBadRequest ", 400)
}

func (o *GetIdentityIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdentityIDNotFound creates a GetIdentityIDNotFound with default headers values
func NewGetIdentityIDNotFound() *GetIdentityIDNotFound {
	return &GetIdentityIDNotFound{}
}

/*GetIdentityIDNotFound handles this case with default header values.

Identity not found
*/
type GetIdentityIDNotFound struct {
}

func (o *GetIdentityIDNotFound) Error() string {
	return fmt.Sprintf("[GET /identity/{id}][%d] getIdentityIdNotFound ", 404)
}

func (o *GetIdentityIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdentityIDUnreachable creates a GetIdentityIDUnreachable with default headers values
func NewGetIdentityIDUnreachable() *GetIdentityIDUnreachable {
	return &GetIdentityIDUnreachable{}
}

/*GetIdentityIDUnreachable handles this case with default header values.

Identity storage unreachable. Likely a network problem.
*/
type GetIdentityIDUnreachable struct {
	Payload models.Error
}

func (o *GetIdentityIDUnreachable) Error() string {
	return fmt.Sprintf("[GET /identity/{id}][%d] getIdentityIdUnreachable  %+v", 520, o.Payload)
}

func (o *GetIdentityIDUnreachable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityIDInvalidStorageFormat creates a GetIdentityIDInvalidStorageFormat with default headers values
func NewGetIdentityIDInvalidStorageFormat() *GetIdentityIDInvalidStorageFormat {
	return &GetIdentityIDInvalidStorageFormat{}
}

/*GetIdentityIDInvalidStorageFormat handles this case with default header values.

Invalid identity format in storage
*/
type GetIdentityIDInvalidStorageFormat struct {
	Payload models.Error
}

func (o *GetIdentityIDInvalidStorageFormat) Error() string {
	return fmt.Sprintf("[GET /identity/{id}][%d] getIdentityIdInvalidStorageFormat  %+v", 521, o.Payload)
}

func (o *GetIdentityIDInvalidStorageFormat) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
