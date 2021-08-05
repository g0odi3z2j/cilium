// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a 2048-bit RSA key pair with the specified name. Amazon EC2 stores the
// public key and displays the private key for you to save to a file. The private
// key is returned as an unencrypted PEM encoded PKCS#1 private key. If a key with
// the specified name already exists, Amazon EC2 returns an error. You can have up
// to five thousand key pairs per Region. The key pair returned to you is available
// only in the Region in which you create it. If you prefer, you can create your
// own key pair using a third-party tool and upload it to any Region using
// ImportKeyPair. For more information, see Key Pairs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateKeyPair(ctx context.Context, params *CreateKeyPairInput, optFns ...func(*Options)) (*CreateKeyPairOutput, error) {
	if params == nil {
		params = &CreateKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKeyPair", params, optFns, c.addOperationCreateKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyPairInput struct {

	// A unique name for the key pair. Constraints: Up to 255 ASCII characters
	//
	// This member is required.
	KeyName *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to apply to the new key pair.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

// Describes a key pair.
type CreateKeyPairOutput struct {

	// The SHA-1 digest of the DER encoded private key.
	KeyFingerprint *string

	// An unencrypted PEM encoded RSA private key.
	KeyMaterial *string

	// The name of the key pair.
	KeyName *string

	// The ID of the key pair.
	KeyPairId *string

	// Any tags applied to the key pair.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateKeyPair{}, middleware.After)
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
	if err = addOpCreateKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateKeyPair",
	}
}
