package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadNodeSystemConfigTextFileDataSource)(nil)

// DownloadNodeSystemConfigTextFileDataSource is the generated Terraform data source implementation.
type DownloadNodeSystemConfigTextFileDataSource struct {
}

// DownloadNodeSystemConfigTextFileDataSourceModel describes the data source state shape.
type DownloadNodeSystemConfigTextFileDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Filename  types.String `tfsdk:"filename"`
}

// NewDownloadNodeSystemConfigTextFileDataSource returns a new instance of the generated data source.
func NewDownloadNodeSystemConfigTextFileDataSource() datasource.DataSource {
	return &DownloadNodeSystemConfigTextFileDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadNodeSystemConfigTextFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_node_system_config_text_file"
}

// Schema returns the data source schema.
func (d *DownloadNodeSystemConfigTextFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download a system text configuration file", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "filename of the text config file", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadNodeSystemConfigTextFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadNodeSystemConfigTextFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
