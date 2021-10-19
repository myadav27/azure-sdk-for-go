//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OriginsClient contains the methods for the Origins group.
// Don't use this type directly, use NewOriginsClient() instead.
type OriginsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOriginsClient creates a new instance of OriginsClient with the specified values.
func NewOriginsClient(con *arm.Connection, subscriptionID string) *OriginsClient {
	return &OriginsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a new origin within the specified endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsBeginCreateOptions) (OriginsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
	if err != nil {
		return OriginsCreatePollerResponse{}, err
	}
	result := OriginsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return OriginsCreatePollerResponse{}, err
	}
	result.Poller = &OriginsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new origin within the specified endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OriginsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, origin)
}

// createHandleError handles the Create error response.
func (client *OriginsClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes an existing origin within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsBeginDeleteOptions) (OriginsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, originName, options)
	if err != nil {
		return OriginsDeletePollerResponse{}, err
	}
	result := OriginsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return OriginsDeletePollerResponse{}, err
	}
	result.Poller = &OriginsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing origin within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OriginsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *OriginsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets an existing origin within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsGetOptions) (OriginsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
	if err != nil {
		return OriginsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OriginsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OriginsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OriginsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OriginsClient) getHandleResponse(resp *http.Response) (OriginsGetResponse, error) {
	result := OriginsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Origin); err != nil {
		return OriginsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OriginsClient) getHandleError(resp *http.Response) error {
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

// ListByEndpoint - Lists all of the existing origins within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) ListByEndpoint(resourceGroupName string, profileName string, endpointName string, options *OriginsListByEndpointOptions) *OriginsListByEndpointPager {
	return &OriginsListByEndpointPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
		},
		advancer: func(ctx context.Context, resp OriginsListByEndpointResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OriginListResult.NextLink)
		},
	}
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *OriginsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *OriginsListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *OriginsClient) listByEndpointHandleResponse(resp *http.Response) (OriginsListByEndpointResponse, error) {
	result := OriginsListByEndpointResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginListResult); err != nil {
		return OriginsListByEndpointResponse{}, err
	}
	return result, nil
}

// listByEndpointHandleError handles the ListByEndpoint error response.
func (client *OriginsClient) listByEndpointHandleError(resp *http.Response) error {
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

// BeginUpdate - Updates an existing origin within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsBeginUpdateOptions) (OriginsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
	if err != nil {
		return OriginsUpdatePollerResponse{}, err
	}
	result := OriginsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return OriginsUpdatePollerResponse{}, err
	}
	result.Poller = &OriginsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing origin within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OriginsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, originUpdateProperties)
}

// updateHandleError handles the Update error response.
func (client *OriginsClient) updateHandleError(resp *http.Response) error {
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