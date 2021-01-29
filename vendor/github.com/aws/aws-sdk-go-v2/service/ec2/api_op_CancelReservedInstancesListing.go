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

// Cancels the specified Reserved Instance listing in the Reserved Instance
// Marketplace. For more information, see Reserved Instance Marketplace
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CancelReservedInstancesListing(ctx context.Context, params *CancelReservedInstancesListingInput, optFns ...func(*Options)) (*CancelReservedInstancesListingOutput, error) {
	if params == nil {
		params = &CancelReservedInstancesListingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelReservedInstancesListing", params, optFns, addOperationCancelReservedInstancesListingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelReservedInstancesListingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelReservedInstancesListing.
type CancelReservedInstancesListingInput struct {

	// The ID of the Reserved Instance listing.
	//
	// This member is required.
	ReservedInstancesListingId *string
}

// Contains the output of CancelReservedInstancesListing.
type CancelReservedInstancesListingOutput struct {

	// The Reserved Instance listing.
	ReservedInstancesListings []types.ReservedInstancesListing

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelReservedInstancesListingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCancelReservedInstancesListing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCancelReservedInstancesListing{}, middleware.After)
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
	if err = addOpCancelReservedInstancesListingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelReservedInstancesListing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelReservedInstancesListing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelReservedInstancesListing",
	}
}
