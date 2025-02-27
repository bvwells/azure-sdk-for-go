//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PipelineRunsClient contains the methods for the PipelineRuns group.
// Don't use this type directly, use NewPipelineRunsClient() instead.
type PipelineRunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPipelineRunsClient creates a new instance of PipelineRunsClient with the specified values.
func NewPipelineRunsClient(con *arm.Connection, subscriptionID string) *PipelineRunsClient {
	return &PipelineRunsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (PipelineRunsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
	if err != nil {
		return PipelineRunsCreatePollerResponse{}, err
	}
	result := PipelineRunsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PipelineRunsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return PipelineRunsCreatePollerResponse{}, err
	}
	result.Poller = &PipelineRunsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) create(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PipelineRunsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, pipelineRunCreateParameters)
}

// createHandleError handles the Create error response.
func (client *PipelineRunsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a pipeline run from a container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (PipelineRunsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return PipelineRunsDeletePollerResponse{}, err
	}
	result := PipelineRunsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PipelineRunsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PipelineRunsDeletePollerResponse{}, err
	}
	result.Poller = &PipelineRunsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a pipeline run from a container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelineRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PipelineRunsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the detailed information for a given pipeline run.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsGetOptions) (PipelineRunsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return PipelineRunsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelineRunsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelineRunsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelineRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineRunsClient) getHandleResponse(resp *http.Response) (PipelineRunsGetResponse, error) {
	result := PipelineRunsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return PipelineRunsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PipelineRunsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists all the pipeline runs for the specified container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) List(resourceGroupName string, registryName string, options *PipelineRunsListOptions) *PipelineRunsListPager {
	return &PipelineRunsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp PipelineRunsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PipelineRunListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PipelineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *PipelineRunsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineRunsClient) listHandleResponse(resp *http.Response) (PipelineRunsListResponse, error) {
	result := PipelineRunsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunListResult); err != nil {
		return PipelineRunsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PipelineRunsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
