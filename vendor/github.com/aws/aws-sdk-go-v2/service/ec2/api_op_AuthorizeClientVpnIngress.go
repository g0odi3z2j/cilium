// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AuthorizeClientVpnIngressInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Active Directory group to grant access.
	AccessGroupId *string `type:"string"`

	// Indicates whether to grant access to all clients. Use true to grant all clients
	// who successfully establish a VPN connection access to the network.
	AuthorizeAllGroups *bool `type:"boolean"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// A brief description of the authorization rule.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IPv4 address range, in CIDR notation, of the network for which access
	// is being authorized.
	//
	// TargetNetworkCidr is a required field
	TargetNetworkCidr *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AuthorizeClientVpnIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeClientVpnIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AuthorizeClientVpnIngressInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if s.TargetNetworkCidr == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetNetworkCidr"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AuthorizeClientVpnIngressOutput struct {
	_ struct{} `type:"structure"`

	// The current state of the authorization rule.
	Status *VpnAuthorizationRuleStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s AuthorizeClientVpnIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opAuthorizeClientVpnIngress = "AuthorizeClientVpnIngress"

// AuthorizeClientVpnIngressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Adds an ingress authorization rule to a Client VPN endpoint. Ingress authorization
// rules act as firewall rules that grant access to networks. You must configure
// ingress authorization rules to enable clients to access resources in AWS
// or on-premises networks.
//
//    // Example sending a request using AuthorizeClientVpnIngressRequest.
//    req := client.AuthorizeClientVpnIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AuthorizeClientVpnIngress
func (c *Client) AuthorizeClientVpnIngressRequest(input *AuthorizeClientVpnIngressInput) AuthorizeClientVpnIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeClientVpnIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeClientVpnIngressInput{}
	}

	req := c.newRequest(op, input, &AuthorizeClientVpnIngressOutput{})
	return AuthorizeClientVpnIngressRequest{Request: req, Input: input, Copy: c.AuthorizeClientVpnIngressRequest}
}

// AuthorizeClientVpnIngressRequest is the request type for the
// AuthorizeClientVpnIngress API operation.
type AuthorizeClientVpnIngressRequest struct {
	*aws.Request
	Input *AuthorizeClientVpnIngressInput
	Copy  func(*AuthorizeClientVpnIngressInput) AuthorizeClientVpnIngressRequest
}

// Send marshals and sends the AuthorizeClientVpnIngress API request.
func (r AuthorizeClientVpnIngressRequest) Send(ctx context.Context) (*AuthorizeClientVpnIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeClientVpnIngressResponse{
		AuthorizeClientVpnIngressOutput: r.Request.Data.(*AuthorizeClientVpnIngressOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeClientVpnIngressResponse is the response type for the
// AuthorizeClientVpnIngress API operation.
type AuthorizeClientVpnIngressResponse struct {
	*AuthorizeClientVpnIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeClientVpnIngress request.
func (r *AuthorizeClientVpnIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
