// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ImportInstanceInput struct {
	_ struct{} `type:"structure"`

	// A description for the instance being imported.
	Description *string `locationName:"description" type:"string"`

	// The disk image.
	DiskImages []DiskImage `locationName:"diskImage" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The launch specification.
	LaunchSpecification *ImportInstanceLaunchSpecification `locationName:"launchSpecification" type:"structure"`

	// The instance operating system.
	//
	// Platform is a required field
	Platform PlatformValues `locationName:"platform" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ImportInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportInstanceInput"}
	if len(s.Platform) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Platform"))
	}
	if s.DiskImages != nil {
		for i, v := range s.DiskImages {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DiskImages", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ImportInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the conversion task.
	ConversionTask *ConversionTask `locationName:"conversionTask" type:"structure"`
}

// String returns the string representation
func (s ImportInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportInstance = "ImportInstance"

// ImportInstanceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an import instance task using metadata from the specified disk image.
// ImportInstance only supports single-volume VMs. To import multi-volume VMs,
// use ImportImage. For more information, see Importing a Virtual Machine Using
// the Amazon EC2 CLI (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/ec2-cli-vmimport-export.html).
//
// For information about the import manifest referenced by this API action,
// see VM Import Manifest (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
//
//    // Example sending a request using ImportInstanceRequest.
//    req := client.ImportInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportInstance
func (c *Client) ImportInstanceRequest(input *ImportInstanceInput) ImportInstanceRequest {
	op := &aws.Operation{
		Name:       opImportInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportInstanceInput{}
	}

	req := c.newRequest(op, input, &ImportInstanceOutput{})
	return ImportInstanceRequest{Request: req, Input: input, Copy: c.ImportInstanceRequest}
}

// ImportInstanceRequest is the request type for the
// ImportInstance API operation.
type ImportInstanceRequest struct {
	*aws.Request
	Input *ImportInstanceInput
	Copy  func(*ImportInstanceInput) ImportInstanceRequest
}

// Send marshals and sends the ImportInstance API request.
func (r ImportInstanceRequest) Send(ctx context.Context) (*ImportInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportInstanceResponse{
		ImportInstanceOutput: r.Request.Data.(*ImportInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportInstanceResponse is the response type for the
// ImportInstance API operation.
type ImportInstanceResponse struct {
	*ImportInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportInstance request.
func (r *ImportInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
