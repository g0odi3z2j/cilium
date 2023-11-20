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

// Lists all AWS accounts assigned to the user. These AWS accounts are assigned by
// the administrator of the account. For more information, see Assign User Access (https://docs.aws.amazon.com/singlesignon/latest/userguide/useraccess.html#assignusers)
// in the IAM Identity Center User Guide. This operation returns a paginated
// response.
func (c *Client) ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) {
	if params == nil {
		params = &ListAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccounts", params, optFns, c.addOperationListAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountsInput struct {

	// The token issued by the CreateToken API call. For more information, see
	// CreateToken (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the IAM Identity Center OIDC API Reference Guide.
	//
	// This member is required.
	AccessToken *string

	// This is the number of items clients can request per page.
	MaxResults *int32

	// (Optional) When requesting subsequent pages, this is the page token from the
	// previous response output.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountsOutput struct {

	// A paginated response with the list of account information and the next token if
	// more results are available.
	AccountList []types.AccountInfo

	// The page token client that is used to retrieve the list of accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccounts"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccounts(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListAccountsAPIClient is a client that implements the ListAccounts operation.
type ListAccountsAPIClient interface {
	ListAccounts(context.Context, *ListAccountsInput, ...func(*Options)) (*ListAccountsOutput, error)
}

var _ ListAccountsAPIClient = (*Client)(nil)

// ListAccountsPaginatorOptions is the paginator options for ListAccounts
type ListAccountsPaginatorOptions struct {
	// This is the number of items clients can request per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountsPaginator is a paginator for ListAccounts
type ListAccountsPaginator struct {
	options   ListAccountsPaginatorOptions
	client    ListAccountsAPIClient
	params    *ListAccountsInput
	nextToken *string
	firstPage bool
}

// NewListAccountsPaginator returns a new ListAccountsPaginator
func NewListAccountsPaginator(client ListAccountsAPIClient, params *ListAccountsInput, optFns ...func(*ListAccountsPaginatorOptions)) *ListAccountsPaginator {
	if params == nil {
		params = &ListAccountsInput{}
	}

	options := ListAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccounts page.
func (p *ListAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountsOutput, error) {
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

	result, err := p.client.ListAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccounts",
	}
}
