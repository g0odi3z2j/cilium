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

// Modifies the configuration of the specified Amazon Web Services Verified Access
// trust provider.
func (c *Client) ModifyVerifiedAccessTrustProvider(ctx context.Context, params *ModifyVerifiedAccessTrustProviderInput, optFns ...func(*Options)) (*ModifyVerifiedAccessTrustProviderOutput, error) {
	if params == nil {
		params = &ModifyVerifiedAccessTrustProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVerifiedAccessTrustProvider", params, optFns, c.addOperationModifyVerifiedAccessTrustProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVerifiedAccessTrustProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVerifiedAccessTrustProviderInput struct {

	// The ID of the Verified Access trust provider.
	//
	// This member is required.
	VerifiedAccessTrustProviderId *string

	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// A description for the Verified Access trust provider.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The options for an OpenID Connect-compatible user-identity trust provider.
	OidcOptions *types.ModifyVerifiedAccessTrustProviderOidcOptions

	noSmithyDocumentSerde
}

type ModifyVerifiedAccessTrustProviderOutput struct {

	// The ID of the Verified Access trust provider.
	VerifiedAccessTrustProvider *types.VerifiedAccessTrustProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVerifiedAccessTrustProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVerifiedAccessTrustProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVerifiedAccessTrustProvider{}, middleware.After)
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
	if err = addIdempotencyToken_opModifyVerifiedAccessTrustProviderMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyVerifiedAccessTrustProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVerifiedAccessTrustProvider(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpModifyVerifiedAccessTrustProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpModifyVerifiedAccessTrustProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpModifyVerifiedAccessTrustProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ModifyVerifiedAccessTrustProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ModifyVerifiedAccessTrustProviderInput ")
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
func addIdempotencyToken_opModifyVerifiedAccessTrustProviderMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpModifyVerifiedAccessTrustProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opModifyVerifiedAccessTrustProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVerifiedAccessTrustProvider",
	}
}
