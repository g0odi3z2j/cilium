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

// Modifies the specified EC2 Fleet. You can only modify an EC2 Fleet request of
// type maintain . While the EC2 Fleet is being modified, it is in the modifying
// state. To scale up your EC2 Fleet, increase its target capacity. The EC2 Fleet
// launches the additional Spot Instances according to the allocation strategy for
// the EC2 Fleet request. If the allocation strategy is lowest-price , the EC2
// Fleet launches instances using the Spot Instance pool with the lowest price. If
// the allocation strategy is diversified , the EC2 Fleet distributes the instances
// across the Spot Instance pools. If the allocation strategy is capacity-optimized
// , EC2 Fleet launches instances from Spot Instance pools with optimal capacity
// for the number of instances that are launching. To scale down your EC2 Fleet,
// decrease its target capacity. First, the EC2 Fleet cancels any open requests
// that exceed the new target capacity. You can request that the EC2 Fleet
// terminate Spot Instances until the size of the fleet no longer exceeds the new
// target capacity. If the allocation strategy is lowest-price , the EC2 Fleet
// terminates the instances with the highest price per unit. If the allocation
// strategy is capacity-optimized , the EC2 Fleet terminates the instances in the
// Spot Instance pools that have the least available Spot Instance capacity. If the
// allocation strategy is diversified , the EC2 Fleet terminates instances across
// the Spot Instance pools. Alternatively, you can request that the EC2 Fleet keep
// the fleet at its current size, but not replace any Spot Instances that are
// interrupted or that you terminate manually. If you are finished with your EC2
// Fleet for now, but will use it again later, you can set the target capacity to
// 0.
func (c *Client) ModifyFleet(ctx context.Context, params *ModifyFleetInput, optFns ...func(*Options)) (*ModifyFleetOutput, error) {
	if params == nil {
		params = &ModifyFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyFleet", params, optFns, c.addOperationModifyFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyFleetInput struct {

	// The ID of the EC2 Fleet.
	//
	// This member is required.
	FleetId *string

	// Reserved.
	Context *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Indicates whether running instances should be terminated if the total target
	// capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	// Supported only for fleets of type maintain .
	ExcessCapacityTerminationPolicy types.FleetExcessCapacityTerminationPolicy

	// The launch template and overrides.
	LaunchTemplateConfigs []types.FleetLaunchTemplateConfigRequest

	// The size of the EC2 Fleet.
	TargetCapacitySpecification *types.TargetCapacitySpecificationRequest

	noSmithyDocumentSerde
}

type ModifyFleetOutput struct {

	// If the request succeeds, the response returns true . If the request fails, no
	// response is returned, and instead an error message is returned.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyFleet{}, middleware.After)
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
	if err = addOpModifyFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyFleet(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opModifyFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyFleet",
	}
}
