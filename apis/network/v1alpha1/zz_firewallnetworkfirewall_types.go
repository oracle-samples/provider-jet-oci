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

type FirewallNetworkFirewallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type FirewallNetworkFirewallParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

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
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	Ipv4Address *string `json:"ipv4address,omitempty" tf:"ipv4address,omitempty"`

	// +kubebuilder:validation:Optional
	Ipv6Address *string `json:"ipv6address,omitempty" tf:"ipv6address,omitempty"`

	// +crossplane:generate:reference:type=FirewallNetworkFirewallPolicy
	// +kubebuilder:validation:Optional
	NetworkFirewallPolicyID *string `json:"networkFirewallPolicyId,omitempty" tf:"network_firewall_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkFirewallPolicyIDRef *v1.Reference `json:"networkFirewallPolicyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkFirewallPolicyIDSelector *v1.Selector `json:"networkFirewallPolicyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIds []*string `json:"networkSecurityGroupIds,omitempty" tf:"network_security_group_ids,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// FirewallNetworkFirewallSpec defines the desired state of FirewallNetworkFirewall
type FirewallNetworkFirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallNetworkFirewallParameters `json:"forProvider"`
}

// FirewallNetworkFirewallStatus defines the observed state of FirewallNetworkFirewall.
type FirewallNetworkFirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallNetworkFirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkFirewall is the Schema for the FirewallNetworkFirewalls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type FirewallNetworkFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNetworkFirewallSpec   `json:"spec"`
	Status            FirewallNetworkFirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkFirewallList contains a list of FirewallNetworkFirewalls
type FirewallNetworkFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallNetworkFirewall `json:"items"`
}

// Repository type metadata.
var (
	FirewallNetworkFirewall_Kind             = "FirewallNetworkFirewall"
	FirewallNetworkFirewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallNetworkFirewall_Kind}.String()
	FirewallNetworkFirewall_KindAPIVersion   = FirewallNetworkFirewall_Kind + "." + CRDGroupVersion.String()
	FirewallNetworkFirewall_GroupVersionKind = CRDGroupVersion.WithKind(FirewallNetworkFirewall_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallNetworkFirewall{}, &FirewallNetworkFirewallList{})
}
