// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Connect peer for a specified transit gateway Connect attachment
// between a transit gateway and an appliance. The peer address and transit gateway
// address must be the same IP address family (IPv4 or IPv6). For more information,
// see Connect peers
// (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html#tgw-connect-peer)
// in the Transit Gateways Guide.
func (c *Client) CreateTransitGatewayConnectPeer(ctx context.Context, params *CreateTransitGatewayConnectPeerInput, optFns ...func(*Options)) (*CreateTransitGatewayConnectPeerOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayConnectPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayConnectPeer", params, optFns, addOperationCreateTransitGatewayConnectPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayConnectPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayConnectPeerInput struct {

	// The range of inside IP addresses that are used for BGP peering. You must specify
	// a size /29 IPv4 CIDR block from the 169.254.0.0/16 range. The first address from
	// the range must be configured on the appliance as the BGP IP address. You can
	// also optionally specify a size /125 IPv6 CIDR block from the fd00::/8 range.
	//
	// This member is required.
	InsideCidrBlocks []string

	// The peer IP address (GRE outer IP address) on the appliance side of the Connect
	// peer.
	//
	// This member is required.
	PeerAddress *string

	// The ID of the Connect attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The BGP options for the Connect peer.
	BgpOptions *types.TransitGatewayConnectRequestBgpOptions

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The tags to apply to the Connect peer.
	TagSpecifications []types.TagSpecification

	// The peer IP address (GRE outer IP address) on the transit gateway side of the
	// Connect peer, which must be specified from a transit gateway CIDR block. If not
	// specified, Amazon automatically assigns the first available IP address from the
	// transit gateway CIDR block.
	TransitGatewayAddress *string
}

type CreateTransitGatewayConnectPeerOutput struct {

	// Information about the Connect peer.
	TransitGatewayConnectPeer *types.TransitGatewayConnectPeer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTransitGatewayConnectPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayConnectPeer{}, middleware.After)
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
	if err = addOpCreateTransitGatewayConnectPeerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayConnectPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGatewayConnectPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayConnectPeer",
	}
}
