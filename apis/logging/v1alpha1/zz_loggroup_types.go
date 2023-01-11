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

type LogGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeLastModified *string `json:"timeLastModified,omitempty" tf:"time_last_modified,omitempty"`
}

type LogGroupParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

// LogGroupSpec defines the desired state of LogGroup
type LogGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogGroupParameters `json:"forProvider"`
}

// LogGroupStatus defines the observed state of LogGroup.
type LogGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogGroup is the Schema for the LogGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type LogGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogGroupSpec   `json:"spec"`
	Status            LogGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogGroupList contains a list of LogGroups
type LogGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogGroup `json:"items"`
}

// Repository type metadata.
var (
	LogGroup_Kind             = "LogGroup"
	LogGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogGroup_Kind}.String()
	LogGroup_KindAPIVersion   = LogGroup_Kind + "." + CRDGroupVersion.String()
	LogGroup_GroupVersionKind = CRDGroupVersion.WithKind(LogGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&LogGroup{}, &LogGroupList{})
}
