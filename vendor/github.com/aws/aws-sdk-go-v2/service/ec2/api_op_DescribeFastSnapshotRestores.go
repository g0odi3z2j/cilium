// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFastSnapshotRestoresInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters. The possible values are:
	//
	//    * availability-zone: The Availability Zone of the snapshot.
	//
	//    * owner-id: The ID of the AWS account that owns the snapshot.
	//
	//    * snapshot-id: The ID of the snapshot.
	//
	//    * state: The state of fast snapshot restores for the snapshot (enabling
	//    | optimizing | enabled | disabling | disabled).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFastSnapshotRestoresInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeFastSnapshotRestoresOutput struct {
	_ struct{} `type:"structure"`

	// Information about the state of fast snapshot restores.
	FastSnapshotRestores []DescribeFastSnapshotRestoreSuccessItem `locationName:"fastSnapshotRestoreSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFastSnapshotRestoresOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFastSnapshotRestores = "DescribeFastSnapshotRestores"

// DescribeFastSnapshotRestoresRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the state of fast snapshot restores for your snapshots.
//
//    // Example sending a request using DescribeFastSnapshotRestoresRequest.
//    req := client.DescribeFastSnapshotRestoresRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFastSnapshotRestores
func (c *Client) DescribeFastSnapshotRestoresRequest(input *DescribeFastSnapshotRestoresInput) DescribeFastSnapshotRestoresRequest {
	op := &aws.Operation{
		Name:       opDescribeFastSnapshotRestores,
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
		input = &DescribeFastSnapshotRestoresInput{}
	}

	req := c.newRequest(op, input, &DescribeFastSnapshotRestoresOutput{})
	return DescribeFastSnapshotRestoresRequest{Request: req, Input: input, Copy: c.DescribeFastSnapshotRestoresRequest}
}

// DescribeFastSnapshotRestoresRequest is the request type for the
// DescribeFastSnapshotRestores API operation.
type DescribeFastSnapshotRestoresRequest struct {
	*aws.Request
	Input *DescribeFastSnapshotRestoresInput
	Copy  func(*DescribeFastSnapshotRestoresInput) DescribeFastSnapshotRestoresRequest
}

// Send marshals and sends the DescribeFastSnapshotRestores API request.
func (r DescribeFastSnapshotRestoresRequest) Send(ctx context.Context) (*DescribeFastSnapshotRestoresResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFastSnapshotRestoresResponse{
		DescribeFastSnapshotRestoresOutput: r.Request.Data.(*DescribeFastSnapshotRestoresOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFastSnapshotRestoresRequestPaginator returns a paginator for DescribeFastSnapshotRestores.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFastSnapshotRestoresRequest(input)
//   p := ec2.NewDescribeFastSnapshotRestoresRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFastSnapshotRestoresPaginator(req DescribeFastSnapshotRestoresRequest) DescribeFastSnapshotRestoresPaginator {
	return DescribeFastSnapshotRestoresPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFastSnapshotRestoresInput
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

// DescribeFastSnapshotRestoresPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFastSnapshotRestoresPaginator struct {
	aws.Pager
}

func (p *DescribeFastSnapshotRestoresPaginator) CurrentPage() *DescribeFastSnapshotRestoresOutput {
	return p.Pager.CurrentPage().(*DescribeFastSnapshotRestoresOutput)
}

// DescribeFastSnapshotRestoresResponse is the response type for the
// DescribeFastSnapshotRestores API operation.
type DescribeFastSnapshotRestoresResponse struct {
	*DescribeFastSnapshotRestoresOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFastSnapshotRestores request.
func (r *DescribeFastSnapshotRestoresResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
