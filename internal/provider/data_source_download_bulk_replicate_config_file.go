package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadBulkReplicateConfigFileDataSource)(nil)

// DownloadBulkReplicateConfigFileDataSource is the generated Terraform data source implementation.
type DownloadBulkReplicateConfigFileDataSource struct {
}

// DownloadBulkReplicateConfigFileDataSourceModel describes the data source state shape.
type DownloadBulkReplicateConfigFileDataSourceModel struct {
	Filename types.String `tfsdk:"filename"`
}

// NewDownloadBulkReplicateConfigFileDataSource returns a new instance of the generated data source.
func NewDownloadBulkReplicateConfigFileDataSource() datasource.DataSource {
	return &DownloadBulkReplicateConfigFileDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadBulkReplicateConfigFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_bulk_replicate_config_file"
}

// Schema returns the data source schema.
func (d *DownloadBulkReplicateConfigFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download bulk replicate config file", Attributes: map[string]schema.Attribute{"filename": schema.StringAttribute{MarkdownDescription: "Target config file name", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadBulkReplicateConfigFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadBulkReplicateConfigFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
