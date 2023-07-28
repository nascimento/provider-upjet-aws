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

type AppCookieStickinessPolicyObservation struct {

	// Application cookie whose lifetime the ELB's cookie should follow.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// Name of load balancer to which the policy
	// should be attached.
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`
}

type AppCookieStickinessPolicyParameters struct {

	// Application cookie whose lifetime the ELB's cookie should follow.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	// +kubebuilder:validation:Required
	LBPort *float64 `json:"lbPort" tf:"lb_port,omitempty"`

	// Name of load balancer to which the policy
	// should be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +kubebuilder:validation:Optional
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// Reference to a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerRef *v1.Reference `json:"loadBalancerRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerSelector *v1.Selector `json:"loadBalancerSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AppCookieStickinessPolicySpec defines the desired state of AppCookieStickinessPolicy
type AppCookieStickinessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppCookieStickinessPolicyParameters `json:"forProvider"`
}

// AppCookieStickinessPolicyStatus defines the observed state of AppCookieStickinessPolicy.
type AppCookieStickinessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppCookieStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppCookieStickinessPolicy is the Schema for the AppCookieStickinessPolicys API. Provides an application cookie stickiness policy, which allows an ELB to wed its stickiness cookie to a cookie generated by your application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AppCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cookieName)",message="cookieName is a required parameter"
	Spec   AppCookieStickinessPolicySpec   `json:"spec"`
	Status AppCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppCookieStickinessPolicyList contains a list of AppCookieStickinessPolicys
type AppCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppCookieStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AppCookieStickinessPolicy_Kind             = "AppCookieStickinessPolicy"
	AppCookieStickinessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppCookieStickinessPolicy_Kind}.String()
	AppCookieStickinessPolicy_KindAPIVersion   = AppCookieStickinessPolicy_Kind + "." + CRDGroupVersion.String()
	AppCookieStickinessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AppCookieStickinessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AppCookieStickinessPolicy{}, &AppCookieStickinessPolicyList{})
}
