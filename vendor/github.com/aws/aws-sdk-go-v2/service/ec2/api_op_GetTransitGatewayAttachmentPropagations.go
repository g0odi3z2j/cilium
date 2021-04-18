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

// Lists the route tables to which the specified resource attachment propagates
// routes.
func (c *Client) GetTransitGatewayAttachmentPropagations(ctx context.Context, params *GetTransitGatewayAttachmentPropagationsInput, optFns ...func(*Options)) (*GetTransitGatewayAttachmentPropagationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayAttachmentPropagationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayAttachmentPropagations", params, optFns, addOperationGetTransitGatewayAttachmentPropagationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayAttachmentPropagationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayAttachmentPropagationsInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters. The possible values are:
	//
	// * transit-gateway-route-table-id
	// - The ID of the transit gateway route table.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string
}

type GetTransitGatewayAttachmentPropagationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the propagation route tables.
	TransitGatewayAttachmentPropagations []types.TransitGatewayAttachmentPropagation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayAttachmentPropagationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayAttachmentPropagations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayAttachmentPropagations{}, middleware.After)
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
	if err = addOpGetTransitGatewayAttachmentPropagationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayAttachmentPropagations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayAttachmentPropagationsAPIClient is a client that implements the
// GetTransitGatewayAttachmentPropagations operation.
type GetTransitGatewayAttachmentPropagationsAPIClient interface {
	GetTransitGatewayAttachmentPropagations(context.Context, *GetTransitGatewayAttachmentPropagationsInput, ...func(*Options)) (*GetTransitGatewayAttachmentPropagationsOutput, error)
}

var _ GetTransitGatewayAttachmentPropagationsAPIClient = (*Client)(nil)

// GetTransitGatewayAttachmentPropagationsPaginatorOptions is the paginator options
// for GetTransitGatewayAttachmentPropagations
type GetTransitGatewayAttachmentPropagationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayAttachmentPropagationsPaginator is a paginator for
// GetTransitGatewayAttachmentPropagations
type GetTransitGatewayAttachmentPropagationsPaginator struct {
	options   GetTransitGatewayAttachmentPropagationsPaginatorOptions
	client    GetTransitGatewayAttachmentPropagationsAPIClient
	params    *GetTransitGatewayAttachmentPropagationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayAttachmentPropagationsPaginator returns a new
// GetTransitGatewayAttachmentPropagationsPaginator
func NewGetTransitGatewayAttachmentPropagationsPaginator(client GetTransitGatewayAttachmentPropagationsAPIClient, params *GetTransitGatewayAttachmentPropagationsInput, optFns ...func(*GetTransitGatewayAttachmentPropagationsPaginatorOptions)) *GetTransitGatewayAttachmentPropagationsPaginator {
	if params == nil {
		params = &GetTransitGatewayAttachmentPropagationsInput{}
	}

	options := GetTransitGatewayAttachmentPropagationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayAttachmentPropagationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayAttachmentPropagationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetTransitGatewayAttachmentPropagations page.
func (p *GetTransitGatewayAttachmentPropagationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayAttachmentPropagationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetTransitGatewayAttachmentPropagations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayAttachmentPropagations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayAttachmentPropagations",
	}
}
