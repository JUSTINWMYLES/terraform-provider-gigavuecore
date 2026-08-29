package provider

import "context"
import (
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*PortConfigResource)(nil)

// PortConfigResource is the generated Terraform managed resource implementation.
type PortConfigResource struct {
}

// PortConfigResourceModel describes the Terraform state and plan shape for PortConfigResource.
type PortConfigResourceModel struct {
	AccessRoles       types.Object   `tfsdk:"access_roles" json:"accessRoles"`
	AlarmThresholds   types.Object   `tfsdk:"alarm_thresholds" json:"alarmThresholds"`
	ClusterId         types.String   `tfsdk:"cluster_id" json:"clusterId"`
	Fec               types.String   `tfsdk:"fec"`
	Gdp               types.Bool     `tfsdk:"gdp"`
	HeaderStrip       types.String   `tfsdk:"header_strip" json:"headerStrip"`
	IngressVlanTag    types.Int64    `tfsdk:"ingress_vlan_tag" json:"ingressVlanTag"`
	L2GreId           types.Int64    `tfsdk:"l2_gre_id" json:"l2greId"`
	Licensed          types.Bool     `tfsdk:"licensed"`
	Lock              types.Object   `tfsdk:"lock"`
	MplsAdvanced      types.Object   `tfsdk:"mpls_advanced" json:"mplsAdvanced"`
	NeighborDiscovery types.String   `tfsdk:"neighbor_discovery" json:"neighborDiscovery"`
	PortId            types.String   `tfsdk:"port_id" json:"portId"`
	Ptp               types.Object   `tfsdk:"ptp"`
	Share             types.Object   `tfsdk:"share"`
	TagProtocolId     types.String   `tfsdk:"tag_protocol_id" json:"tagProtocolId"`
	Taptx             types.String   `tfsdk:"taptx"`
	Timestamp         types.Object   `tfsdk:"timestamp"`
	VxlanId           types.Int64    `tfsdk:"vxlan_id" json:"vxlanId"`
	Timeouts          timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *PortConfigResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_port_config"
}

// Schema returns the Terraform schema for this resource.
func (r *PortConfigResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find PortConfig by portId", Attributes: map[string]schema.Attribute{"access_roles": schema.SingleNestedAttribute{MarkdownDescription: "Port Access RBAC definitions", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"level1": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics", Optional: true, Computed: true, ElementType: types.StringType}, "level2": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair", Optional: true, Computed: true, ElementType: types.StringType}, "level3": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair", Optional: true, Computed: true, ElementType: types.StringType}, "level4": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type", Optional: true, Computed: true, ElementType: types.StringType}}}, "alarm_thresholds": schema.SingleNestedAttribute{MarkdownDescription: "Port Alarm Thresholds definitions", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alarm_buffer_threshold_rx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "alarm_buffer_threshold_tx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "alarm_threshold": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "alarm_threshold_low": schema.Int64Attribute{MarkdownDescription: "in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "fec": schema.StringAttribute{MarkdownDescription: "enable/disable forward error correction", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("cl91", "cl74", "off", "default_val", "cl108", "cl91_kp4")}}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP packets on port", Optional: true, Computed: true}, "header_strip": schema.StringAttribute{MarkdownDescription: "protocol type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "mpls-l3", "vxlan", "mpls", "mpls-advanced")}}, "ingress_vlan_tag": schema.Int64Attribute{MarkdownDescription: "Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4000)}}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the l2greId", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 4294967295)}}, "licensed": schema.BoolAttribute{MarkdownDescription: "for TA series indicates whether a port is licensed. Defaults to 'true'", Optional: true, Computed: true}, "lock": schema.SingleNestedAttribute{MarkdownDescription: "Port Locking definitions", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Optional lock description string", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "locking_user": schema.StringAttribute{MarkdownDescription: "User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "shared_with": schema.ListAttribute{MarkdownDescription: "Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account", Optional: true, Computed: true, ElementType: types.StringType}}}, "mpls_advanced": schema.SingleNestedAttribute{MarkdownDescription: "Select a combination of Mpls-Advanced options", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mpls_advanced_opt": schema.ListNestedAttribute{MarkdownDescription: "List of Mpls-Advanced Options", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"adv_opt": schema.StringAttribute{MarkdownDescription: "Mpls-Advanced Options", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("ip4-wo-options", "ip4-options", "ipv6", "l2-wo-cw-ach", "l2-cw-ach", "all")}}}}}}}, "neighbor_discovery": schema.StringAttribute{MarkdownDescription: "Configures port neighbor discovery options", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "all", "lldp", "cdp")}}, "port_id": schema.StringAttribute{MarkdownDescription: "device port id", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "ptp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP configurations", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP announce messages on an interface.\nThe range for the PTP announcement interval is from -2 to 4 log seconds.\nFor the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4", Optional: true, Computed: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Configures the minimum interval allowed between PTP delay messages when the port is in the source state.\nThe range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second.\nFor the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.", Optional: true, Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "Enable PTP on the port", Optional: true, Computed: true}, "local_priority": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}, "role": schema.StringAttribute{MarkdownDescription: "deprecated: use roleAlias", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("standard", "master", "slave")}}, "role_alias": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("standard", "source", "receiver")}}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP synchronization messages on an interface.\nThe range is from log(-7) to log(1) seconds.\nFor domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4. ", Optional: true, Computed: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP timestamp configurations", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"egress": schema.SingleNestedAttribute{MarkdownDescription: "Tx (Egress) settings", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Optional: true, Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}, "ingress": schema.SingleNestedAttribute{MarkdownDescription: "Rx (Ingress) settings", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Optional: true, Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}}}, "vlan": schema.Int64Attribute{MarkdownDescription: "configures vlan on the PTP configured port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(4080, 4089)}}}}, "share": schema.SingleNestedAttribute{MarkdownDescription: "Port Sharing definitions", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tool_share_roles": schema.ListAttribute{MarkdownDescription: "Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles", Optional: true, Computed: true, ElementType: types.StringType}}}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port ", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("0x8100", "0x9100", "0x88a8")}}, "taptx": schema.StringAttribute{MarkdownDescription: "Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("active", "passive")}}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Timestamping definitions for GigaPORT-X12-TS ports", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"append_ingress": schema.BoolAttribute{MarkdownDescription: "Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports", Optional: true, Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> * 2048) + (<slot-id> * 256) + <port-number>", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "strip_egress": schema.BoolAttribute{MarkdownDescription: "Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports", Optional: true, Computed: true}}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the vxlanId", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 16777215)}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *PortConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PortConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.PortId.IsNull() || plan.PortId.IsUnknown() {
		plan.PortId = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *PortConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PortConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *PortConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PortConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state PortConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.PortId.IsNull() || plan.PortId.IsUnknown() {
		if !state.PortId.IsNull() && !state.PortId.IsUnknown() {
			plan.PortId = state.PortId
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *PortConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PortConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
