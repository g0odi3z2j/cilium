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

// Describes the specified attribute of the specified instance. You can specify
// only one attribute at a time. Valid attribute values are: instanceType | kernel
// | ramdisk | userData | disableApiTermination | instanceInitiatedShutdownBehavior
// | rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
// groupSet | ebsOptimized | sriovNetSupport
func (c *Client) DescribeInstanceAttribute(ctx context.Context, params *DescribeInstanceAttributeInput, optFns ...func(*Options)) (*DescribeInstanceAttributeOutput, error) {
	if params == nil {
		params = &DescribeInstanceAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceAttribute", params, optFns, addOperationDescribeInstanceAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceAttributeInput struct {

	// The instance attribute. Note: The enaSupport attribute is not supported at this
	// time.
	//
	// This member is required.
	Attribute types.InstanceAttributeName

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

// Describes an instance attribute.
type DescribeInstanceAttributeOutput struct {

	// The block device mapping of the instance.
	BlockDeviceMappings []types.InstanceBlockDeviceMapping

	// If the value is true, you can't terminate the instance through the Amazon EC2
	// console, CLI, or API; otherwise, you can.
	DisableApiTermination *types.AttributeBooleanValue

	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized *types.AttributeBooleanValue

	// Indicates whether enhanced networking with ENA is enabled.
	EnaSupport *types.AttributeBooleanValue

	// To enable the instance for AWS Nitro Enclaves, set this parameter to true;
	// otherwise, set it to false.
	EnclaveOptions *types.EnclaveOptions

	// The security groups associated with the instance.
	Groups []types.GroupIdentifier

	// The ID of the instance.
	InstanceId *string

	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior *types.AttributeValue

	// The instance type.
	InstanceType *types.AttributeValue

	// The kernel ID.
	KernelId *types.AttributeValue

	// A list of product codes.
	ProductCodes []types.ProductCode

	// The RAM disk ID.
	RamdiskId *types.AttributeValue

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *types.AttributeValue

	// Indicates whether source/destination checking is enabled. A value of true means
	// that checking is enabled, and false means that checking is disabled. This value
	// must be false for a NAT instance to perform NAT.
	SourceDestCheck *types.AttributeBooleanValue

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *types.AttributeValue

	// The user data.
	UserData *types.AttributeValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstanceAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceAttribute{}, middleware.After)
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
	if err = addOpDescribeInstanceAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceAttribute",
	}
}
