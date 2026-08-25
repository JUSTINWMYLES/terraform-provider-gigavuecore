package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GenerateRunningTextBackupConfigDataSource)(nil)

// GenerateRunningTextBackupConfigDataSource is the generated Terraform data source implementation.
type GenerateRunningTextBackupConfigDataSource struct {
}

// GenerateRunningTextBackupConfigDataSourceModel describes the data source state shape.
type GenerateRunningTextBackupConfigDataSourceModel struct {
	BackupId  types.String `tfsdk:"backup_id" json:"backupId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
}

// NewGenerateRunningTextBackupConfigDataSource returns a new instance of the generated data source.
func NewGenerateRunningTextBackupConfigDataSource() datasource.DataSource {
	return &GenerateRunningTextBackupConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *GenerateRunningTextBackupConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_generate_running_text_backup_config"
}

// Schema returns the data source schema.
func (d *GenerateRunningTextBackupConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Generate running text backup config", Attributes: map[string]schema.Attribute{"backup_id": schema.StringAttribute{MarkdownDescription: "id of the config backup snapshot to download", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "ID of the target device", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GenerateRunningTextBackupConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GenerateRunningTextBackupConfigDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
