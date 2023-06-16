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

// Creates a local gateway route table virtual interface group association.
func (c *Client) CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx context.Context, params *CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, optFns ...func(*Options)) (*CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	if params == nil {
		params = &CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation", params, optFns, c.addOperationCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput struct {

	// The ID of the local gateway route table.
	//
	// This member is required.
	LocalGatewayRouteTableId *string

	// The ID of the local gateway route table virtual interface group association.
	//
	// This member is required.
	LocalGatewayVirtualInterfaceGroupId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags assigned to the local gateway route table virtual interface group
	// association.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput struct {

	// Information about the local gateway route table virtual interface group
	// association.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociation *types.LocalGatewayRouteTableVirtualInterfaceGroupAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation{}, middleware.After)
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
	if err = addOpCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation",
	}
}
