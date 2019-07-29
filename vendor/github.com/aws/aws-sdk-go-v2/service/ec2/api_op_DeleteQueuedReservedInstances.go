// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteQueuedReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the Reserved Instances.
	//
	// ReservedInstancesIds is a required field
	ReservedInstancesIds []string `locationName:"ReservedInstancesId" locationNameList:"item" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteQueuedReservedInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteQueuedReservedInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteQueuedReservedInstancesInput"}

	if s.ReservedInstancesIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstancesIds"))
	}
	if s.ReservedInstancesIds != nil && len(s.ReservedInstancesIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReservedInstancesIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteQueuedReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the queued purchases that could not be deleted.
	FailedQueuedPurchaseDeletions []FailedQueuedPurchaseDeletion `locationName:"failedQueuedPurchaseDeletionSet" locationNameList:"item" type:"list"`

	// Information about the queued purchases that were successfully deleted.
	SuccessfulQueuedPurchaseDeletions []SuccessfulQueuedPurchaseDeletion `locationName:"successfulQueuedPurchaseDeletionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteQueuedReservedInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteQueuedReservedInstances = "DeleteQueuedReservedInstances"

// DeleteQueuedReservedInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the queued purchases for the specified Reserved Instances.
//
//    // Example sending a request using DeleteQueuedReservedInstancesRequest.
//    req := client.DeleteQueuedReservedInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteQueuedReservedInstances
func (c *Client) DeleteQueuedReservedInstancesRequest(input *DeleteQueuedReservedInstancesInput) DeleteQueuedReservedInstancesRequest {
	op := &aws.Operation{
		Name:       opDeleteQueuedReservedInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteQueuedReservedInstancesInput{}
	}

	req := c.newRequest(op, input, &DeleteQueuedReservedInstancesOutput{})
	return DeleteQueuedReservedInstancesRequest{Request: req, Input: input, Copy: c.DeleteQueuedReservedInstancesRequest}
}

// DeleteQueuedReservedInstancesRequest is the request type for the
// DeleteQueuedReservedInstances API operation.
type DeleteQueuedReservedInstancesRequest struct {
	*aws.Request
	Input *DeleteQueuedReservedInstancesInput
	Copy  func(*DeleteQueuedReservedInstancesInput) DeleteQueuedReservedInstancesRequest
}

// Send marshals and sends the DeleteQueuedReservedInstances API request.
func (r DeleteQueuedReservedInstancesRequest) Send(ctx context.Context) (*DeleteQueuedReservedInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteQueuedReservedInstancesResponse{
		DeleteQueuedReservedInstancesOutput: r.Request.Data.(*DeleteQueuedReservedInstancesOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteQueuedReservedInstancesResponse is the response type for the
// DeleteQueuedReservedInstances API operation.
type DeleteQueuedReservedInstancesResponse struct {
	*DeleteQueuedReservedInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteQueuedReservedInstances request.
func (r *DeleteQueuedReservedInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
