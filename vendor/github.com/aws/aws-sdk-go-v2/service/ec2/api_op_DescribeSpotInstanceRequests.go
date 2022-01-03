// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified Spot Instance requests. You can use
// DescribeSpotInstanceRequests to find a running Spot Instance by examining the
// response. If the status of the Spot Instance is fulfilled, the instance ID
// appears in the response and contains the identifier of the instance.
// Alternatively, you can use DescribeInstances
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances)
// with a filter to look for instances where the instance lifecycle is spot. We
// recommend that you set MaxResults to a value between 5 and 1000 to limit the
// number of results returned. This paginates the output, which makes the list more
// manageable and returns the results faster. If the list of results exceeds your
// MaxResults value, then that number of results is returned along with a NextToken
// value that can be passed to a subsequent DescribeSpotInstanceRequests request to
// retrieve the remaining results. Spot Instance requests are deleted four hours
// after they are canceled and their instances are terminated.
func (c *Client) DescribeSpotInstanceRequests(ctx context.Context, params *DescribeSpotInstanceRequestsInput, optFns ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error) {
	if params == nil {
		params = &DescribeSpotInstanceRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotInstanceRequests", params, optFns, c.addOperationDescribeSpotInstanceRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotInstanceRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * availability-zone-group - The Availability Zone
	// group.
	//
	// * create-time - The time stamp when the Spot Instance request was
	// created.
	//
	// * fault-code - The fault code related to the request.
	//
	// * fault-message
	// - The fault message related to the request.
	//
	// * instance-id - The ID of the
	// instance that fulfilled the request.
	//
	// * launch-group - The Spot Instance launch
	// group.
	//
	// * launch.block-device-mapping.delete-on-termination - Indicates whether
	// the EBS volume is deleted on instance termination.
	//
	// *
	// launch.block-device-mapping.device-name - The device name for the volume in the
	// block device mapping (for example, /dev/sdh or xvdh).
	//
	// *
	// launch.block-device-mapping.snapshot-id - The ID of the snapshot for the EBS
	// volume.
	//
	// * launch.block-device-mapping.volume-size - The size of the EBS volume,
	// in GiB.
	//
	// * launch.block-device-mapping.volume-type - The type of EBS volume: gp2
	// for General Purpose SSD, io1 or io2 for Provisioned IOPS SSD, st1 for Throughput
	// Optimized HDD, sc1for Cold HDD, or standard for Magnetic.
	//
	// * launch.group-id -
	// The ID of the security group for the instance.
	//
	// * launch.group-name - The name
	// of the security group for the instance.
	//
	// * launch.image-id - The ID of the
	// AMI.
	//
	// * launch.instance-type - The type of instance (for example, m3.medium).
	//
	// *
	// launch.kernel-id - The kernel ID.
	//
	// * launch.key-name - The name of the key pair
	// the instance launched with.
	//
	// * launch.monitoring-enabled - Whether detailed
	// monitoring is enabled for the Spot Instance.
	//
	// * launch.ramdisk-id - The RAM disk
	// ID.
	//
	// * launched-availability-zone - The Availability Zone in which the request
	// is launched.
	//
	// * network-interface.addresses.primary - Indicates whether the IP
	// address is the primary private IP address.
	//
	// *
	// network-interface.delete-on-termination - Indicates whether the network
	// interface is deleted when the instance is terminated.
	//
	// *
	// network-interface.description - A description of the network interface.
	//
	// *
	// network-interface.device-index - The index of the device for the network
	// interface attachment on the instance.
	//
	// * network-interface.group-id - The ID of
	// the security group associated with the network interface.
	//
	// *
	// network-interface.network-interface-id - The ID of the network interface.
	//
	// *
	// network-interface.private-ip-address - The primary private IP address of the
	// network interface.
	//
	// * network-interface.subnet-id - The ID of the subnet for the
	// instance.
	//
	// * product-description - The product description associated with the
	// instance (Linux/UNIX | Windows).
	//
	// * spot-instance-request-id - The Spot Instance
	// request ID.
	//
	// * spot-price - The maximum hourly price for any Spot Instance
	// launched to fulfill the request.
	//
	// * state - The state of the Spot Instance
	// request (open | active | closed | cancelled | failed). Spot request status
	// information can help you track your Amazon EC2 Spot Instance requests. For more
	// information, see Spot request status
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	//
	// * status-code - The short code
	// describing the most recent evaluation of your Spot Instance request.
	//
	// *
	// status-message - The message explaining the status of the Spot Instance
	// request.
	//
	// * tag: - The key/value combination of a tag assigned to the resource.
	// Use the tag key in the filter name and the tag value as the filter value. For
	// example, to find all resources that have a tag with the key Owner and the value
	// TeamA, specify tag:Owner for the filter name and TeamA for the filter value.
	//
	// *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	//
	// *
	// type - The type of Spot Instance request (one-time | persistent).
	//
	// * valid-from
	// - The start date of the request.
	//
	// * valid-until - The end date of the request.
	Filters []types.Filter

	// The maximum number of results to return in a single call. Specify a value
	// between 5 and 1000. To retrieve the remaining results, make another call with
	// the returned NextToken value.
	MaxResults *int32

	// The token to request the next set of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// One or more Spot Instance request IDs.
	SpotInstanceRequestIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsOutput struct {

	// The token to use to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// One or more Spot Instance requests.
	SpotInstanceRequests []types.SpotInstanceRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSpotInstanceRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotInstanceRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotInstanceRequests{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotInstanceRequests(options.Region), middleware.Before); err != nil {
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

// DescribeSpotInstanceRequestsAPIClient is a client that implements the
// DescribeSpotInstanceRequests operation.
type DescribeSpotInstanceRequestsAPIClient interface {
	DescribeSpotInstanceRequests(context.Context, *DescribeSpotInstanceRequestsInput, ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error)
}

var _ DescribeSpotInstanceRequestsAPIClient = (*Client)(nil)

// DescribeSpotInstanceRequestsPaginatorOptions is the paginator options for
// DescribeSpotInstanceRequests
type DescribeSpotInstanceRequestsPaginatorOptions struct {
	// The maximum number of results to return in a single call. Specify a value
	// between 5 and 1000. To retrieve the remaining results, make another call with
	// the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSpotInstanceRequestsPaginator is a paginator for
// DescribeSpotInstanceRequests
type DescribeSpotInstanceRequestsPaginator struct {
	options   DescribeSpotInstanceRequestsPaginatorOptions
	client    DescribeSpotInstanceRequestsAPIClient
	params    *DescribeSpotInstanceRequestsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSpotInstanceRequestsPaginator returns a new
// DescribeSpotInstanceRequestsPaginator
func NewDescribeSpotInstanceRequestsPaginator(client DescribeSpotInstanceRequestsAPIClient, params *DescribeSpotInstanceRequestsInput, optFns ...func(*DescribeSpotInstanceRequestsPaginatorOptions)) *DescribeSpotInstanceRequestsPaginator {
	if params == nil {
		params = &DescribeSpotInstanceRequestsInput{}
	}

	options := DescribeSpotInstanceRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSpotInstanceRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSpotInstanceRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSpotInstanceRequests page.
func (p *DescribeSpotInstanceRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error) {
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

	result, err := p.client.DescribeSpotInstanceRequests(ctx, &params, optFns...)
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

// SpotInstanceRequestFulfilledWaiterOptions are waiter options for
// SpotInstanceRequestFulfilledWaiter
type SpotInstanceRequestFulfilledWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SpotInstanceRequestFulfilledWaiter will use default minimum delay of 15 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SpotInstanceRequestFulfilledWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeSpotInstanceRequestsInput, *DescribeSpotInstanceRequestsOutput, error) (bool, error)
}

// SpotInstanceRequestFulfilledWaiter defines the waiters for
// SpotInstanceRequestFulfilled
type SpotInstanceRequestFulfilledWaiter struct {
	client DescribeSpotInstanceRequestsAPIClient

	options SpotInstanceRequestFulfilledWaiterOptions
}

// NewSpotInstanceRequestFulfilledWaiter constructs a
// SpotInstanceRequestFulfilledWaiter.
func NewSpotInstanceRequestFulfilledWaiter(client DescribeSpotInstanceRequestsAPIClient, optFns ...func(*SpotInstanceRequestFulfilledWaiterOptions)) *SpotInstanceRequestFulfilledWaiter {
	options := SpotInstanceRequestFulfilledWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = spotInstanceRequestFulfilledStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SpotInstanceRequestFulfilledWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SpotInstanceRequestFulfilled waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *SpotInstanceRequestFulfilledWaiter) Wait(ctx context.Context, params *DescribeSpotInstanceRequestsInput, maxWaitDur time.Duration, optFns ...func(*SpotInstanceRequestFulfilledWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SpotInstanceRequestFulfilled waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *SpotInstanceRequestFulfilledWaiter) WaitForOutput(ctx context.Context, params *DescribeSpotInstanceRequestsInput, maxWaitDur time.Duration, optFns ...func(*SpotInstanceRequestFulfilledWaiterOptions)) (*DescribeSpotInstanceRequestsOutput, error) {
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

		out, err := w.client.DescribeSpotInstanceRequests(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for SpotInstanceRequestFulfilled waiter")
}

func spotInstanceRequestFulfilledStateRetryable(ctx context.Context, input *DescribeSpotInstanceRequestsInput, output *DescribeSpotInstanceRequestsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "fulfilled"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "request-canceled-and-instance-running"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "schedule-expired"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "canceled-before-fulfillment"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "bad-parameters"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("SpotInstanceRequests[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "system-error"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidSpotInstanceRequestID.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeSpotInstanceRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotInstanceRequests",
	}
}
