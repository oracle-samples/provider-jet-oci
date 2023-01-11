/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VaultObservation struct {
	CryptoEndpoint *string `json:"cryptoEndpoint,omitempty" tf:"crypto_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	ManagementEndpoint *string `json:"managementEndpoint,omitempty" tf:"management_endpoint,omitempty"`

	ReplicaDetails []VaultReplicaDetailsObservation `json:"replicaDetails,omitempty" tf:"replica_details,omitempty"`

	RestoredFromVaultID *string `json:"restoredFromVaultId,omitempty" tf:"restored_from_vault_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type VaultParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	RestoreFromFile []VaultRestoreFromFileParameters `json:"restoreFromFile,omitempty" tf:"restore_from_file,omitempty"`

	// +kubebuilder:validation:Optional
	RestoreFromObjectStore []VaultRestoreFromObjectStoreParameters `json:"restoreFromObjectStore,omitempty" tf:"restore_from_object_store,omitempty"`

	// +kubebuilder:validation:Optional
	RestoreTrigger *bool `json:"restoreTrigger,omitempty" tf:"restore_trigger,omitempty"`

	// +kubebuilder:validation:Optional
	TimeOfDeletion *string `json:"timeOfDeletion,omitempty" tf:"time_of_deletion,omitempty"`

	// +kubebuilder:validation:Required
	VaultType *string `json:"vaultType" tf:"vault_type,omitempty"`
}

type VaultReplicaDetailsObservation struct {
	ReplicationID *string `json:"replicationId,omitempty" tf:"replication_id,omitempty"`
}

type VaultReplicaDetailsParameters struct {
}

type VaultRestoreFromFileObservation struct {
}

type VaultRestoreFromFileParameters struct {

	// +kubebuilder:validation:Required
	ContentLength *string `json:"contentLength" tf:"content_length,omitempty"`

	// +kubebuilder:validation:Optional
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// +kubebuilder:validation:Required
	RestoreVaultFromFileDetails *string `json:"restoreVaultFromFileDetails" tf:"restore_vault_from_file_details,omitempty"`
}

type VaultRestoreFromObjectStoreObservation struct {
}

type VaultRestoreFromObjectStoreParameters struct {

	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

// VaultSpec defines the desired state of Vault
type VaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultParameters `json:"forProvider"`
}

// VaultStatus defines the observed state of Vault.
type VaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vault is the Schema for the Vaults API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSpec   `json:"spec"`
	Status            VaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultList contains a list of Vaults
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Repository type metadata.
var (
	Vault_Kind             = "Vault"
	Vault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vault_Kind}.String()
	Vault_KindAPIVersion   = Vault_Kind + "." + CRDGroupVersion.String()
	Vault_GroupVersionKind = CRDGroupVersion.WithKind(Vault_Kind)
)

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
