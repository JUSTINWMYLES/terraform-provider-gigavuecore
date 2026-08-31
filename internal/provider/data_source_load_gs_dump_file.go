package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*LoadGsDumpFileDataSource)(nil)

// LoadGsDumpFileDataSource is the generated Terraform data source implementation.
type LoadGsDumpFileDataSource struct {
}

// LoadGsDumpFileDataSourceModel describes the data source state shape.
type LoadGsDumpFileDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Filename  types.String `tfsdk:"filename"`
	Hostname  types.String `tfsdk:"hostname"`
}

// NewLoadGsDumpFileDataSource returns a new instance of the generated data source.
func NewLoadGsDumpFileDataSource() datasource.DataSource {
	return &LoadGsDumpFileDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadGsDumpFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_gs_dump_file"
}

// Schema returns the data source schema.
func (d *LoadGsDumpFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download Gigasmart dump File", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster Id", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "Gigasmart dump filename to download", Required: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Target Host Name", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadGsDumpFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadGsDumpFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
