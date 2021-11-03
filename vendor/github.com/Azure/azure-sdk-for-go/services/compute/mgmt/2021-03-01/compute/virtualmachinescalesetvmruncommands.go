package compute

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

// VirtualMachineScaleSetVMRunCommandsClient is the compute Client
type VirtualMachineScaleSetVMRunCommandsClient struct {
	BaseClient
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates an instance of the VirtualMachineScaleSetVMRunCommandsClient
// client.
func NewVirtualMachineScaleSetVMRunCommandsClient(subscriptionID string) VirtualMachineScaleSetVMRunCommandsClient {
	return NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI creates an instance of the
// VirtualMachineScaleSetVMRunCommandsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMRunCommandsClient {
	return VirtualMachineScaleSetVMRunCommandsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the VMSS VM run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// runCommandName - the name of the virtual machine run command.
// runCommand - parameters supplied to the Create Virtual Machine RunCommand operation.
func (client VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand) (result VirtualMachineScaleSetVMRunCommandsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, runCommandName, runCommand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}", pathParameters),
		autorest.WithJSON(runCommand),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdateSender(req *http.Request) (future VirtualMachineScaleSetVMRunCommandsCreateOrUpdateFuture, err error) {
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
func (client VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the VMSS VM run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// runCommandName - the name of the virtual machine run command.
func (client VirtualMachineScaleSetVMRunCommandsClient) Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string) (result VirtualMachineScaleSetVMRunCommandsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, runCommandName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineScaleSetVMRunCommandsClient) DeletePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMRunCommandsClient) DeleteSender(req *http.Request) (future VirtualMachineScaleSetVMRunCommandsDeleteFuture, err error) {
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
func (client VirtualMachineScaleSetVMRunCommandsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the VMSS VM run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// runCommandName - the name of the virtual machine run command.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineScaleSetVMRunCommandsClient) Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, expand string) (result VirtualMachineRunCommand, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, runCommandName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineScaleSetVMRunCommandsClient) GetPreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMRunCommandsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMRunCommandsClient) GetResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to get all run commands of an instance in Virtual Machine Scaleset.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineScaleSetVMRunCommandsClient) List(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (result VirtualMachineRunCommandsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.List")
		defer func() {
			sc := -1
			if result.vmrclr.Response.Response != nil {
				sc = result.vmrclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vmrclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "List", resp, "Failure sending request")
		return
	}

	result.vmrclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vmrclr.hasNextLink() && result.vmrclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineScaleSetVMRunCommandsClient) ListPreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMRunCommandsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMRunCommandsClient) ListResponder(resp *http.Response) (result VirtualMachineRunCommandsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualMachineScaleSetVMRunCommandsClient) listNextResults(ctx context.Context, lastResults VirtualMachineRunCommandsListResult) (result VirtualMachineRunCommandsListResult, err error) {
	req, err := lastResults.virtualMachineRunCommandsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineScaleSetVMRunCommandsClient) ListComplete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (result VirtualMachineRunCommandsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, VMScaleSetName, instanceID, expand)
	return
}

// Update the operation to update the VMSS VM run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// runCommandName - the name of the virtual machine run command.
// runCommand - parameters supplied to the Update Virtual Machine RunCommand operation.
func (client VirtualMachineScaleSetVMRunCommandsClient) Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate) (result VirtualMachineScaleSetVMRunCommandsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMRunCommandsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, runCommandName, runCommand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMRunCommandsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachineScaleSetVMRunCommandsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}", pathParameters),
		autorest.WithJSON(runCommand),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMRunCommandsClient) UpdateSender(req *http.Request) (future VirtualMachineScaleSetVMRunCommandsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMRunCommandsClient) UpdateResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
