// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// We are retiring EC2-Classic on August 15, 2022. We recommend that you migrate
// from EC2-Classic to a VPC. For more information, see Migrate from EC2-Classic to
// a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in
// the Amazon Elastic Compute Cloud User Guide. Enables a VPC for ClassicLink. You
// can then link EC2-Classic instances to your ClassicLink-enabled VPC to allow
// communication over private IP addresses. You cannot enable your VPC for
// ClassicLink if any of your VPC route tables have existing routes for address
// ranges within the 10.0.0.0/8 IP address range, excluding local routes for VPCs
// in the 10.0.0.0/16 and 10.1.0.0/16 IP address ranges. For more information, see
// ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) EnableVpcClassicLink(ctx context.Context, params *EnableVpcClassicLinkInput, optFns ...func(*Options)) (*EnableVpcClassicLinkOutput, error) {
	if params == nil {
		params = &EnableVpcClassicLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableVpcClassicLink", params, optFns, c.addOperationEnableVpcClassicLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableVpcClassicLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableVpcClassicLinkInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type EnableVpcClassicLinkOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableVpcClassicLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableVpcClassicLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableVpcClassicLink{}, middleware.After)
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
	if err = addOpEnableVpcClassicLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableVpcClassicLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableVpcClassicLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableVpcClassicLink",
	}
}
