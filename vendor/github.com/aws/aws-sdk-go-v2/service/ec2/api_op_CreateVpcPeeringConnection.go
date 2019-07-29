// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVpcPeeringConnectionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The AWS account ID of the owner of the accepter VPC.
	//
	// Default: Your AWS account ID
	PeerOwnerId *string `locationName:"peerOwnerId" type:"string"`

	// The Region code for the accepter VPC, if the accepter VPC is located in a
	// Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string `type:"string"`

	// The ID of the VPC with which you are creating the VPC peering connection.
	// You must specify this parameter in the request.
	PeerVpcId *string `locationName:"peerVpcId" type:"string"`

	// The ID of the requester VPC. You must specify this parameter in the request.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s CreateVpcPeeringConnectionInput) String() string {
	return awsutil.Prettify(s)
}

type CreateVpcPeeringConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the VPC peering connection.
	VpcPeeringConnection *VpcPeeringConnection `locationName:"vpcPeeringConnection" type:"structure"`
}

// String returns the string representation
func (s CreateVpcPeeringConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcPeeringConnection = "CreateVpcPeeringConnection"

// CreateVpcPeeringConnectionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Requests a VPC peering connection between two VPCs: a requester VPC that
// you own and an accepter VPC with which to create the connection. The accepter
// VPC can belong to another AWS account and can be in a different Region to
// the requester VPC. The requester VPC and accepter VPC cannot have overlapping
// CIDR blocks.
//
// Limitations and rules apply to a VPC peering connection. For more information,
// see the limitations (https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations)
// section in the VPC Peering Guide.
//
// The owner of the accepter VPC must accept the peering request to activate
// the peering connection. The VPC peering connection request expires after
// 7 days, after which it cannot be accepted or rejected.
//
// If you create a VPC peering connection request between VPCs with overlapping
// CIDR blocks, the VPC peering connection has a status of failed.
//
//    // Example sending a request using CreateVpcPeeringConnectionRequest.
//    req := client.CreateVpcPeeringConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcPeeringConnection
func (c *Client) CreateVpcPeeringConnectionRequest(input *CreateVpcPeeringConnectionInput) CreateVpcPeeringConnectionRequest {
	op := &aws.Operation{
		Name:       opCreateVpcPeeringConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcPeeringConnectionInput{}
	}

	req := c.newRequest(op, input, &CreateVpcPeeringConnectionOutput{})
	return CreateVpcPeeringConnectionRequest{Request: req, Input: input, Copy: c.CreateVpcPeeringConnectionRequest}
}

// CreateVpcPeeringConnectionRequest is the request type for the
// CreateVpcPeeringConnection API operation.
type CreateVpcPeeringConnectionRequest struct {
	*aws.Request
	Input *CreateVpcPeeringConnectionInput
	Copy  func(*CreateVpcPeeringConnectionInput) CreateVpcPeeringConnectionRequest
}

// Send marshals and sends the CreateVpcPeeringConnection API request.
func (r CreateVpcPeeringConnectionRequest) Send(ctx context.Context) (*CreateVpcPeeringConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcPeeringConnectionResponse{
		CreateVpcPeeringConnectionOutput: r.Request.Data.(*CreateVpcPeeringConnectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcPeeringConnectionResponse is the response type for the
// CreateVpcPeeringConnection API operation.
type CreateVpcPeeringConnectionResponse struct {
	*CreateVpcPeeringConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcPeeringConnection request.
func (r *CreateVpcPeeringConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
