// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type ModifyVolumeAttributeInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the volume should be auto-enabled for I/O operations.
	AutoEnableIO *AttributeBooleanValue `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVolumeAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVolumeAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVolumeAttributeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVolumeAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyVolumeAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVolumeAttribute = "ModifyVolumeAttribute"

// ModifyVolumeAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies a volume attribute.
//
// By default, all I/O operations for the volume are suspended when the data
// on the volume is determined to be potentially inconsistent, to prevent undetectable,
// latent data corruption. The I/O access to the volume can be resumed by first
// enabling I/O access and then checking the data consistency on your volume.
//
// You can change the default behavior to resume I/O operations. We recommend
// that you change this only for boot volumes or for volumes that are stateless
// or disposable.
//
//    // Example sending a request using ModifyVolumeAttributeRequest.
//    req := client.ModifyVolumeAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVolumeAttribute
func (c *Client) ModifyVolumeAttributeRequest(input *ModifyVolumeAttributeInput) ModifyVolumeAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyVolumeAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVolumeAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyVolumeAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifyVolumeAttributeRequest{Request: req, Input: input, Copy: c.ModifyVolumeAttributeRequest}
}

// ModifyVolumeAttributeRequest is the request type for the
// ModifyVolumeAttribute API operation.
type ModifyVolumeAttributeRequest struct {
	*aws.Request
	Input *ModifyVolumeAttributeInput
	Copy  func(*ModifyVolumeAttributeInput) ModifyVolumeAttributeRequest
}

// Send marshals and sends the ModifyVolumeAttribute API request.
func (r ModifyVolumeAttributeRequest) Send(ctx context.Context) (*ModifyVolumeAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVolumeAttributeResponse{
		ModifyVolumeAttributeOutput: r.Request.Data.(*ModifyVolumeAttributeOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVolumeAttributeResponse is the response type for the
// ModifyVolumeAttribute API operation.
type ModifyVolumeAttributeResponse struct {
	*ModifyVolumeAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVolumeAttribute request.
func (r *ModifyVolumeAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
