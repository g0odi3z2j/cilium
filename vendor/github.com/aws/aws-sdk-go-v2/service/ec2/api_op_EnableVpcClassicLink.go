// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableVpcClassicLinkInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableVpcClassicLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableVpcClassicLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableVpcClassicLinkInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableVpcClassicLinkOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s EnableVpcClassicLinkOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableVpcClassicLink = "EnableVpcClassicLink"

// EnableVpcClassicLinkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables a VPC for ClassicLink. You can then link EC2-Classic instances to
// your ClassicLink-enabled VPC to allow communication over private IP addresses.
// You cannot enable your VPC for ClassicLink if any of your VPC route tables
// have existing routes for address ranges within the 10.0.0.0/8 IP address
// range, excluding local routes for VPCs in the 10.0.0.0/16 and 10.1.0.0/16
// IP address ranges. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using EnableVpcClassicLinkRequest.
//    req := client.EnableVpcClassicLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVpcClassicLink
func (c *Client) EnableVpcClassicLinkRequest(input *EnableVpcClassicLinkInput) EnableVpcClassicLinkRequest {
	op := &aws.Operation{
		Name:       opEnableVpcClassicLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableVpcClassicLinkInput{}
	}

	req := c.newRequest(op, input, &EnableVpcClassicLinkOutput{})
	return EnableVpcClassicLinkRequest{Request: req, Input: input, Copy: c.EnableVpcClassicLinkRequest}
}

// EnableVpcClassicLinkRequest is the request type for the
// EnableVpcClassicLink API operation.
type EnableVpcClassicLinkRequest struct {
	*aws.Request
	Input *EnableVpcClassicLinkInput
	Copy  func(*EnableVpcClassicLinkInput) EnableVpcClassicLinkRequest
}

// Send marshals and sends the EnableVpcClassicLink API request.
func (r EnableVpcClassicLinkRequest) Send(ctx context.Context) (*EnableVpcClassicLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableVpcClassicLinkResponse{
		EnableVpcClassicLinkOutput: r.Request.Data.(*EnableVpcClassicLinkOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableVpcClassicLinkResponse is the response type for the
// EnableVpcClassicLink API operation.
type EnableVpcClassicLinkResponse struct {
	*EnableVpcClassicLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableVpcClassicLink request.
func (r *EnableVpcClassicLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
