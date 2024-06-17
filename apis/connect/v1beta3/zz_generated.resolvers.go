// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta3

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *HoursOfOperation) ResolveReferences( // ResolveReferences of this HoursOfOperation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta3", "HoursOfOperation", "HoursOfOperationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HoursOfOperationID),
			Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
			Reference:    mg.Spec.ForProvider.HoursOfOperationIDRef,
			Selector:     mg.Spec.ForProvider.HoursOfOperationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HoursOfOperationID")
	}
	mg.Spec.ForProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HoursOfOperationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta3", "HoursOfOperation", "HoursOfOperationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HoursOfOperationID),
			Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
			Reference:    mg.Spec.InitProvider.HoursOfOperationIDRef,
			Selector:     mg.Spec.InitProvider.HoursOfOperationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HoursOfOperationID")
	}
	mg.Spec.InitProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HoursOfOperationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}
