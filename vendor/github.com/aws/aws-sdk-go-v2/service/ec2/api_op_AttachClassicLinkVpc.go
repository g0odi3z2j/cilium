// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Links an EC2-Classic instance to a ClassicLink-enabled VPC through one or more
// of the VPC's security groups. You cannot link an EC2-Classic instance to more
// than one VPC at a time. You can only link an instance that's in the running
// state. An instance is automatically unlinked from a VPC when it's stopped - you
// can link it to the VPC again when you restart it. After you've linked an
// instance, you cannot change the VPC security groups that are associated with it.
// To change the security groups, you must first unlink the instance, and then link
// it again. Linking your instance to a VPC is sometimes referred to as attaching
// your instance.
func (c *Client) AttachClassicLinkVpc(ctx context.Context, params *AttachClassicLinkVpcInput, optFns ...func(*Options)) (*AttachClassicLinkVpcOutput, error) {
	if params == nil {
		params = &AttachClassicLinkVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachClassicLinkVpc", params, optFns, addOperationAttachClassicLinkVpcMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachClassicLinkVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachClassicLinkVpcInput struct {

	// The ID of one or more of the VPC's security groups. You cannot specify security
	// groups from a different VPC.
	//
	// This member is required.
	Groups []string

	// The ID of an EC2-Classic instance to link to the ClassicLink-enabled VPC.
	//
	// This member is required.
	InstanceId *string

	// The ID of a ClassicLink-enabled VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AttachClassicLinkVpcOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachClassicLinkVpcMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAttachClassicLinkVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAttachClassicLinkVpc{}, middleware.After)
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
	if err = addOpAttachClassicLinkVpcValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachClassicLinkVpc(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachClassicLinkVpc(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AttachClassicLinkVpc",
	}
}
