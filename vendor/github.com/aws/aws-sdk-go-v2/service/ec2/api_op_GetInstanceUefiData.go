// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A binary representation of the UEFI variable store. Only non-volatile variables
// are stored. This is a base64 encoded and zlib compressed binary value that must
// be properly encoded. When you use register-image
// (https://docs.aws.amazon.com/cli/latest/reference/ec2/register-image.html) to
// create an AMI, you can create an exact copy of your variable store by passing
// the UEFI data in the UefiData parameter. You can modify the UEFI data by using
// the python-uefivars tool (https://github.com/awslabs/python-uefivars) on GitHub.
// You can use the tool to convert the UEFI data into a human-readable format
// (JSON), which you can inspect and modify, and then convert back into the binary
// format to use with register-image. For more information, see UEFI Secure Boot
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html) in
// the Amazon EC2 User Guide.
func (c *Client) GetInstanceUefiData(ctx context.Context, params *GetInstanceUefiDataInput, optFns ...func(*Options)) (*GetInstanceUefiDataOutput, error) {
	if params == nil {
		params = &GetInstanceUefiDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceUefiData", params, optFns, c.addOperationGetInstanceUefiDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceUefiDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceUefiDataInput struct {

	// The ID of the instance from which to retrieve the UEFI data.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type GetInstanceUefiDataOutput struct {

	// The ID of the instance from which to retrieve the UEFI data.
	InstanceId *string

	// Base64 representation of the non-volatile UEFI variable store.
	UefiData *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInstanceUefiDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetInstanceUefiData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetInstanceUefiData{}, middleware.After)
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
	if err = addOpGetInstanceUefiDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceUefiData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInstanceUefiData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetInstanceUefiData",
	}
}
