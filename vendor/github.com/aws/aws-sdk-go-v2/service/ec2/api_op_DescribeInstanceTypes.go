// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the details of the instance types that are offered in a location. The
// results can be filtered by the attributes of the instance types.
func (c *Client) DescribeInstanceTypes(ctx context.Context, params *DescribeInstanceTypesInput, optFns ...func(*Options)) (*DescribeInstanceTypesOutput, error) {
	if params == nil {
		params = &DescribeInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceTypes", params, optFns, c.addOperationDescribeInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceTypesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//   - auto-recovery-supported - Indicates whether Amazon CloudWatch action based
	//   recovery is supported ( true | false ).
	//   - bare-metal - Indicates whether it is a bare metal instance type ( true |
	//   false ).
	//   - burstable-performance-supported - Indicates whether the instance type is a
	//   burstable performance T instance type ( true | false ).
	//   - current-generation - Indicates whether this instance type is the latest
	//   generation instance type of an instance family ( true | false ).
	//   - ebs-info.ebs-optimized-info.baseline-bandwidth-in-mbps - The baseline
	//   bandwidth performance for an EBS-optimized instance type, in Mbps.
	//   - ebs-info.ebs-optimized-info.baseline-iops - The baseline input/output
	//   storage operations per second for an EBS-optimized instance type.
	//   - ebs-info.ebs-optimized-info.baseline-throughput-in-mbps - The baseline
	//   throughput performance for an EBS-optimized instance type, in MB/s.
	//   - ebs-info.ebs-optimized-info.maximum-bandwidth-in-mbps - The maximum
	//   bandwidth performance for an EBS-optimized instance type, in Mbps.
	//   - ebs-info.ebs-optimized-info.maximum-iops - The maximum input/output storage
	//   operations per second for an EBS-optimized instance type.
	//   - ebs-info.ebs-optimized-info.maximum-throughput-in-mbps - The maximum
	//   throughput performance for an EBS-optimized instance type, in MB/s.
	//   - ebs-info.ebs-optimized-support - Indicates whether the instance type is
	//   EBS-optimized ( supported | unsupported | default ).
	//   - ebs-info.encryption-support - Indicates whether EBS encryption is supported (
	//   supported | unsupported ).
	//   - ebs-info.nvme-support - Indicates whether non-volatile memory express (NVMe)
	//   is supported for EBS volumes ( required | supported | unsupported ).
	//   - free-tier-eligible - Indicates whether the instance type is eligible to use
	//   in the free tier ( true | false ).
	//   - hibernation-supported - Indicates whether On-Demand hibernation is supported
	//   ( true | false ).
	//   - hypervisor - The hypervisor ( nitro | xen ).
	//   - instance-storage-info.disk.count - The number of local disks.
	//   - instance-storage-info.disk.size-in-gb - The storage size of each instance
	//   storage disk, in GB.
	//   - instance-storage-info.disk.type - The storage technology for the local
	//   instance storage disks ( hdd | ssd ).
	//   - instance-storage-info.encryption-support - Indicates whether data is
	//   encrypted at rest ( required | supported | unsupported ).
	//   - instance-storage-info.nvme-support - Indicates whether non-volatile memory
	//   express (NVMe) is supported for instance store ( required | supported |
	//   unsupported ).
	//   - instance-storage-info.total-size-in-gb - The total amount of storage
	//   available from all local instance storage, in GB.
	//   - instance-storage-supported - Indicates whether the instance type has local
	//   instance storage ( true | false ).
	//   - instance-type - The instance type (for example c5.2xlarge or c5*).
	//   - memory-info.size-in-mib - The memory size.
	//   - network-info.efa-info.maximum-efa-interfaces - The maximum number of Elastic
	//   Fabric Adapters (EFAs) per instance.
	//   - network-info.efa-supported - Indicates whether the instance type supports
	//   Elastic Fabric Adapter (EFA) ( true | false ).
	//   - network-info.ena-support - Indicates whether Elastic Network Adapter (ENA)
	//   is supported or required ( required | supported | unsupported ).
	//   - network-info.encryption-in-transit-supported - Indicates whether the
	//   instance type automatically encrypts in-transit traffic between instances (
	//   true | false ).
	//   - network-info.ipv4-addresses-per-interface - The maximum number of private
	//   IPv4 addresses per network interface.
	//   - network-info.ipv6-addresses-per-interface - The maximum number of private
	//   IPv6 addresses per network interface.
	//   - network-info.ipv6-supported - Indicates whether the instance type supports
	//   IPv6 ( true | false ).
	//   - network-info.maximum-network-cards - The maximum number of network cards per
	//   instance.
	//   - network-info.maximum-network-interfaces - The maximum number of network
	//   interfaces per instance.
	//   - network-info.network-performance - The network performance (for example, "25
	//   Gigabit").
	//   - nitro-enclaves-support - Indicates whether Nitro Enclaves is supported (
	//   supported | unsupported ).
	//   - nitro-tpm-support - Indicates whether NitroTPM is supported ( supported |
	//   unsupported ).
	//   - nitro-tpm-info.supported-versions - The supported NitroTPM version ( 2.0 ).
	//   - processor-info.supported-architecture - The CPU architecture ( arm64 | i386
	//   | x86_64 ).
	//   - processor-info.sustained-clock-speed-in-ghz - The CPU clock speed, in GHz.
	//   - processor-info.supported-features - The supported CPU features ( amd-sev-snp
	//   ).
	//   - supported-boot-mode - The boot mode ( legacy-bios | uefi ).
	//   - supported-root-device-type - The root device type ( ebs | instance-store ).
	//   - supported-usage-class - The usage class ( on-demand | spot ).
	//   - supported-virtualization-type - The virtualization type ( hvm | paravirtual
	//   ).
	//   - vcpu-info.default-cores - The default number of cores for the instance type.
	//   - vcpu-info.default-threads-per-core - The default number of threads per core
	//   for the instance type.
	//   - vcpu-info.default-vcpus - The default number of vCPUs for the instance type.
	//   - vcpu-info.valid-cores - The number of cores that can be configured for the
	//   instance type.
	//   - vcpu-info.valid-threads-per-core - The number of threads per core that can
	//   be configured for the instance type. For example, "1" or "1,2".
	Filters []types.Filter

	// The instance types. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon EC2 User Guide.
	InstanceTypes []types.InstanceType

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

type DescribeInstanceTypesOutput struct {

	// The instance type. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon EC2 User Guide.
	InstanceTypes []types.InstanceTypeInfo

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceTypes{}, middleware.After)
	if err != nil {
		return err
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
	if err = addDescribeInstanceTypesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceTypes(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeInstanceTypesAPIClient is a client that implements the
// DescribeInstanceTypes operation.
type DescribeInstanceTypesAPIClient interface {
	DescribeInstanceTypes(context.Context, *DescribeInstanceTypesInput, ...func(*Options)) (*DescribeInstanceTypesOutput, error)
}

var _ DescribeInstanceTypesAPIClient = (*Client)(nil)

// DescribeInstanceTypesPaginatorOptions is the paginator options for
// DescribeInstanceTypes
type DescribeInstanceTypesPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceTypesPaginator is a paginator for DescribeInstanceTypes
type DescribeInstanceTypesPaginator struct {
	options   DescribeInstanceTypesPaginatorOptions
	client    DescribeInstanceTypesAPIClient
	params    *DescribeInstanceTypesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceTypesPaginator returns a new DescribeInstanceTypesPaginator
func NewDescribeInstanceTypesPaginator(client DescribeInstanceTypesAPIClient, params *DescribeInstanceTypesInput, optFns ...func(*DescribeInstanceTypesPaginatorOptions)) *DescribeInstanceTypesPaginator {
	if params == nil {
		params = &DescribeInstanceTypesInput{}
	}

	options := DescribeInstanceTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceTypes page.
func (p *DescribeInstanceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceTypesOutput, error) {
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

	result, err := p.client.DescribeInstanceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceTypes",
	}
}

type opDescribeInstanceTypesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeInstanceTypesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeInstanceTypesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeInstanceTypesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeInstanceTypesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
