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

// Returns the IAM roles that are associated with the specified AWS Certificate
// Manager (ACM) certificate. It also returns the name of the Amazon S3 bucket and
// the Amazon S3 object key where the certificate, certificate chain, and encrypted
// private key bundle are stored, and the ARN of the AWS Key Management Service
// (KMS) customer master key (CMK) that's used to encrypt the private key.
func (c *Client) GetAssociatedEnclaveCertificateIamRoles(ctx context.Context, params *GetAssociatedEnclaveCertificateIamRolesInput, optFns ...func(*Options)) (*GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	if params == nil {
		params = &GetAssociatedEnclaveCertificateIamRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssociatedEnclaveCertificateIamRoles", params, optFns, addOperationGetAssociatedEnclaveCertificateIamRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssociatedEnclaveCertificateIamRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociatedEnclaveCertificateIamRolesInput struct {

	// The ARN of the ACM certificate for which to view the associated IAM roles,
	// encryption keys, and Amazon S3 object information.
	CertificateArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type GetAssociatedEnclaveCertificateIamRolesOutput struct {

	// Information about the associated IAM roles.
	AssociatedRoles []types.AssociatedRole

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAssociatedEnclaveCertificateIamRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetAssociatedEnclaveCertificateIamRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetAssociatedEnclaveCertificateIamRoles{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociatedEnclaveCertificateIamRoles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAssociatedEnclaveCertificateIamRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetAssociatedEnclaveCertificateIamRoles",
	}
}
