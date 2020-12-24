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

// Describes one or more of your NAT gateways.
func (c *Client) DescribeNatGateways(ctx context.Context, params *DescribeNatGatewaysInput, optFns ...func(*Options)) (*DescribeNatGatewaysOutput, error) {
	if params == nil {
		params = &DescribeNatGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNatGateways", params, optFns, addOperationDescribeNatGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNatGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNatGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * nat-gateway-id - The ID of the NAT gateway.
	//
	// * state -
	// The state of the NAT gateway (pending | failed | available | deleting |
	// deleted).
	//
	// * subnet-id - The ID of the subnet in which the NAT gateway
	// resides.
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
	// vpc-id - The ID of the VPC in which the NAT gateway resides.
	Filter []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// One or more NAT gateway IDs.
	NatGatewayIds []string

	// The token for the next page of results.
	NextToken *string
}

type DescribeNatGatewaysOutput struct {

	// Information about the NAT gateways.
	NatGateways []types.NatGateway

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNatGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNatGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNatGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNatGateways(options.Region), middleware.Before); err != nil {
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

// DescribeNatGatewaysAPIClient is a client that implements the DescribeNatGateways
// operation.
type DescribeNatGatewaysAPIClient interface {
	DescribeNatGateways(context.Context, *DescribeNatGatewaysInput, ...func(*Options)) (*DescribeNatGatewaysOutput, error)
}

var _ DescribeNatGatewaysAPIClient = (*Client)(nil)

// DescribeNatGatewaysPaginatorOptions is the paginator options for
// DescribeNatGateways
type DescribeNatGatewaysPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNatGatewaysPaginator is a paginator for DescribeNatGateways
type DescribeNatGatewaysPaginator struct {
	options   DescribeNatGatewaysPaginatorOptions
	client    DescribeNatGatewaysAPIClient
	params    *DescribeNatGatewaysInput
	nextToken *string
	firstPage bool
}

// NewDescribeNatGatewaysPaginator returns a new DescribeNatGatewaysPaginator
func NewDescribeNatGatewaysPaginator(client DescribeNatGatewaysAPIClient, params *DescribeNatGatewaysInput, optFns ...func(*DescribeNatGatewaysPaginatorOptions)) *DescribeNatGatewaysPaginator {
	options := DescribeNatGatewaysPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeNatGatewaysInput{}
	}

	return &DescribeNatGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNatGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeNatGateways page.
func (p *DescribeNatGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNatGatewaysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeNatGateways(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeNatGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNatGateways",
	}
}
