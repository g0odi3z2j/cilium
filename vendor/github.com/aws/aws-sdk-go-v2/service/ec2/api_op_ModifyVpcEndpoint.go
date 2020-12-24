// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies attributes of a specified VPC endpoint. The attributes that you can
// modify depend on the type of VPC endpoint (interface, gateway, or Gateway Load
// Balancer). For more information, see VPC Endpoints
// (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) ModifyVpcEndpoint(ctx context.Context, params *ModifyVpcEndpointInput, optFns ...func(*Options)) (*ModifyVpcEndpointOutput, error) {
	if params == nil {
		params = &ModifyVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcEndpoint", params, optFns, addOperationModifyVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyVpcEndpoint.
type ModifyVpcEndpointInput struct {

	// The ID of the endpoint.
	//
	// This member is required.
	VpcEndpointId *string

	// (Gateway endpoint) One or more route tables IDs to associate with the endpoint.
	AddRouteTableIds []string

	// (Interface endpoint) One or more security group IDs to associate with the
	// network interface.
	AddSecurityGroupIds []string

	// (Interface and Gateway Load Balancer endpoints) One or more subnet IDs in which
	// to serve the endpoint. For a Gateway Load Balancer endpoint, you can specify
	// only one subnet.
	AddSubnetIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// (Interface and gateway endpoints) A policy to attach to the endpoint that
	// controls access to the service. The policy must be in valid JSON format.
	PolicyDocument *string

	// (Interface endpoint) Indicates whether a private hosted zone is associated with
	// the VPC.
	PrivateDnsEnabled bool

	// (Gateway endpoint) One or more route table IDs to disassociate from the
	// endpoint.
	RemoveRouteTableIds []string

	// (Interface endpoint) One or more security group IDs to disassociate from the
	// network interface.
	RemoveSecurityGroupIds []string

	// (Interface endpoint) One or more subnets IDs in which to remove the endpoint.
	RemoveSubnetIds []string

	// (Gateway endpoint) Specify true to reset the policy document to the default
	// policy. The default policy allows full access to the service.
	ResetPolicy bool
}

type ModifyVpcEndpointOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpoint{}, middleware.After)
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
	if err = addOpModifyVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpoint",
	}
}
