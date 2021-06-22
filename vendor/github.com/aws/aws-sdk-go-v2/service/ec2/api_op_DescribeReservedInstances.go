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

// Describes one or more of the Reserved Instances that you purchased. For more
// information about Reserved Instances, see Reserved Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
// in the Amazon EC2 User Guide.
func (c *Client) DescribeReservedInstances(ctx context.Context, params *DescribeReservedInstancesInput, optFns ...func(*Options)) (*DescribeReservedInstancesOutput, error) {
	if params == nil {
		params = &DescribeReservedInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstances", params, optFns, addOperationDescribeReservedInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeReservedInstances.
type DescribeReservedInstancesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * availability-zone - The Availability Zone where the
	// Reserved Instance can be used.
	//
	// * duration - The duration of the Reserved
	// Instance (one year or three years), in seconds (31536000 | 94608000).
	//
	// * end -
	// The time when the Reserved Instance expires (for example,
	// 2015-08-07T11:54:42.000Z).
	//
	// * fixed-price - The purchase price of the Reserved
	// Instance (for example, 9800.0).
	//
	// * instance-type - The instance type that is
	// covered by the reservation.
	//
	// * scope - The scope of the Reserved Instance
	// (Region or Availability Zone).
	//
	// * product-description - The Reserved Instance
	// product platform description. Instances that include (Amazon VPC) in the product
	// platform description will only be displayed to EC2-Classic account holders and
	// are for use with Amazon VPC (Linux/UNIX | Linux/UNIX (Amazon VPC) | SUSE Linux |
	// SUSE Linux (Amazon VPC) | Red Hat Enterprise Linux | Red Hat Enterprise Linux
	// (Amazon VPC) | Red Hat Enterprise Linux with HA (Amazon VPC) | Windows | Windows
	// (Amazon VPC) | Windows with SQL Server Standard | Windows with SQL Server
	// Standard (Amazon VPC) | Windows with SQL Server Web | Windows with SQL Server
	// Web (Amazon VPC) | Windows with SQL Server Enterprise | Windows with SQL Server
	// Enterprise (Amazon VPC)).
	//
	// * reserved-instances-id - The ID of the Reserved
	// Instance.
	//
	// * start - The time at which the Reserved Instance purchase request
	// was placed (for example, 2014-08-07T11:54:42.000Z).
	//
	// * state - The state of the
	// Reserved Instance (payment-pending | active | payment-failed | retired).
	//
	// * tag:
	// - The key/value combination of a tag assigned to the resource. Use the tag key
	// in the filter name and the tag value as the filter value. For example, to find
	// all resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key - The
	// key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	//
	// * usage-price
	// - The usage price of the Reserved Instance, per hour (for example, 0.84).
	Filters []types.Filter

	// Describes whether the Reserved Instance is Standard or Convertible.
	OfferingClass types.OfferingClassType

	// The Reserved Instance offering type. If you are using tools that predate the
	// 2011-11-01 API version, you only have access to the Medium Utilization Reserved
	// Instance offering type.
	OfferingType types.OfferingTypeValues

	// One or more Reserved Instance IDs. Default: Describes all your Reserved
	// Instances, or only those otherwise specified.
	ReservedInstancesIds []string
}

// Contains the output for DescribeReservedInstances.
type DescribeReservedInstancesOutput struct {

	// A list of Reserved Instances.
	ReservedInstances []types.ReservedInstances

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReservedInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeReservedInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReservedInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReservedInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeReservedInstances",
	}
}
