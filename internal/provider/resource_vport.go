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
	_ resource.Resource                = (*VportResource)(nil)
	_ resource.ResourceWithImportState = (*VportResource)(nil)
)

// VportResource is the generated Terraform managed resource implementation.
type VportResource struct {
}

// VportResourceModel describes the Terraform state and plan shape for VportResource.
type VportResourceModel struct {
	Alias              types.String `tfsdk:"alias"`
	DeferredBinding    types.Bool   `tfsdk:"deferred_binding" json:"deferredBinding"`
	FailOverAction     types.String `tfsdk:"fail_over_action" json:"failOverAction"`
	GsGroup            types.String `tfsdk:"gs_group" json:"gsGroup"`
	HealthState        types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineStatus       types.String `tfsdk:"inline_status" json:"inlineStatus"`
	InnerTrafficPath   types.String `tfsdk:"inner_traffic_path" json:"innerTrafficPath"`
	MetadataMonitoring types.Object `tfsdk:"metadata_monitoring" json:"metadataMonitoring"`
	Mode               types.String `tfsdk:"mode"`
	OuterTrafficPath   types.String `tfsdk:"outer_traffic_path" json:"outerTrafficPath"`
	SaApfProfile       types.String `tfsdk:"sa_apf_profile" json:"saApfProfile"`
}

// Metadata returns the resource type name.
func (r *VportResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_vport"
}

// Schema returns the Terraform schema for this resource.
func (r *VportResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find vPort by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "deferred_binding": schema.BoolAttribute{MarkdownDescription: "enable/disable deferred-binding", Optional: true, Computed: true}, "fail_over_action": schema.StringAttribute{Optional: true, Computed: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "inline_status": schema.StringAttribute{Optional: true, Computed: true}, "inner_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-network traffic-path, applicable for inner map", Optional: true, Computed: true}, "metadata_monitoring": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "metadata monitoring action", Optional: true, Computed: true}, "exporters": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "mode": schema.StringAttribute{Optional: true, Computed: true}, "outer_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map", Optional: true, Computed: true}, "sa_apf_profile": schema.StringAttribute{MarkdownDescription: "ASF session profile", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *VportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VportResourceModel
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
func (r *VportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VportResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *VportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VportResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state VportResourceModel
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
func (r *VportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VportResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *VportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
