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

// CommunityGalleryImagesClient is the compute Client
type CommunityGalleryImagesClient struct {
	BaseClient
}

// NewCommunityGalleryImagesClient creates an instance of the CommunityGalleryImagesClient client.
func NewCommunityGalleryImagesClient(subscriptionID string) CommunityGalleryImagesClient {
	return NewCommunityGalleryImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommunityGalleryImagesClientWithBaseURI creates an instance of the CommunityGalleryImagesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewCommunityGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) CommunityGalleryImagesClient {
	return CommunityGalleryImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a community gallery image.
// Parameters:
// location - resource location.
// publicGalleryName - the public name of the community gallery.
// galleryImageName - the name of the community gallery image definition.
func (client CommunityGalleryImagesClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string) (result CommunityGalleryImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunityGalleryImagesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, publicGalleryName, galleryImageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImagesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommunityGalleryImagesClient) GetPreparer(ctx context.Context, location string, publicGalleryName string, galleryImageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"location":          autorest.Encode("path", location),
		"publicGalleryName": autorest.Encode("path", publicGalleryName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommunityGalleryImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommunityGalleryImagesClient) GetResponder(resp *http.Response) (result CommunityGalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
