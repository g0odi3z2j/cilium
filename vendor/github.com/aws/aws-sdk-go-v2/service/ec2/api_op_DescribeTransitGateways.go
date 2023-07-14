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

// Describes one or more transit gateways. By default, all transit gateways are
// described. Alternatively, you can filter the results.
func (c *Client) DescribeTransitGateways(ctx context.Context, params *DescribeTransitGatewaysInput, optFns ...func(*Options)) (*DescribeTransitGatewaysOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGateways", params, optFns, c.addOperationDescribeTransitGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - options.propagation-default-route-table-id - The ID of the default
	//   propagation route table.
	//   - options.amazon-side-asn - The private ASN for the Amazon side of a BGP
	//   session.
	//   - options.association-default-route-table-id - The ID of the default
	//   association route table.
	//   - options.auto-accept-shared-attachments - Indicates whether there is
	//   automatic acceptance of attachment requests ( enable | disable ).
	//   - options.default-route-table-association - Indicates whether resource
	//   attachments are automatically associated with the default association route
	//   table ( enable | disable ).
	//   - options.default-route-table-propagation - Indicates whether resource
	//   attachments automatically propagate routes to the default propagation route
	//   table ( enable | disable ).
	//   - options.dns-support - Indicates whether DNS support is enabled ( enable |
	//   disable ).
	//   - options.vpn-ecmp-support - Indicates whether Equal Cost Multipath Protocol
	//   support is enabled ( enable | disable ).
	//   - owner-id - The ID of the Amazon Web Services account that owns the transit
	//   gateway.
	//   - state - The state of the transit gateway ( available | deleted | deleting |
	//   modifying | pending ).
	//   - transit-gateway-id - The ID of the transit gateway.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the transit gateways.
	TransitGatewayIds []string

	noSmithyDocumentSerde
}

type DescribeTransitGatewaysOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the transit gateways.
	TransitGateways []types.TransitGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGateways(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewaysAPIClient is a client that implements the
// DescribeTransitGateways operation.
type DescribeTransitGatewaysAPIClient interface {
	DescribeTransitGateways(context.Context, *DescribeTransitGatewaysInput, ...func(*Options)) (*DescribeTransitGatewaysOutput, error)
}

var _ DescribeTransitGatewaysAPIClient = (*Client)(nil)

// DescribeTransitGatewaysPaginatorOptions is the paginator options for
// DescribeTransitGateways
type DescribeTransitGatewaysPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewaysPaginator is a paginator for DescribeTransitGateways
type DescribeTransitGatewaysPaginator struct {
	options   DescribeTransitGatewaysPaginatorOptions
	client    DescribeTransitGatewaysAPIClient
	params    *DescribeTransitGatewaysInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewaysPaginator returns a new
// DescribeTransitGatewaysPaginator
func NewDescribeTransitGatewaysPaginator(client DescribeTransitGatewaysAPIClient, params *DescribeTransitGatewaysInput, optFns ...func(*DescribeTransitGatewaysPaginatorOptions)) *DescribeTransitGatewaysPaginator {
	if params == nil {
		params = &DescribeTransitGatewaysInput{}
	}

	options := DescribeTransitGatewaysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTransitGateways page.
func (p *DescribeTransitGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewaysOutput, error) {
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

	result, err := p.client.DescribeTransitGateways(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGateways",
	}
}
