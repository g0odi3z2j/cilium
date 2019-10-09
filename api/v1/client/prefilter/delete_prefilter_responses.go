// Code generated by go-swagger; DO NOT EDIT.

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// DeletePrefilterReader is a Reader for the DeletePrefilter structure.
type DeletePrefilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrefilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePrefilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 461:
		result := NewDeletePrefilterInvalidCIDR()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeletePrefilterFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePrefilterOK creates a DeletePrefilterOK with default headers values
func NewDeletePrefilterOK() *DeletePrefilterOK {
	return &DeletePrefilterOK{}
}

/*DeletePrefilterOK handles this case with default header values.

Deleted
*/
type DeletePrefilterOK struct {
	Payload *models.Prefilter
}

func (o *DeletePrefilterOK) Error() string {
	return fmt.Sprintf("[DELETE /prefilter][%d] deletePrefilterOK  %+v", 200, o.Payload)
}

func (o *DeletePrefilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrefilterInvalidCIDR creates a DeletePrefilterInvalidCIDR with default headers values
func NewDeletePrefilterInvalidCIDR() *DeletePrefilterInvalidCIDR {
	return &DeletePrefilterInvalidCIDR{}
}

/*DeletePrefilterInvalidCIDR handles this case with default header values.

Invalid CIDR prefix
*/
type DeletePrefilterInvalidCIDR struct {
	Payload models.Error
}

func (o *DeletePrefilterInvalidCIDR) Error() string {
	return fmt.Sprintf("[DELETE /prefilter][%d] deletePrefilterInvalidCIdR  %+v", 461, o.Payload)
}

func (o *DeletePrefilterInvalidCIDR) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrefilterFailure creates a DeletePrefilterFailure with default headers values
func NewDeletePrefilterFailure() *DeletePrefilterFailure {
	return &DeletePrefilterFailure{}
}

/*DeletePrefilterFailure handles this case with default header values.

Prefilter delete failed
*/
type DeletePrefilterFailure struct {
	Payload models.Error
}

func (o *DeletePrefilterFailure) Error() string {
	return fmt.Sprintf("[DELETE /prefilter][%d] deletePrefilterFailure  %+v", 500, o.Payload)
}

func (o *DeletePrefilterFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
