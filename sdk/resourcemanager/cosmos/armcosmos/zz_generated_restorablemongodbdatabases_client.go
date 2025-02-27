//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// RestorableMongodbDatabasesClient contains the methods for the RestorableMongodbDatabases group.
// Don't use this type directly, use NewRestorableMongodbDatabasesClient() instead.
type RestorableMongodbDatabasesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRestorableMongodbDatabasesClient creates a new instance of RestorableMongodbDatabasesClient with the specified values.
func NewRestorableMongodbDatabasesClient(con *arm.Connection, subscriptionID string) *RestorableMongodbDatabasesClient {
	return &RestorableMongodbDatabasesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Show the event feed of all mutations done on all the Azure Cosmos DB MongoDB databases under the restorable account. This helps in scenario where
// database was accidentally deleted to get the deletion
// time. This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns the *CloudError error type.
func (client *RestorableMongodbDatabasesClient) List(ctx context.Context, location string, instanceID string, options *RestorableMongodbDatabasesListOptions) (RestorableMongodbDatabasesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableMongodbDatabasesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableMongodbDatabasesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableMongodbDatabasesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbDatabasesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbDatabasesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbDatabasesClient) listHandleResponse(resp *http.Response) (RestorableMongodbDatabasesListResponse, error) {
	result := RestorableMongodbDatabasesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbDatabasesListResult); err != nil {
		return RestorableMongodbDatabasesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableMongodbDatabasesClient) listHandleError(resp *http.Response) error {
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
