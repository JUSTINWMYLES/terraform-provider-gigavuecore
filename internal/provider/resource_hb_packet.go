package provider

import "context"
import (
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*HbPacketResource)(nil)
	_ resource.ResourceWithImportState = (*HbPacketResource)(nil)
)

// HbPacketResource is the generated Terraform managed resource implementation.
type HbPacketResource struct {
}

// HbPacketResourceModel describes the Terraform state and plan shape for HbPacketResource.
type HbPacketResourceModel struct {
	Alias        types.String `tfsdk:"alias"`
	CustomPacket types.String `tfsdk:"custom_packet" json:"customPacket"`
}

// Metadata returns the resource type name.
func (r *HbPacketResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_hb_packet"
}

// Schema returns the Terraform schema for this resource.
func (r *HbPacketResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Heartbeat Packet by alias (deprecated: use GET /inline/hbProfiles/{alias})", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Heartbeat Custom Packet alias. Maps directly (foreign key) into InlineHbProfile alias", Required: true}, "custom_packet": schema.StringAttribute{MarkdownDescription: "Base64-encoded custom pcap packet", Required: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *HbPacketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HbPacketResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		plan.Alias = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *HbPacketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HbPacketResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *HbPacketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HbPacketResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state HbPacketResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *HbPacketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HbPacketResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *HbPacketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
