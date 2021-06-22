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

// Adds an ingress authorization rule to a Client VPN endpoint. Ingress
// authorization rules act as firewall rules that grant access to networks. You
// must configure ingress authorization rules to enable clients to access resources
// in AWS or on-premises networks.
func (c *Client) AuthorizeClientVpnIngress(ctx context.Context, params *AuthorizeClientVpnIngressInput, optFns ...func(*Options)) (*AuthorizeClientVpnIngressOutput, error) {
	if params == nil {
		params = &AuthorizeClientVpnIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeClientVpnIngress", params, optFns, addOperationAuthorizeClientVpnIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeClientVpnIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeClientVpnIngressInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IPv4 address range, in CIDR notation, of the network for which access is
	// being authorized.
	//
	// This member is required.
	TargetNetworkCidr *string

	// The ID of the group to grant access to, for example, the Active Directory group
	// or identity provider (IdP) group. Required if AuthorizeAllGroups is false or not
	// specified.
	AccessGroupId *string

	// Indicates whether to grant access to all clients. Specify true to grant all
	// clients who successfully establish a VPN connection access to the network. Must
	// be set to true if AccessGroupId is not specified.
	AuthorizeAllGroups *bool

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// A brief description of the authorization rule.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AuthorizeClientVpnIngressOutput struct {

	// The current state of the authorization rule.
	Status *types.ClientVpnAuthorizationRuleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAuthorizeClientVpnIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAuthorizeClientVpnIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAuthorizeClientVpnIngress{}, middleware.After)
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
	if err = addIdempotencyToken_opAuthorizeClientVpnIngressMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAuthorizeClientVpnIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeClientVpnIngress(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpAuthorizeClientVpnIngress struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAuthorizeClientVpnIngress) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAuthorizeClientVpnIngress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AuthorizeClientVpnIngressInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AuthorizeClientVpnIngressInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAuthorizeClientVpnIngressMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAuthorizeClientVpnIngress{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAuthorizeClientVpnIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AuthorizeClientVpnIngress",
	}
}
