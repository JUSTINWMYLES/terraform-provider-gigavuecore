package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadConfigFileDataSource)(nil)

// DownloadConfigFileDataSource is the generated Terraform data source implementation.
type DownloadConfigFileDataSource struct {
}

// DownloadConfigFileDataSourceModel describes the data source state shape.
type DownloadConfigFileDataSourceModel struct {
	BackupId  types.String `tfsdk:"backup_id" json:"backupId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
}

// NewDownloadConfigFileDataSource returns a new instance of the generated data source.
func NewDownloadConfigFileDataSource() datasource.DataSource {
	return &DownloadConfigFileDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadConfigFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_config_file"
}

// Schema returns the data source schema.
func (d *DownloadConfigFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download a config file", Attributes: map[string]schema.Attribute{"backup_id": schema.StringAttribute{MarkdownDescription: "id of the config backup snapshot to download", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "ID of the target device", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadConfigFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadConfigFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
