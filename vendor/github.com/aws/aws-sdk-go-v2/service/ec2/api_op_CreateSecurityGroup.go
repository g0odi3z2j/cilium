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

// Creates a security group. A security group acts as a virtual firewall for your
// instance to control inbound and outbound traffic. For more information, see
// Amazon EC2 security groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
// in the Amazon Elastic Compute Cloud User Guide and Security groups for your VPC (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
// in the Amazon Virtual Private Cloud User Guide. When you create a security
// group, you specify a friendly name of your choice. You can have a security group
// for use in EC2-Classic with the same name as a security group for use in a VPC.
// However, you can't have two security groups for use in EC2-Classic with the same
// name or two security groups for use in a VPC with the same name. You have a
// default security group for use in EC2-Classic and a default security group for
// use in your VPC. If you don't specify a security group when you launch an
// instance, the instance is launched into the appropriate default security group.
// A default security group includes a default rule that grants instances
// unrestricted network access to each other. You can add or remove rules from your
// security groups using AuthorizeSecurityGroupIngress ,
// AuthorizeSecurityGroupEgress , RevokeSecurityGroupIngress , and
// RevokeSecurityGroupEgress . For more information about VPC security group
// limits, see Amazon VPC Limits (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html)
// . We are retiring EC2-Classic. We recommend that you migrate from EC2-Classic to
// a VPC. For more information, see Migrate from EC2-Classic to a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateSecurityGroup(ctx context.Context, params *CreateSecurityGroupInput, optFns ...func(*Options)) (*CreateSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityGroup", params, optFns, c.addOperationCreateSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityGroupInput struct {

	// A description for the security group. Constraints: Up to 255 characters in
	// length Constraints for EC2-Classic: ASCII characters Constraints for EC2-VPC:
	// a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	//
	// This member is required.
	Description *string

	// The name of the security group. Constraints: Up to 255 characters in length.
	// Cannot start with sg- . Constraints for EC2-Classic: ASCII characters
	// Constraints for EC2-VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	//
	// This member is required.
	GroupName *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to assign to the security group.
	TagSpecifications []types.TagSpecification

	// [EC2-VPC] The ID of the VPC. Required for EC2-VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type CreateSecurityGroupOutput struct {

	// The ID of the security group.
	GroupId *string

	// The tags assigned to the security group.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSecurityGroup{}, middleware.After)
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
	if err = addOpCreateSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSecurityGroup",
	}
}
