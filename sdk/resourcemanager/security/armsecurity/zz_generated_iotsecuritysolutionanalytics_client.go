//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// IotSecuritySolutionAnalyticsClient contains the methods for the IotSecuritySolutionAnalytics group.
// Don't use this type directly, use NewIotSecuritySolutionAnalyticsClient() instead.
type IotSecuritySolutionAnalyticsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIotSecuritySolutionAnalyticsClient creates a new instance of IotSecuritySolutionAnalyticsClient with the specified values.
func NewIotSecuritySolutionAnalyticsClient(con *arm.Connection, subscriptionID string) *IotSecuritySolutionAnalyticsClient {
	return &IotSecuritySolutionAnalyticsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Use this method to get IoT Security Analytics metrics.
// If the operation fails it returns the *CloudError error type.
func (client *IotSecuritySolutionAnalyticsClient) Get(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsGetOptions) (IotSecuritySolutionAnalyticsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionAnalyticsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionAnalyticsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionAnalyticsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotSecuritySolutionAnalyticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotSecuritySolutionAnalyticsClient) getHandleResponse(resp *http.Response) (IotSecuritySolutionAnalyticsGetResponse, error) {
	result := IotSecuritySolutionAnalyticsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionAnalyticsModel); err != nil {
		return IotSecuritySolutionAnalyticsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IotSecuritySolutionAnalyticsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Use this method to get IoT security Analytics metrics in an array.
// If the operation fails it returns the *CloudError error type.
func (client *IotSecuritySolutionAnalyticsClient) List(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsListOptions) (IotSecuritySolutionAnalyticsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionAnalyticsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionAnalyticsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionAnalyticsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *IotSecuritySolutionAnalyticsClient) listCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IotSecuritySolutionAnalyticsClient) listHandleResponse(resp *http.Response) (IotSecuritySolutionAnalyticsListResponse, error) {
	result := IotSecuritySolutionAnalyticsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionAnalyticsModelList); err != nil {
		return IotSecuritySolutionAnalyticsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IotSecuritySolutionAnalyticsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
