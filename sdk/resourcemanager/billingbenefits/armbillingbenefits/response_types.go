//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// RPClientValidatePurchaseResponse contains the response from method RPClient.ValidatePurchase.
type RPClientValidatePurchaseResponse struct {
	SavingsPlanValidateResponse
}

// ReservationOrderAliasClientCreateResponse contains the response from method ReservationOrderAliasClient.Create.
type ReservationOrderAliasClientCreateResponse struct {
	ReservationOrderAliasResponse
}

// ReservationOrderAliasClientGetResponse contains the response from method ReservationOrderAliasClient.Get.
type ReservationOrderAliasClientGetResponse struct {
	ReservationOrderAliasResponse
}

// SavingsPlanClientGetResponse contains the response from method SavingsPlanClient.Get.
type SavingsPlanClientGetResponse struct {
	SavingsPlanModel
}

// SavingsPlanClientListAllResponse contains the response from method SavingsPlanClient.ListAll.
type SavingsPlanClientListAllResponse struct {
	SavingsPlanModelListResult
}

// SavingsPlanClientListResponse contains the response from method SavingsPlanClient.List.
type SavingsPlanClientListResponse struct {
	SavingsPlanModelList
}

// SavingsPlanClientUpdateResponse contains the response from method SavingsPlanClient.Update.
type SavingsPlanClientUpdateResponse struct {
	SavingsPlanModel
	// Location contains the information returned from the Location header response.
	Location *string
}

// SavingsPlanClientValidateUpdateResponse contains the response from method SavingsPlanClient.ValidateUpdate.
type SavingsPlanClientValidateUpdateResponse struct {
	SavingsPlanValidateResponse
}

// SavingsPlanOrderAliasClientCreateResponse contains the response from method SavingsPlanOrderAliasClient.Create.
type SavingsPlanOrderAliasClientCreateResponse struct {
	SavingsPlanOrderAliasModel
}

// SavingsPlanOrderAliasClientGetResponse contains the response from method SavingsPlanOrderAliasClient.Get.
type SavingsPlanOrderAliasClientGetResponse struct {
	SavingsPlanOrderAliasModel
}

// SavingsPlanOrderClientElevateResponse contains the response from method SavingsPlanOrderClient.Elevate.
type SavingsPlanOrderClientElevateResponse struct {
	RoleAssignmentEntity
}

// SavingsPlanOrderClientGetResponse contains the response from method SavingsPlanOrderClient.Get.
type SavingsPlanOrderClientGetResponse struct {
	SavingsPlanOrderModel
}

// SavingsPlanOrderClientListResponse contains the response from method SavingsPlanOrderClient.List.
type SavingsPlanOrderClientListResponse struct {
	SavingsPlanOrderModelList
}
