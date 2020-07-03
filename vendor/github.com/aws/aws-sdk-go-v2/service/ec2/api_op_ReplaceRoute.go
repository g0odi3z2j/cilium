// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type ReplaceRouteInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 CIDR address block used for the destination match. The value that
	// you provide must match the CIDR of an existing route in the table.
	DestinationCidrBlock *string `locationName:"destinationCidrBlock" type:"string"`

	// The IPv6 CIDR address block used for the destination match. The value that
	// you provide must match the CIDR of an existing route in the table.
	DestinationIpv6CidrBlock *string `locationName:"destinationIpv6CidrBlock" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string `locationName:"egressOnlyInternetGatewayId" type:"string"`

	// The ID of an internet gateway or virtual private gateway.
	GatewayId *string `locationName:"gatewayId" type:"string"`

	// The ID of a NAT instance in your VPC.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// The ID of the local gateway.
	LocalGatewayId *string `type:"string"`

	// Specifies whether to reset the local route to its default target (local).
	LocalTarget *bool `type:"boolean"`

	// [IPv4 traffic only] The ID of a NAT gateway.
	NatGatewayId *string `locationName:"natGatewayId" type:"string"`

	// The ID of a network interface.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`

	// The ID of the route table.
	//
	// RouteTableId is a required field
	RouteTableId *string `locationName:"routeTableId" type:"string" required:"true"`

	// The ID of a transit gateway.
	TransitGatewayId *string `type:"string"`

	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string `locationName:"vpcPeeringConnectionId" type:"string"`
}

// String returns the string representation
func (s ReplaceRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceRouteInput"}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReplaceRouteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ReplaceRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opReplaceRoute = "ReplaceRoute"

// ReplaceRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Replaces an existing route within a route table in a VPC. You must provide
// only one of the following: internet gateway, virtual private gateway, NAT
// instance, NAT gateway, VPC peering connection, network interface, egress-only
// internet gateway, or transit gateway.
//
// For more information, see Route Tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using ReplaceRouteRequest.
//    req := client.ReplaceRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceRoute
func (c *Client) ReplaceRouteRequest(input *ReplaceRouteInput) ReplaceRouteRequest {
	op := &aws.Operation{
		Name:       opReplaceRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceRouteInput{}
	}

	req := c.newRequest(op, input, &ReplaceRouteOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ReplaceRouteRequest{Request: req, Input: input, Copy: c.ReplaceRouteRequest}
}

// ReplaceRouteRequest is the request type for the
// ReplaceRoute API operation.
type ReplaceRouteRequest struct {
	*aws.Request
	Input *ReplaceRouteInput
	Copy  func(*ReplaceRouteInput) ReplaceRouteRequest
}

// Send marshals and sends the ReplaceRoute API request.
func (r ReplaceRouteRequest) Send(ctx context.Context) (*ReplaceRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceRouteResponse{
		ReplaceRouteOutput: r.Request.Data.(*ReplaceRouteOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceRouteResponse is the response type for the
// ReplaceRoute API operation.
type ReplaceRouteResponse struct {
	*ReplaceRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceRoute request.
func (r *ReplaceRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
