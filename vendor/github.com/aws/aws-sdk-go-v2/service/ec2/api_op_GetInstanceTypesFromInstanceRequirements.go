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

// Returns a list of instance types with the specified instance attributes. You
// can use the response to preview the instance types without launching instances.
// Note that the response does not consider capacity. When you specify multiple
// parameters, you get instance types that satisfy all of the specified parameters.
// If you specify multiple values for a parameter, you get instance types that
// satisfy any of the specified values. For more information, see Preview instance
// types with specified attributes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html#spotfleet-get-instance-types-from-instance-requirements)
// , Attribute-based instance type selection for EC2 Fleet (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html)
// , Attribute-based instance type selection for Spot Fleet (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html)
// , and Spot placement score (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html)
// in the Amazon EC2 User Guide, and Creating an Auto Scaling group using
// attribute-based instance type selection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) GetInstanceTypesFromInstanceRequirements(ctx context.Context, params *GetInstanceTypesFromInstanceRequirementsInput, optFns ...func(*Options)) (*GetInstanceTypesFromInstanceRequirementsOutput, error) {
	if params == nil {
		params = &GetInstanceTypesFromInstanceRequirementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceTypesFromInstanceRequirements", params, optFns, c.addOperationGetInstanceTypesFromInstanceRequirementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceTypesFromInstanceRequirementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceTypesFromInstanceRequirementsInput struct {

	// The processor architecture type.
	//
	// This member is required.
	ArchitectureTypes []types.ArchitectureType

	// The attributes required for the instance types.
	//
	// This member is required.
	InstanceRequirements *types.InstanceRequirementsRequest

	// The virtualization type.
	//
	// This member is required.
	VirtualizationTypes []types.VirtualizationType

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetInstanceTypesFromInstanceRequirementsOutput struct {

	// The instance types with the specified instance attributes.
	InstanceTypes []types.InstanceTypeInfoFromInstanceRequirements

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInstanceTypesFromInstanceRequirementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetInstanceTypesFromInstanceRequirements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetInstanceTypesFromInstanceRequirements{}, middleware.After)
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
	if err = addOpGetInstanceTypesFromInstanceRequirementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceTypesFromInstanceRequirements(options.Region), middleware.Before); err != nil {
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

// GetInstanceTypesFromInstanceRequirementsAPIClient is a client that implements
// the GetInstanceTypesFromInstanceRequirements operation.
type GetInstanceTypesFromInstanceRequirementsAPIClient interface {
	GetInstanceTypesFromInstanceRequirements(context.Context, *GetInstanceTypesFromInstanceRequirementsInput, ...func(*Options)) (*GetInstanceTypesFromInstanceRequirementsOutput, error)
}

var _ GetInstanceTypesFromInstanceRequirementsAPIClient = (*Client)(nil)

// GetInstanceTypesFromInstanceRequirementsPaginatorOptions is the paginator
// options for GetInstanceTypesFromInstanceRequirements
type GetInstanceTypesFromInstanceRequirementsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInstanceTypesFromInstanceRequirementsPaginator is a paginator for
// GetInstanceTypesFromInstanceRequirements
type GetInstanceTypesFromInstanceRequirementsPaginator struct {
	options   GetInstanceTypesFromInstanceRequirementsPaginatorOptions
	client    GetInstanceTypesFromInstanceRequirementsAPIClient
	params    *GetInstanceTypesFromInstanceRequirementsInput
	nextToken *string
	firstPage bool
}

// NewGetInstanceTypesFromInstanceRequirementsPaginator returns a new
// GetInstanceTypesFromInstanceRequirementsPaginator
func NewGetInstanceTypesFromInstanceRequirementsPaginator(client GetInstanceTypesFromInstanceRequirementsAPIClient, params *GetInstanceTypesFromInstanceRequirementsInput, optFns ...func(*GetInstanceTypesFromInstanceRequirementsPaginatorOptions)) *GetInstanceTypesFromInstanceRequirementsPaginator {
	if params == nil {
		params = &GetInstanceTypesFromInstanceRequirementsInput{}
	}

	options := GetInstanceTypesFromInstanceRequirementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInstanceTypesFromInstanceRequirementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInstanceTypesFromInstanceRequirementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInstanceTypesFromInstanceRequirements page.
func (p *GetInstanceTypesFromInstanceRequirementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInstanceTypesFromInstanceRequirementsOutput, error) {
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

	result, err := p.client.GetInstanceTypesFromInstanceRequirements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInstanceTypesFromInstanceRequirements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetInstanceTypesFromInstanceRequirements",
	}
}
