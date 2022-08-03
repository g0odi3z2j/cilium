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

// Create an IP address pool for Amazon VPC IP Address Manager (IPAM). In IPAM, a
// pool is a collection of contiguous IP addresses CIDRs. Pools enable you to
// organize your IP addresses according to your routing and security needs. For
// example, if you have separate routing and security needs for development and
// production applications, you can create a pool for each. For more information,
// see Create a top-level pool
// (https://docs.aws.amazon.com/vpc/latest/ipam/create-top-ipam.html) in the Amazon
// VPC IPAM User Guide.
func (c *Client) CreateIpamPool(ctx context.Context, params *CreateIpamPoolInput, optFns ...func(*Options)) (*CreateIpamPoolOutput, error) {
	if params == nil {
		params = &CreateIpamPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIpamPool", params, optFns, c.addOperationCreateIpamPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIpamPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIpamPoolInput struct {

	// The IP protocol assigned to this IPAM pool. You must choose either IPv4 or IPv6
	// protocol for a pool.
	//
	// This member is required.
	AddressFamily types.AddressFamily

	// The ID of the scope in which you would like to create the IPAM pool.
	//
	// This member is required.
	IpamScopeId *string

	// The default netmask length for allocations added to this pool. If, for example,
	// the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new
	// allocations will default to 10.0.0.0/16.
	AllocationDefaultNetmaskLength *int32

	// The maximum netmask length possible for CIDR allocations in this IPAM pool to be
	// compliant. The maximum netmask length must be greater than the minimum netmask
	// length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask
	// lengths for IPv6 addresses are 0 - 128.
	AllocationMaxNetmaskLength *int32

	// The minimum netmask length required for CIDR allocations in this IPAM pool to be
	// compliant. The minimum netmask length must be less than the maximum netmask
	// length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask
	// lengths for IPv6 addresses are 0 - 128.
	AllocationMinNetmaskLength *int32

	// Tags that are required for resources that use CIDRs from this IPAM pool.
	// Resources that do not have these tags will not be allowed to allocate space from
	// the pool. If the resources have their tags changed after they have allocated
	// space or if the allocation tagging requirements are changed on the pool, the
	// resource may be marked as noncompliant.
	AllocationResourceTags []types.RequestIpamResourceTag

	// If selected, IPAM will continuously look for resources within the CIDR range of
	// this pool and automatically import them as allocations into your IPAM. The CIDRs
	// that will be allocated for these resources must not already be allocated to
	// other resources in order for the import to succeed. IPAM will import a CIDR
	// regardless of its compliance with the pool's allocation rules, so a resource
	// might be imported and subsequently marked as noncompliant. If IPAM discovers
	// multiple CIDRs that overlap, IPAM will import the largest CIDR only. If IPAM
	// discovers multiple CIDRs with matching CIDRs, IPAM will randomly import one of
	// them only. A locale must be set on the pool for this feature to work.
	AutoImport *bool

	// Limits which service in Amazon Web Services that the pool can be used in. "ec2",
	// for example, allows users to use space for Elastic IP addresses and VPCs.
	AwsService types.IpamPoolAwsService

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// A description for the IPAM pool.
	Description *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// In IPAM, the locale is the Amazon Web Services Region where you want to make an
	// IPAM pool available for allocations. Only resources in the same Region as the
	// locale of the pool can get IP address allocations from the pool. You can only
	// allocate a CIDR for a VPC, for example, from an IPAM pool that shares a locale
	// with the VPC’s Region. Note that once you choose a Locale for a pool, you cannot
	// modify it. If you do not choose a locale, resources in Regions others than the
	// IPAM's home region cannot use CIDRs from this pool. Possible values: Any Amazon
	// Web Services Region, such as us-east-1.
	Locale *string

	// Determines if the pool is publicly advertisable. This option is not available
	// for pools with AddressFamily set to ipv4.
	PubliclyAdvertisable *bool

	// The ID of the source IPAM pool. Use this option to create a pool within an
	// existing pool. Note that the CIDR you provision for the pool within the source
	// pool must be available in the source pool's CIDR range.
	SourceIpamPoolId *string

	// The key/value combination of a tag assigned to the resource. Use the tag key in
	// the filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateIpamPoolOutput struct {

	// Information about the IPAM pool created.
	IpamPool *types.IpamPool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIpamPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateIpamPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateIpamPool{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateIpamPoolMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIpamPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIpamPool(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIpamPool struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIpamPool) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIpamPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIpamPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIpamPoolInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateIpamPoolMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIpamPool{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIpamPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateIpamPool",
	}
}
