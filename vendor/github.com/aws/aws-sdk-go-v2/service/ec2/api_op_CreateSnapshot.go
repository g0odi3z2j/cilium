// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a snapshot of an EBS volume and stores it in Amazon S3. You can use
// snapshots for backups, to make copies of EBS volumes, and to save data before
// shutting down an instance. You can create snapshots of volumes in a Region and
// volumes on an Outpost. If you create a snapshot of a volume in a Region, the
// snapshot must be stored in the same Region as the volume. If you create a
// snapshot of a volume on an Outpost, the snapshot can be stored on the same
// Outpost as the volume, or in the Region for that Outpost. When a snapshot is
// created, any Amazon Web Services Marketplace product codes that are associated
// with the source volume are propagated to the snapshot. You can take a snapshot
// of an attached volume that is in use. However, snapshots only capture data that
// has been written to your Amazon EBS volume at the time the snapshot command is
// issued; this might exclude any data that has been cached by any applications or
// the operating system. If you can pause any file systems on the volume long
// enough to take a snapshot, your snapshot should be complete. However, if you
// cannot pause all file writes to the volume, you should unmount the volume from
// within the instance, issue the snapshot command, and then remount the volume to
// ensure a consistent and complete snapshot. You may remount and use your volume
// while the snapshot status is pending . When you create a snapshot for an EBS
// volume that serves as a root device, we recommend that you stop the instance
// before taking the snapshot. Snapshots that are taken from encrypted volumes are
// automatically encrypted. Volumes that are created from encrypted snapshots are
// also automatically encrypted. Your encrypted volumes and any associated
// snapshots always remain protected. You can tag your snapshots during creation.
// For more information, see Tag your Amazon EC2 resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information, see
// Amazon Elastic Block Store (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html)
// and Amazon EBS encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) {
	if params == nil {
		params = &CreateSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshot", params, optFns, c.addOperationCreateSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotInput struct {

	// The ID of the Amazon EBS volume.
	//
	// This member is required.
	VolumeId *string

	// A description for the snapshot.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Amazon Resource Name (ARN) of the Outpost on which to create a local
	// snapshot.
	//   - To create a snapshot of a volume in a Region, omit this parameter. The
	//   snapshot is created in the same Region as the volume.
	//   - To create a snapshot of a volume on an Outpost and store the snapshot in
	//   the Region, omit this parameter. The snapshot is created in the Region for the
	//   Outpost.
	//   - To create a snapshot of a volume on an Outpost and store the snapshot on an
	//   Outpost, specify the ARN of the destination Outpost. The snapshot must be
	//   created on the same Outpost as the volume.
	// For more information, see Create local snapshots from volumes on an Outpost (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html#create-snapshot)
	// in the Amazon Elastic Compute Cloud User Guide.
	OutpostArn *string

	// The tags to apply to the snapshot during creation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

// Describes a snapshot.
type CreateSnapshotOutput struct {

	// The data encryption key identifier for the snapshot. This value is a unique
	// identifier that corresponds to the data encryption key that was used to encrypt
	// the original volume or snapshot copy. Because data encryption keys are inherited
	// by volumes created from snapshots, and vice versa, if snapshots share the same
	// data encryption key identifier, then they belong to the same volume/snapshot
	// lineage. This parameter is only returned by DescribeSnapshots .
	DataEncryptionKeyId *string

	// The description for the snapshot.
	Description *string

	// Indicates whether the snapshot is encrypted.
	Encrypted *bool

	// The Amazon Resource Name (ARN) of the Key Management Service (KMS) KMS key that
	// was used to protect the volume encryption key for the parent volume.
	KmsKeyId *string

	// The ARN of the Outpost on which the snapshot is stored. For more information,
	// see Amazon EBS local snapshots on Outposts (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	OutpostArn *string

	// The Amazon Web Services owner alias, from an Amazon-maintained list ( amazon ).
	// This is not the user-configured Amazon Web Services account alias set using the
	// IAM console.
	OwnerAlias *string

	// The ID of the Amazon Web Services account that owns the EBS snapshot.
	OwnerId *string

	// The progress of the snapshot, as a percentage.
	Progress *string

	// Only for archived snapshots that are temporarily restored. Indicates the date
	// and time when a temporarily restored snapshot will be automatically re-archived.
	RestoreExpiryTime *time.Time

	// The ID of the snapshot. Each snapshot receives a unique identifier when it is
	// created.
	SnapshotId *string

	// The time stamp when the snapshot was initiated.
	StartTime *time.Time

	// The snapshot state.
	State types.SnapshotState

	// Encrypted Amazon EBS snapshots are copied asynchronously. If a snapshot copy
	// operation fails (for example, if the proper Key Management Service (KMS)
	// permissions are not obtained) this field displays error state details to help
	// you diagnose why the error occurred. This parameter is only returned by
	// DescribeSnapshots .
	StateMessage *string

	// The storage tier in which the snapshot is stored. standard indicates that the
	// snapshot is stored in the standard snapshot storage tier and that it is ready
	// for use. archive indicates that the snapshot is currently archived and that it
	// must be restored before it can be used.
	StorageTier types.StorageTier

	// Any tags assigned to the snapshot.
	Tags []types.Tag

	// The ID of the volume that was used to create the snapshot. Snapshots created by
	// the CopySnapshot action have an arbitrary volume ID that should not be used for
	// any purpose.
	VolumeId *string

	// The size of the volume, in GiB.
	VolumeSize *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSnapshot{}, middleware.After)
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
	if err = addOpCreateSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSnapshot",
	}
}
