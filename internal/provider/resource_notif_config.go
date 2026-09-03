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
var _ resource.Resource = (*NotifConfigResource)(nil)

// NotifConfigResource is the generated Terraform managed resource implementation.
type NotifConfigResource struct {
}

// NotifConfigResourceModel describes the Terraform state and plan shape for NotifConfigResource.
type NotifConfigResourceModel struct {
	InterfaceName types.String   `tfsdk:"interface_name" json:"interfaceName"`
	InterfaceType types.String   `tfsdk:"interface_type" json:"interfaceType"`
	TargetAddress types.String   `tfsdk:"target_address" json:"targetAddress"`
	Timeouts      timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *NotifConfigResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_notif_config"
}

// Schema returns the Terraform schema for this resource.
func (r *NotifConfigResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Modify Notification Config", Attributes: map[string]schema.Attribute{"interface_name": schema.StringAttribute{MarkdownDescription: "Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.", Optional: true, Computed: true}, "interface_type": schema.StringAttribute{MarkdownDescription: "Notification target interface type", Required: true, Validators: []validator.String{stringvalidator.OneOf("management", "data")}}, "target_address": schema.StringAttribute{MarkdownDescription: "Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NotifConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NotifConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.InterfaceType.IsNull() || plan.InterfaceType.IsUnknown() {
		plan.InterfaceType = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *NotifConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NotifConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NotifConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NotifConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NotifConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.InterfaceType.IsNull() || plan.InterfaceType.IsUnknown() {
		if !state.InterfaceType.IsNull() && !state.InterfaceType.IsUnknown() {
			plan.InterfaceType = state.InterfaceType
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *NotifConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NotifConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
