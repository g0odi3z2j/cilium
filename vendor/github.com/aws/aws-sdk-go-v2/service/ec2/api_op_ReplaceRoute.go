// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces an existing route within a route table in a VPC. You must specify
// either a destination CIDR block or a prefix list ID. You must also specify
// exactly one of the resources from the parameter list, or reset the local route
// to its default target. For more information, see Route tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) ReplaceRoute(ctx context.Context, params *ReplaceRouteInput, optFns ...func(*Options)) (*ReplaceRouteOutput, error) {
	if params == nil {
		params = &ReplaceRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceRoute", params, optFns, c.addOperationReplaceRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceRouteInput struct {

	// The ID of the route table.
	//
	// This member is required.
	RouteTableId *string

	// [IPv4 traffic only] The ID of a carrier gateway.
	CarrierGatewayId *string

	// The Amazon Resource Name (ARN) of the core network.
	CoreNetworkArn *string

	// The IPv4 CIDR address block used for the destination match. The value that you
	// provide must match the CIDR of an existing route in the table.
	DestinationCidrBlock *string

	// The IPv6 CIDR address block used for the destination match. The value that you
	// provide must match the CIDR of an existing route in the table.
	DestinationIpv6CidrBlock *string

	// The ID of the prefix list for the route.
	DestinationPrefixListId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string

	// The ID of an internet gateway or virtual private gateway.
	GatewayId *string

	// The ID of a NAT instance in your VPC.
	InstanceId *string

	// The ID of the local gateway.
	LocalGatewayId *string

	// Specifies whether to reset the local route to its default target ( local ).
	LocalTarget *bool

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

type ReplaceRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplaceRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReplaceRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceRoute{}, middleware.After)
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
	if err = addOpReplaceRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplaceRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceRoute",
	}
}
