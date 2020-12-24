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

// Modifies a Traffic Mirror session.
func (c *Client) ModifyTrafficMirrorSession(ctx context.Context, params *ModifyTrafficMirrorSessionInput, optFns ...func(*Options)) (*ModifyTrafficMirrorSessionOutput, error) {
	if params == nil {
		params = &ModifyTrafficMirrorSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTrafficMirrorSession", params, optFns, addOperationModifyTrafficMirrorSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTrafficMirrorSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTrafficMirrorSessionInput struct {

	// The ID of the Traffic Mirror session.
	//
	// This member is required.
	TrafficMirrorSessionId *string

	// The description to assign to the Traffic Mirror session.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The number of bytes in each packet to mirror. These are bytes after the VXLAN
	// header. To mirror a subset, set this to the length (in bytes) to mirror. For
	// example, if you set this value to 100, then the first 100 bytes that meet the
	// filter criteria are copied to the target. Do not specify this parameter when you
	// want to mirror the entire packet.
	PacketLength int32

	// The properties that you want to remove from the Traffic Mirror session. When you
	// remove a property from a Traffic Mirror session, the property is set to the
	// default.
	RemoveFields []types.TrafficMirrorSessionField

	// The session number determines the order in which sessions are evaluated when an
	// interface is used by multiple sessions. The first session with a matching filter
	// is the one that mirrors the packets. Valid values are 1-32766.
	SessionNumber int32

	// The ID of the Traffic Mirror filter.
	TrafficMirrorFilterId *string

	// The Traffic Mirror target. The target must be in the same VPC as the source, or
	// have a VPC peering connection with the source.
	TrafficMirrorTargetId *string

	// The virtual network ID of the Traffic Mirror session.
	VirtualNetworkId int32
}

type ModifyTrafficMirrorSessionOutput struct {

	// Information about the Traffic Mirror session.
	TrafficMirrorSession *types.TrafficMirrorSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyTrafficMirrorSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyTrafficMirrorSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTrafficMirrorSession{}, middleware.After)
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
	if err = addOpModifyTrafficMirrorSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTrafficMirrorSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTrafficMirrorSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyTrafficMirrorSession",
	}
}
