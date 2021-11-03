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

// SharedGalleryImagesClient is the compute Client
type SharedGalleryImagesClient struct {
	BaseClient
}

// NewSharedGalleryImagesClient creates an instance of the SharedGalleryImagesClient client.
func NewSharedGalleryImagesClient(subscriptionID string) SharedGalleryImagesClient {
	return NewSharedGalleryImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharedGalleryImagesClientWithBaseURI creates an instance of the SharedGalleryImagesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSharedGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) SharedGalleryImagesClient {
	return SharedGalleryImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a shared gallery image by subscription id or tenant id.
// Parameters:
// location - resource location.
// galleryUniqueName - the unique name of the Shared Gallery.
// galleryImageName - the name of the Shared Gallery Image Definition from which the Image Versions are to be
// listed.
func (client SharedGalleryImagesClient) Get(ctx context.Context, location string, galleryUniqueName string, galleryImageName string) (result SharedGalleryImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedGalleryImagesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, galleryUniqueName, galleryImageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SharedGalleryImagesClient) GetPreparer(ctx context.Context, location string, galleryUniqueName string, galleryImageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"galleryUniqueName": autorest.Encode("path", galleryUniqueName),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SharedGalleryImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SharedGalleryImagesClient) GetResponder(resp *http.Response) (result SharedGalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list shared gallery images by subscription id or tenant id.
// Parameters:
// location - resource location.
// galleryUniqueName - the unique name of the Shared Gallery.
// sharedTo - the query parameter to decide what shared galleries to fetch when doing listing operations.
func (client SharedGalleryImagesClient) List(ctx context.Context, location string, galleryUniqueName string, sharedTo SharedToValues) (result SharedGalleryImageListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedGalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.sgil.Response.Response != nil {
				sc = result.sgil.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, galleryUniqueName, sharedTo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sgil.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "List", resp, "Failure sending request")
		return
	}

	result.sgil, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sgil.hasNextLink() && result.sgil.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SharedGalleryImagesClient) ListPreparer(ctx context.Context, location string, galleryUniqueName string, sharedTo SharedToValues) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryUniqueName": autorest.Encode("path", galleryUniqueName),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(sharedTo)) > 0 {
		queryParameters["sharedTo"] = autorest.Encode("query", sharedTo)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SharedGalleryImagesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SharedGalleryImagesClient) ListResponder(resp *http.Response) (result SharedGalleryImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SharedGalleryImagesClient) listNextResults(ctx context.Context, lastResults SharedGalleryImageList) (result SharedGalleryImageList, err error) {
	req, err := lastResults.sharedGalleryImageListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SharedGalleryImagesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharedGalleryImagesClient) ListComplete(ctx context.Context, location string, galleryUniqueName string, sharedTo SharedToValues) (result SharedGalleryImageListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedGalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location, galleryUniqueName, sharedTo)
	return
}
