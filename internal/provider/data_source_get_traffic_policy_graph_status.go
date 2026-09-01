package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetTrafficPolicyGraphStatusDataSource)(nil)

// GetTrafficPolicyGraphStatusDataSource is the generated Terraform data source implementation.
type GetTrafficPolicyGraphStatusDataSource struct {
}

// GetTrafficPolicyGraphStatusDataSourceModel describes the data source state shape.
type GetTrafficPolicyGraphStatusDataSourceModel struct {
	Alias types.String `tfsdk:"alias"`
}

// NewGetTrafficPolicyGraphStatusDataSource returns a new instance of the generated data source.
func NewGetTrafficPolicyGraphStatusDataSource() datasource.DataSource {
	return &GetTrafficPolicyGraphStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTrafficPolicyGraphStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_traffic_policy_graph_status"
}

// Schema returns the data source schema.
func (d *GetTrafficPolicyGraphStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get the traffic policy graph status", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Traffic Policy Graph alias", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetTrafficPolicyGraphStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTrafficPolicyGraphStatusDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
