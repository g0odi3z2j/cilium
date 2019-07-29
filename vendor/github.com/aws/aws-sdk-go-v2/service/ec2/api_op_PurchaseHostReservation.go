// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PurchaseHostReservationInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// The currency in which the totalUpfrontPrice, LimitPrice, and totalHourlyPrice
	// amounts are specified. At this time, the only supported currency is USD.
	CurrencyCode CurrencyCodeValues `type:"string" enum:"true"`

	// The IDs of the Dedicated Hosts with which the reservation will be associated.
	//
	// HostIdSet is a required field
	HostIdSet []string `locationNameList:"item" type:"list" required:"true"`

	// The specified limit is checked against the total upfront cost of the reservation
	// (calculated as the offering's upfront cost multiplied by the host count).
	// If the total upfront cost is greater than the specified price limit, the
	// request fails. This is used to ensure that the purchase does not exceed the
	// expected upfront cost of the purchase. At this time, the only supported currency
	// is USD. For example, to indicate a limit price of USD 100, specify 100.00.
	LimitPrice *string `type:"string"`

	// The ID of the offering.
	//
	// OfferingId is a required field
	OfferingId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurchaseHostReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseHostReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseHostReservationInput"}

	if s.HostIdSet == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostIdSet"))
	}

	if s.OfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PurchaseHostReservationOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The currency in which the totalUpfrontPrice and totalHourlyPrice amounts
	// are specified. At this time, the only supported currency is USD.
	CurrencyCode CurrencyCodeValues `locationName:"currencyCode" type:"string" enum:"true"`

	// Describes the details of the purchase.
	Purchase []Purchase `locationName:"purchase" locationNameList:"item" type:"list"`

	// The total hourly price of the reservation calculated per hour.
	TotalHourlyPrice *string `locationName:"totalHourlyPrice" type:"string"`

	// The total amount charged to your account when you purchase the reservation.
	TotalUpfrontPrice *string `locationName:"totalUpfrontPrice" type:"string"`
}

// String returns the string representation
func (s PurchaseHostReservationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseHostReservation = "PurchaseHostReservation"

// PurchaseHostReservationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Purchase a reservation with configurations that match those of your Dedicated
// Host. You must have active Dedicated Hosts in your account before you purchase
// a reservation. This action results in the specified reservation being purchased
// and charged to your account.
//
//    // Example sending a request using PurchaseHostReservationRequest.
//    req := client.PurchaseHostReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseHostReservation
func (c *Client) PurchaseHostReservationRequest(input *PurchaseHostReservationInput) PurchaseHostReservationRequest {
	op := &aws.Operation{
		Name:       opPurchaseHostReservation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseHostReservationInput{}
	}

	req := c.newRequest(op, input, &PurchaseHostReservationOutput{})
	return PurchaseHostReservationRequest{Request: req, Input: input, Copy: c.PurchaseHostReservationRequest}
}

// PurchaseHostReservationRequest is the request type for the
// PurchaseHostReservation API operation.
type PurchaseHostReservationRequest struct {
	*aws.Request
	Input *PurchaseHostReservationInput
	Copy  func(*PurchaseHostReservationInput) PurchaseHostReservationRequest
}

// Send marshals and sends the PurchaseHostReservation API request.
func (r PurchaseHostReservationRequest) Send(ctx context.Context) (*PurchaseHostReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseHostReservationResponse{
		PurchaseHostReservationOutput: r.Request.Data.(*PurchaseHostReservationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseHostReservationResponse is the response type for the
// PurchaseHostReservation API operation.
type PurchaseHostReservationResponse struct {
	*PurchaseHostReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseHostReservation request.
func (r *PurchaseHostReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
