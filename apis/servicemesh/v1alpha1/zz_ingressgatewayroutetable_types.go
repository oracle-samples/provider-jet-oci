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

type DestinationsObservation struct {
}

type DestinationsParameters struct {

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +crossplane:generate:reference:type=VirtualService
	// +kubebuilder:validation:Optional
	VirtualServiceID *string `json:"virtualServiceId,omitempty" tf:"virtual_service_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualServiceIDRef *v1.Reference `json:"virtualServiceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualServiceIDSelector *v1.Selector `json:"virtualServiceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type IngressGatewayHostObservation struct {
}

type IngressGatewayHostParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type IngressGatewayRouteTableObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type IngressGatewayRouteTableParameters struct {

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

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +crossplane:generate:reference:type=IngressGateway
	// +kubebuilder:validation:Optional
	IngressGatewayID *string `json:"ingressGatewayId,omitempty" tf:"ingress_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	IngressGatewayIDRef *v1.Reference `json:"ingressGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IngressGatewayIDSelector *v1.Selector `json:"ingressGatewayIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	RouteRules []RouteRulesParameters `json:"routeRules" tf:"route_rules,omitempty"`
}

type RouteRulesObservation struct {
}

type RouteRulesParameters struct {

	// +kubebuilder:validation:Required
	Destinations []DestinationsParameters `json:"destinations" tf:"destinations,omitempty"`

	// +kubebuilder:validation:Optional
	IngressGatewayHost []IngressGatewayHostParameters `json:"ingressGatewayHost,omitempty" tf:"ingress_gateway_host,omitempty"`

	// +kubebuilder:validation:Optional
	IsGRPC *bool `json:"isGrpc,omitempty" tf:"is_grpc,omitempty"`

	// +kubebuilder:validation:Optional
	IsHostRewriteEnabled *bool `json:"isHostRewriteEnabled,omitempty" tf:"is_host_rewrite_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IsPathRewriteEnabled *bool `json:"isPathRewriteEnabled,omitempty" tf:"is_path_rewrite_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	PathType *string `json:"pathType,omitempty" tf:"path_type,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// IngressGatewayRouteTableSpec defines the desired state of IngressGatewayRouteTable
type IngressGatewayRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IngressGatewayRouteTableParameters `json:"forProvider"`
}

// IngressGatewayRouteTableStatus defines the observed state of IngressGatewayRouteTable.
type IngressGatewayRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IngressGatewayRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IngressGatewayRouteTable is the Schema for the IngressGatewayRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type IngressGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressGatewayRouteTableSpec   `json:"spec"`
	Status            IngressGatewayRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IngressGatewayRouteTableList contains a list of IngressGatewayRouteTables
type IngressGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressGatewayRouteTable `json:"items"`
}

// Repository type metadata.
var (
	IngressGatewayRouteTable_Kind             = "IngressGatewayRouteTable"
	IngressGatewayRouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IngressGatewayRouteTable_Kind}.String()
	IngressGatewayRouteTable_KindAPIVersion   = IngressGatewayRouteTable_Kind + "." + CRDGroupVersion.String()
	IngressGatewayRouteTable_GroupVersionKind = CRDGroupVersion.WithKind(IngressGatewayRouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&IngressGatewayRouteTable{}, &IngressGatewayRouteTableList{})
}
