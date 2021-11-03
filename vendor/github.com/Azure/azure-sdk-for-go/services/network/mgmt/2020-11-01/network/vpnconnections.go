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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VpnConnectionsClient is the network Client
type VpnConnectionsClient struct {
	BaseClient
}

// NewVpnConnectionsClient creates an instance of the VpnConnectionsClient client.
func NewVpnConnectionsClient(subscriptionID string) VpnConnectionsClient {
	return NewVpnConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnConnectionsClientWithBaseURI creates an instance of the VpnConnectionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVpnConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VpnConnectionsClient {
	return VpnConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing
// connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the connection.
// vpnConnectionParameters - parameters supplied to create or Update a VPN Connection.
func (client VpnConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection) (result VpnConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VpnConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	vpnConnectionParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithJSON(vpnConnectionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) CreateOrUpdateSender(req *http.Request) (future VpnConnectionsCreateOrUpdateFuture, err error) {
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
func (client VpnConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result VpnConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the connection.
func (client VpnConnectionsClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result VpnConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VpnConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) DeleteSender(req *http.Request) (future VpnConnectionsDeleteFuture, err error) {
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
func (client VpnConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
func (client VpnConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result VpnConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) GetResponder(resp *http.Response) (result VpnConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnGateway retrieves all vpn connections for a particular virtual wan vpn gateway.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
func (client VpnConnectionsClient) ListByVpnGateway(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.lvcr.Response.Response != nil {
				sc = result.lvcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnGatewayNextResults
	req, err := client.ListByVpnGatewayPreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.lvcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", resp, "Failure sending request")
		return
	}

	result.lvcr, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", resp, "Failure responding to request")
		return
	}
	if result.lvcr.hasNextLink() && result.lvcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVpnGatewayPreparer prepares the ListByVpnGateway request.
func (client VpnConnectionsClient) ListByVpnGatewayPreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnGatewaySender sends the ListByVpnGateway request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) ListByVpnGatewaySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnGatewayResponder handles the response to the ListByVpnGateway request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) ListByVpnGatewayResponder(resp *http.Response) (result ListVpnConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnGatewayNextResults retrieves the next set of results, if any.
func (client VpnConnectionsClient) listByVpnGatewayNextResults(ctx context.Context, lastResults ListVpnConnectionsResult) (result ListVpnConnectionsResult, err error) {
	req, err := lastResults.listVpnConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnGatewayComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnConnectionsClient) ListByVpnGatewayComplete(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnGateway(ctx, resourceGroupName, gatewayName)
	return
}

// StartPacketCapture starts packet capture on Vpn connection in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the gateway.
// vpnConnectionName - the name of the vpn connection.
// parameters - vpn Connection packet capture parameters supplied to start packet capture on gateway
// connection.
func (client VpnConnectionsClient) StartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, parameters *VpnConnectionPacketCaptureStartParameters) (result VpnConnectionsStartPacketCaptureFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.StartPacketCapture")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPacketCapturePreparer(ctx, resourceGroupName, gatewayName, vpnConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "StartPacketCapture", nil, "Failure preparing request")
		return
	}

	result, err = client.StartPacketCaptureSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "StartPacketCapture", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPacketCapturePreparer prepares the StartPacketCapture request.
func (client VpnConnectionsClient) StartPacketCapturePreparer(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, parameters *VpnConnectionPacketCaptureStartParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnConnectionName": autorest.Encode("path", vpnConnectionName),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/startpacketcapture", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartPacketCaptureSender sends the StartPacketCapture request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) StartPacketCaptureSender(req *http.Request) (future VpnConnectionsStartPacketCaptureFuture, err error) {
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

// StartPacketCaptureResponder handles the response to the StartPacketCapture request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) StartPacketCaptureResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// StopPacketCapture stops packet capture on Vpn connection in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the gateway.
// vpnConnectionName - the name of the vpn connection.
// parameters - vpn Connection packet capture parameters supplied to stop packet capture on gateway connection.
func (client VpnConnectionsClient) StopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, parameters *VpnConnectionPacketCaptureStopParameters) (result VpnConnectionsStopPacketCaptureFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.StopPacketCapture")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPacketCapturePreparer(ctx, resourceGroupName, gatewayName, vpnConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "StopPacketCapture", nil, "Failure preparing request")
		return
	}

	result, err = client.StopPacketCaptureSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "StopPacketCapture", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPacketCapturePreparer prepares the StopPacketCapture request.
func (client VpnConnectionsClient) StopPacketCapturePreparer(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, parameters *VpnConnectionPacketCaptureStopParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnConnectionName": autorest.Encode("path", vpnConnectionName),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/stoppacketcapture", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopPacketCaptureSender sends the StopPacketCapture request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) StopPacketCaptureSender(req *http.Request) (future VpnConnectionsStopPacketCaptureFuture, err error) {
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

// StopPacketCaptureResponder handles the response to the StopPacketCapture request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) StopPacketCaptureResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
