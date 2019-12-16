// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for PurchaseReservedInstancesOffering.
type PurchaseReservedInstancesOfferingInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The number of Reserved Instances to purchase.
	//
	// InstanceCount is a required field
	InstanceCount *int64 `type:"integer" required:"true"`

	// Specified for Reserved Instance Marketplace offerings to limit the total
	// order and ensure that the Reserved Instances are not purchased at unexpected
	// prices.
	LimitPrice *ReservedInstanceLimitPrice `locationName:"limitPrice" type:"structure"`

	// The time at which to purchase the Reserved Instance, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	PurchaseTime *time.Time `type:"timestamp"`

	// The ID of the Reserved Instance offering to purchase.
	//
	// ReservedInstancesOfferingId is a required field
	ReservedInstancesOfferingId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurchaseReservedInstancesOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseReservedInstancesOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseReservedInstancesOfferingInput"}

	if s.InstanceCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceCount"))
	}

	if s.ReservedInstancesOfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstancesOfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of PurchaseReservedInstancesOffering.
type PurchaseReservedInstancesOfferingOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the purchased Reserved Instances.
	ReservedInstancesId *string `locationName:"reservedInstancesId" type:"string"`
}

// String returns the string representation
func (s PurchaseReservedInstancesOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseReservedInstancesOffering = "PurchaseReservedInstancesOffering"

// PurchaseReservedInstancesOfferingRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Purchases a Reserved Instance for use with your account. With Reserved Instances,
// you pay a lower hourly rate compared to On-Demand instance pricing.
//
// Use DescribeReservedInstancesOfferings to get a list of Reserved Instance
// offerings that match your specifications. After you've purchased a Reserved
// Instance, you can check for your new Reserved Instance with DescribeReservedInstances.
//
// To queue a purchase for a future date and time, specify a purchase time.
// If you do not specify a purchase time, the default is the current time.
//
// For more information, see Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
// and Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using PurchaseReservedInstancesOfferingRequest.
//    req := client.PurchaseReservedInstancesOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseReservedInstancesOffering
func (c *Client) PurchaseReservedInstancesOfferingRequest(input *PurchaseReservedInstancesOfferingInput) PurchaseReservedInstancesOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseReservedInstancesOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseReservedInstancesOfferingInput{}
	}

	req := c.newRequest(op, input, &PurchaseReservedInstancesOfferingOutput{})
	return PurchaseReservedInstancesOfferingRequest{Request: req, Input: input, Copy: c.PurchaseReservedInstancesOfferingRequest}
}

// PurchaseReservedInstancesOfferingRequest is the request type for the
// PurchaseReservedInstancesOffering API operation.
type PurchaseReservedInstancesOfferingRequest struct {
	*aws.Request
	Input *PurchaseReservedInstancesOfferingInput
	Copy  func(*PurchaseReservedInstancesOfferingInput) PurchaseReservedInstancesOfferingRequest
}

// Send marshals and sends the PurchaseReservedInstancesOffering API request.
func (r PurchaseReservedInstancesOfferingRequest) Send(ctx context.Context) (*PurchaseReservedInstancesOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseReservedInstancesOfferingResponse{
		PurchaseReservedInstancesOfferingOutput: r.Request.Data.(*PurchaseReservedInstancesOfferingOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseReservedInstancesOfferingResponse is the response type for the
// PurchaseReservedInstancesOffering API operation.
type PurchaseReservedInstancesOfferingResponse struct {
	*PurchaseReservedInstancesOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseReservedInstancesOffering request.
func (r *PurchaseReservedInstancesOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
