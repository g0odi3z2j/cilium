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

// Imports a disk into an EBS snapshot.
func (c *Client) ImportSnapshot(ctx context.Context, params *ImportSnapshotInput, optFns ...func(*Options)) (*ImportSnapshotOutput, error) {
	if params == nil {
		params = &ImportSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportSnapshot", params, optFns, addOperationImportSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSnapshotInput struct {

	// The client-specific data.
	ClientData *types.ClientData

	// Token to enable idempotency for VM import requests.
	ClientToken *string

	// The description string for the import snapshot task.
	Description *string

	// Information about the disk container.
	DiskContainer *types.SnapshotDiskContainer

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Specifies whether the destination snapshot of the imported image should be
	// encrypted. The default CMK for EBS is used unless you specify a non-default AWS
	// Key Management Service (AWS KMS) CMK using KmsKeyId. For more information, see
	// Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted bool

	// An identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating the encrypted snapshot. This parameter is
	// only required if you want to use a non-default CMK; if this parameter is not
	// specified, the default CMK for EBS is used. If a KmsKeyId is specified, the
	// Encrypted flag must also be set. The CMK identifier may be provided in any of
	// the following formats:
	//
	// * Key ID
	//
	// * Key alias. The alias ARN contains the
	// arn:aws:kms namespace, followed by the Region of the CMK, the AWS account ID of
	// the CMK owner, the alias namespace, and then the CMK alias. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// * ARN using key ID. The
	// ID ARN contains the arn:aws:kms namespace, followed by the Region of the CMK,
	// the AWS account ID of the CMK owner, the key namespace, and then the CMK ID. For
	// example,
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	// *
	// ARN using key alias. The alias ARN contains the arn:aws:kms namespace, followed
	// by the Region of the CMK, the AWS account ID of the CMK owner, the alias
	// namespace, and then the CMK alias. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS parses KmsKeyId
	// asynchronously, meaning that the action you call may appear to complete even
	// though you provided an invalid identifier. This action will eventually report
	// failure. The specified CMK must exist in the Region that the snapshot is being
	// copied to. Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string

	// The tags to apply to the snapshot being imported.
	TagSpecifications []types.TagSpecification
}

type ImportSnapshotOutput struct {

	// A description of the import snapshot task.
	Description *string

	// The ID of the import snapshot task.
	ImportTaskId *string

	// Information about the import snapshot task.
	SnapshotTaskDetail *types.SnapshotTaskDetail

	// Any tags assigned to the snapshot being imported.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpImportSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpImportSnapshot{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ImportSnapshot",
	}
}
