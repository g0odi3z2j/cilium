// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchLocalGatewayRoutesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	// Filters is a required field
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list" required:"true"`

	// The ID of the local gateway route table.
	//
	// LocalGatewayRouteTableId is a required field
	LocalGatewayRouteTableId *string `type:"string" required:"true"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s SearchLocalGatewayRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchLocalGatewayRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchLocalGatewayRoutesInput"}

	if s.Filters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filters"))
	}

	if s.LocalGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocalGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchLocalGatewayRoutesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the routes.
	Routes []LocalGatewayRoute `locationName:"routeSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s SearchLocalGatewayRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchLocalGatewayRoutes = "SearchLocalGatewayRoutes"

// SearchLocalGatewayRoutesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Searches for routes in the specified local gateway route table.
//
//    // Example sending a request using SearchLocalGatewayRoutesRequest.
//    req := client.SearchLocalGatewayRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/SearchLocalGatewayRoutes
func (c *Client) SearchLocalGatewayRoutesRequest(input *SearchLocalGatewayRoutesInput) SearchLocalGatewayRoutesRequest {
	op := &aws.Operation{
		Name:       opSearchLocalGatewayRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SearchLocalGatewayRoutesInput{}
	}

	req := c.newRequest(op, input, &SearchLocalGatewayRoutesOutput{})
	return SearchLocalGatewayRoutesRequest{Request: req, Input: input, Copy: c.SearchLocalGatewayRoutesRequest}
}

// SearchLocalGatewayRoutesRequest is the request type for the
// SearchLocalGatewayRoutes API operation.
type SearchLocalGatewayRoutesRequest struct {
	*aws.Request
	Input *SearchLocalGatewayRoutesInput
	Copy  func(*SearchLocalGatewayRoutesInput) SearchLocalGatewayRoutesRequest
}

// Send marshals and sends the SearchLocalGatewayRoutes API request.
func (r SearchLocalGatewayRoutesRequest) Send(ctx context.Context) (*SearchLocalGatewayRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchLocalGatewayRoutesResponse{
		SearchLocalGatewayRoutesOutput: r.Request.Data.(*SearchLocalGatewayRoutesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchLocalGatewayRoutesResponse is the response type for the
// SearchLocalGatewayRoutes API operation.
type SearchLocalGatewayRoutesResponse struct {
	*SearchLocalGatewayRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchLocalGatewayRoutes request.
func (r *SearchLocalGatewayRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
