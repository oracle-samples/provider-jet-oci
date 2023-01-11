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

type SteeringPolicyAttachmentObservation struct {
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Rtypes []*string `json:"rtypes,omitempty" tf:"rtypes,omitempty"`

	Self *string `json:"self,omitempty" tf:"self,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type SteeringPolicyAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +crossplane:generate:reference:type=SteeringPolicy
	// +kubebuilder:validation:Optional
	SteeringPolicyID *string `json:"steeringPolicyId,omitempty" tf:"steering_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	SteeringPolicyIDRef *v1.Reference `json:"steeringPolicyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SteeringPolicyIDSelector *v1.Selector `json:"steeringPolicyIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// SteeringPolicyAttachmentSpec defines the desired state of SteeringPolicyAttachment
type SteeringPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SteeringPolicyAttachmentParameters `json:"forProvider"`
}

// SteeringPolicyAttachmentStatus defines the observed state of SteeringPolicyAttachment.
type SteeringPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SteeringPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SteeringPolicyAttachment is the Schema for the SteeringPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type SteeringPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SteeringPolicyAttachmentSpec   `json:"spec"`
	Status            SteeringPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SteeringPolicyAttachmentList contains a list of SteeringPolicyAttachments
type SteeringPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SteeringPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	SteeringPolicyAttachment_Kind             = "SteeringPolicyAttachment"
	SteeringPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SteeringPolicyAttachment_Kind}.String()
	SteeringPolicyAttachment_KindAPIVersion   = SteeringPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	SteeringPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(SteeringPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&SteeringPolicyAttachment{}, &SteeringPolicyAttachmentList{})
}
