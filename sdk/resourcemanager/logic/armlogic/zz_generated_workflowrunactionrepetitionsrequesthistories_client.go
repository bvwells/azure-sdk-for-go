//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// WorkflowRunActionRepetitionsRequestHistoriesClient contains the methods for the WorkflowRunActionRepetitionsRequestHistories group.
// Don't use this type directly, use NewWorkflowRunActionRepetitionsRequestHistoriesClient() instead.
type WorkflowRunActionRepetitionsRequestHistoriesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkflowRunActionRepetitionsRequestHistoriesClient creates a new instance of WorkflowRunActionRepetitionsRequestHistoriesClient with the specified values.
func NewWorkflowRunActionRepetitionsRequestHistoriesClient(con *arm.Connection, subscriptionID string) *WorkflowRunActionRepetitionsRequestHistoriesClient {
	return &WorkflowRunActionRepetitionsRequestHistoriesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a workflow run repetition request history.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string, options *WorkflowRunActionRepetitionsRequestHistoriesGetOptions) (WorkflowRunActionRepetitionsRequestHistoriesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName, requestHistoryName, options)
	if err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionRepetitionsRequestHistoriesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string, options *WorkflowRunActionRepetitionsRequestHistoriesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories/{requestHistoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	if requestHistoryName == "" {
		return nil, errors.New("parameter requestHistoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{requestHistoryName}", url.PathEscape(requestHistoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) getHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsRequestHistoriesGetResponse, error) {
	result := WorkflowRunActionRepetitionsRequestHistoriesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestHistory); err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List a workflow run repetition request history.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) List(resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsRequestHistoriesListOptions) *WorkflowRunActionRepetitionsRequestHistoriesListPager {
	return &WorkflowRunActionRepetitionsRequestHistoriesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowRunActionRepetitionsRequestHistoriesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RequestHistoryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsRequestHistoriesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) listHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsRequestHistoriesListResponse, error) {
	result := WorkflowRunActionRepetitionsRequestHistoriesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestHistoryListResult); err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
