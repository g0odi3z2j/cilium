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

// Describes the ID format settings for your resources on a per-Region basis, for
// example, to view which resource types are enabled for longer IDs. This request
// only returns information about resource types whose ID formats can be modified;
// it does not return information about other resource types. The following
// resource types support longer IDs: bundle | conversion-task | customer-gateway
// | dhcp-options | elastic-ip-allocation | elastic-ip-association | export-task |
// flow-log | image | import-task | instance | internet-gateway | network-acl |
// network-acl-association | network-interface | network-interface-attachment |
// prefix-list | reservation | route-table | route-table-association |
// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
// vpn-connection | vpn-gateway . These settings apply to the IAM user who makes
// the request; they do not apply to the entire Amazon Web Services account. By
// default, an IAM user defaults to the same settings as the root user, unless they
// explicitly override the settings by running the ModifyIdFormat command.
// Resources created with longer IDs are visible to all IAM users, regardless of
// these settings and provided that they have permission to use the relevant
// Describe command for the resource type.
func (c *Client) DescribeIdFormat(ctx context.Context, params *DescribeIdFormatInput, optFns ...func(*Options)) (*DescribeIdFormatOutput, error) {
	if params == nil {
		params = &DescribeIdFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdFormat", params, optFns, c.addOperationDescribeIdFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIdFormatInput struct {

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log |
	// image | import-task | instance | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | reservation | route-table | route-table-association |
	// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
	// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
	// vpn-connection | vpn-gateway
	Resource *string

	noSmithyDocumentSerde
}

type DescribeIdFormatOutput struct {

	// Information about the ID format for the resource.
	Statuses []types.IdFormat

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIdFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIdFormat{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdFormat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIdFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeIdFormat",
	}
}
