// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified Amazon Web Services Verified Access endpoint policy.
func (c *Client) ModifyVerifiedAccessEndpointPolicy(ctx context.Context, params *ModifyVerifiedAccessEndpointPolicyInput, optFns ...func(*Options)) (*ModifyVerifiedAccessEndpointPolicyOutput, error) {
	if params == nil {
		params = &ModifyVerifiedAccessEndpointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVerifiedAccessEndpointPolicy", params, optFns, c.addOperationModifyVerifiedAccessEndpointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVerifiedAccessEndpointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVerifiedAccessEndpointPolicyInput struct {

	// The status of the Verified Access policy.
	//
	// This member is required.
	PolicyEnabled *bool

	// The ID of the Verified Access endpoint.
	//
	// This member is required.
	VerifiedAccessEndpointId *string

	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Verified Access policy document.
	PolicyDocument *string

	noSmithyDocumentSerde
}

type ModifyVerifiedAccessEndpointPolicyOutput struct {

	// The Verified Access policy document.
	PolicyDocument *string

	// The status of the Verified Access policy.
	PolicyEnabled *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVerifiedAccessEndpointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVerifiedAccessEndpointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVerifiedAccessEndpointPolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opModifyVerifiedAccessEndpointPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyVerifiedAccessEndpointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVerifiedAccessEndpointPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpModifyVerifiedAccessEndpointPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpModifyVerifiedAccessEndpointPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpModifyVerifiedAccessEndpointPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ModifyVerifiedAccessEndpointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ModifyVerifiedAccessEndpointPolicyInput ")
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
func addIdempotencyToken_opModifyVerifiedAccessEndpointPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpModifyVerifiedAccessEndpointPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opModifyVerifiedAccessEndpointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVerifiedAccessEndpointPolicy",
	}
}
