package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*ActivationResource)(nil)

// ActivationResource is the generated Terraform managed resource implementation.
type ActivationResource struct {
}

// ActivationResourceModel describes the Terraform state and plan shape for ActivationResource.
type ActivationResourceModel struct {
	Activations types.List   `tfsdk:"activations"`
	Context     types.Object `tfsdk:"context"`
	EliId       types.String `tfsdk:"eli_id" json:"eliId"`
	Id          types.String `tfsdk:"id"`
	Quantity    types.Int64  `tfsdk:"quantity"`
}

// Metadata returns the resource type name.
func (r *ActivationResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_activation"
}

// Schema returns the Terraform schema for this resource.
func (r *ActivationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get license activations(for a specific entitlement,or for all entitlements)", Attributes: map[string]schema.Attribute{"activations": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"eid": schema.StringAttribute{Computed: true}, "entl_item_id": schema.StringAttribute{Computed: true}, "gid": schema.StringAttribute{Computed: true}, "imported": schema.BoolAttribute{Computed: true}, "num_licenses": schema.Int64Attribute{Computed: true}, "owner_fm": schema.BoolAttribute{Computed: true}, "vmac": schema.StringAttribute{Computed: true}}}}, "context": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{Computed: true}, "page_size": schema.Int64Attribute{Computed: true}, "sort": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{Computed: true}}}, "eli_id": schema.StringAttribute{Optional: true}, "id": schema.StringAttribute{Computed: true}, "quantity": schema.Int64Attribute{Optional: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ActivationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ActivationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		plan.Id = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *ActivationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ActivationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ActivationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ActivationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ActivationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		if !state.Id.IsNull() && !state.Id.IsUnknown() {
			plan.Id = state.Id
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *ActivationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ActivationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
