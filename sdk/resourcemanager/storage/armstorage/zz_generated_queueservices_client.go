//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// QueueServicesClient contains the methods for the QueueServices group.
// Don't use this type directly, use NewQueueServicesClient() instead.
type QueueServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewQueueServicesClient creates a new instance of QueueServicesClient with the specified values.
func NewQueueServicesClient(con *arm.Connection, subscriptionID string) *QueueServicesClient {
	return &QueueServicesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetServiceProperties - Gets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *QueueServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesGetServicePropertiesOptions) (QueueServicesGetServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return QueueServicesGetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueServicesGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueueServicesGetServicePropertiesResponse{}, client.getServicePropertiesHandleError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *QueueServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *QueueServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (QueueServicesGetServicePropertiesResponse, error) {
	result := QueueServicesGetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueueServiceProperties); err != nil {
		return QueueServicesGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *QueueServicesClient) getServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List all queue services for the storage account
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *QueueServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesListOptions) (QueueServicesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return QueueServicesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueServicesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueueServicesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *QueueServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QueueServicesClient) listHandleResponse(resp *http.Response) (QueueServicesListResponse, error) {
	result := QueueServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListQueueServices); err != nil {
		return QueueServicesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *QueueServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// SetServiceProperties - Sets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *QueueServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties, options *QueueServicesSetServicePropertiesOptions) (QueueServicesSetServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return QueueServicesSetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueServicesSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueueServicesSetServicePropertiesResponse{}, client.setServicePropertiesHandleError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *QueueServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties, options *QueueServicesSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *QueueServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (QueueServicesSetServicePropertiesResponse, error) {
	result := QueueServicesSetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueueServiceProperties); err != nil {
		return QueueServicesSetServicePropertiesResponse{}, err
	}
	return result, nil
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *QueueServicesClient) setServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
