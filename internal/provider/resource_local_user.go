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
	_ resource.Resource                = (*LocalUserResource)(nil)
	_ resource.ResourceWithImportState = (*LocalUserResource)(nil)
)

// LocalUserResource is the generated Terraform managed resource implementation.
type LocalUserResource struct {
}

// LocalUserResourceModel describes the Terraform state and plan shape for LocalUserResource.
type LocalUserResourceModel struct {
	AccountStatus   types.String `tfsdk:"account_status" json:"accountStatus"`
	Capability      types.String `tfsdk:"capability"`
	CurrentPassword types.String `tfsdk:"current_password" json:"currentPassword"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	FullName        types.String `tfsdk:"full_name" json:"fullName"`
	Roles           types.List   `tfsdk:"roles"`
	UserPwd         types.String `tfsdk:"user_pwd" json:"userPwd"`
	Username        types.String `tfsdk:"username"`
}

// Metadata returns the resource type name.
func (r *LocalUserResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_local_user"
}

// Schema returns the Terraform schema for this resource.
func (r *LocalUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Local User by username", Attributes: map[string]schema.Attribute{"account_status": schema.StringAttribute{Optional: true, Computed: true}, "capability": schema.StringAttribute{Optional: true, Computed: true}, "current_password": schema.StringAttribute{MarkdownDescription: "Specifies the logged in password", Required: true, Sensitive: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Temporarily enable/disable logins for the specified account. Disabling an account closes any currently open sessions for the specified account", Optional: true, Computed: true}, "full_name": schema.StringAttribute{MarkdownDescription: "Full name for the account (referred to sometimes as the gecos)", Optional: true, Computed: true}, "roles": schema.ListAttribute{MarkdownDescription: "References to the Roles defined for the User", Optional: true, Computed: true, ElementType: types.StringType}, "user_pwd": schema.StringAttribute{Required: true}, "username": schema.StringAttribute{MarkdownDescription: "Specifies the username of the local user", Required: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *LocalUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LocalUserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.Username.IsNull() || plan.Username.IsUnknown() {
		plan.Username = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *LocalUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LocalUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *LocalUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan LocalUserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state LocalUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Username.IsNull() || plan.Username.IsUnknown() {
		if !state.Username.IsNull() && !state.Username.IsUnknown() {
			plan.Username = state.Username
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *LocalUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LocalUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *LocalUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("username"), req.ID)...)
}
