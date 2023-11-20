// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a static route associated with a VPN connection between an existing
// virtual private gateway and a VPN customer gateway. The static route allows
// traffic to be routed from the virtual private gateway to the VPN customer
// gateway. For more information, see Amazon Web Services Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the Amazon Web Services Site-to-Site VPN User Guide.
func (c *Client) CreateVpnConnectionRoute(ctx context.Context, params *CreateVpnConnectionRouteInput, optFns ...func(*Options)) (*CreateVpnConnectionRouteOutput, error) {
	if params == nil {
		params = &CreateVpnConnectionRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpnConnectionRoute", params, optFns, c.addOperationCreateVpnConnectionRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpnConnectionRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateVpnConnectionRoute.
type CreateVpnConnectionRouteInput struct {

	// The CIDR block associated with the local subnet of the customer network.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	noSmithyDocumentSerde
}

type CreateVpnConnectionRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpnConnectionRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpnConnectionRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpnConnectionRoute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVpnConnectionRoute"); err != nil {
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
	if err = addOpCreateVpnConnectionRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpnConnectionRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpnConnectionRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVpnConnectionRoute",
	}
}
