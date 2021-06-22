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

// Restores an Elastic IP address that was previously moved to the EC2-VPC platform
// back to the EC2-Classic platform. You cannot move an Elastic IP address that was
// originally allocated for use in EC2-VPC. The Elastic IP address must not be
// associated with an instance or network interface.
func (c *Client) RestoreAddressToClassic(ctx context.Context, params *RestoreAddressToClassicInput, optFns ...func(*Options)) (*RestoreAddressToClassicOutput, error) {
	if params == nil {
		params = &RestoreAddressToClassicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreAddressToClassic", params, optFns, addOperationRestoreAddressToClassicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreAddressToClassicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreAddressToClassicInput struct {

	// The Elastic IP address.
	//
	// This member is required.
	PublicIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type RestoreAddressToClassicOutput struct {

	// The Elastic IP address.
	PublicIp *string

	// The move status for the IP address.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreAddressToClassicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRestoreAddressToClassic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRestoreAddressToClassic{}, middleware.After)
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
	if err = addOpRestoreAddressToClassicValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreAddressToClassic(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreAddressToClassic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RestoreAddressToClassic",
	}
}
