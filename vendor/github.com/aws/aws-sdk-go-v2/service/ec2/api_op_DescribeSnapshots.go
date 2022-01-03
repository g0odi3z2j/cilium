// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified EBS snapshots available to you or all of the EBS
// snapshots available to you. The snapshots available to you include public
// snapshots, private snapshots that you own, and private snapshots owned by other
// Amazon Web Services accounts for which you have explicit create volume
// permissions. The create volume permissions fall into the following
// categories:
//
// * public: The owner of the snapshot granted create volume
// permissions for the snapshot to the all group. All Amazon Web Services accounts
// have create volume permissions for these snapshots.
//
// * explicit: The owner of
// the snapshot granted create volume permissions to a specific Amazon Web Services
// account.
//
// * implicit: An Amazon Web Services account has implicit create volume
// permissions for all snapshots it owns.
//
// The list of snapshots returned can be
// filtered by specifying snapshot IDs, snapshot owners, or Amazon Web Services
// accounts with create volume permissions. If no options are specified, Amazon EC2
// returns all snapshots for which you have create volume permissions. If you
// specify one or more snapshot IDs, only snapshots that have the specified IDs are
// returned. If you specify an invalid snapshot ID, an error is returned. If you
// specify a snapshot ID for which you do not have access, it is not included in
// the returned results. If you specify one or more snapshot owners using the
// OwnerIds option, only snapshots from the specified owners and for which you have
// access are returned. The results can include the Amazon Web Services account IDs
// of the specified owners, amazon for snapshots owned by Amazon, or self for
// snapshots that you own. If you specify a list of restorable users, only
// snapshots with create snapshot permissions for those users are returned. You can
// specify Amazon Web Services account IDs (if you own the snapshots), self for
// snapshots for which you own or have explicit permissions, or all for public
// snapshots. If you are describing a long list of snapshots, we recommend that you
// paginate the output to make the list more manageable. The MaxResults parameter
// sets the maximum number of results returned in a single page. If the list of
// results exceeds your MaxResults value, then that number of results is returned
// along with a NextToken value that can be passed to a subsequent
// DescribeSnapshots request to retrieve the remaining results. To get the state of
// fast snapshot restores for a snapshot, use DescribeFastSnapshotRestores. For
// more information about EBS snapshots, see Amazon EBS snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshots", params, optFns, c.addOperationDescribeSnapshotsMiddlewares)
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
	DryRun *bool

	// The filters.
	//
	// * description - A description of the snapshot.
	//
	// * encrypted -
	// Indicates whether the snapshot is encrypted (true | false)
	//
	// * owner-alias - The
	// owner alias, from an Amazon-maintained list (amazon). This is not the
	// user-configured Amazon Web Services account alias set using the IAM console. We
	// recommend that you use the related parameter instead of this filter.
	//
	// * owner-id
	// - The Amazon Web Services account ID of the owner. We recommend that you use the
	// related parameter instead of this filter.
	//
	// * progress - The progress of the
	// snapshot, as a percentage (for example, 80%).
	//
	// * snapshot-id - The snapshot
	// ID.
	//
	// * start-time - The time stamp when the snapshot was initiated.
	//
	// * status -
	// The status of the snapshot (pending | completed | error).
	//
	// * storage-tier - The
	// storage tier of the snapshot (archive | standard).
	//
	// * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	// * tag-key - The key of a tag
	// assigned to the resource. Use this filter to find all resources assigned a tag
	// with a specific key, regardless of the tag value.
	//
	// * volume-id - The ID of the
	// volume the snapshot is for.
	//
	// * volume-size - The size of the volume, in GiB.
	Filters []types.Filter

	// The maximum number of snapshot results returned by DescribeSnapshots in
	// paginated output. When this parameter is used, DescribeSnapshots only returns
	// MaxResults results in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeSnapshots request with the returned NextToken value. This value can be
	// between 5 and 1,000; if MaxResults is given a value larger than 1,000, only
	// 1,000 results are returned. If this parameter is not used, then
	// DescribeSnapshots returns all results. You cannot specify this parameter and the
	// snapshot IDs parameter in the same request.
	MaxResults *int32

	// The NextToken value returned from a previous paginated DescribeSnapshots request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string

	// Scopes the results to snapshots with the specified owners. You can specify a
	// combination of Amazon Web Services account IDs, self, and amazon.
	OwnerIds []string

	// The IDs of the Amazon Web Services accounts that can create volumes from the
	// snapshot.
	RestorableByUserIds []string

	// The snapshot IDs. Default: Describes the snapshots for which you have create
	// volume permissions.
	SnapshotIds []string

	noSmithyDocumentSerde
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

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
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
	// between 5 and 1,000; if MaxResults is given a value larger than 1,000, only
	// 1,000 results are returned. If this parameter is not used, then
	// DescribeSnapshots returns all results. You cannot specify this parameter and the
	// snapshot IDs parameter in the same request.
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
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	options := DescribeSnapshotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSnapshots page.
func (p *DescribeSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
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

	result, err := p.client.DescribeSnapshots(ctx, &params, optFns...)
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

// SnapshotCompletedWaiterOptions are waiter options for SnapshotCompletedWaiter
type SnapshotCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SnapshotCompletedWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SnapshotCompletedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeSnapshotsInput, *DescribeSnapshotsOutput, error) (bool, error)
}

// SnapshotCompletedWaiter defines the waiters for SnapshotCompleted
type SnapshotCompletedWaiter struct {
	client DescribeSnapshotsAPIClient

	options SnapshotCompletedWaiterOptions
}

// NewSnapshotCompletedWaiter constructs a SnapshotCompletedWaiter.
func NewSnapshotCompletedWaiter(client DescribeSnapshotsAPIClient, optFns ...func(*SnapshotCompletedWaiterOptions)) *SnapshotCompletedWaiter {
	options := SnapshotCompletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = snapshotCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SnapshotCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SnapshotCompleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *SnapshotCompletedWaiter) Wait(ctx context.Context, params *DescribeSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*SnapshotCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SnapshotCompleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *SnapshotCompletedWaiter) WaitForOutput(ctx context.Context, params *DescribeSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*SnapshotCompletedWaiterOptions)) (*DescribeSnapshotsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeSnapshots(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for SnapshotCompleted waiter")
}

func snapshotCompletedStateRetryable(ctx context.Context, input *DescribeSnapshotsInput, output *DescribeSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Snapshots[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "completed"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.SnapshotState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.SnapshotState value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Snapshots[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "error"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.SnapshotState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.SnapshotState value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSnapshots",
	}
}
