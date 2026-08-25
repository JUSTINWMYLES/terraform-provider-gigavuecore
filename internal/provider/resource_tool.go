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
	_ resource.Resource                = (*ToolResource)(nil)
	_ resource.ResourceWithImportState = (*ToolResource)(nil)
)

// ToolResource is the generated Terraform managed resource implementation.
type ToolResource struct {
}

// ToolResourceModel describes the Terraform state and plan shape for ToolResource.
type ToolResourceModel struct {
	Alias                   types.String `tfsdk:"alias"`
	CombinedHeartBeatStatus types.String `tfsdk:"combined_heart_beat_status" json:"combinedHeartBeatStatus"`
	Comment                 types.String `tfsdk:"comment"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	FailoverAction          types.String `tfsdk:"failover_action" json:"failoverAction"`
	FlexStatus              types.String `tfsdk:"flex_status" json:"flexStatus"`
	FlexTrafficPath         types.String `tfsdk:"flex_traffic_path" json:"flexTrafficPath"`
	HealthState             types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons      types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Heartbeat               types.Object `tfsdk:"heartbeat"`
	InlineToolType          types.String `tfsdk:"inline_tool_type" json:"inlineToolType"`
	NegativeHeartbeat       types.Object `tfsdk:"negative_heartbeat" json:"negativeHeartbeat"`
	OperationalState        types.String `tfsdk:"operational_state" json:"operationalState"`
	PortA                   types.String `tfsdk:"port_a" json:"portA"`
	PortAStatus             types.String `tfsdk:"port_a_status" json:"portAStatus"`
	PortB                   types.String `tfsdk:"port_b" json:"portB"`
	PortBStatus             types.String `tfsdk:"port_b_status" json:"portBStatus"`
	RecoveryMode            types.String `tfsdk:"recovery_mode" json:"recoveryMode"`
	Shared                  types.Bool   `tfsdk:"shared"`
	Timestamp               types.String `tfsdk:"timestamp"`
}

// Metadata returns the resource type name.
func (r *ToolResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tool"
}

// Schema returns the Terraform schema for this resource.
func (r *ToolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Tool by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool alias. Unique within a cluster", Required: true}, "combined_heart_beat_status": schema.StringAttribute{MarkdownDescription: "combined Heartbeat status", Optional: true, Computed: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "failover_action": schema.StringAttribute{Optional: true, Computed: true}, "flex_status": schema.StringAttribute{Optional: true, Computed: true}, "flex_traffic_path": schema.StringAttribute{Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "heartbeat": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Required: true}, "ip_address_a": schema.StringAttribute{MarkdownDescription: "the destination IP address to be used in heartbeat packets send from side A to side B (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)", Required: true}, "ip_address_b": schema.StringAttribute{MarkdownDescription: "the destination IP address to be used in heartbeat packets send from side B to side A (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)", Required: true}, "profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Heartbeat Profile. the default is the heartbeat profile named 'default'", Required: true}, "status": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"heartbeat_passing": schema.StringAttribute{MarkdownDescription: "indicates whether heartbeat packets from portA are reaching portB and vice versa", Optional: true, Computed: true}, "stats_ato_b": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}, "stats_bto_a": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}}}}}, "inline_tool_type": schema.StringAttribute{MarkdownDescription: "inlineTool type", Optional: true, Computed: true}, "negative_heartbeat": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Required: true}, "profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Negative Heartbeat Profile", Optional: true, Computed: true}, "status": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"heartbeat_passing": schema.StringAttribute{MarkdownDescription: "indicates whether heartbeat packets from portA are reaching portB and vice versa", Optional: true, Computed: true}, "stats_ato_b": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}, "stats_bto_a": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}}}}}, "operational_state": schema.StringAttribute{MarkdownDescription: "Operational State", Optional: true, Computed: true}, "port_a": schema.StringAttribute{MarkdownDescription: "portId of side A inline tool port", Required: true}, "port_a_status": schema.StringAttribute{MarkdownDescription: "port A status", Optional: true, Computed: true}, "port_b": schema.StringAttribute{MarkdownDescription: "portId of side B inline tool port", Required: true}, "port_b_status": schema.StringAttribute{MarkdownDescription: "port B status", Optional: true, Computed: true}, "recovery_mode": schema.StringAttribute{Optional: true, Computed: true}, "shared": schema.BoolAttribute{MarkdownDescription: "inline tool sharing mode", Optional: true, Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ToolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolResourceModel
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
func (r *ToolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ToolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ToolResourceModel
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
func (r *ToolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *ToolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
