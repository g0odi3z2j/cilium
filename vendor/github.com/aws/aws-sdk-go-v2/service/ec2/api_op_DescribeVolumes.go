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

// Describes the specified EBS volumes or all of your EBS volumes. If you are
// describing a long list of volumes, we recommend that you paginate the output to
// make the list more manageable. The MaxResults parameter sets the maximum number
// of results returned in a single page. If the list of results exceeds your
// MaxResults value, then that number of results is returned along with a NextToken
// value that can be passed to a subsequent DescribeVolumes request to retrieve the
// remaining results. For more information about EBS volumes, see Amazon EBS
// volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumes(ctx context.Context, params *DescribeVolumesInput, optFns ...func(*Options)) (*DescribeVolumesOutput, error) {
	if params == nil {
		params = &DescribeVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumes", params, optFns, addOperationDescribeVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	// * attachment.attach-time - The time stamp when the attachment
	// initiated.
	//
	// * attachment.delete-on-termination - Whether the volume is deleted
	// on instance termination.
	//
	// * attachment.device - The device name specified in the
	// block device mapping (for example, /dev/sda1).
	//
	// * attachment.instance-id - The
	// ID of the instance the volume is attached to.
	//
	// * attachment.status - The
	// attachment state (attaching | attached | detaching).
	//
	// * availability-zone - The
	// Availability Zone in which the volume was created.
	//
	// * create-time - The time
	// stamp when the volume was created.
	//
	// * encrypted - Indicates whether the volume
	// is encrypted (true | false)
	//
	// * multi-attach-enabled - Indicates whether the
	// volume is enabled for Multi-Attach (true | false)
	//
	// * fast-restored - Indicates
	// whether the volume was created from a snapshot that is enabled for fast snapshot
	// restore (true | false).
	//
	// * size - The size of the volume, in GiB.
	//
	// * snapshot-id
	// - The snapshot from which the volume was created.
	//
	// * status - The state of the
	// volume (creating | available | in-use | deleting | deleted | error).
	//
	// * tag: -
	// The key/value combination of a tag assigned to the resource. Use the tag key in
	// the filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key - The
	// key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	//
	// * volume-id -
	// The volume ID.
	//
	// * volume-type - The Amazon EBS volume type (gp2 | gp3 | io1 |
	// io2 | st1 | sc1| standard)
	Filters []types.Filter

	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and 500;
	// if MaxResults is given a value larger than 500, only 500 results are returned.
	// If this parameter is not used, then DescribeVolumes returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int32

	// The NextToken value returned from a previous paginated DescribeVolumes request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string

	// The volume IDs.
	VolumeIds []string
}

type DescribeVolumesOutput struct {

	// The NextToken value to include in a future DescribeVolumes request. When the
	// results of a DescribeVolumes request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Information about the volumes.
	Volumes []types.Volume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumes(options.Region), middleware.Before); err != nil {
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

// DescribeVolumesAPIClient is a client that implements the DescribeVolumes
// operation.
type DescribeVolumesAPIClient interface {
	DescribeVolumes(context.Context, *DescribeVolumesInput, ...func(*Options)) (*DescribeVolumesOutput, error)
}

var _ DescribeVolumesAPIClient = (*Client)(nil)

// DescribeVolumesPaginatorOptions is the paginator options for DescribeVolumes
type DescribeVolumesPaginatorOptions struct {
	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and 500;
	// if MaxResults is given a value larger than 500, only 500 results are returned.
	// If this parameter is not used, then DescribeVolumes returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVolumesPaginator is a paginator for DescribeVolumes
type DescribeVolumesPaginator struct {
	options   DescribeVolumesPaginatorOptions
	client    DescribeVolumesAPIClient
	params    *DescribeVolumesInput
	nextToken *string
	firstPage bool
}

// NewDescribeVolumesPaginator returns a new DescribeVolumesPaginator
func NewDescribeVolumesPaginator(client DescribeVolumesAPIClient, params *DescribeVolumesInput, optFns ...func(*DescribeVolumesPaginatorOptions)) *DescribeVolumesPaginator {
	if params == nil {
		params = &DescribeVolumesInput{}
	}

	options := DescribeVolumesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVolumesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVolumesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeVolumes page.
func (p *DescribeVolumesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVolumesOutput, error) {
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

	result, err := p.client.DescribeVolumes(ctx, &params, optFns...)
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

// VolumeAvailableWaiterOptions are waiter options for VolumeAvailableWaiter
type VolumeAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VolumeAvailableWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VolumeAvailableWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribeVolumesInput, *DescribeVolumesOutput, error) (bool, error)
}

// VolumeAvailableWaiter defines the waiters for VolumeAvailable
type VolumeAvailableWaiter struct {
	client DescribeVolumesAPIClient

	options VolumeAvailableWaiterOptions
}

// NewVolumeAvailableWaiter constructs a VolumeAvailableWaiter.
func NewVolumeAvailableWaiter(client DescribeVolumesAPIClient, optFns ...func(*VolumeAvailableWaiterOptions)) *VolumeAvailableWaiter {
	options := VolumeAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = volumeAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VolumeAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VolumeAvailable waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VolumeAvailableWaiter) Wait(ctx context.Context, params *DescribeVolumesInput, maxWaitDur time.Duration, optFns ...func(*VolumeAvailableWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
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

		out, err := w.client.DescribeVolumes(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
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
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for VolumeAvailable waiter")
}

func volumeAvailableStateRetryable(ctx context.Context, input *DescribeVolumesInput, output *DescribeVolumesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Volumes[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.VolumeState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VolumeState value, got %T", pathValue)
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
		pathValue, err := jmespath.Search("Volumes[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.VolumeState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VolumeState value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// VolumeInUseWaiterOptions are waiter options for VolumeInUseWaiter
type VolumeInUseWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VolumeInUseWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VolumeInUseWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeVolumesInput, *DescribeVolumesOutput, error) (bool, error)
}

// VolumeInUseWaiter defines the waiters for VolumeInUse
type VolumeInUseWaiter struct {
	client DescribeVolumesAPIClient

	options VolumeInUseWaiterOptions
}

// NewVolumeInUseWaiter constructs a VolumeInUseWaiter.
func NewVolumeInUseWaiter(client DescribeVolumesAPIClient, optFns ...func(*VolumeInUseWaiterOptions)) *VolumeInUseWaiter {
	options := VolumeInUseWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = volumeInUseStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VolumeInUseWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VolumeInUse waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VolumeInUseWaiter) Wait(ctx context.Context, params *DescribeVolumesInput, maxWaitDur time.Duration, optFns ...func(*VolumeInUseWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
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

		out, err := w.client.DescribeVolumes(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
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
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for VolumeInUse waiter")
}

func volumeInUseStateRetryable(ctx context.Context, input *DescribeVolumesInput, output *DescribeVolumesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Volumes[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "in-use"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.VolumeState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VolumeState value, got %T", pathValue)
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
		pathValue, err := jmespath.Search("Volumes[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.VolumeState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VolumeState value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumes",
	}
}
