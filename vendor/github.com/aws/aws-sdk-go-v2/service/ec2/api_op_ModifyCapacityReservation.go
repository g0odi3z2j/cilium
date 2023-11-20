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
	"time"
)

// Modifies a Capacity Reservation's capacity and the conditions under which it is
// to be released. You cannot change a Capacity Reservation's instance type, EBS
// optimization, instance store settings, platform, Availability Zone, or instance
// eligibility. If you need to modify any of these attributes, we recommend that
// you cancel the Capacity Reservation, and then create a new one with the required
// attributes.
func (c *Client) ModifyCapacityReservation(ctx context.Context, params *ModifyCapacityReservationInput, optFns ...func(*Options)) (*ModifyCapacityReservationOutput, error) {
	if params == nil {
		params = &ModifyCapacityReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCapacityReservation", params, optFns, c.addOperationModifyCapacityReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCapacityReservationInput struct {

	// The ID of the Capacity Reservation.
	//
	// This member is required.
	CapacityReservationId *string

	// Reserved. Capacity Reservations you have created are accepted by default.
	Accept *bool

	// Reserved for future use.
	AdditionalInfo *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The date and time at which the Capacity Reservation expires. When a Capacity
	// Reservation expires, the reserved capacity is released and you can no longer
	// launch instances into it. The Capacity Reservation's state changes to expired
	// when it reaches its end date and time. The Capacity Reservation is cancelled
	// within an hour from the specified time. For example, if you specify 5/31/2019,
	// 13:30:55, the Capacity Reservation is guaranteed to end between 13:30:55 and
	// 14:30:55 on 5/31/2019. You must provide an EndDate value if EndDateType is
	// limited . Omit EndDate if EndDateType is unlimited .
	EndDate *time.Time

	// Indicates the way in which the Capacity Reservation ends. A Capacity
	// Reservation can have one of the following end types:
	//   - unlimited - The Capacity Reservation remains active until you explicitly
	//   cancel it. Do not provide an EndDate value if EndDateType is unlimited .
	//   - limited - The Capacity Reservation expires automatically at a specified date
	//   and time. You must provide an EndDate value if EndDateType is limited .
	EndDateType types.EndDateType

	// The number of instances for which to reserve capacity. The number of instances
	// can't be increased or decreased by more than 1000 in a single request.
	InstanceCount *int32

	noSmithyDocumentSerde
}

type ModifyCapacityReservationOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCapacityReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyCapacityReservation"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyCapacityReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCapacityReservation(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyCapacityReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyCapacityReservation",
	}
}
