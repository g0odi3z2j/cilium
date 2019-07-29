// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeImportSnapshotTasksInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// A list of import snapshot task IDs.
	ImportTaskIds []string `locationName:"ImportTaskId" locationNameList:"ImportTaskId" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// A token that indicates the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeImportSnapshotTasksInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeImportSnapshotTasksOutput struct {
	_ struct{} `type:"structure"`

	// A list of zero or more import snapshot tasks that are currently active or
	// were completed or canceled in the previous 7 days.
	ImportSnapshotTasks []ImportSnapshotTask `locationName:"importSnapshotTaskSet" locationNameList:"item" type:"list"`

	// The token to use to get the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeImportSnapshotTasksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeImportSnapshotTasks = "DescribeImportSnapshotTasks"

// DescribeImportSnapshotTasksRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes your import snapshot tasks.
//
//    // Example sending a request using DescribeImportSnapshotTasksRequest.
//    req := client.DescribeImportSnapshotTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeImportSnapshotTasks
func (c *Client) DescribeImportSnapshotTasksRequest(input *DescribeImportSnapshotTasksInput) DescribeImportSnapshotTasksRequest {
	op := &aws.Operation{
		Name:       opDescribeImportSnapshotTasks,
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
		input = &DescribeImportSnapshotTasksInput{}
	}

	req := c.newRequest(op, input, &DescribeImportSnapshotTasksOutput{})
	return DescribeImportSnapshotTasksRequest{Request: req, Input: input, Copy: c.DescribeImportSnapshotTasksRequest}
}

// DescribeImportSnapshotTasksRequest is the request type for the
// DescribeImportSnapshotTasks API operation.
type DescribeImportSnapshotTasksRequest struct {
	*aws.Request
	Input *DescribeImportSnapshotTasksInput
	Copy  func(*DescribeImportSnapshotTasksInput) DescribeImportSnapshotTasksRequest
}

// Send marshals and sends the DescribeImportSnapshotTasks API request.
func (r DescribeImportSnapshotTasksRequest) Send(ctx context.Context) (*DescribeImportSnapshotTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImportSnapshotTasksResponse{
		DescribeImportSnapshotTasksOutput: r.Request.Data.(*DescribeImportSnapshotTasksOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeImportSnapshotTasksRequestPaginator returns a paginator for DescribeImportSnapshotTasks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeImportSnapshotTasksRequest(input)
//   p := ec2.NewDescribeImportSnapshotTasksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeImportSnapshotTasksPaginator(req DescribeImportSnapshotTasksRequest) DescribeImportSnapshotTasksPaginator {
	return DescribeImportSnapshotTasksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeImportSnapshotTasksInput
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

// DescribeImportSnapshotTasksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeImportSnapshotTasksPaginator struct {
	aws.Pager
}

func (p *DescribeImportSnapshotTasksPaginator) CurrentPage() *DescribeImportSnapshotTasksOutput {
	return p.Pager.CurrentPage().(*DescribeImportSnapshotTasksOutput)
}

// DescribeImportSnapshotTasksResponse is the response type for the
// DescribeImportSnapshotTasks API operation.
type DescribeImportSnapshotTasksResponse struct {
	*DescribeImportSnapshotTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImportSnapshotTasks request.
func (r *DescribeImportSnapshotTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
