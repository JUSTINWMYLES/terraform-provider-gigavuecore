package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetHighestVblBundleDataSource)(nil)

// GetHighestVblBundleDataSource is the generated Terraform data source implementation.
type GetHighestVblBundleDataSource struct {
}

// GetHighestVblBundleDataSourceModel describes the data source state shape.
type GetHighestVblBundleDataSourceModel struct {
}

// NewGetHighestVblBundleDataSource returns a new instance of the generated data source.
func NewGetHighestVblBundleDataSource() datasource.DataSource {
	return &GetHighestVblBundleDataSource{}
}

// Metadata returns the data source type name.
func (d *GetHighestVblBundleDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_highest_vbl_bundle"
}

// Schema returns the data source schema.
func (d *GetHighestVblBundleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns the currently installed highest VBL bundle family that is not expired or deactivated (CoreVUE < NetVUE < SecureVUEPlus), or 'Unbundled' if no VBL license exists"}
}

// Read fetches remote state into the data source model.
func (d *GetHighestVblBundleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetHighestVblBundleDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
