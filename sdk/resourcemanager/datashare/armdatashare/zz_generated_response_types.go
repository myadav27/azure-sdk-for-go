//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	OperationResponse
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResponse struct {
	AccountList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.ListBySubscription.
type AccountsClientListBySubscriptionResponse struct {
	AccountList
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	Account
}

// ConsumerInvitationsClientGetResponse contains the response from method ConsumerInvitationsClient.Get.
type ConsumerInvitationsClientGetResponse struct {
	ConsumerInvitation
}

// ConsumerInvitationsClientListInvitationsResponse contains the response from method ConsumerInvitationsClient.ListInvitations.
type ConsumerInvitationsClientListInvitationsResponse struct {
	ConsumerInvitationList
}

// ConsumerInvitationsClientRejectInvitationResponse contains the response from method ConsumerInvitationsClient.RejectInvitation.
type ConsumerInvitationsClientRejectInvitationResponse struct {
	ConsumerInvitation
}

// ConsumerSourceDataSetsClientListByShareSubscriptionResponse contains the response from method ConsumerSourceDataSetsClient.ListByShareSubscription.
type ConsumerSourceDataSetsClientListByShareSubscriptionResponse struct {
	ConsumerSourceDataSetList
}

// DataSetMappingsClientCreateResponse contains the response from method DataSetMappingsClient.Create.
type DataSetMappingsClientCreateResponse struct {
	DataSetMappingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataSetMappingsClientCreateResponse.
func (d *DataSetMappingsClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataSetMappingClassification(data)
	if err != nil {
		return err
	}
	d.DataSetMappingClassification = res
	return nil
}

// DataSetMappingsClientDeleteResponse contains the response from method DataSetMappingsClient.Delete.
type DataSetMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataSetMappingsClientGetResponse contains the response from method DataSetMappingsClient.Get.
type DataSetMappingsClientGetResponse struct {
	DataSetMappingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataSetMappingsClientGetResponse.
func (d *DataSetMappingsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataSetMappingClassification(data)
	if err != nil {
		return err
	}
	d.DataSetMappingClassification = res
	return nil
}

// DataSetMappingsClientListByShareSubscriptionResponse contains the response from method DataSetMappingsClient.ListByShareSubscription.
type DataSetMappingsClientListByShareSubscriptionResponse struct {
	DataSetMappingList
}

// DataSetsClientCreateResponse contains the response from method DataSetsClient.Create.
type DataSetsClientCreateResponse struct {
	DataSetClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataSetsClientCreateResponse.
func (d *DataSetsClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataSetClassification(data)
	if err != nil {
		return err
	}
	d.DataSetClassification = res
	return nil
}

// DataSetsClientDeleteResponse contains the response from method DataSetsClient.Delete.
type DataSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataSetsClientGetResponse contains the response from method DataSetsClient.Get.
type DataSetsClientGetResponse struct {
	DataSetClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataSetsClientGetResponse.
func (d *DataSetsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataSetClassification(data)
	if err != nil {
		return err
	}
	d.DataSetClassification = res
	return nil
}

// DataSetsClientListByShareResponse contains the response from method DataSetsClient.ListByShare.
type DataSetsClientListByShareResponse struct {
	DataSetList
}

// EmailRegistrationsClientActivateEmailResponse contains the response from method EmailRegistrationsClient.ActivateEmail.
type EmailRegistrationsClientActivateEmailResponse struct {
	EmailRegistration
}

// EmailRegistrationsClientRegisterEmailResponse contains the response from method EmailRegistrationsClient.RegisterEmail.
type EmailRegistrationsClientRegisterEmailResponse struct {
	EmailRegistration
}

// InvitationsClientCreateResponse contains the response from method InvitationsClient.Create.
type InvitationsClientCreateResponse struct {
	Invitation
}

// InvitationsClientDeleteResponse contains the response from method InvitationsClient.Delete.
type InvitationsClientDeleteResponse struct {
	// placeholder for future response values
}

// InvitationsClientGetResponse contains the response from method InvitationsClient.Get.
type InvitationsClientGetResponse struct {
	Invitation
}

// InvitationsClientListByShareResponse contains the response from method InvitationsClient.ListByShare.
type InvitationsClientListByShareResponse struct {
	InvitationList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationList
}

// ProviderShareSubscriptionsClientAdjustResponse contains the response from method ProviderShareSubscriptionsClient.Adjust.
type ProviderShareSubscriptionsClientAdjustResponse struct {
	ProviderShareSubscription
}

// ProviderShareSubscriptionsClientGetByShareResponse contains the response from method ProviderShareSubscriptionsClient.GetByShare.
type ProviderShareSubscriptionsClientGetByShareResponse struct {
	ProviderShareSubscription
}

// ProviderShareSubscriptionsClientListByShareResponse contains the response from method ProviderShareSubscriptionsClient.ListByShare.
type ProviderShareSubscriptionsClientListByShareResponse struct {
	ProviderShareSubscriptionList
}

// ProviderShareSubscriptionsClientReinstateResponse contains the response from method ProviderShareSubscriptionsClient.Reinstate.
type ProviderShareSubscriptionsClientReinstateResponse struct {
	ProviderShareSubscription
}

// ProviderShareSubscriptionsClientRevokeResponse contains the response from method ProviderShareSubscriptionsClient.Revoke.
type ProviderShareSubscriptionsClientRevokeResponse struct {
	ProviderShareSubscription
}

// ShareSubscriptionsClientCancelSynchronizationResponse contains the response from method ShareSubscriptionsClient.CancelSynchronization.
type ShareSubscriptionsClientCancelSynchronizationResponse struct {
	ShareSubscriptionSynchronization
}

// ShareSubscriptionsClientCreateResponse contains the response from method ShareSubscriptionsClient.Create.
type ShareSubscriptionsClientCreateResponse struct {
	ShareSubscription
}

// ShareSubscriptionsClientDeleteResponse contains the response from method ShareSubscriptionsClient.Delete.
type ShareSubscriptionsClientDeleteResponse struct {
	OperationResponse
}

// ShareSubscriptionsClientGetResponse contains the response from method ShareSubscriptionsClient.Get.
type ShareSubscriptionsClientGetResponse struct {
	ShareSubscription
}

// ShareSubscriptionsClientListByAccountResponse contains the response from method ShareSubscriptionsClient.ListByAccount.
type ShareSubscriptionsClientListByAccountResponse struct {
	ShareSubscriptionList
}

// ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse contains the response from method ShareSubscriptionsClient.ListSourceShareSynchronizationSettings.
type ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse struct {
	SourceShareSynchronizationSettingList
}

// ShareSubscriptionsClientListSynchronizationDetailsResponse contains the response from method ShareSubscriptionsClient.ListSynchronizationDetails.
type ShareSubscriptionsClientListSynchronizationDetailsResponse struct {
	SynchronizationDetailsList
}

// ShareSubscriptionsClientListSynchronizationsResponse contains the response from method ShareSubscriptionsClient.ListSynchronizations.
type ShareSubscriptionsClientListSynchronizationsResponse struct {
	ShareSubscriptionSynchronizationList
}

// ShareSubscriptionsClientSynchronizeResponse contains the response from method ShareSubscriptionsClient.Synchronize.
type ShareSubscriptionsClientSynchronizeResponse struct {
	ShareSubscriptionSynchronization
}

// SharesClientCreateResponse contains the response from method SharesClient.Create.
type SharesClientCreateResponse struct {
	Share
}

// SharesClientDeleteResponse contains the response from method SharesClient.Delete.
type SharesClientDeleteResponse struct {
	OperationResponse
}

// SharesClientGetResponse contains the response from method SharesClient.Get.
type SharesClientGetResponse struct {
	Share
}

// SharesClientListByAccountResponse contains the response from method SharesClient.ListByAccount.
type SharesClientListByAccountResponse struct {
	ShareList
}

// SharesClientListSynchronizationDetailsResponse contains the response from method SharesClient.ListSynchronizationDetails.
type SharesClientListSynchronizationDetailsResponse struct {
	SynchronizationDetailsList
}

// SharesClientListSynchronizationsResponse contains the response from method SharesClient.ListSynchronizations.
type SharesClientListSynchronizationsResponse struct {
	ShareSynchronizationList
}

// SynchronizationSettingsClientCreateResponse contains the response from method SynchronizationSettingsClient.Create.
type SynchronizationSettingsClientCreateResponse struct {
	SynchronizationSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SynchronizationSettingsClientCreateResponse.
func (s *SynchronizationSettingsClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSynchronizationSettingClassification(data)
	if err != nil {
		return err
	}
	s.SynchronizationSettingClassification = res
	return nil
}

// SynchronizationSettingsClientDeleteResponse contains the response from method SynchronizationSettingsClient.Delete.
type SynchronizationSettingsClientDeleteResponse struct {
	OperationResponse
}

// SynchronizationSettingsClientGetResponse contains the response from method SynchronizationSettingsClient.Get.
type SynchronizationSettingsClientGetResponse struct {
	SynchronizationSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SynchronizationSettingsClientGetResponse.
func (s *SynchronizationSettingsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSynchronizationSettingClassification(data)
	if err != nil {
		return err
	}
	s.SynchronizationSettingClassification = res
	return nil
}

// SynchronizationSettingsClientListByShareResponse contains the response from method SynchronizationSettingsClient.ListByShare.
type SynchronizationSettingsClientListByShareResponse struct {
	SynchronizationSettingList
}

// TriggersClientCreateResponse contains the response from method TriggersClient.Create.
type TriggersClientCreateResponse struct {
	TriggerClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientCreateResponse.
func (t *TriggersClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientDeleteResponse contains the response from method TriggersClient.Delete.
type TriggersClientDeleteResponse struct {
	OperationResponse
}

// TriggersClientGetResponse contains the response from method TriggersClient.Get.
type TriggersClientGetResponse struct {
	TriggerClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientGetResponse.
func (t *TriggersClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientListByShareSubscriptionResponse contains the response from method TriggersClient.ListByShareSubscription.
type TriggersClientListByShareSubscriptionResponse struct {
	TriggerList
}
