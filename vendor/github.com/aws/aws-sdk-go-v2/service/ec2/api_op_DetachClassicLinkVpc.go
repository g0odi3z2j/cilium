// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unlinks (detaches) a linked EC2-Classic instance from a VPC. After the instance
// has been unlinked, the VPC security groups are no longer associated with it. An
// instance is automatically unlinked from a VPC when it's stopped.
func (c *Client) DetachClassicLinkVpc(ctx context.Context, params *DetachClassicLinkVpcInput, optFns ...func(*Options)) (*DetachClassicLinkVpcOutput, error) {
	if params == nil {
		params = &DetachClassicLinkVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachClassicLinkVpc", params, optFns, addOperationDetachClassicLinkVpcMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachClassicLinkVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachClassicLinkVpcInput struct {

	// The ID of the instance to unlink from the VPC.
	//
	// This member is required.
	InstanceId *string

	// The ID of the VPC to which the instance is linked.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type DetachClassicLinkVpcOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachClassicLinkVpcMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDetachClassicLinkVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDetachClassicLinkVpc{}, middleware.After)
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
	if err = addOpDetachClassicLinkVpcValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachClassicLinkVpc(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachClassicLinkVpc(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DetachClassicLinkVpc",
	}
}
