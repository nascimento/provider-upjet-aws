// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SQLInjectionMatchSetInitParameters struct {

	// The name or description of the SizeConstraintSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SQLInjectionMatchTuple []SQLInjectionMatchTupleInitParameters `json:"sqlInjectionMatchTuple,omitempty" tf:"sql_injection_match_tuple,omitempty"`
}

type SQLInjectionMatchSetObservation struct {

	// The ID of the WAF SqlInjectionMatchSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the SizeConstraintSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SQLInjectionMatchTuple []SQLInjectionMatchTupleObservation `json:"sqlInjectionMatchTuple,omitempty" tf:"sql_injection_match_tuple,omitempty"`
}

type SQLInjectionMatchSetParameters struct {

	// The name or description of the SizeConstraintSet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	// +kubebuilder:validation:Optional
	SQLInjectionMatchTuple []SQLInjectionMatchTupleParameters `json:"sqlInjectionMatchTuple,omitempty" tf:"sql_injection_match_tuple,omitempty"`
}

type SQLInjectionMatchTupleFieldToMatchInitParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLInjectionMatchTupleFieldToMatchObservation struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLInjectionMatchTupleFieldToMatchParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SQLInjectionMatchTupleInitParameters struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch *SQLInjectionMatchTupleFieldToMatchInitParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type SQLInjectionMatchTupleObservation struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch *SQLInjectionMatchTupleFieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type SQLInjectionMatchTupleParameters struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	// +kubebuilder:validation:Optional
	FieldToMatch *SQLInjectionMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// SQLInjectionMatchSetSpec defines the desired state of SQLInjectionMatchSet
type SQLInjectionMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLInjectionMatchSetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SQLInjectionMatchSetInitParameters `json:"initProvider,omitempty"`
}

// SQLInjectionMatchSetStatus defines the observed state of SQLInjectionMatchSet.
type SQLInjectionMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLInjectionMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SQLInjectionMatchSet is the Schema for the SQLInjectionMatchSets API. Provides a AWS WAF Regional SqlInjectionMatchSet resource for use with ALB.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SQLInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SQLInjectionMatchSetSpec   `json:"spec"`
	Status SQLInjectionMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLInjectionMatchSetList contains a list of SQLInjectionMatchSets
type SQLInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLInjectionMatchSet `json:"items"`
}

// Repository type metadata.
var (
	SQLInjectionMatchSet_Kind             = "SQLInjectionMatchSet"
	SQLInjectionMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLInjectionMatchSet_Kind}.String()
	SQLInjectionMatchSet_KindAPIVersion   = SQLInjectionMatchSet_Kind + "." + CRDGroupVersion.String()
	SQLInjectionMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(SQLInjectionMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLInjectionMatchSet{}, &SQLInjectionMatchSetList{})
}
