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
	_ resource.Resource                = (*PortFilterResource)(nil)
	_ resource.ResourceWithImportState = (*PortFilterResource)(nil)
)

// PortFilterResource is the generated Terraform managed resource implementation.
type PortFilterResource struct {
}

// PortFilterResourceModel describes the Terraform state and plan shape for PortFilterResource.
type PortFilterResourceModel struct {
	Port   types.String `tfsdk:"port"`
	PortId types.String `tfsdk:"port_id"`
	Rules  types.Object `tfsdk:"rules"`
}

// Metadata returns the resource type name.
func (r *PortFilterResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_port_filter"
}

// Schema returns the Terraform schema for this resource.
func (r *PortFilterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get filter of a Port by port ID", Attributes: map[string]schema.Attribute{"port": schema.StringAttribute{Required: true}, "port_id": schema.StringAttribute{Computed: true}, "rules": schema.SingleNestedAttribute{MarkdownDescription: "Port Filter Rules Container. Private class", Required: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *PortFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PortFilterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.PortId.IsNull() || plan.PortId.IsUnknown() {
		plan.PortId = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *PortFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PortFilterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *PortFilterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PortFilterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state PortFilterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.PortId.IsNull() || plan.PortId.IsUnknown() {
		if !state.PortId.IsNull() && !state.PortId.IsUnknown() {
			plan.PortId = state.PortId
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *PortFilterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PortFilterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *PortFilterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("port_id"), req.ID)...)
}
