// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the default KMS key for EBS encryption by default for your account in
// this Region. Amazon Web Services creates a unique Amazon Web Services managed
// KMS key in each Region for use with encryption by default. If you change the
// default KMS key to a symmetric customer managed KMS key, it is used instead of
// the Amazon Web Services managed KMS key. To reset the default KMS key to the
// Amazon Web Services managed KMS key for EBS, use ResetEbsDefaultKmsKeyId .
// Amazon EBS does not support asymmetric KMS keys. If you delete or disable the
// customer managed KMS key that you specified for use with encryption by default,
// your instances will fail to launch. For more information, see Amazon EBS
// encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifyEbsDefaultKmsKeyId(ctx context.Context, params *ModifyEbsDefaultKmsKeyIdInput, optFns ...func(*Options)) (*ModifyEbsDefaultKmsKeyIdOutput, error) {
	if params == nil {
		params = &ModifyEbsDefaultKmsKeyIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyEbsDefaultKmsKeyId", params, optFns, c.addOperationModifyEbsDefaultKmsKeyIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyEbsDefaultKmsKeyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyEbsDefaultKmsKeyIdInput struct {

	// The identifier of the Key Management Service (KMS) KMS key to use for Amazon
	// EBS encryption. If this parameter is not specified, your KMS key for Amazon EBS
	// is used. If KmsKeyId is specified, the encrypted state must be true . You can
	// specify the KMS key using any of the following:
	//   - Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//   - Key alias. For example, alias/ExampleAlias.
	//   - Key ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//   - Alias ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	// Amazon Web Services authenticates the KMS key asynchronously. Therefore, if you
	// specify an ID, alias, or ARN that is not valid, the action can appear to
	// complete, but eventually fails. Amazon EBS does not support asymmetric KMS keys.
	//
	// This member is required.
	KmsKeyId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyEbsDefaultKmsKeyIdOutput struct {

	// The Amazon Resource Name (ARN) of the default KMS key for encryption by default.
	KmsKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyEbsDefaultKmsKeyIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyEbsDefaultKmsKeyId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyEbsDefaultKmsKeyId{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyEbsDefaultKmsKeyIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyEbsDefaultKmsKeyId(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opModifyEbsDefaultKmsKeyId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyEbsDefaultKmsKeyId",
	}
}
