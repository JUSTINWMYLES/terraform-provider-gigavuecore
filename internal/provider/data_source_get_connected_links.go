package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetConnectedLinksDataSource)(nil)

// GetConnectedLinksDataSource is the generated Terraform data source implementation.
type GetConnectedLinksDataSource struct {
}

// GetConnectedLinksDataSourceModel describes the data source state shape.
type GetConnectedLinksDataSourceModel struct {
}

// NewGetConnectedLinksDataSource returns a new instance of the generated data source.
func NewGetConnectedLinksDataSource() datasource.DataSource {
	return &GetConnectedLinksDataSource{}
}

// Metadata returns the data source type name.
func (d *GetConnectedLinksDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_connected_links"
}

// Schema returns the data source schema.
func (d *GetConnectedLinksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve connected links for troubleshooting flows"}
}

// Read fetches remote state into the data source model.
func (d *GetConnectedLinksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetConnectedLinksDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
