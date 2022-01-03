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

// Describes one or more of your VPCs.
func (c *Client) DescribeVpcs(ctx context.Context, params *DescribeVpcsInput, optFns ...func(*Options)) (*DescribeVpcsOutput, error) {
	if params == nil {
		params = &DescribeVpcsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcs", params, optFns, c.addOperationDescribeVpcsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * cidr - The primary IPv4 CIDR block of the VPC. The CIDR
	// block you specify must exactly match the VPC's CIDR block for information to be
	// returned for the VPC. Must contain the slash followed by one or two digits (for
	// example, /28).
	//
	// * cidr-block-association.cidr-block - An IPv4 CIDR block
	// associated with the VPC.
	//
	// * cidr-block-association.association-id - The
	// association ID for an IPv4 CIDR block associated with the VPC.
	//
	// *
	// cidr-block-association.state - The state of an IPv4 CIDR block associated with
	// the VPC.
	//
	// * dhcp-options-id - The ID of a set of DHCP options.
	//
	// *
	// ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated with
	// the VPC.
	//
	// * ipv6-cidr-block-association.ipv6-pool - The ID of the IPv6 address
	// pool from which the IPv6 CIDR block is allocated.
	//
	// *
	// ipv6-cidr-block-association.association-id - The association ID for an IPv6 CIDR
	// block associated with the VPC.
	//
	// * ipv6-cidr-block-association.state - The state
	// of an IPv6 CIDR block associated with the VPC.
	//
	// * is-default - Indicates whether
	// the VPC is the default VPC.
	//
	// * owner-id - The ID of the Amazon Web Services
	// account that owns the VPC.
	//
	// * state - The state of the VPC (pending |
	// available).
	//
	// * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use this filter
	// to find all resources assigned a tag with a specific key, regardless of the tag
	// value.
	//
	// * vpc-id - The ID of the VPC.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more VPC IDs. Default: Describes all your VPCs.
	VpcIds []string

	noSmithyDocumentSerde
}

type DescribeVpcsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more VPCs.
	Vpcs []types.Vpc

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcs(options.Region), middleware.Before); err != nil {
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

// DescribeVpcsAPIClient is a client that implements the DescribeVpcs operation.
type DescribeVpcsAPIClient interface {
	DescribeVpcs(context.Context, *DescribeVpcsInput, ...func(*Options)) (*DescribeVpcsOutput, error)
}

var _ DescribeVpcsAPIClient = (*Client)(nil)

// DescribeVpcsPaginatorOptions is the paginator options for DescribeVpcs
type DescribeVpcsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcsPaginator is a paginator for DescribeVpcs
type DescribeVpcsPaginator struct {
	options   DescribeVpcsPaginatorOptions
	client    DescribeVpcsAPIClient
	params    *DescribeVpcsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcsPaginator returns a new DescribeVpcsPaginator
func NewDescribeVpcsPaginator(client DescribeVpcsAPIClient, params *DescribeVpcsInput, optFns ...func(*DescribeVpcsPaginatorOptions)) *DescribeVpcsPaginator {
	if params == nil {
		params = &DescribeVpcsInput{}
	}

	options := DescribeVpcsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcs page.
func (p *DescribeVpcsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcsOutput, error) {
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

	result, err := p.client.DescribeVpcs(ctx, &params, optFns...)
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

// VpcAvailableWaiterOptions are waiter options for VpcAvailableWaiter
type VpcAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VpcAvailableWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VpcAvailableWaiter will use default max delay of 120 seconds. Note that
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
	Retryable func(context.Context, *DescribeVpcsInput, *DescribeVpcsOutput, error) (bool, error)
}

// VpcAvailableWaiter defines the waiters for VpcAvailable
type VpcAvailableWaiter struct {
	client DescribeVpcsAPIClient

	options VpcAvailableWaiterOptions
}

// NewVpcAvailableWaiter constructs a VpcAvailableWaiter.
func NewVpcAvailableWaiter(client DescribeVpcsAPIClient, optFns ...func(*VpcAvailableWaiterOptions)) *VpcAvailableWaiter {
	options := VpcAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vpcAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VpcAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VpcAvailable waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VpcAvailableWaiter) Wait(ctx context.Context, params *DescribeVpcsInput, maxWaitDur time.Duration, optFns ...func(*VpcAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for VpcAvailable waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *VpcAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeVpcsInput, maxWaitDur time.Duration, optFns ...func(*VpcAvailableWaiterOptions)) (*DescribeVpcsOutput, error) {
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

		out, err := w.client.DescribeVpcs(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for VpcAvailable waiter")
}

func vpcAvailableStateRetryable(ctx context.Context, input *DescribeVpcsInput, output *DescribeVpcsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Vpcs[].State", output)
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
			value, ok := v.(types.VpcState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VpcState value, got %T", pathValue)
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

// VpcExistsWaiterOptions are waiter options for VpcExistsWaiter
type VpcExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VpcExistsWaiter will use default minimum delay of 1 seconds. Note that MinDelay
	// must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VpcExistsWaiter will use default max delay of 120 seconds. Note that
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
	Retryable func(context.Context, *DescribeVpcsInput, *DescribeVpcsOutput, error) (bool, error)
}

// VpcExistsWaiter defines the waiters for VpcExists
type VpcExistsWaiter struct {
	client DescribeVpcsAPIClient

	options VpcExistsWaiterOptions
}

// NewVpcExistsWaiter constructs a VpcExistsWaiter.
func NewVpcExistsWaiter(client DescribeVpcsAPIClient, optFns ...func(*VpcExistsWaiterOptions)) *VpcExistsWaiter {
	options := VpcExistsWaiterOptions{}
	options.MinDelay = 1 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vpcExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VpcExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VpcExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VpcExistsWaiter) Wait(ctx context.Context, params *DescribeVpcsInput, maxWaitDur time.Duration, optFns ...func(*VpcExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for VpcExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *VpcExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeVpcsInput, maxWaitDur time.Duration, optFns ...func(*VpcExistsWaiterOptions)) (*DescribeVpcsOutput, error) {
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

		out, err := w.client.DescribeVpcs(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for VpcExists waiter")
}

func vpcExistsStateRetryable(ctx context.Context, input *DescribeVpcsInput, output *DescribeVpcsOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidVpcID.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeVpcs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcs",
	}
}
