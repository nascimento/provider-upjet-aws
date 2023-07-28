/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ThingGroupMembershipObservation struct {

	// The membership ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Override dynamic thing groups with static thing groups when 10-group limit is reached. If a thing belongs to 10 thing groups, and one or more of those groups are dynamic thing groups, adding a thing to a static group removes the thing from the last dynamic group.
	OverrideDynamicGroup *bool `json:"overrideDynamicGroup,omitempty" tf:"override_dynamic_group,omitempty"`

	// The name of the group to which you are adding a thing.
	ThingGroupName *string `json:"thingGroupName,omitempty" tf:"thing_group_name,omitempty"`

	// The name of the thing to add to a group.
	ThingName *string `json:"thingName,omitempty" tf:"thing_name,omitempty"`
}

type ThingGroupMembershipParameters struct {

	// Override dynamic thing groups with static thing groups when 10-group limit is reached. If a thing belongs to 10 thing groups, and one or more of those groups are dynamic thing groups, adding a thing to a static group removes the thing from the last dynamic group.
	// +kubebuilder:validation:Optional
	OverrideDynamicGroup *bool `json:"overrideDynamicGroup,omitempty" tf:"override_dynamic_group,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the group to which you are adding a thing.
	// +kubebuilder:validation:Optional
	ThingGroupName *string `json:"thingGroupName,omitempty" tf:"thing_group_name,omitempty"`

	// The name of the thing to add to a group.
	// +kubebuilder:validation:Optional
	ThingName *string `json:"thingName,omitempty" tf:"thing_name,omitempty"`
}

// ThingGroupMembershipSpec defines the desired state of ThingGroupMembership
type ThingGroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingGroupMembershipParameters `json:"forProvider"`
}

// ThingGroupMembershipStatus defines the observed state of ThingGroupMembership.
type ThingGroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingGroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ThingGroupMembership is the Schema for the ThingGroupMemberships API. Adds an IoT Thing to an IoT Thing Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ThingGroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.thingGroupName)",message="thingGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.thingName)",message="thingName is a required parameter"
	Spec   ThingGroupMembershipSpec   `json:"spec"`
	Status ThingGroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingGroupMembershipList contains a list of ThingGroupMemberships
type ThingGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThingGroupMembership `json:"items"`
}

// Repository type metadata.
var (
	ThingGroupMembership_Kind             = "ThingGroupMembership"
	ThingGroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThingGroupMembership_Kind}.String()
	ThingGroupMembership_KindAPIVersion   = ThingGroupMembership_Kind + "." + CRDGroupVersion.String()
	ThingGroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(ThingGroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&ThingGroupMembership{}, &ThingGroupMembershipList{})
}
