package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetConnectionStatusDataSource)(nil)

// GetConnectionStatusDataSource is the generated Terraform data source implementation.
type GetConnectionStatusDataSource struct {
}

// GetConnectionStatusDataSourceModel describes the data source state shape.
type GetConnectionStatusDataSourceModel struct {
	EnvId   types.String `tfsdk:"env_id" json:"envId"`
	UnifyId types.String `tfsdk:"unify_id" json:"unifyId"`
}

// NewGetConnectionStatusDataSource returns a new instance of the generated data source.
func NewGetConnectionStatusDataSource() datasource.DataSource {
	return &GetConnectionStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *GetConnectionStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_connection_status"
}

// Schema returns the data source schema.
func (d *GetConnectionStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Obtain the connection status for underlying unified environment and unified resource id", Attributes: map[string]schema.Attribute{"env_id": schema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "unify_id": schema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetConnectionStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetConnectionStatusDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
