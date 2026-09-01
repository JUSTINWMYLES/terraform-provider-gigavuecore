package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadBackupConfigFileOfFormatTextDataSource)(nil)

// DownloadBackupConfigFileOfFormatTextDataSource is the generated Terraform data source implementation.
type DownloadBackupConfigFileOfFormatTextDataSource struct {
}

// DownloadBackupConfigFileOfFormatTextDataSourceModel describes the data source state shape.
type DownloadBackupConfigFileOfFormatTextDataSourceModel struct {
	BackupId  types.String `tfsdk:"backup_id" json:"backupId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
}

// NewDownloadBackupConfigFileOfFormatTextDataSource returns a new instance of the generated data source.
func NewDownloadBackupConfigFileOfFormatTextDataSource() datasource.DataSource {
	return &DownloadBackupConfigFileOfFormatTextDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadBackupConfigFileOfFormatTextDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_backup_config_file_of_format_text"
}

// Schema returns the data source schema.
func (d *DownloadBackupConfigFileOfFormatTextDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download a Backup config file of format text", Attributes: map[string]schema.Attribute{"backup_id": schema.StringAttribute{MarkdownDescription: "id of the config backup snapshot to download", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "ID of the target device", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadBackupConfigFileOfFormatTextDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadBackupConfigFileOfFormatTextDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
