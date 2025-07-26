// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"go.elastic.co/apm/module/apmhttp"
)

var _ provider.Provider = &ObservabilityProvider{}

type ObservabilityProvider struct {
	version string
}

type ObservabilityProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ObservabilityProvider{
			version: version,
		}
	}
}

func (o *ObservabilityProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "observability"
	resp.Version = o.version
}

func (o *ObservabilityProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The endpoint of the APM server (e.g., http://localhost:8200)",
				Optional:            true,
			},
		},
	}
}

func (o *ObservabilityProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ObservabilityProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Set APM server URL if provided
	if !data.Endpoint.IsNull() && !data.Endpoint.IsUnknown() {
		os.Setenv("ELASTIC_APM_SERVER_URL", data.Endpoint.ValueString())
	}

	client := apmhttp.WrapClient(http.DefaultClient)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (o *ObservabilityProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (o *ObservabilityProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
