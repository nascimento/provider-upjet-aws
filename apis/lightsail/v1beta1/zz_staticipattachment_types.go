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

type StaticIPAttachmentInitParameters struct {

	// The name of the Lightsail instance to attach the IP to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// The name of the allocated static IP
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.StaticIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StaticIPName *string `json:"staticIpName,omitempty" tf:"static_ip_name,omitempty"`

	// Reference to a StaticIP in lightsail to populate staticIpName.
	// +kubebuilder:validation:Optional
	StaticIPNameRef *v1.Reference `json:"staticIpNameRef,omitempty" tf:"-"`

	// Selector for a StaticIP in lightsail to populate staticIpName.
	// +kubebuilder:validation:Optional
	StaticIPNameSelector *v1.Selector `json:"staticIpNameSelector,omitempty" tf:"-"`
}

type StaticIPAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The allocated static IP address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name of the Lightsail instance to attach the IP to
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// The name of the allocated static IP
	StaticIPName *string `json:"staticIpName,omitempty" tf:"static_ip_name,omitempty"`
}

type StaticIPAttachmentParameters struct {

	// The name of the Lightsail instance to attach the IP to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the allocated static IP
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.StaticIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StaticIPName *string `json:"staticIpName,omitempty" tf:"static_ip_name,omitempty"`

	// Reference to a StaticIP in lightsail to populate staticIpName.
	// +kubebuilder:validation:Optional
	StaticIPNameRef *v1.Reference `json:"staticIpNameRef,omitempty" tf:"-"`

	// Selector for a StaticIP in lightsail to populate staticIpName.
	// +kubebuilder:validation:Optional
	StaticIPNameSelector *v1.Selector `json:"staticIpNameSelector,omitempty" tf:"-"`
}

// StaticIPAttachmentSpec defines the desired state of StaticIPAttachment
type StaticIPAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StaticIPAttachmentParameters `json:"forProvider"`
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
	InitProvider StaticIPAttachmentInitParameters `json:"initProvider,omitempty"`
}

// StaticIPAttachmentStatus defines the observed state of StaticIPAttachment.
type StaticIPAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StaticIPAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StaticIPAttachment is the Schema for the StaticIPAttachments API. Provides an Lightsail Static IP Attachment
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StaticIPAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StaticIPAttachmentSpec   `json:"spec"`
	Status            StaticIPAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticIPAttachmentList contains a list of StaticIPAttachments
type StaticIPAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticIPAttachment `json:"items"`
}

// Repository type metadata.
var (
	StaticIPAttachment_Kind             = "StaticIPAttachment"
	StaticIPAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StaticIPAttachment_Kind}.String()
	StaticIPAttachment_KindAPIVersion   = StaticIPAttachment_Kind + "." + CRDGroupVersion.String()
	StaticIPAttachment_GroupVersionKind = CRDGroupVersion.WithKind(StaticIPAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&StaticIPAttachment{}, &StaticIPAttachmentList{})
}
