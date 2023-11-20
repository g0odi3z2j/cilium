// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and returns access and refresh tokens for clients and applications that
// are authenticated using IAM entities. The access token can be used to fetch
// short-term credentials for the assigned AWS accounts or to access application
// APIs using bearer authentication.
func (c *Client) CreateTokenWithIAM(ctx context.Context, params *CreateTokenWithIAMInput, optFns ...func(*Options)) (*CreateTokenWithIAMOutput, error) {
	if params == nil {
		params = &CreateTokenWithIAMInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTokenWithIAM", params, optFns, c.addOperationCreateTokenWithIAMMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTokenWithIAMOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTokenWithIAMInput struct {

	// The unique identifier string for the client or application. This value is an
	// application ARN that has OAuth grants configured.
	//
	// This member is required.
	ClientId *string

	// Supports the following OAuth grant types: Authorization Code, Refresh Token,
	// JWT Bearer, and Token Exchange. Specify one of the following values, depending
	// on the grant type that you want: * Authorization Code - authorization_code *
	// Refresh Token - refresh_token * JWT Bearer -
	// urn:ietf:params:oauth:grant-type:jwt-bearer * Token Exchange -
	// urn:ietf:params:oauth:grant-type:token-exchange
	//
	// This member is required.
	GrantType *string

	// Used only when calling this API for the JWT Bearer grant type. This value
	// specifies the JSON Web Token (JWT) issued by a trusted token issuer. To
	// authorize a trusted token issuer, configure the JWT Bearer GrantOptions for the
	// application.
	Assertion *string

	// Used only when calling this API for the Authorization Code grant type. This
	// short-term code is used to identify this authorization request. The code is
	// obtained through a redirect from IAM Identity Center to a redirect URI persisted
	// in the Authorization Code GrantOptions for the application.
	Code *string

	// Used only when calling this API for the Authorization Code grant type. This
	// value specifies the location of the client or application that has registered to
	// receive the authorization code.
	RedirectUri *string

	// Used only when calling this API for the Refresh Token grant type. This token is
	// used to refresh short-term tokens, such as the access token, that might expire.
	// For more information about the features and limitations of the current IAM
	// Identity Center OIDC implementation, see Considerations for Using this Guide in
	// the IAM Identity Center OIDC API Reference (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/Welcome.html)
	// .
	RefreshToken *string

	// Used only when calling this API for the Token Exchange grant type. This value
	// specifies the type of token that the requester can receive. The following values
	// are supported: * Access Token - urn:ietf:params:oauth:token-type:access_token *
	// Refresh Token - urn:ietf:params:oauth:token-type:refresh_token
	RequestedTokenType *string

	// The list of scopes for which authorization is requested. The access token that
	// is issued is limited to the scopes that are granted. If the value is not
	// specified, IAM Identity Center authorizes all scopes configured for the
	// application, including the following default scopes: openid , aws ,
	// sts:identity_context .
	Scope []string

	// Used only when calling this API for the Token Exchange grant type. This value
	// specifies the subject of the exchange. The value of the subject token must be an
	// access token issued by IAM Identity Center to a different client or application.
	// The access token must have authorized scopes that indicate the requested
	// application as a target audience.
	SubjectToken *string

	// Used only when calling this API for the Token Exchange grant type. This value
	// specifies the type of token that is passed as the subject of the exchange. The
	// following value is supported: * Access Token -
	// urn:ietf:params:oauth:token-type:access_token
	SubjectTokenType *string

	noSmithyDocumentSerde
}

type CreateTokenWithIAMOutput struct {

	// A bearer token to access AWS accounts and applications assigned to a user.
	AccessToken *string

	// Indicates the time in seconds when an access token will expire.
	ExpiresIn int32

	// A JSON Web Token (JWT) that identifies the user associated with the issued
	// access token.
	IdToken *string

	// Indicates the type of tokens that are issued by IAM Identity Center. The
	// following values are supported: * Access Token -
	// urn:ietf:params:oauth:token-type:access_token * Refresh Token -
	// urn:ietf:params:oauth:token-type:refresh_token
	IssuedTokenType *string

	// A token that, if present, can be used to refresh a previously issued access
	// token that might have expired. For more information about the features and
	// limitations of the current IAM Identity Center OIDC implementation, see
	// Considerations for Using this Guide in the IAM Identity Center OIDC API
	// Reference (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/Welcome.html)
	// .
	RefreshToken *string

	// The list of scopes for which authorization is granted. The access token that is
	// issued is limited to the scopes that are granted.
	Scope []string

	// Used to notify the requester that the returned token is an access token. The
	// supported token type is Bearer .
	TokenType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTokenWithIAMMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTokenWithIAM{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTokenWithIAM{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTokenWithIAM"); err != nil {
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
	if err = addOpCreateTokenWithIAMValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTokenWithIAM(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTokenWithIAM(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTokenWithIAM",
	}
}
