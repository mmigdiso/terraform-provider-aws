// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rum

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
			Factory:  ResourceAppMonitor,
			TypeName: "aws_rum_app_monitor",
		},
		{
			Factory:  ResourceMetricsDestination,
			TypeName: "aws_rum_metrics_destination",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RUM
}

var ServicePackage = &servicePackage{}
