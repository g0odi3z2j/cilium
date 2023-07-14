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

// Information about one or more Traffic Mirror targets.
func (c *Client) DescribeTrafficMirrorTargets(ctx context.Context, params *DescribeTrafficMirrorTargetsInput, optFns ...func(*Options)) (*DescribeTrafficMirrorTargetsOutput, error) {
	if params == nil {
		params = &DescribeTrafficMirrorTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrafficMirrorTargets", params, optFns, c.addOperationDescribeTrafficMirrorTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrafficMirrorTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrafficMirrorTargetsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - description : The Traffic Mirror target description.
	//   - network-interface-id : The ID of the Traffic Mirror session network
	//   interface.
	//   - network-load-balancer-arn : The Amazon Resource Name (ARN) of the Network
	//   Load Balancer that is associated with the session.
	//   - owner-id : The ID of the account that owns the Traffic Mirror session.
	//   - traffic-mirror-target-id : The ID of the Traffic Mirror target.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the Traffic Mirror targets.
	TrafficMirrorTargetIds []string

	noSmithyDocumentSerde
}

type DescribeTrafficMirrorTargetsOutput struct {

	// The token to use to retrieve the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more Traffic Mirror targets.
	TrafficMirrorTargets []types.TrafficMirrorTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrafficMirrorTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTrafficMirrorTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTrafficMirrorTargets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrafficMirrorTargets(options.Region), middleware.Before); err != nil {
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

// DescribeTrafficMirrorTargetsAPIClient is a client that implements the
// DescribeTrafficMirrorTargets operation.
type DescribeTrafficMirrorTargetsAPIClient interface {
	DescribeTrafficMirrorTargets(context.Context, *DescribeTrafficMirrorTargetsInput, ...func(*Options)) (*DescribeTrafficMirrorTargetsOutput, error)
}

var _ DescribeTrafficMirrorTargetsAPIClient = (*Client)(nil)

// DescribeTrafficMirrorTargetsPaginatorOptions is the paginator options for
// DescribeTrafficMirrorTargets
type DescribeTrafficMirrorTargetsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTrafficMirrorTargetsPaginator is a paginator for
// DescribeTrafficMirrorTargets
type DescribeTrafficMirrorTargetsPaginator struct {
	options   DescribeTrafficMirrorTargetsPaginatorOptions
	client    DescribeTrafficMirrorTargetsAPIClient
	params    *DescribeTrafficMirrorTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTrafficMirrorTargetsPaginator returns a new
// DescribeTrafficMirrorTargetsPaginator
func NewDescribeTrafficMirrorTargetsPaginator(client DescribeTrafficMirrorTargetsAPIClient, params *DescribeTrafficMirrorTargetsInput, optFns ...func(*DescribeTrafficMirrorTargetsPaginatorOptions)) *DescribeTrafficMirrorTargetsPaginator {
	if params == nil {
		params = &DescribeTrafficMirrorTargetsInput{}
	}

	options := DescribeTrafficMirrorTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTrafficMirrorTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTrafficMirrorTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTrafficMirrorTargets page.
func (p *DescribeTrafficMirrorTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTrafficMirrorTargetsOutput, error) {
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

	result, err := p.client.DescribeTrafficMirrorTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTrafficMirrorTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTrafficMirrorTargets",
	}
}
