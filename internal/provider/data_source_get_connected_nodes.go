package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetConnectedNodesDataSource)(nil)

// GetConnectedNodesDataSource is the generated Terraform data source implementation.
type GetConnectedNodesDataSource struct {
}

// GetConnectedNodesDataSourceModel describes the data source state shape.
type GetConnectedNodesDataSourceModel struct {
	ClusterName types.String `tfsdk:"cluster_name" json:"clusterName"`
}

// NewGetConnectedNodesDataSource returns a new instance of the generated data source.
func NewGetConnectedNodesDataSource() datasource.DataSource {
	return &GetConnectedNodesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetConnectedNodesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_connected_nodes"
}

// Schema returns the data source schema.
func (d *GetConnectedNodesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve connected nodes for troubleshooting flows", Attributes: map[string]schema.Attribute{"cluster_name": schema.StringAttribute{MarkdownDescription: "Target Cluster Name", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetConnectedNodesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetConnectedNodesDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
