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

// Get a list of all the CIDR allocations in an IPAM pool. The Region you use
// should be the IPAM pool locale. The locale is the Amazon Web Services Region
// where this IPAM pool is available for allocations. If you use this action after
// AllocateIpamPoolCidr (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllocateIpamPoolCidr.html)
// or ReleaseIpamPoolAllocation (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseIpamPoolAllocation.html)
// , note that all EC2 API actions follow an eventual consistency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/query-api-troubleshooting.html#eventual-consistency)
// model.
func (c *Client) GetIpamPoolAllocations(ctx context.Context, params *GetIpamPoolAllocationsInput, optFns ...func(*Options)) (*GetIpamPoolAllocationsOutput, error) {
	if params == nil {
		params = &GetIpamPoolAllocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIpamPoolAllocations", params, optFns, c.addOperationGetIpamPoolAllocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIpamPoolAllocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIpamPoolAllocationsInput struct {

	// The ID of the IPAM pool you want to see the allocations for.
	//
	// This member is required.
	IpamPoolId *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters for the request. For more information about filtering, see
	// Filtering CLI output (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html)
	// .
	Filters []types.Filter

	// The ID of the allocation.
	IpamPoolAllocationId *string

	// The maximum number of results you would like returned per page.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIpamPoolAllocationsOutput struct {

	// The IPAM pool allocations you want information on.
	IpamPoolAllocations []types.IpamPoolAllocation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIpamPoolAllocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetIpamPoolAllocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetIpamPoolAllocations{}, middleware.After)
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
	if err = addOpGetIpamPoolAllocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIpamPoolAllocations(options.Region), middleware.Before); err != nil {
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

// GetIpamPoolAllocationsAPIClient is a client that implements the
// GetIpamPoolAllocations operation.
type GetIpamPoolAllocationsAPIClient interface {
	GetIpamPoolAllocations(context.Context, *GetIpamPoolAllocationsInput, ...func(*Options)) (*GetIpamPoolAllocationsOutput, error)
}

var _ GetIpamPoolAllocationsAPIClient = (*Client)(nil)

// GetIpamPoolAllocationsPaginatorOptions is the paginator options for
// GetIpamPoolAllocations
type GetIpamPoolAllocationsPaginatorOptions struct {
	// The maximum number of results you would like returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIpamPoolAllocationsPaginator is a paginator for GetIpamPoolAllocations
type GetIpamPoolAllocationsPaginator struct {
	options   GetIpamPoolAllocationsPaginatorOptions
	client    GetIpamPoolAllocationsAPIClient
	params    *GetIpamPoolAllocationsInput
	nextToken *string
	firstPage bool
}

// NewGetIpamPoolAllocationsPaginator returns a new GetIpamPoolAllocationsPaginator
func NewGetIpamPoolAllocationsPaginator(client GetIpamPoolAllocationsAPIClient, params *GetIpamPoolAllocationsInput, optFns ...func(*GetIpamPoolAllocationsPaginatorOptions)) *GetIpamPoolAllocationsPaginator {
	if params == nil {
		params = &GetIpamPoolAllocationsInput{}
	}

	options := GetIpamPoolAllocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIpamPoolAllocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIpamPoolAllocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIpamPoolAllocations page.
func (p *GetIpamPoolAllocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIpamPoolAllocationsOutput, error) {
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

	result, err := p.client.GetIpamPoolAllocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIpamPoolAllocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetIpamPoolAllocations",
	}
}
