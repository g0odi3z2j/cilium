// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpcEndpointServiceConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether requests to create an endpoint to your service must be
	// accepted.
	AcceptanceRequired *bool `type:"boolean"`

	// The Amazon Resource Names (ARNs) of Network Load Balancers to add to your
	// service configuration.
	AddNetworkLoadBalancerArns []string `locationName:"AddNetworkLoadBalancerArn" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The private DNS name to assign to the endpoint service.
	PrivateDnsName *string `type:"string"`

	// The Amazon Resource Names (ARNs) of Network Load Balancers to remove from
	// your service configuration.
	RemoveNetworkLoadBalancerArns []string `locationName:"RemoveNetworkLoadBalancerArn" locationNameList:"item" type:"list"`

	// Removes the private DNS name of the endpoint service.
	RemovePrivateDnsName *bool `type:"boolean"`

	// The ID of the service.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcEndpointServiceConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointServiceConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcEndpointServiceConfigurationInput"}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcEndpointServiceConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointServiceConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpcEndpointServiceConfiguration = "ModifyVpcEndpointServiceConfiguration"

// ModifyVpcEndpointServiceConfigurationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the attributes of your VPC endpoint service configuration. You can
// change the Network Load Balancers for your service, and you can specify whether
// acceptance is required for requests to connect to your endpoint service through
// an interface VPC endpoint.
//
// If you set or modify the private DNS name, you must prove that you own the
// private DNS domain name. For more information, see VPC Endpoint Service Private
// DNS Name Verification (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using ModifyVpcEndpointServiceConfigurationRequest.
//    req := client.ModifyVpcEndpointServiceConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointServiceConfiguration
func (c *Client) ModifyVpcEndpointServiceConfigurationRequest(input *ModifyVpcEndpointServiceConfigurationInput) ModifyVpcEndpointServiceConfigurationRequest {
	op := &aws.Operation{
		Name:       opModifyVpcEndpointServiceConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointServiceConfigurationInput{}
	}

	req := c.newRequest(op, input, &ModifyVpcEndpointServiceConfigurationOutput{})

	return ModifyVpcEndpointServiceConfigurationRequest{Request: req, Input: input, Copy: c.ModifyVpcEndpointServiceConfigurationRequest}
}

// ModifyVpcEndpointServiceConfigurationRequest is the request type for the
// ModifyVpcEndpointServiceConfiguration API operation.
type ModifyVpcEndpointServiceConfigurationRequest struct {
	*aws.Request
	Input *ModifyVpcEndpointServiceConfigurationInput
	Copy  func(*ModifyVpcEndpointServiceConfigurationInput) ModifyVpcEndpointServiceConfigurationRequest
}

// Send marshals and sends the ModifyVpcEndpointServiceConfiguration API request.
func (r ModifyVpcEndpointServiceConfigurationRequest) Send(ctx context.Context) (*ModifyVpcEndpointServiceConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcEndpointServiceConfigurationResponse{
		ModifyVpcEndpointServiceConfigurationOutput: r.Request.Data.(*ModifyVpcEndpointServiceConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcEndpointServiceConfigurationResponse is the response type for the
// ModifyVpcEndpointServiceConfiguration API operation.
type ModifyVpcEndpointServiceConfigurationResponse struct {
	*ModifyVpcEndpointServiceConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcEndpointServiceConfiguration request.
func (r *ModifyVpcEndpointServiceConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
