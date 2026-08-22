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
	_ resource.Resource                = (*HbProfileResource)(nil)
	_ resource.ResourceWithImportState = (*HbProfileResource)(nil)
)

// HbProfileResource is the generated Terraform managed resource implementation.
type HbProfileResource struct {
}

// HbProfileResourceModel describes the Terraform state and plan shape for HbProfileResource.
type HbProfileResourceModel struct {
	Alias                types.String `tfsdk:"alias"`
	CustomPacket         types.String `tfsdk:"custom_packet" json:"customPacket"`
	CustomPacketAlias    types.String `tfsdk:"custom_packet_alias" json:"customPacketAlias"`
	CustomPacketFileName types.String `tfsdk:"custom_packet_file_name" json:"customPacketFileName"`
	Direction            types.String `tfsdk:"direction"`
	PacketFormat         types.String `tfsdk:"packet_format" json:"packetFormat"`
	Period               types.Int64  `tfsdk:"period"`
	RecoveryTime         types.Int64  `tfsdk:"recovery_time" json:"recoveryTime"`
	Retries              types.Int64  `tfsdk:"retries"`
	Timeout              types.Int64  `tfsdk:"timeout"`
}

// Metadata returns the resource type name.
func (r *HbProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_hb_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *HbProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Heartbeat Profile by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Heartbeat Profile alias. Unique within a cluster", Required: true}, "custom_packet": schema.StringAttribute{MarkdownDescription: "Base64-encoded custom pcap packet. If omitted, a standard ICMP ARP packet will be used as a heartbeat packet. Custom heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6", Optional: true, Computed: true}, "custom_packet_alias": schema.StringAttribute{MarkdownDescription: "(deprecated) Alias of referenced 'custom heartbeat packet' entry. Used for 'inline heartbeat profile' creation", Optional: true, Computed: true}, "custom_packet_file_name": schema.StringAttribute{MarkdownDescription: "File name of referenced 'custom heartbeat packet' entry", Optional: true, Computed: true}, "direction": schema.StringAttribute{Optional: true, Computed: true}, "packet_format": schema.StringAttribute{Optional: true, Computed: true}, "period": schema.Int64Attribute{MarkdownDescription: "number of milliseconds between sending subsequent heartbeat packets", Optional: true, Computed: true}, "recovery_time": schema.Int64Attribute{MarkdownDescription: " the minimum number of seconds with successfully received packet to declare that the inline tool is up", Optional: true, Computed: true}, "retries": schema.Int64Attribute{MarkdownDescription: "number of consecutive timed-out heartbeat packets at which the system will trigger a failover condition", Optional: true, Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "number of milliseconds allowed for a heartbeat packet between sending and receiving", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *HbProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HbProfileResourceModel
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
func (r *HbProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HbProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *HbProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HbProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state HbProfileResourceModel
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
func (r *HbProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HbProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *HbProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
