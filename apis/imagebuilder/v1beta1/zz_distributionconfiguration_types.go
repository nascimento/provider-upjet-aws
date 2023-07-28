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

type AMIDistributionConfigurationObservation struct {

	// Key-value map of tags to apply to the distributed AMI.
	AMITags map[string]*string `json:"amiTags,omitempty" tf:"ami_tags,omitempty"`

	// Description to apply to the distributed AMI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key to encrypt the distributed AMI.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Configuration block of EC2 launch permissions to apply to the distributed AMI. Detailed below.
	LaunchPermission []LaunchPermissionObservation `json:"launchPermission,omitempty" tf:"launch_permission,omitempty"`

	// Name to apply to the distributed AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of AWS Account identifiers to distribute the AMI.
	TargetAccountIds []*string `json:"targetAccountIds,omitempty" tf:"target_account_ids,omitempty"`
}

type AMIDistributionConfigurationParameters struct {

	// Key-value map of tags to apply to the distributed AMI.
	// +kubebuilder:validation:Optional
	AMITags map[string]*string `json:"amiTags,omitempty" tf:"ami_tags,omitempty"`

	// Description to apply to the distributed AMI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key to encrypt the distributed AMI.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Configuration block of EC2 launch permissions to apply to the distributed AMI. Detailed below.
	// +kubebuilder:validation:Optional
	LaunchPermission []LaunchPermissionParameters `json:"launchPermission,omitempty" tf:"launch_permission,omitempty"`

	// Name to apply to the distributed AMI.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of AWS Account identifiers to distribute the AMI.
	// +kubebuilder:validation:Optional
	TargetAccountIds []*string `json:"targetAccountIds,omitempty" tf:"target_account_ids,omitempty"`
}

type ContainerDistributionConfigurationObservation struct {

	// Set of tags that are attached to the container distribution configuration.
	ContainerTags []*string `json:"containerTags,omitempty" tf:"container_tags,omitempty"`

	// Description of the container distribution configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block with the destination repository for the container distribution configuration.
	TargetRepository []ContainerDistributionConfigurationTargetRepositoryObservation `json:"targetRepository,omitempty" tf:"target_repository,omitempty"`
}

type ContainerDistributionConfigurationParameters struct {

	// Set of tags that are attached to the container distribution configuration.
	// +kubebuilder:validation:Optional
	ContainerTags []*string `json:"containerTags,omitempty" tf:"container_tags,omitempty"`

	// Description of the container distribution configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block with the destination repository for the container distribution configuration.
	// +kubebuilder:validation:Required
	TargetRepository []ContainerDistributionConfigurationTargetRepositoryParameters `json:"targetRepository" tf:"target_repository,omitempty"`
}

type ContainerDistributionConfigurationTargetRepositoryObservation struct {

	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// The service in which this image is registered. Valid values: ECR.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ContainerDistributionConfigurationTargetRepositoryParameters struct {

	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	// +kubebuilder:validation:Required
	RepositoryName *string `json:"repositoryName" tf:"repository_name,omitempty"`

	// The service in which this image is registered. Valid values: ECR.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

type DistributionConfigurationObservation struct {

	// Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date the distribution configuration was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Date the distribution configuration was updated.
	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	// Description of the distribution configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more configuration blocks with distribution settings. Detailed below.
	Distribution []DistributionObservation `json:"distribution,omitempty" tf:"distribution,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the distribution configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DistributionConfigurationParameters struct {

	// Description of the distribution configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more configuration blocks with distribution settings. Detailed below.
	// +kubebuilder:validation:Optional
	Distribution []DistributionParameters `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// Name of the distribution configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// AWS Region for the distribution.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DistributionObservation struct {

	// Configuration block with Amazon Machine Image (AMI) distribution settings. Detailed below.
	AMIDistributionConfiguration []AMIDistributionConfigurationObservation `json:"amiDistributionConfiguration,omitempty" tf:"ami_distribution_configuration,omitempty"`

	// Configuration block with container distribution settings. Detailed below.
	ContainerDistributionConfiguration []ContainerDistributionConfigurationObservation `json:"containerDistributionConfiguration,omitempty" tf:"container_distribution_configuration,omitempty"`

	// Set of Windows faster-launching configurations to use for AMI distribution. Detailed below.
	FastLaunchConfiguration []FastLaunchConfigurationObservation `json:"fastLaunchConfiguration,omitempty" tf:"fast_launch_configuration,omitempty"`

	// Set of launch template configuration settings that apply to image distribution. Detailed below.
	LaunchTemplateConfiguration []LaunchTemplateConfigurationObservation `json:"launchTemplateConfiguration,omitempty" tf:"launch_template_configuration,omitempty"`

	// Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
	LicenseConfigurationArns []*string `json:"licenseConfigurationArns,omitempty" tf:"license_configuration_arns,omitempty"`

	// AWS Region for the distribution.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DistributionParameters struct {

	// Configuration block with Amazon Machine Image (AMI) distribution settings. Detailed below.
	// +kubebuilder:validation:Optional
	AMIDistributionConfiguration []AMIDistributionConfigurationParameters `json:"amiDistributionConfiguration,omitempty" tf:"ami_distribution_configuration,omitempty"`

	// Configuration block with container distribution settings. Detailed below.
	// +kubebuilder:validation:Optional
	ContainerDistributionConfiguration []ContainerDistributionConfigurationParameters `json:"containerDistributionConfiguration,omitempty" tf:"container_distribution_configuration,omitempty"`

	// Set of Windows faster-launching configurations to use for AMI distribution. Detailed below.
	// +kubebuilder:validation:Optional
	FastLaunchConfiguration []FastLaunchConfigurationParameters `json:"fastLaunchConfiguration,omitempty" tf:"fast_launch_configuration,omitempty"`

	// Set of launch template configuration settings that apply to image distribution. Detailed below.
	// +kubebuilder:validation:Optional
	LaunchTemplateConfiguration []LaunchTemplateConfigurationParameters `json:"launchTemplateConfiguration,omitempty" tf:"launch_template_configuration,omitempty"`

	// Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
	// +kubebuilder:validation:Optional
	LicenseConfigurationArns []*string `json:"licenseConfigurationArns,omitempty" tf:"license_configuration_arns,omitempty"`

	// AWS Region for the distribution.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type FastLaunchConfigurationObservation struct {

	// The owner account ID for the fast-launch enabled Windows AMI.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Configuration block for the launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots. Detailed below.
	LaunchTemplate []LaunchTemplateObservation `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// The maximum number of parallel instances that are launched for creating resources.
	MaxParallelLaunches *float64 `json:"maxParallelLaunches,omitempty" tf:"max_parallel_launches,omitempty"`

	// Configuration block for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled. Detailed below.
	SnapshotConfiguration []SnapshotConfigurationObservation `json:"snapshotConfiguration,omitempty" tf:"snapshot_configuration,omitempty"`
}

type FastLaunchConfigurationParameters struct {

	// The owner account ID for the fast-launch enabled Windows AMI.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Configuration block for the launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots. Detailed below.
	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// The maximum number of parallel instances that are launched for creating resources.
	// +kubebuilder:validation:Optional
	MaxParallelLaunches *float64 `json:"maxParallelLaunches,omitempty" tf:"max_parallel_launches,omitempty"`

	// Configuration block for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled. Detailed below.
	// +kubebuilder:validation:Optional
	SnapshotConfiguration []SnapshotConfigurationParameters `json:"snapshotConfiguration,omitempty" tf:"snapshot_configuration,omitempty"`
}

type LaunchPermissionObservation struct {

	// Set of AWS Organization ARNs to assign.
	OrganizationArns []*string `json:"organizationArns,omitempty" tf:"organization_arns,omitempty"`

	// Set of AWS Organizational Unit ARNs to assign.
	OrganizationalUnitArns []*string `json:"organizationalUnitArns,omitempty" tf:"organizational_unit_arns,omitempty"`

	// Set of EC2 launch permission user groups to assign. Use all to distribute a public AMI.
	UserGroups []*string `json:"userGroups,omitempty" tf:"user_groups,omitempty"`

	// Set of AWS Account identifiers to assign.
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`
}

type LaunchPermissionParameters struct {

	// Set of AWS Organization ARNs to assign.
	// +kubebuilder:validation:Optional
	OrganizationArns []*string `json:"organizationArns,omitempty" tf:"organization_arns,omitempty"`

	// Set of AWS Organizational Unit ARNs to assign.
	// +kubebuilder:validation:Optional
	OrganizationalUnitArns []*string `json:"organizationalUnitArns,omitempty" tf:"organizational_unit_arns,omitempty"`

	// Set of EC2 launch permission user groups to assign. Use all to distribute a public AMI.
	// +kubebuilder:validation:Optional
	UserGroups []*string `json:"userGroups,omitempty" tf:"user_groups,omitempty"`

	// Set of AWS Account identifiers to assign.
	// +kubebuilder:validation:Optional
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`
}

type LaunchTemplateConfigurationObservation struct {

	// The account ID that this configuration applies to.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether to set the specified Amazon EC2 launch template as the default launch template. Defaults to true.
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	// The ID of the Amazon EC2 launch template to use.
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`
}

type LaunchTemplateConfigurationParameters struct {

	// The account ID that this configuration applies to.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether to set the specified Amazon EC2 launch template as the default launch template. Defaults to true.
	// +kubebuilder:validation:Optional
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	// The ID of the Amazon EC2 launch template to use.
	// +kubebuilder:validation:Required
	LaunchTemplateID *string `json:"launchTemplateId" tf:"launch_template_id,omitempty"`
}

type LaunchTemplateObservation struct {

	// The ID of the launch template to use for faster launching for a Windows AMI.
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// The name of the launch template to use for faster launching for a Windows AMI.
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// The version of the launch template to use for faster launching for a Windows AMI.
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version,omitempty"`
}

type LaunchTemplateParameters struct {

	// The ID of the launch template to use for faster launching for a Windows AMI.
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// The name of the launch template to use for faster launching for a Windows AMI.
	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// The version of the launch template to use for faster launching for a Windows AMI.
	// +kubebuilder:validation:Optional
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version,omitempty"`
}

type SnapshotConfigurationObservation struct {

	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	TargetResourceCount *float64 `json:"targetResourceCount,omitempty" tf:"target_resource_count,omitempty"`
}

type SnapshotConfigurationParameters struct {

	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	// +kubebuilder:validation:Optional
	TargetResourceCount *float64 `json:"targetResourceCount,omitempty" tf:"target_resource_count,omitempty"`
}

// DistributionConfigurationSpec defines the desired state of DistributionConfiguration
type DistributionConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DistributionConfigurationParameters `json:"forProvider"`
}

// DistributionConfigurationStatus defines the observed state of DistributionConfiguration.
type DistributionConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DistributionConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DistributionConfiguration is the Schema for the DistributionConfigurations API. Manage an Image Builder Distribution Configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DistributionConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.distribution)",message="distribution is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   DistributionConfigurationSpec   `json:"spec"`
	Status DistributionConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributionConfigurationList contains a list of DistributionConfigurations
type DistributionConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributionConfiguration `json:"items"`
}

// Repository type metadata.
var (
	DistributionConfiguration_Kind             = "DistributionConfiguration"
	DistributionConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DistributionConfiguration_Kind}.String()
	DistributionConfiguration_KindAPIVersion   = DistributionConfiguration_Kind + "." + CRDGroupVersion.String()
	DistributionConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(DistributionConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&DistributionConfiguration{}, &DistributionConfigurationList{})
}
