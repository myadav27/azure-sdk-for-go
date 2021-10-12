//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleAssignmentsClient contains the methods for the RoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
func NewRoleAssignmentsClient(con *arm.Connection, subscriptionID string) *RoleAssignmentsClient {
	return &RoleAssignmentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Create or update a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Create(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (RoleAssignmentsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RoleAssignmentsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentsClient) createHandleResponse(resp *http.Response) (RoleAssignmentsCreateResponse, error) {
	result := RoleAssignmentsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *RoleAssignmentsClient) createHandleError(resp *http.Response) error {
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

// CreateByID - Create or update a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) CreateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateByIDOptions) (RoleAssignmentsCreateByIDResponse, error) {
	req, err := client.createByIDCreateRequest(ctx, roleAssignmentID, parameters, options)
	if err != nil {
		return RoleAssignmentsCreateByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsCreateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RoleAssignmentsCreateByIDResponse{}, client.createByIDHandleError(resp)
	}
	return client.createByIDHandleResponse(resp)
}

// createByIDCreateRequest creates the CreateByID request.
func (client *RoleAssignmentsClient) createByIDCreateRequest(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createByIDHandleResponse handles the CreateByID response.
func (client *RoleAssignmentsClient) createByIDHandleResponse(resp *http.Response) (RoleAssignmentsCreateByIDResponse, error) {
	result := RoleAssignmentsCreateByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsCreateByIDResponse{}, err
	}
	return result, nil
}

// createByIDHandleError handles the CreateByID error response.
func (client *RoleAssignmentsClient) createByIDHandleError(resp *http.Response) error {
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

// Delete - Delete a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Delete(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (RoleAssignmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RoleAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RoleAssignmentsClient) deleteHandleResponse(resp *http.Response) (RoleAssignmentsDeleteResponse, error) {
	result := RoleAssignmentsDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoleAssignmentsClient) deleteHandleError(resp *http.Response) error {
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

// DeleteByID - Delete a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) DeleteByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsDeleteByIDOptions) (RoleAssignmentsDeleteByIDResponse, error) {
	req, err := client.deleteByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentsDeleteByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentsDeleteByIDResponse{}, client.deleteByIDHandleError(resp)
	}
	return client.deleteByIDHandleResponse(resp)
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *RoleAssignmentsClient) deleteByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteByIDHandleResponse handles the DeleteByID response.
func (client *RoleAssignmentsClient) deleteByIDHandleResponse(resp *http.Response) (RoleAssignmentsDeleteByIDResponse, error) {
	result := RoleAssignmentsDeleteByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsDeleteByIDResponse{}, err
	}
	return result, nil
}

// deleteByIDHandleError handles the DeleteByID error response.
func (client *RoleAssignmentsClient) deleteByIDHandleError(resp *http.Response) error {
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

// Get - Get a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Get(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (RoleAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentsClient) getHandleResponse(resp *http.Response) (RoleAssignmentsGetResponse, error) {
	result := RoleAssignmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleAssignmentsClient) getHandleError(resp *http.Response) error {
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

// GetByID - Get a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) GetByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsGetByIDOptions) (RoleAssignmentsGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentsGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *RoleAssignmentsClient) getByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsGetByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *RoleAssignmentsClient) getByIDHandleResponse(resp *http.Response) (RoleAssignmentsGetByIDResponse, error) {
	result := RoleAssignmentsGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsGetByIDResponse{}, err
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *RoleAssignmentsClient) getByIDHandleError(resp *http.Response) error {
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

// ListForResource - List all role assignments that apply to a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, resourceType string, resourceName string, options *RoleAssignmentsListForResourceOptions) *RoleAssignmentsListForResourcePager {
	return &RoleAssignmentsListForResourcePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, resourceType, resourceName, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsListForResourceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *RoleAssignmentsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, resourceType string, resourceName string, options *RoleAssignmentsListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", resourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *RoleAssignmentsClient) listForResourceHandleResponse(resp *http.Response) (RoleAssignmentsListForResourceResponse, error) {
	result := RoleAssignmentsListForResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsListForResourceResponse{}, err
	}
	return result, nil
}

// listForResourceHandleError handles the ListForResource error response.
func (client *RoleAssignmentsClient) listForResourceHandleError(resp *http.Response) error {
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

// ListForResourceGroup - List all role assignments that apply to a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForResourceGroup(resourceGroupName string, options *RoleAssignmentsListForResourceGroupOptions) *RoleAssignmentsListForResourceGroupPager {
	return &RoleAssignmentsListForResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsListForResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *RoleAssignmentsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RoleAssignmentsListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *RoleAssignmentsClient) listForResourceGroupHandleResponse(resp *http.Response) (RoleAssignmentsListForResourceGroupResponse, error) {
	result := RoleAssignmentsListForResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsListForResourceGroupResponse{}, err
	}
	return result, nil
}

// listForResourceGroupHandleError handles the ListForResourceGroup error response.
func (client *RoleAssignmentsClient) listForResourceGroupHandleError(resp *http.Response) error {
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

// ListForScope - List all role assignments that apply to a scope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForScope(scope string, options *RoleAssignmentsListForScopeOptions) *RoleAssignmentsListForScopePager {
	return &RoleAssignmentsListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentsListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentsListForScopeResponse, error) {
	result := RoleAssignmentsListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleAssignmentsClient) listForScopeHandleError(resp *http.Response) error {
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

// ListForSubscription - List all role assignments that apply to a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForSubscription(options *RoleAssignmentsListForSubscriptionOptions) *RoleAssignmentsListForSubscriptionPager {
	return &RoleAssignmentsListForSubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForSubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsListForSubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listForSubscriptionCreateRequest creates the ListForSubscription request.
func (client *RoleAssignmentsClient) listForSubscriptionCreateRequest(ctx context.Context, options *RoleAssignmentsListForSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForSubscriptionHandleResponse handles the ListForSubscription response.
func (client *RoleAssignmentsClient) listForSubscriptionHandleResponse(resp *http.Response) (RoleAssignmentsListForSubscriptionResponse, error) {
	result := RoleAssignmentsListForSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsListForSubscriptionResponse{}, err
	}
	return result, nil
}

// listForSubscriptionHandleError handles the ListForSubscription error response.
func (client *RoleAssignmentsClient) listForSubscriptionHandleError(resp *http.Response) error {
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

// Validate - Validate a role assignment create or update operation by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Validate(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateOptions) (RoleAssignmentsValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentsValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsValidateResponse{}, client.validateHandleError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *RoleAssignmentsClient) validateCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}/validate"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateHandleResponse handles the Validate response.
func (client *RoleAssignmentsClient) validateHandleResponse(resp *http.Response) (RoleAssignmentsValidateResponse, error) {
	result := RoleAssignmentsValidateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidationResponse); err != nil {
		return RoleAssignmentsValidateResponse{}, err
	}
	return result, nil
}

// validateHandleError handles the Validate error response.
func (client *RoleAssignmentsClient) validateHandleError(resp *http.Response) error {
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

// ValidateByID - Validate a role assignment create or update operation by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ValidateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateByIDOptions) (RoleAssignmentsValidateByIDResponse, error) {
	req, err := client.validateByIDCreateRequest(ctx, roleAssignmentID, parameters, options)
	if err != nil {
		return RoleAssignmentsValidateByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsValidateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsValidateByIDResponse{}, client.validateByIDHandleError(resp)
	}
	return client.validateByIDHandleResponse(resp)
}

// validateByIDCreateRequest creates the ValidateByID request.
func (client *RoleAssignmentsClient) validateByIDCreateRequest(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}/validate"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateByIDHandleResponse handles the ValidateByID response.
func (client *RoleAssignmentsClient) validateByIDHandleResponse(resp *http.Response) (RoleAssignmentsValidateByIDResponse, error) {
	result := RoleAssignmentsValidateByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidationResponse); err != nil {
		return RoleAssignmentsValidateByIDResponse{}, err
	}
	return result, nil
}

// validateByIDHandleError handles the ValidateByID error response.
func (client *RoleAssignmentsClient) validateByIDHandleError(resp *http.Response) error {
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