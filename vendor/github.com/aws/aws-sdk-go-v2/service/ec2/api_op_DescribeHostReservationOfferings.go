// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHostReservationOfferingsInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	//
	//    * instance-family - The instance family of the offering (for example,
	//    m4).
	//
	//    * payment-option - The payment option (NoUpfront | PartialUpfront | AllUpfront).
	Filter []Filter `locationNameList:"Filter" type:"list"`

	// This is the maximum duration of the reservation to purchase, specified in
	// seconds. Reservations are available in one-year and three-year terms. The
	// number of seconds specified must be the number of seconds in a year (365x24x60x60)
	// times one of the supported durations (1 or 3). For example, specify 94608000
	// for three years.
	MaxDuration *int64 `type:"integer"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given
	// a larger value than 500, you receive an error.
	MaxResults *int64 `min:"5" type:"integer"`

	// This is the minimum duration of the reservation you'd like to purchase, specified
	// in seconds. Reservations are available in one-year and three-year terms.
	// The number of seconds specified must be the number of seconds in a year (365x24x60x60)
	// times one of the supported durations (1 or 3). For example, specify 31536000
	// for one year.
	MinDuration *int64 `type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The ID of the reservation offering.
	OfferingId *string `type:"string"`
}

// String returns the string representation
func (s DescribeHostReservationOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHostReservationOfferingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeHostReservationOfferingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeHostReservationOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the offerings.
	OfferingSet []HostOffering `locationName:"offeringSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeHostReservationOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHostReservationOfferings = "DescribeHostReservationOfferings"

// DescribeHostReservationOfferingsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Dedicated Host reservations that are available to purchase.
//
// The results describe all of the Dedicated Host reservation offerings, including
// offerings that might not match the instance family and Region of your Dedicated
// Hosts. When purchasing an offering, ensure that the instance family and Region
// of the offering matches that of the Dedicated Hosts with which it is to be
// associated. For more information about supported instance types, see Dedicated
// Hosts Overview (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeHostReservationOfferingsRequest.
//    req := client.DescribeHostReservationOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeHostReservationOfferings
func (c *Client) DescribeHostReservationOfferingsRequest(input *DescribeHostReservationOfferingsInput) DescribeHostReservationOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeHostReservationOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeHostReservationOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeHostReservationOfferingsOutput{})
	return DescribeHostReservationOfferingsRequest{Request: req, Input: input, Copy: c.DescribeHostReservationOfferingsRequest}
}

// DescribeHostReservationOfferingsRequest is the request type for the
// DescribeHostReservationOfferings API operation.
type DescribeHostReservationOfferingsRequest struct {
	*aws.Request
	Input *DescribeHostReservationOfferingsInput
	Copy  func(*DescribeHostReservationOfferingsInput) DescribeHostReservationOfferingsRequest
}

// Send marshals and sends the DescribeHostReservationOfferings API request.
func (r DescribeHostReservationOfferingsRequest) Send(ctx context.Context) (*DescribeHostReservationOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHostReservationOfferingsResponse{
		DescribeHostReservationOfferingsOutput: r.Request.Data.(*DescribeHostReservationOfferingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeHostReservationOfferingsRequestPaginator returns a paginator for DescribeHostReservationOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeHostReservationOfferingsRequest(input)
//   p := ec2.NewDescribeHostReservationOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeHostReservationOfferingsPaginator(req DescribeHostReservationOfferingsRequest) DescribeHostReservationOfferingsPaginator {
	return DescribeHostReservationOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeHostReservationOfferingsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeHostReservationOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeHostReservationOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeHostReservationOfferingsPaginator) CurrentPage() *DescribeHostReservationOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeHostReservationOfferingsOutput)
}

// DescribeHostReservationOfferingsResponse is the response type for the
// DescribeHostReservationOfferings API operation.
type DescribeHostReservationOfferingsResponse struct {
	*DescribeHostReservationOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHostReservationOfferings request.
func (r *DescribeHostReservationOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
