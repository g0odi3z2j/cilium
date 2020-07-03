// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterInstanceEventNotificationAttributesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Information about the tag keys to deregister.
	InstanceTagAttribute *DeregisterInstanceTagAttributeRequest `type:"structure"`
}

// String returns the string representation
func (s DeregisterInstanceEventNotificationAttributesInput) String() string {
	return awsutil.Prettify(s)
}

type DeregisterInstanceEventNotificationAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The resulting set of tag keys.
	InstanceTagAttribute *InstanceTagNotificationAttribute `locationName:"instanceTagAttribute" type:"structure"`
}

// String returns the string representation
func (s DeregisterInstanceEventNotificationAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterInstanceEventNotificationAttributes = "DeregisterInstanceEventNotificationAttributes"

// DeregisterInstanceEventNotificationAttributesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deregisters tag keys to prevent tags that have the specified tag keys from
// being included in scheduled event notifications for resources in the Region.
//
//    // Example sending a request using DeregisterInstanceEventNotificationAttributesRequest.
//    req := client.DeregisterInstanceEventNotificationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeregisterInstanceEventNotificationAttributes
func (c *Client) DeregisterInstanceEventNotificationAttributesRequest(input *DeregisterInstanceEventNotificationAttributesInput) DeregisterInstanceEventNotificationAttributesRequest {
	op := &aws.Operation{
		Name:       opDeregisterInstanceEventNotificationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterInstanceEventNotificationAttributesInput{}
	}

	req := c.newRequest(op, input, &DeregisterInstanceEventNotificationAttributesOutput{})

	return DeregisterInstanceEventNotificationAttributesRequest{Request: req, Input: input, Copy: c.DeregisterInstanceEventNotificationAttributesRequest}
}

// DeregisterInstanceEventNotificationAttributesRequest is the request type for the
// DeregisterInstanceEventNotificationAttributes API operation.
type DeregisterInstanceEventNotificationAttributesRequest struct {
	*aws.Request
	Input *DeregisterInstanceEventNotificationAttributesInput
	Copy  func(*DeregisterInstanceEventNotificationAttributesInput) DeregisterInstanceEventNotificationAttributesRequest
}

// Send marshals and sends the DeregisterInstanceEventNotificationAttributes API request.
func (r DeregisterInstanceEventNotificationAttributesRequest) Send(ctx context.Context) (*DeregisterInstanceEventNotificationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterInstanceEventNotificationAttributesResponse{
		DeregisterInstanceEventNotificationAttributesOutput: r.Request.Data.(*DeregisterInstanceEventNotificationAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterInstanceEventNotificationAttributesResponse is the response type for the
// DeregisterInstanceEventNotificationAttributes API operation.
type DeregisterInstanceEventNotificationAttributesResponse struct {
	*DeregisterInstanceEventNotificationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterInstanceEventNotificationAttributes request.
func (r *DeregisterInstanceEventNotificationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
