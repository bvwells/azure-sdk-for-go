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

// AgentRegistrationInformationClient contains the methods for the AgentRegistrationInformation group.
// Don't use this type directly, use NewAgentRegistrationInformationClient() instead.
type AgentRegistrationInformationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAgentRegistrationInformationClient creates a new instance of AgentRegistrationInformationClient with the specified values.
func NewAgentRegistrationInformationClient(con *arm.Connection, subscriptionID string) *AgentRegistrationInformationClient {
	return &AgentRegistrationInformationClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieve the automation agent registration information.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AgentRegistrationInformationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, options *AgentRegistrationInformationGetOptions) (AgentRegistrationInformationGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return AgentRegistrationInformationGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentRegistrationInformationGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentRegistrationInformationGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentRegistrationInformationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *AgentRegistrationInformationGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/agentRegistrationInformation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
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
func (client *AgentRegistrationInformationClient) getHandleResponse(resp *http.Response) (AgentRegistrationInformationGetResponse, error) {
	result := AgentRegistrationInformationGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentRegistration); err != nil {
		return AgentRegistrationInformationGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AgentRegistrationInformationClient) getHandleError(resp *http.Response) error {
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

// RegenerateKey - Regenerate a primary or secondary agent registration key
// If the operation fails it returns the *ErrorResponse error type.
func (client *AgentRegistrationInformationClient) RegenerateKey(ctx context.Context, resourceGroupName string, automationAccountName string, parameters AgentRegistrationRegenerateKeyParameter, options *AgentRegistrationInformationRegenerateKeyOptions) (AgentRegistrationInformationRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, automationAccountName, parameters, options)
	if err != nil {
		return AgentRegistrationInformationRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentRegistrationInformationRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentRegistrationInformationRegenerateKeyResponse{}, client.regenerateKeyHandleError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *AgentRegistrationInformationClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, parameters AgentRegistrationRegenerateKeyParameter, options *AgentRegistrationInformationRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/agentRegistrationInformation/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *AgentRegistrationInformationClient) regenerateKeyHandleResponse(resp *http.Response) (AgentRegistrationInformationRegenerateKeyResponse, error) {
	result := AgentRegistrationInformationRegenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentRegistration); err != nil {
		return AgentRegistrationInformationRegenerateKeyResponse{}, err
	}
	return result, nil
}

// regenerateKeyHandleError handles the RegenerateKey error response.
func (client *AgentRegistrationInformationClient) regenerateKeyHandleError(resp *http.Response) error {
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
