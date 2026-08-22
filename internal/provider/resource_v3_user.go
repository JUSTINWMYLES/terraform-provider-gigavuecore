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
	_ resource.Resource                = (*V3UserResource)(nil)
	_ resource.ResourceWithImportState = (*V3UserResource)(nil)
)

// V3UserResource is the generated Terraform managed resource implementation.
type V3UserResource struct {
}

// V3UserResourceModel describes the Terraform state and plan shape for V3UserResource.
type V3UserResourceModel struct {
	AuthKey      types.String `tfsdk:"auth_key" json:"authKey"`
	AuthProtocol types.String `tfsdk:"auth_protocol" json:"authProtocol"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	PrivKey      types.String `tfsdk:"priv_key" json:"privKey"`
	PrivProtocol types.String `tfsdk:"priv_protocol" json:"privProtocol"`
	ReadOnly     types.Bool   `tfsdk:"read_only" json:"readOnly"`
	UserName     types.String `tfsdk:"user_name"`
	Username     types.String `tfsdk:"username"`
}

// Metadata returns the resource type name.
func (r *V3UserResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_v3_user"
}

// Schema returns the Terraform schema for this resource.
func (r *V3UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find SNMPv3 User by name", Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{Optional: true, Computed: true}, "auth_protocol": schema.StringAttribute{Optional: true, Computed: true}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "priv_key": schema.StringAttribute{Optional: true, Computed: true}, "priv_protocol": schema.StringAttribute{Optional: true, Computed: true}, "read_only": schema.BoolAttribute{MarkdownDescription: "This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.", Optional: true, Computed: true}, "user_name": schema.StringAttribute{Computed: true}, "username": schema.StringAttribute{Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *V3UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V3UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.UserName.IsNull() || plan.UserName.IsUnknown() {
		plan.UserName = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *V3UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state V3UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *V3UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan V3UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state V3UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.UserName.IsNull() || plan.UserName.IsUnknown() {
		if !state.UserName.IsNull() && !state.UserName.IsUnknown() {
			plan.UserName = state.UserName
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *V3UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state V3UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *V3UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("user_name"), req.ID)...)
}
