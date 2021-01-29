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

// Creates a new version for a launch template. You can specify an existing version
// of launch template from which to base the new version. Launch template versions
// are numbered in the order in which they are created. You cannot specify, change,
// or replace the numbering of launch template versions. For more information, see
// Managing launch template versions
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#manage-launch-template-versions)in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateLaunchTemplateVersion(ctx context.Context, params *CreateLaunchTemplateVersionInput, optFns ...func(*Options)) (*CreateLaunchTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateLaunchTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchTemplateVersion", params, optFns, addOperationCreateLaunchTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchTemplateVersionInput struct {

	// The information for the launch template.
	//
	// This member is required.
	LaunchTemplateData *types.RequestLaunchTemplateData

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The ID of the launch template. You must specify either the launch template ID or
	// launch template name in the request.
	LaunchTemplateId *string

	// The name of the launch template. You must specify either the launch template ID
	// or launch template name in the request.
	LaunchTemplateName *string

	// The version number of the launch template version on which to base the new
	// version. The new version inherits the same launch parameters as the source
	// version, except for parameters that you specify in LaunchTemplateData. Snapshots
	// applied to the block device mapping are ignored when creating a new version
	// unless they are explicitly included.
	SourceVersion *string

	// A description for the version of the launch template.
	VersionDescription *string
}

type CreateLaunchTemplateVersionOutput struct {

	// Information about the launch template version.
	LaunchTemplateVersion *types.LaunchTemplateVersion

	// If the new version of the launch template contains parameters or parameter
	// combinations that are not valid, an error code and an error message are returned
	// for each issue that's found.
	Warning *types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLaunchTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateLaunchTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLaunchTemplateVersion{}, middleware.After)
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
	if err = addOpCreateLaunchTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchTemplateVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLaunchTemplateVersion",
	}
}
