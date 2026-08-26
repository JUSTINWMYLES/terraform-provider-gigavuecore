package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*LoadPcapFileDataSource)(nil)

// LoadPcapFileDataSource is the generated Terraform data source implementation.
type LoadPcapFileDataSource struct {
}

// LoadPcapFileDataSourceModel describes the data source state shape.
type LoadPcapFileDataSourceModel struct {
	BoxId     types.String `tfsdk:"box_id" json:"boxId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Filename  types.String `tfsdk:"filename"`
}

// NewLoadPcapFileDataSource returns a new instance of the generated data source.
func NewLoadPcapFileDataSource() datasource.DataSource {
	return &LoadPcapFileDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadPcapFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_pcap_file"
}

// Schema returns the data source schema.
func (d *LoadPcapFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download Pcap File", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Box ID range from 1 to 64(inclusive)", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "pcap filename to download", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadPcapFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadPcapFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
