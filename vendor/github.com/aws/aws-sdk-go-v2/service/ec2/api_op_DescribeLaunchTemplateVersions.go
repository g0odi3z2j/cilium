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

// Describes one or more versions of a specified launch template. You can describe
// all versions, individual versions, or a range of versions. You can also describe
// all the latest versions or all the default versions of all the launch templates
// in your account.
func (c *Client) DescribeLaunchTemplateVersions(ctx context.Context, params *DescribeLaunchTemplateVersionsInput, optFns ...func(*Options)) (*DescribeLaunchTemplateVersionsOutput, error) {
	if params == nil {
		params = &DescribeLaunchTemplateVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLaunchTemplateVersions", params, optFns, addOperationDescribeLaunchTemplateVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLaunchTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLaunchTemplateVersionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * create-time - The time the launch template version was
	// created.
	//
	// * ebs-optimized - A boolean that indicates whether the instance is
	// optimized for Amazon EBS I/O.
	//
	// * iam-instance-profile - The ARN of the IAM
	// instance profile.
	//
	// * image-id - The ID of the AMI.
	//
	// * instance-type - The
	// instance type.
	//
	// * is-default-version - A boolean that indicates whether the
	// launch template version is the default version.
	//
	// * kernel-id - The kernel ID.
	//
	// *
	// ram-disk-id - The RAM disk ID.
	Filters []types.Filter

	// The ID of the launch template. To describe one or more versions of a specified
	// launch template, you must specify either the launch template ID or the launch
	// template name in the request. To describe all the latest or default launch
	// template versions in your account, you must omit this parameter.
	LaunchTemplateId *string

	// The name of the launch template. To describe one or more versions of a specified
	// launch template, you must specify either the launch template ID or the launch
	// template name in the request. To describe all the latest or default launch
	// template versions in your account, you must omit this parameter.
	LaunchTemplateName *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	MaxResults int32

	// The version number up to which to describe launch template versions.
	MaxVersion *string

	// The version number after which to describe launch template versions.
	MinVersion *string

	// The token to request the next page of results.
	NextToken *string

	// One or more versions of the launch template. Valid values depend on whether you
	// are describing a specified launch template (by ID or name) or all launch
	// templates in your account. To describe one or more versions of a specified
	// launch template, valid values are $Latest, $Default, and numbers. To describe
	// all launch templates in your account that are defined as the latest version, the
	// valid value is $Latest. To describe all launch templates in your account that
	// are defined as the default version, the valid value is $Default. You can specify
	// $Latest and $Default in the same call. You cannot specify numbers.
	Versions []string
}

type DescribeLaunchTemplateVersionsOutput struct {

	// Information about the launch template versions.
	LaunchTemplateVersions []types.LaunchTemplateVersion

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLaunchTemplateVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLaunchTemplateVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLaunchTemplateVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLaunchTemplateVersions(options.Region), middleware.Before); err != nil {
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

// DescribeLaunchTemplateVersionsAPIClient is a client that implements the
// DescribeLaunchTemplateVersions operation.
type DescribeLaunchTemplateVersionsAPIClient interface {
	DescribeLaunchTemplateVersions(context.Context, *DescribeLaunchTemplateVersionsInput, ...func(*Options)) (*DescribeLaunchTemplateVersionsOutput, error)
}

var _ DescribeLaunchTemplateVersionsAPIClient = (*Client)(nil)

// DescribeLaunchTemplateVersionsPaginatorOptions is the paginator options for
// DescribeLaunchTemplateVersions
type DescribeLaunchTemplateVersionsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLaunchTemplateVersionsPaginator is a paginator for
// DescribeLaunchTemplateVersions
type DescribeLaunchTemplateVersionsPaginator struct {
	options   DescribeLaunchTemplateVersionsPaginatorOptions
	client    DescribeLaunchTemplateVersionsAPIClient
	params    *DescribeLaunchTemplateVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLaunchTemplateVersionsPaginator returns a new
// DescribeLaunchTemplateVersionsPaginator
func NewDescribeLaunchTemplateVersionsPaginator(client DescribeLaunchTemplateVersionsAPIClient, params *DescribeLaunchTemplateVersionsInput, optFns ...func(*DescribeLaunchTemplateVersionsPaginatorOptions)) *DescribeLaunchTemplateVersionsPaginator {
	options := DescribeLaunchTemplateVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeLaunchTemplateVersionsInput{}
	}

	return &DescribeLaunchTemplateVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLaunchTemplateVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeLaunchTemplateVersions page.
func (p *DescribeLaunchTemplateVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLaunchTemplateVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeLaunchTemplateVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLaunchTemplateVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLaunchTemplateVersions",
	}
}
