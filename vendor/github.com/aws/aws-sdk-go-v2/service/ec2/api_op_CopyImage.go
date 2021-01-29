// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates the copy of an AMI from the specified source Region to the current
// Region. You specify the destination Region by using its endpoint when making the
// request. Copies of encrypted backing snapshots for the AMI are encrypted. Copies
// of unencrypted backing snapshots remain unencrypted, unless you set Encrypted
// during the copy operation. You cannot create an unencrypted copy of an encrypted
// backing snapshot. For more information about the prerequisites and limits when
// copying an AMI, see Copying an AMI
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) CopyImage(ctx context.Context, params *CopyImageInput, optFns ...func(*Options)) (*CopyImageOutput, error) {
	if params == nil {
		params = &CopyImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyImage", params, optFns, addOperationCopyImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CopyImage.
type CopyImageInput struct {

	// The name of the new AMI in the destination Region.
	//
	// This member is required.
	Name *string

	// The ID of the AMI to copy.
	//
	// This member is required.
	SourceImageId *string

	// The name of the Region that contains the AMI to copy.
	//
	// This member is required.
	SourceRegion *string

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ClientToken *string

	// A description for the new AMI in the destination Region.
	Description *string

	DestinationOutpostArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Specifies whether the destination snapshots of the copied image should be
	// encrypted. You can encrypt a copy of an unencrypted snapshot, but you cannot
	// create an unencrypted copy of an encrypted snapshot. The default CMK for EBS is
	// used unless you specify a non-default AWS Key Management Service (AWS KMS) CMK
	// using KmsKeyId. For more information, see Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted bool

	// The identifier of the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating encrypted volumes. If this parameter is
	// not specified, your AWS managed CMK for EBS is used. If you specify a CMK, you
	// must also set the encrypted state to true. You can specify a CMK using any of
	// the following:
	//
	// * Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// *
	// Key alias. For example, alias/ExampleAlias.
	//
	// * Key ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// *
	// Alias ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS authenticates the
	// CMK asynchronously. Therefore, if you specify an identifier that is not valid,
	// the action can appear to complete, but eventually fails. The specified CMK must
	// exist in the destination Region. Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string
}

// Contains the output of CopyImage.
type CopyImageOutput struct {

	// The ID of the new AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCopyImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCopyImage{}, middleware.After)
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
	if err = addOpCopyImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopyImage",
	}
}
