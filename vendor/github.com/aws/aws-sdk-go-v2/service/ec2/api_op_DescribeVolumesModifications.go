// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVolumesModificationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters. Supported filters: volume-id, modification-state, target-size,
	// target-iops, target-volume-type, original-size, original-iops, original-volume-type,
	// start-time.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results (up to a limit of 500) to be returned in a
	// paginated request.
	MaxResults *int64 `type:"integer"`

	// The nextToken value returned by a previous paginated request.
	NextToken *string `type:"string"`

	// The IDs of the volumes for which in-progress modifications will be described.
	VolumeIds []string `locationName:"VolumeId" locationNameList:"VolumeId" type:"list"`
}

// String returns the string representation
func (s DescribeVolumesModificationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVolumesModificationsOutput struct {
	_ struct{} `type:"structure"`

	// Token for pagination, null if there are no more results
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the volume modifications.
	VolumesModifications []VolumeModification `locationName:"volumeModificationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVolumesModificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVolumesModifications = "DescribeVolumesModifications"

// DescribeVolumesModificationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Reports the current modification status of EBS volumes.
//
// Current-generation EBS volumes support modification of attributes including
// type, size, and (for io1 volumes) IOPS provisioning while either attached
// to or detached from an instance. Following an action from the API or the
// console to modify a volume, the status of the modification may be modifying,
// optimizing, completed, or failed. If a volume has never been modified, then
// certain elements of the returned VolumeModification objects are null.
//
// You can also use CloudWatch Events to check the status of a modification
// to an EBS volume. For information about CloudWatch Events, see the Amazon
// CloudWatch Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/).
// For more information, see Monitoring Volume Modifications" (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#monitoring_mods)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeVolumesModificationsRequest.
//    req := client.DescribeVolumesModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumesModifications
func (c *Client) DescribeVolumesModificationsRequest(input *DescribeVolumesModificationsInput) DescribeVolumesModificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumesModifications,
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
		input = &DescribeVolumesModificationsInput{}
	}

	req := c.newRequest(op, input, &DescribeVolumesModificationsOutput{})
	return DescribeVolumesModificationsRequest{Request: req, Input: input, Copy: c.DescribeVolumesModificationsRequest}
}

// DescribeVolumesModificationsRequest is the request type for the
// DescribeVolumesModifications API operation.
type DescribeVolumesModificationsRequest struct {
	*aws.Request
	Input *DescribeVolumesModificationsInput
	Copy  func(*DescribeVolumesModificationsInput) DescribeVolumesModificationsRequest
}

// Send marshals and sends the DescribeVolumesModifications API request.
func (r DescribeVolumesModificationsRequest) Send(ctx context.Context) (*DescribeVolumesModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumesModificationsResponse{
		DescribeVolumesModificationsOutput: r.Request.Data.(*DescribeVolumesModificationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVolumesModificationsRequestPaginator returns a paginator for DescribeVolumesModifications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVolumesModificationsRequest(input)
//   p := ec2.NewDescribeVolumesModificationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVolumesModificationsPaginator(req DescribeVolumesModificationsRequest) DescribeVolumesModificationsPaginator {
	return DescribeVolumesModificationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVolumesModificationsInput
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

// DescribeVolumesModificationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVolumesModificationsPaginator struct {
	aws.Pager
}

func (p *DescribeVolumesModificationsPaginator) CurrentPage() *DescribeVolumesModificationsOutput {
	return p.Pager.CurrentPage().(*DescribeVolumesModificationsOutput)
}

// DescribeVolumesModificationsResponse is the response type for the
// DescribeVolumesModifications API operation.
type DescribeVolumesModificationsResponse struct {
	*DescribeVolumesModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumesModifications request.
func (r *DescribeVolumesModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
