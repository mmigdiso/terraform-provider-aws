// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package mediastore

import (
	"context"

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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceContainer,
			TypeName: "aws_media_store_container",
		},
		{
			Factory:  ResourceContainerPolicy,
			TypeName: "aws_media_store_container_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.MediaStore
}

var ServicePackage = &servicePackage{}
