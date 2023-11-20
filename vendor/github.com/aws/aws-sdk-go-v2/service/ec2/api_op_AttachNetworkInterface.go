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

// Attaches a network interface to an instance.
func (c *Client) AttachNetworkInterface(ctx context.Context, params *AttachNetworkInterfaceInput, optFns ...func(*Options)) (*AttachNetworkInterfaceOutput, error) {
	if params == nil {
		params = &AttachNetworkInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachNetworkInterface", params, optFns, c.addOperationAttachNetworkInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachNetworkInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for AttachNetworkInterface.
type AttachNetworkInterfaceInput struct {

	// The index of the device for the network interface attachment.
	//
	// This member is required.
	DeviceIndex *int32

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Configures ENA Express for the network interface that this action attaches to
	// the instance.
	EnaSrdSpecification *types.EnaSrdSpecification

	// The index of the network card. Some instance types support multiple network
	// cards. The primary network interface must be assigned to network card index 0.
	// The default is network card index 0.
	NetworkCardIndex *int32

	noSmithyDocumentSerde
}

// Contains the output of AttachNetworkInterface.
type AttachNetworkInterfaceOutput struct {

	// The ID of the network interface attachment.
	AttachmentId *string

	// The index of the network card.
	NetworkCardIndex *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachNetworkInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAttachNetworkInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAttachNetworkInterface{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AttachNetworkInterface"); err != nil {
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
	if err = addOpAttachNetworkInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachNetworkInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachNetworkInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AttachNetworkInterface",
	}
}
