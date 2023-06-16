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

// Describes one or more VPC attachments. By default, all VPC attachments are
// described. Alternatively, you can filter the results.
func (c *Client) DescribeTransitGatewayVpcAttachments(ctx context.Context, params *DescribeTransitGatewayVpcAttachmentsInput, optFns ...func(*Options)) (*DescribeTransitGatewayVpcAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayVpcAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayVpcAttachments", params, optFns, c.addOperationDescribeTransitGatewayVpcAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayVpcAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayVpcAttachmentsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - state - The state of the attachment. Valid values are available | deleted |
	//   deleting | failed | failing | initiatingRequest | modifying |
	//   pendingAcceptance | pending | rollingBack | rejected | rejecting .
	//   - transit-gateway-attachment-id - The ID of the attachment.
	//   - transit-gateway-id - The ID of the transit gateway.
	//   - vpc-id - The ID of the VPC.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the attachments.
	TransitGatewayAttachmentIds []string

	noSmithyDocumentSerde
}

type DescribeTransitGatewayVpcAttachmentsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the VPC attachments.
	TransitGatewayVpcAttachments []types.TransitGatewayVpcAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewayVpcAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayVpcAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayVpcAttachments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayVpcAttachments(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayVpcAttachmentsAPIClient is a client that implements the
// DescribeTransitGatewayVpcAttachments operation.
type DescribeTransitGatewayVpcAttachmentsAPIClient interface {
	DescribeTransitGatewayVpcAttachments(context.Context, *DescribeTransitGatewayVpcAttachmentsInput, ...func(*Options)) (*DescribeTransitGatewayVpcAttachmentsOutput, error)
}

var _ DescribeTransitGatewayVpcAttachmentsAPIClient = (*Client)(nil)

// DescribeTransitGatewayVpcAttachmentsPaginatorOptions is the paginator options
// for DescribeTransitGatewayVpcAttachments
type DescribeTransitGatewayVpcAttachmentsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayVpcAttachmentsPaginator is a paginator for
// DescribeTransitGatewayVpcAttachments
type DescribeTransitGatewayVpcAttachmentsPaginator struct {
	options   DescribeTransitGatewayVpcAttachmentsPaginatorOptions
	client    DescribeTransitGatewayVpcAttachmentsAPIClient
	params    *DescribeTransitGatewayVpcAttachmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayVpcAttachmentsPaginator returns a new
// DescribeTransitGatewayVpcAttachmentsPaginator
func NewDescribeTransitGatewayVpcAttachmentsPaginator(client DescribeTransitGatewayVpcAttachmentsAPIClient, params *DescribeTransitGatewayVpcAttachmentsInput, optFns ...func(*DescribeTransitGatewayVpcAttachmentsPaginatorOptions)) *DescribeTransitGatewayVpcAttachmentsPaginator {
	if params == nil {
		params = &DescribeTransitGatewayVpcAttachmentsInput{}
	}

	options := DescribeTransitGatewayVpcAttachmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayVpcAttachmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayVpcAttachmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTransitGatewayVpcAttachments page.
func (p *DescribeTransitGatewayVpcAttachmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayVpcAttachmentsOutput, error) {
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

	result, err := p.client.DescribeTransitGatewayVpcAttachments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayVpcAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayVpcAttachments",
	}
}
