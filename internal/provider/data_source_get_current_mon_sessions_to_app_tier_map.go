package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetCurrentMonSessionsToAppTierMapDataSource)(nil)

// GetCurrentMonSessionsToAppTierMapDataSource is the generated Terraform data source implementation.
type GetCurrentMonSessionsToAppTierMapDataSource struct {
}

// GetCurrentMonSessionsToAppTierMapDataSourceModel describes the data source state shape.
type GetCurrentMonSessionsToAppTierMapDataSourceModel struct {
}

// NewGetCurrentMonSessionsToAppTierMapDataSource returns a new instance of the generated data source.
func NewGetCurrentMonSessionsToAppTierMapDataSource() datasource.DataSource {
	return &GetCurrentMonSessionsToAppTierMapDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCurrentMonSessionsToAppTierMapDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_current_mon_sessions_to_app_tier_map"
}

// Schema returns the data source schema.
func (d *GetCurrentMonSessionsToAppTierMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives the mapping from current monitoring session id's to corresponding highest app tier as determined from the apps in the monitoring session"}
}

// Read fetches remote state into the data source model.
func (d *GetCurrentMonSessionsToAppTierMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCurrentMonSessionsToAppTierMapDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
