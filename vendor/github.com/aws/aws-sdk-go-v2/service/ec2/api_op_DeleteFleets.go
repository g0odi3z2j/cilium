// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFleetsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the EC2 Fleets.
	//
	// FleetIds is a required field
	FleetIds []string `locationName:"FleetId" type:"list" required:"true"`

	// Indicates whether to terminate instances for an EC2 Fleet if it is deleted
	// successfully.
	//
	// TerminateInstances is a required field
	TerminateInstances *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s DeleteFleetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFleetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFleetsInput"}

	if s.FleetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetIds"))
	}

	if s.TerminateInstances == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminateInstances"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFleetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the EC2 Fleets that are successfully deleted.
	SuccessfulFleetDeletions []DeleteFleetSuccessItem `locationName:"successfulFleetDeletionSet" locationNameList:"item" type:"list"`

	// Information about the EC2 Fleets that are not successfully deleted.
	UnsuccessfulFleetDeletions []DeleteFleetErrorItem `locationName:"unsuccessfulFleetDeletionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteFleetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFleets = "DeleteFleets"

// DeleteFleetsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified EC2 Fleet.
//
// After you delete an EC2 Fleet, it launches no new instances. You must specify
// whether an EC2 Fleet should also terminate its instances. If you terminate
// the instances, the EC2 Fleet enters the deleted_terminating state. Otherwise,
// the EC2 Fleet enters the deleted_running state, and the instances continue
// to run until they are interrupted or you terminate them manually.
//
//    // Example sending a request using DeleteFleetsRequest.
//    req := client.DeleteFleetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteFleets
func (c *Client) DeleteFleetsRequest(input *DeleteFleetsInput) DeleteFleetsRequest {
	op := &aws.Operation{
		Name:       opDeleteFleets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFleetsInput{}
	}

	req := c.newRequest(op, input, &DeleteFleetsOutput{})

	return DeleteFleetsRequest{Request: req, Input: input, Copy: c.DeleteFleetsRequest}
}

// DeleteFleetsRequest is the request type for the
// DeleteFleets API operation.
type DeleteFleetsRequest struct {
	*aws.Request
	Input *DeleteFleetsInput
	Copy  func(*DeleteFleetsInput) DeleteFleetsRequest
}

// Send marshals and sends the DeleteFleets API request.
func (r DeleteFleetsRequest) Send(ctx context.Context) (*DeleteFleetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFleetsResponse{
		DeleteFleetsOutput: r.Request.Data.(*DeleteFleetsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFleetsResponse is the response type for the
// DeleteFleets API operation.
type DeleteFleetsResponse struct {
	*DeleteFleetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFleets request.
func (r *DeleteFleetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
