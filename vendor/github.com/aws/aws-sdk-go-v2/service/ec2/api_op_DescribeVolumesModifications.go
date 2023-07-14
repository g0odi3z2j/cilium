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

// Describes the most recent volume modification request for the specified EBS
// volumes. If a volume has never been modified, some information in the output
// will be null. If a volume has been modified more than once, the output includes
// only the most recent modification request. You can also use CloudWatch Events to
// check the status of a modification to an EBS volume. For information about
// CloudWatch Events, see the Amazon CloudWatch Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/)
// . For more information, see Monitor the progress of volume modifications (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-modifications.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumesModifications(ctx context.Context, params *DescribeVolumesModificationsInput, optFns ...func(*Options)) (*DescribeVolumesModificationsOutput, error) {
	if params == nil {
		params = &DescribeVolumesModificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumesModifications", params, optFns, c.addOperationDescribeVolumesModificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumesModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesModificationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - modification-state - The current modification state (modifying | optimizing
	//   | completed | failed).
	//   - original-iops - The original IOPS rate of the volume.
	//   - original-size - The original size of the volume, in GiB.
	//   - original-volume-type - The original volume type of the volume (standard |
	//   io1 | io2 | gp2 | sc1 | st1).
	//   - originalMultiAttachEnabled - Indicates whether Multi-Attach support was
	//   enabled (true | false).
	//   - start-time - The modification start time.
	//   - target-iops - The target IOPS rate of the volume.
	//   - target-size - The target size of the volume, in GiB.
	//   - target-volume-type - The target volume type of the volume (standard | io1 |
	//   io2 | gp2 | sc1 | st1).
	//   - targetMultiAttachEnabled - Indicates whether Multi-Attach support is to be
	//   enabled (true | false).
	//   - volume-id - The ID of the volume.
	Filters []types.Filter

	// The maximum number of results (up to a limit of 500) to be returned in a
	// paginated request. For more information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned by a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// The IDs of the volumes.
	VolumeIds []string

	noSmithyDocumentSerde
}

type DescribeVolumesModificationsOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null if there are no more items to return.
	NextToken *string

	// Information about the volume modifications.
	VolumesModifications []types.VolumeModification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVolumesModificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumesModifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumesModifications{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumesModifications(options.Region), middleware.Before); err != nil {
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

// DescribeVolumesModificationsAPIClient is a client that implements the
// DescribeVolumesModifications operation.
type DescribeVolumesModificationsAPIClient interface {
	DescribeVolumesModifications(context.Context, *DescribeVolumesModificationsInput, ...func(*Options)) (*DescribeVolumesModificationsOutput, error)
}

var _ DescribeVolumesModificationsAPIClient = (*Client)(nil)

// DescribeVolumesModificationsPaginatorOptions is the paginator options for
// DescribeVolumesModifications
type DescribeVolumesModificationsPaginatorOptions struct {
	// The maximum number of results (up to a limit of 500) to be returned in a
	// paginated request. For more information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVolumesModificationsPaginator is a paginator for
// DescribeVolumesModifications
type DescribeVolumesModificationsPaginator struct {
	options   DescribeVolumesModificationsPaginatorOptions
	client    DescribeVolumesModificationsAPIClient
	params    *DescribeVolumesModificationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVolumesModificationsPaginator returns a new
// DescribeVolumesModificationsPaginator
func NewDescribeVolumesModificationsPaginator(client DescribeVolumesModificationsAPIClient, params *DescribeVolumesModificationsInput, optFns ...func(*DescribeVolumesModificationsPaginatorOptions)) *DescribeVolumesModificationsPaginator {
	if params == nil {
		params = &DescribeVolumesModificationsInput{}
	}

	options := DescribeVolumesModificationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVolumesModificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVolumesModificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVolumesModifications page.
func (p *DescribeVolumesModificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVolumesModificationsOutput, error) {
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

	result, err := p.client.DescribeVolumesModifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVolumesModifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumesModifications",
	}
}
