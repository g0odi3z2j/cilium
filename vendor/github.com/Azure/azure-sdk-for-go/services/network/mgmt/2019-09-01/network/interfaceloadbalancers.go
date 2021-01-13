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

// InterfaceLoadBalancersClient is the network Client
type InterfaceLoadBalancersClient struct {
	BaseClient
}

// NewInterfaceLoadBalancersClient creates an instance of the InterfaceLoadBalancersClient client.
func NewInterfaceLoadBalancersClient(subscriptionID string) InterfaceLoadBalancersClient {
	return NewInterfaceLoadBalancersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInterfaceLoadBalancersClientWithBaseURI creates an instance of the InterfaceLoadBalancersClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewInterfaceLoadBalancersClientWithBaseURI(baseURI string, subscriptionID string) InterfaceLoadBalancersClient {
	return InterfaceLoadBalancersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list all load balancers in a network interface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
func (client InterfaceLoadBalancersClient) List(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceLoadBalancerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceLoadBalancersClient.List")
		defer func() {
			sc := -1
			if result.ilblr.Response.Response != nil {
				sc = result.ilblr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilblr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "List", resp, "Failure sending request")
		return
	}

	result.ilblr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ilblr.hasNextLink() && result.ilblr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client InterfaceLoadBalancersClient) ListPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceLoadBalancersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfaceLoadBalancersClient) ListResponder(resp *http.Response) (result InterfaceLoadBalancerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InterfaceLoadBalancersClient) listNextResults(ctx context.Context, lastResults InterfaceLoadBalancerListResult) (result InterfaceLoadBalancerListResult, err error) {
	req, err := lastResults.interfaceLoadBalancerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceLoadBalancersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfaceLoadBalancersClient) ListComplete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceLoadBalancerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceLoadBalancersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkInterfaceName)
	return
}
