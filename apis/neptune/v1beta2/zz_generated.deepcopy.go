//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInitParameters) DeepCopyInto(out *ClusterInitParameters) {
	*out = *in
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableCloudwatchLogsExports != nil {
		in, out := &in.EnableCloudwatchLogsExports, &out.EnableCloudwatchLogsExports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalClusterIdentifier != nil {
		in, out := &in.GlobalClusterIdentifier, &out.GlobalClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.IAMDatabaseAuthenticationEnabled != nil {
		in, out := &in.IAMDatabaseAuthenticationEnabled, &out.IAMDatabaseAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IAMRoleRefs != nil {
		in, out := &in.IAMRoleRefs, &out.IAMRoleRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IAMRoleSelector != nil {
		in, out := &in.IAMRoleSelector, &out.IAMRoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoles != nil {
		in, out := &in.IAMRoles, &out.IAMRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KMSKeyArn != nil {
		in, out := &in.KMSKeyArn, &out.KMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArnRef != nil {
		in, out := &in.KMSKeyArnRef, &out.KMSKeyArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyArnSelector != nil {
		in, out := &in.KMSKeyArnSelector, &out.KMSKeyArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneClusterParameterGroupName != nil {
		in, out := &in.NeptuneClusterParameterGroupName, &out.NeptuneClusterParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneClusterParameterGroupNameRef != nil {
		in, out := &in.NeptuneClusterParameterGroupNameRef, &out.NeptuneClusterParameterGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneClusterParameterGroupNameSelector != nil {
		in, out := &in.NeptuneClusterParameterGroupNameSelector, &out.NeptuneClusterParameterGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneInstanceParameterGroupName != nil {
		in, out := &in.NeptuneInstanceParameterGroupName, &out.NeptuneInstanceParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneSubnetGroupName != nil {
		in, out := &in.NeptuneSubnetGroupName, &out.NeptuneSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneSubnetGroupNameRef != nil {
		in, out := &in.NeptuneSubnetGroupNameRef, &out.NeptuneSubnetGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneSubnetGroupNameSelector != nil {
		in, out := &in.NeptuneSubnetGroupNameSelector, &out.NeptuneSubnetGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifier != nil {
		in, out := &in.ReplicationSourceIdentifier, &out.ReplicationSourceIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifierRef != nil {
		in, out := &in.ReplicationSourceIdentifierRef, &out.ReplicationSourceIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicationSourceIdentifierSelector != nil {
		in, out := &in.ReplicationSourceIdentifierSelector, &out.ReplicationSourceIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerlessV2ScalingConfiguration != nil {
		in, out := &in.ServerlessV2ScalingConfiguration, &out.ServerlessV2ScalingConfiguration
		*out = new(ServerlessV2ScalingConfigurationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipFinalSnapshot != nil {
		in, out := &in.SkipFinalSnapshot, &out.SkipFinalSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotIdentifier != nil {
		in, out := &in.SnapshotIdentifier, &out.SnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.SnapshotIdentifierRef != nil {
		in, out := &in.SnapshotIdentifierRef, &out.SnapshotIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SnapshotIdentifierSelector != nil {
		in, out := &in.SnapshotIdentifierSelector, &out.SnapshotIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCSecurityGroupIDRefs != nil {
		in, out := &in.VPCSecurityGroupIDRefs, &out.VPCSecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VPCSecurityGroupIDSelector != nil {
		in, out := &in.VPCSecurityGroupIDSelector, &out.VPCSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCSecurityGroupIds != nil {
		in, out := &in.VPCSecurityGroupIds, &out.VPCSecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInitParameters.
func (in *ClusterInitParameters) DeepCopy() *ClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ClusterMembers != nil {
		in, out := &in.ClusterMembers, &out.ClusterMembers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ClusterResourceID != nil {
		in, out := &in.ClusterResourceID, &out.ClusterResourceID
		*out = new(string)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableCloudwatchLogsExports != nil {
		in, out := &in.EnableCloudwatchLogsExports, &out.EnableCloudwatchLogsExports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalClusterIdentifier != nil {
		in, out := &in.GlobalClusterIdentifier, &out.GlobalClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.HostedZoneID != nil {
		in, out := &in.HostedZoneID, &out.HostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.IAMDatabaseAuthenticationEnabled != nil {
		in, out := &in.IAMDatabaseAuthenticationEnabled, &out.IAMDatabaseAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IAMRoles != nil {
		in, out := &in.IAMRoles, &out.IAMRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArn != nil {
		in, out := &in.KMSKeyArn, &out.KMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.NeptuneClusterParameterGroupName != nil {
		in, out := &in.NeptuneClusterParameterGroupName, &out.NeptuneClusterParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneInstanceParameterGroupName != nil {
		in, out := &in.NeptuneInstanceParameterGroupName, &out.NeptuneInstanceParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneSubnetGroupName != nil {
		in, out := &in.NeptuneSubnetGroupName, &out.NeptuneSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ReaderEndpoint != nil {
		in, out := &in.ReaderEndpoint, &out.ReaderEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifier != nil {
		in, out := &in.ReplicationSourceIdentifier, &out.ReplicationSourceIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ServerlessV2ScalingConfiguration != nil {
		in, out := &in.ServerlessV2ScalingConfiguration, &out.ServerlessV2ScalingConfiguration
		*out = new(ServerlessV2ScalingConfigurationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipFinalSnapshot != nil {
		in, out := &in.SkipFinalSnapshot, &out.SkipFinalSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotIdentifier != nil {
		in, out := &in.SnapshotIdentifier, &out.SnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCSecurityGroupIds != nil {
		in, out := &in.VPCSecurityGroupIds, &out.VPCSecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableCloudwatchLogsExports != nil {
		in, out := &in.EnableCloudwatchLogsExports, &out.EnableCloudwatchLogsExports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalClusterIdentifier != nil {
		in, out := &in.GlobalClusterIdentifier, &out.GlobalClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.IAMDatabaseAuthenticationEnabled != nil {
		in, out := &in.IAMDatabaseAuthenticationEnabled, &out.IAMDatabaseAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IAMRoleRefs != nil {
		in, out := &in.IAMRoleRefs, &out.IAMRoleRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IAMRoleSelector != nil {
		in, out := &in.IAMRoleSelector, &out.IAMRoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoles != nil {
		in, out := &in.IAMRoles, &out.IAMRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KMSKeyArn != nil {
		in, out := &in.KMSKeyArn, &out.KMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArnRef != nil {
		in, out := &in.KMSKeyArnRef, &out.KMSKeyArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyArnSelector != nil {
		in, out := &in.KMSKeyArnSelector, &out.KMSKeyArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneClusterParameterGroupName != nil {
		in, out := &in.NeptuneClusterParameterGroupName, &out.NeptuneClusterParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneClusterParameterGroupNameRef != nil {
		in, out := &in.NeptuneClusterParameterGroupNameRef, &out.NeptuneClusterParameterGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneClusterParameterGroupNameSelector != nil {
		in, out := &in.NeptuneClusterParameterGroupNameSelector, &out.NeptuneClusterParameterGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneInstanceParameterGroupName != nil {
		in, out := &in.NeptuneInstanceParameterGroupName, &out.NeptuneInstanceParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneSubnetGroupName != nil {
		in, out := &in.NeptuneSubnetGroupName, &out.NeptuneSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.NeptuneSubnetGroupNameRef != nil {
		in, out := &in.NeptuneSubnetGroupNameRef, &out.NeptuneSubnetGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NeptuneSubnetGroupNameSelector != nil {
		in, out := &in.NeptuneSubnetGroupNameSelector, &out.NeptuneSubnetGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifier != nil {
		in, out := &in.ReplicationSourceIdentifier, &out.ReplicationSourceIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifierRef != nil {
		in, out := &in.ReplicationSourceIdentifierRef, &out.ReplicationSourceIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicationSourceIdentifierSelector != nil {
		in, out := &in.ReplicationSourceIdentifierSelector, &out.ReplicationSourceIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerlessV2ScalingConfiguration != nil {
		in, out := &in.ServerlessV2ScalingConfiguration, &out.ServerlessV2ScalingConfiguration
		*out = new(ServerlessV2ScalingConfigurationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipFinalSnapshot != nil {
		in, out := &in.SkipFinalSnapshot, &out.SkipFinalSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotIdentifier != nil {
		in, out := &in.SnapshotIdentifier, &out.SnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.SnapshotIdentifierRef != nil {
		in, out := &in.SnapshotIdentifierRef, &out.SnapshotIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SnapshotIdentifierSelector != nil {
		in, out := &in.SnapshotIdentifierSelector, &out.SnapshotIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCSecurityGroupIDRefs != nil {
		in, out := &in.VPCSecurityGroupIDRefs, &out.VPCSecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VPCSecurityGroupIDSelector != nil {
		in, out := &in.VPCSecurityGroupIDSelector, &out.VPCSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCSecurityGroupIds != nil {
		in, out := &in.VPCSecurityGroupIds, &out.VPCSecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessV2ScalingConfigurationInitParameters) DeepCopyInto(out *ServerlessV2ScalingConfigurationInitParameters) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(float64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessV2ScalingConfigurationInitParameters.
func (in *ServerlessV2ScalingConfigurationInitParameters) DeepCopy() *ServerlessV2ScalingConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerlessV2ScalingConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessV2ScalingConfigurationObservation) DeepCopyInto(out *ServerlessV2ScalingConfigurationObservation) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(float64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessV2ScalingConfigurationObservation.
func (in *ServerlessV2ScalingConfigurationObservation) DeepCopy() *ServerlessV2ScalingConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ServerlessV2ScalingConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessV2ScalingConfigurationParameters) DeepCopyInto(out *ServerlessV2ScalingConfigurationParameters) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(float64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessV2ScalingConfigurationParameters.
func (in *ServerlessV2ScalingConfigurationParameters) DeepCopy() *ServerlessV2ScalingConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ServerlessV2ScalingConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
