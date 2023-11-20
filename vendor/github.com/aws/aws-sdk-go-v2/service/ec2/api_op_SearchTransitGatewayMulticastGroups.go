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

// Searches one or more transit gateway multicast groups and returns the group
// membership information.
func (c *Client) SearchTransitGatewayMulticastGroups(ctx context.Context, params *SearchTransitGatewayMulticastGroupsInput, optFns ...func(*Options)) (*SearchTransitGatewayMulticastGroupsOutput, error) {
	if params == nil {
		params = &SearchTransitGatewayMulticastGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTransitGatewayMulticastGroups", params, optFns, c.addOperationSearchTransitGatewayMulticastGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTransitGatewayMulticastGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTransitGatewayMulticastGroupsInput struct {

	// The ID of the transit gateway multicast domain.
	//
	// This member is required.
	TransitGatewayMulticastDomainId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - group-ip-address - The IP address of the transit gateway multicast group.
	//   - is-group-member - The resource is a group member. Valid values are true |
	//   false .
	//   - is-group-source - The resource is a group source. Valid values are true |
	//   false .
	//   - member-type - The member type. Valid values are igmp | static .
	//   - resource-id - The ID of the resource.
	//   - resource-type - The type of resource. Valid values are vpc | vpn |
	//   direct-connect-gateway | tgw-peering .
	//   - source-type - The source type. Valid values are igmp | static .
	//   - subnet-id - The ID of the subnet.
	//   - transit-gateway-attachment-id - The id of the transit gateway attachment.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchTransitGatewayMulticastGroupsOutput struct {

	// Information about the transit gateway multicast group.
	MulticastGroups []types.TransitGatewayMulticastGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTransitGatewayMulticastGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpSearchTransitGatewayMulticastGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpSearchTransitGatewayMulticastGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchTransitGatewayMulticastGroups"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSearchTransitGatewayMulticastGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTransitGatewayMulticastGroups(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// SearchTransitGatewayMulticastGroupsAPIClient is a client that implements the
// SearchTransitGatewayMulticastGroups operation.
type SearchTransitGatewayMulticastGroupsAPIClient interface {
	SearchTransitGatewayMulticastGroups(context.Context, *SearchTransitGatewayMulticastGroupsInput, ...func(*Options)) (*SearchTransitGatewayMulticastGroupsOutput, error)
}

var _ SearchTransitGatewayMulticastGroupsAPIClient = (*Client)(nil)

// SearchTransitGatewayMulticastGroupsPaginatorOptions is the paginator options
// for SearchTransitGatewayMulticastGroups
type SearchTransitGatewayMulticastGroupsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchTransitGatewayMulticastGroupsPaginator is a paginator for
// SearchTransitGatewayMulticastGroups
type SearchTransitGatewayMulticastGroupsPaginator struct {
	options   SearchTransitGatewayMulticastGroupsPaginatorOptions
	client    SearchTransitGatewayMulticastGroupsAPIClient
	params    *SearchTransitGatewayMulticastGroupsInput
	nextToken *string
	firstPage bool
}

// NewSearchTransitGatewayMulticastGroupsPaginator returns a new
// SearchTransitGatewayMulticastGroupsPaginator
func NewSearchTransitGatewayMulticastGroupsPaginator(client SearchTransitGatewayMulticastGroupsAPIClient, params *SearchTransitGatewayMulticastGroupsInput, optFns ...func(*SearchTransitGatewayMulticastGroupsPaginatorOptions)) *SearchTransitGatewayMulticastGroupsPaginator {
	if params == nil {
		params = &SearchTransitGatewayMulticastGroupsInput{}
	}

	options := SearchTransitGatewayMulticastGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchTransitGatewayMulticastGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchTransitGatewayMulticastGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchTransitGatewayMulticastGroups page.
func (p *SearchTransitGatewayMulticastGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchTransitGatewayMulticastGroupsOutput, error) {
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

	result, err := p.client.SearchTransitGatewayMulticastGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchTransitGatewayMulticastGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchTransitGatewayMulticastGroups",
	}
}
