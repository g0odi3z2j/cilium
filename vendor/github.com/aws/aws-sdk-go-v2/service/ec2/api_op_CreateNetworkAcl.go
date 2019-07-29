// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateNetworkAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkAclInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNetworkAclOutput struct {
	_ struct{} `type:"structure"`

	// Information about the network ACL.
	NetworkAcl *NetworkAcl `locationName:"networkAcl" type:"structure"`
}

// String returns the string representation
func (s CreateNetworkAclOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkAcl = "CreateNetworkAcl"

// CreateNetworkAclRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a network ACL in a VPC. Network ACLs provide an optional layer of
// security (in addition to security groups) for the instances in your VPC.
//
// For more information, see Network ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNetworkAclRequest.
//    req := client.CreateNetworkAclRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkAcl
func (c *Client) CreateNetworkAclRequest(input *CreateNetworkAclInput) CreateNetworkAclRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkAcl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkAclInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkAclOutput{})
	return CreateNetworkAclRequest{Request: req, Input: input, Copy: c.CreateNetworkAclRequest}
}

// CreateNetworkAclRequest is the request type for the
// CreateNetworkAcl API operation.
type CreateNetworkAclRequest struct {
	*aws.Request
	Input *CreateNetworkAclInput
	Copy  func(*CreateNetworkAclInput) CreateNetworkAclRequest
}

// Send marshals and sends the CreateNetworkAcl API request.
func (r CreateNetworkAclRequest) Send(ctx context.Context) (*CreateNetworkAclResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkAclResponse{
		CreateNetworkAclOutput: r.Request.Data.(*CreateNetworkAclOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkAclResponse is the response type for the
// CreateNetworkAcl API operation.
type CreateNetworkAclResponse struct {
	*CreateNetworkAclOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkAcl request.
func (r *CreateNetworkAclResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
