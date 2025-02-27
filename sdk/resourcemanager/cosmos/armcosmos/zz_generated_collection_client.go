//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// CollectionClient contains the methods for the Collection group.
// Don't use this type directly, use NewCollectionClient() instead.
type CollectionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCollectionClient creates a new instance of CollectionClient with the specified values.
func NewCollectionClient(con *arm.Connection, subscriptionID string) *CollectionClient {
	return &CollectionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListMetricDefinitions - Retrieves metric definitions for the given collection.
// If the operation fails it returns a generic error.
func (client *CollectionClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionListMetricDefinitionsOptions) (CollectionListMetricDefinitionsResponse, error) {
	req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
	if err != nil {
		return CollectionListMetricDefinitionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CollectionListMetricDefinitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CollectionListMetricDefinitionsResponse{}, client.listMetricDefinitionsHandleError(resp)
	}
	return client.listMetricDefinitionsHandleResponse(resp)
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *CollectionClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *CollectionClient) listMetricDefinitionsHandleResponse(resp *http.Response) (CollectionListMetricDefinitionsResponse, error) {
	result := CollectionListMetricDefinitionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionsListResult); err != nil {
		return CollectionListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// listMetricDefinitionsHandleError handles the ListMetricDefinitions error response.
func (client *CollectionClient) listMetricDefinitionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given database account and collection.
// If the operation fails it returns a generic error.
func (client *CollectionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionListMetricsOptions) (CollectionListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter, options)
	if err != nil {
		return CollectionListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CollectionListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CollectionListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionClient) listMetricsHandleResponse(resp *http.Response) (CollectionListMetricsResponse, error) {
	result := CollectionListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return CollectionListMetricsResponse{}, err
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *CollectionClient) listMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListUsages - Retrieves the usages (most recent storage data) for the given collection.
// If the operation fails it returns a generic error.
func (client *CollectionClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionListUsagesOptions) (CollectionListUsagesResponse, error) {
	req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
	if err != nil {
		return CollectionListUsagesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CollectionListUsagesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CollectionListUsagesResponse{}, client.listUsagesHandleError(resp)
	}
	return client.listUsagesHandleResponse(resp)
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *CollectionClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *CollectionClient) listUsagesHandleResponse(resp *http.Response) (CollectionListUsagesResponse, error) {
	result := CollectionListUsagesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesResult); err != nil {
		return CollectionListUsagesResponse{}, err
	}
	return result, nil
}

// listUsagesHandleError handles the ListUsages error response.
func (client *CollectionClient) listUsagesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
