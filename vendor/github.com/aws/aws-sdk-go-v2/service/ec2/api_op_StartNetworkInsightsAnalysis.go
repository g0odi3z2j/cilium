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

// Starts analyzing the specified path. If the path is reachable, the operation
// returns the shortest feasible path.
func (c *Client) StartNetworkInsightsAnalysis(ctx context.Context, params *StartNetworkInsightsAnalysisInput, optFns ...func(*Options)) (*StartNetworkInsightsAnalysisOutput, error) {
	if params == nil {
		params = &StartNetworkInsightsAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNetworkInsightsAnalysis", params, optFns, c.addOperationStartNetworkInsightsAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNetworkInsightsAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNetworkInsightsAnalysisInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	//
	// This member is required.
	ClientToken *string

	// The ID of the path.
	//
	// This member is required.
	NetworkInsightsPathId *string

	// The member accounts that contain resources that the path can traverse.
	AdditionalAccounts []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Amazon Resource Names (ARN) of the resources that the path must traverse.
	FilterInArns []string

	// The tags to apply.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type StartNetworkInsightsAnalysisOutput struct {

	// Information about the network insights analysis.
	NetworkInsightsAnalysis *types.NetworkInsightsAnalysis

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartNetworkInsightsAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpStartNetworkInsightsAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpStartNetworkInsightsAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartNetworkInsightsAnalysis"); err != nil {
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
	if err = addIdempotencyToken_opStartNetworkInsightsAnalysisMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartNetworkInsightsAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartNetworkInsightsAnalysis(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartNetworkInsightsAnalysis struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartNetworkInsightsAnalysis) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartNetworkInsightsAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartNetworkInsightsAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartNetworkInsightsAnalysisInput ")
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
func addIdempotencyToken_opStartNetworkInsightsAnalysisMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartNetworkInsightsAnalysis{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartNetworkInsightsAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartNetworkInsightsAnalysis",
	}
}
