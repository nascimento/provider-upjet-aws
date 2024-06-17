// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Addon.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Addon) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceAccountRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceAccountRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountRoleArn")
	}
	mg.Spec.ForProvider.ServiceAccountRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceAccountRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceAccountRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountRoleArn")
	}
	mg.Spec.InitProvider.ServiceAccountRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIds")
		}
		mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig[i3].SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this ClusterAuth.
func (mg *ClusterAuth) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: mg.Spec.ForProvider.ClusterName,
			Extract:      ExternalNameIfClusterActive(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = rsp.ResolvedValue
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FargateProfile.
func (mg *FargateProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PodExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PodExecutionRoleArnRef,
			Selector:     mg.Spec.ForProvider.PodExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PodExecutionRoleArn")
	}
	mg.Spec.ForProvider.PodExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PodExecutionRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PodExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.PodExecutionRoleArnRef,
			Selector:     mg.Spec.InitProvider.PodExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PodExecutionRoleArn")
	}
	mg.Spec.InitProvider.PodExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PodExecutionRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetIDRefs,
			Selector:      mg.Spec.InitProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this IdentityProviderConfig.
func (mg *IdentityProviderConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodeGroup.
func (mg *NodeGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      ExternalNameIfClusterActive(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.NodeRoleArnRef,
			Selector:     mg.Spec.ForProvider.NodeRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NodeRoleArn")
	}
	mg.Spec.ForProvider.NodeRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NodeRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RemoteAccess); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIDRefs,
				Selector:      mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIds")
		}
		mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.RemoteAccess[i3].SourceSecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Version),
			Extract:      resource.ExtractParamPath("version", false),
			Reference:    mg.Spec.ForProvider.VersionRef,
			Selector:     mg.Spec.ForProvider.VersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Version")
	}
	mg.Spec.ForProvider.Version = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VersionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NodeRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.NodeRoleArnRef,
			Selector:     mg.Spec.InitProvider.NodeRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NodeRoleArn")
	}
	mg.Spec.InitProvider.NodeRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NodeRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.RemoteAccess); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIDRefs,
				Selector:      mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIds")
		}
		mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.RemoteAccess[i3].SourceSecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetIDRefs,
			Selector:      mg.Spec.InitProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Version),
			Extract:      resource.ExtractParamPath("version", false),
			Reference:    mg.Spec.InitProvider.VersionRef,
			Selector:     mg.Spec.InitProvider.VersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Version")
	}
	mg.Spec.InitProvider.Version = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VersionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PodIdentityAssociation.
func (mg *PodIdentityAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eks.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
