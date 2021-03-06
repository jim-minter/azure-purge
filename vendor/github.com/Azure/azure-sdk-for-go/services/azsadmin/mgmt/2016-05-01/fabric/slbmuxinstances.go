package fabric

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
	"net/http"
)

// SlbMuxInstancesClient is the fabric Admin Client
type SlbMuxInstancesClient struct {
	BaseClient
}

// NewSlbMuxInstancesClient creates an instance of the SlbMuxInstancesClient client.
func NewSlbMuxInstancesClient(subscriptionID string) SlbMuxInstancesClient {
	return NewSlbMuxInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSlbMuxInstancesClientWithBaseURI creates an instance of the SlbMuxInstancesClient client.
func NewSlbMuxInstancesClientWithBaseURI(baseURI string, subscriptionID string) SlbMuxInstancesClient {
	return SlbMuxInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a software load balancer multiplexer instance.
//
// location is location of the resource. slbMuxInstance is name of a SLB Mux instance.
func (client SlbMuxInstancesClient) Get(ctx context.Context, location string, slbMuxInstance string) (result SlbMuxInstance, err error) {
	req, err := client.GetPreparer(ctx, location, slbMuxInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SlbMuxInstancesClient) GetPreparer(ctx context.Context, location string, slbMuxInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"slbMuxInstance": autorest.Encode("path", slbMuxInstance),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/slbMuxInstances/{slbMuxInstance}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SlbMuxInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SlbMuxInstancesClient) GetResponder(resp *http.Response) (result SlbMuxInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of all software load balancer instances at a location.
//
// location is location of the resource. filter is oData filter parameter.
func (client SlbMuxInstancesClient) List(ctx context.Context, location string, filter string) (result SlbMuxInstanceListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.smil.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.smil, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SlbMuxInstancesClient) ListPreparer(ctx context.Context, location string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/slbMuxInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SlbMuxInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SlbMuxInstancesClient) ListResponder(resp *http.Response) (result SlbMuxInstanceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SlbMuxInstancesClient) listNextResults(lastResults SlbMuxInstanceList) (result SlbMuxInstanceList, err error) {
	req, err := lastResults.slbMuxInstanceListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.SlbMuxInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SlbMuxInstancesClient) ListComplete(ctx context.Context, location string, filter string) (result SlbMuxInstanceListIterator, err error) {
	result.page, err = client.List(ctx, location, filter)
	return
}
