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

// Describes your transit gateway peering attachments.
func (c *Client) DescribeTransitGatewayPeeringAttachments(ctx context.Context, params *DescribeTransitGatewayPeeringAttachmentsInput, optFns ...func(*Options)) (*DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayPeeringAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayPeeringAttachments", params, optFns, addOperationDescribeTransitGatewayPeeringAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayPeeringAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayPeeringAttachmentsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters. The possible values are:
	//
	// * transit-gateway-attachment-id -
	// The ID of the transit gateway attachment.
	//
	// * local-owner-id - The ID of your AWS
	// account.
	//
	// * remote-owner-id - The ID of the AWS account in the remote Region
	// that owns the transit gateway.
	//
	// * state - The state of the peering attachment.
	// Valid values are available | deleted | deleting | failed | failing |
	// initiatingRequest | modifying | pendingAcceptance | pending | rollingBack |
	// rejected | rejecting).
	//
	// * tag: - The key/value combination of a tag assigned to
	// the resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use this filter
	// to find all resources that have a tag with a specific key, regardless of the tag
	// value.
	//
	// * transit-gateway-id - The ID of the transit gateway.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// One or more IDs of the transit gateway peering attachments.
	TransitGatewayAttachmentIds []string
}

type DescribeTransitGatewayPeeringAttachmentsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The transit gateway peering attachments.
	TransitGatewayPeeringAttachments []types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewayPeeringAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayPeeringAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayPeeringAttachments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayPeeringAttachments(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayPeeringAttachmentsAPIClient is a client that implements
// the DescribeTransitGatewayPeeringAttachments operation.
type DescribeTransitGatewayPeeringAttachmentsAPIClient interface {
	DescribeTransitGatewayPeeringAttachments(context.Context, *DescribeTransitGatewayPeeringAttachmentsInput, ...func(*Options)) (*DescribeTransitGatewayPeeringAttachmentsOutput, error)
}

var _ DescribeTransitGatewayPeeringAttachmentsAPIClient = (*Client)(nil)

// DescribeTransitGatewayPeeringAttachmentsPaginatorOptions is the paginator
// options for DescribeTransitGatewayPeeringAttachments
type DescribeTransitGatewayPeeringAttachmentsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayPeeringAttachmentsPaginator is a paginator for
// DescribeTransitGatewayPeeringAttachments
type DescribeTransitGatewayPeeringAttachmentsPaginator struct {
	options   DescribeTransitGatewayPeeringAttachmentsPaginatorOptions
	client    DescribeTransitGatewayPeeringAttachmentsAPIClient
	params    *DescribeTransitGatewayPeeringAttachmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayPeeringAttachmentsPaginator returns a new
// DescribeTransitGatewayPeeringAttachmentsPaginator
func NewDescribeTransitGatewayPeeringAttachmentsPaginator(client DescribeTransitGatewayPeeringAttachmentsAPIClient, params *DescribeTransitGatewayPeeringAttachmentsInput, optFns ...func(*DescribeTransitGatewayPeeringAttachmentsPaginatorOptions)) *DescribeTransitGatewayPeeringAttachmentsPaginator {
	if params == nil {
		params = &DescribeTransitGatewayPeeringAttachmentsInput{}
	}

	options := DescribeTransitGatewayPeeringAttachmentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayPeeringAttachmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayPeeringAttachmentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeTransitGatewayPeeringAttachments page.
func (p *DescribeTransitGatewayPeeringAttachmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeTransitGatewayPeeringAttachments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayPeeringAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayPeeringAttachments",
	}
}
