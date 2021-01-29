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

// [VPC only] Describes the stale security group rules for security groups in a
// specified VPC. Rules are stale when they reference a deleted security group in a
// peer VPC, or a security group in a peer VPC for which the VPC peering connection
// has been deleted.
func (c *Client) DescribeStaleSecurityGroups(ctx context.Context, params *DescribeStaleSecurityGroupsInput, optFns ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeStaleSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStaleSecurityGroups", params, optFns, addOperationDescribeStaleSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStaleSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStaleSecurityGroupsInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// prior call.)
	NextToken *string
}

type DescribeStaleSecurityGroupsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the stale security groups.
	StaleSecurityGroupSet []types.StaleSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStaleSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeStaleSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeStaleSecurityGroups{}, middleware.After)
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
	if err = addOpDescribeStaleSecurityGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStaleSecurityGroups(options.Region), middleware.Before); err != nil {
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

// DescribeStaleSecurityGroupsAPIClient is a client that implements the
// DescribeStaleSecurityGroups operation.
type DescribeStaleSecurityGroupsAPIClient interface {
	DescribeStaleSecurityGroups(context.Context, *DescribeStaleSecurityGroupsInput, ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error)
}

var _ DescribeStaleSecurityGroupsAPIClient = (*Client)(nil)

// DescribeStaleSecurityGroupsPaginatorOptions is the paginator options for
// DescribeStaleSecurityGroups
type DescribeStaleSecurityGroupsPaginatorOptions struct {
	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStaleSecurityGroupsPaginator is a paginator for
// DescribeStaleSecurityGroups
type DescribeStaleSecurityGroupsPaginator struct {
	options   DescribeStaleSecurityGroupsPaginatorOptions
	client    DescribeStaleSecurityGroupsAPIClient
	params    *DescribeStaleSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeStaleSecurityGroupsPaginator returns a new
// DescribeStaleSecurityGroupsPaginator
func NewDescribeStaleSecurityGroupsPaginator(client DescribeStaleSecurityGroupsAPIClient, params *DescribeStaleSecurityGroupsInput, optFns ...func(*DescribeStaleSecurityGroupsPaginatorOptions)) *DescribeStaleSecurityGroupsPaginator {
	options := DescribeStaleSecurityGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeStaleSecurityGroupsInput{}
	}

	return &DescribeStaleSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStaleSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeStaleSecurityGroups page.
func (p *DescribeStaleSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeStaleSecurityGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStaleSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeStaleSecurityGroups",
	}
}
