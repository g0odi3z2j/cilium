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

// [VPC only] Adds the specified outbound (egress) rules to a security group for
// use with a VPC. An outbound rule permits instances to send traffic to the
// specified IPv4 or IPv6 CIDR address ranges, or to the instances that are
// associated with the specified source security groups. When specifying an
// outbound rule for your security group in a VPC, the IpPermissions must include
// a destination for the traffic. You specify a protocol for each rule (for
// example, TCP). For the TCP and UDP protocols, you must also specify the
// destination port or port range. For the ICMP protocol, you must also specify the
// ICMP type and code. You can use -1 for the type or code to mean all types or all
// codes. Rule changes are propagated to affected instances as quickly as possible.
// However, a small delay might occur. For information about VPC security group
// quotas, see Amazon VPC quotas (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html)
// .
func (c *Client) AuthorizeSecurityGroupEgress(ctx context.Context, params *AuthorizeSecurityGroupEgressInput, optFns ...func(*Options)) (*AuthorizeSecurityGroupEgressOutput, error) {
	if params == nil {
		params = &AuthorizeSecurityGroupEgressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeSecurityGroupEgress", params, optFns, c.addOperationAuthorizeSecurityGroupEgressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeSecurityGroupEgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeSecurityGroupEgressInput struct {

	// The ID of the security group.
	//
	// This member is required.
	GroupId *string

	// Not supported. Use a set of IP permissions to specify the CIDR.
	CidrIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Not supported. Use a set of IP permissions to specify the port.
	FromPort *int32

	// The sets of IP permissions. You can't specify a destination security group and
	// a CIDR IP address range in the same set of permissions.
	IpPermissions []types.IpPermission

	// Not supported. Use a set of IP permissions to specify the protocol name or
	// number.
	IpProtocol *string

	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupName *string

	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupOwnerId *string

	// The tags applied to the security group rule.
	TagSpecifications []types.TagSpecification

	// Not supported. Use a set of IP permissions to specify the port.
	ToPort *int32

	noSmithyDocumentSerde
}

type AuthorizeSecurityGroupEgressOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool

	// Information about the outbound (egress) security group rules that were added.
	SecurityGroupRules []types.SecurityGroupRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeSecurityGroupEgressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAuthorizeSecurityGroupEgress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAuthorizeSecurityGroupEgress{}, middleware.After)
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
	if err = addOpAuthorizeSecurityGroupEgressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeSecurityGroupEgress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeSecurityGroupEgress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AuthorizeSecurityGroupEgress",
	}
}
