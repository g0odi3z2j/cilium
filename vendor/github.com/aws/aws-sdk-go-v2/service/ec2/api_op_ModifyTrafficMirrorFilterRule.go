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

// Modifies the specified Traffic Mirror rule. DestinationCidrBlock and
// SourceCidrBlock must both be an IPv4 range or an IPv6 range.
func (c *Client) ModifyTrafficMirrorFilterRule(ctx context.Context, params *ModifyTrafficMirrorFilterRuleInput, optFns ...func(*Options)) (*ModifyTrafficMirrorFilterRuleOutput, error) {
	if params == nil {
		params = &ModifyTrafficMirrorFilterRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTrafficMirrorFilterRule", params, optFns, addOperationModifyTrafficMirrorFilterRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTrafficMirrorFilterRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTrafficMirrorFilterRuleInput struct {

	// The ID of the Traffic Mirror rule.
	//
	// This member is required.
	TrafficMirrorFilterRuleId *string

	// The description to assign to the Traffic Mirror rule.
	Description *string

	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string

	// The destination ports that are associated with the Traffic Mirror rule.
	DestinationPortRange *types.TrafficMirrorPortRangeRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The protocol, for example TCP, to assign to the Traffic Mirror rule.
	Protocol int32

	// The properties that you want to remove from the Traffic Mirror filter rule. When
	// you remove a property from a Traffic Mirror filter rule, the property is set to
	// the default.
	RemoveFields []types.TrafficMirrorFilterRuleField

	// The action to assign to the rule.
	RuleAction types.TrafficMirrorRuleAction

	// The number of the Traffic Mirror rule. This number must be unique for each
	// Traffic Mirror rule in a given direction. The rules are processed in ascending
	// order by rule number.
	RuleNumber int32

	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string

	// The port range to assign to the Traffic Mirror rule.
	SourcePortRange *types.TrafficMirrorPortRangeRequest

	// The type of traffic (ingress | egress) to assign to the rule.
	TrafficDirection types.TrafficDirection
}

type ModifyTrafficMirrorFilterRuleOutput struct {

	// Modifies a Traffic Mirror rule.
	TrafficMirrorFilterRule *types.TrafficMirrorFilterRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyTrafficMirrorFilterRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyTrafficMirrorFilterRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTrafficMirrorFilterRule{}, middleware.After)
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
	if err = addOpModifyTrafficMirrorFilterRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTrafficMirrorFilterRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTrafficMirrorFilterRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyTrafficMirrorFilterRule",
	}
}
