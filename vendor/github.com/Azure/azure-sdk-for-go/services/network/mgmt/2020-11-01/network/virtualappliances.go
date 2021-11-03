package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualAppliancesClient is the network Client
type VirtualAppliancesClient struct {
	BaseClient
}

// NewVirtualAppliancesClient creates an instance of the VirtualAppliancesClient client.
func NewVirtualAppliancesClient(subscriptionID string) VirtualAppliancesClient {
	return NewVirtualAppliancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualAppliancesClientWithBaseURI creates an instance of the VirtualAppliancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVirtualAppliancesClientWithBaseURI(baseURI string, subscriptionID string) VirtualAppliancesClient {
	return VirtualAppliancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified Network Virtual Appliance.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of Network Virtual Appliance.
// parameters - parameters supplied to the create or update Network Virtual Appliance.
func (client VirtualAppliancesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance) (result VirtualAppliancesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VirtualAppliancePropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.VirtualAppliancePropertiesFormat.VirtualApplianceAsn", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.VirtualAppliancePropertiesFormat.VirtualApplianceAsn", Name: validation.InclusiveMaximum, Rule: int64(4294967295), Chain: nil},
						{Target: "parameters.VirtualAppliancePropertiesFormat.VirtualApplianceAsn", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.VirtualAppliancesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkVirtualApplianceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualAppliancesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) CreateOrUpdateSender(req *http.Request) (future VirtualAppliancesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualAppliance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Network Virtual Appliance.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of Network Virtual Appliance.
func (client VirtualAppliancesClient) Delete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string) (result VirtualAppliancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkVirtualApplianceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualAppliancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) DeleteSender(req *http.Request) (future VirtualAppliancesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Network Virtual Appliance.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of Network Virtual Appliance.
// expand - expands referenced resources.
func (client VirtualAppliancesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, expand string) (result VirtualAppliance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkVirtualApplianceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualAppliancesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) GetResponder(resp *http.Response) (result VirtualAppliance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all Network Virtual Appliances in a subscription.
func (client VirtualAppliancesClient) List(ctx context.Context) (result VirtualApplianceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.List")
		defer func() {
			sc := -1
			if result.valr.Response.Response != nil {
				sc = result.valr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.valr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "List", resp, "Failure sending request")
		return
	}

	result.valr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.valr.hasNextLink() && result.valr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualAppliancesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) ListResponder(resp *http.Response) (result VirtualApplianceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualAppliancesClient) listNextResults(ctx context.Context, lastResults VirtualApplianceListResult) (result VirtualApplianceListResult, err error) {
	req, err := lastResults.virtualApplianceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualAppliancesClient) ListComplete(ctx context.Context) (result VirtualApplianceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all Network Virtual Appliances in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client VirtualAppliancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result VirtualApplianceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.valr.Response.Response != nil {
				sc = result.valr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.valr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.valr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.valr.hasNextLink() && result.valr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VirtualAppliancesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) ListByResourceGroupResponder(resp *http.Response) (result VirtualApplianceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VirtualAppliancesClient) listByResourceGroupNextResults(ctx context.Context, lastResults VirtualApplianceListResult) (result VirtualApplianceListResult, err error) {
	req, err := lastResults.virtualApplianceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualAppliancesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result VirtualApplianceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// UpdateTags updates a Network Virtual Appliance.
// Parameters:
// resourceGroupName - the resource group name of Network Virtual Appliance.
// networkVirtualApplianceName - the name of Network Virtual Appliance being updated.
// parameters - parameters supplied to Update Network Virtual Appliance Tags.
func (client VirtualAppliancesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject) (result VirtualAppliance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualAppliancesClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, networkVirtualApplianceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualAppliancesClient", "UpdateTags", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client VirtualAppliancesClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualAppliancesClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client VirtualAppliancesClient) UpdateTagsResponder(resp *http.Response) (result VirtualAppliance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
