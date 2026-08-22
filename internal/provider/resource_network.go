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
	_ resource.Resource                = (*NetworkResource)(nil)
	_ resource.ResourceWithImportState = (*NetworkResource)(nil)
)

// NetworkResource is the generated Terraform managed resource implementation.
type NetworkResource struct {
}

// NetworkResourceModel describes the Terraform state and plan shape for NetworkResource.
type NetworkResourceModel struct {
	Alias                  types.String `tfsdk:"alias"`
	Comment                types.String `tfsdk:"comment"`
	ForwardingState        types.String `tfsdk:"forwarding_state" json:"forwardingState"`
	HealthState            types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons     types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Heartbeat              types.Object `tfsdk:"heartbeat"`
	Lfp                    types.Bool   `tfsdk:"lfp"`
	PhysicalBypass         types.Bool   `tfsdk:"physical_bypass" json:"physicalBypass"`
	PortA                  types.String `tfsdk:"port_a" json:"portA"`
	PortB                  types.String `tfsdk:"port_b" json:"portB"`
	RedundancyControlState types.String `tfsdk:"redundancy_control_state" json:"redundancyControlState"`
	RedundancyProfile      types.String `tfsdk:"redundancy_profile" json:"redundancyProfile"`
	TrafficPath            types.String `tfsdk:"traffic_path" json:"trafficPath"`
	Type                   types.String `tfsdk:"type"`
}

// Metadata returns the resource type name.
func (r *NetworkResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_network"
}

// Schema returns the Terraform schema for this resource.
func (r *NetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Network by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Network alias. Unique within a cluster", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "forwarding_state": schema.StringAttribute{Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "heartbeat": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Network.", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Required: true}}}, "lfp": schema.BoolAttribute{MarkdownDescription: "Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side", Optional: true, Computed: true}, "physical_bypass": schema.BoolAttribute{MarkdownDescription: "only applicable for 'protected' inline networks", Optional: true, Computed: true}, "port_a": schema.StringAttribute{MarkdownDescription: "portId of side A inline network port", Required: true}, "port_b": schema.StringAttribute{MarkdownDescription: "portId of side B inline network port", Required: true}, "redundancy_control_state": schema.StringAttribute{MarkdownDescription: "Redundancy Control State. For 'protected' inline networks", Optional: true, Computed: true}, "redundancy_profile": schema.StringAttribute{MarkdownDescription: "Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks", Optional: true, Computed: true}, "traffic_path": schema.StringAttribute{Optional: true, Computed: true}, "type": schema.StringAttribute{Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkResourceModel
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
func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NetworkResourceModel
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
func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *NetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
