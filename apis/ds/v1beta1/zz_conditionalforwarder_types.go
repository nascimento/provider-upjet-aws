// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionalForwarderInitParameters struct {

	// A list of forwarder IP addresses.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`
}

type ConditionalForwarderObservation struct {

	// A list of forwarder IP addresses.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// ID of directory.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fully qualified domain name of the remote domain for which forwarders will be used.
	RemoteDomainName *string `json:"remoteDomainName,omitempty" tf:"remote_domain_name,omitempty"`
}

type ConditionalForwarderParameters struct {

	// A list of forwarder IP addresses.
	// +kubebuilder:validation:Optional
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// ID of directory.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta2.Directory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The fully qualified domain name of the remote domain for which forwarders will be used.
	// +kubebuilder:validation:Required
	RemoteDomainName *string `json:"remoteDomainName" tf:"remote_domain_name,omitempty"`
}

// ConditionalForwarderSpec defines the desired state of ConditionalForwarder
type ConditionalForwarderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConditionalForwarderParameters `json:"forProvider"`
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
	InitProvider ConditionalForwarderInitParameters `json:"initProvider,omitempty"`
}

// ConditionalForwarderStatus defines the observed state of ConditionalForwarder.
type ConditionalForwarderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConditionalForwarderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConditionalForwarder is the Schema for the ConditionalForwarders API. Provides a conditional forwarder for managed Microsoft AD in AWS Directory Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsIps) || (has(self.initProvider) && has(self.initProvider.dnsIps))",message="spec.forProvider.dnsIps is a required parameter"
	Spec   ConditionalForwarderSpec   `json:"spec"`
	Status ConditionalForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionalForwarderList contains a list of ConditionalForwarders
type ConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConditionalForwarder `json:"items"`
}

// Repository type metadata.
var (
	ConditionalForwarder_Kind             = "ConditionalForwarder"
	ConditionalForwarder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConditionalForwarder_Kind}.String()
	ConditionalForwarder_KindAPIVersion   = ConditionalForwarder_Kind + "." + CRDGroupVersion.String()
	ConditionalForwarder_GroupVersionKind = CRDGroupVersion.WithKind(ConditionalForwarder_Kind)
)

func init() {
	SchemeBuilder.Register(&ConditionalForwarder{}, &ConditionalForwarderList{})
}
