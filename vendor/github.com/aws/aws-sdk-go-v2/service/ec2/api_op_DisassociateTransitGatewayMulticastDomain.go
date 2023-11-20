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

// Disassociates the specified subnets from the transit gateway multicast domain.
func (c *Client) DisassociateTransitGatewayMulticastDomain(ctx context.Context, params *DisassociateTransitGatewayMulticastDomainInput, optFns ...func(*Options)) (*DisassociateTransitGatewayMulticastDomainOutput, error) {
	if params == nil {
		params = &DisassociateTransitGatewayMulticastDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTransitGatewayMulticastDomain", params, optFns, c.addOperationDisassociateTransitGatewayMulticastDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTransitGatewayMulticastDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTransitGatewayMulticastDomainInput struct {

	// The IDs of the subnets;
	//
	// This member is required.
	SubnetIds []string

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The ID of the transit gateway multicast domain.
	//
	// This member is required.
	TransitGatewayMulticastDomainId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DisassociateTransitGatewayMulticastDomainOutput struct {

	// Information about the association.
	Associations *types.TransitGatewayMulticastDomainAssociations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateTransitGatewayMulticastDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateTransitGatewayMulticastDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateTransitGatewayMulticastDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateTransitGatewayMulticastDomain"); err != nil {
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
	if err = addOpDisassociateTransitGatewayMulticastDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTransitGatewayMulticastDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateTransitGatewayMulticastDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateTransitGatewayMulticastDomain",
	}
}
