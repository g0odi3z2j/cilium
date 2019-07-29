// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePrincipalIdFormatInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log
	// | image | import-task | instance | internet-gateway | network-acl | network-acl-association
	// | network-interface | network-interface-attachment | prefix-list | reservation
	// | route-table | route-table-association | security-group | snapshot | subnet
	// | subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association
	// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway
	Resources []string `locationName:"Resource" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePrincipalIdFormatInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePrincipalIdFormatInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePrincipalIdFormatInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePrincipalIdFormatOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the ID format settings for the ARN.
	Principals []PrincipalIdFormat `locationName:"principalSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePrincipalIdFormatOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePrincipalIdFormat = "DescribePrincipalIdFormat"

// DescribePrincipalIdFormatRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the ID format settings for the root user and all IAM roles and
// IAM users that have explicitly specified a longer ID (17-character ID) preference.
//
// By default, all IAM roles and IAM users default to the same ID settings as
// the root user, unless they explicitly override the settings. This request
// is useful for identifying those IAM users and IAM roles that have overridden
// the default ID settings.
//
// The following resource types support longer IDs: bundle | conversion-task
// | customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | instance | internet-gateway
// | network-acl | network-acl-association | network-interface | network-interface-attachment
// | prefix-list | reservation | route-table | route-table-association | security-group
// | snapshot | subnet | subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association
// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
//
//    // Example sending a request using DescribePrincipalIdFormatRequest.
//    req := client.DescribePrincipalIdFormatRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePrincipalIdFormat
func (c *Client) DescribePrincipalIdFormatRequest(input *DescribePrincipalIdFormatInput) DescribePrincipalIdFormatRequest {
	op := &aws.Operation{
		Name:       opDescribePrincipalIdFormat,
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
		input = &DescribePrincipalIdFormatInput{}
	}

	req := c.newRequest(op, input, &DescribePrincipalIdFormatOutput{})
	return DescribePrincipalIdFormatRequest{Request: req, Input: input, Copy: c.DescribePrincipalIdFormatRequest}
}

// DescribePrincipalIdFormatRequest is the request type for the
// DescribePrincipalIdFormat API operation.
type DescribePrincipalIdFormatRequest struct {
	*aws.Request
	Input *DescribePrincipalIdFormatInput
	Copy  func(*DescribePrincipalIdFormatInput) DescribePrincipalIdFormatRequest
}

// Send marshals and sends the DescribePrincipalIdFormat API request.
func (r DescribePrincipalIdFormatRequest) Send(ctx context.Context) (*DescribePrincipalIdFormatResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePrincipalIdFormatResponse{
		DescribePrincipalIdFormatOutput: r.Request.Data.(*DescribePrincipalIdFormatOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePrincipalIdFormatRequestPaginator returns a paginator for DescribePrincipalIdFormat.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePrincipalIdFormatRequest(input)
//   p := ec2.NewDescribePrincipalIdFormatRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePrincipalIdFormatPaginator(req DescribePrincipalIdFormatRequest) DescribePrincipalIdFormatPaginator {
	return DescribePrincipalIdFormatPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribePrincipalIdFormatInput
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

// DescribePrincipalIdFormatPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePrincipalIdFormatPaginator struct {
	aws.Pager
}

func (p *DescribePrincipalIdFormatPaginator) CurrentPage() *DescribePrincipalIdFormatOutput {
	return p.Pager.CurrentPage().(*DescribePrincipalIdFormatOutput)
}

// DescribePrincipalIdFormatResponse is the response type for the
// DescribePrincipalIdFormat API operation.
type DescribePrincipalIdFormatResponse struct {
	*DescribePrincipalIdFormatOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePrincipalIdFormat request.
func (r *DescribePrincipalIdFormatResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
