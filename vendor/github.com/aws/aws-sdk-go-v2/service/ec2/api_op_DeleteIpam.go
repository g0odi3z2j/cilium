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

// Delete an IPAM. Deleting an IPAM removes all monitored data associated with the
// IPAM including the historical data for CIDRs. For more information, see Delete
// an IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/delete-ipam.html) in the
// Amazon VPC IPAM User Guide.
func (c *Client) DeleteIpam(ctx context.Context, params *DeleteIpamInput, optFns ...func(*Options)) (*DeleteIpamOutput, error) {
	if params == nil {
		params = &DeleteIpamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIpam", params, optFns, c.addOperationDeleteIpamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIpamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIpamInput struct {

	// The ID of the IPAM to delete.
	//
	// This member is required.
	IpamId *string

	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes,
	// and any allocations in the pools in private scopes. You cannot delete the IPAM
	// with this option if there is a pool in your public scope. If you use this
	// option, IPAM does the following:
	//   - Deallocates any CIDRs allocated to VPC resources (such as VPCs) in pools in
	//   private scopes. No VPC resources are deleted as a result of enabling this
	//   option. The CIDR associated with the resource will no longer be allocated from
	//   an IPAM pool, but the CIDR itself will remain unchanged.
	//   - Deprovisions all IPv4 CIDRs provisioned to IPAM pools in private scopes.
	//   - Deletes all IPAM pools in private scopes.
	//   - Deletes all non-default private scopes in the IPAM.
	//   - Deletes the default public and private scopes and the IPAM.
	Cascade *bool

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DeleteIpamOutput struct {

	// Information about the results of the deletion.
	Ipam *types.Ipam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIpamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteIpam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteIpam{}, middleware.After)
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
	if err = addOpDeleteIpamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIpam(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIpam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteIpam",
	}
}
