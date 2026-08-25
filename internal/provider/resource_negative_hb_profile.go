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
	_ resource.Resource                = (*NegativeHbProfileResource)(nil)
	_ resource.ResourceWithImportState = (*NegativeHbProfileResource)(nil)
)

// NegativeHbProfileResource is the generated Terraform managed resource implementation.
type NegativeHbProfileResource struct {
}

// NegativeHbProfileResourceModel describes the Terraform state and plan shape for NegativeHbProfileResource.
type NegativeHbProfileResourceModel struct {
	Alias                types.String `tfsdk:"alias"`
	CustomPacket         types.String `tfsdk:"custom_packet" json:"customPacket"`
	CustomPacketFileName types.String `tfsdk:"custom_packet_file_name" json:"customPacketFileName"`
	Direction            types.String `tfsdk:"direction"`
	Period               types.Int64  `tfsdk:"period"`
	RecoveryTime         types.Int64  `tfsdk:"recovery_time" json:"recoveryTime"`
}

// Metadata returns the resource type name.
func (r *NegativeHbProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_negative_hb_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *NegativeHbProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Negative Heartbeat Profile by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Negative Heartbeat Profile alias. Unique within a cluster", Required: true}, "custom_packet": schema.StringAttribute{MarkdownDescription: "Base64-encoded custom pcap packet. Custom negative heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6", Optional: true, Computed: true}, "custom_packet_file_name": schema.StringAttribute{MarkdownDescription: "File name of referenced 'custom negative heartbeat packet' entry. Used for 'inline negative heartbeat profile' creation", Optional: true, Computed: true}, "direction": schema.StringAttribute{Optional: true, Computed: true}, "period": schema.Int64Attribute{MarkdownDescription: "number of milliseconds between sending subsequent heartbeat packets", Optional: true, Computed: true}, "recovery_time": schema.Int64Attribute{MarkdownDescription: " the minimum number of seconds with successfully received packet to declare that the inline tool is up", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NegativeHbProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NegativeHbProfileResourceModel
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
func (r *NegativeHbProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NegativeHbProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *NegativeHbProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NegativeHbProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NegativeHbProfileResourceModel
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
func (r *NegativeHbProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NegativeHbProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *NegativeHbProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
