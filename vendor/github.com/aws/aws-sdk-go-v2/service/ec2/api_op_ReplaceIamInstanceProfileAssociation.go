// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReplaceIamInstanceProfileAssociationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the existing IAM instance profile association.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`

	// The IAM instance profile.
	//
	// IamInstanceProfile is a required field
	IamInstanceProfile *IamInstanceProfileSpecification `type:"structure" required:"true"`
}

// String returns the string representation
func (s ReplaceIamInstanceProfileAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceIamInstanceProfileAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceIamInstanceProfileAssociationInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.IamInstanceProfile == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamInstanceProfile"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReplaceIamInstanceProfileAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *IamInstanceProfileAssociation `locationName:"iamInstanceProfileAssociation" type:"structure"`
}

// String returns the string representation
func (s ReplaceIamInstanceProfileAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opReplaceIamInstanceProfileAssociation = "ReplaceIamInstanceProfileAssociation"

// ReplaceIamInstanceProfileAssociationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Replaces an IAM instance profile for the specified running instance. You
// can use this action to change the IAM instance profile that's associated
// with an instance without having to disassociate the existing IAM instance
// profile first.
//
// Use DescribeIamInstanceProfileAssociations to get the association ID.
//
//    // Example sending a request using ReplaceIamInstanceProfileAssociationRequest.
//    req := client.ReplaceIamInstanceProfileAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceIamInstanceProfileAssociation
func (c *Client) ReplaceIamInstanceProfileAssociationRequest(input *ReplaceIamInstanceProfileAssociationInput) ReplaceIamInstanceProfileAssociationRequest {
	op := &aws.Operation{
		Name:       opReplaceIamInstanceProfileAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceIamInstanceProfileAssociationInput{}
	}

	req := c.newRequest(op, input, &ReplaceIamInstanceProfileAssociationOutput{})

	return ReplaceIamInstanceProfileAssociationRequest{Request: req, Input: input, Copy: c.ReplaceIamInstanceProfileAssociationRequest}
}

// ReplaceIamInstanceProfileAssociationRequest is the request type for the
// ReplaceIamInstanceProfileAssociation API operation.
type ReplaceIamInstanceProfileAssociationRequest struct {
	*aws.Request
	Input *ReplaceIamInstanceProfileAssociationInput
	Copy  func(*ReplaceIamInstanceProfileAssociationInput) ReplaceIamInstanceProfileAssociationRequest
}

// Send marshals and sends the ReplaceIamInstanceProfileAssociation API request.
func (r ReplaceIamInstanceProfileAssociationRequest) Send(ctx context.Context) (*ReplaceIamInstanceProfileAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceIamInstanceProfileAssociationResponse{
		ReplaceIamInstanceProfileAssociationOutput: r.Request.Data.(*ReplaceIamInstanceProfileAssociationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceIamInstanceProfileAssociationResponse is the response type for the
// ReplaceIamInstanceProfileAssociation API operation.
type ReplaceIamInstanceProfileAssociationResponse struct {
	*ReplaceIamInstanceProfileAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceIamInstanceProfileAssociation request.
func (r *ReplaceIamInstanceProfileAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
