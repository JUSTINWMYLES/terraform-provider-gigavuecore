package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*LoadSysdumpFileDataSource)(nil)

// LoadSysdumpFileDataSource is the generated Terraform data source implementation.
type LoadSysdumpFileDataSource struct {
}

// LoadSysdumpFileDataSourceModel describes the data source state shape.
type LoadSysdumpFileDataSourceModel struct {
	BoxId     types.String `tfsdk:"box_id" json:"boxId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Filename  types.String `tfsdk:"filename"`
}

// NewLoadSysdumpFileDataSource returns a new instance of the generated data source.
func NewLoadSysdumpFileDataSource() datasource.DataSource {
	return &LoadSysdumpFileDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSysdumpFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_sysdump_file"
}

// Schema returns the data source schema.
func (d *LoadSysdumpFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download Sysdump File", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Box ID range from 1 to 64(inclusive)", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "Sysdump filename to download", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSysdumpFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSysdumpFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
