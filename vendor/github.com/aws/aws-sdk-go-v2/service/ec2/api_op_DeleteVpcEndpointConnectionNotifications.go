// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteVpcEndpointConnectionNotificationsInput struct {
	_ struct{} `type:"structure"`

	// One or more notification IDs.
	//
	// ConnectionNotificationIds is a required field
	ConnectionNotificationIds []string `locationName:"ConnectionNotificationId" locationNameList:"item" type:"list" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteVpcEndpointConnectionNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcEndpointConnectionNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpcEndpointConnectionNotificationsInput"}

	if s.ConnectionNotificationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionNotificationIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVpcEndpointConnectionNotificationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the notifications that could not be deleted successfully.
	Unsuccessful []UnsuccessfulItem `locationName:"unsuccessful" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteVpcEndpointConnectionNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVpcEndpointConnectionNotifications = "DeleteVpcEndpointConnectionNotifications"

// DeleteVpcEndpointConnectionNotificationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes one or more VPC endpoint connection notifications.
//
//    // Example sending a request using DeleteVpcEndpointConnectionNotificationsRequest.
//    req := client.DeleteVpcEndpointConnectionNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications
func (c *Client) DeleteVpcEndpointConnectionNotificationsRequest(input *DeleteVpcEndpointConnectionNotificationsInput) DeleteVpcEndpointConnectionNotificationsRequest {
	op := &aws.Operation{
		Name:       opDeleteVpcEndpointConnectionNotifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcEndpointConnectionNotificationsInput{}
	}

	req := c.newRequest(op, input, &DeleteVpcEndpointConnectionNotificationsOutput{})

	return DeleteVpcEndpointConnectionNotificationsRequest{Request: req, Input: input, Copy: c.DeleteVpcEndpointConnectionNotificationsRequest}
}

// DeleteVpcEndpointConnectionNotificationsRequest is the request type for the
// DeleteVpcEndpointConnectionNotifications API operation.
type DeleteVpcEndpointConnectionNotificationsRequest struct {
	*aws.Request
	Input *DeleteVpcEndpointConnectionNotificationsInput
	Copy  func(*DeleteVpcEndpointConnectionNotificationsInput) DeleteVpcEndpointConnectionNotificationsRequest
}

// Send marshals and sends the DeleteVpcEndpointConnectionNotifications API request.
func (r DeleteVpcEndpointConnectionNotificationsRequest) Send(ctx context.Context) (*DeleteVpcEndpointConnectionNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcEndpointConnectionNotificationsResponse{
		DeleteVpcEndpointConnectionNotificationsOutput: r.Request.Data.(*DeleteVpcEndpointConnectionNotificationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcEndpointConnectionNotificationsResponse is the response type for the
// DeleteVpcEndpointConnectionNotifications API operation.
type DeleteVpcEndpointConnectionNotificationsResponse struct {
	*DeleteVpcEndpointConnectionNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpcEndpointConnectionNotifications request.
func (r *DeleteVpcEndpointConnectionNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
