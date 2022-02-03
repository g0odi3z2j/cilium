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

// Discontinue faster launching for a Windows AMI, and clean up existing
// pre-provisioned snapshots. When you disable faster launching, the AMI uses the
// standard launch process for each instance. All pre-provisioned snapshots must be
// removed before you can enable faster launching again.
func (c *Client) DisableFastLaunch(ctx context.Context, params *DisableFastLaunchInput, optFns ...func(*Options)) (*DisableFastLaunchOutput, error) {
	if params == nil {
		params = &DisableFastLaunchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableFastLaunch", params, optFns, c.addOperationDisableFastLaunchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableFastLaunchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableFastLaunchInput struct {

	// The ID of the image for which you’re turning off faster launching, and removing
	// pre-provisioned snapshots.
	//
	// This member is required.
	ImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Forces the image settings to turn off faster launching for your Windows AMI.
	// This parameter overrides any errors that are encountered while cleaning up
	// resources in your account.
	Force *bool

	noSmithyDocumentSerde
}

type DisableFastLaunchOutput struct {

	// The ID of the image for which faster-launching has been turned off.
	ImageId *string

	// The launch template that was used to launch Windows instances from
	// pre-provisioned snapshots.
	LaunchTemplate *types.FastLaunchLaunchTemplateSpecificationResponse

	// The maximum number of parallel instances to launch for creating resources.
	MaxParallelLaunches *int32

	// The owner of the Windows AMI for which faster launching was turned off.
	OwnerId *string

	// The pre-provisioning resource type that must be cleaned after turning off faster
	// launching for the Windows AMI. Supported values include: snapshot.
	ResourceType types.FastLaunchResourceType

	// Parameters that were used for faster launching for the Windows AMI before faster
	// launching was turned off. This informs the clean-up process.
	SnapshotConfiguration *types.FastLaunchSnapshotConfigurationResponse

	// The current state of faster launching for the specified Windows AMI.
	State types.FastLaunchStateCode

	// The reason that the state changed for faster launching for the Windows AMI.
	StateTransitionReason *string

	// The time that the state changed for faster launching for the Windows AMI.
	StateTransitionTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableFastLaunchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisableFastLaunch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisableFastLaunch{}, middleware.After)
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
	if err = addOpDisableFastLaunchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableFastLaunch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableFastLaunch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableFastLaunch",
	}
}
