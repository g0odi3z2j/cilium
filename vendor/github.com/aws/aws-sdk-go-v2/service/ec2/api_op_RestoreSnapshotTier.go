// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores an archived Amazon EBS snapshot for use temporarily or permanently, or
// modifies the restore period or restore type for a snapshot that was previously
// temporarily restored. For more information see  Restore an archived snapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-snapshot-archiving.html#restore-archived-snapshot)
// and  modify the restore period or restore type for a temporarily restored
// snapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-snapshot-archiving.html#modify-temp-restore-period)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) RestoreSnapshotTier(ctx context.Context, params *RestoreSnapshotTierInput, optFns ...func(*Options)) (*RestoreSnapshotTierOutput, error) {
	if params == nil {
		params = &RestoreSnapshotTierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreSnapshotTier", params, optFns, c.addOperationRestoreSnapshotTierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreSnapshotTierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreSnapshotTierInput struct {

	// The ID of the snapshot to restore.
	//
	// This member is required.
	SnapshotId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Indicates whether to permanently restore an archived snapshot. To permanently
	// restore an archived snapshot, specify true and omit the
	// RestoreSnapshotTierRequest$TemporaryRestoreDays parameter.
	PermanentRestore *bool

	// Specifies the number of days for which to temporarily restore an archived
	// snapshot. Required for temporary restores only. The snapshot will be
	// automatically re-archived after this period. To temporarily restore an archived
	// snapshot, specify the number of days and omit the PermanentRestore parameter or
	// set it to false.
	TemporaryRestoreDays *int32

	noSmithyDocumentSerde
}

type RestoreSnapshotTierOutput struct {

	// Indicates whether the snapshot is permanently restored. true indicates a
	// permanent restore. false indicates a temporary restore.
	IsPermanentRestore *bool

	// For temporary restores only. The number of days for which the archived snapshot
	// is temporarily restored.
	RestoreDuration *int32

	// The date and time when the snapshot restore process started.
	RestoreStartTime *time.Time

	// The ID of the snapshot.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreSnapshotTierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRestoreSnapshotTier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRestoreSnapshotTier{}, middleware.After)
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
	if err = addOpRestoreSnapshotTierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreSnapshotTier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreSnapshotTier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RestoreSnapshotTier",
	}
}
