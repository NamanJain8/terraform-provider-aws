// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ivschat

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ivschat_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ivschat"
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
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLoggingConfiguration,
			TypeName: "aws_ivschat_logging_configuration",
			Name:     "Logging Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceRoom,
			TypeName: "aws_ivschat_room",
			Name:     "Room",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IVSChat
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context) (*ivschat_sdkv2.Client, error) {
	cfg := *(p.config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ivschat_sdkv2.NewFromConfig(cfg, func(o *ivschat_sdkv2.Options) {
		if endpoint := p.config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = ivschat_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
