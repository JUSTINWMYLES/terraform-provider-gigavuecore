package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetSerialNumberOfCardDataSource)(nil)

// GetSerialNumberOfCardDataSource is the generated Terraform data source implementation.
type GetSerialNumberOfCardDataSource struct {
}

// GetSerialNumberOfCardDataSourceModel describes the data source state shape.
type GetSerialNumberOfCardDataSourceModel struct {
	BoxSlashSlotId types.String `tfsdk:"box_slash_slot_id" json:"boxSlashSlotId"`
	ClusterName    types.String `tfsdk:"cluster_name" json:"clusterName"`
	NodeId         types.String `tfsdk:"node_id" json:"nodeId"`
}

// NewGetSerialNumberOfCardDataSource returns a new instance of the generated data source.
func NewGetSerialNumberOfCardDataSource() datasource.DataSource {
	return &GetSerialNumberOfCardDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSerialNumberOfCardDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_serial_number_of_card"
}

// Schema returns the data source schema.
func (d *GetSerialNumberOfCardDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gets the serial numbers of card specified by input parameters", Attributes: map[string]schema.Attribute{"box_slash_slot_id": schema.StringAttribute{MarkdownDescription: "<box-id>/<slot-id>", Required: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "name of chassis cluster", Required: true}, "node_id": schema.StringAttribute{MarkdownDescription: "node ID, the IP address of the chassis", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSerialNumberOfCardDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSerialNumberOfCardDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
