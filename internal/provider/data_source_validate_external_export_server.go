package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*ValidateExternalExportServerDataSource)(nil)

// ValidateExternalExportServerDataSource is the generated Terraform data source implementation.
type ValidateExternalExportServerDataSource struct {
}

// ValidateExternalExportServerDataSourceModel describes the data source state shape.
type ValidateExternalExportServerDataSourceModel struct {
	ExportTargetAlias types.String `tfsdk:"export_target_alias" json:"exportTargetAlias"`
}

// NewValidateExternalExportServerDataSource returns a new instance of the generated data source.
func NewValidateExternalExportServerDataSource() datasource.DataSource {
	return &ValidateExternalExportServerDataSource{}
}

// Metadata returns the data source type name.
func (d *ValidateExternalExportServerDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_validate_external_export_server"
}

// Schema returns the data source schema.
func (d *ValidateExternalExportServerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Validate External Export Server", Attributes: map[string]schema.Attribute{"export_target_alias": schema.StringAttribute{MarkdownDescription: "Alias for external export target target server", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *ValidateExternalExportServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ValidateExternalExportServerDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
