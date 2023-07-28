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

type ThingTypeObservation struct {

	// The ARN of the created AWS IoT Thing Type.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated *bool `json:"deprecated,omitempty" tf:"deprecated,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the thing type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// , Configuration block that can contain the following properties of the thing type:
	Properties []ThingTypePropertiesObservation `json:"properties,omitempty" tf:"properties,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ThingTypeParameters struct {

	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	// +kubebuilder:validation:Optional
	Deprecated *bool `json:"deprecated,omitempty" tf:"deprecated,omitempty"`

	// The name of the thing type.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// , Configuration block that can contain the following properties of the thing type:
	// +kubebuilder:validation:Optional
	Properties []ThingTypePropertiesParameters `json:"properties,omitempty" tf:"properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ThingTypePropertiesObservation struct {

	// The description of the thing type.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of searchable thing attribute names.
	SearchableAttributes []*string `json:"searchableAttributes,omitempty" tf:"searchable_attributes,omitempty"`
}

type ThingTypePropertiesParameters struct {

	// The description of the thing type.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of searchable thing attribute names.
	// +kubebuilder:validation:Optional
	SearchableAttributes []*string `json:"searchableAttributes,omitempty" tf:"searchable_attributes,omitempty"`
}

// ThingTypeSpec defines the desired state of ThingType
type ThingTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingTypeParameters `json:"forProvider"`
}

// ThingTypeStatus defines the observed state of ThingType.
type ThingTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ThingType is the Schema for the ThingTypes API. Creates and manages an AWS IoT Thing Type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ThingType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ThingTypeSpec   `json:"spec"`
	Status ThingTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingTypeList contains a list of ThingTypes
type ThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThingType `json:"items"`
}

// Repository type metadata.
var (
	ThingType_Kind             = "ThingType"
	ThingType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThingType_Kind}.String()
	ThingType_KindAPIVersion   = ThingType_Kind + "." + CRDGroupVersion.String()
	ThingType_GroupVersionKind = CRDGroupVersion.WithKind(ThingType_Kind)
)

func init() {
	SchemeBuilder.Register(&ThingType{}, &ThingTypeList{})
}
