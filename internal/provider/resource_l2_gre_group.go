package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*L2GreGroupResource)(nil)

// L2GreGroupResource is the generated Terraform managed resource implementation.
type L2GreGroupResource struct {
}

// L2GreGroupResourceModel describes the Terraform state and plan shape for L2GreGroupResource.
type L2GreGroupResourceModel struct {
	Alias                    types.String `tfsdk:"alias"`
	BoxId                    types.String `tfsdk:"box_id" json:"boxId"`
	CircuitTunnelL2GreGroups types.List   `tfsdk:"circuit_tunnel_l2_gre_groups" json:"circuitTunnelL2GreGroups"`
	Comment                  types.String `tfsdk:"comment"`
	Context                  types.Object `tfsdk:"context"`
	Id                       types.String `tfsdk:"id"`
	L2GreIds                 types.List   `tfsdk:"l2_gre_ids" json:"l2GreIds"`
}

// Metadata returns the resource type name.
func (r *L2GreGroupResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_l2_gre_group"
}

// Schema returns the Terraform schema for this resource.
func (r *L2GreGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Circuit Tunnel L2Gre Groups", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Optional: true}, "circuit_tunnel_l2_gre_groups": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "l2_gre_ids": schema.ListAttribute{Computed: true, ElementType: types.Int64Type}}}}, "comment": schema.StringAttribute{Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "id": schema.StringAttribute{Computed: true}, "l2_gre_ids": schema.ListAttribute{Optional: true, ElementType: types.Int64Type}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *L2GreGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan L2GreGroupResourceModel
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
func (r *L2GreGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state L2GreGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *L2GreGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan L2GreGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state L2GreGroupResourceModel
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
func (r *L2GreGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state L2GreGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
