// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all roles that are assigned to the user for a given AWS account.
func (c *Client) ListAccountRoles(ctx context.Context, params *ListAccountRolesInput, optFns ...func(*Options)) (*ListAccountRolesOutput, error) {
	if params == nil {
		params = &ListAccountRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountRoles", params, optFns, c.addOperationListAccountRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountRolesInput struct {

	// The token issued by the CreateToken API call. For more information, see
	// CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// This member is required.
	AccessToken *string

	// The identifier for the AWS account that is assigned to the user.
	//
	// This member is required.
	AccountId *string

	// The number of items that clients can request per page.
	MaxResults *int32

	// The page token from the previous response output when you request subsequent
	// pages.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountRolesOutput struct {

	// The page token client that is used to retrieve the list of accounts.
	NextToken *string

	// A paginated response with the list of roles and the next token if more results
	// are available.
	RoleList []types.RoleInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccountRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccountRoles{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccountRolesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountRoles(options.Region), middleware.Before); err != nil {
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

// ListAccountRolesAPIClient is a client that implements the ListAccountRoles
// operation.
type ListAccountRolesAPIClient interface {
	ListAccountRoles(context.Context, *ListAccountRolesInput, ...func(*Options)) (*ListAccountRolesOutput, error)
}

var _ ListAccountRolesAPIClient = (*Client)(nil)

// ListAccountRolesPaginatorOptions is the paginator options for ListAccountRoles
type ListAccountRolesPaginatorOptions struct {
	// The number of items that clients can request per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountRolesPaginator is a paginator for ListAccountRoles
type ListAccountRolesPaginator struct {
	options   ListAccountRolesPaginatorOptions
	client    ListAccountRolesAPIClient
	params    *ListAccountRolesInput
	nextToken *string
	firstPage bool
}

// NewListAccountRolesPaginator returns a new ListAccountRolesPaginator
func NewListAccountRolesPaginator(client ListAccountRolesAPIClient, params *ListAccountRolesInput, optFns ...func(*ListAccountRolesPaginatorOptions)) *ListAccountRolesPaginator {
	if params == nil {
		params = &ListAccountRolesInput{}
	}

	options := ListAccountRolesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountRolesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountRolesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountRoles page.
func (p *ListAccountRolesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountRolesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListAccountRoles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAccountRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountRoles",
	}
}
