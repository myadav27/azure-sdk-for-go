//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RestorableSQLContainersClient contains the methods for the RestorableSQLContainers group.
// Don't use this type directly, use NewRestorableSQLContainersClient() instead.
type RestorableSQLContainersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableSQLContainersClient creates a new instance of RestorableSQLContainersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableSQLContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableSQLContainersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableSQLContainersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Show the event feed of all mutations done on all the Azure Cosmos DB SQL containers under a specific database.
// This helps in scenario where container was accidentally deleted. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-15-preview
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableSQLContainersClientListOptions contains the optional parameters for the RestorableSQLContainersClient.List
// method.
func (client *RestorableSQLContainersClient) NewListPager(location string, instanceID string, options *RestorableSQLContainersClientListOptions) *runtime.Pager[RestorableSQLContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableSQLContainersClientListResponse]{
		More: func(page RestorableSQLContainersClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableSQLContainersClientListResponse) (RestorableSQLContainersClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableSQLContainersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorableSQLContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableSQLContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorableSQLContainersClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableSQLContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlContainers"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	if options != nil && options.RestorableSQLDatabaseRid != nil {
		reqQP.Set("restorableSqlDatabaseRid", *options.RestorableSQLDatabaseRid)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableSQLContainersClient) listHandleResponse(resp *http.Response) (RestorableSQLContainersClientListResponse, error) {
	result := RestorableSQLContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableSQLContainersListResult); err != nil {
		return RestorableSQLContainersClientListResponse{}, err
	}
	return result, nil
}
