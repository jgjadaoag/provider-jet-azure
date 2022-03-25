/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubEndpointEventHub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubEndpointEventHub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubEndpointEventHub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubEndpointEventHub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubEndpointEventHub.
func (mg *IOTHubEndpointEventHub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubEndpointServiceBusQueue.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubEndpointServiceBusQueue) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubEndpointServiceBusQueue.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubEndpointServiceBusQueue) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubEndpointServiceBusQueue.
func (mg *IOTHubEndpointServiceBusQueue) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubEndpointServiceBusTopic.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubEndpointServiceBusTopic) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubEndpointServiceBusTopic.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubEndpointServiceBusTopic) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubEndpointServiceBusTopic.
func (mg *IOTHubEndpointServiceBusTopic) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubEnrichment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubEnrichment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubEnrichment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubEnrichment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubEnrichment.
func (mg *IOTHubEnrichment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IOTHubRoute.
func (mg *IOTHubRoute) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubRoute.
func (mg *IOTHubRoute) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IOTHubRoute.
func (mg *IOTHubRoute) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubRoute.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubRoute) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubRoute.
func (mg *IOTHubRoute) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubRoute.
func (mg *IOTHubRoute) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubRoute.
func (mg *IOTHubRoute) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubRoute.
func (mg *IOTHubRoute) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IOTHubRoute.
func (mg *IOTHubRoute) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubRoute.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubRoute) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubRoute.
func (mg *IOTHubRoute) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubRoute.
func (mg *IOTHubRoute) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
