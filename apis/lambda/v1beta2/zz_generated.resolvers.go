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
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Alias.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Alias) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.FunctionNameRef,
			Selector:     mg.Spec.ForProvider.FunctionNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CodeSigningConfig.
func (mg *CodeSigningConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.AllowedPublishers != nil {
		{
			m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta2", "SigningProfile", "SigningProfileList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArns),
				Extract:       common.ARNExtractor(),
				References:    mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArnsRefs,
				Selector:      mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArnsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArns")
		}
		mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArns = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.AllowedPublishers.SigningProfileVersionArnsRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.AllowedPublishers != nil {
		{
			m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta2", "SigningProfile", "SigningProfileList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArns),
				Extract:       common.ARNExtractor(),
				References:    mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArnsRefs,
				Selector:      mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArnsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArns")
		}
		mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArns = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.AllowedPublishers.SigningProfileVersionArnsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this EventSourceMapping.
func (mg *EventSourceMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.FunctionNameRef,
			Selector:     mg.Spec.ForProvider.FunctionNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FunctionName),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.FunctionNameRef,
			Selector:     mg.Spec.InitProvider.FunctionNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FunctionName")
	}
	mg.Spec.InitProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Function.
func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.FileSystemConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("efs.aws.upbound.io", "v1beta2", "AccessPoint", "AccessPointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemConfig.Arn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.FileSystemConfig.ArnRef,
				Selector:     mg.Spec.ForProvider.FileSystemConfig.ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemConfig.Arn")
		}
		mg.Spec.ForProvider.FileSystemConfig.Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.FileSystemConfig.ArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ReplacementSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.ReplacementSecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.ReplacementSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplacementSecurityGroupIds")
	}
	mg.Spec.ForProvider.ReplacementSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ReplacementSecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.S3BucketRef,
			Selector:     mg.Spec.ForProvider.S3BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3Bucket")
	}
	mg.Spec.ForProvider.S3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.VPCConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig.SecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig.SecurityGroupIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig.SecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig.SecurityGroupIds")
		}
		mg.Spec.ForProvider.VPCConfig.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig.SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.VPCConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig.SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig.SubnetIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig.SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig.SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfig.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig.SubnetIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.FileSystemConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("efs.aws.upbound.io", "v1beta2", "AccessPoint", "AccessPointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemConfig.Arn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.FileSystemConfig.ArnRef,
				Selector:     mg.Spec.InitProvider.FileSystemConfig.ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemConfig.Arn")
		}
		mg.Spec.InitProvider.FileSystemConfig.Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.FileSystemConfig.ArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.ReplacementSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.ReplacementSecurityGroupIDRefs,
			Selector:      mg.Spec.InitProvider.ReplacementSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ReplacementSecurityGroupIds")
	}
	mg.Spec.InitProvider.ReplacementSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ReplacementSecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.S3Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.S3BucketRef,
			Selector:     mg.Spec.InitProvider.S3BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.S3Bucket")
	}
	mg.Spec.InitProvider.S3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.S3BucketRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.VPCConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig.SecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig.SecurityGroupIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig.SecurityGroupIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig.SecurityGroupIds")
		}
		mg.Spec.InitProvider.VPCConfig.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig.SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.VPCConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig.SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig.SubnetIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig.SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig.SubnetIds")
		}
		mg.Spec.InitProvider.VPCConfig.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig.SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this FunctionEventInvokeConfig.
func (mg *FunctionEventInvokeConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.DestinationConfig != nil {
		if mg.Spec.ForProvider.DestinationConfig.OnFailure != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationConfig.OnFailure.Destination),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.DestinationConfig.OnFailure.DestinationRef,
					Selector:     mg.Spec.ForProvider.DestinationConfig.OnFailure.DestinationSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DestinationConfig.OnFailure.Destination")
			}
			mg.Spec.ForProvider.DestinationConfig.OnFailure.Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DestinationConfig.OnFailure.DestinationRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.DestinationConfig != nil {
		if mg.Spec.ForProvider.DestinationConfig.OnSuccess != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationConfig.OnSuccess.Destination),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.DestinationConfig.OnSuccess.DestinationRef,
					Selector:     mg.Spec.ForProvider.DestinationConfig.OnSuccess.DestinationSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DestinationConfig.OnSuccess.Destination")
			}
			mg.Spec.ForProvider.DestinationConfig.OnSuccess.Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DestinationConfig.OnSuccess.DestinationRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.DestinationConfig != nil {
		if mg.Spec.InitProvider.DestinationConfig.OnFailure != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationConfig.OnFailure.Destination),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.DestinationConfig.OnFailure.DestinationRef,
					Selector:     mg.Spec.InitProvider.DestinationConfig.OnFailure.DestinationSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.DestinationConfig.OnFailure.Destination")
			}
			mg.Spec.InitProvider.DestinationConfig.OnFailure.Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.DestinationConfig.OnFailure.DestinationRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.DestinationConfig != nil {
		if mg.Spec.InitProvider.DestinationConfig.OnSuccess != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationConfig.OnSuccess.Destination),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.DestinationConfig.OnSuccess.DestinationRef,
					Selector:     mg.Spec.InitProvider.DestinationConfig.OnSuccess.DestinationSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.DestinationConfig.OnSuccess.Destination")
			}
			mg.Spec.InitProvider.DestinationConfig.OnSuccess.Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.DestinationConfig.OnSuccess.DestinationRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this FunctionURL.
func (mg *FunctionURL) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.FunctionNameRef,
			Selector:     mg.Spec.ForProvider.FunctionNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FunctionName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.FunctionNameRef,
			Selector:     mg.Spec.InitProvider.FunctionNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FunctionName")
	}
	mg.Spec.InitProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}
