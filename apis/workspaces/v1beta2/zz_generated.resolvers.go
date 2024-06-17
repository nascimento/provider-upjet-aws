// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Directory) ResolveReferences( // ResolveReferences of this Directory.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta2", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DirectoryIDRef,
			Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference
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

	if mg.Spec.ForProvider.WorkspaceCreationProperties != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupIDRef,
				Selector:     mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupID")
		}
		mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.WorkspaceCreationProperties.CustomSecurityGroupIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta2", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DirectoryIDRef,
			Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference
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

	if mg.Spec.InitProvider.WorkspaceCreationProperties != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupIDRef,
				Selector:     mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupID")
		}
		mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.WorkspaceCreationProperties.CustomSecurityGroupIDRef = rsp.ResolvedReference

	}

	return nil
}
