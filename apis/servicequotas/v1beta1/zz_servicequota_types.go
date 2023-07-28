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

type MetricDimensionsObservation struct {
	Class *string `json:"class,omitempty" tf:"class,omitempty"`

	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MetricDimensionsParameters struct {
}

type ServiceQuotaObservation struct {

	// Whether the service quota can be increased.
	Adjustable *bool `json:"adjustable,omitempty" tf:"adjustable,omitempty"`

	// Amazon Resource Name (ARN) of the service quota.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Default value of the service quota.
	DefaultValue *float64 `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// Service code and quota code, separated by a front slash (/)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Code of the service quota to track. For example: L-F678F1CE. Available values can be found with the AWS CLI service-quotas list-service-quotas command.
	QuotaCode *string `json:"quotaCode,omitempty" tf:"quota_code,omitempty"`

	// Name of the quota.
	QuotaName *string `json:"quotaName,omitempty" tf:"quota_name,omitempty"`

	// Service code and quota code, separated by a front slash (/)
	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	RequestStatus *string `json:"requestStatus,omitempty" tf:"request_status,omitempty"`

	// Code of the service to track. For example: vpc. Available values can be found with the AWS CLI service-quotas list-services command.
	ServiceCode *string `json:"serviceCode,omitempty" tf:"service_code,omitempty"`

	// Name of the service.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Information about the measurement.
	UsageMetric []UsageMetricObservation `json:"usageMetric,omitempty" tf:"usage_metric,omitempty"`

	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type ServiceQuotaParameters struct {

	// Code of the service quota to track. For example: L-F678F1CE. Available values can be found with the AWS CLI service-quotas list-service-quotas command.
	// +kubebuilder:validation:Optional
	QuotaCode *string `json:"quotaCode,omitempty" tf:"quota_code,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Code of the service to track. For example: vpc. Available values can be found with the AWS CLI service-quotas list-services command.
	// +kubebuilder:validation:Optional
	ServiceCode *string `json:"serviceCode,omitempty" tf:"service_code,omitempty"`

	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type UsageMetricObservation struct {

	// The metric dimensions.
	MetricDimensions []MetricDimensionsObservation `json:"metricDimensions,omitempty" tf:"metric_dimensions,omitempty"`

	// The name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace of the metric.
	MetricNamespace *string `json:"metricNamespace,omitempty" tf:"metric_namespace,omitempty"`

	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation *string `json:"metricStatisticRecommendation,omitempty" tf:"metric_statistic_recommendation,omitempty"`
}

type UsageMetricParameters struct {
}

// ServiceQuotaSpec defines the desired state of ServiceQuota
type ServiceQuotaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceQuotaParameters `json:"forProvider"`
}

// ServiceQuotaStatus defines the observed state of ServiceQuota.
type ServiceQuotaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceQuotaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceQuota is the Schema for the ServiceQuotas API. Manages an individual Service Quota
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServiceQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.quotaCode)",message="quotaCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceCode)",message="serviceCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value)",message="value is a required parameter"
	Spec   ServiceQuotaSpec   `json:"spec"`
	Status ServiceQuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceQuotaList contains a list of ServiceQuotas
type ServiceQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceQuota `json:"items"`
}

// Repository type metadata.
var (
	ServiceQuota_Kind             = "ServiceQuota"
	ServiceQuota_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceQuota_Kind}.String()
	ServiceQuota_KindAPIVersion   = ServiceQuota_Kind + "." + CRDGroupVersion.String()
	ServiceQuota_GroupVersionKind = CRDGroupVersion.WithKind(ServiceQuota_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceQuota{}, &ServiceQuotaList{})
}
