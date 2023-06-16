// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deprovision a CIDR from a public IPv4 pool.
func (c *Client) DeprovisionPublicIpv4PoolCidr(ctx context.Context, params *DeprovisionPublicIpv4PoolCidrInput, optFns ...func(*Options)) (*DeprovisionPublicIpv4PoolCidrOutput, error) {
	if params == nil {
		params = &DeprovisionPublicIpv4PoolCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeprovisionPublicIpv4PoolCidr", params, optFns, c.addOperationDeprovisionPublicIpv4PoolCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeprovisionPublicIpv4PoolCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeprovisionPublicIpv4PoolCidrInput struct {

	// The CIDR you want to deprovision from the pool. Enter the CIDR you want to
	// deprovision with a netmask of /32 . You must rerun this command for each IP
	// address in the CIDR range. If your CIDR is a /24 , you will have to run this
	// command to deprovision each of the 256 IP addresses in the /24 CIDR.
	//
	// This member is required.
	Cidr *string

	// The ID of the pool that you want to deprovision the CIDR from.
	//
	// This member is required.
	PoolId *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DeprovisionPublicIpv4PoolCidrOutput struct {

	// The deprovisioned CIDRs.
	DeprovisionedAddresses []string

	// The ID of the pool that you deprovisioned the CIDR from.
	PoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeprovisionPublicIpv4PoolCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeprovisionPublicIpv4PoolCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeprovisionPublicIpv4PoolCidr{}, middleware.After)
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
	if err = addOpDeprovisionPublicIpv4PoolCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeprovisionPublicIpv4PoolCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeprovisionPublicIpv4PoolCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeprovisionPublicIpv4PoolCidr",
	}
}
