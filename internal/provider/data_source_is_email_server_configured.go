package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*IsEmailServerConfiguredDataSource)(nil)

// IsEmailServerConfiguredDataSource is the generated Terraform data source implementation.
type IsEmailServerConfiguredDataSource struct {
}

// IsEmailServerConfiguredDataSourceModel describes the data source state shape.
type IsEmailServerConfiguredDataSourceModel struct {
}

// NewIsEmailServerConfiguredDataSource returns a new instance of the generated data source.
func NewIsEmailServerConfiguredDataSource() datasource.DataSource {
	return &IsEmailServerConfiguredDataSource{}
}

// Metadata returns the data source type name.
func (d *IsEmailServerConfiguredDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_is_email_server_configured"
}

// Schema returns the data source schema.
func (d *IsEmailServerConfiguredDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns whether email server has been configured, or not"}
}

// Read fetches remote state into the data source model.
func (d *IsEmailServerConfiguredDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IsEmailServerConfiguredDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
