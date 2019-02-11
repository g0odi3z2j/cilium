// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// PostIPAMReader is a Reader for the PostIPAM structure.
type PostIPAMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostIPAMCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 502:
		result := NewPostIPAMFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMCreated creates a PostIPAMCreated with default headers values
func NewPostIPAMCreated() *PostIPAMCreated {
	return &PostIPAMCreated{}
}

/*PostIPAMCreated handles this case with default header values.

Success
*/
type PostIPAMCreated struct {
	Payload *models.IPAMResponse
}

func (o *PostIPAMCreated) Error() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpAMCreated  %+v", 201, o.Payload)
}

func (o *PostIPAMCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPAMFailure creates a PostIPAMFailure with default headers values
func NewPostIPAMFailure() *PostIPAMFailure {
	return &PostIPAMFailure{}
}

/*PostIPAMFailure handles this case with default header values.

Allocation failure
*/
type PostIPAMFailure struct {
	Payload models.Error
}

func (o *PostIPAMFailure) Error() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpAMFailure  %+v", 502, o.Payload)
}

func (o *PostIPAMFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
