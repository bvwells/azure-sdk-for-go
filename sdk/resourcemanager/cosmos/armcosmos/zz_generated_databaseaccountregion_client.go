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

// DatabaseAccountRegionClient contains the methods for the DatabaseAccountRegion group.
// Don't use this type directly, use NewDatabaseAccountRegionClient() instead.
type DatabaseAccountRegionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabaseAccountRegionClient creates a new instance of DatabaseAccountRegionClient with the specified values.
func NewDatabaseAccountRegionClient(con *arm.Connection, subscriptionID string) *DatabaseAccountRegionClient {
	return &DatabaseAccountRegionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given database account and region.
// If the operation fails it returns a generic error.
func (client *DatabaseAccountRegionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, filter string, options *DatabaseAccountRegionListMetricsOptions) (DatabaseAccountRegionListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, region, filter, options)
	if err != nil {
		return DatabaseAccountRegionListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseAccountRegionListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseAccountRegionListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *DatabaseAccountRegionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, region string, filter string, options *DatabaseAccountRegionListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics"
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
	if region == "" {
		return nil, errors.New("parameter region cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{region}", url.PathEscape(region))
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
func (client *DatabaseAccountRegionClient) listMetricsHandleResponse(resp *http.Response) (DatabaseAccountRegionListMetricsResponse, error) {
	result := DatabaseAccountRegionListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return DatabaseAccountRegionListMetricsResponse{}, err
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *DatabaseAccountRegionClient) listMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
