// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type ReportInstanceStatusInput struct {
	_ struct{} `type:"structure"`

	// Descriptive text about the health state of your instance.
	Description *string `locationName:"description" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The time at which the reported instance health state ended.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The instances.
	//
	// Instances is a required field
	Instances []string `locationName:"instanceId" locationNameList:"InstanceId" type:"list" required:"true"`

	// The reason codes that describe the health state of your instance.
	//
	//    * instance-stuck-in-state: My instance is stuck in a state.
	//
	//    * unresponsive: My instance is unresponsive.
	//
	//    * not-accepting-credentials: My instance is not accepting my credentials.
	//
	//    * password-not-available: A password is not available for my instance.
	//
	//    * performance-network: My instance is experiencing performance problems
	//    that I believe are network related.
	//
	//    * performance-instance-store: My instance is experiencing performance
	//    problems that I believe are related to the instance stores.
	//
	//    * performance-ebs-volume: My instance is experiencing performance problems
	//    that I believe are related to an EBS volume.
	//
	//    * performance-other: My instance is experiencing performance problems.
	//
	//    * other: [explain using the description parameter]
	//
	// ReasonCodes is a required field
	ReasonCodes []ReportInstanceReasonCodes `locationName:"reasonCode" locationNameList:"item" type:"list" required:"true"`

	// The time at which the reported instance health state began.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// The status of all instances listed.
	//
	// Status is a required field
	Status ReportStatusType `locationName:"status" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ReportInstanceStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReportInstanceStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReportInstanceStatusInput"}

	if s.Instances == nil {
		invalidParams.Add(aws.NewErrParamRequired("Instances"))
	}

	if s.ReasonCodes == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReasonCodes"))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReportInstanceStatusOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ReportInstanceStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opReportInstanceStatus = "ReportInstanceStatus"

// ReportInstanceStatusRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Submits feedback about the status of an instance. The instance must be in
// the running state. If your experience with the instance differs from the
// instance status returned by DescribeInstanceStatus, use ReportInstanceStatus
// to report your experience with the instance. Amazon EC2 collects this information
// to improve the accuracy of status checks.
//
// Use of this action does not change the value returned by DescribeInstanceStatus.
//
//    // Example sending a request using ReportInstanceStatusRequest.
//    req := client.ReportInstanceStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReportInstanceStatus
func (c *Client) ReportInstanceStatusRequest(input *ReportInstanceStatusInput) ReportInstanceStatusRequest {
	op := &aws.Operation{
		Name:       opReportInstanceStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReportInstanceStatusInput{}
	}

	req := c.newRequest(op, input, &ReportInstanceStatusOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ReportInstanceStatusRequest{Request: req, Input: input, Copy: c.ReportInstanceStatusRequest}
}

// ReportInstanceStatusRequest is the request type for the
// ReportInstanceStatus API operation.
type ReportInstanceStatusRequest struct {
	*aws.Request
	Input *ReportInstanceStatusInput
	Copy  func(*ReportInstanceStatusInput) ReportInstanceStatusRequest
}

// Send marshals and sends the ReportInstanceStatus API request.
func (r ReportInstanceStatusRequest) Send(ctx context.Context) (*ReportInstanceStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReportInstanceStatusResponse{
		ReportInstanceStatusOutput: r.Request.Data.(*ReportInstanceStatusOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReportInstanceStatusResponse is the response type for the
// ReportInstanceStatus API operation.
type ReportInstanceStatusResponse struct {
	*ReportInstanceStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReportInstanceStatus request.
func (r *ReportInstanceStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
