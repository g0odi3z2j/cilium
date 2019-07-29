// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVpcEndpointConnectionNotificationInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// One or more endpoint events for which to receive notifications. Valid values
	// are Accept, Connect, Delete, and Reject.
	//
	// ConnectionEvents is a required field
	ConnectionEvents []string `locationNameList:"item" type:"list" required:"true"`

	// The ARN of the SNS topic for the notifications.
	//
	// ConnectionNotificationArn is a required field
	ConnectionNotificationArn *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the endpoint service.
	ServiceId *string `type:"string"`

	// The ID of the endpoint.
	VpcEndpointId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpcEndpointConnectionNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcEndpointConnectionNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcEndpointConnectionNotificationInput"}

	if s.ConnectionEvents == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionEvents"))
	}

	if s.ConnectionNotificationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionNotificationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVpcEndpointConnectionNotificationOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the notification.
	ConnectionNotification *ConnectionNotification `locationName:"connectionNotification" type:"structure"`
}

// String returns the string representation
func (s CreateVpcEndpointConnectionNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcEndpointConnectionNotification = "CreateVpcEndpointConnectionNotification"

// CreateVpcEndpointConnectionNotificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a connection notification for a specified VPC endpoint or VPC endpoint
// service. A connection notification notifies you of specific endpoint events.
// You must create an SNS topic to receive notifications. For more information,
// see Create a Topic (https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html)
// in the Amazon Simple Notification Service Developer Guide.
//
// You can create a connection notification for interface endpoints only.
//
//    // Example sending a request using CreateVpcEndpointConnectionNotificationRequest.
//    req := client.CreateVpcEndpointConnectionNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointConnectionNotification
func (c *Client) CreateVpcEndpointConnectionNotificationRequest(input *CreateVpcEndpointConnectionNotificationInput) CreateVpcEndpointConnectionNotificationRequest {
	op := &aws.Operation{
		Name:       opCreateVpcEndpointConnectionNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcEndpointConnectionNotificationInput{}
	}

	req := c.newRequest(op, input, &CreateVpcEndpointConnectionNotificationOutput{})
	return CreateVpcEndpointConnectionNotificationRequest{Request: req, Input: input, Copy: c.CreateVpcEndpointConnectionNotificationRequest}
}

// CreateVpcEndpointConnectionNotificationRequest is the request type for the
// CreateVpcEndpointConnectionNotification API operation.
type CreateVpcEndpointConnectionNotificationRequest struct {
	*aws.Request
	Input *CreateVpcEndpointConnectionNotificationInput
	Copy  func(*CreateVpcEndpointConnectionNotificationInput) CreateVpcEndpointConnectionNotificationRequest
}

// Send marshals and sends the CreateVpcEndpointConnectionNotification API request.
func (r CreateVpcEndpointConnectionNotificationRequest) Send(ctx context.Context) (*CreateVpcEndpointConnectionNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcEndpointConnectionNotificationResponse{
		CreateVpcEndpointConnectionNotificationOutput: r.Request.Data.(*CreateVpcEndpointConnectionNotificationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcEndpointConnectionNotificationResponse is the response type for the
// CreateVpcEndpointConnectionNotification API operation.
type CreateVpcEndpointConnectionNotificationResponse struct {
	*CreateVpcEndpointConnectionNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcEndpointConnectionNotification request.
func (r *CreateVpcEndpointConnectionNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
