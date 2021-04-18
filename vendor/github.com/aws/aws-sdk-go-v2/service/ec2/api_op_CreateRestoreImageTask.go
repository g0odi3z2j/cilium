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

// Starts a task that restores an AMI from an S3 object that was previously created
// by using CreateStoreImageTask
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html).
// To use this API, you must have the required permissions. For more information,
// see Permissions for storing and restoring AMIs using S3
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html#ami-s3-permissions)
// in the Amazon Elastic Compute Cloud User Guide. For more information, see Store
// and restore an AMI using S3
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateRestoreImageTask(ctx context.Context, params *CreateRestoreImageTaskInput, optFns ...func(*Options)) (*CreateRestoreImageTaskOutput, error) {
	if params == nil {
		params = &CreateRestoreImageTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRestoreImageTask", params, optFns, addOperationCreateRestoreImageTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRestoreImageTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRestoreImageTaskInput struct {

	// The name of the S3 bucket that contains the stored AMI object.
	//
	// This member is required.
	Bucket *string

	// The name of the stored AMI object in the bucket.
	//
	// This member is required.
	ObjectKey *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The name for the restored AMI. The name must be unique for AMIs in the Region
	// for this account. If you do not provide a name, the new AMI gets the same name
	// as the original AMI.
	Name *string

	// The tags to apply to the AMI and snapshots on restoration. You can tag the AMI,
	// the snapshots, or both.
	//
	// * To tag the AMI, the value for ResourceType must be
	// image.
	//
	// * To tag the snapshots, the value for ResourceType must be snapshot. The
	// same tag is applied to all of the snapshots that are created.
	TagSpecifications []types.TagSpecification
}

type CreateRestoreImageTaskOutput struct {

	// The AMI ID.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRestoreImageTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateRestoreImageTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateRestoreImageTask{}, middleware.After)
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
	if err = addOpCreateRestoreImageTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRestoreImageTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRestoreImageTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateRestoreImageTask",
	}
}
