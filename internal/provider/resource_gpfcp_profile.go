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
	_ resource.Resource                = (*GpfcpProfileResource)(nil)
	_ resource.ResourceWithImportState = (*GpfcpProfileResource)(nil)
)

// GpfcpProfileResource is the generated Terraform managed resource implementation.
type GpfcpProfileResource struct {
}

// GpfcpProfileResourceModel describes the Terraform state and plan shape for GpfcpProfileResource.
type GpfcpProfileResourceModel struct {
	Alias     types.String `tfsdk:"alias"`
	Comment   types.String `tfsdk:"comment"`
	GProfiles types.List   `tfsdk:"g_profiles" json:"gProfiles"`
}

// Metadata returns the resource type name.
func (r *GpfcpProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_gpfcp_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *GpfcpProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "new in H 6.8", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the Gpfcp Profile", Required: true}, "comment": schema.StringAttribute{MarkdownDescription: "Description of the Gpfcp Profile", Optional: true, Computed: true}, "g_profiles": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{MarkdownDescription: "Description of the Gpfcp Profile Rule", Optional: true, Computed: true}, "g_interface": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"ip_addresses": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "ip_interface": schema.StringAttribute{Required: true}, "node_type": schema.StringAttribute{Required: true}, "port_list": schema.ListAttribute{Required: true, ElementType: types.Int64Type}}}}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *GpfcpProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GpfcpProfileResourceModel
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
func (r *GpfcpProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GpfcpProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *GpfcpProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GpfcpProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state GpfcpProfileResourceModel
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
func (r *GpfcpProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GpfcpProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *GpfcpProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
