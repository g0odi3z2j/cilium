// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisableEbsEncryptionByDefaultInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DisableEbsEncryptionByDefaultInput) String() string {
	return awsutil.Prettify(s)
}

type DisableEbsEncryptionByDefaultOutput struct {
	_ struct{} `type:"structure"`

	// The updated status of encryption by default.
	EbsEncryptionByDefault *bool `locationName:"ebsEncryptionByDefault" type:"boolean"`
}

// String returns the string representation
func (s DisableEbsEncryptionByDefaultOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableEbsEncryptionByDefault = "DisableEbsEncryptionByDefault"

// DisableEbsEncryptionByDefaultRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disables EBS encryption by default for your account in the current Region.
//
// After you disable encryption by default, you can still create encrypted volumes
// by enabling encryption when you create each volume.
//
// Disabling encryption by default does not change the encryption status of
// your existing volumes.
//
// For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DisableEbsEncryptionByDefaultRequest.
//    req := client.DisableEbsEncryptionByDefaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisableEbsEncryptionByDefault
func (c *Client) DisableEbsEncryptionByDefaultRequest(input *DisableEbsEncryptionByDefaultInput) DisableEbsEncryptionByDefaultRequest {
	op := &aws.Operation{
		Name:       opDisableEbsEncryptionByDefault,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableEbsEncryptionByDefaultInput{}
	}

	req := c.newRequest(op, input, &DisableEbsEncryptionByDefaultOutput{})
	return DisableEbsEncryptionByDefaultRequest{Request: req, Input: input, Copy: c.DisableEbsEncryptionByDefaultRequest}
}

// DisableEbsEncryptionByDefaultRequest is the request type for the
// DisableEbsEncryptionByDefault API operation.
type DisableEbsEncryptionByDefaultRequest struct {
	*aws.Request
	Input *DisableEbsEncryptionByDefaultInput
	Copy  func(*DisableEbsEncryptionByDefaultInput) DisableEbsEncryptionByDefaultRequest
}

// Send marshals and sends the DisableEbsEncryptionByDefault API request.
func (r DisableEbsEncryptionByDefaultRequest) Send(ctx context.Context) (*DisableEbsEncryptionByDefaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableEbsEncryptionByDefaultResponse{
		DisableEbsEncryptionByDefaultOutput: r.Request.Data.(*DisableEbsEncryptionByDefaultOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableEbsEncryptionByDefaultResponse is the response type for the
// DisableEbsEncryptionByDefault API operation.
type DisableEbsEncryptionByDefaultResponse struct {
	*DisableEbsEncryptionByDefaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableEbsEncryptionByDefault request.
func (r *DisableEbsEncryptionByDefaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
