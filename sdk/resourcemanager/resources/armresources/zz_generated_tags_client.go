//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
func NewTagsClient(con *arm.Connection, subscriptionID string) *TagsClient {
	return &TagsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - This operation allows adding a name to the list of predefined tag names for the given subscription. A tag name can have a maximum of
// 512 characters and is case-insensitive. Tag names cannot have the
// following prefixes which are reserved for Azure use: 'microsoft', 'azure', 'windows'.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) CreateOrUpdate(ctx context.Context, tagName string, options *TagsCreateOrUpdateOptions) (TagsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, tagName, options)
	if err != nil {
		return TagsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TagsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagsClient) createOrUpdateCreateRequest(ctx context.Context, tagName string, options *TagsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagsClient) createOrUpdateHandleResponse(resp *http.Response) (TagsCreateOrUpdateResponse, error) {
	result := TagsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagDetails); err != nil {
		return TagsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TagsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// CreateOrUpdateAtScope - This operation allows adding or replacing the entire set of tags on the specified resource or subscription. The specified entity
// can have a maximum of 50 tags.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) CreateOrUpdateAtScope(ctx context.Context, scope string, parameters TagsResource, options *TagsCreateOrUpdateAtScopeOptions) (TagsCreateOrUpdateAtScopeResponse, error) {
	req, err := client.createOrUpdateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsCreateOrUpdateAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsCreateOrUpdateAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsCreateOrUpdateAtScopeResponse{}, client.createOrUpdateAtScopeHandleError(resp)
	}
	return client.createOrUpdateAtScopeHandleResponse(resp)
}

// createOrUpdateAtScopeCreateRequest creates the CreateOrUpdateAtScope request.
func (client *TagsClient) createOrUpdateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsResource, options *TagsCreateOrUpdateAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAtScopeHandleResponse handles the CreateOrUpdateAtScope response.
func (client *TagsClient) createOrUpdateAtScopeHandleResponse(resp *http.Response) (TagsCreateOrUpdateAtScopeResponse, error) {
	result := TagsCreateOrUpdateAtScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsCreateOrUpdateAtScopeResponse{}, err
	}
	return result, nil
}

// createOrUpdateAtScopeHandleError handles the CreateOrUpdateAtScope error response.
func (client *TagsClient) createOrUpdateAtScopeHandleError(resp *http.Response) error {
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

// CreateOrUpdateValue - This operation allows adding a value to the list of predefined values for an existing predefined tag name. A tag value can have
// a maximum of 256 characters.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string, options *TagsCreateOrUpdateValueOptions) (TagsCreateOrUpdateValueResponse, error) {
	req, err := client.createOrUpdateValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return TagsCreateOrUpdateValueResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsCreateOrUpdateValueResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TagsCreateOrUpdateValueResponse{}, client.createOrUpdateValueHandleError(resp)
	}
	return client.createOrUpdateValueHandleResponse(resp)
}

// createOrUpdateValueCreateRequest creates the CreateOrUpdateValue request.
func (client *TagsClient) createOrUpdateValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsCreateOrUpdateValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if tagValue == "" {
		return nil, errors.New("parameter tagValue cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateValueHandleResponse handles the CreateOrUpdateValue response.
func (client *TagsClient) createOrUpdateValueHandleResponse(resp *http.Response) (TagsCreateOrUpdateValueResponse, error) {
	result := TagsCreateOrUpdateValueResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagValue); err != nil {
		return TagsCreateOrUpdateValueResponse{}, err
	}
	return result, nil
}

// createOrUpdateValueHandleError handles the CreateOrUpdateValue error response.
func (client *TagsClient) createOrUpdateValueHandleError(resp *http.Response) error {
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

// Delete - This operation allows deleting a name from the list of predefined tag names for the given subscription. The name being deleted must not be in
// use as a tag name for any resource. All predefined values
// for the given name must have already been deleted.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) Delete(ctx context.Context, tagName string, options *TagsDeleteOptions) (TagsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, tagName, options)
	if err != nil {
		return TagsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return TagsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagsClient) deleteCreateRequest(ctx context.Context, tagName string, options *TagsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TagsClient) deleteHandleError(resp *http.Response) error {
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

// DeleteAtScope - Deletes the entire set of tags on a resource or subscription.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) DeleteAtScope(ctx context.Context, scope string, options *TagsDeleteAtScopeOptions) (TagsDeleteAtScopeResponse, error) {
	req, err := client.deleteAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsDeleteAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsDeleteAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsDeleteAtScopeResponse{}, client.deleteAtScopeHandleError(resp)
	}
	return TagsDeleteAtScopeResponse{RawResponse: resp}, nil
}

// deleteAtScopeCreateRequest creates the DeleteAtScope request.
func (client *TagsClient) deleteAtScopeCreateRequest(ctx context.Context, scope string, options *TagsDeleteAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAtScopeHandleError handles the DeleteAtScope error response.
func (client *TagsClient) deleteAtScopeHandleError(resp *http.Response) error {
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

// DeleteValue - This operation allows deleting a value from the list of predefined values for an existing predefined tag name. The value being deleted
// must not be in use as a tag value for the given tag name for any
// resource.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) DeleteValue(ctx context.Context, tagName string, tagValue string, options *TagsDeleteValueOptions) (TagsDeleteValueResponse, error) {
	req, err := client.deleteValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return TagsDeleteValueResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsDeleteValueResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsDeleteValueResponse{}, client.deleteValueHandleError(resp)
	}
	return TagsDeleteValueResponse{RawResponse: resp}, nil
}

// deleteValueCreateRequest creates the DeleteValue request.
func (client *TagsClient) deleteValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsDeleteValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if tagValue == "" {
		return nil, errors.New("parameter tagValue cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteValueHandleError handles the DeleteValue error response.
func (client *TagsClient) deleteValueHandleError(resp *http.Response) error {
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

// GetAtScope - Gets the entire set of tags on a resource or subscription.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) GetAtScope(ctx context.Context, scope string, options *TagsGetAtScopeOptions) (TagsGetAtScopeResponse, error) {
	req, err := client.getAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsGetAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsGetAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsGetAtScopeResponse{}, client.getAtScopeHandleError(resp)
	}
	return client.getAtScopeHandleResponse(resp)
}

// getAtScopeCreateRequest creates the GetAtScope request.
func (client *TagsClient) getAtScopeCreateRequest(ctx context.Context, scope string, options *TagsGetAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAtScopeHandleResponse handles the GetAtScope response.
func (client *TagsClient) getAtScopeHandleResponse(resp *http.Response) (TagsGetAtScopeResponse, error) {
	result := TagsGetAtScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsGetAtScopeResponse{}, err
	}
	return result, nil
}

// getAtScopeHandleError handles the GetAtScope error response.
func (client *TagsClient) getAtScopeHandleError(resp *http.Response) error {
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

// List - This operation performs a union of predefined tags, resource tags, resource group tags and subscription tags, and returns a summary of usage for
// each tag name and value under the given subscription.
// In case of a large number of tags, this operation may return a previously cached result.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) List(options *TagsListOptions) *TagsListPager {
	return &TagsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TagsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TagsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TagsClient) listCreateRequest(ctx context.Context, options *TagsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagsClient) listHandleResponse(resp *http.Response) (TagsListResponse, error) {
	result := TagsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsListResult); err != nil {
		return TagsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *TagsClient) listHandleError(resp *http.Response) error {
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

// UpdateAtScope - This operation allows replacing, merging or selectively deleting tags on the specified resource or subscription. The specified entity
// can have a maximum of 50 tags at the end of the operation. The
// 'replace' option replaces the entire set of existing tags with a new set. The 'merge' option allows adding tags with new names and updating the values
// of tags with existing names. The 'delete' option
// allows selectively deleting tags based on given names or name/value pairs.
// If the operation fails it returns the *CloudError error type.
func (client *TagsClient) UpdateAtScope(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsUpdateAtScopeOptions) (TagsUpdateAtScopeResponse, error) {
	req, err := client.updateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsUpdateAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsUpdateAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsUpdateAtScopeResponse{}, client.updateAtScopeHandleError(resp)
	}
	return client.updateAtScopeHandleResponse(resp)
}

// updateAtScopeCreateRequest creates the UpdateAtScope request.
func (client *TagsClient) updateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsUpdateAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateAtScopeHandleResponse handles the UpdateAtScope response.
func (client *TagsClient) updateAtScopeHandleResponse(resp *http.Response) (TagsUpdateAtScopeResponse, error) {
	result := TagsUpdateAtScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsUpdateAtScopeResponse{}, err
	}
	return result, nil
}

// updateAtScopeHandleError handles the UpdateAtScope error response.
func (client *TagsClient) updateAtScopeHandleError(resp *http.Response) error {
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
