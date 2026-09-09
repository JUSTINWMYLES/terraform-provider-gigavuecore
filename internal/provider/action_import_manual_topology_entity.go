package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ImportManualTopologyEntityAction)(nil)

// ImportManualTopologyEntityAction is the generated Terraform action implementation.
type ImportManualTopologyEntityAction struct {
}

// ImportManualTopologyEntityActionModel describes the action configuration shape.
type ImportManualTopologyEntityActionModel struct {
	Input types.String `tfsdk:"input"`
	Type  types.String `tfsdk:"type"`
}

// NewImportManualTopologyEntityAction returns a new instance of the generated action.
func NewImportManualTopologyEntityAction() action.Action {
	return &ImportManualTopologyEntityAction{}
}

// Metadata returns the action type name.
func (r *ImportManualTopologyEntityAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_manual_topology_entity"
}

// Schema returns the action schema.
func (r *ImportManualTopologyEntityAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Imports topology nodes and links from the information provided in csv file. For importing network devices \"Tool Name\", \"Vendor\", \"Type\", \"Model\", \"Comment\" and for tools the additional properties like \"Max Throughput(Gbps)\",\"Storage Capacity(TB)\", \"Compression Ratio\" column headers are required. Similarly for links \"Source Device\"(device alias/hostname), \"Source Cluster Name\", \"Source Device Type\"(gigamon/manual), \"Source EndPoint Type\"(Port/Gigastream), \"Source EndPoint Alias\"(port/gigastream alias), \"Destination Device\", \"Destination Cluster Name\", \"Destination Device Type\", \"Destination EndPoint Alias\", \"Destination EndPoint Type\", \"Link Type\"(port/gigaStream) are required in the csv file.", Attributes: map[string]schema.Attribute{"input": schema.StringAttribute{MarkdownDescription: "User uploaded file, only csv format is supported", Required: true}, "type": schema.StringAttribute{MarkdownDescription: "Topology Entity Type", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *ImportManualTopologyEntityAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportManualTopologyEntityActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
