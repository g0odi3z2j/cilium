// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a CIDR block with your subnet. You can only associate a single IPv6
// CIDR block with your subnet.
func (c *Client) AssociateSubnetCidrBlock(ctx context.Context, params *AssociateSubnetCidrBlockInput, optFns ...func(*Options)) (*AssociateSubnetCidrBlockOutput, error) {
	if params == nil {
		params = &AssociateSubnetCidrBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSubnetCidrBlock", params, optFns, c.addOperationAssociateSubnetCidrBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSubnetCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSubnetCidrBlockInput struct {

	// The ID of your subnet.
	//
	// This member is required.
	SubnetId *string

	// The IPv6 CIDR block for your subnet.
	Ipv6CidrBlock *string

	// An IPv6 IPAM pool ID.
	Ipv6IpamPoolId *string

	// An IPv6 netmask length.
	Ipv6NetmaskLength *int32

	noSmithyDocumentSerde
}

type AssociateSubnetCidrBlockOutput struct {

	// Information about the IPv6 association.
	Ipv6CidrBlockAssociation *types.SubnetIpv6CidrBlockAssociation

	// The ID of the subnet.
	SubnetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateSubnetCidrBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateSubnetCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateSubnetCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateSubnetCidrBlock"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateSubnetCidrBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSubnetCidrBlock(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateSubnetCidrBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateSubnetCidrBlock",
	}
}
