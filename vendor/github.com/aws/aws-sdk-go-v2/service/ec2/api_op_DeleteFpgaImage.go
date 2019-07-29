// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFpgaImageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFpgaImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFpgaImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFpgaImageInput"}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFpgaImageOutput struct {
	_ struct{} `type:"structure"`

	// Is true if the request succeeds, and an error otherwise.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteFpgaImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFpgaImage = "DeleteFpgaImage"

// DeleteFpgaImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Amazon FPGA Image (AFI).
//
//    // Example sending a request using DeleteFpgaImageRequest.
//    req := client.DeleteFpgaImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteFpgaImage
func (c *Client) DeleteFpgaImageRequest(input *DeleteFpgaImageInput) DeleteFpgaImageRequest {
	op := &aws.Operation{
		Name:       opDeleteFpgaImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFpgaImageInput{}
	}

	req := c.newRequest(op, input, &DeleteFpgaImageOutput{})
	return DeleteFpgaImageRequest{Request: req, Input: input, Copy: c.DeleteFpgaImageRequest}
}

// DeleteFpgaImageRequest is the request type for the
// DeleteFpgaImage API operation.
type DeleteFpgaImageRequest struct {
	*aws.Request
	Input *DeleteFpgaImageInput
	Copy  func(*DeleteFpgaImageInput) DeleteFpgaImageRequest
}

// Send marshals and sends the DeleteFpgaImage API request.
func (r DeleteFpgaImageRequest) Send(ctx context.Context) (*DeleteFpgaImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFpgaImageResponse{
		DeleteFpgaImageOutput: r.Request.Data.(*DeleteFpgaImageOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFpgaImageResponse is the response type for the
// DeleteFpgaImage API operation.
type DeleteFpgaImageResponse struct {
	*DeleteFpgaImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFpgaImage request.
func (r *DeleteFpgaImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
