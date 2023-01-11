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

type GenericArtifactObservation struct {
	ArtifactPath *string `json:"artifactPath,omitempty" tf:"artifact_path,omitempty"`

	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	SizeInBytes *string `json:"sizeInBytes,omitempty" tf:"size_in_bytes,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type GenericArtifactParameters struct {

	// +kubebuilder:validation:Required
	ArtifactID *string `json:"artifactId" tf:"artifact_id,omitempty"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

// GenericArtifactSpec defines the desired state of GenericArtifact
type GenericArtifactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GenericArtifactParameters `json:"forProvider"`
}

// GenericArtifactStatus defines the observed state of GenericArtifact.
type GenericArtifactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GenericArtifactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GenericArtifact is the Schema for the GenericArtifacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type GenericArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GenericArtifactSpec   `json:"spec"`
	Status            GenericArtifactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GenericArtifactList contains a list of GenericArtifacts
type GenericArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GenericArtifact `json:"items"`
}

// Repository type metadata.
var (
	GenericArtifact_Kind             = "GenericArtifact"
	GenericArtifact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GenericArtifact_Kind}.String()
	GenericArtifact_KindAPIVersion   = GenericArtifact_Kind + "." + CRDGroupVersion.String()
	GenericArtifact_GroupVersionKind = CRDGroupVersion.WithKind(GenericArtifact_Kind)
)

func init() {
	SchemeBuilder.Register(&GenericArtifact{}, &GenericArtifactList{})
}
