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
	_ resource.Resource                = (*NetworkLagResource)(nil)
	_ resource.ResourceWithImportState = (*NetworkLagResource)(nil)
)

// NetworkLagResource is the generated Terraform managed resource implementation.
type NetworkLagResource struct {
}

// NetworkLagResourceModel describes the Terraform state and plan shape for NetworkLagResource.
type NetworkLagResourceModel struct {
	Alias                  types.String `tfsdk:"alias"`
	Cdp                    types.Bool   `tfsdk:"cdp"`
	Comment                types.String `tfsdk:"comment"`
	ForwardingState        types.String `tfsdk:"forwarding_state" json:"forwardingState"`
	HealthState            types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons     types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineNetworks         types.Set    `tfsdk:"inline_networks" json:"inlineNetworks"`
	Lacp                   types.Bool   `tfsdk:"lacp"`
	Lfp                    types.Bool   `tfsdk:"lfp"`
	PhysicalBypass         types.Bool   `tfsdk:"physical_bypass" json:"physicalBypass"`
	RedundancyControlState types.String `tfsdk:"redundancy_control_state" json:"redundancyControlState"`
	RedundancyProfile      types.String `tfsdk:"redundancy_profile" json:"redundancyProfile"`
	TrafficPath            types.String `tfsdk:"traffic_path" json:"trafficPath"`
}

// Metadata returns the resource type name.
func (r *NetworkLagResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_network_lag"
}

// Schema returns the Terraform schema for this resource.
func (r *NetworkLagResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Network LAG by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Network LAG alias. Unique within a cluster", Required: true}, "cdp": schema.BoolAttribute{MarkdownDescription: "Enable or Disable Cisco Discovery Protocol", Optional: true, Computed: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "forwarding_state": schema.StringAttribute{Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "inline_networks": schema.SetAttribute{MarkdownDescription: "list of inline-network-lag aliases", Required: true, ElementType: types.StringType}, "lacp": schema.BoolAttribute{MarkdownDescription: "Enable or Disable Link Aggregation Control Protocol", Optional: true, Computed: true}, "lfp": schema.BoolAttribute{MarkdownDescription: "Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side. Changes will be pushed down to all inline-network members", Optional: true, Computed: true}, "physical_bypass": schema.BoolAttribute{MarkdownDescription: "changes will be pushed down to all inline-network members", Optional: true, Computed: true}, "redundancy_control_state": schema.StringAttribute{MarkdownDescription: "For Redundancy Control State. From 'protected' inline networks", Optional: true, Computed: true}, "redundancy_profile": schema.StringAttribute{MarkdownDescription: "Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks", Optional: true, Computed: true}, "traffic_path": schema.StringAttribute{Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NetworkLagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkLagResourceModel
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
func (r *NetworkLagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkLagResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NetworkLagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworkLagResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NetworkLagResourceModel
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
func (r *NetworkLagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkLagResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *NetworkLagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
