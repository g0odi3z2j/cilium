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

// Modifies a volume attribute. By default, all I/O operations for the volume are
// suspended when the data on the volume is determined to be potentially
// inconsistent, to prevent undetectable, latent data corruption. The I/O access to
// the volume can be resumed by first enabling I/O access and then checking the
// data consistency on your volume. You can change the default behavior to resume
// I/O operations. We recommend that you change this only for boot volumes or for
// volumes that are stateless or disposable.
func (c *Client) ModifyVolumeAttribute(ctx context.Context, params *ModifyVolumeAttributeInput, optFns ...func(*Options)) (*ModifyVolumeAttributeOutput, error) {
	if params == nil {
		params = &ModifyVolumeAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVolumeAttribute", params, optFns, addOperationModifyVolumeAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVolumeAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVolumeAttributeInput struct {

	// The ID of the volume.
	//
	// This member is required.
	VolumeId *string

	// Indicates whether the volume should be auto-enabled for I/O operations.
	AutoEnableIO *types.AttributeBooleanValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type ModifyVolumeAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVolumeAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVolumeAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVolumeAttribute{}, middleware.After)
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
	if err = addOpModifyVolumeAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVolumeAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVolumeAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVolumeAttribute",
	}
}
