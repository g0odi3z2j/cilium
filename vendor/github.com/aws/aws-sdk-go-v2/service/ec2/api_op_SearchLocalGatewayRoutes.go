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

// Searches for routes in the specified local gateway route table.
func (c *Client) SearchLocalGatewayRoutes(ctx context.Context, params *SearchLocalGatewayRoutesInput, optFns ...func(*Options)) (*SearchLocalGatewayRoutesOutput, error) {
	if params == nil {
		params = &SearchLocalGatewayRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchLocalGatewayRoutes", params, optFns, addOperationSearchLocalGatewayRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchLocalGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchLocalGatewayRoutesInput struct {

	// One or more filters.
	//
	// This member is required.
	Filters []types.Filter

	// The ID of the local gateway route table.
	//
	// This member is required.
	LocalGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type SearchLocalGatewayRoutesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the routes.
	Routes []types.LocalGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchLocalGatewayRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpSearchLocalGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpSearchLocalGatewayRoutes{}, middleware.After)
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
	if err = addOpSearchLocalGatewayRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchLocalGatewayRoutes(options.Region), middleware.Before); err != nil {
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

// SearchLocalGatewayRoutesAPIClient is a client that implements the
// SearchLocalGatewayRoutes operation.
type SearchLocalGatewayRoutesAPIClient interface {
	SearchLocalGatewayRoutes(context.Context, *SearchLocalGatewayRoutesInput, ...func(*Options)) (*SearchLocalGatewayRoutesOutput, error)
}

var _ SearchLocalGatewayRoutesAPIClient = (*Client)(nil)

// SearchLocalGatewayRoutesPaginatorOptions is the paginator options for
// SearchLocalGatewayRoutes
type SearchLocalGatewayRoutesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchLocalGatewayRoutesPaginator is a paginator for SearchLocalGatewayRoutes
type SearchLocalGatewayRoutesPaginator struct {
	options   SearchLocalGatewayRoutesPaginatorOptions
	client    SearchLocalGatewayRoutesAPIClient
	params    *SearchLocalGatewayRoutesInput
	nextToken *string
	firstPage bool
}

// NewSearchLocalGatewayRoutesPaginator returns a new
// SearchLocalGatewayRoutesPaginator
func NewSearchLocalGatewayRoutesPaginator(client SearchLocalGatewayRoutesAPIClient, params *SearchLocalGatewayRoutesInput, optFns ...func(*SearchLocalGatewayRoutesPaginatorOptions)) *SearchLocalGatewayRoutesPaginator {
	if params == nil {
		params = &SearchLocalGatewayRoutesInput{}
	}

	options := SearchLocalGatewayRoutesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchLocalGatewayRoutesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchLocalGatewayRoutesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SearchLocalGatewayRoutes page.
func (p *SearchLocalGatewayRoutesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchLocalGatewayRoutesOutput, error) {
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

	result, err := p.client.SearchLocalGatewayRoutes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchLocalGatewayRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "SearchLocalGatewayRoutes",
	}
}
