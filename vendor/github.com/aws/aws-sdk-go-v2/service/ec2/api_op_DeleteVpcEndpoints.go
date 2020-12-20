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

// Deletes one or more specified VPC endpoints. Deleting a gateway endpoint also
// deletes the endpoint routes in the route tables that were associated with the
// endpoint. Deleting an interface endpoint or a Gateway Load Balancer endpoint
// deletes the endpoint network interfaces. Gateway Load Balancer endpoints can
// only be deleted if the routes that are associated with the endpoint are deleted.
func (c *Client) DeleteVpcEndpoints(ctx context.Context, params *DeleteVpcEndpointsInput, optFns ...func(*Options)) (*DeleteVpcEndpointsOutput, error) {
	if params == nil {
		params = &DeleteVpcEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpcEndpoints", params, optFns, addOperationDeleteVpcEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteVpcEndpoints.
type DeleteVpcEndpointsInput struct {

	// One or more VPC endpoint IDs.
	//
	// This member is required.
	VpcEndpointIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

// Contains the output of DeleteVpcEndpoints.
type DeleteVpcEndpointsOutput struct {

	// Information about the VPC endpoints that were not successfully deleted.
	Unsuccessful []types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVpcEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteVpcEndpoints{}, middleware.After)
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
	if err = addOpDeleteVpcEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVpcEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteVpcEndpoints",
	}
}
