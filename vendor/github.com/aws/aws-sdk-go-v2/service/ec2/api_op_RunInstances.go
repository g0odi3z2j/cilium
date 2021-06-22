// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Launches the specified number of instances using an AMI for which you have
// permissions. You can specify a number of options, or leave the default options.
// The following rules apply:
//
// * [EC2-VPC] If you don't specify a subnet ID, we
// choose a default subnet from your default VPC for you. If you don't have a
// default VPC, you must specify a subnet ID in the request.
//
// * [EC2-Classic] If
// don't specify an Availability Zone, we choose one for you.
//
// * Some instance
// types must be launched into a VPC. If you do not have a default VPC, or if you
// do not specify a subnet ID, the request fails. For more information, see
// Instance types available only in a VPC
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-vpc.html#vpc-only-instance-types).
//
// *
// [EC2-VPC] All instances have a network interface with a primary private IPv4
// address. If you don't specify this address, we choose one from the IPv4 range of
// your subnet.
//
// * Not all instance types support IPv6 addresses. For more
// information, see Instance types
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
//
// * If
// you don't specify a security group ID, we use the default security group. For
// more information, see Security groups
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).
//
// *
// If any of the AMIs have a product code attached for which the user has not
// subscribed, the request fails.
//
// You can create a launch template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html),
// which is a resource that contains the parameters to launch an instance. When you
// launch an instance using RunInstances, you can specify the launch template
// instead of specifying the launch parameters. To ensure faster instance launches,
// break up large requests into smaller batches. For example, create five separate
// launch requests for 100 instances each instead of one launch request for 500
// instances. An instance is ready for you to use when it's in the running state.
// You can check the state of your instance using DescribeInstances. You can tag
// instances and EBS volumes during launch, after launch, or both. For more
// information, see CreateTags and Tagging your Amazon EC2 resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html). Linux
// instances have access to the public key of the key pair at boot. You can use
// this key to provide secure access to the instance. Amazon EC2 public images use
// this feature to provide secure access without passwords. For more information,
// see Key pairs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html). For
// troubleshooting, see What to do if an instance immediately terminates
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_InstanceStraightToTerminated.html),
// and Troubleshooting connecting to your instance
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesConnecting.html).
func (c *Client) RunInstances(ctx context.Context, params *RunInstancesInput, optFns ...func(*Options)) (*RunInstancesOutput, error) {
	if params == nil {
		params = &RunInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RunInstances", params, optFns, addOperationRunInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RunInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RunInstancesInput struct {

	// The maximum number of instances to launch. If you specify more instances than
	// Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches the
	// largest possible number of instances above MinCount. Constraints: Between 1 and
	// the maximum number you're allowed for the specified instance type. For more
	// information about the default limits, and how to request an increase, see How
	// many instances can I run in Amazon EC2
	// (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in
	// the Amazon EC2 FAQ.
	//
	// This member is required.
	MaxCount *int32

	// The minimum number of instances to launch. If you specify a minimum that is more
	// instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2
	// launches no instances. Constraints: Between 1 and the maximum number you're
	// allowed for the specified instance type. For more information about the default
	// limits, and how to request an increase, see How many instances can I run in
	// Amazon EC2
	// (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in
	// the Amazon EC2 General FAQ.
	//
	// This member is required.
	MinCount *int32

	// Reserved.
	AdditionalInfo *string

	// The block device mapping entries.
	BlockDeviceMappings []types.BlockDeviceMapping

	// Information about the Capacity Reservation targeting option. If you do not
	// specify this parameter, the instance's Capacity Reservation preference defaults
	// to open, which enables it to run in any open Capacity Reservation that has
	// matching attributes (instance type, platform, Availability Zone).
	CapacityReservationSpecification *types.CapacityReservationSpecification

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If you do not specify a client token, a randomly generated token is
	// used for the request to ensure idempotency. For more information, see Ensuring
	// Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraints: Maximum 64 ASCII characters
	ClientToken *string

	// The CPU options for the instance. For more information, see Optimizing CPU
	// options
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html)
	// in the Amazon EC2 User Guide.
	CpuOptions *types.CpuOptionsRequest

	// The credit option for CPU usage of the burstable performance instance. Valid
	// values are standard and unlimited. To change this attribute after launch, use
	// ModifyInstanceCreditSpecification
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html).
	// For more information, see Burstable performance instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide. Default: standard (T2 instances) or unlimited
	// (T3/T3a instances)
	CreditSpecification *types.CreditSpecificationRequest

	// If you set this parameter to true, you can't terminate the instance using the
	// Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute
	// after launch, use ModifyInstanceAttribute
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html).
	// Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you
	// can terminate the instance by running the shutdown command from the instance.
	// Default: false
	DisableApiTermination *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Indicates whether the instance is optimized for Amazon EBS I/O. This
	// optimization provides dedicated throughput to Amazon EBS and an optimized
	// configuration stack to provide optimal Amazon EBS I/O performance. This
	// optimization isn't available with all instance types. Additional usage charges
	// apply when using an EBS-optimized instance. Default: false
	EbsOptimized *bool

	// An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource
	// that you can attach to your Windows instance to accelerate the graphics
	// performance of your applications. For more information, see Amazon EC2 Elastic
	// GPUs
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html)
	// in the Amazon EC2 User Guide.
	ElasticGpuSpecification []types.ElasticGpuSpecification

	// An elastic inference accelerator to associate with the instance. Elastic
	// inference accelerators are a resource you can attach to your Amazon EC2
	// instances to accelerate your Deep Learning (DL) inference workloads. You cannot
	// specify accelerators from different generations in the same request.
	ElasticInferenceAccelerators []types.ElasticInferenceAccelerator

	// Indicates whether the instance is enabled for AWS Nitro Enclaves. For more
	// information, see  What is AWS Nitro Enclaves?
	// (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the AWS
	// Nitro Enclaves User Guide. You can't enable AWS Nitro Enclaves and hibernation
	// on the same instance.
	EnclaveOptions *types.EnclaveOptionsRequest

	// Indicates whether an instance is enabled for hibernation. For more information,
	// see Hibernate your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the
	// Amazon EC2 User Guide. You can't enable hibernation and AWS Nitro Enclaves on
	// the same instance.
	HibernationOptions *types.HibernationOptionsRequest

	// The name or Amazon Resource Name (ARN) of an IAM instance profile.
	IamInstanceProfile *types.IamInstanceProfileSpecification

	// The ID of the AMI. An AMI ID is required to launch an instance and must be
	// specified here or in a launch template.
	ImageId *string

	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	// Default: stop
	InstanceInitiatedShutdownBehavior types.ShutdownBehavior

	// The market (purchasing) option for the instances. For RunInstances, persistent
	// Spot Instance requests are only supported when InstanceInterruptionBehavior is
	// set to either hibernate or stop.
	InstanceMarketOptions *types.InstanceMarketOptionsRequest

	// The instance type. For more information, see Instance types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon EC2 User Guide. Default: m1.small
	InstanceType types.InstanceType

	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network
	// interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	// You cannot specify this option and the option to assign specific IPv6 addresses
	// in the same request. You can specify this option if you've specified a minimum
	// number of instances to launch. You cannot specify this option and the network
	// interfaces option in the same request.
	Ipv6AddressCount *int32

	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the
	// primary network interface. You cannot specify this option and the option to
	// assign a number of IPv6 addresses in the same request. You cannot specify this
	// option if you've specified a minimum number of instances to launch. You cannot
	// specify this option and the network interfaces option in the same request.
	Ipv6Addresses []types.InstanceIpv6Address

	// The ID of the kernel. We recommend that you use PV-GRUB instead of kernels and
	// RAM disks. For more information, see  PV-GRUB
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon EC2 User Guide.
	KernelId *string

	// The name of the key pair. You can create a key pair using CreateKeyPair
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html)
	// or ImportKeyPair
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).
	// If you do not specify a key pair, you can't connect to the instance unless you
	// choose an AMI that is configured to allow users another way to log in.
	KeyName *string

	// The launch template to use to launch the instances. Any parameters that you
	// specify in RunInstances override the same parameters in the launch template. You
	// can specify either the name or ID of a launch template, but not both.
	LaunchTemplate *types.LaunchTemplateSpecification

	// The license configurations.
	LicenseSpecifications []types.LicenseConfigurationRequest

	// The metadata options for the instance. For more information, see Instance
	// metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
	MetadataOptions *types.InstanceMetadataOptionsRequest

	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring *types.RunInstancesMonitoringEnabled

	// The network interfaces to associate with the instance. If you specify a network
	// interface, you must specify any security groups and subnets as part of the
	// network interface.
	NetworkInterfaces []types.InstanceNetworkInterfaceSpecification

	// The placement for the instance.
	Placement *types.Placement

	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4
	// address range of the subnet. Only one private IP address can be designated as
	// primary. You can't specify this option if you've specified the option to
	// designate a private IP address as the primary IP address in a network interface
	// specification. You cannot specify this option if you're launching more than one
	// instance in the request. You cannot specify this option and the network
	// interfaces option in the same request.
	PrivateIpAddress *string

	// The ID of the RAM disk to select. Some kernels require additional drivers at
	// launch. Check the kernel requirements for information about whether you need to
	// specify a RAM disk. To find kernel requirements, go to the AWS Resource Center
	// and search for the kernel ID. We recommend that you use PV-GRUB instead of
	// kernels and RAM disks. For more information, see  PV-GRUB
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon EC2 User Guide.
	RamdiskId *string

	// The IDs of the security groups. You can create a security group using
	// CreateSecurityGroup
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html).
	// If you specify a network interface, you must specify any security groups as part
	// of the network interface.
	SecurityGroupIds []string

	// [EC2-Classic, default VPC] The names of the security groups. For a nondefault
	// VPC, you must use security group IDs instead. If you specify a network
	// interface, you must specify any security groups as part of the network
	// interface. Default: Amazon EC2 uses the default security group.
	SecurityGroups []string

	// [EC2-VPC] The ID of the subnet to launch the instance into. If you specify a
	// network interface, you must specify any subnets as part of the network
	// interface.
	SubnetId *string

	// The tags to apply to the resources during launch. You can only tag instances and
	// volumes on launch. The specified tags are applied to all instances or volumes
	// that are created during launch. To tag a resource after it has been created, see
	// CreateTags
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	TagSpecifications []types.TagSpecification

	// The user data to make available to the instance. For more information, see
	// Running commands on your Linux instance at launch
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) (Linux) and
	// Adding User Data
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html#instancedata-add-user-data)
	// (Windows). If you are using a command line tool, base64-encoding is performed
	// for you, and you can load the text from a file. Otherwise, you must provide
	// base64-encoded text. User data is limited to 16 KB.
	UserData *string
}

// Describes a launch request for one or more instances, and includes owner,
// requester, and security group information that applies to all instances in the
// launch request.
type RunInstancesOutput struct {

	// [EC2-Classic only] The security groups.
	Groups []types.GroupIdentifier

	// The instances.
	Instances []types.Instance

	// The ID of the AWS account that owns the reservation.
	OwnerId *string

	// The ID of the requester that launched the instances on your behalf (for example,
	// AWS Management Console or Auto Scaling).
	RequesterId *string

	// The ID of the reservation.
	ReservationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRunInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRunInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRunInstances{}, middleware.After)
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
	if err = addIdempotencyToken_opRunInstancesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRunInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRunInstances(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRunInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRunInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRunInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RunInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RunInstancesInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRunInstancesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRunInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRunInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RunInstances",
	}
}
