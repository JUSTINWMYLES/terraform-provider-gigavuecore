package provider

import "context"
import (
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*HeaderStripAgingResource)(nil)
	_ resource.ResourceWithImportState = (*HeaderStripAgingResource)(nil)
)

// HeaderStripAgingResource is the generated Terraform managed resource implementation.
type HeaderStripAgingResource struct {
}

// HeaderStripAgingResourceModel describes the Terraform state and plan shape for HeaderStripAgingResource.
type HeaderStripAgingResourceModel struct {
	AgingInterval types.Int64    `tfsdk:"aging_interval" json:"agingInterval"`
	BoxId         types.String   `tfsdk:"box_id" json:"boxId"`
	DstPort       types.Int64    `tfsdk:"dst_port" json:"dstPort"`
	ProtocolType  types.String   `tfsdk:"protocol_type" json:"protocolType"`
	Timeouts      timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *HeaderStripAgingResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_header_strip_aging"
}

// Schema returns the Terraform schema for this resource.
func (r *HeaderStripAgingResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load header strip aging for target box", Attributes: map[string]schema.Attribute{"aging_interval": schema.Int64Attribute{MarkdownDescription: "Interval in sec. Valid range 300-1000000. Enter 0 to disable.", Optional: true, Computed: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64. all is applicable only for post request.", Required: true}, "dst_port": schema.Int64Attribute{MarkdownDescription: "L4 destination port number.Valid value is between 0 to 65535.", Optional: true, Computed: true}, "protocol_type": schema.StringAttribute{MarkdownDescription: "protocol type", Optional: true, Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *HeaderStripAgingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.BoxId.IsNull() || plan.BoxId.IsUnknown() {
		plan.BoxId = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *HeaderStripAgingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *HeaderStripAgingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.BoxId.IsNull() || plan.BoxId.IsUnknown() {
		if !state.BoxId.IsNull() && !state.BoxId.IsUnknown() {
			plan.BoxId = state.BoxId
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *HeaderStripAgingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *HeaderStripAgingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("box_id"), req.ID)...)
}
