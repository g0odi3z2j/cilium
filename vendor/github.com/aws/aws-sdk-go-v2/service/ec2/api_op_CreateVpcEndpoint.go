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

// Creates a VPC endpoint. A VPC endpoint provides a private connection between
// the specified VPC and the specified endpoint service. You can use an endpoint
// service provided by Amazon Web Services, an Amazon Web Services Marketplace
// Partner, or another Amazon Web Services account. For more information, see the
// Amazon Web Services PrivateLink User Guide (https://docs.aws.amazon.com/vpc/latest/privatelink/)
// .
func (c *Client) CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) {
	if params == nil {
		params = &CreateVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcEndpoint", params, optFns, c.addOperationCreateVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcEndpointInput struct {

	// The name of the endpoint service.
	//
	// This member is required.
	ServiceName *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The DNS options for the endpoint.
	DnsOptions *types.DnsOptionsSpecification

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The IP address type for the endpoint.
	IpAddressType types.IpAddressType

	// (Interface and gateway endpoints) A policy to attach to the endpoint that
	// controls access to the service. The policy must be in valid JSON format. If this
	// parameter is not specified, we attach a default policy that allows full access
	// to the service.
	PolicyDocument *string

	// (Interface endpoint) Indicates whether to associate a private hosted zone with
	// the specified VPC. The private hosted zone contains a record set for the default
	// public DNS name for the service for the Region (for example,
	// kinesis.us-east-1.amazonaws.com ), which resolves to the private IP addresses of
	// the endpoint network interfaces in the VPC. This enables you to make requests to
	// the default public DNS name for the service instead of the public DNS names that
	// are automatically generated by the VPC endpoint service. To use a private hosted
	// zone, you must set the following VPC attributes to true : enableDnsHostnames
	// and enableDnsSupport . Use ModifyVpcAttribute to set the VPC attributes.
	// Default: true
	PrivateDnsEnabled *bool

	// (Gateway endpoint) The route table IDs.
	RouteTableIds []string

	// (Interface endpoint) The IDs of the security groups to associate with the
	// endpoint network interfaces. If this parameter is not specified, we use the
	// default security group for the VPC.
	SecurityGroupIds []string

	// The subnet configurations for the endpoint.
	SubnetConfigurations []types.SubnetConfiguration

	// (Interface and Gateway Load Balancer endpoints) The IDs of the subnets in which
	// to create endpoint network interfaces. For a Gateway Load Balancer endpoint, you
	// can specify only one subnet.
	SubnetIds []string

	// The tags to associate with the endpoint.
	TagSpecifications []types.TagSpecification

	// The type of endpoint. Default: Gateway
	VpcEndpointType types.VpcEndpointType

	noSmithyDocumentSerde
}

type CreateVpcEndpointOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// Information about the endpoint.
	VpcEndpoint *types.VpcEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVpcEndpoint"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcEndpoint(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVpcEndpoint",
	}
}
