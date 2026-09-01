package provider

import "context"
import (
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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
	Eid         types.String   `tfsdk:"eid"`
	EliId       types.String   `tfsdk:"eli_id" json:"eliId"`
	EntlItemId  types.String   `tfsdk:"entl_item_id" json:"entlItemId"`
	Gid         types.String   `tfsdk:"gid"`
	Imported    types.Bool     `tfsdk:"imported"`
	NumLicenses types.Int64    `tfsdk:"num_licenses" json:"numLicenses"`
	OwnerFm     types.Bool     `tfsdk:"owner_fm" json:"ownerFm"`
	Page        types.String   `tfsdk:"page"`
	Quantity    types.Int64    `tfsdk:"quantity"`
	Sort        types.String   `tfsdk:"sort"`
	Vmac        types.String   `tfsdk:"vmac"`
	Timeouts    timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *ActivationResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_activation"
}

// Schema returns the Terraform schema for this resource.
func (r *ActivationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create a new activation from an entitlement", Attributes: map[string]schema.Attribute{"eid": schema.StringAttribute{Computed: true}, "eli_id": schema.StringAttribute{Optional: true}, "entl_item_id": schema.StringAttribute{Computed: true}, "gid": schema.StringAttribute{Computed: true}, "imported": schema.BoolAttribute{Computed: true}, "num_licenses": schema.Int64Attribute{Computed: true}, "owner_fm": schema.BoolAttribute{Computed: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "quantity": schema.Int64Attribute{Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (fieldName:sortOrder) format. The default sort order is DESC. Example: sort=(numLicenses:DESC)", Optional: true}, "vmac": schema.StringAttribute{Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ActivationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ActivationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.EntlItemId.IsNull() || plan.EntlItemId.IsUnknown() {
		plan.EntlItemId = types.StringValue("generated")
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
	if plan.EntlItemId.IsNull() || plan.EntlItemId.IsUnknown() {
		if !state.EntlItemId.IsNull() && !state.EntlItemId.IsUnknown() {
			plan.EntlItemId = state.EntlItemId
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
