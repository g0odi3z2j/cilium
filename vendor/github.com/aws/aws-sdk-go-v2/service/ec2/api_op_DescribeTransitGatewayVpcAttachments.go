// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTransitGatewayVpcAttachmentsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * state - The state of the attachment (available | deleted | deleting
	//    | failed | modifying | pendingAcceptance | pending | rollingBack | rejected
	//    | rejecting).
	//
	//    * transit-gateway-attachment-id - The ID of the attachment.
	//
	//    * transit-gateway-id - The ID of the transit gateway.
	//
	//    * vpc-id - The ID of the VPC.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The IDs of the attachments.
	TransitGatewayAttachmentIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayVpcAttachmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitGatewayVpcAttachmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTransitGatewayVpcAttachmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTransitGatewayVpcAttachmentsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the VPC attachments.
	TransitGatewayVpcAttachments []TransitGatewayVpcAttachment `locationName:"transitGatewayVpcAttachments" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayVpcAttachmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTransitGatewayVpcAttachments = "DescribeTransitGatewayVpcAttachments"

// DescribeTransitGatewayVpcAttachmentsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more VPC attachments. By default, all VPC attachments are
// described. Alternatively, you can filter the results.
//
//    // Example sending a request using DescribeTransitGatewayVpcAttachmentsRequest.
//    req := client.DescribeTransitGatewayVpcAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGatewayVpcAttachments
func (c *Client) DescribeTransitGatewayVpcAttachmentsRequest(input *DescribeTransitGatewayVpcAttachmentsInput) DescribeTransitGatewayVpcAttachmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeTransitGatewayVpcAttachments,
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
		input = &DescribeTransitGatewayVpcAttachmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeTransitGatewayVpcAttachmentsOutput{})
	return DescribeTransitGatewayVpcAttachmentsRequest{Request: req, Input: input, Copy: c.DescribeTransitGatewayVpcAttachmentsRequest}
}

// DescribeTransitGatewayVpcAttachmentsRequest is the request type for the
// DescribeTransitGatewayVpcAttachments API operation.
type DescribeTransitGatewayVpcAttachmentsRequest struct {
	*aws.Request
	Input *DescribeTransitGatewayVpcAttachmentsInput
	Copy  func(*DescribeTransitGatewayVpcAttachmentsInput) DescribeTransitGatewayVpcAttachmentsRequest
}

// Send marshals and sends the DescribeTransitGatewayVpcAttachments API request.
func (r DescribeTransitGatewayVpcAttachmentsRequest) Send(ctx context.Context) (*DescribeTransitGatewayVpcAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTransitGatewayVpcAttachmentsResponse{
		DescribeTransitGatewayVpcAttachmentsOutput: r.Request.Data.(*DescribeTransitGatewayVpcAttachmentsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTransitGatewayVpcAttachmentsRequestPaginator returns a paginator for DescribeTransitGatewayVpcAttachments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTransitGatewayVpcAttachmentsRequest(input)
//   p := ec2.NewDescribeTransitGatewayVpcAttachmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTransitGatewayVpcAttachmentsPaginator(req DescribeTransitGatewayVpcAttachmentsRequest) DescribeTransitGatewayVpcAttachmentsPaginator {
	return DescribeTransitGatewayVpcAttachmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTransitGatewayVpcAttachmentsInput
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

// DescribeTransitGatewayVpcAttachmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTransitGatewayVpcAttachmentsPaginator struct {
	aws.Pager
}

func (p *DescribeTransitGatewayVpcAttachmentsPaginator) CurrentPage() *DescribeTransitGatewayVpcAttachmentsOutput {
	return p.Pager.CurrentPage().(*DescribeTransitGatewayVpcAttachmentsOutput)
}

// DescribeTransitGatewayVpcAttachmentsResponse is the response type for the
// DescribeTransitGatewayVpcAttachments API operation.
type DescribeTransitGatewayVpcAttachmentsResponse struct {
	*DescribeTransitGatewayVpcAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTransitGatewayVpcAttachments request.
func (r *DescribeTransitGatewayVpcAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
