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

// Describes the principals (service consumers) that are permitted to discover your
// VPC endpoint service.
func (c *Client) DescribeVpcEndpointServicePermissions(ctx context.Context, params *DescribeVpcEndpointServicePermissionsInput, optFns ...func(*Options)) (*DescribeVpcEndpointServicePermissionsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointServicePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpointServicePermissions", params, optFns, c.addOperationDescribeVpcEndpointServicePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointServicePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointServicePermissionsInput struct {

	// The ID of the service.
	//
	// This member is required.
	ServiceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * principal - The ARN of the principal.
	//
	// * principal-type
	// - The principal type (All | Service | OrganizationUnit | Account | User | Role).
	Filters []types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeVpcEndpointServicePermissionsOutput struct {

	// Information about one or more allowed principals.
	AllowedPrincipals []types.AllowedPrincipal

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcEndpointServicePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointServicePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointServicePermissions{}, middleware.After)
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
	if err = addOpDescribeVpcEndpointServicePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointServicePermissions(options.Region), middleware.Before); err != nil {
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

// DescribeVpcEndpointServicePermissionsAPIClient is a client that implements the
// DescribeVpcEndpointServicePermissions operation.
type DescribeVpcEndpointServicePermissionsAPIClient interface {
	DescribeVpcEndpointServicePermissions(context.Context, *DescribeVpcEndpointServicePermissionsInput, ...func(*Options)) (*DescribeVpcEndpointServicePermissionsOutput, error)
}

var _ DescribeVpcEndpointServicePermissionsAPIClient = (*Client)(nil)

// DescribeVpcEndpointServicePermissionsPaginatorOptions is the paginator options
// for DescribeVpcEndpointServicePermissions
type DescribeVpcEndpointServicePermissionsPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcEndpointServicePermissionsPaginator is a paginator for
// DescribeVpcEndpointServicePermissions
type DescribeVpcEndpointServicePermissionsPaginator struct {
	options   DescribeVpcEndpointServicePermissionsPaginatorOptions
	client    DescribeVpcEndpointServicePermissionsAPIClient
	params    *DescribeVpcEndpointServicePermissionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcEndpointServicePermissionsPaginator returns a new
// DescribeVpcEndpointServicePermissionsPaginator
func NewDescribeVpcEndpointServicePermissionsPaginator(client DescribeVpcEndpointServicePermissionsAPIClient, params *DescribeVpcEndpointServicePermissionsInput, optFns ...func(*DescribeVpcEndpointServicePermissionsPaginatorOptions)) *DescribeVpcEndpointServicePermissionsPaginator {
	if params == nil {
		params = &DescribeVpcEndpointServicePermissionsInput{}
	}

	options := DescribeVpcEndpointServicePermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcEndpointServicePermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcEndpointServicePermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcEndpointServicePermissions page.
func (p *DescribeVpcEndpointServicePermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcEndpointServicePermissionsOutput, error) {
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

	result, err := p.client.DescribeVpcEndpointServicePermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVpcEndpointServicePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointServicePermissions",
	}
}
