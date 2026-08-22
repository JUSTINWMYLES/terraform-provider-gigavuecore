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
	_ resource.Resource                = (*NotifTargetResource)(nil)
	_ resource.ResourceWithImportState = (*NotifTargetResource)(nil)
)

// NotifTargetResource is the generated Terraform managed resource implementation.
type NotifTargetResource struct {
}

// NotifTargetResourceModel describes the Terraform state and plan shape for NotifTargetResource.
type NotifTargetResourceModel struct {
	Enabled            types.Bool   `tfsdk:"enabled"`
	Host               types.String `tfsdk:"host"`
	NotifTargetAddress types.String `tfsdk:"notif_target_address"`
	NotifyConfig       types.Object `tfsdk:"notify_config" json:"notifyConfig"`
	NotifyType         types.String `tfsdk:"notify_type" json:"notifyType"`
}

// Metadata returns the resource type name.
func (r *NotifTargetResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_notif_target"
}

// Schema returns the Terraform schema for this resource.
func (r *NotifTargetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find SNMP Notification Target by address", Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "temporarily enable/disable the notification destination", Optional: true, Computed: true}, "host": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or domain name", Required: true}, "notif_target_address": schema.StringAttribute{Computed: true}, "notify_config": schema.SingleNestedAttribute{MarkdownDescription: "Notification Target configuration for specific Notification type (Trap/Inform)", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{MarkdownDescription: "authentication password. required with 'v3user'", Optional: true, Computed: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "authentication hash algorithm. required with 'v3user'", Optional: true, Computed: true}, "community": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v2c'", Optional: true, Computed: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "remote engineID. only valid with notifyType 'inform' and 'version' v3", Optional: true, Computed: true}, "port": schema.Int64Attribute{Optional: true, Computed: true}, "priv_key": schema.StringAttribute{MarkdownDescription: "privacy password", Optional: true, Computed: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "privacy encryption", Optional: true, Computed: true}, "v3_user": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v3'", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "SNMP version to use. v1 is only valid for traps. for v3, user name should be provided", Optional: true, Computed: true}}}, "notify_type": schema.StringAttribute{MarkdownDescription: "SNMP notification type to use", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NotifTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NotifTargetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.NotifTargetAddress.IsNull() || plan.NotifTargetAddress.IsUnknown() {
		plan.NotifTargetAddress = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *NotifTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NotifTargetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NotifTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NotifTargetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NotifTargetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.NotifTargetAddress.IsNull() || plan.NotifTargetAddress.IsUnknown() {
		if !state.NotifTargetAddress.IsNull() && !state.NotifTargetAddress.IsUnknown() {
			plan.NotifTargetAddress = state.NotifTargetAddress
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *NotifTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NotifTargetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *NotifTargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("notif_target_address"), req.ID)...)
}
