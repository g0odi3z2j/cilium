// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Decodes additional information about the authorization status of a request from
// an encoded message returned in response to an AWS request. For example, if a
// user is not authorized to perform an operation that he or she has requested, the
// request returns a Client.UnauthorizedOperation response (an HTTP 403 response).
// Some AWS operations additionally return an encoded message that can provide
// details about this authorization failure. Only certain AWS operations return an
// encoded authorization message. The documentation for an individual operation
// indicates whether that operation returns an encoded message in addition to
// returning an HTTP code. The message is encoded because the details of the
// authorization status can constitute privileged information that the user who
// requested the operation should not see. To decode an authorization status
// message, a user must be granted permissions via an IAM policy to request the
// DecodeAuthorizationMessage (sts:DecodeAuthorizationMessage) action. The decoded
// message includes the following type of information:
//
// * Whether the request was
// denied due to an explicit deny or due to the absence of an explicit allow. For
// more information, see Determining Whether a Request is Allowed or Denied
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-denyallow)
// in the IAM User Guide.
//
// * The principal who made the request.
//
// * The requested
// action.
//
// * The requested resource.
//
// * The values of condition keys in the
// context of the user's request.
func (c *Client) DecodeAuthorizationMessage(ctx context.Context, params *DecodeAuthorizationMessageInput, optFns ...func(*Options)) (*DecodeAuthorizationMessageOutput, error) {
	if params == nil {
		params = &DecodeAuthorizationMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecodeAuthorizationMessage", params, optFns, addOperationDecodeAuthorizationMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecodeAuthorizationMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecodeAuthorizationMessageInput struct {

	// The encoded message that was returned with the response.
	//
	// This member is required.
	EncodedMessage *string
}

// A document that contains additional information about the authorization status
// of a request from an encoded message that is returned in response to an AWS
// request.
type DecodeAuthorizationMessageOutput struct {

	// An XML document that contains the decoded message.
	DecodedMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDecodeAuthorizationMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDecodeAuthorizationMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDecodeAuthorizationMessage{}, middleware.After)
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
	if err = addOpDecodeAuthorizationMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecodeAuthorizationMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDecodeAuthorizationMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "DecodeAuthorizationMessage",
	}
}
