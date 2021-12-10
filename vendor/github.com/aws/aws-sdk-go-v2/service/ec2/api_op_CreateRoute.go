// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a route in a route table within a VPC. You must specify one of the
// following targets: internet gateway or virtual private gateway, NAT instance,
// NAT gateway, VPC peering connection, network interface, egress-only internet
// gateway, or transit gateway. When determining how to route traffic, we use the
// route with the most specific match. For example, traffic is destined for the
// IPv4 address 192.0.2.3, and the route table includes the following two IPv4
// routes:
//
// * 192.0.2.0/24 (goes to some target A)
//
// * 192.0.2.0/28 (goes to some
// target B)
//
// Both routes apply to the traffic destined for 192.0.2.3. However, the
// second route in the list covers a smaller number of IP addresses and is
// therefore more specific, so we use that route to determine where to target the
// traffic. For more information about route tables, see Route tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) {
	if params == nil {
		params = &CreateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoute", params, optFns, c.addOperationCreateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRouteInput struct {

	// The ID of the route table for the route.
	//
	// This member is required.
	RouteTableId *string

	// The ID of the carrier gateway. You can only use this option when the VPC
	// contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId *string

	// The Amazon Resource Name (ARN) of the core network.
	CoreNetworkArn *string

	// The IPv4 CIDR address block used for the destination match. Routing decisions
	// are based on the most specific match. We modify the specified CIDR block to its
	// canonical form; for example, if you specify 100.68.0.18/18, we modify it to
	// 100.68.0.0/18.
	DestinationCidrBlock *string

	// The IPv6 CIDR block used for the destination match. Routing decisions are based
	// on the most specific match.
	DestinationIpv6CidrBlock *string

	// The ID of a prefix list used for the destination match.
	DestinationPrefixListId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string

	// The ID of an internet gateway or virtual private gateway attached to your VPC.
	GatewayId *string

	// The ID of a NAT instance in your VPC. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId *string

	// The ID of the local gateway.
	LocalGatewayId *string

	// [IPv4 traffic only] The ID of a NAT gateway.
	NatGatewayId *string

	// The ID of a network interface.
	NetworkInterfaceId *string

	// The ID of a transit gateway.
	TransitGatewayId *string

	// The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.
	VpcEndpointId *string

	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string

	noSmithyDocumentSerde
}

type CreateRouteOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateRoute{}, middleware.After)
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
	if err = addOpCreateRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateRoute",
	}
}
