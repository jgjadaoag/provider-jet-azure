/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CaptureDescriptionObservation struct {
}

type CaptureDescriptionParameters struct {

	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Encoding *string `json:"encoding" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	SizeLimitInBytes *int64 `json:"sizeLimitInBytes,omitempty" tf:"size_limit_in_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty" tf:"skip_empty_archives,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Required
	ArchiveNameFormat *string `json:"archiveNameFormat" tf:"archive_name_format,omitempty"`

	// +kubebuilder:validation:Required
	BlobContainerName *string `json:"blobContainerName" tf:"blob_container_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type EventHubObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PartitionIds []*string `json:"partitionIds,omitempty" tf:"partition_ids,omitempty"`
}

type EventHubParameters struct {

	// +kubebuilder:validation:Optional
	CaptureDescription []CaptureDescriptionParameters `json:"captureDescription,omitempty" tf:"capture_description,omitempty"`

	// +kubebuilder:validation:Required
	MessageRetention *int64 `json:"messageRetention" tf:"message_retention,omitempty"`

	// +crossplane:generate:reference:type=EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	PartitionCount *int64 `json:"partitionCount" tf:"partition_count,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// EventHubSpec defines the desired state of EventHub
type EventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventHubParameters `json:"forProvider"`
}

// EventHubStatus defines the observed state of EventHub.
type EventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventHub is the Schema for the EventHubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventHubSpec   `json:"spec"`
	Status            EventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubList contains a list of EventHubs
type EventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventHub `json:"items"`
}

// Repository type metadata.
var (
	EventHub_Kind             = "EventHub"
	EventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventHub_Kind}.String()
	EventHub_KindAPIVersion   = EventHub_Kind + "." + CRDGroupVersion.String()
	EventHub_GroupVersionKind = CRDGroupVersion.WithKind(EventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&EventHub{}, &EventHubList{})
}
