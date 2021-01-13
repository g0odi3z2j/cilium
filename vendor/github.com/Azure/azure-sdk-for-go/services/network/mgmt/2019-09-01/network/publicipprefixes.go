package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PublicIPPrefixesClient is the network Client
type PublicIPPrefixesClient struct {
	BaseClient
}

// NewPublicIPPrefixesClient creates an instance of the PublicIPPrefixesClient client.
func NewPublicIPPrefixesClient(subscriptionID string) PublicIPPrefixesClient {
	return NewPublicIPPrefixesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPublicIPPrefixesClientWithBaseURI creates an instance of the PublicIPPrefixesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPublicIPPrefixesClientWithBaseURI(baseURI string, subscriptionID string) PublicIPPrefixesClient {
	return PublicIPPrefixesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a static or dynamic public IP prefix.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPPrefixName - the name of the public IP prefix.
// parameters - parameters supplied to the create or update public IP prefix operation.
func (client PublicIPPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix) (result PublicIPPrefixesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, publicIPPrefixName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PublicIPPrefixesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpPrefixName": autorest.Encode("path", publicIPPrefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) CreateOrUpdateSender(req *http.Request) (future PublicIPPrefixesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PublicIPPrefixesClient) (pip PublicIPPrefix, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.PublicIPPrefixesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if pip.Response.Response, err = future.GetResult(sender); err == nil && pip.Response.Response.StatusCode != http.StatusNoContent {
			pip, err = client.CreateOrUpdateResponder(pip.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesCreateOrUpdateFuture", "Result", pip.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) CreateOrUpdateResponder(resp *http.Response) (result PublicIPPrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified public IP prefix.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPPrefixName - the name of the PublicIpPrefix.
func (client PublicIPPrefixesClient) Delete(ctx context.Context, resourceGroupName string, publicIPPrefixName string) (result PublicIPPrefixesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, publicIPPrefixName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublicIPPrefixesClient) DeletePreparer(ctx context.Context, resourceGroupName string, publicIPPrefixName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpPrefixName": autorest.Encode("path", publicIPPrefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) DeleteSender(req *http.Request) (future PublicIPPrefixesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PublicIPPrefixesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.PublicIPPrefixesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified public IP prefix in a specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPPrefixName - the name of the public IP prefix.
// expand - expands referenced resources.
func (client PublicIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, expand string) (result PublicIPPrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, publicIPPrefixName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PublicIPPrefixesClient) GetPreparer(ctx context.Context, resourceGroupName string, publicIPPrefixName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpPrefixName": autorest.Encode("path", publicIPPrefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) GetResponder(resp *http.Response) (result PublicIPPrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all public IP prefixes in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client PublicIPPrefixesClient) List(ctx context.Context, resourceGroupName string) (result PublicIPPrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.List")
		defer func() {
			sc := -1
			if result.piplr.Response.Response != nil {
				sc = result.piplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.piplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "List", resp, "Failure sending request")
		return
	}

	result.piplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.piplr.hasNextLink() && result.piplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PublicIPPrefixesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) ListResponder(resp *http.Response) (result PublicIPPrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PublicIPPrefixesClient) listNextResults(ctx context.Context, lastResults PublicIPPrefixListResult) (result PublicIPPrefixListResult, err error) {
	req, err := lastResults.publicIPPrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPPrefixesClient) ListComplete(ctx context.Context, resourceGroupName string) (result PublicIPPrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all the public IP prefixes in a subscription.
func (client PublicIPPrefixesClient) ListAll(ctx context.Context) (result PublicIPPrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.ListAll")
		defer func() {
			sc := -1
			if result.piplr.Response.Response != nil {
				sc = result.piplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.piplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.piplr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.piplr.hasNextLink() && result.piplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client PublicIPPrefixesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) ListAllResponder(resp *http.Response) (result PublicIPPrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client PublicIPPrefixesClient) listAllNextResults(ctx context.Context, lastResults PublicIPPrefixListResult) (result PublicIPPrefixListResult, err error) {
	req, err := lastResults.publicIPPrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPPrefixesClient) ListAllComplete(ctx context.Context) (result PublicIPPrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// UpdateTags updates public IP prefix tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPPrefixName - the name of the public IP prefix.
// parameters - parameters supplied to update public IP prefix tags.
func (client PublicIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject) (result PublicIPPrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPPrefixesClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, publicIPPrefixName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPPrefixesClient", "UpdateTags", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client PublicIPPrefixesClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpPrefixName": autorest.Encode("path", publicIPPrefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPPrefixesClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client PublicIPPrefixesClient) UpdateTagsResponder(resp *http.Response) (result PublicIPPrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
