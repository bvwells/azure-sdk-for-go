//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// VirtualNetworkGatewayNatRulesClient contains the methods for the VirtualNetworkGatewayNatRules group.
// Don't use this type directly, use NewVirtualNetworkGatewayNatRulesClient() instead.
type VirtualNetworkGatewayNatRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualNetworkGatewayNatRulesClient creates a new instance of VirtualNetworkGatewayNatRulesClient with the specified values.
func NewVirtualNetworkGatewayNatRulesClient(con *arm.Connection, subscriptionID string) *VirtualNetworkGatewayNatRulesClient {
	return &VirtualNetworkGatewayNatRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a nat rule to a scalable virtual network gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesBeginCreateOrUpdateOptions) (VirtualNetworkGatewayNatRulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, natRuleParameters, options)
	if err != nil {
		return VirtualNetworkGatewayNatRulesCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualNetworkGatewayNatRulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayNatRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkGatewayNatRulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayNatRulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a nat rule to a scalable virtual network gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, natRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkGatewayNatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, natRuleParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkGatewayNatRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a nat rule.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesBeginDeleteOptions) (VirtualNetworkGatewayNatRulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
	if err != nil {
		return VirtualNetworkGatewayNatRulesDeletePollerResponse{}, err
	}
	result := VirtualNetworkGatewayNatRulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayNatRulesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualNetworkGatewayNatRulesDeletePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayNatRulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a nat rule.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
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
func (client *VirtualNetworkGatewayNatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkGatewayNatRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the details of a nat rule.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesGetOptions) (VirtualNetworkGatewayNatRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
	if err != nil {
		return VirtualNetworkGatewayNatRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkGatewayNatRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkGatewayNatRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkGatewayNatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkGatewayNatRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkGatewayNatRulesGetResponse, error) {
	result := VirtualNetworkGatewayNatRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkGatewayNatRule); err != nil {
		return VirtualNetworkGatewayNatRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkGatewayNatRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByVirtualNetworkGateway - Retrieves all nat rules for a particular virtual network gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGateway(resourceGroupName string, virtualNetworkGatewayName string, options *VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayOptions) *VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayPager {
	return &VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVirtualNetworkGatewayCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualNetworkGatewayNatRulesResult.NextLink)
		},
	}
}

// listByVirtualNetworkGatewayCreateRequest creates the ListByVirtualNetworkGateway request.
func (client *VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, options *VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVirtualNetworkGatewayHandleResponse handles the ListByVirtualNetworkGateway response.
func (client *VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayHandleResponse(resp *http.Response) (VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayResponse, error) {
	result := VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualNetworkGatewayNatRulesResult); err != nil {
		return VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayResponse{}, err
	}
	return result, nil
}

// listByVirtualNetworkGatewayHandleError handles the ListByVirtualNetworkGateway error response.
func (client *VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
