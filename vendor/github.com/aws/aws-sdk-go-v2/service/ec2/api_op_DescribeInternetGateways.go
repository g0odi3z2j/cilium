// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInternetGatewaysInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * attachment.state - The current state of the attachment between the gateway
	//    and the VPC (available). Present only if a VPC is attached.
	//
	//    * attachment.vpc-id - The ID of an attached VPC.
	//
	//    * internet-gateway-id - The ID of the Internet gateway.
	//
	//    * owner-id - The ID of the AWS account that owns the internet gateway.
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
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more internet gateway IDs.
	//
	// Default: Describes all your internet gateways.
	InternetGatewayIds []string `locationName:"internetGatewayId" locationNameList:"item" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInternetGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInternetGatewaysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInternetGatewaysInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInternetGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more internet gateways.
	InternetGateways []InternetGateway `locationName:"internetGatewaySet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInternetGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInternetGateways = "DescribeInternetGateways"

// DescribeInternetGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your internet gateways.
//
//    // Example sending a request using DescribeInternetGatewaysRequest.
//    req := client.DescribeInternetGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInternetGateways
func (c *Client) DescribeInternetGatewaysRequest(input *DescribeInternetGatewaysInput) DescribeInternetGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeInternetGateways,
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
		input = &DescribeInternetGatewaysInput{}
	}

	req := c.newRequest(op, input, &DescribeInternetGatewaysOutput{})

	return DescribeInternetGatewaysRequest{Request: req, Input: input, Copy: c.DescribeInternetGatewaysRequest}
}

// DescribeInternetGatewaysRequest is the request type for the
// DescribeInternetGateways API operation.
type DescribeInternetGatewaysRequest struct {
	*aws.Request
	Input *DescribeInternetGatewaysInput
	Copy  func(*DescribeInternetGatewaysInput) DescribeInternetGatewaysRequest
}

// Send marshals and sends the DescribeInternetGateways API request.
func (r DescribeInternetGatewaysRequest) Send(ctx context.Context) (*DescribeInternetGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInternetGatewaysResponse{
		DescribeInternetGatewaysOutput: r.Request.Data.(*DescribeInternetGatewaysOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInternetGatewaysRequestPaginator returns a paginator for DescribeInternetGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInternetGatewaysRequest(input)
//   p := ec2.NewDescribeInternetGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInternetGatewaysPaginator(req DescribeInternetGatewaysRequest) DescribeInternetGatewaysPaginator {
	return DescribeInternetGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInternetGatewaysInput
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

// DescribeInternetGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInternetGatewaysPaginator struct {
	aws.Pager
}

func (p *DescribeInternetGatewaysPaginator) CurrentPage() *DescribeInternetGatewaysOutput {
	return p.Pager.CurrentPage().(*DescribeInternetGatewaysOutput)
}

// DescribeInternetGatewaysResponse is the response type for the
// DescribeInternetGateways API operation.
type DescribeInternetGatewaysResponse struct {
	*DescribeInternetGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInternetGateways request.
func (r *DescribeInternetGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
