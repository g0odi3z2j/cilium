// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified EBS snapshots available to you or all of the EBS
// snapshots available to you. The snapshots available to you include public
// snapshots, private snapshots that you own, and private snapshots owned by other
// AWS accounts for which you have explicit create volume permissions. The create
// volume permissions fall into the following categories:
//
// * public: The owner of
// the snapshot granted create volume permissions for the snapshot to the all
// group. All AWS accounts have create volume permissions for these snapshots.
//
// *
// explicit: The owner of the snapshot granted create volume permissions to a
// specific AWS account.
//
// * implicit: An AWS account has implicit create volume
// permissions for all snapshots it owns.
//
// The list of snapshots returned can be
// filtered by specifying snapshot IDs, snapshot owners, or AWS accounts with
// create volume permissions. If no options are specified, Amazon EC2 returns all
// snapshots for which you have create volume permissions. If you specify one or
// more snapshot IDs, only snapshots that have the specified IDs are returned. If
// you specify an invalid snapshot ID, an error is returned. If you specify a
// snapshot ID for which you do not have access, it is not included in the returned
// results. If you specify one or more snapshot owners using the OwnerIds option,
// only snapshots from the specified owners and for which you have access are
// returned. The results can include the AWS account IDs of the specified owners,
// amazon for snapshots owned by Amazon, or self for snapshots that you own. If you
// specify a list of restorable users, only snapshots with create snapshot
// permissions for those users are returned. You can specify AWS account IDs (if
// you own the snapshots), self for snapshots for which you own or have explicit
// permissions, or all for public snapshots. If you are describing a long list of
// snapshots, we recommend that you paginate the output to make the list more
// manageable. The MaxResults parameter sets the maximum number of results returned
// in a single page. If the list of results exceeds your MaxResults value, then
// that number of results is returned along with a NextToken value that can be
// passed to a subsequent DescribeSnapshots request to retrieve the remaining
// results. To get the state of fast snapshot restores for a snapshot, use
// DescribeFastSnapshotRestores. For more information about EBS snapshots, see
// Amazon EBS Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshots", params, optFns, addOperationDescribeSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The filters.
	//
	// * description - A description of the snapshot.
	//
	// * encrypted -
	// Indicates whether the snapshot is encrypted (true | false)
	//
	// * owner-alias - The
	// owner alias, from an Amazon-maintained list (amazon). This is not the
	// user-configured AWS account alias set using the IAM console. We recommend that
	// you use the related parameter instead of this filter.
	//
	// * owner-id - The AWS
	// account ID of the owner. We recommend that you use the related parameter instead
	// of this filter.
	//
	// * progress - The progress of the snapshot, as a percentage (for
	// example, 80%).
	//
	// * snapshot-id - The snapshot ID.
	//
	// * start-time - The time stamp
	// when the snapshot was initiated.
	//
	// * status - The status of the snapshot (pending
	// | completed | error).
	//
	// * tag: - The key/value combination of a tag assigned to
	// the resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use this filter
	// to find all resources assigned a tag with a specific key, regardless of the tag
	// value.
	//
	// * volume-id - The ID of the volume the snapshot is for.
	//
	// * volume-size -
	// The size of the volume, in GiB.
	Filters []types.Filter

	// The maximum number of snapshot results returned by DescribeSnapshots in
	// paginated output. When this parameter is used, DescribeSnapshots only returns
	// MaxResults results in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeSnapshots request with the returned NextToken value. This value can be
	// between 5 and 1000; if MaxResults is given a value larger than 1000, only 1000
	// results are returned. If this parameter is not used, then DescribeSnapshots
	// returns all results. You cannot specify this parameter and the snapshot IDs
	// parameter in the same request.
	MaxResults int32

	// The NextToken value returned from a previous paginated DescribeSnapshots request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string

	// Scopes the results to snapshots with the specified owners. You can specify a
	// combination of AWS account IDs, self, and amazon.
	OwnerIds []string

	// The IDs of the AWS accounts that can create volumes from the snapshot.
	RestorableByUserIds []string

	// The snapshot IDs. Default: Describes the snapshots for which you have create
	// volume permissions.
	SnapshotIds []string
}

type DescribeSnapshotsOutput struct {

	// The NextToken value to include in a future DescribeSnapshots request. When the
	// results of a DescribeSnapshots request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Information about the snapshots.
	Snapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSnapshots{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeSnapshotsAPIClient is a client that implements the DescribeSnapshots
// operation.
type DescribeSnapshotsAPIClient interface {
	DescribeSnapshots(context.Context, *DescribeSnapshotsInput, ...func(*Options)) (*DescribeSnapshotsOutput, error)
}

var _ DescribeSnapshotsAPIClient = (*Client)(nil)

// DescribeSnapshotsPaginatorOptions is the paginator options for DescribeSnapshots
type DescribeSnapshotsPaginatorOptions struct {
	// The maximum number of snapshot results returned by DescribeSnapshots in
	// paginated output. When this parameter is used, DescribeSnapshots only returns
	// MaxResults results in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeSnapshots request with the returned NextToken value. This value can be
	// between 5 and 1000; if MaxResults is given a value larger than 1000, only 1000
	// results are returned. If this parameter is not used, then DescribeSnapshots
	// returns all results. You cannot specify this parameter and the snapshot IDs
	// parameter in the same request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSnapshotsPaginator is a paginator for DescribeSnapshots
type DescribeSnapshotsPaginator struct {
	options   DescribeSnapshotsPaginatorOptions
	client    DescribeSnapshotsAPIClient
	params    *DescribeSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSnapshotsPaginator returns a new DescribeSnapshotsPaginator
func NewDescribeSnapshotsPaginator(client DescribeSnapshotsAPIClient, params *DescribeSnapshotsInput, optFns ...func(*DescribeSnapshotsPaginatorOptions)) *DescribeSnapshotsPaginator {
	options := DescribeSnapshotsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	return &DescribeSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSnapshots page.
func (p *DescribeSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeSnapshots(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSnapshots",
	}
}
