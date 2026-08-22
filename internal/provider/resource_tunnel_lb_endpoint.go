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
	_ resource.Resource                = (*TunnelLbEndpointResource)(nil)
	_ resource.ResourceWithImportState = (*TunnelLbEndpointResource)(nil)
)

// TunnelLbEndpointResource is the generated Terraform managed resource implementation.
type TunnelLbEndpointResource struct {
}

// TunnelLbEndpointResourceModel describes the Terraform state and plan shape for TunnelLbEndpointResource.
type TunnelLbEndpointResourceModel struct {
	Alias     types.String `tfsdk:"alias"`
	IpAddress types.String `tfsdk:"ip_address" json:"ipAddress"`
	TeId      types.String `tfsdk:"te_id" json:"teId"`
}

// Metadata returns the resource type name.
func (r *TunnelLbEndpointResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tunnel_lb_endpoint"
}

// Schema returns the Terraform schema for this resource.
func (r *TunnelLbEndpointResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Tunnel Endpoint by id", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "tunnel endpoint alias", Optional: true, Computed: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "tunnel endpoint remote ip address. IPv4 or IPv6", Optional: true, Computed: true}, "te_id": schema.StringAttribute{MarkdownDescription: "tunnel endpoint alias, format: teN , where 1 <= N <= 128", Required: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *TunnelLbEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TunnelLbEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.TeId.IsNull() || plan.TeId.IsUnknown() {
		plan.TeId = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *TunnelLbEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TunnelLbEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *TunnelLbEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TunnelLbEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state TunnelLbEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.TeId.IsNull() || plan.TeId.IsUnknown() {
		if !state.TeId.IsNull() && !state.TeId.IsUnknown() {
			plan.TeId = state.TeId
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *TunnelLbEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TunnelLbEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *TunnelLbEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("te_id"), req.ID)...)
}
