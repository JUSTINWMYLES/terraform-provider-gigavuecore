package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*ConnectionResource)(nil)

// ConnectionResource is the generated Terraform managed resource implementation.
type ConnectionResource struct {
}

// ConnectionResourceModel describes the Terraform state and plan shape for ConnectionResource.
type ConnectionResourceModel struct {
	UnifiedResourcesConnectionGetQueryResponse types.Dynamic `tfsdk:"unified_resources_connection_get_query_response"`
}

// Metadata returns the resource type name.
func (r *ConnectionResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_connection"
}

// Schema returns the Terraform schema for this resource.
func (r *ConnectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get unified resource connection by environment id", Attributes: map[string]schema.Attribute{"unified_resources_connection_get_query_response": schema.DynamicAttribute{MarkdownDescription: "Unified Resource connection", Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConnectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *ConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConnectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConnectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ConnectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *ConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConnectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
