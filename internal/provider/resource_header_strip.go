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
	_ resource.Resource                = (*HeaderStripResource)(nil)
	_ resource.ResourceWithImportState = (*HeaderStripResource)(nil)
)

// HeaderStripResource is the generated Terraform managed resource implementation.
type HeaderStripResource struct {
}

// HeaderStripResourceModel describes the Terraform state and plan shape for HeaderStripResource.
type HeaderStripResourceModel struct {
	BoxId      types.String `tfsdk:"box_id" json:"boxId"`
	MplsLabels types.List   `tfsdk:"mpls_labels" json:"mplsLabels"`
}

// Metadata returns the resource type name.
func (r *HeaderStripResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_header_strip"
}

// Schema returns the Terraform schema for this resource.
func (r *HeaderStripResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load header strip for target box", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Required: true}, "mpls_labels": schema.ListAttribute{MarkdownDescription: "mpls ids, valid and required. Range can be specified. Example:1..200", Optional: true, Computed: true, ElementType: types.StringType}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *HeaderStripResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HeaderStripResourceModel
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
func (r *HeaderStripResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HeaderStripResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *HeaderStripResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HeaderStripResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state HeaderStripResourceModel
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
func (r *HeaderStripResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HeaderStripResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *HeaderStripResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("box_id"), req.ID)...)
}
