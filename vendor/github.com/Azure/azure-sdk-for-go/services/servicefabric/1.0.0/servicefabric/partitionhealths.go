package servicefabric

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

// PartitionHealthsClient is the client for the PartitionHealths methods of the Servicefabric service.
type PartitionHealthsClient struct {
	BaseClient
}

// NewPartitionHealthsClient creates an instance of the PartitionHealthsClient client.
func NewPartitionHealthsClient(timeout *int32) PartitionHealthsClient {
	return NewPartitionHealthsClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewPartitionHealthsClientWithBaseURI creates an instance of the PartitionHealthsClient client.
func NewPartitionHealthsClientWithBaseURI(baseURI string, timeout *int32) PartitionHealthsClient {
	return PartitionHealthsClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get partition healths
//
// partitionID is the id of the partition eventsHealthStateFilter is the filter of the events health state
// replicasHealthStateFilter is the filter of the replicas health state
func (client PartitionHealthsClient) Get(ctx context.Context, partitionID string, eventsHealthStateFilter string, replicasHealthStateFilter string) (result PartitionHealth, err error) {
	req, err := client.GetPreparer(ctx, partitionID, eventsHealthStateFilter, replicasHealthStateFilter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PartitionHealthsClient) GetPreparer(ctx context.Context, partitionID string, eventsHealthStateFilter string, replicasHealthStateFilter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partitionId": autorest.Encode("path", partitionID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(eventsHealthStateFilter) > 0 {
		queryParameters["EventsHealthStateFilter"] = autorest.Encode("query", eventsHealthStateFilter)
	}
	if len(replicasHealthStateFilter) > 0 {
		queryParameters["ReplicasHealthStateFilter"] = autorest.Encode("query", replicasHealthStateFilter)
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Partitions/{partitionId}/$/GetHealth", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PartitionHealthsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PartitionHealthsClient) GetResponder(resp *http.Response) (result PartitionHealth, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Send send partition health
//
// partitionID is the id of the partition partitionHealthReport is the report of the partition health
func (client PartitionHealthsClient) Send(ctx context.Context, partitionID string, partitionHealthReport PartitionHealthReport) (result String, err error) {
	req, err := client.SendPreparer(ctx, partitionID, partitionHealthReport)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Send", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Send", resp, "Failure sending request")
		return
	}

	result, err = client.SendResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.PartitionHealthsClient", "Send", resp, "Failure responding to request")
	}

	return
}

// SendPreparer prepares the Send request.
func (client PartitionHealthsClient) SendPreparer(ctx context.Context, partitionID string, partitionHealthReport PartitionHealthReport) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partitionId": autorest.Encode("path", partitionID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Partitions/{partitionId}/$/ReportHealth", pathParameters),
		autorest.WithJSON(partitionHealthReport),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendSender sends the Send request. The method will close the
// http.Response Body if it receives an error.
func (client PartitionHealthsClient) SendSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SendResponder handles the response to the Send request. The method always
// closes the http.Response Body.
func (client PartitionHealthsClient) SendResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
