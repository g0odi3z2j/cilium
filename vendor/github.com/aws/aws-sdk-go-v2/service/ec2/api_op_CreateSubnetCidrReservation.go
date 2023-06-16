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

// Creates a subnet CIDR reservation. For information about subnet CIDR
// reservations, see Subnet CIDR reservations (https://docs.aws.amazon.com/vpc/latest/userguide/subnet-cidr-reservation.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateSubnetCidrReservation(ctx context.Context, params *CreateSubnetCidrReservationInput, optFns ...func(*Options)) (*CreateSubnetCidrReservationOutput, error) {
	if params == nil {
		params = &CreateSubnetCidrReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubnetCidrReservation", params, optFns, c.addOperationCreateSubnetCidrReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubnetCidrReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubnetCidrReservationInput struct {

	// The IPv4 or IPV6 CIDR range to reserve.
	//
	// This member is required.
	Cidr *string

	// The type of reservation. The following are valid values:
	//   - prefix : The Amazon EC2 Prefix Delegation feature assigns the IP addresses
	//   to network interfaces that are associated with an instance. For information
	//   about Prefix Delegation, see Prefix Delegation for Amazon EC2 network
	//   interfaces (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-delegation.html)
	//   in the Amazon Elastic Compute Cloud User Guide.
	//   - explicit : You manually assign the IP addresses to resources that reside in
	//   your subnet.
	//
	// This member is required.
	ReservationType types.SubnetCidrReservationType

	// The ID of the subnet.
	//
	// This member is required.
	SubnetId *string

	// The description to assign to the subnet CIDR reservation.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to assign to the subnet CIDR reservation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateSubnetCidrReservationOutput struct {

	// Information about the created subnet CIDR reservation.
	SubnetCidrReservation *types.SubnetCidrReservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubnetCidrReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSubnetCidrReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSubnetCidrReservation{}, middleware.After)
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
	if err = addOpCreateSubnetCidrReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubnetCidrReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubnetCidrReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSubnetCidrReservation",
	}
}
