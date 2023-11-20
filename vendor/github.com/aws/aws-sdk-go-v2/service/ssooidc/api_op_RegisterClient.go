// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a client with IAM Identity Center. This allows clients to initiate
// device authorization. The output should be persisted for reuse through many
// authentication requests.
func (c *Client) RegisterClient(ctx context.Context, params *RegisterClientInput, optFns ...func(*Options)) (*RegisterClientOutput, error) {
	if params == nil {
		params = &RegisterClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterClient", params, optFns, c.addOperationRegisterClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterClientInput struct {

	// The friendly name of the client.
	//
	// This member is required.
	ClientName *string

	// The type of client. The service supports only public as a client type. Anything
	// other than public will be rejected by the service.
	//
	// This member is required.
	ClientType *string

	// The list of scopes that are defined by the client. Upon authorization, this
	// list is used to restrict permissions when granting an access token.
	Scopes []string

	noSmithyDocumentSerde
}

type RegisterClientOutput struct {

	// An endpoint that the client can use to request authorization.
	AuthorizationEndpoint *string

	// The unique identifier string for each client. This client uses this identifier
	// to get authenticated by the service in subsequent calls.
	ClientId *string

	// Indicates the time at which the clientId and clientSecret were issued.
	ClientIdIssuedAt int64

	// A secret string generated for the client. The client will use this string to
	// get authenticated by the service in subsequent calls.
	ClientSecret *string

	// Indicates the time at which the clientId and clientSecret will become invalid.
	ClientSecretExpiresAt int64

	// An endpoint that the client can use to create tokens.
	TokenEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterClient{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterClient"); err != nil {
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
	if err = addOpRegisterClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterClient",
	}
}
