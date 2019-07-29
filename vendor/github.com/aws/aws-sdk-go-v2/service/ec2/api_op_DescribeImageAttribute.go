// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeImageAttribute.
type DescribeImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The AMI attribute.
	//
	// Note: Depending on your account privileges, the blockDeviceMapping attribute
	// may return a Client.AuthFailure error. If this happens, use DescribeImages
	// to get information about the block device mapping for the AMI.
	//
	// Attribute is a required field
	Attribute ImageAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the AMI.
	//
	// ImageId is a required field
	ImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeImageAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes an image attribute.
type DescribeImageAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The block device mapping entries.
	BlockDeviceMappings []BlockDeviceMapping `locationName:"blockDeviceMapping" locationNameList:"item" type:"list"`

	// A description for the AMI.
	Description *AttributeValue `locationName:"description" type:"structure"`

	// The ID of the AMI.
	ImageId *string `locationName:"imageId" type:"string"`

	// The kernel ID.
	KernelId *AttributeValue `locationName:"kernel" type:"structure"`

	// The launch permissions.
	LaunchPermissions []LaunchPermission `locationName:"launchPermission" locationNameList:"item" type:"list"`

	// The product codes.
	ProductCodes []ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`

	// The RAM disk ID.
	RamdiskId *AttributeValue `locationName:"ramdisk" type:"structure"`

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *AttributeValue `locationName:"sriovNetSupport" type:"structure"`
}

// String returns the string representation
func (s DescribeImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeImageAttribute = "DescribeImageAttribute"

// DescribeImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified AMI. You can specify only
// one attribute at a time.
//
//    // Example sending a request using DescribeImageAttributeRequest.
//    req := client.DescribeImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeImageAttribute
func (c *Client) DescribeImageAttributeRequest(input *DescribeImageAttributeInput) DescribeImageAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeImageAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeImageAttributeOutput{})
	return DescribeImageAttributeRequest{Request: req, Input: input, Copy: c.DescribeImageAttributeRequest}
}

// DescribeImageAttributeRequest is the request type for the
// DescribeImageAttribute API operation.
type DescribeImageAttributeRequest struct {
	*aws.Request
	Input *DescribeImageAttributeInput
	Copy  func(*DescribeImageAttributeInput) DescribeImageAttributeRequest
}

// Send marshals and sends the DescribeImageAttribute API request.
func (r DescribeImageAttributeRequest) Send(ctx context.Context) (*DescribeImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageAttributeResponse{
		DescribeImageAttributeOutput: r.Request.Data.(*DescribeImageAttributeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeImageAttributeResponse is the response type for the
// DescribeImageAttribute API operation.
type DescribeImageAttributeResponse struct {
	*DescribeImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImageAttribute request.
func (r *DescribeImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
