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
	_ resource.Resource                = (*RadiusServerResource)(nil)
	_ resource.ResourceWithImportState = (*RadiusServerResource)(nil)
)

// RadiusServerResource is the generated Terraform managed resource implementation.
type RadiusServerResource struct {
}

// RadiusServerResourceModel describes the Terraform state and plan shape for RadiusServerResource.
type RadiusServerResourceModel struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	Port          types.Int64  `tfsdk:"port"`
	Retries       types.Int64  `tfsdk:"retries"`
	SecretKey     types.String `tfsdk:"secret_key" json:"secretKey"`
	ServerAddress types.String `tfsdk:"server_address" json:"serverAddress"`
	Timeout       types.Int64  `tfsdk:"timeout"`
}

// Metadata returns the resource type name.
func (r *RadiusServerResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_radius_server"
}

// Schema returns the Terraform schema for this resource.
func (r *RadiusServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find RADIUS Server by address", Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true, Computed: true}, "port": schema.Int64Attribute{Optional: true, Computed: true}, "retries": schema.Int64Attribute{MarkdownDescription: "value of 0 disables retries. Defaults to the value defined in the RadiusServerDefaults", Optional: true, Computed: true}, "secret_key": schema.StringAttribute{MarkdownDescription: "if not included, defaults to the value defined in the RadiusServerDefaults", Required: true, Sensitive: true}, "server_address": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname", Required: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds. Defaults to the value defined in the RadiusServerDefaults", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *RadiusServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RadiusServerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.ServerAddress.IsNull() || plan.ServerAddress.IsUnknown() {
		plan.ServerAddress = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *RadiusServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RadiusServerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *RadiusServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RadiusServerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state RadiusServerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.ServerAddress.IsNull() || plan.ServerAddress.IsUnknown() {
		if !state.ServerAddress.IsNull() && !state.ServerAddress.IsUnknown() {
			plan.ServerAddress = state.ServerAddress
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *RadiusServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RadiusServerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *RadiusServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("server_address"), req.ID)...)
}
