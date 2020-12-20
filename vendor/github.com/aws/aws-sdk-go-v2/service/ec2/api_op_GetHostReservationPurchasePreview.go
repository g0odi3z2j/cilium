// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Preview a reservation purchase with configurations that match those of your
// Dedicated Host. You must have active Dedicated Hosts in your account before you
// purchase a reservation. This is a preview of the PurchaseHostReservation action
// and does not result in the offering being purchased.
func (c *Client) GetHostReservationPurchasePreview(ctx context.Context, params *GetHostReservationPurchasePreviewInput, optFns ...func(*Options)) (*GetHostReservationPurchasePreviewOutput, error) {
	if params == nil {
		params = &GetHostReservationPurchasePreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHostReservationPurchasePreview", params, optFns, addOperationGetHostReservationPurchasePreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHostReservationPurchasePreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHostReservationPurchasePreviewInput struct {

	// The IDs of the Dedicated Hosts with which the reservation is associated.
	//
	// This member is required.
	HostIdSet []string

	// The offering ID of the reservation.
	//
	// This member is required.
	OfferingId *string
}

type GetHostReservationPurchasePreviewOutput struct {

	// The currency in which the totalUpfrontPrice and totalHourlyPrice amounts are
	// specified. At this time, the only supported currency is USD.
	CurrencyCode types.CurrencyCodeValues

	// The purchase information of the Dedicated Host reservation and the Dedicated
	// Hosts associated with it.
	Purchase []types.Purchase

	// The potential total hourly price of the reservation per hour.
	TotalHourlyPrice *string

	// The potential total upfront price. This is billed immediately.
	TotalUpfrontPrice *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHostReservationPurchasePreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetHostReservationPurchasePreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetHostReservationPurchasePreview{}, middleware.After)
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
	if err = addOpGetHostReservationPurchasePreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostReservationPurchasePreview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHostReservationPurchasePreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetHostReservationPurchasePreview",
	}
}
