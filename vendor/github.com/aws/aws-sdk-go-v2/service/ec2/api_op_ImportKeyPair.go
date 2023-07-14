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

// Imports the public key from an RSA or ED25519 key pair that you created with a
// third-party tool. Compare this with CreateKeyPair , in which Amazon Web Services
// creates the key pair and gives the keys to you (Amazon Web Services keeps a copy
// of the public key). With ImportKeyPair, you create the key pair and give Amazon
// Web Services just the public key. The private key is never transferred between
// you and Amazon Web Services. For more information about key pairs, see Amazon
// EC2 key pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ImportKeyPair(ctx context.Context, params *ImportKeyPairInput, optFns ...func(*Options)) (*ImportKeyPairOutput, error) {
	if params == nil {
		params = &ImportKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportKeyPair", params, optFns, c.addOperationImportKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportKeyPairInput struct {

	// A unique name for the key pair.
	//
	// This member is required.
	KeyName *string

	// The public key. For API calls, the text must be base64-encoded. For command
	// line tools, base64 encoding is performed for you.
	//
	// This member is required.
	PublicKeyMaterial []byte

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to apply to the imported key pair.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type ImportKeyPairOutput struct {

	//   - For RSA key pairs, the key fingerprint is the MD5 public key fingerprint as
	//   specified in section 4 of RFC 4716.
	//   - For ED25519 key pairs, the key fingerprint is the base64-encoded SHA-256
	//   digest, which is the default for OpenSSH, starting with OpenSSH 6.8 (http://www.openssh.com/txt/release-6.8)
	//   .
	KeyFingerprint *string

	// The key pair name that you provided.
	KeyName *string

	// The ID of the resulting key pair.
	KeyPairId *string

	// The tags applied to the imported key pair.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpImportKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpImportKeyPair{}, middleware.After)
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
	if err = addOpImportKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ImportKeyPair",
	}
}
