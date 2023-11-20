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

// Purchase a reservation with configurations that match those of your Dedicated
// Host. You must have active Dedicated Hosts in your account before you purchase a
// reservation. This action results in the specified reservation being purchased
// and charged to your account.
func (c *Client) PurchaseHostReservation(ctx context.Context, params *PurchaseHostReservationInput, optFns ...func(*Options)) (*PurchaseHostReservationOutput, error) {
	if params == nil {
		params = &PurchaseHostReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseHostReservation", params, optFns, c.addOperationPurchaseHostReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseHostReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PurchaseHostReservationInput struct {

	// The IDs of the Dedicated Hosts with which the reservation will be associated.
	//
	// This member is required.
	HostIdSet []string

	// The ID of the offering.
	//
	// This member is required.
	OfferingId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The currency in which the totalUpfrontPrice , LimitPrice , and totalHourlyPrice
	// amounts are specified. At this time, the only supported currency is USD .
	CurrencyCode types.CurrencyCodeValues

	// The specified limit is checked against the total upfront cost of the
	// reservation (calculated as the offering's upfront cost multiplied by the host
	// count). If the total upfront cost is greater than the specified price limit, the
	// request fails. This is used to ensure that the purchase does not exceed the
	// expected upfront cost of the purchase. At this time, the only supported currency
	// is USD . For example, to indicate a limit price of USD 100, specify 100.00.
	LimitPrice *string

	// The tags to apply to the Dedicated Host Reservation during purchase.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type PurchaseHostReservationOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The currency in which the totalUpfrontPrice and totalHourlyPrice amounts are
	// specified. At this time, the only supported currency is USD .
	CurrencyCode types.CurrencyCodeValues

	// Describes the details of the purchase.
	Purchase []types.Purchase

	// The total hourly price of the reservation calculated per hour.
	TotalHourlyPrice *string

	// The total amount charged to your account when you purchase the reservation.
	TotalUpfrontPrice *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseHostReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpPurchaseHostReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpPurchaseHostReservation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PurchaseHostReservation"); err != nil {
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
	if err = addOpPurchaseHostReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseHostReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseHostReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PurchaseHostReservation",
	}
}
