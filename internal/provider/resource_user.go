package provider

import "context"
import (
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*UserResource)(nil)

// UserResource is the generated Terraform managed resource implementation.
type UserResource struct {
}

// UserResourceModel describes the Terraform state and plan shape for UserResource.
type UserResourceModel struct {
	EmailId  types.String   `tfsdk:"email_id" json:"emailId"`
	Enabled  types.Bool     `tfsdk:"enabled"`
	FullName types.String   `tfsdk:"full_name" json:"fullName"`
	Groups   types.List     `tfsdk:"groups"`
	Password types.String   `tfsdk:"password"`
	Username types.String   `tfsdk:"username"`
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *UserResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_user"
}

// Schema returns the Terraform schema for this resource.
func (r *UserResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load User by Username", Attributes: map[string]schema.Attribute{"email_id": schema.StringAttribute{MarkdownDescription: "email ID", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "full_name": schema.StringAttribute{MarkdownDescription: "user's full name", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "groups": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "password": schema.StringAttribute{MarkdownDescription: "password", Required: true, Sensitive: true, Validators: []validator.String{stringvalidator.LengthAtLeast(8)}}, "username": schema.StringAttribute{MarkdownDescription: "username", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserResourceModel
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
func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state UserResourceModel
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
func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
