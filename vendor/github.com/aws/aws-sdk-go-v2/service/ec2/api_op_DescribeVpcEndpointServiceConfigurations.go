// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcEndpointServiceConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * service-name - The name of the service.
	//
	//    * service-id - The ID of the service.
	//
	//    * service-state - The state of the service (Pending | Available | Deleting
	//    | Deleted | Failed).
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

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. This value can be between 5 and
	// 1,000; if MaxResults is given a value larger than 1,000, only 1,000 results
	// are returned.
	MaxResults *int64 `type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The IDs of one or more services.
	ServiceIds []string `locationName:"ServiceId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpcEndpointServiceConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about one or more services.
	ServiceConfigurations []ServiceConfiguration `locationName:"serviceConfigurationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcEndpointServiceConfigurations = "DescribeVpcEndpointServiceConfigurations"

// DescribeVpcEndpointServiceConfigurationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the VPC endpoint service configurations in your account (your services).
//
//    // Example sending a request using DescribeVpcEndpointServiceConfigurationsRequest.
//    req := client.DescribeVpcEndpointServiceConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcEndpointServiceConfigurations
func (c *Client) DescribeVpcEndpointServiceConfigurationsRequest(input *DescribeVpcEndpointServiceConfigurationsInput) DescribeVpcEndpointServiceConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcEndpointServiceConfigurations,
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
		input = &DescribeVpcEndpointServiceConfigurationsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcEndpointServiceConfigurationsOutput{})

	return DescribeVpcEndpointServiceConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeVpcEndpointServiceConfigurationsRequest}
}

// DescribeVpcEndpointServiceConfigurationsRequest is the request type for the
// DescribeVpcEndpointServiceConfigurations API operation.
type DescribeVpcEndpointServiceConfigurationsRequest struct {
	*aws.Request
	Input *DescribeVpcEndpointServiceConfigurationsInput
	Copy  func(*DescribeVpcEndpointServiceConfigurationsInput) DescribeVpcEndpointServiceConfigurationsRequest
}

// Send marshals and sends the DescribeVpcEndpointServiceConfigurations API request.
func (r DescribeVpcEndpointServiceConfigurationsRequest) Send(ctx context.Context) (*DescribeVpcEndpointServiceConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcEndpointServiceConfigurationsResponse{
		DescribeVpcEndpointServiceConfigurationsOutput: r.Request.Data.(*DescribeVpcEndpointServiceConfigurationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVpcEndpointServiceConfigurationsRequestPaginator returns a paginator for DescribeVpcEndpointServiceConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVpcEndpointServiceConfigurationsRequest(input)
//   p := ec2.NewDescribeVpcEndpointServiceConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVpcEndpointServiceConfigurationsPaginator(req DescribeVpcEndpointServiceConfigurationsRequest) DescribeVpcEndpointServiceConfigurationsPaginator {
	return DescribeVpcEndpointServiceConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVpcEndpointServiceConfigurationsInput
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

// DescribeVpcEndpointServiceConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVpcEndpointServiceConfigurationsPaginator struct {
	aws.Pager
}

func (p *DescribeVpcEndpointServiceConfigurationsPaginator) CurrentPage() *DescribeVpcEndpointServiceConfigurationsOutput {
	return p.Pager.CurrentPage().(*DescribeVpcEndpointServiceConfigurationsOutput)
}

// DescribeVpcEndpointServiceConfigurationsResponse is the response type for the
// DescribeVpcEndpointServiceConfigurations API operation.
type DescribeVpcEndpointServiceConfigurationsResponse struct {
	*DescribeVpcEndpointServiceConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcEndpointServiceConfigurations request.
func (r *DescribeVpcEndpointServiceConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
