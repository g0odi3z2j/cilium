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

// Describes the attributes of the specified Elastic IP addresses. For
// requirements, see Using reverse DNS for email applications
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#Using_Elastic_Addressing_Reverse_DNS).
func (c *Client) DescribeAddressesAttribute(ctx context.Context, params *DescribeAddressesAttributeInput, optFns ...func(*Options)) (*DescribeAddressesAttributeOutput, error) {
	if params == nil {
		params = &DescribeAddressesAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddressesAttribute", params, optFns, c.addOperationDescribeAddressesAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddressesAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressesAttributeInput struct {

	// [EC2-VPC] The allocation IDs.
	AllocationIds []string

	// The attribute of the IP address.
	Attribute types.AddressAttributeName

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

	noSmithyDocumentSerde
}

type DescribeAddressesAttributeOutput struct {

	// Information about the IP addresses.
	Addresses []types.AddressAttribute

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAddressesAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeAddressesAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAddressesAttribute{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddressesAttribute(options.Region), middleware.Before); err != nil {
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

// DescribeAddressesAttributeAPIClient is a client that implements the
// DescribeAddressesAttribute operation.
type DescribeAddressesAttributeAPIClient interface {
	DescribeAddressesAttribute(context.Context, *DescribeAddressesAttributeInput, ...func(*Options)) (*DescribeAddressesAttributeOutput, error)
}

var _ DescribeAddressesAttributeAPIClient = (*Client)(nil)

// DescribeAddressesAttributePaginatorOptions is the paginator options for
// DescribeAddressesAttribute
type DescribeAddressesAttributePaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddressesAttributePaginator is a paginator for
// DescribeAddressesAttribute
type DescribeAddressesAttributePaginator struct {
	options   DescribeAddressesAttributePaginatorOptions
	client    DescribeAddressesAttributeAPIClient
	params    *DescribeAddressesAttributeInput
	nextToken *string
	firstPage bool
}

// NewDescribeAddressesAttributePaginator returns a new
// DescribeAddressesAttributePaginator
func NewDescribeAddressesAttributePaginator(client DescribeAddressesAttributeAPIClient, params *DescribeAddressesAttributeInput, optFns ...func(*DescribeAddressesAttributePaginatorOptions)) *DescribeAddressesAttributePaginator {
	if params == nil {
		params = &DescribeAddressesAttributeInput{}
	}

	options := DescribeAddressesAttributePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAddressesAttributePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddressesAttributePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAddressesAttribute page.
func (p *DescribeAddressesAttributePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddressesAttributeOutput, error) {
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

	result, err := p.client.DescribeAddressesAttribute(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAddressesAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAddressesAttribute",
	}
}
