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

// Creates an Amazon FPGA Image (AFI) from the specified design checkpoint (DCP).
// The create operation is asynchronous. To verify that the AFI is ready for use,
// check the output logs. An AFI contains the FPGA bitstream that is ready to
// download to an FPGA. You can securely deploy an AFI on multiple FPGA-accelerated
// instances. For more information, see the AWS FPGA Hardware Development Kit
// (https://github.com/aws/aws-fpga/).
func (c *Client) CreateFpgaImage(ctx context.Context, params *CreateFpgaImageInput, optFns ...func(*Options)) (*CreateFpgaImageOutput, error) {
	if params == nil {
		params = &CreateFpgaImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFpgaImage", params, optFns, addOperationCreateFpgaImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFpgaImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFpgaImageInput struct {

	// The location of the encrypted design checkpoint in Amazon S3. The input must be
	// a tarball.
	//
	// This member is required.
	InputStorageLocation *types.StorageLocation

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	// A description for the AFI.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The location in Amazon S3 for the output logs.
	LogsStorageLocation *types.StorageLocation

	// A name for the AFI.
	Name *string

	// The tags to apply to the FPGA image during creation.
	TagSpecifications []types.TagSpecification
}

type CreateFpgaImageOutput struct {

	// The global FPGA image identifier (AGFI ID).
	FpgaImageGlobalId *string

	// The FPGA image identifier (AFI ID).
	FpgaImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFpgaImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateFpgaImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFpgaImage{}, middleware.After)
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
	if err = addOpCreateFpgaImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFpgaImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFpgaImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateFpgaImage",
	}
}
