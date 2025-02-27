//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// LinkedWorkspaceClient contains the methods for the LinkedWorkspace group.
// Don't use this type directly, use NewLinkedWorkspaceClient() instead.
type LinkedWorkspaceClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLinkedWorkspaceClient creates a new instance of LinkedWorkspaceClient with the specified values.
func NewLinkedWorkspaceClient(con *arm.Connection, subscriptionID string) *LinkedWorkspaceClient {
	return &LinkedWorkspaceClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieve the linked workspace for the account id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedWorkspaceClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, options *LinkedWorkspaceGetOptions) (LinkedWorkspaceGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return LinkedWorkspaceGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedWorkspaceGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedWorkspaceGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedWorkspaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *LinkedWorkspaceGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/linkedWorkspace"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedWorkspaceClient) getHandleResponse(resp *http.Response) (LinkedWorkspaceGetResponse, error) {
	result := LinkedWorkspaceGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedWorkspace); err != nil {
		return LinkedWorkspaceGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LinkedWorkspaceClient) getHandleError(resp *http.Response) error {
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
