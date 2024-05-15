// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package organizations

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	organizations_sdkv2 "github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDelegatedAdministrators,
			TypeName: "aws_organizations_delegated_administrators",
			Name:     "Delegated Administrators",
		},
		{
			Factory:  dataSourceDelegatedServices,
			TypeName: "aws_organizations_delegated_services",
			Name:     "Delegated Services",
		},
		{
			Factory:  dataSourceOrganization,
			TypeName: "aws_organizations_organization",
			Name:     "Organization",
		},
		{
			Factory:  DataSourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
			Name:     "Organizational Unit",
		},
		{
			Factory:  dataSourceOrganizationalUnitChildAccounts,
			TypeName: "aws_organizations_organizational_unit_child_accounts",
			Name:     "Organizational Unit Child Accounts",
		},
		{
			Factory:  DataSourceOrganizationalUnitDescendantAccounts,
			TypeName: "aws_organizations_organizational_unit_descendant_accounts",
		},
		{
			Factory:  DataSourceOrganizationalUnits,
			TypeName: "aws_organizations_organizational_units",
		},
		{
			Factory:  DataSourcePolicies,
			TypeName: "aws_organizations_policies",
		},
		{
			Factory:  DataSourcePoliciesForTarget,
			TypeName: "aws_organizations_policies_for_target",
		},
		{
			Factory:  DataSourcePolicy,
			TypeName: "aws_organizations_policy",
		},
		{
			Factory:  DataSourceResourceTags,
			TypeName: "aws_organizations_resource_tags",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccount,
			TypeName: "aws_organizations_account",
			Name:     "Account",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDelegatedAdministrator,
			TypeName: "aws_organizations_delegated_administrator",
			Name:     "Delegated Administrator",
		},
		{
			Factory:  resourceOrganization,
			TypeName: "aws_organizations_organization",
			Name:     "Organization",
		},
		{
			Factory:  ResourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
			Name:     "Organizational Unit",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourcePolicy,
			TypeName: "aws_organizations_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourcePolicyAttachment,
			TypeName: "aws_organizations_policy_attachment",
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_organizations_resource_policy",
			Name:     "Resource Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Organizations
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*organizations_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return organizations_sdkv2.NewFromConfig(cfg, func(o *organizations_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
