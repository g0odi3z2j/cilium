// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the attributes of your VPC endpoint service configuration. You can
// change the Network Load Balancers or Gateway Load Balancers for your service,
// and you can specify whether acceptance is required for requests to connect to
// your endpoint service through an interface VPC endpoint. If you set or modify
// the private DNS name, you must prove that you own the private DNS domain name.
// For more information, see VPC Endpoint Service Private DNS Name Verification
// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) ModifyVpcEndpointServiceConfiguration(ctx context.Context, params *ModifyVpcEndpointServiceConfigurationInput, optFns ...func(*Options)) (*ModifyVpcEndpointServiceConfigurationOutput, error) {
	if params == nil {
		params = &ModifyVpcEndpointServiceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcEndpointServiceConfiguration", params, optFns, addOperationModifyVpcEndpointServiceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcEndpointServiceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointServiceConfigurationInput struct {

	// The ID of the service.
	//
	// This member is required.
	ServiceId *string

	// Indicates whether requests to create an endpoint to your service must be
	// accepted.
	AcceptanceRequired bool

	// The Amazon Resource Names (ARNs) of Gateway Load Balancers to add to your
	// service configuration.
	AddGatewayLoadBalancerArns []string

	// The Amazon Resource Names (ARNs) of Network Load Balancers to add to your
	// service configuration.
	AddNetworkLoadBalancerArns []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// (Interface endpoint configuration) The private DNS name to assign to the
	// endpoint service.
	PrivateDnsName *string

	// The Amazon Resource Names (ARNs) of Gateway Load Balancers to remove from your
	// service configuration.
	RemoveGatewayLoadBalancerArns []string

	// The Amazon Resource Names (ARNs) of Network Load Balancers to remove from your
	// service configuration.
	RemoveNetworkLoadBalancerArns []string

	// (Interface endpoint configuration) Removes the private DNS name of the endpoint
	// service.
	RemovePrivateDnsName bool
}

type ModifyVpcEndpointServiceConfigurationOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcEndpointServiceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpointServiceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpointServiceConfiguration{}, middleware.After)
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
	if err = addOpModifyVpcEndpointServiceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpointServiceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcEndpointServiceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointServiceConfiguration",
	}
}
