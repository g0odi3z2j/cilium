// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the prefix list references in a specified transit gateway
// route table.
func (c *Client) GetTransitGatewayPrefixListReferences(ctx context.Context, params *GetTransitGatewayPrefixListReferencesInput, optFns ...func(*Options)) (*GetTransitGatewayPrefixListReferencesOutput, error) {
	if params == nil {
		params = &GetTransitGatewayPrefixListReferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayPrefixListReferences", params, optFns, addOperationGetTransitGatewayPrefixListReferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayPrefixListReferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayPrefixListReferencesInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters. The possible values are:
	//
	// * attachment.resource-id - The ID
	// of the resource for the attachment.
	//
	// * attachment.resource-type - The type of
	// resource for the attachment. Valid values are vpc | vpn | direct-connect-gateway
	// | peering.
	//
	// * attachment.transit-gateway-attachment-id - The ID of the
	// attachment.
	//
	// * is-blackhole - Whether traffic matching the route is blocked
	// (true | false).
	//
	// * prefix-list-id - The ID of the prefix list.
	//
	// *
	// prefix-list-owner-id - The ID of the owner of the prefix list.
	//
	// * state - The
	// state of the prefix list reference (pending | available | modifying | deleting).
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string
}

type GetTransitGatewayPrefixListReferencesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the prefix list references.
	TransitGatewayPrefixListReferences []types.TransitGatewayPrefixListReference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayPrefixListReferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayPrefixListReferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayPrefixListReferences{}, middleware.After)
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
	if err = addOpGetTransitGatewayPrefixListReferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayPrefixListReferences(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayPrefixListReferencesAPIClient is a client that implements the
// GetTransitGatewayPrefixListReferences operation.
type GetTransitGatewayPrefixListReferencesAPIClient interface {
	GetTransitGatewayPrefixListReferences(context.Context, *GetTransitGatewayPrefixListReferencesInput, ...func(*Options)) (*GetTransitGatewayPrefixListReferencesOutput, error)
}

var _ GetTransitGatewayPrefixListReferencesAPIClient = (*Client)(nil)

// GetTransitGatewayPrefixListReferencesPaginatorOptions is the paginator options
// for GetTransitGatewayPrefixListReferences
type GetTransitGatewayPrefixListReferencesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayPrefixListReferencesPaginator is a paginator for
// GetTransitGatewayPrefixListReferences
type GetTransitGatewayPrefixListReferencesPaginator struct {
	options   GetTransitGatewayPrefixListReferencesPaginatorOptions
	client    GetTransitGatewayPrefixListReferencesAPIClient
	params    *GetTransitGatewayPrefixListReferencesInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayPrefixListReferencesPaginator returns a new
// GetTransitGatewayPrefixListReferencesPaginator
func NewGetTransitGatewayPrefixListReferencesPaginator(client GetTransitGatewayPrefixListReferencesAPIClient, params *GetTransitGatewayPrefixListReferencesInput, optFns ...func(*GetTransitGatewayPrefixListReferencesPaginatorOptions)) *GetTransitGatewayPrefixListReferencesPaginator {
	options := GetTransitGatewayPrefixListReferencesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetTransitGatewayPrefixListReferencesInput{}
	}

	return &GetTransitGatewayPrefixListReferencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayPrefixListReferencesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetTransitGatewayPrefixListReferences page.
func (p *GetTransitGatewayPrefixListReferencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayPrefixListReferencesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetTransitGatewayPrefixListReferences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayPrefixListReferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayPrefixListReferences",
	}
}
