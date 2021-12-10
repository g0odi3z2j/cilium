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

// Get the CIDRs provisioned to an IPAM pool.
func (c *Client) GetIpamPoolCidrs(ctx context.Context, params *GetIpamPoolCidrsInput, optFns ...func(*Options)) (*GetIpamPoolCidrsOutput, error) {
	if params == nil {
		params = &GetIpamPoolCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIpamPoolCidrs", params, optFns, c.addOperationGetIpamPoolCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIpamPoolCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIpamPoolCidrsInput struct {

	// The ID of the IPAM pool you want the CIDR for.
	//
	// This member is required.
	IpamPoolId *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters for the request. For more information about filtering, see
	// Filtering CLI output
	// (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html).
	Filters []types.Filter

	// The maximum number of results to return in the request.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIpamPoolCidrsOutput struct {

	// Information about the CIDRs provisioned to an IPAM pool.
	IpamPoolCidrs []types.IpamPoolCidr

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIpamPoolCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetIpamPoolCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetIpamPoolCidrs{}, middleware.After)
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
	if err = addOpGetIpamPoolCidrsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIpamPoolCidrs(options.Region), middleware.Before); err != nil {
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

// GetIpamPoolCidrsAPIClient is a client that implements the GetIpamPoolCidrs
// operation.
type GetIpamPoolCidrsAPIClient interface {
	GetIpamPoolCidrs(context.Context, *GetIpamPoolCidrsInput, ...func(*Options)) (*GetIpamPoolCidrsOutput, error)
}

var _ GetIpamPoolCidrsAPIClient = (*Client)(nil)

// GetIpamPoolCidrsPaginatorOptions is the paginator options for GetIpamPoolCidrs
type GetIpamPoolCidrsPaginatorOptions struct {
	// The maximum number of results to return in the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIpamPoolCidrsPaginator is a paginator for GetIpamPoolCidrs
type GetIpamPoolCidrsPaginator struct {
	options   GetIpamPoolCidrsPaginatorOptions
	client    GetIpamPoolCidrsAPIClient
	params    *GetIpamPoolCidrsInput
	nextToken *string
	firstPage bool
}

// NewGetIpamPoolCidrsPaginator returns a new GetIpamPoolCidrsPaginator
func NewGetIpamPoolCidrsPaginator(client GetIpamPoolCidrsAPIClient, params *GetIpamPoolCidrsInput, optFns ...func(*GetIpamPoolCidrsPaginatorOptions)) *GetIpamPoolCidrsPaginator {
	if params == nil {
		params = &GetIpamPoolCidrsInput{}
	}

	options := GetIpamPoolCidrsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIpamPoolCidrsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIpamPoolCidrsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetIpamPoolCidrs page.
func (p *GetIpamPoolCidrsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIpamPoolCidrsOutput, error) {
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

	result, err := p.client.GetIpamPoolCidrs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIpamPoolCidrs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetIpamPoolCidrs",
	}
}
