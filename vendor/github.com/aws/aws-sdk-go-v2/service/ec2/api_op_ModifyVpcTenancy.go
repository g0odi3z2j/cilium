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

// Modifies the instance tenancy attribute of the specified VPC. You can change the
// instance tenancy attribute of a VPC to default only. You cannot change the
// instance tenancy attribute to dedicated. After you modify the tenancy of the
// VPC, any new instances that you launch into the VPC have a tenancy of default,
// unless you specify otherwise during launch. The tenancy of any existing
// instances in the VPC is not affected. For more information, see Dedicated
// Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifyVpcTenancy(ctx context.Context, params *ModifyVpcTenancyInput, optFns ...func(*Options)) (*ModifyVpcTenancyOutput, error) {
	if params == nil {
		params = &ModifyVpcTenancyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcTenancy", params, optFns, addOperationModifyVpcTenancyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcTenancyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcTenancyInput struct {

	// The instance tenancy attribute for the VPC.
	//
	// This member is required.
	InstanceTenancy types.VpcTenancy

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyVpcTenancyOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcTenancyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcTenancy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcTenancy{}, middleware.After)
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
	if err = addOpModifyVpcTenancyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcTenancy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcTenancy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcTenancy",
	}
}
