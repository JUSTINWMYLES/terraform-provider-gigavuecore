package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*NtpServerResource)(nil)

// NtpServerResource is the generated Terraform managed resource implementation.
type NtpServerResource struct {
}

// NtpServerResourceModel describes the Terraform state and plan shape for NtpServerResource.
type NtpServerResourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Id         types.String `tfsdk:"id"`
	KeyEnabled types.Bool   `tfsdk:"key_enabled" json:"keyEnabled"`
	KeyNumber  types.Int64  `tfsdk:"key_number" json:"keyNumber"`
	Preferred  types.Bool   `tfsdk:"preferred"`
	Server     types.String `tfsdk:"server"`
	Version    types.String `tfsdk:"version"`
}

// Metadata returns the resource type name.
func (r *NtpServerResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_ntp_server"
}

// Schema returns the Terraform schema for this resource.
func (r *NtpServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find NtpServer by address", Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true, Computed: true}, "id": schema.StringAttribute{Computed: true}, "key_enabled": schema.BoolAttribute{Optional: true, Computed: true}, "key_number": schema.Int64Attribute{Optional: true, Computed: true}, "preferred": schema.BoolAttribute{Optional: true, Computed: true}, "server": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname", Required: true}, "version": schema.StringAttribute{Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NtpServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NtpServerResourceModel
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
func (r *NtpServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NtpServerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NtpServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NtpServerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NtpServerResourceModel
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
func (r *NtpServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NtpServerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
