// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the status of the specified volumes. Volume status provides the
// result of the checks performed on your volumes to determine events that can
// impair the performance of your volumes. The performance of a volume can be
// affected if an issue occurs on the volume's underlying host. If the volume's
// underlying host experiences a power outage or system issue, after the system is
// restored, there could be data inconsistencies on the volume. Volume events
// notify you if this occurs. Volume actions notify you if any action needs to be
// taken in response to the event. The DescribeVolumeStatus operation provides the
// following information about the specified volumes: Status: Reflects the current
// status of the volume. The possible values are ok , impaired , warning , or
// insufficient-data . If all checks pass, the overall status of the volume is ok .
// If the check fails, the overall status is impaired . If the status is
// insufficient-data , then the checks might still be taking place on your volume
// at the time. We recommend that you retry the request. For more information about
// volume status, see Monitor the status of your volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-status.html)
// in the Amazon Elastic Compute Cloud User Guide. Events: Reflect the cause of a
// volume status and might require you to take action. For example, if your volume
// returns an impaired status, then the volume event might be
// potential-data-inconsistency . This means that your volume has been affected by
// an issue with the underlying host, has all I/O operations disabled, and might
// have inconsistent data. Actions: Reflect the actions you might have to take in
// response to an event. For example, if the status of the volume is impaired and
// the volume event shows potential-data-inconsistency , then the action shows
// enable-volume-io . This means that you may want to enable the I/O operations for
// the volume by calling the EnableVolumeIO action and then check the volume for
// data consistency. Volume status is based on the volume status checks, and does
// not reflect the volume state. Therefore, volume status does not indicate volumes
// in the error state (for example, when a volume is incapable of accepting I/O.)
func (c *Client) DescribeVolumeStatus(ctx context.Context, params *DescribeVolumeStatusInput, optFns ...func(*Options)) (*DescribeVolumeStatusOutput, error) {
	if params == nil {
		params = &DescribeVolumeStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumeStatus", params, optFns, c.addOperationDescribeVolumeStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumeStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumeStatusInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - action.code - The action code for the event (for example, enable-volume-io
	//   ).
	//   - action.description - A description of the action.
	//   - action.event-id - The event ID associated with the action.
	//   - availability-zone - The Availability Zone of the instance.
	//   - event.description - A description of the event.
	//   - event.event-id - The event ID.
	//   - event.event-type - The event type (for io-enabled : passed | failed ; for
	//   io-performance : io-performance:degraded | io-performance:severely-degraded |
	//   io-performance:stalled ).
	//   - event.not-after - The latest end time for the event.
	//   - event.not-before - The earliest start time for the event.
	//   - volume-status.details-name - The cause for volume-status.status ( io-enabled
	//   | io-performance ).
	//   - volume-status.details-status - The status of volume-status.details-name (for
	//   io-enabled : passed | failed ; for io-performance : normal | degraded |
	//   severely-degraded | stalled ).
	//   - volume-status.status - The status of the volume ( ok | impaired | warning |
	//   insufficient-data ).
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. This value
	// can be between 5 and 1,000; if the value is larger than 1,000, only 1,000
	// results are returned. If this parameter is not used, then all items are
	// returned. You cannot specify this parameter and the volume IDs parameter in the
	// same request. For more information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// The IDs of the volumes. Default: Describes all your volumes.
	VolumeIds []string

	noSmithyDocumentSerde
}

type DescribeVolumeStatusOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the status of the volumes.
	VolumeStatuses []types.VolumeStatusItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVolumeStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumeStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeVolumeStatusAPIClient is a client that implements the
// DescribeVolumeStatus operation.
type DescribeVolumeStatusAPIClient interface {
	DescribeVolumeStatus(context.Context, *DescribeVolumeStatusInput, ...func(*Options)) (*DescribeVolumeStatusOutput, error)
}

var _ DescribeVolumeStatusAPIClient = (*Client)(nil)

// DescribeVolumeStatusPaginatorOptions is the paginator options for
// DescribeVolumeStatus
type DescribeVolumeStatusPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. This value
	// can be between 5 and 1,000; if the value is larger than 1,000, only 1,000
	// results are returned. If this parameter is not used, then all items are
	// returned. You cannot specify this parameter and the volume IDs parameter in the
	// same request. For more information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVolumeStatusPaginator is a paginator for DescribeVolumeStatus
type DescribeVolumeStatusPaginator struct {
	options   DescribeVolumeStatusPaginatorOptions
	client    DescribeVolumeStatusAPIClient
	params    *DescribeVolumeStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeVolumeStatusPaginator returns a new DescribeVolumeStatusPaginator
func NewDescribeVolumeStatusPaginator(client DescribeVolumeStatusAPIClient, params *DescribeVolumeStatusInput, optFns ...func(*DescribeVolumeStatusPaginatorOptions)) *DescribeVolumeStatusPaginator {
	if params == nil {
		params = &DescribeVolumeStatusInput{}
	}

	options := DescribeVolumeStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVolumeStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVolumeStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVolumeStatus page.
func (p *DescribeVolumeStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVolumeStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeVolumeStatus(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeVolumeStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumeStatus",
	}
}
