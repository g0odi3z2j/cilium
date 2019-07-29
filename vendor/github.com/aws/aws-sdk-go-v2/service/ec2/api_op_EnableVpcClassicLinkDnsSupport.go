// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableVpcClassicLinkDnsSupportInput struct {
	_ struct{} `type:"structure"`

	// The ID of the VPC.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s EnableVpcClassicLinkDnsSupportInput) String() string {
	return awsutil.Prettify(s)
}

type EnableVpcClassicLinkDnsSupportOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s EnableVpcClassicLinkDnsSupportOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableVpcClassicLinkDnsSupport = "EnableVpcClassicLinkDnsSupport"

// EnableVpcClassicLinkDnsSupportRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables a VPC to support DNS hostname resolution for ClassicLink. If enabled,
// the DNS hostname of a linked EC2-Classic instance resolves to its private
// IP address when addressed from an instance in the VPC to which it's linked.
// Similarly, the DNS hostname of an instance in a VPC resolves to its private
// IP address when addressed from a linked EC2-Classic instance. For more information,
// see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using EnableVpcClassicLinkDnsSupportRequest.
//    req := client.EnableVpcClassicLinkDnsSupportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVpcClassicLinkDnsSupport
func (c *Client) EnableVpcClassicLinkDnsSupportRequest(input *EnableVpcClassicLinkDnsSupportInput) EnableVpcClassicLinkDnsSupportRequest {
	op := &aws.Operation{
		Name:       opEnableVpcClassicLinkDnsSupport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableVpcClassicLinkDnsSupportInput{}
	}

	req := c.newRequest(op, input, &EnableVpcClassicLinkDnsSupportOutput{})
	return EnableVpcClassicLinkDnsSupportRequest{Request: req, Input: input, Copy: c.EnableVpcClassicLinkDnsSupportRequest}
}

// EnableVpcClassicLinkDnsSupportRequest is the request type for the
// EnableVpcClassicLinkDnsSupport API operation.
type EnableVpcClassicLinkDnsSupportRequest struct {
	*aws.Request
	Input *EnableVpcClassicLinkDnsSupportInput
	Copy  func(*EnableVpcClassicLinkDnsSupportInput) EnableVpcClassicLinkDnsSupportRequest
}

// Send marshals and sends the EnableVpcClassicLinkDnsSupport API request.
func (r EnableVpcClassicLinkDnsSupportRequest) Send(ctx context.Context) (*EnableVpcClassicLinkDnsSupportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableVpcClassicLinkDnsSupportResponse{
		EnableVpcClassicLinkDnsSupportOutput: r.Request.Data.(*EnableVpcClassicLinkDnsSupportOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableVpcClassicLinkDnsSupportResponse is the response type for the
// EnableVpcClassicLinkDnsSupport API operation.
type EnableVpcClassicLinkDnsSupportResponse struct {
	*EnableVpcClassicLinkDnsSupportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableVpcClassicLinkDnsSupport request.
func (r *EnableVpcClassicLinkDnsSupportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
