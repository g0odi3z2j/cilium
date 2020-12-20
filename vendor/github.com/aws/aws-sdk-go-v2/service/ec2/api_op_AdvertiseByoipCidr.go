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

// Advertises an IPv4 or IPv6 address range that is provisioned for use with your
// AWS resources through bring your own IP addresses (BYOIP). You can perform this
// operation at most once every 10 seconds, even if you specify different address
// ranges each time. We recommend that you stop advertising the BYOIP CIDR from
// other locations when you advertise it from AWS. To minimize down time, you can
// configure your AWS resources to use an address from a BYOIP CIDR before it is
// advertised, and then simultaneously stop advertising it from the current
// location and start advertising it through AWS. It can take a few minutes before
// traffic to the specified addresses starts routing to AWS because of BGP
// propagation delays. To stop advertising the BYOIP CIDR, use WithdrawByoipCidr.
func (c *Client) AdvertiseByoipCidr(ctx context.Context, params *AdvertiseByoipCidrInput, optFns ...func(*Options)) (*AdvertiseByoipCidrOutput, error) {
	if params == nil {
		params = &AdvertiseByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdvertiseByoipCidr", params, optFns, addOperationAdvertiseByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdvertiseByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdvertiseByoipCidrInput struct {

	// The address range, in CIDR notation. This must be the exact range that you
	// provisioned. You can't advertise only a portion of the provisioned range.
	//
	// This member is required.
	Cidr *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type AdvertiseByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdvertiseByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAdvertiseByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAdvertiseByoipCidr{}, middleware.After)
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
	if err = addOpAdvertiseByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdvertiseByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdvertiseByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AdvertiseByoipCidr",
	}
}
