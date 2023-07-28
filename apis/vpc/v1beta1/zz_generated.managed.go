/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkPerformanceMetricSubscription.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkPerformanceMetricSubscription) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkPerformanceMetricSubscription.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkPerformanceMetricSubscription) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this NetworkPerformanceMetricSubscription.
func (mg *NetworkPerformanceMetricSubscription) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
