package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadNodeSystemConfigFileDataSource)(nil)

// DownloadNodeSystemConfigFileDataSource is the generated Terraform data source implementation.
type DownloadNodeSystemConfigFileDataSource struct {
}

// DownloadNodeSystemConfigFileDataSourceModel describes the data source state shape.
type DownloadNodeSystemConfigFileDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Filename  types.String `tfsdk:"filename"`
}

// NewDownloadNodeSystemConfigFileDataSource returns a new instance of the generated data source.
func NewDownloadNodeSystemConfigFileDataSource() datasource.DataSource {
	return &DownloadNodeSystemConfigFileDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadNodeSystemConfigFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_node_system_config_file"
}

// Schema returns the data source schema.
func (d *DownloadNodeSystemConfigFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download a system configuration file", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "filename of the target config file", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadNodeSystemConfigFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadNodeSystemConfigFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
