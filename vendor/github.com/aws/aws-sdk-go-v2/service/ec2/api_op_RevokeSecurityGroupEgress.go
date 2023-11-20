// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified outbound (egress) rules from the specified security
// group. You can specify rules using either rule IDs or security group rule
// properties. If you use rule properties, the values that you specify (for
// example, ports) must match the existing rule's values exactly. Each rule has a
// protocol, from and to ports, and destination (CIDR range, security group, or
// prefix list). For the TCP and UDP protocols, you must also specify the
// destination port or range of ports. For the ICMP protocol, you must also specify
// the ICMP type and code. If the security group rule has a description, you do not
// need to specify the description to revoke the rule. For a default VPC, if the
// values you specify do not match the existing rule's values, no error is
// returned, and the output describes the security group rules that were not
// revoked. Amazon Web Services recommends that you describe the security group to
// verify that the rules were removed. Rule changes are propagated to instances
// within the security group as quickly as possible. However, a small delay might
// occur.
func (c *Client) RevokeSecurityGroupEgress(ctx context.Context, params *RevokeSecurityGroupEgressInput, optFns ...func(*Options)) (*RevokeSecurityGroupEgressOutput, error) {
	if params == nil {
		params = &RevokeSecurityGroupEgressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeSecurityGroupEgress", params, optFns, c.addOperationRevokeSecurityGroupEgressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeSecurityGroupEgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeSecurityGroupEgressInput struct {

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

	// The IDs of the security group rules.
	SecurityGroupRuleIds []string

	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupName *string

	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupOwnerId *string

	// Not supported. Use a set of IP permissions to specify the port.
	ToPort *int32

	noSmithyDocumentSerde
}

type RevokeSecurityGroupEgressOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool

	// The outbound rules that were unknown to the service. In some cases,
	// unknownIpPermissionSet might be in a different format from the request parameter.
	UnknownIpPermissions []types.IpPermission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeSecurityGroupEgressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpRevokeSecurityGroupEgress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRevokeSecurityGroupEgress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeSecurityGroupEgress"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRevokeSecurityGroupEgressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSecurityGroupEgress(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRevokeSecurityGroupEgress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeSecurityGroupEgress",
	}
}
