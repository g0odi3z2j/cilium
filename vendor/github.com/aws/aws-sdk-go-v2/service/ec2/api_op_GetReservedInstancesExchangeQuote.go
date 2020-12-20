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

// Returns a quote and exchange information for exchanging one or more specified
// Convertible Reserved Instances for a new Convertible Reserved Instance. If the
// exchange cannot be performed, the reason is returned in the response. Use
// AcceptReservedInstancesExchangeQuote to perform the exchange.
func (c *Client) GetReservedInstancesExchangeQuote(ctx context.Context, params *GetReservedInstancesExchangeQuoteInput, optFns ...func(*Options)) (*GetReservedInstancesExchangeQuoteOutput, error) {
	if params == nil {
		params = &GetReservedInstancesExchangeQuoteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservedInstancesExchangeQuote", params, optFns, addOperationGetReservedInstancesExchangeQuoteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservedInstancesExchangeQuoteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for GetReservedInstanceExchangeQuote.
type GetReservedInstancesExchangeQuoteInput struct {

	// The IDs of the Convertible Reserved Instances to exchange.
	//
	// This member is required.
	ReservedInstanceIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The configuration of the target Convertible Reserved Instance to exchange for
	// your current Convertible Reserved Instances.
	TargetConfigurations []types.TargetConfigurationRequest
}

// Contains the output of GetReservedInstancesExchangeQuote.
type GetReservedInstancesExchangeQuoteOutput struct {

	// The currency of the transaction.
	CurrencyCode *string

	// If true, the exchange is valid. If false, the exchange cannot be completed.
	IsValidExchange bool

	// The new end date of the reservation term.
	OutputReservedInstancesWillExpireAt *time.Time

	// The total true upfront charge for the exchange.
	PaymentDue *string

	// The cost associated with the Reserved Instance.
	ReservedInstanceValueRollup *types.ReservationValue

	// The configuration of your Convertible Reserved Instances.
	ReservedInstanceValueSet []types.ReservedInstanceReservationValue

	// The cost associated with the Reserved Instance.
	TargetConfigurationValueRollup *types.ReservationValue

	// The values of the target Convertible Reserved Instances.
	TargetConfigurationValueSet []types.TargetReservationValue

	// Describes the reason why the exchange cannot be completed.
	ValidationFailureReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetReservedInstancesExchangeQuoteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetReservedInstancesExchangeQuote{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetReservedInstancesExchangeQuote{}, middleware.After)
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
	if err = addOpGetReservedInstancesExchangeQuoteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservedInstancesExchangeQuote(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReservedInstancesExchangeQuote(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetReservedInstancesExchangeQuote",
	}
}
