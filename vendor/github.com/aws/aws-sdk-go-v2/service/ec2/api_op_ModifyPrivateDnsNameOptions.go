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

// Modifies the options for instance hostnames for the specified instance.
func (c *Client) ModifyPrivateDnsNameOptions(ctx context.Context, params *ModifyPrivateDnsNameOptionsInput, optFns ...func(*Options)) (*ModifyPrivateDnsNameOptionsOutput, error) {
	if params == nil {
		params = &ModifyPrivateDnsNameOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyPrivateDnsNameOptions", params, optFns, c.addOperationModifyPrivateDnsNameOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyPrivateDnsNameOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyPrivateDnsNameOptionsInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Indicates whether to respond to DNS queries for instance hostnames with DNS
	// AAAA records.
	EnableResourceNameDnsAAAARecord *bool

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A
	// records.
	EnableResourceNameDnsARecord *bool

	// The type of hostname for EC2 instances. For IPv4 only subnets, an instance DNS
	// name must be based on the instance IPv4 address. For IPv6 only subnets, an
	// instance DNS name must be based on the instance ID. For dual-stack subnets, you
	// can specify whether DNS names use the instance IPv4 address or the instance ID.
	PrivateDnsHostnameType types.HostnameType

	noSmithyDocumentSerde
}

type ModifyPrivateDnsNameOptionsOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyPrivateDnsNameOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyPrivateDnsNameOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyPrivateDnsNameOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyPrivateDnsNameOptions"); err != nil {
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
	if err = addOpModifyPrivateDnsNameOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyPrivateDnsNameOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyPrivateDnsNameOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyPrivateDnsNameOptions",
	}
}
