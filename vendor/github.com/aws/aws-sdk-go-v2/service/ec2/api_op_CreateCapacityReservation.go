// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new Capacity Reservation with the specified attributes. Capacity
// Reservations enable you to reserve capacity for your Amazon EC2 instances in a
// specific Availability Zone for any duration. This gives you the flexibility to
// selectively add capacity reservations and still get the Regional RI discounts
// for that usage. By creating Capacity Reservations, you ensure that you always
// have access to Amazon EC2 capacity when you need it, for as long as you need it.
// For more information, see Capacity Reservations
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html)
// in the Amazon Elastic Compute Cloud User Guide. Your request to create a
// Capacity Reservation could fail if Amazon EC2 does not have sufficient capacity
// to fulfill the request. If your request fails due to Amazon EC2 capacity
// constraints, either try again at a later time, try in a different Availability
// Zone, or request a smaller capacity reservation. If your application is flexible
// across instance types and sizes, try to create a Capacity Reservation with
// different instance attributes. Your request could also fail if the requested
// quantity exceeds your On-Demand Instance limit for the selected instance type.
// If your request fails due to limit constraints, increase your On-Demand Instance
// limit for the required instance type and try again. For more information about
// increasing your instance limits, see Amazon EC2 Service Limits
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateCapacityReservation(ctx context.Context, params *CreateCapacityReservationInput, optFns ...func(*Options)) (*CreateCapacityReservationOutput, error) {
	if params == nil {
		params = &CreateCapacityReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCapacityReservation", params, optFns, addOperationCreateCapacityReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCapacityReservationInput struct {

	// The number of instances for which to reserve capacity.
	//
	// This member is required.
	InstanceCount int32

	// The type of operating system for which to reserve capacity.
	//
	// This member is required.
	InstancePlatform types.CapacityReservationInstancePlatform

	// The instance type for which to reserve capacity. For more information, see
	// Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	// This member is required.
	InstanceType *string

	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone *string

	// The ID of the Availability Zone in which to create the Capacity Reservation.
	AvailabilityZoneId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	// This optimization provides dedicated throughput to Amazon EBS and an optimized
	// configuration stack to provide optimal I/O performance. This optimization isn't
	// available with all instance types. Additional usage charges apply when using an
	// EBS- optimized instance.
	EbsOptimized bool

	// The date and time at which the Capacity Reservation expires. When a Capacity
	// Reservation expires, the reserved capacity is released and you can no longer
	// launch instances into it. The Capacity Reservation's state changes to expired
	// when it reaches its end date and time. You must provide an EndDate value if
	// EndDateType is limited. Omit EndDate if EndDateType is unlimited. If the
	// EndDateType is limited, the Capacity Reservation is cancelled within an hour
	// from the specified time. For example, if you specify 5/31/2019, 13:30:55, the
	// Capacity Reservation is guaranteed to end between 13:30:55 and 14:30:55 on
	// 5/31/2019.
	EndDate *time.Time

	// Indicates the way in which the Capacity Reservation ends. A Capacity Reservation
	// can have one of the following end types:
	//
	// * unlimited - The Capacity Reservation
	// remains active until you explicitly cancel it. Do not provide an EndDate if the
	// EndDateType is unlimited.
	//
	// * limited - The Capacity Reservation expires
	// automatically at a specified date and time. You must provide an EndDate value if
	// the EndDateType value is limited.
	EndDateType types.EndDateType

	// Indicates whether the Capacity Reservation supports instances with temporary,
	// block-level storage.
	EphemeralStorage bool

	// Indicates the type of instance launches that the Capacity Reservation accepts.
	// The options include:
	//
	// * open - The Capacity Reservation automatically matches
	// all instances that have matching attributes (instance type, platform, and
	// Availability Zone). Instances that have matching attributes run in the Capacity
	// Reservation automatically without specifying any additional parameters.
	//
	// *
	// targeted - The Capacity Reservation only accepts instances that have matching
	// attributes (instance type, platform, and Availability Zone), and explicitly
	// target the Capacity Reservation. This ensures that only permitted instances can
	// use the reserved capacity.
	//
	// Default: open
	InstanceMatchCriteria types.InstanceMatchCriteria

	// The tags to apply to the Capacity Reservation during launch.
	TagSpecifications []types.TagSpecification

	// Indicates the tenancy of the Capacity Reservation. A Capacity Reservation can
	// have one of the following tenancy settings:
	//
	// * default - The Capacity
	// Reservation is created on hardware that is shared with other AWS accounts.
	//
	// *
	// dedicated - The Capacity Reservation is created on single-tenant hardware that
	// is dedicated to a single AWS account.
	Tenancy types.CapacityReservationTenancy
}

type CreateCapacityReservationOutput struct {

	// Information about the Capacity Reservation.
	CapacityReservation *types.CapacityReservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCapacityReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateCapacityReservation{}, middleware.After)
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
	if err = addOpCreateCapacityReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCapacityReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCapacityReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateCapacityReservation",
	}
}
