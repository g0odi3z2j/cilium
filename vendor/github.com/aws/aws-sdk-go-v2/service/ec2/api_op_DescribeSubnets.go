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

// Describes one or more of your subnets. For more information, see Your VPC and
// subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in
// the Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeSubnets(ctx context.Context, params *DescribeSubnetsInput, optFns ...func(*Options)) (*DescribeSubnetsOutput, error) {
	if params == nil {
		params = &DescribeSubnetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSubnets", params, optFns, c.addOperationDescribeSubnetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSubnetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSubnetsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * availability-zone - The Availability Zone for the
	// subnet. You can also use availabilityZone as the filter name.
	//
	// *
	// availability-zone-id - The ID of the Availability Zone for the subnet. You can
	// also use availabilityZoneId as the filter name.
	//
	// * available-ip-address-count -
	// The number of IPv4 addresses in the subnet that are available.
	//
	// * cidr-block -
	// The IPv4 CIDR block of the subnet. The CIDR block you specify must exactly match
	// the subnet's CIDR block for information to be returned for the subnet. You can
	// also use cidr or cidrBlock as the filter names.
	//
	// * default-for-az - Indicates
	// whether this is the default subnet for the Availability Zone (true | false). You
	// can also use defaultForAz as the filter name.
	//
	// *
	// ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated with
	// the subnet.
	//
	// * ipv6-cidr-block-association.association-id - An association ID
	// for an IPv6 CIDR block associated with the subnet.
	//
	// *
	// ipv6-cidr-block-association.state - The state of an IPv6 CIDR block associated
	// with the subnet.
	//
	// * ipv6-native - Indicates whether this is an IPv6 only subnet
	// (true | false).
	//
	// * outpost-arn - The Amazon Resource Name (ARN) of the
	// Outpost.
	//
	// * owner-id - The ID of the Amazon Web Services account that owns the
	// subnet.
	//
	// * state - The state of the subnet (pending | available).
	//
	// * subnet-arn
	// - The Amazon Resource Name (ARN) of the subnet.
	//
	// * subnet-id - The ID of the
	// subnet.
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
	// vpc-id - The ID of the VPC for the subnet.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more subnet IDs. Default: Describes all your subnets.
	SubnetIds []string

	noSmithyDocumentSerde
}

type DescribeSubnetsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more subnets.
	Subnets []types.Subnet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSubnetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSubnets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSubnets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubnets(options.Region), middleware.Before); err != nil {
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

// DescribeSubnetsAPIClient is a client that implements the DescribeSubnets
// operation.
type DescribeSubnetsAPIClient interface {
	DescribeSubnets(context.Context, *DescribeSubnetsInput, ...func(*Options)) (*DescribeSubnetsOutput, error)
}

var _ DescribeSubnetsAPIClient = (*Client)(nil)

// DescribeSubnetsPaginatorOptions is the paginator options for DescribeSubnets
type DescribeSubnetsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSubnetsPaginator is a paginator for DescribeSubnets
type DescribeSubnetsPaginator struct {
	options   DescribeSubnetsPaginatorOptions
	client    DescribeSubnetsAPIClient
	params    *DescribeSubnetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSubnetsPaginator returns a new DescribeSubnetsPaginator
func NewDescribeSubnetsPaginator(client DescribeSubnetsAPIClient, params *DescribeSubnetsInput, optFns ...func(*DescribeSubnetsPaginatorOptions)) *DescribeSubnetsPaginator {
	if params == nil {
		params = &DescribeSubnetsInput{}
	}

	options := DescribeSubnetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSubnetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSubnetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSubnets page.
func (p *DescribeSubnetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSubnetsOutput, error) {
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

	result, err := p.client.DescribeSubnets(ctx, &params, optFns...)
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

// SubnetAvailableWaiterOptions are waiter options for SubnetAvailableWaiter
type SubnetAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SubnetAvailableWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SubnetAvailableWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribeSubnetsInput, *DescribeSubnetsOutput, error) (bool, error)
}

// SubnetAvailableWaiter defines the waiters for SubnetAvailable
type SubnetAvailableWaiter struct {
	client DescribeSubnetsAPIClient

	options SubnetAvailableWaiterOptions
}

// NewSubnetAvailableWaiter constructs a SubnetAvailableWaiter.
func NewSubnetAvailableWaiter(client DescribeSubnetsAPIClient, optFns ...func(*SubnetAvailableWaiterOptions)) *SubnetAvailableWaiter {
	options := SubnetAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = subnetAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SubnetAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SubnetAvailable waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *SubnetAvailableWaiter) Wait(ctx context.Context, params *DescribeSubnetsInput, maxWaitDur time.Duration, optFns ...func(*SubnetAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SubnetAvailable waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *SubnetAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeSubnetsInput, maxWaitDur time.Duration, optFns ...func(*SubnetAvailableWaiterOptions)) (*DescribeSubnetsOutput, error) {
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

		out, err := w.client.DescribeSubnets(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for SubnetAvailable waiter")
}

func subnetAvailableStateRetryable(ctx context.Context, input *DescribeSubnetsInput, output *DescribeSubnetsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Subnets[].State", output)
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
			value, ok := v.(types.SubnetState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.SubnetState value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeSubnets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSubnets",
	}
}
