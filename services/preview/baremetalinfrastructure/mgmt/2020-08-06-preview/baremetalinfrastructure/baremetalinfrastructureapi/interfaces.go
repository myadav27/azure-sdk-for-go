package baremetalinfrastructureapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/baremetalinfrastructure/mgmt/2020-08-06-preview/baremetalinfrastructure"
)

// AzureBareMetalInstancesClientAPI contains the set of methods on the AzureBareMetalInstancesClient type.
type AzureBareMetalInstancesClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result baremetalinfrastructure.AzureBareMetalInstancesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result baremetalinfrastructure.AzureBareMetalInstance, err error)
	List(ctx context.Context, resourceGroupName string) (result baremetalinfrastructure.AzureBareMetalInstancesListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result baremetalinfrastructure.AzureBareMetalInstancesListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result baremetalinfrastructure.AzureBareMetalInstancesListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result baremetalinfrastructure.AzureBareMetalInstancesListResultIterator, err error)
	Restart(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result baremetalinfrastructure.AzureBareMetalInstancesRestartFuture, err error)
	Shutdown(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result baremetalinfrastructure.AzureBareMetalInstancesShutdownFuture, err error)
	Start(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result baremetalinfrastructure.AzureBareMetalInstancesStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter baremetalinfrastructure.Tags) (result baremetalinfrastructure.AzureBareMetalInstance, err error)
}

var _ AzureBareMetalInstancesClientAPI = (*baremetalinfrastructure.AzureBareMetalInstancesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result baremetalinfrastructure.OperationList, err error)
}

var _ OperationsClientAPI = (*baremetalinfrastructure.OperationsClient)(nil)