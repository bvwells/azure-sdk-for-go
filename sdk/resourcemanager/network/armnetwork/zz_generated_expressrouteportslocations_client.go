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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExpressRoutePortsLocationsClient contains the methods for the ExpressRoutePortsLocations group.
// Don't use this type directly, use NewExpressRoutePortsLocationsClient() instead.
type ExpressRoutePortsLocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExpressRoutePortsLocationsClient creates a new instance of ExpressRoutePortsLocationsClient with the specified values.
func NewExpressRoutePortsLocationsClient(con *arm.Connection, subscriptionID string) *ExpressRoutePortsLocationsClient {
	return &ExpressRoutePortsLocationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said peering location.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsLocationsClient) Get(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsGetOptions) (ExpressRoutePortsLocationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, options)
	if err != nil {
		return ExpressRoutePortsLocationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsLocationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsLocationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsLocationsClient) getCreateRequest(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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
func (client *ExpressRoutePortsLocationsClient) getHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsGetResponse, error) {
	result := ExpressRoutePortsLocationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocation); err != nil {
		return ExpressRoutePortsLocationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRoutePortsLocationsClient) getHandleError(resp *http.Response) error {
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

// List - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location. Available bandwidths can only be obtained
// when retrieving a specific peering location.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsLocationsClient) List(options *ExpressRoutePortsLocationsListOptions) *ExpressRoutePortsLocationsListPager {
	return &ExpressRoutePortsLocationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsLocationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortsLocationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsLocationsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsLocationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsLocationsClient) listHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsListResponse, error) {
	result := ExpressRoutePortsLocationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocationListResult); err != nil {
		return ExpressRoutePortsLocationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRoutePortsLocationsClient) listHandleError(resp *http.Response) error {
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
