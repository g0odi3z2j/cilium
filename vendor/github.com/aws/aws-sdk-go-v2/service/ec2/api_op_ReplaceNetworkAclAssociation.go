// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes which network ACL a subnet is associated with. By default when you
// create a subnet, it's automatically associated with the default network ACL. For
// more information, see Network ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html)
// in the Amazon Virtual Private Cloud User Guide. This is an idempotent operation.
func (c *Client) ReplaceNetworkAclAssociation(ctx context.Context, params *ReplaceNetworkAclAssociationInput, optFns ...func(*Options)) (*ReplaceNetworkAclAssociationOutput, error) {
	if params == nil {
		params = &ReplaceNetworkAclAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceNetworkAclAssociation", params, optFns, c.addOperationReplaceNetworkAclAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceNetworkAclAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceNetworkAclAssociationInput struct {

	// The ID of the current association between the original network ACL and the
	// subnet.
	//
	// This member is required.
	AssociationId *string

	// The ID of the new network ACL to associate with the subnet.
	//
	// This member is required.
	NetworkAclId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ReplaceNetworkAclAssociationOutput struct {

	// The ID of the new association.
	NewAssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplaceNetworkAclAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReplaceNetworkAclAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceNetworkAclAssociation{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpReplaceNetworkAclAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceNetworkAclAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplaceNetworkAclAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceNetworkAclAssociation",
	}
}
