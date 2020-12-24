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

// Registers an AMI. When you're creating an AMI, this is the final step you must
// complete before you can launch an instance from the AMI. For more information
// about creating AMIs, see Creating your own AMIs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html) in
// the Amazon Elastic Compute Cloud User Guide. For Amazon EBS-backed instances,
// CreateImage creates and registers the AMI in a single request, so you don't have
// to register the AMI yourself. You can also use RegisterImage to create an Amazon
// EBS-backed Linux AMI from a snapshot of a root device volume. You specify the
// snapshot using the block device mapping. For more information, see Launching a
// Linux instance from a backup
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-launch-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide. If any snapshots have AWS
// Marketplace product codes, they are copied to the new AMI. Windows and some
// Linux distributions, such as Red Hat Enterprise Linux (RHEL) and SUSE Linux
// Enterprise Server (SLES), use the EC2 billing product code associated with an
// AMI to verify the subscription status for package updates. To create a new AMI
// for operating systems that require a billing product code, instead of
// registering the AMI, do the following to preserve the billing product code
// association:
//
// * Launch an instance from an existing AMI with that billing
// product code.
//
// * Customize the instance.
//
// * Create an AMI from the instance
// using CreateImage.
//
// If you purchase a Reserved Instance to apply to an On-Demand
// Instance that was launched from an AMI with a billing product code, make sure
// that the Reserved Instance has the matching billing product code. If you
// purchase a Reserved Instance without the matching billing product code, the
// Reserved Instance will not be applied to the On-Demand Instance. For information
// about how to obtain the platform details and billing information of an AMI, see
// Obtaining billing information
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html) in
// the Amazon Elastic Compute Cloud User Guide.
//
// If needed, you can deregister an
// AMI at any time. Any modifications you make to an AMI backed by an instance
// store volume invalidates its registration. If you make changes to an image,
// deregister the previous image and register the new image.
func (c *Client) RegisterImage(ctx context.Context, params *RegisterImageInput, optFns ...func(*Options)) (*RegisterImageOutput, error) {
	if params == nil {
		params = &RegisterImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterImage", params, optFns, addOperationRegisterImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterImage.
type RegisterImageInput struct {

	// A name for your AMI. Constraints: 3-128 alphanumeric characters, parentheses
	// (()), square brackets ([]), spaces ( ), periods (.), slashes (/), dashes (-),
	// single quotes ('), at-signs (@), or underscores(_)
	//
	// This member is required.
	Name *string

	// The architecture of the AMI. Default: For Amazon EBS-backed AMIs, i386. For
	// instance store-backed AMIs, the architecture specified in the manifest file.
	Architecture types.ArchitectureValues

	// The billing product codes. Your account must be authorized to specify billing
	// product codes. Otherwise, you can use the AWS Marketplace to bill for the use of
	// an AMI.
	BillingProducts []string

	// The block device mapping entries.
	BlockDeviceMappings []types.BlockDeviceMapping

	BootMode types.BootModeValues

	// A description for your AMI.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Set to true to enable enhanced networking with ENA for the AMI and any instances
	// that you launch from the AMI. This option is supported only for HVM AMIs.
	// Specifying this option with a PV AMI can make instances launched from the AMI
	// unreachable.
	EnaSupport bool

	// The full path to your AMI manifest in Amazon S3 storage. The specified bucket
	// must have the aws-exec-read canned access control list (ACL) to ensure that it
	// can be accessed by Amazon EC2. For more information, see Canned ACLs
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
	// in the Amazon S3 Service Developer Guide.
	ImageLocation *string

	// The ID of the kernel.
	KernelId *string

	// The ID of the RAM disk.
	RamdiskId *string

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *string

	// Set to simple to enable enhanced networking with the Intel 82599 Virtual
	// Function interface for the AMI and any instances that you launch from the AMI.
	// There is no way to disable sriovNetSupport at this time. This option is
	// supported only for HVM AMIs. Specifying this option with a PV AMI can make
	// instances launched from the AMI unreachable.
	SriovNetSupport *string

	// The type of virtualization (hvm | paravirtual). Default: paravirtual
	VirtualizationType *string
}

// Contains the output of RegisterImage.
type RegisterImageOutput struct {

	// The ID of the newly registered AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRegisterImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterImage{}, middleware.After)
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
	if err = addOpRegisterImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterImage",
	}
}
