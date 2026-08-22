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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
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
	AccessRoles       types.Dynamic `tfsdk:"access_roles" json:"accessRoles"`
	AlarmThresholds   types.Dynamic `tfsdk:"alarm_thresholds" json:"alarmThresholds"`
	BodyPortId        types.String  `tfsdk:"body_port_id" json:"portId"`
	ClusterId         types.String  `tfsdk:"cluster_id"`
	Fec               types.String  `tfsdk:"fec"`
	Gdp               types.Bool    `tfsdk:"gdp"`
	HeaderStrip       types.String  `tfsdk:"header_strip" json:"headerStrip"`
	IngressVlanTag    types.Int64   `tfsdk:"ingress_vlan_tag" json:"ingressVlanTag"`
	L2GreId           types.Int64   `tfsdk:"l2_gre_id" json:"l2greId"`
	Licensed          types.Bool    `tfsdk:"licensed"`
	Lock              types.Dynamic `tfsdk:"lock"`
	MplsAdvanced      types.Dynamic `tfsdk:"mpls_advanced" json:"mplsAdvanced"`
	NeighborDiscovery types.String  `tfsdk:"neighbor_discovery" json:"neighborDiscovery"`
	PortId            types.String  `tfsdk:"port_id"`
	Ptp               types.Dynamic `tfsdk:"ptp"`
	Share             types.Dynamic `tfsdk:"share"`
	TagProtocolId     types.String  `tfsdk:"tag_protocol_id" json:"tagProtocolId"`
	Taptx             types.String  `tfsdk:"taptx"`
	Timestamp         types.Dynamic `tfsdk:"timestamp"`
	VxlanId           types.Int64   `tfsdk:"vxlan_id" json:"vxlanId"`
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
	resp.Schema = schema.Schema{Description: "Update a PortConfig configuration", Attributes: map[string]schema.Attribute{"access_roles": schema.DynamicAttribute{MarkdownDescription: "Port Access RBAC definitions", Optional: true}, "alarm_thresholds": schema.DynamicAttribute{MarkdownDescription: "Port Alarm Thresholds definitions", Optional: true}, "body_port_id": schema.StringAttribute{MarkdownDescription: "device port id", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "fec": schema.StringAttribute{MarkdownDescription: "enable/disable forward error correction", Optional: true}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP packets on port", Optional: true}, "header_strip": schema.StringAttribute{MarkdownDescription: "protocol type", Optional: true}, "ingress_vlan_tag": schema.Int64Attribute{MarkdownDescription: "Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port", Optional: true}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the l2greId", Optional: true}, "licensed": schema.BoolAttribute{MarkdownDescription: "for TA series indicates whether a port is licensed. Defaults to 'true'", Optional: true}, "lock": schema.DynamicAttribute{MarkdownDescription: "Port Locking definitions", Optional: true}, "mpls_advanced": schema.DynamicAttribute{MarkdownDescription: "Select a combination of Mpls-Advanced options", Optional: true}, "neighbor_discovery": schema.StringAttribute{MarkdownDescription: "Configures port neighbor discovery options", Optional: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Device Port ID (format: boxId_slotId_port, example: 1_1_c1)", Required: true}, "ptp": schema.DynamicAttribute{MarkdownDescription: "Port PTP configurations", Optional: true}, "share": schema.DynamicAttribute{MarkdownDescription: "Port Sharing definitions", Optional: true}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port ", Optional: true}, "taptx": schema.StringAttribute{MarkdownDescription: "Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay", Optional: true}, "timestamp": schema.DynamicAttribute{MarkdownDescription: "Timestamping definitions for GigaPORT-X12-TS ports", Optional: true}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the vxlanId", Optional: true}}}
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
