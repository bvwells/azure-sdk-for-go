//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// MaintenanceWindowsClient contains the methods for the MaintenanceWindows group.
// Don't use this type directly, use NewMaintenanceWindowsClient() instead.
type MaintenanceWindowsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMaintenanceWindowsClient creates a new instance of MaintenanceWindowsClient with the specified values.
func NewMaintenanceWindowsClient(con *arm.Connection, subscriptionID string) *MaintenanceWindowsClient {
	return &MaintenanceWindowsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Sets maintenance windows settings for a database.
// If the operation fails it returns a generic error.
func (client *MaintenanceWindowsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowName string, parameters MaintenanceWindows, options *MaintenanceWindowsCreateOrUpdateOptions) (MaintenanceWindowsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, maintenanceWindowName, parameters, options)
	if err != nil {
		return MaintenanceWindowsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceWindowsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceWindowsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return MaintenanceWindowsCreateOrUpdateResponse{RawResponse: resp}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MaintenanceWindowsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowName string, parameters MaintenanceWindows, options *MaintenanceWindowsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/maintenanceWindows/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("maintenanceWindowName", maintenanceWindowName)
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *MaintenanceWindowsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets maintenance windows settings for a database.
// If the operation fails it returns a generic error.
func (client *MaintenanceWindowsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowName string, options *MaintenanceWindowsGetOptions) (MaintenanceWindowsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, maintenanceWindowName, options)
	if err != nil {
		return MaintenanceWindowsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceWindowsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceWindowsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MaintenanceWindowsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowName string, options *MaintenanceWindowsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/maintenanceWindows/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("maintenanceWindowName", maintenanceWindowName)
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MaintenanceWindowsClient) getHandleResponse(resp *http.Response) (MaintenanceWindowsGetResponse, error) {
	result := MaintenanceWindowsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceWindows); err != nil {
		return MaintenanceWindowsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MaintenanceWindowsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
