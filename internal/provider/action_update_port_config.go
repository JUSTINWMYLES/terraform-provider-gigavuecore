package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdatePortConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdatePortConfigAction)(nil)

// UpdatePortConfigAction is the generated Terraform action implementation.
type UpdatePortConfigAction struct {
	client *client.Client
}

// UpdatePortConfigActionModel describes the action configuration shape.
type UpdatePortConfigActionModel struct {
	AccessRoles       types.Object `tfsdk:"access_roles" json:"accessRoles"`
	AlarmThresholds   types.Object `tfsdk:"alarm_thresholds" json:"alarmThresholds"`
	BodyPortId        types.String `tfsdk:"body_port_id" json:"portId"`
	ClusterId         types.String `tfsdk:"cluster_id"`
	Fec               types.String `tfsdk:"fec"`
	Gdp               types.Bool   `tfsdk:"gdp"`
	HeaderStrip       types.String `tfsdk:"header_strip" json:"headerStrip"`
	IngressVlanTag    types.Int64  `tfsdk:"ingress_vlan_tag" json:"ingressVlanTag"`
	L2GreId           types.Int64  `tfsdk:"l2_gre_id" json:"l2greId"`
	Licensed          types.Bool   `tfsdk:"licensed"`
	Lock              types.Object `tfsdk:"lock"`
	MplsAdvanced      types.Object `tfsdk:"mpls_advanced" json:"mplsAdvanced"`
	NeighborDiscovery types.String `tfsdk:"neighbor_discovery" json:"neighborDiscovery"`
	PortId            types.String `tfsdk:"port_id"`
	Ptp               types.Object `tfsdk:"ptp"`
	Share             types.Object `tfsdk:"share"`
	TagProtocolId     types.String `tfsdk:"tag_protocol_id" json:"tagProtocolId"`
	Taptx             types.String `tfsdk:"taptx"`
	Timestamp         types.Object `tfsdk:"timestamp"`
	VxlanId           types.Int64  `tfsdk:"vxlan_id" json:"vxlanId"`
}

// NewUpdatePortConfigAction returns a new instance of the generated action.
func NewUpdatePortConfigAction() action.Action {
	return &UpdatePortConfigAction{}
}

// Metadata returns the action type name.
func (r *UpdatePortConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_port_config"
}

// Schema returns the action schema.
func (r *UpdatePortConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update a PortConfig configuration", Attributes: map[string]schema.Attribute{"access_roles": schema.SingleNestedAttribute{MarkdownDescription: "Port Access RBAC definitions", Optional: true, Attributes: map[string]schema.Attribute{"level1": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics", Optional: true, ElementType: types.StringType}, "level2": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair", Optional: true, ElementType: types.StringType}, "level3": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair", Optional: true, ElementType: types.StringType}, "level4": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type", Optional: true, ElementType: types.StringType}}}, "alarm_thresholds": schema.SingleNestedAttribute{MarkdownDescription: "Port Alarm Thresholds definitions", Optional: true, Attributes: map[string]schema.Attribute{"alarm_buffer_threshold_rx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Optional: true}, "alarm_buffer_threshold_tx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Optional: true}, "alarm_threshold": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold", Optional: true}, "alarm_threshold_low": schema.Int64Attribute{MarkdownDescription: "in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold", Optional: true}}}, "body_port_id": schema.StringAttribute{MarkdownDescription: "device port id", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "fec": schema.StringAttribute{MarkdownDescription: "enable/disable forward error correction", Optional: true}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP packets on port", Optional: true}, "header_strip": schema.StringAttribute{MarkdownDescription: "protocol type", Optional: true}, "ingress_vlan_tag": schema.Int64Attribute{MarkdownDescription: "Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port", Optional: true}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the l2greId", Optional: true}, "licensed": schema.BoolAttribute{MarkdownDescription: "for TA series indicates whether a port is licensed. Defaults to 'true'", Optional: true}, "lock": schema.SingleNestedAttribute{MarkdownDescription: "Port Locking definitions", Optional: true, Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Optional lock description string", Optional: true}, "locking_user": schema.StringAttribute{MarkdownDescription: "User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name", Required: true}, "shared_with": schema.ListAttribute{MarkdownDescription: "Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account", Optional: true, ElementType: types.StringType}}}, "mpls_advanced": schema.SingleNestedAttribute{MarkdownDescription: "Select a combination of Mpls-Advanced options", Optional: true, Attributes: map[string]schema.Attribute{"mpls_advanced_opt": schema.ListNestedAttribute{MarkdownDescription: "List of Mpls-Advanced Options", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"adv_opt": schema.StringAttribute{MarkdownDescription: "Mpls-Advanced Options", Optional: true}}}}}}, "neighbor_discovery": schema.StringAttribute{MarkdownDescription: "Configures port neighbor discovery options", Optional: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Device Port ID (format: boxId_slotId_port, example: 1_1_c1)", Required: true}, "ptp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP configurations", Optional: true, Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP announce messages on an interface.\nThe range for the PTP announcement interval is from -2 to 4 log seconds.\nFor the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4", Optional: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Configures the minimum interval allowed between PTP delay messages when the port is in the source state.\nThe range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second.\nFor the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.", Optional: true}, "enable": schema.BoolAttribute{MarkdownDescription: "Enable PTP on the port", Optional: true}, "local_priority": schema.Int64Attribute{Optional: true}, "role": schema.StringAttribute{MarkdownDescription: "deprecated: use roleAlias", Optional: true}, "role_alias": schema.StringAttribute{Optional: true}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP synchronization messages on an interface.\nThe range is from log(-7) to log(1) seconds.\nFor domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4. ", Optional: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP timestamp configurations", Optional: true, Attributes: map[string]schema.Attribute{"egress": schema.SingleNestedAttribute{MarkdownDescription: "Tx (Egress) settings", Optional: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Optional: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Optional: true}}}, "ingress": schema.SingleNestedAttribute{MarkdownDescription: "Rx (Ingress) settings", Optional: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Optional: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Optional: true}}}}}, "vlan": schema.Int64Attribute{MarkdownDescription: "configures vlan on the PTP configured port", Optional: true}}}, "share": schema.SingleNestedAttribute{MarkdownDescription: "Port Sharing definitions", Optional: true, Attributes: map[string]schema.Attribute{"tool_share_roles": schema.ListAttribute{MarkdownDescription: "Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles", Optional: true, ElementType: types.StringType}}}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port ", Optional: true}, "taptx": schema.StringAttribute{MarkdownDescription: "Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay", Optional: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Timestamping definitions for GigaPORT-X12-TS ports", Optional: true, Attributes: map[string]schema.Attribute{"append_ingress": schema.BoolAttribute{MarkdownDescription: "Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports", Optional: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> * 2048) + (<slot-id> * 256) + <port-number>", Optional: true}, "strip_egress": schema.BoolAttribute{MarkdownDescription: "Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports", Optional: true}}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the vxlanId", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdatePortConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdatePortConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdatePortConfigAction) invokeRemote(ctx context.Context, config *UpdatePortConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/portConfig/portConfigs/{portId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(config.PortId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_port_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdatePortConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
