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

// Requests a VPC peering connection between two VPCs: a requester VPC that you own
// and an accepter VPC with which to create the connection. The accepter VPC can
// belong to another AWS account and can be in a different Region to the requester
// VPC. The requester VPC and accepter VPC cannot have overlapping CIDR blocks.
// Limitations and rules apply to a VPC peering connection. For more information,
// see the limitations
// (https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations)
// section in the VPC Peering Guide. The owner of the accepter VPC must accept the
// peering request to activate the peering connection. The VPC peering connection
// request expires after 7 days, after which it cannot be accepted or rejected. If
// you create a VPC peering connection request between VPCs with overlapping CIDR
// blocks, the VPC peering connection has a status of failed.
func (c *Client) CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &CreateVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcPeeringConnection", params, optFns, addOperationCreateVpcPeeringConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcPeeringConnectionInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The AWS account ID of the owner of the accepter VPC. Default: Your AWS account
	// ID
	PeerOwnerId *string

	// The Region code for the accepter VPC, if the accepter VPC is located in a Region
	// other than the Region in which you make the request. Default: The Region in
	// which you make the request.
	PeerRegion *string

	// The ID of the VPC with which you are creating the VPC peering connection. You
	// must specify this parameter in the request.
	PeerVpcId *string

	// The tags to assign to the peering connection.
	TagSpecifications []types.TagSpecification

	// The ID of the requester VPC. You must specify this parameter in the request.
	VpcId *string
}

type CreateVpcPeeringConnectionOutput struct {

	// Information about the VPC peering connection.
	VpcPeeringConnection *types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVpcPeeringConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcPeeringConnection{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcPeeringConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpcPeeringConnection",
	}
}
