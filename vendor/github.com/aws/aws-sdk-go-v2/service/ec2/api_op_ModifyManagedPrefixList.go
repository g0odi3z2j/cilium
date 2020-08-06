// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyManagedPrefixListInput struct {
	_ struct{} `type:"structure"`

	// One or more entries to add to the prefix list.
	AddEntries []AddPrefixListEntry `locationName:"AddEntry" type:"list"`

	// The current version of the prefix list.
	CurrentVersion *int64 `type:"long"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the prefix list.
	//
	// PrefixListId is a required field
	PrefixListId *string `type:"string" required:"true"`

	// A name for the prefix list.
	PrefixListName *string `type:"string"`

	// One or more entries to remove from the prefix list.
	RemoveEntries []RemovePrefixListEntry `locationName:"RemoveEntry" type:"list"`
}

// String returns the string representation
func (s ModifyManagedPrefixListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyManagedPrefixListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyManagedPrefixListInput"}

	if s.PrefixListId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrefixListId"))
	}
	if s.AddEntries != nil {
		for i, v := range s.AddEntries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AddEntries", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RemoveEntries != nil {
		for i, v := range s.RemoveEntries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RemoveEntries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyManagedPrefixListOutput struct {
	_ struct{} `type:"structure"`

	// Information about the prefix list.
	PrefixList *ManagedPrefixList `locationName:"prefixList" type:"structure"`
}

// String returns the string representation
func (s ModifyManagedPrefixListOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyManagedPrefixList = "ModifyManagedPrefixList"

// ModifyManagedPrefixListRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified managed prefix list.
//
// Adding or removing entries in a prefix list creates a new version of the
// prefix list. Changing the name of the prefix list does not affect the version.
//
// If you specify a current version number that does not match the true current
// version number, the request fails.
//
//    // Example sending a request using ModifyManagedPrefixListRequest.
//    req := client.ModifyManagedPrefixListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyManagedPrefixList
func (c *Client) ModifyManagedPrefixListRequest(input *ModifyManagedPrefixListInput) ModifyManagedPrefixListRequest {
	op := &aws.Operation{
		Name:       opModifyManagedPrefixList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyManagedPrefixListInput{}
	}

	req := c.newRequest(op, input, &ModifyManagedPrefixListOutput{})

	return ModifyManagedPrefixListRequest{Request: req, Input: input, Copy: c.ModifyManagedPrefixListRequest}
}

// ModifyManagedPrefixListRequest is the request type for the
// ModifyManagedPrefixList API operation.
type ModifyManagedPrefixListRequest struct {
	*aws.Request
	Input *ModifyManagedPrefixListInput
	Copy  func(*ModifyManagedPrefixListInput) ModifyManagedPrefixListRequest
}

// Send marshals and sends the ModifyManagedPrefixList API request.
func (r ModifyManagedPrefixListRequest) Send(ctx context.Context) (*ModifyManagedPrefixListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyManagedPrefixListResponse{
		ModifyManagedPrefixListOutput: r.Request.Data.(*ModifyManagedPrefixListOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyManagedPrefixListResponse is the response type for the
// ModifyManagedPrefixList API operation.
type ModifyManagedPrefixListResponse struct {
	*ModifyManagedPrefixListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyManagedPrefixList request.
func (r *ModifyManagedPrefixListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
