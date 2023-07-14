// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified managed prefix list. Adding or removing entries in a
// prefix list creates a new version of the prefix list. Changing the name of the
// prefix list does not affect the version. If you specify a current version number
// that does not match the true current version number, the request fails.
func (c *Client) ModifyManagedPrefixList(ctx context.Context, params *ModifyManagedPrefixListInput, optFns ...func(*Options)) (*ModifyManagedPrefixListOutput, error) {
	if params == nil {
		params = &ModifyManagedPrefixListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyManagedPrefixList", params, optFns, c.addOperationModifyManagedPrefixListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyManagedPrefixListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyManagedPrefixListInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// One or more entries to add to the prefix list.
	AddEntries []types.AddPrefixListEntry

	// The current version of the prefix list.
	CurrentVersion *int64

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of entries for the prefix list. You cannot modify the
	// entries of a prefix list and modify the size of a prefix list at the same time.
	// If any of the resources that reference the prefix list cannot support the new
	// maximum size, the modify operation fails. Check the state message for the IDs of
	// the first ten resources that do not support the new maximum size.
	MaxEntries *int32

	// A name for the prefix list.
	PrefixListName *string

	// One or more entries to remove from the prefix list.
	RemoveEntries []types.RemovePrefixListEntry

	noSmithyDocumentSerde
}

type ModifyManagedPrefixListOutput struct {

	// Information about the prefix list.
	PrefixList *types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyManagedPrefixListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyManagedPrefixList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyManagedPrefixList{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyManagedPrefixListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyManagedPrefixList(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyManagedPrefixList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyManagedPrefixList",
	}
}
