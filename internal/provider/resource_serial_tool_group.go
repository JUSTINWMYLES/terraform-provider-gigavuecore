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
	_ resource.Resource                = (*SerialToolGroupResource)(nil)
	_ resource.ResourceWithImportState = (*SerialToolGroupResource)(nil)
)

// SerialToolGroupResource is the generated Terraform managed resource implementation.
type SerialToolGroupResource struct {
}

// SerialToolGroupResourceModel describes the Terraform state and plan shape for SerialToolGroupResource.
type SerialToolGroupResourceModel struct {
	Alias              types.String `tfsdk:"alias"`
	Comment            types.String `tfsdk:"comment"`
	Enabled            types.Bool   `tfsdk:"enabled"`
	FailoverAction     types.String `tfsdk:"failover_action" json:"failoverAction"`
	HealthState        types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineTools        types.List   `tfsdk:"inline_tools" json:"inlineTools"`
	PerDirectionOrder  types.String `tfsdk:"per_direction_order" json:"perDirectionOrder"`
}

// Metadata returns the resource type name.
func (r *SerialToolGroupResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_serial_tool_group"
}

// Schema returns the Terraform schema for this resource.
func (r *SerialToolGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Serial Tool Group by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool Group alias. Unique within a cluster", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "setting to false is equivalent to forcing the inline serial tool failure (useful for taking the inline serial tool out of commission for maintenance or other purposes)", Optional: true, Computed: true}, "failover_action": schema.StringAttribute{Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "inline_tools": schema.ListAttribute{MarkdownDescription: "a list of aliases of the inline tools or inline tool groups that participate in the series", Required: true, ElementType: types.StringType}, "per_direction_order": schema.StringAttribute{MarkdownDescription: "direction of traffic flow in reference to the order that the inline-tools are configured", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *SerialToolGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SerialToolGroupResourceModel
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
func (r *SerialToolGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SerialToolGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *SerialToolGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SerialToolGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state SerialToolGroupResourceModel
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
func (r *SerialToolGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SerialToolGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *SerialToolGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
