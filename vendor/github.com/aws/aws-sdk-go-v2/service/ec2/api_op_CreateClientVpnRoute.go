// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateClientVpnRouteInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The ID of the Client VPN endpoint to which to add the route.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// A brief description of the route.
	Description *string `type:"string"`

	// The IPv4 address range, in CIDR notation, of the route destination. For example:
	//
	//    * To add a route for Internet access, enter 0.0.0.0/0
	//
	//    * To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range
	//
	//    * To add a route for an on-premises network, enter the AWS Site-to-Site
	//    VPN connection's IPv4 CIDR range
	//
	// Route address ranges cannot overlap with the CIDR range specified for client
	// allocation.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the subnet through which you want to route traffic. The specified
	// subnet must be an existing target network of the Client VPN endpoint.
	//
	// TargetVpcSubnetId is a required field
	TargetVpcSubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateClientVpnRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClientVpnRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClientVpnRouteInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}

	if s.TargetVpcSubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetVpcSubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateClientVpnRouteOutput struct {
	_ struct{} `type:"structure"`

	// The current state of the route.
	Status *VpnRouteStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s CreateClientVpnRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateClientVpnRoute = "CreateClientVpnRoute"

// CreateClientVpnRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Adds a route to a network to a Client VPN endpoint. Each Client VPN endpoint
// has a route table that describes the available destination network routes.
// Each route in the route table specifies the path for traﬃc to speciﬁc
// resources or networks.
//
//    // Example sending a request using CreateClientVpnRouteRequest.
//    req := client.CreateClientVpnRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateClientVpnRoute
func (c *Client) CreateClientVpnRouteRequest(input *CreateClientVpnRouteInput) CreateClientVpnRouteRequest {
	op := &aws.Operation{
		Name:       opCreateClientVpnRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClientVpnRouteInput{}
	}

	req := c.newRequest(op, input, &CreateClientVpnRouteOutput{})
	return CreateClientVpnRouteRequest{Request: req, Input: input, Copy: c.CreateClientVpnRouteRequest}
}

// CreateClientVpnRouteRequest is the request type for the
// CreateClientVpnRoute API operation.
type CreateClientVpnRouteRequest struct {
	*aws.Request
	Input *CreateClientVpnRouteInput
	Copy  func(*CreateClientVpnRouteInput) CreateClientVpnRouteRequest
}

// Send marshals and sends the CreateClientVpnRoute API request.
func (r CreateClientVpnRouteRequest) Send(ctx context.Context) (*CreateClientVpnRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClientVpnRouteResponse{
		CreateClientVpnRouteOutput: r.Request.Data.(*CreateClientVpnRouteOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClientVpnRouteResponse is the response type for the
// CreateClientVpnRoute API operation.
type CreateClientVpnRouteResponse struct {
	*CreateClientVpnRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClientVpnRoute request.
func (r *CreateClientVpnRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
