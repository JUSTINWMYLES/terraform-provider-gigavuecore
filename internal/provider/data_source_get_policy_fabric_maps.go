package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetPolicyFabricMapsDataSource)(nil)

// GetPolicyFabricMapsDataSource is the generated Terraform data source implementation.
type GetPolicyFabricMapsDataSource struct {
}

// GetPolicyFabricMapsDataSourceModel describes the data source state shape.
type GetPolicyFabricMapsDataSourceModel struct {
	Name types.Dynamic `tfsdk:"name"`
}

// NewGetPolicyFabricMapsDataSource returns a new instance of the generated data source.
func NewGetPolicyFabricMapsDataSource() datasource.DataSource {
	return &GetPolicyFabricMapsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPolicyFabricMapsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_policy_fabric_maps"
}

// Schema returns the data source schema.
func (d *GetPolicyFabricMapsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Tools", Attributes: map[string]schema.Attribute{"name": schema.DynamicAttribute{MarkdownDescription: "policy name", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPolicyFabricMapsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPolicyFabricMapsDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
