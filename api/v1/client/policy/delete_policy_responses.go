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

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeletePolicyInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeletePolicyFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePolicyOK creates a DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {
	return &DeletePolicyOK{}
}

/*DeletePolicyOK handles this case with default header values.

Success
*/
type DeletePolicyOK struct {
	Payload *models.Policy
}

func (o *DeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /policy][%d] deletePolicyOK  %+v", 200, o.Payload)
}

func (o *DeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyInvalid creates a DeletePolicyInvalid with default headers values
func NewDeletePolicyInvalid() *DeletePolicyInvalid {
	return &DeletePolicyInvalid{}
}

/*DeletePolicyInvalid handles this case with default header values.

Invalid request
*/
type DeletePolicyInvalid struct {
	Payload models.Error
}

func (o *DeletePolicyInvalid) Error() string {
	return fmt.Sprintf("[DELETE /policy][%d] deletePolicyInvalid  %+v", 400, o.Payload)
}

func (o *DeletePolicyInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyNotFound creates a DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

/*DeletePolicyNotFound handles this case with default header values.

Policy not found
*/
type DeletePolicyNotFound struct {
}

func (o *DeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policy][%d] deletePolicyNotFound ", 404)
}

func (o *DeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyFailure creates a DeletePolicyFailure with default headers values
func NewDeletePolicyFailure() *DeletePolicyFailure {
	return &DeletePolicyFailure{}
}

/*DeletePolicyFailure handles this case with default header values.

Error while deleting policy
*/
type DeletePolicyFailure struct {
	Payload models.Error
}

func (o *DeletePolicyFailure) Error() string {
	return fmt.Sprintf("[DELETE /policy][%d] deletePolicyFailure  %+v", 500, o.Payload)
}

func (o *DeletePolicyFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
