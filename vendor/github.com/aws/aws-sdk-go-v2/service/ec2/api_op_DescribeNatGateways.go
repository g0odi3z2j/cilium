// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeNatGatewaysInput struct {
	_ struct{} `type:"structure"`

	// One or more filters.
	//
	//    * nat-gateway-id - The ID of the NAT gateway.
	//
	//    * state - The state of the NAT gateway (pending | failed | available |
	//    deleting | deleted).
	//
	//    * subnet-id - The ID of the subnet in which the NAT gateway resides.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-id - The ID of the VPC in which the NAT gateway resides.
	Filter []Filter `locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// One or more NAT gateway IDs.
	NatGatewayIds []string `locationName:"NatGatewayId" locationNameList:"item" type:"list"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeNatGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNatGatewaysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNatGatewaysInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeNatGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// Information about the NAT gateways.
	NatGateways []NatGateway `locationName:"natGatewaySet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeNatGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNatGateways = "DescribeNatGateways"

// DescribeNatGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your NAT gateways.
//
//    // Example sending a request using DescribeNatGatewaysRequest.
//    req := client.DescribeNatGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNatGateways
func (c *Client) DescribeNatGatewaysRequest(input *DescribeNatGatewaysInput) DescribeNatGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeNatGateways,
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
		input = &DescribeNatGatewaysInput{}
	}

	req := c.newRequest(op, input, &DescribeNatGatewaysOutput{})
	return DescribeNatGatewaysRequest{Request: req, Input: input, Copy: c.DescribeNatGatewaysRequest}
}

// DescribeNatGatewaysRequest is the request type for the
// DescribeNatGateways API operation.
type DescribeNatGatewaysRequest struct {
	*aws.Request
	Input *DescribeNatGatewaysInput
	Copy  func(*DescribeNatGatewaysInput) DescribeNatGatewaysRequest
}

// Send marshals and sends the DescribeNatGateways API request.
func (r DescribeNatGatewaysRequest) Send(ctx context.Context) (*DescribeNatGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNatGatewaysResponse{
		DescribeNatGatewaysOutput: r.Request.Data.(*DescribeNatGatewaysOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNatGatewaysRequestPaginator returns a paginator for DescribeNatGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNatGatewaysRequest(input)
//   p := ec2.NewDescribeNatGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNatGatewaysPaginator(req DescribeNatGatewaysRequest) DescribeNatGatewaysPaginator {
	return DescribeNatGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeNatGatewaysInput
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

// DescribeNatGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNatGatewaysPaginator struct {
	aws.Pager
}

func (p *DescribeNatGatewaysPaginator) CurrentPage() *DescribeNatGatewaysOutput {
	return p.Pager.CurrentPage().(*DescribeNatGatewaysOutput)
}

// DescribeNatGatewaysResponse is the response type for the
// DescribeNatGateways API operation.
type DescribeNatGatewaysResponse struct {
	*DescribeNatGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNatGateways request.
func (r *DescribeNatGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
