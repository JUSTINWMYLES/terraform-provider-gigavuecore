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
var _ action.Action = (*UpdateFabricMapWithLastUpdateTimestampValidationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateFabricMapWithLastUpdateTimestampValidationAction)(nil)

// UpdateFabricMapWithLastUpdateTimestampValidationAction is the generated Terraform action implementation.
type UpdateFabricMapWithLastUpdateTimestampValidationAction struct {
	client *client.Client
}

// UpdateFabricMapWithLastUpdateTimestampValidationActionModel describes the action configuration shape.
type UpdateFabricMapWithLastUpdateTimestampValidationActionModel struct {
	Alias                     types.String  `tfsdk:"alias"`
	ApRules                   types.Dynamic `tfsdk:"ap_rules" json:"apRules"`
	BodyAlias                 types.String  `tfsdk:"body_alias" json:"alias"`
	ClusterId                 types.String  `tfsdk:"cluster_id" json:"clusterId"`
	Comment                   types.String  `tfsdk:"comment"`
	DstPorts                  types.List    `tfsdk:"dst_ports" json:"dstPorts"`
	EgressGigastream          types.List    `tfsdk:"egress_gigastream" json:"egressGigastream"`
	Enable                    types.Bool    `tfsdk:"enable"`
	EncapTunnel               types.String  `tfsdk:"encap_tunnel" json:"encapTunnel"`
	FlexInline                types.Dynamic `tfsdk:"flex_inline" json:"flexInline"`
	FlexInlineFailover        types.String  `tfsdk:"flex_inline_failover" json:"flexInlineFailover"`
	FlexInlineVlanId          types.Int64   `tfsdk:"flex_inline_vlan_id" json:"flexInlineVlanId"`
	FlowRules                 types.Dynamic `tfsdk:"flow_rules" json:"flowRules"`
	FlowSampleDiameterRules   types.Dynamic `tfsdk:"flow_sample_diameter_rules" json:"flowSampleDiameterRules"`
	FlowSampleOverlapRules    types.Dynamic `tfsdk:"flow_sample_overlap_rules" json:"flowSampleOverlapRules"`
	FlowSampleRules           types.Dynamic `tfsdk:"flow_sample_rules" json:"flowSampleRules"`
	FlowSampleSipRules        types.Dynamic `tfsdk:"flow_sample_sip_rules" json:"flowSampleSipRules"`
	FlowWhitelistOverlapRules types.Dynamic `tfsdk:"flow_whitelist_overlap_rules" json:"flowWhitelistOverlapRules"`
	FlowWhitelistRules        types.Dynamic `tfsdk:"flow_whitelist_rules" json:"flowWhitelistRules"`
	GsRules                   types.Dynamic `tfsdk:"gs_rules" json:"gsRules"`
	Gsop                      types.String  `tfsdk:"gsop"`
	HealthState               types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons        types.Dynamic `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineTrafficPath         types.String  `tfsdk:"inline_traffic_path" json:"inlineTrafficPath"`
	InlineTrafficType         types.String  `tfsdk:"inline_traffic_type" json:"inlineTrafficType"`
	ModTime                   types.Int64   `tfsdk:"mod_time" json:"modTime"`
	Order                     types.Int64   `tfsdk:"order"`
	Roles                     types.Dynamic `tfsdk:"roles"`
	RuleMatching              types.String  `tfsdk:"rule_matching" json:"ruleMatching"`
	Rules                     types.Dynamic `tfsdk:"rules"`
	RxClusterPorts            types.List    `tfsdk:"rx_cluster_ports" json:"rxClusterPorts"`
	SrcPorts                  types.List    `tfsdk:"src_ports" json:"srcPorts"`
	SubType                   types.String  `tfsdk:"sub_type" json:"subType"`
	TrafficType               types.String  `tfsdk:"traffic_type" json:"trafficType"`
	TxClusterPorts            types.List    `tfsdk:"tx_cluster_ports" json:"txClusterPorts"`
	Type                      types.String  `tfsdk:"type"`
	UpdatedTime               types.Int64   `tfsdk:"updated_time" json:"updatedTime"`
}

// NewUpdateFabricMapWithLastUpdateTimestampValidationAction returns a new instance of the generated action.
func NewUpdateFabricMapWithLastUpdateTimestampValidationAction() action.Action {
	return &UpdateFabricMapWithLastUpdateTimestampValidationAction{}
}

// Metadata returns the action type name.
func (r *UpdateFabricMapWithLastUpdateTimestampValidationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_fabric_map_with_last_update_timestamp_validation"
}

// Schema returns the action schema.
func (r *UpdateFabricMapWithLastUpdateTimestampValidationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "since FM 6.8.00", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the fabric map", Required: true}, "ap_rules": schema.DynamicAttribute{MarkdownDescription: "pass and drop application profile rules", Optional: true}, "body_alias": schema.StringAttribute{MarkdownDescription: "unique map alias", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "comment": schema.StringAttribute{Optional: true}, "dst_ports": schema.ListAttribute{MarkdownDescription: "List of the 'to' ports", Optional: true, ElementType: types.StringType}, "egress_gigastream": schema.ListAttribute{MarkdownDescription: "Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias", Optional: true, ElementType: types.StringType}, "enable": schema.BoolAttribute{MarkdownDescription: "enable/disable map, applicable only to first level maps", Optional: true}, "encap_tunnel": schema.StringAttribute{MarkdownDescription: "tunnel alias", Optional: true}, "flex_inline": schema.DynamicAttribute{MarkdownDescription: "When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured", Optional: true}, "flex_inline_failover": schema.StringAttribute{MarkdownDescription: "only valid for flexInline maps", Optional: true}, "flex_inline_vlan_id": schema.Int64Attribute{MarkdownDescription: "VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport", Optional: true}, "flow_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Rules Container. Private class", Optional: true}, "flow_sample_diameter_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Sample Diameter Rules Container. Private class", Optional: true}, "flow_sample_overlap_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Sample Overlap Rules Container. Private class", Optional: true}, "flow_sample_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Sample Rules Container. Private class", Optional: true}, "flow_sample_sip_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Sample Sip Rules Container. Private class", Optional: true}, "flow_whitelist_overlap_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Whitelist Overlap Rules Container. Private class", Optional: true}, "flow_whitelist_rules": schema.DynamicAttribute{MarkdownDescription: "Map Flow Whitelist Rules Container. Private class", Optional: true}, "gs_rules": schema.DynamicAttribute{MarkdownDescription: "Map GigaSMART Rules Container. Private class", Optional: true}, "gsop": schema.StringAttribute{MarkdownDescription: "Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types", Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "health_state_reasons": schema.DynamicAttribute{Optional: true}, "inline_traffic_path": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty", Optional: true}, "inline_traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'", Optional: true}, "mod_time": schema.Int64Attribute{MarkdownDescription: "Last modification time of the map in milliseconds since the epoch", Optional: true}, "order": schema.Int64Attribute{MarkdownDescription: "relative order within per-source port map chain", Optional: true}, "roles": schema.DynamicAttribute{MarkdownDescription: "Map Roles Container. Private class", Optional: true}, "rule_matching": schema.StringAttribute{MarkdownDescription: "If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)", Optional: true}, "rules": schema.DynamicAttribute{MarkdownDescription: "Map Rules Container. Private class", Optional: true}, "rx_cluster_ports": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "src_ports": schema.ListAttribute{MarkdownDescription: "list of the 'from' ports", Optional: true, ElementType: types.StringType}, "sub_type": schema.StringAttribute{MarkdownDescription: "'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap' and 'flowWhitelistOverlap' are applicable to 'secondLevel' maps", Optional: true}, "traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'", Optional: true}, "tx_cluster_ports": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport", Optional: true}, "updated_time": schema.Int64Attribute{MarkdownDescription: "Last Updated Timestamp of the fabric map", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateFabricMapWithLastUpdateTimestampValidationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateFabricMapWithLastUpdateTimestampValidationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateFabricMapWithLastUpdateTimestampValidationAction) invokeRemote(ctx context.Context, config *UpdateFabricMapWithLastUpdateTimestampValidationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/{alias}/preventLostUpdate"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Deployment failed. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fabric_map_with_last_update_timestamp_validation", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateFabricMapWithLastUpdateTimestampValidationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
