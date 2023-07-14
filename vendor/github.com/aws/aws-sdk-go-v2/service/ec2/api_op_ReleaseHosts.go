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

// When you no longer want to use an On-Demand Dedicated Host it can be released.
// On-Demand billing is stopped and the host goes into released state. The host ID
// of Dedicated Hosts that have been released can no longer be specified in another
// request, for example, to modify the host. You must stop or terminate all
// instances on a host before it can be released. When Dedicated Hosts are
// released, it may take some time for them to stop counting toward your limit and
// you may receive capacity errors when trying to allocate new Dedicated Hosts.
// Wait a few minutes and then try again. Released hosts still appear in a
// DescribeHosts response.
func (c *Client) ReleaseHosts(ctx context.Context, params *ReleaseHostsInput, optFns ...func(*Options)) (*ReleaseHostsOutput, error) {
	if params == nil {
		params = &ReleaseHostsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleaseHosts", params, optFns, c.addOperationReleaseHostsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleaseHostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleaseHostsInput struct {

	// The IDs of the Dedicated Hosts to release.
	//
	// This member is required.
	HostIds []string

	noSmithyDocumentSerde
}

type ReleaseHostsOutput struct {

	// The IDs of the Dedicated Hosts that were successfully released.
	Successful []string

	// The IDs of the Dedicated Hosts that could not be released, including an error
	// message.
	Unsuccessful []types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReleaseHostsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReleaseHosts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReleaseHosts{}, middleware.After)
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
	if err = addOpReleaseHostsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReleaseHosts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReleaseHosts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReleaseHosts",
	}
}
