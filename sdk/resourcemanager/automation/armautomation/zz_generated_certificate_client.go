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

// CertificateClient contains the methods for the Certificate group.
// Don't use this type directly, use NewCertificateClient() instead.
type CertificateClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCertificateClient creates a new instance of CertificateClient with the specified values.
func NewCertificateClient(con *arm.Connection, subscriptionID string) *CertificateClient {
	return &CertificateClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create a certificate.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateCreateOrUpdateOptions) (CertificateCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, parameters, options)
	if err != nil {
		return CertificateCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CertificateCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CertificateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CertificateClient) createOrUpdateHandleResponse(resp *http.Response) (CertificateCreateOrUpdateResponse, error) {
	result := CertificateCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CertificateClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete the certificate.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CertificateClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateDeleteOptions) (CertificateDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, options)
	if err != nil {
		return CertificateDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateDeleteResponse{}, client.deleteHandleError(resp)
	}
	return CertificateDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CertificateClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieve the certificate identified by certificate name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CertificateClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateGetOptions) (CertificateGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, options)
	if err != nil {
		return CertificateGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CertificateClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
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
func (client *CertificateClient) getHandleResponse(resp *http.Response) (CertificateGetResponse, error) {
	result := CertificateGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CertificateClient) getHandleError(resp *http.Response) error {
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

// ListByAutomationAccount - Retrieve a list of certificates.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CertificateClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *CertificateListByAutomationAccountOptions) *CertificateListByAutomationAccountPager {
	return &CertificateListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp CertificateListByAutomationAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CertificateListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *CertificateClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *CertificateListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates"
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

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *CertificateClient) listByAutomationAccountHandleResponse(resp *http.Response) (CertificateListByAutomationAccountResponse, error) {
	result := CertificateListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListResult); err != nil {
		return CertificateListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *CertificateClient) listByAutomationAccountHandleError(resp *http.Response) error {
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

// Update - Update a certificate.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CertificateClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateUpdateParameters, options *CertificateUpdateOptions) (CertificateUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, parameters, options)
	if err != nil {
		return CertificateUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CertificateClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateUpdateParameters, options *CertificateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CertificateClient) updateHandleResponse(resp *http.Response) (CertificateUpdateResponse, error) {
	result := CertificateUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *CertificateClient) updateHandleError(resp *http.Response) error {
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
