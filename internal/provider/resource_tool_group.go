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
	_ resource.Resource                = (*ToolGroupResource)(nil)
	_ resource.ResourceWithImportState = (*ToolGroupResource)(nil)
)

// ToolGroupResource is the generated Terraform managed resource implementation.
type ToolGroupResource struct {
}

// ToolGroupResourceModel describes the Terraform state and plan shape for ToolGroupResource.
type ToolGroupResourceModel struct {
	Alias                  types.String `tfsdk:"alias"`
	Comment                types.String `tfsdk:"comment"`
	CurrentState           types.Object `tfsdk:"current_state" json:"currentState"`
	Enabled                types.Bool   `tfsdk:"enabled"`
	FailoverAction         types.String `tfsdk:"failover_action" json:"failoverAction"`
	FailoverMode           types.String `tfsdk:"failover_mode" json:"failoverMode"`
	FlexStatus             types.String `tfsdk:"flex_status" json:"flexStatus"`
	FlexTrafficPath        types.String `tfsdk:"flex_traffic_path" json:"flexTrafficPath"`
	Hash                   types.String `tfsdk:"hash"`
	HealthState            types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons     types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineTools            types.List   `tfsdk:"inline_tools" json:"inlineTools"`
	MinGroupSize           types.Int64  `tfsdk:"min_group_size" json:"minGroupSize"`
	OperStatus             types.String `tfsdk:"oper_status" json:"operStatus"`
	ReleaseSpareIfPossible types.Bool   `tfsdk:"release_spare_if_possible" json:"releaseSpareIfPossible"`
	SpareInlineTool        types.String `tfsdk:"spare_inline_tool" json:"spareInlineTool"`
}

// Metadata returns the resource type name.
func (r *ToolGroupResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tool_group"
}

// Schema returns the Terraform schema for this resource.
func (r *ToolGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Tool Group by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool Group alias. Unique within a cluster", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "current_state": schema.SingleNestedAttribute{MarkdownDescription: "Inline Tool Group Smart Load Balancing definition", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inline_tools": schema.ListAttribute{MarkdownDescription: "Current list of inlineTools", Optional: true, Computed: true, ElementType: types.StringType}, "spare_tool": schema.StringAttribute{MarkdownDescription: "Current spare tool", Optional: true, Computed: true}, "spare_tool_status": schema.StringAttribute{MarkdownDescription: "Current spare tool status", Optional: true, Computed: true}, "switched_inline_tool": schema.StringAttribute{MarkdownDescription: "Current switched inlineTool", Optional: true, Computed: true}}}, "enabled": schema.BoolAttribute{MarkdownDescription: "setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)", Optional: true, Computed: true}, "failover_action": schema.StringAttribute{Required: true}, "failover_mode": schema.StringAttribute{MarkdownDescription: "the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed", Required: true}, "flex_status": schema.StringAttribute{Optional: true, Computed: true}, "flex_traffic_path": schema.StringAttribute{Optional: true, Computed: true}, "hash": schema.StringAttribute{Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "inline_tools": schema.ListAttribute{MarkdownDescription: "If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise", Required: true, ElementType: types.StringType}, "min_group_size": schema.Int64Attribute{MarkdownDescription: " the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up", Optional: true, Computed: true}, "oper_status": schema.StringAttribute{Optional: true, Computed: true}, "release_spare_if_possible": schema.BoolAttribute{MarkdownDescription: "when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool", Optional: true, Computed: true}, "spare_inline_tool": schema.StringAttribute{MarkdownDescription: "alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ToolGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolGroupResourceModel
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
func (r *ToolGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ToolGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ToolGroupResourceModel
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
func (r *ToolGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *ToolGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
