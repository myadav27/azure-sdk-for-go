//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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
	"strconv"
	"strings"
)

// SystemTopicsClient contains the methods for the SystemTopics group.
// Don't use this type directly, use NewSystemTopicsClient() instead.
type SystemTopicsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSystemTopicsClient creates a new instance of SystemTopicsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSystemTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemTopicsClient, error) {
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
	client := &SystemTopicsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// systemTopicInfo - System Topic information.
// options - SystemTopicsClientBeginCreateOrUpdateOptions contains the optional parameters for the SystemTopicsClient.BeginCreateOrUpdate
// method.
func (client *SystemTopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SystemTopicsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SystemTopicsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SystemTopicsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *SystemTopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, systemTopicInfo)
}

// BeginDelete - Delete existing system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// options - SystemTopicsClientBeginDeleteOptions contains the optional parameters for the SystemTopicsClient.BeginDelete
// method.
func (client *SystemTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsClientBeginDeleteOptions) (*runtime.Poller[SystemTopicsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, systemTopicName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SystemTopicsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SystemTopicsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete existing system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *SystemTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SystemTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// options - SystemTopicsClientGetOptions contains the optional parameters for the SystemTopicsClient.Get method.
func (client *SystemTopicsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsClientGetOptions) (SystemTopicsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return SystemTopicsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SystemTopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SystemTopicsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SystemTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemTopicsClient) getHandleResponse(resp *http.Response) (SystemTopicsClientGetResponse, error) {
	result := SystemTopicsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopic); err != nil {
		return SystemTopicsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the system topics under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// options - SystemTopicsClientListByResourceGroupOptions contains the optional parameters for the SystemTopicsClient.ListByResourceGroup
// method.
func (client *SystemTopicsClient) NewListByResourceGroupPager(resourceGroupName string, options *SystemTopicsClientListByResourceGroupOptions) *runtime.Pager[SystemTopicsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemTopicsClientListByResourceGroupResponse]{
		More: func(page SystemTopicsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemTopicsClientListByResourceGroupResponse) (SystemTopicsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SystemTopicsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SystemTopicsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SystemTopicsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SystemTopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SystemTopicsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SystemTopicsClient) listByResourceGroupHandleResponse(resp *http.Response) (SystemTopicsClientListByResourceGroupResponse, error) {
	result := SystemTopicsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopicsListResult); err != nil {
		return SystemTopicsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the system topics under an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// options - SystemTopicsClientListBySubscriptionOptions contains the optional parameters for the SystemTopicsClient.ListBySubscription
// method.
func (client *SystemTopicsClient) NewListBySubscriptionPager(options *SystemTopicsClientListBySubscriptionOptions) *runtime.Pager[SystemTopicsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemTopicsClientListBySubscriptionResponse]{
		More: func(page SystemTopicsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemTopicsClientListBySubscriptionResponse) (SystemTopicsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SystemTopicsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SystemTopicsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SystemTopicsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SystemTopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SystemTopicsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/systemTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SystemTopicsClient) listBySubscriptionHandleResponse(resp *http.Response) (SystemTopicsClientListBySubscriptionResponse, error) {
	result := SystemTopicsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopicsListResult); err != nil {
		return SystemTopicsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// systemTopicUpdateParameters - SystemTopic update information.
// options - SystemTopicsClientBeginUpdateOptions contains the optional parameters for the SystemTopicsClient.BeginUpdate
// method.
func (client *SystemTopicsClient) BeginUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsClientBeginUpdateOptions) (*runtime.Poller[SystemTopicsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SystemTopicsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SystemTopicsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *SystemTopicsClient) update(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SystemTopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, systemTopicUpdateParameters)
}
