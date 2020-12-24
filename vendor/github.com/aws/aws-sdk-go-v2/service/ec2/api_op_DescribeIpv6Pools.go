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

// Describes your IPv6 address pools.
func (c *Client) DescribeIpv6Pools(ctx context.Context, params *DescribeIpv6PoolsInput, optFns ...func(*Options)) (*DescribeIpv6PoolsOutput, error) {
	if params == nil {
		params = &DescribeIpv6PoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIpv6Pools", params, optFns, addOperationDescribeIpv6PoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIpv6PoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIpv6PoolsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
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
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the IPv6 address pools.
	PoolIds []string
}

type DescribeIpv6PoolsOutput struct {

	// Information about the IPv6 address pools.
	Ipv6Pools []types.Ipv6Pool

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIpv6PoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeIpv6Pools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIpv6Pools{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIpv6Pools(options.Region), middleware.Before); err != nil {
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

// DescribeIpv6PoolsAPIClient is a client that implements the DescribeIpv6Pools
// operation.
type DescribeIpv6PoolsAPIClient interface {
	DescribeIpv6Pools(context.Context, *DescribeIpv6PoolsInput, ...func(*Options)) (*DescribeIpv6PoolsOutput, error)
}

var _ DescribeIpv6PoolsAPIClient = (*Client)(nil)

// DescribeIpv6PoolsPaginatorOptions is the paginator options for DescribeIpv6Pools
type DescribeIpv6PoolsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeIpv6PoolsPaginator is a paginator for DescribeIpv6Pools
type DescribeIpv6PoolsPaginator struct {
	options   DescribeIpv6PoolsPaginatorOptions
	client    DescribeIpv6PoolsAPIClient
	params    *DescribeIpv6PoolsInput
	nextToken *string
	firstPage bool
}

// NewDescribeIpv6PoolsPaginator returns a new DescribeIpv6PoolsPaginator
func NewDescribeIpv6PoolsPaginator(client DescribeIpv6PoolsAPIClient, params *DescribeIpv6PoolsInput, optFns ...func(*DescribeIpv6PoolsPaginatorOptions)) *DescribeIpv6PoolsPaginator {
	options := DescribeIpv6PoolsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeIpv6PoolsInput{}
	}

	return &DescribeIpv6PoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeIpv6PoolsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeIpv6Pools page.
func (p *DescribeIpv6PoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeIpv6PoolsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeIpv6Pools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeIpv6Pools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeIpv6Pools",
	}
}
