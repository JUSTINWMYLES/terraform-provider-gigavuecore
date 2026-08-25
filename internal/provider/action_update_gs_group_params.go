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
var _ action.Action = (*UpdateGsGroupParamsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateGsGroupParamsAction)(nil)

// UpdateGsGroupParamsAction is the generated Terraform action implementation.
type UpdateGsGroupParamsAction struct {
	client *client.Client
}

// UpdateGsGroupParamsActionModel describes the action configuration shape.
type UpdateGsGroupParamsActionModel struct {
	Alias                 types.String  `tfsdk:"alias"`
	AppTcp                types.Dynamic `tfsdk:"app_tcp" json:"appTcp"`
	ClusterId             types.String  `tfsdk:"cluster_id"`
	Dedup                 types.Dynamic `tfsdk:"dedup"`
	DiameterPacket        types.Dynamic `tfsdk:"diameter_packet" json:"diameterPacket"`
	DiameterS6ASession    types.Dynamic `tfsdk:"diameter_s6_a_session" json:"diameterS6aSession"`
	DiameterWhitelist     types.Dynamic `tfsdk:"diameter_whitelist" json:"diameterWhitelist"`
	Eflow                 types.Dynamic `tfsdk:"eflow"`
	EngineWatchdogTimer   types.Dynamic `tfsdk:"engine_watchdog_timer" json:"engineWatchdogTimer"`
	Erspan3               types.Dynamic `tfsdk:"erspan3"`
	FlowMask              types.Dynamic `tfsdk:"flow_mask" json:"flowMask"`
	FlowSampling          types.Dynamic `tfsdk:"flow_sampling" json:"flowSampling"`
	GenericSessionTimeout types.Dynamic `tfsdk:"generic_session_timeout" json:"genericSessionTimeout"`
	GpfcpProfiles         types.Dynamic `tfsdk:"gpfcp_profiles" json:"gpfcpProfiles"`
	GsGroupSystem         types.Dynamic `tfsdk:"gs_group_system" json:"gsGroupSystem"`
	GtaProfiles           types.Dynamic `tfsdk:"gta_profiles" json:"gtaProfiles"`
	GtpControlSampling    types.Dynamic `tfsdk:"gtp_control_sampling" json:"gtpControlSampling"`
	GtpFlow               types.Dynamic `tfsdk:"gtp_flow" json:"gtpFlow"`
	GtpGpfcpDelay         types.Dynamic `tfsdk:"gtp_gpfcp_delay" json:"gtpGpfcpDelay"`
	GtpPersistence        types.Dynamic `tfsdk:"gtp_persistence" json:"gtpPersistence"`
	GtpRandomSampling     types.Dynamic `tfsdk:"gtp_random_sampling" json:"gtpRandomSampling"`
	GtpWhitelist          types.Dynamic `tfsdk:"gtp_whitelist" json:"gtpWhitelist"`
	HealthCheck           types.Dynamic `tfsdk:"health_check" json:"healthCheck"`
	HsmGroup              types.Dynamic `tfsdk:"hsm_group" json:"hsmGroup"`
	IpFrag                types.Dynamic `tfsdk:"ip_frag" json:"ipFrag"`
	LoadBalance           types.Dynamic `tfsdk:"load_balance" json:"loadBalance"`
	Netflow               types.Dynamic `tfsdk:"netflow"`
	NodeRole              types.Dynamic `tfsdk:"node_role" json:"nodeRole"`
	PortThrottleSip       types.Dynamic `tfsdk:"port_throttle_sip" json:"portThrottleSip"`
	Resource              types.Dynamic `tfsdk:"resource"`
	RtpPorts              types.Dynamic `tfsdk:"rtp_ports" json:"rtpPorts"`
	SaApf                 types.Dynamic `tfsdk:"sa_apf" json:"saApf"`
	SessionLogging        types.Dynamic `tfsdk:"session_logging" json:"sessionLogging"`
	SffpProfiles          types.Dynamic `tfsdk:"sffp_profiles" json:"sffpProfiles"`
	SipMedia              types.Dynamic `tfsdk:"sip_media" json:"sipMedia"`
	SipPorts              types.Dynamic `tfsdk:"sip_ports" json:"sipPorts"`
	SipSession            types.Dynamic `tfsdk:"sip_session" json:"sipSession"`
	SipTcpIdleTimeout     types.Dynamic `tfsdk:"sip_tcp_idle_timeout" json:"sipTcpIdleTimeout"`
	SipWhitelist          types.Dynamic `tfsdk:"sip_whitelist" json:"sipWhitelist"`
	SslDecrypt            types.Dynamic `tfsdk:"ssl_decrypt" json:"sslDecrypt"`
	XpktMatch             types.Dynamic `tfsdk:"xpkt_match" json:"xpktMatch"`
}

// NewUpdateGsGroupParamsAction returns a new instance of the generated action.
func NewUpdateGsGroupParamsAction() action.Action {
	return &UpdateGsGroupParamsAction{}
}

// Metadata returns the action type name.
func (r *UpdateGsGroupParamsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_gs_group_params"
}

// Schema returns the action schema.
func (r *UpdateGsGroupParamsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Incremental updates to GS Group's Params", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "app_tcp": schema.DynamicAttribute{MarkdownDescription: "GsGroup TCP Parameters", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dedup": schema.DynamicAttribute{MarkdownDescription: "GsGroup Dedup Parameters", Optional: true}, "diameter_packet": schema.DynamicAttribute{MarkdownDescription: "GsGroup Diameter Packet Timeout Parameters", Optional: true}, "diameter_s6_a_session": schema.DynamicAttribute{MarkdownDescription: "GsGroup Diameter s6a Session Timeout Parameters", Optional: true}, "diameter_whitelist": schema.DynamicAttribute{MarkdownDescription: "GsGroup Diameter Whitelist Parameters", Optional: true}, "eflow": schema.DynamicAttribute{MarkdownDescription: "GsGroup Eflow Parameters", Optional: true}, "engine_watchdog_timer": schema.DynamicAttribute{MarkdownDescription: "GsGroup EngineWatchdogTimer Parameters", Optional: true}, "erspan3": schema.DynamicAttribute{MarkdownDescription: "GsGroup ERSPAN III Parameters", Optional: true}, "flow_mask": schema.DynamicAttribute{MarkdownDescription: "GsGroup Flow Mask Parameters", Optional: true}, "flow_sampling": schema.DynamicAttribute{MarkdownDescription: "GsGroup Flow Sampling Parameters", Optional: true}, "generic_session_timeout": schema.DynamicAttribute{MarkdownDescription: "GsGroup Generic Session Timeout Parameters", Optional: true}, "gpfcp_profiles": schema.DynamicAttribute{MarkdownDescription: "Enriched CUPS Gpfcp profile aliases", Optional: true}, "gs_group_system": schema.DynamicAttribute{MarkdownDescription: "GsGroup System Monitoring Parameters", Optional: true}, "gta_profiles": schema.DynamicAttribute{MarkdownDescription: "3GPP CUPS gta profile aliases", Optional: true}, "gtp_control_sampling": schema.DynamicAttribute{MarkdownDescription: "GsGroup Gtp Control Sampling Parameters", Optional: true}, "gtp_flow": schema.DynamicAttribute{MarkdownDescription: "GsGroup Gtp Flow Parameters", Optional: true}, "gtp_gpfcp_delay": schema.DynamicAttribute{MarkdownDescription: "GsGroup Gtp GPFCP Delay time", Optional: true}, "gtp_persistence": schema.DynamicAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Optional: true}, "gtp_random_sampling": schema.DynamicAttribute{MarkdownDescription: "GsGroup Gtp Random Sampling Parameters", Optional: true}, "gtp_whitelist": schema.DynamicAttribute{MarkdownDescription: "GsGroup GTP Whitelist Parameters", Optional: true}, "health_check": schema.DynamicAttribute{MarkdownDescription: "health check configuration", Optional: true}, "hsm_group": schema.DynamicAttribute{MarkdownDescription: "GsGroup Hsm Group Parameters", Optional: true}, "ip_frag": schema.DynamicAttribute{MarkdownDescription: "GsGroup IP Fragmentation Parameters", Optional: true}, "load_balance": schema.DynamicAttribute{MarkdownDescription: "GsGroup Load Balancing Parameters", Optional: true}, "netflow": schema.DynamicAttribute{MarkdownDescription: "GsGroup Netflow Parameters", Optional: true}, "node_role": schema.DynamicAttribute{Optional: true}, "port_throttle_sip": schema.DynamicAttribute{MarkdownDescription: "GsGroup SIP Port Throttle Parameters", Optional: true}, "resource": schema.DynamicAttribute{MarkdownDescription: "GsGroup Resource Parameters", Optional: true}, "rtp_ports": schema.DynamicAttribute{MarkdownDescription: "GsGroup Rtp Ports Parameters", Optional: true}, "sa_apf": schema.DynamicAttribute{MarkdownDescription: "GsGroup Session Aware APF Parameters", Optional: true}, "session_logging": schema.DynamicAttribute{MarkdownDescription: "GsGroup Session Logging Configuration", Optional: true}, "sffp_profiles": schema.DynamicAttribute{MarkdownDescription: "3GPP CUPS sffp profile aliases", Optional: true}, "sip_media": schema.DynamicAttribute{MarkdownDescription: "GsGroup Sip Media Parameters", Optional: true}, "sip_ports": schema.DynamicAttribute{MarkdownDescription: "GsGroup Sip Ports Parameters", Optional: true}, "sip_session": schema.DynamicAttribute{MarkdownDescription: "GsGroup Sip Session Parameters", Optional: true}, "sip_tcp_idle_timeout": schema.DynamicAttribute{MarkdownDescription: "GsGroup Sip Tcp Idle Parameters", Optional: true}, "sip_whitelist": schema.DynamicAttribute{MarkdownDescription: "GsGroup SIP Whitelist Parameters", Optional: true}, "ssl_decrypt": schema.DynamicAttribute{MarkdownDescription: "GsGroup SSL Decrypt Parameters", Optional: true}, "xpkt_match": schema.DynamicAttribute{MarkdownDescription: "cross packet match configuration", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateGsGroupParamsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateGsGroupParamsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateGsGroupParamsAction) invokeRemote(ctx context.Context, config *UpdateGsGroupParamsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/params"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gs_group_params", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateGsGroupParamsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
