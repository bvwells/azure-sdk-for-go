//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

type hsmSecurityDomainClient struct {
	con *connection
}

// BeginDownload - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned managed HSM resource.
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) BeginDownload(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (HSMSecurityDomainDownloadPollerResponse, error) {
	resp, err := client.download(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return HSMSecurityDomainDownloadPollerResponse{}, err
	}
	result := HSMSecurityDomainDownloadPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("hsmSecurityDomainClient.Download", resp, client.con.Pipeline(), client.downloadHandleError)
	if err != nil {
		return HSMSecurityDomainDownloadPollerResponse{}, err
	}
	result.Poller = &HSMSecurityDomainDownloadPoller{
		pt: pt,
	}
	return result, nil
}

// Download - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned managed HSM resource.
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) download(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.downloadHandleError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *hsmSecurityDomainClient) downloadCreateRequest(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, certificateInfoObject)
}

// downloadHandleError handles the Download error response.
func (client *hsmSecurityDomainClient) downloadHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// DownloadPending - Retrieves the Security Domain download operation status
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) DownloadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainDownloadPendingOptions) (HSMSecurityDomainDownloadPendingResponse, error) {
	req, err := client.downloadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainDownloadPendingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HSMSecurityDomainDownloadPendingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainDownloadPendingResponse{}, client.downloadPendingHandleError(resp)
	}
	return client.downloadPendingHandleResponse(resp)
}

// downloadPendingCreateRequest creates the DownloadPending request.
func (client *hsmSecurityDomainClient) downloadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainDownloadPendingOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download/pending"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// downloadPendingHandleResponse handles the DownloadPending response.
func (client *hsmSecurityDomainClient) downloadPendingHandleResponse(resp *http.Response) (HSMSecurityDomainDownloadPendingResponse, error) {
	result := HSMSecurityDomainDownloadPendingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityDomainOperationStatus); err != nil {
		return HSMSecurityDomainDownloadPendingResponse{}, err
	}
	return result, nil
}

// downloadPendingHandleError handles the DownloadPending error response.
func (client *hsmSecurityDomainClient) downloadPendingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// TransferKey - Retrieve Security Domain transfer key
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) TransferKey(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainTransferKeyOptions) (HSMSecurityDomainTransferKeyResponse, error) {
	req, err := client.transferKeyCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainTransferKeyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HSMSecurityDomainTransferKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainTransferKeyResponse{}, client.transferKeyHandleError(resp)
	}
	return client.transferKeyHandleResponse(resp)
}

// transferKeyCreateRequest creates the TransferKey request.
func (client *hsmSecurityDomainClient) transferKeyCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainTransferKeyOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// transferKeyHandleResponse handles the TransferKey response.
func (client *hsmSecurityDomainClient) transferKeyHandleResponse(resp *http.Response) (HSMSecurityDomainTransferKeyResponse, error) {
	result := HSMSecurityDomainTransferKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransferKey); err != nil {
		return HSMSecurityDomainTransferKeyResponse{}, err
	}
	return result, nil
}

// transferKeyHandleError handles the TransferKey error response.
func (client *hsmSecurityDomainClient) transferKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpload - Restore the provided Security Domain.
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) BeginUpload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (HSMSecurityDomainUploadPollerResponse, error) {
	resp, err := client.upload(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return HSMSecurityDomainUploadPollerResponse{}, err
	}
	result := HSMSecurityDomainUploadPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("hsmSecurityDomainClient.Upload", resp, client.con.Pipeline(), client.uploadHandleError)
	if err != nil {
		return HSMSecurityDomainUploadPollerResponse{}, err
	}
	result.Poller = &HSMSecurityDomainUploadPoller{
		pt: pt,
	}
	return result, nil
}

// Upload - Restore the provided Security Domain.
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) upload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (*http.Response, error) {
	req, err := client.uploadCreateRequest(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.uploadHandleError(resp)
	}
	return resp, nil
}

// uploadCreateRequest creates the Upload request.
func (client *hsmSecurityDomainClient) uploadCreateRequest(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, securityDomain)
}

// uploadHandleError handles the Upload error response.
func (client *hsmSecurityDomainClient) uploadHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UploadPending - Get Security Domain upload operation status
// If the operation fails it returns the *KeyVaultError error type.
func (client *hsmSecurityDomainClient) UploadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainUploadPendingOptions) (HSMSecurityDomainUploadPendingResponse, error) {
	req, err := client.uploadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainUploadPendingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HSMSecurityDomainUploadPendingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainUploadPendingResponse{}, client.uploadPendingHandleError(resp)
	}
	return client.uploadPendingHandleResponse(resp)
}

// uploadPendingCreateRequest creates the UploadPending request.
func (client *hsmSecurityDomainClient) uploadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainUploadPendingOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload/pending"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// uploadPendingHandleResponse handles the UploadPending response.
func (client *hsmSecurityDomainClient) uploadPendingHandleResponse(resp *http.Response) (HSMSecurityDomainUploadPendingResponse, error) {
	result := HSMSecurityDomainUploadPendingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityDomainOperationStatus); err != nil {
		return HSMSecurityDomainUploadPendingResponse{}, err
	}
	return result, nil
}

// uploadPendingHandleError handles the UploadPending error response.
func (client *hsmSecurityDomainClient) uploadPendingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
