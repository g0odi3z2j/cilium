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

// Describes one or more transit gateway multicast domains.
func (c *Client) DescribeTransitGatewayMulticastDomains(ctx context.Context, params *DescribeTransitGatewayMulticastDomainsInput, optFns ...func(*Options)) (*DescribeTransitGatewayMulticastDomainsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayMulticastDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayMulticastDomains", params, optFns, addOperationDescribeTransitGatewayMulticastDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayMulticastDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayMulticastDomainsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	// * state - The state of the
	// transit gateway multicast domain. Valid values are pending | available |
	// deleting | deleted.
	//
	// * transit-gateway-id - The ID of the transit gateway.
	//
	// *
	// transit-gateway-multicast-domain-id - The ID of the transit gateway multicast
	// domain.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainIds []string
}

type DescribeTransitGatewayMulticastDomainsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the transit gateway multicast domains.
	TransitGatewayMulticastDomains []types.TransitGatewayMulticastDomain

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewayMulticastDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayMulticastDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayMulticastDomains{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayMulticastDomains(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayMulticastDomainsAPIClient is a client that implements the
// DescribeTransitGatewayMulticastDomains operation.
type DescribeTransitGatewayMulticastDomainsAPIClient interface {
	DescribeTransitGatewayMulticastDomains(context.Context, *DescribeTransitGatewayMulticastDomainsInput, ...func(*Options)) (*DescribeTransitGatewayMulticastDomainsOutput, error)
}

var _ DescribeTransitGatewayMulticastDomainsAPIClient = (*Client)(nil)

// DescribeTransitGatewayMulticastDomainsPaginatorOptions is the paginator options
// for DescribeTransitGatewayMulticastDomains
type DescribeTransitGatewayMulticastDomainsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayMulticastDomainsPaginator is a paginator for
// DescribeTransitGatewayMulticastDomains
type DescribeTransitGatewayMulticastDomainsPaginator struct {
	options   DescribeTransitGatewayMulticastDomainsPaginatorOptions
	client    DescribeTransitGatewayMulticastDomainsAPIClient
	params    *DescribeTransitGatewayMulticastDomainsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayMulticastDomainsPaginator returns a new
// DescribeTransitGatewayMulticastDomainsPaginator
func NewDescribeTransitGatewayMulticastDomainsPaginator(client DescribeTransitGatewayMulticastDomainsAPIClient, params *DescribeTransitGatewayMulticastDomainsInput, optFns ...func(*DescribeTransitGatewayMulticastDomainsPaginatorOptions)) *DescribeTransitGatewayMulticastDomainsPaginator {
	if params == nil {
		params = &DescribeTransitGatewayMulticastDomainsInput{}
	}

	options := DescribeTransitGatewayMulticastDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayMulticastDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayMulticastDomainsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeTransitGatewayMulticastDomains page.
func (p *DescribeTransitGatewayMulticastDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayMulticastDomainsOutput, error) {
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

	result, err := p.client.DescribeTransitGatewayMulticastDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayMulticastDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayMulticastDomains",
	}
}
