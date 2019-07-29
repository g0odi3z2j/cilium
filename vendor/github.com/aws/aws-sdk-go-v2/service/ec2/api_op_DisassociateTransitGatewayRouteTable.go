// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateTransitGatewayRouteTableInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateTransitGatewayRouteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateTransitGatewayRouteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateTransitGatewayRouteTableInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateTransitGatewayRouteTableOutput struct {
	_ struct{} `type:"structure"`

	// Information about the association.
	Association *TransitGatewayAssociation `locationName:"association" type:"structure"`
}

// String returns the string representation
func (s DisassociateTransitGatewayRouteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateTransitGatewayRouteTable = "DisassociateTransitGatewayRouteTable"

// DisassociateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates a resource attachment from a transit gateway route table.
//
//    // Example sending a request using DisassociateTransitGatewayRouteTableRequest.
//    req := client.DisassociateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateTransitGatewayRouteTable
func (c *Client) DisassociateTransitGatewayRouteTableRequest(input *DisassociateTransitGatewayRouteTableInput) DisassociateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opDisassociateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &DisassociateTransitGatewayRouteTableOutput{})
	return DisassociateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.DisassociateTransitGatewayRouteTableRequest}
}

// DisassociateTransitGatewayRouteTableRequest is the request type for the
// DisassociateTransitGatewayRouteTable API operation.
type DisassociateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *DisassociateTransitGatewayRouteTableInput
	Copy  func(*DisassociateTransitGatewayRouteTableInput) DisassociateTransitGatewayRouteTableRequest
}

// Send marshals and sends the DisassociateTransitGatewayRouteTable API request.
func (r DisassociateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*DisassociateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateTransitGatewayRouteTableResponse{
		DisassociateTransitGatewayRouteTableOutput: r.Request.Data.(*DisassociateTransitGatewayRouteTableOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateTransitGatewayRouteTableResponse is the response type for the
// DisassociateTransitGatewayRouteTable API operation.
type DisassociateTransitGatewayRouteTableResponse struct {
	*DisassociateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateTransitGatewayRouteTable request.
func (r *DisassociateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
