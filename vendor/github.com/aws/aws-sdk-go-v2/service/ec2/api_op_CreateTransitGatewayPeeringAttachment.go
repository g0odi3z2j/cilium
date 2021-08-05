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

// Requests a transit gateway peering attachment between the specified transit
// gateway (requester) and a peer transit gateway (accepter). The transit gateways
// must be in different Regions. The peer transit gateway can be in your account or
// a different AWS account. After you create the peering attachment, the owner of
// the accepter transit gateway must accept the attachment request.
func (c *Client) CreateTransitGatewayPeeringAttachment(ctx context.Context, params *CreateTransitGatewayPeeringAttachmentInput, optFns ...func(*Options)) (*CreateTransitGatewayPeeringAttachmentOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayPeeringAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayPeeringAttachment", params, optFns, c.addOperationCreateTransitGatewayPeeringAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayPeeringAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayPeeringAttachmentInput struct {

	// The AWS account ID of the owner of the peer transit gateway.
	//
	// This member is required.
	PeerAccountId *string

	// The Region where the peer transit gateway is located.
	//
	// This member is required.
	PeerRegion *string

	// The ID of the peer transit gateway with which to create the peering attachment.
	//
	// This member is required.
	PeerTransitGatewayId *string

	// The ID of the transit gateway.
	//
	// This member is required.
	TransitGatewayId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to apply to the transit gateway peering attachment.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateTransitGatewayPeeringAttachmentOutput struct {

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransitGatewayPeeringAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayPeeringAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayPeeringAttachment{}, middleware.After)
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
	if err = addOpCreateTransitGatewayPeeringAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayPeeringAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGatewayPeeringAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayPeeringAttachment",
	}
}
