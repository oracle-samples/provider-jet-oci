/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteTableAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteTableAttachmentParameters struct {

	// The OCID of the route table.
	// +kubebuilder:validation:Required
	RouteTableID *string `json:"routeTableId" tf:"route_table_id,omitempty"`

	// The OCID of the subnet.
	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// RouteTableAttachmentSpec defines the desired state of RouteTableAttachment
type RouteTableAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableAttachmentParameters `json:"forProvider"`
}

// RouteTableAttachmentStatus defines the observed state of RouteTableAttachment.
type RouteTableAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAttachment is the Schema for the RouteTableAttachments API. Provides the ability to associate a route table for a subnet in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type RouteTableAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableAttachmentSpec   `json:"spec"`
	Status            RouteTableAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAttachmentList contains a list of RouteTableAttachments
type RouteTableAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAttachment `json:"items"`
}

// Repository type metadata.
var (
	RouteTableAttachment_Kind             = "RouteTableAttachment"
	RouteTableAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTableAttachment_Kind}.String()
	RouteTableAttachment_KindAPIVersion   = RouteTableAttachment_Kind + "." + CRDGroupVersion.String()
	RouteTableAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RouteTableAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTableAttachment{}, &RouteTableAttachmentList{})
}
