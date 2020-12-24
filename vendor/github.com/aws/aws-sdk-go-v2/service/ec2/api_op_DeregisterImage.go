// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters the specified AMI. After you deregister an AMI, it can't be used to
// launch new instances; however, it doesn't affect any instances that you've
// already launched from the AMI. You'll continue to incur usage costs for those
// instances until you terminate them. When you deregister an Amazon EBS-backed
// AMI, it doesn't affect the snapshot that was created for the root volume of the
// instance during the AMI creation process. When you deregister an instance
// store-backed AMI, it doesn't affect the files that you uploaded to Amazon S3
// when you created the AMI.
func (c *Client) DeregisterImage(ctx context.Context, params *DeregisterImageInput, optFns ...func(*Options)) (*DeregisterImageOutput, error) {
	if params == nil {
		params = &DeregisterImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterImage", params, optFns, addOperationDeregisterImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeregisterImage.
type DeregisterImageInput struct {

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type DeregisterImageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeregisterImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterImage{}, middleware.After)
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
	if err = addOpDeregisterImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterImage",
	}
}
