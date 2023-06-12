// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package healthlake

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	healthlake_sdkv2 "github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

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
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.HealthLake
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context) (*healthlake_sdkv2.Client, error) {
	cfg := *(p.config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return healthlake_sdkv2.NewFromConfig(cfg, func(o *healthlake_sdkv2.Options) {
		if endpoint := p.config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = healthlake_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
