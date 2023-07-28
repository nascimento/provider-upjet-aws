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

type AuditLogConfigurationObservation struct {

	// The Amazon Resource Name (ARN) for the destination of the audit logs. The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN. Can be specified when file_access_audit_log_level and file_share_access_audit_log_level are not set to DISABLED. The name of the Amazon CloudWatch Logs log group must begin with the /aws/fsx prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the aws-fsx prefix. If you do not provide a destination in audit_log_destionation, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/windows log group.
	AuditLogDestination *string `json:"auditLogDestination,omitempty" tf:"audit_log_destination,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file and folder accesses. Valid values are SUCCESS_ONLY, FAILURE_ONLY, SUCCESS_AND_FAILURE, and DISABLED. Default value is DISABLED.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel,omitempty" tf:"file_access_audit_log_level,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file share accesses. Valid values are SUCCESS_ONLY, FAILURE_ONLY, SUCCESS_AND_FAILURE, and DISABLED. Default value is DISABLED.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel,omitempty" tf:"file_share_access_audit_log_level,omitempty"`
}

type AuditLogConfigurationParameters struct {

	// The Amazon Resource Name (ARN) for the destination of the audit logs. The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN. Can be specified when file_access_audit_log_level and file_share_access_audit_log_level are not set to DISABLED. The name of the Amazon CloudWatch Logs log group must begin with the /aws/fsx prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the aws-fsx prefix. If you do not provide a destination in audit_log_destionation, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/windows log group.
	// +kubebuilder:validation:Optional
	AuditLogDestination *string `json:"auditLogDestination,omitempty" tf:"audit_log_destination,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file and folder accesses. Valid values are SUCCESS_ONLY, FAILURE_ONLY, SUCCESS_AND_FAILURE, and DISABLED. Default value is DISABLED.
	// +kubebuilder:validation:Optional
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel,omitempty" tf:"file_access_audit_log_level,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file share accesses. Valid values are SUCCESS_ONLY, FAILURE_ONLY, SUCCESS_AND_FAILURE, and DISABLED. Default value is DISABLED.
	// +kubebuilder:validation:Optional
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel,omitempty" tf:"file_share_access_audit_log_level,omitempty"`
}

type SelfManagedActiveDirectoryObservation struct {

	// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in RFC 1918.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// The fully qualified domain name of the self-managed AD directory. For example, corp.example.com.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to Domain Admins.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group,omitempty"`

	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, OU=FSx,DC=yourdomain,DC=corp,DC=com. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see RFC 2253.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`

	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SelfManagedActiveDirectoryParameters struct {

	// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in RFC 1918.
	// +kubebuilder:validation:Required
	DNSIps []*string `json:"dnsIps" tf:"dns_ips,omitempty"`

	// The fully qualified domain name of the self-managed AD directory. For example, corp.example.com.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to Domain Admins.
	// +kubebuilder:validation:Optional
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group,omitempty"`

	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, OU=FSx,DC=yourdomain,DC=corp,DC=com. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see RFC 2253.
	// +kubebuilder:validation:Optional
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`

	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type WindowsFileSystemObservation struct {

	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with self_managed_active_directory.
	ActiveDirectoryID *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id,omitempty"`

	// An array DNS alias names that you want to associate with the Amazon FSx file system.  For more information, see Working with DNS Aliases
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// Amazon Resource Name of the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system. See below.
	AuditLogConfiguration []AuditLogConfigurationObservation `json:"auditLogConfiguration,omitempty" tf:"audit_log_configuration,omitempty"`

	// The number of days to retain automatic backups. Minimum of 0 and maximum of 90. Defaults to 7. Set to 0 to disable.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to false.
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// DNS name for the file system, e.g., fs-12345678.corp.example.com (domain name matching the Active Directory domain name)
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The preferred time (in HH:MM format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// Specifies the file system deployment type, valid values are MULTI_AZ_1, SINGLE_AZ_1 and SINGLE_AZ_2. Default value is SINGLE_AZ_1.
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// Identifier of the file system (e.g. fs-12345678).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// AWS account identifier that created the file system.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The IP address of the primary, or preferred, file server.
	PreferredFileServerIP *string `json:"preferredFileServerIp,omitempty" tf:"preferred_file_server_ip,omitempty"`

	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is MULTI_AZ_1.
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// For MULTI_AZ_1 deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For SINGLE_AZ_1 deployment types, this is the DNS name of the file system.
	RemoteAdministrationEndpoint *string `json:"remoteAdministrationEndpoint,omitempty" tf:"remote_administration_endpoint,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with active_directory_id. Detailed below.
	SelfManagedActiveDirectory []SelfManagedActiveDirectoryObservation `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to false.
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to HDD the minimum value is 2000. Required when not creating filesystem for a backup.
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// Specifies the storage type, Valid values are SSD and HDD. HDD is supported on SINGLE_AZ_2 and MULTI_AZ_1 Windows file system deployment types. Default value is SSD.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set deployment_type to MULTI_AZ_1.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of 8 and maximum of 2048.
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`

	// Identifier of the Virtual Private Cloud for the file system.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type WindowsFileSystemParameters struct {

	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with self_managed_active_directory.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ActiveDirectoryID *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id,omitempty"`

	// Reference to a Directory in ds to populate activeDirectoryId.
	// +kubebuilder:validation:Optional
	ActiveDirectoryIDRef *v1.Reference `json:"activeDirectoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate activeDirectoryId.
	// +kubebuilder:validation:Optional
	ActiveDirectoryIDSelector *v1.Selector `json:"activeDirectoryIdSelector,omitempty" tf:"-"`

	// An array DNS alias names that you want to associate with the Amazon FSx file system.  For more information, see Working with DNS Aliases
	// +kubebuilder:validation:Optional
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system. See below.
	// +kubebuilder:validation:Optional
	AuditLogConfiguration []AuditLogConfigurationParameters `json:"auditLogConfiguration,omitempty" tf:"audit_log_configuration,omitempty"`

	// The number of days to retain automatic backups. Minimum of 0 and maximum of 90. Defaults to 7. Set to 0 to disable.
	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// The ID of the source backup to create the filesystem from.
	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to false.
	// +kubebuilder:validation:Optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// The preferred time (in HH:MM format) to take daily automatic backups, in the UTC time zone.
	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// Specifies the file system deployment type, valid values are MULTI_AZ_1, SINGLE_AZ_1 and SINGLE_AZ_2. Default value is SINGLE_AZ_1.
	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is MULTI_AZ_1.
	// +kubebuilder:validation:Optional
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with active_directory_id. Detailed below.
	// +kubebuilder:validation:Optional
	SelfManagedActiveDirectory []SelfManagedActiveDirectoryParameters `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to false.
	// +kubebuilder:validation:Optional
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to HDD the minimum value is 2000. Required when not creating filesystem for a backup.
	// +kubebuilder:validation:Optional
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// Specifies the storage type, Valid values are SSD and HDD. HDD is supported on SINGLE_AZ_2 and MULTI_AZ_1 Windows file system deployment types. Default value is SSD.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set deployment_type to MULTI_AZ_1.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of 8 and maximum of 2048.
	// +kubebuilder:validation:Optional
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

// WindowsFileSystemSpec defines the desired state of WindowsFileSystem
type WindowsFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowsFileSystemParameters `json:"forProvider"`
}

// WindowsFileSystemStatus defines the observed state of WindowsFileSystem.
type WindowsFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsFileSystem is the Schema for the WindowsFileSystems API. Manages a FSx Windows File System.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.throughputCapacity)",message="throughputCapacity is a required parameter"
	Spec   WindowsFileSystemSpec   `json:"spec"`
	Status WindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsFileSystemList contains a list of WindowsFileSystems
type WindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WindowsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	WindowsFileSystem_Kind             = "WindowsFileSystem"
	WindowsFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WindowsFileSystem_Kind}.String()
	WindowsFileSystem_KindAPIVersion   = WindowsFileSystem_Kind + "." + CRDGroupVersion.String()
	WindowsFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(WindowsFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&WindowsFileSystem{}, &WindowsFileSystemList{})
}
