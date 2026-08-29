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
var _ action.Action = (*RedefineGsGroupParamsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineGsGroupParamsAction)(nil)

// RedefineGsGroupParamsAction is the generated Terraform action implementation.
type RedefineGsGroupParamsAction struct {
	client *client.Client
}

// RedefineGsGroupParamsActionModel describes the action configuration shape.
type RedefineGsGroupParamsActionModel struct {
	Alias                 types.String `tfsdk:"alias"`
	AppTcp                types.Object `tfsdk:"app_tcp" json:"appTcp"`
	ClusterId             types.String `tfsdk:"cluster_id"`
	Dedup                 types.Object `tfsdk:"dedup"`
	DiameterPacket        types.Object `tfsdk:"diameter_packet" json:"diameterPacket"`
	DiameterS6ASession    types.Object `tfsdk:"diameter_s6_a_session" json:"diameterS6aSession"`
	DiameterWhitelist     types.Object `tfsdk:"diameter_whitelist" json:"diameterWhitelist"`
	Eflow                 types.Object `tfsdk:"eflow"`
	EngineWatchdogTimer   types.Object `tfsdk:"engine_watchdog_timer" json:"engineWatchdogTimer"`
	Erspan3               types.Object `tfsdk:"erspan3"`
	FlowMask              types.Object `tfsdk:"flow_mask" json:"flowMask"`
	FlowSampling          types.Object `tfsdk:"flow_sampling" json:"flowSampling"`
	GenericSessionTimeout types.Object `tfsdk:"generic_session_timeout" json:"genericSessionTimeout"`
	GpfcpProfiles         types.Object `tfsdk:"gpfcp_profiles" json:"gpfcpProfiles"`
	GsGroupSystem         types.Object `tfsdk:"gs_group_system" json:"gsGroupSystem"`
	GtaProfiles           types.Object `tfsdk:"gta_profiles" json:"gtaProfiles"`
	GtpControlSampling    types.Object `tfsdk:"gtp_control_sampling" json:"gtpControlSampling"`
	GtpFlow               types.Object `tfsdk:"gtp_flow" json:"gtpFlow"`
	GtpGpfcpDelay         types.Object `tfsdk:"gtp_gpfcp_delay" json:"gtpGpfcpDelay"`
	GtpPersistence        types.Object `tfsdk:"gtp_persistence" json:"gtpPersistence"`
	GtpRandomSampling     types.Object `tfsdk:"gtp_random_sampling" json:"gtpRandomSampling"`
	GtpWhitelist          types.Object `tfsdk:"gtp_whitelist" json:"gtpWhitelist"`
	HealthCheck           types.Object `tfsdk:"health_check" json:"healthCheck"`
	HsmGroup              types.Object `tfsdk:"hsm_group" json:"hsmGroup"`
	IpFrag                types.Object `tfsdk:"ip_frag" json:"ipFrag"`
	LoadBalance           types.Object `tfsdk:"load_balance" json:"loadBalance"`
	Netflow               types.Object `tfsdk:"netflow"`
	NodeRole              types.Object `tfsdk:"node_role" json:"nodeRole"`
	PortThrottleSip       types.Object `tfsdk:"port_throttle_sip" json:"portThrottleSip"`
	Resource              types.Object `tfsdk:"resource"`
	RtpPorts              types.Object `tfsdk:"rtp_ports" json:"rtpPorts"`
	SaApf                 types.Object `tfsdk:"sa_apf" json:"saApf"`
	SessionLogging        types.Object `tfsdk:"session_logging" json:"sessionLogging"`
	SffpProfiles          types.Object `tfsdk:"sffp_profiles" json:"sffpProfiles"`
	SipMedia              types.Object `tfsdk:"sip_media" json:"sipMedia"`
	SipPorts              types.Object `tfsdk:"sip_ports" json:"sipPorts"`
	SipSession            types.Object `tfsdk:"sip_session" json:"sipSession"`
	SipTcpIdleTimeout     types.Object `tfsdk:"sip_tcp_idle_timeout" json:"sipTcpIdleTimeout"`
	SipWhitelist          types.Object `tfsdk:"sip_whitelist" json:"sipWhitelist"`
	SslDecrypt            types.Object `tfsdk:"ssl_decrypt" json:"sslDecrypt"`
	XpktMatch             types.Object `tfsdk:"xpkt_match" json:"xpktMatch"`
}

// NewRedefineGsGroupParamsAction returns a new instance of the generated action.
func NewRedefineGsGroupParamsAction() action.Action {
	return &RedefineGsGroupParamsAction{}
}

// Metadata returns the action type name.
func (r *RedefineGsGroupParamsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_gs_group_params"
}

// Schema returns the action schema.
func (r *RedefineGsGroupParamsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine GS Group's Params", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "app_tcp": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup TCP Parameters", Optional: true, Attributes: map[string]schema.Attribute{"application": schema.StringAttribute{MarkdownDescription: "To choose the action on Unknown Application Data", Optional: true}, "load_balance": schema.BoolAttribute{MarkdownDescription: "Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only", Optional: true}, "tcp_control": schema.StringAttribute{MarkdownDescription: "To choose the action on TCP Control messages", Optional: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dedup": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Dedup Parameters", Optional: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Optional: true}, "ip_tclass": schema.StringAttribute{Optional: true}, "ip_tos": schema.StringAttribute{Optional: true}, "tcp_seq": schema.StringAttribute{Optional: true}, "timer": schema.Int64Attribute{MarkdownDescription: "in microseconds", Optional: true}, "vlan": schema.StringAttribute{Optional: true}}}, "diameter_packet": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter Packet Timeout Parameters", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "timeout in seconds to remove OOO or fragmented packet", Optional: true}}}, "diameter_s6_a_session": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter s6a Session Timeout Parameters", Optional: true, Attributes: map[string]schema.Attribute{"limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for Diameter S6A", Optional: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "timeout in seconds used to clean inactive sessions", Optional: true}}}, "diameter_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter Whitelist Parameters", Optional: true, Attributes: map[string]schema.Attribute{"whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced diameter Whitelist", Required: true}}}, "eflow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Eflow Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/Disable elephant flow detection and handling", Optional: true}, "interval": schema.Int64Attribute{MarkdownDescription: "time interval in seconds", Optional: true}, "log_enabled": schema.BoolAttribute{MarkdownDescription: "Enable/Disable logging of elephant flow parameters into gs logs", Optional: true}, "packet_count": schema.Int64Attribute{MarkdownDescription: "Number of packets to be received by a flow", Optional: true}, "packet_ratio": schema.Int64Attribute{MarkdownDescription: "Percentage of packets in a flow vs overall packet count", Optional: true}}}, "engine_watchdog_timer": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup EngineWatchdogTimer Parameters", Optional: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable", Optional: true}}}, "erspan3": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup ERSPAN III Parameters", Optional: true, Attributes: map[string]schema.Attribute{"timestamp_format": schema.StringAttribute{MarkdownDescription: "erspan III tunnelDecap timestamp format", Optional: true}}}, "flow_mask": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Flow Mask Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true}, "length": schema.Int64Attribute{Optional: true}, "offset": schema.Int64Attribute{Optional: true}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Flow Sampling Parameters", Optional: true, Attributes: map[string]schema.Attribute{"ip_ranges": schema.DynamicAttribute{Optional: true}, "rate": schema.Int64Attribute{MarkdownDescription: "in percent", Optional: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in minutes", Optional: true}, "type": schema.StringAttribute{Optional: true}}}, "generic_session_timeout": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Generic Session Timeout Parameters", Optional: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Maximum timeout for session entry", Optional: true}}}, "gpfcp_profiles": schema.SingleNestedAttribute{MarkdownDescription: "Enriched CUPS Gpfcp profile aliases", Optional: true, Attributes: map[string]schema.Attribute{"g_pfcp_profiles": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}, "gs_group_system": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup System Monitoring Parameters", Optional: true, Attributes: map[string]schema.Attribute{"cpu_load_alarm_threshold": schema.Int64Attribute{MarkdownDescription: "CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated", Optional: true}}}, "gta_profiles": schema.SingleNestedAttribute{MarkdownDescription: "3GPP CUPS gta profile aliases", Optional: true, Attributes: map[string]schema.Attribute{"gta_profiles": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}, "gtp_control_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Control Sampling Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.", Optional: true}}}, "gtp_flow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Flow Parameters", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Session Timeout. in units of 10 minutes. Default of 48 is 8 hours", Optional: true}}}, "gtp_gpfcp_delay": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp GPFCP Delay time", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{Optional: true}}}, "gtp_persistence": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "GTP Persistence Status", Optional: true}, "file_age_timeout": schema.Int64Attribute{MarkdownDescription: "GTP Persistence File Age Timeout(mins)", Optional: true}, "interval": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Interval(mins)", Optional: true}, "restart_age_time": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Restart Age Time(mins)", Optional: true}}}, "gtp_random_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Random Sampling Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, sampling of subscriber's sessions happens in random fashion", Optional: true}, "interval": schema.Int64Attribute{MarkdownDescription: "Rotation Interval in multiples of 12 (hrs)", Optional: true}}}, "gtp_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup GTP Whitelist Parameters", Optional: true, Attributes: map[string]schema.Attribute{"multi_whitelists": schema.ListAttribute{MarkdownDescription: "Alias/Aliases of referenced GTP Whitelists.", Optional: true, ElementType: types.StringType}, "whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced GTP Whitelist.Deprecated since H 5.12", Required: true}}}, "health_check": schema.SingleNestedAttribute{MarkdownDescription: "health check configuration", Optional: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Optional: true}, "dst_port": schema.Int64Attribute{Optional: true}, "enabled": schema.BoolAttribute{Optional: true}, "interval": schema.Int64Attribute{Optional: true}, "protocol": schema.StringAttribute{Optional: true}, "rcv_port": schema.Int64Attribute{Optional: true}, "retries": schema.Int64Attribute{Optional: true}, "round_trip_time": schema.Int64Attribute{Optional: true}, "src_port": schema.Int64Attribute{Optional: true}}}, "hsm_group": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Hsm Group Parameters", Optional: true, Attributes: map[string]schema.Attribute{"hsm_group": schema.StringAttribute{MarkdownDescription: "alias of Hsm Group", Optional: true}}}, "ip_frag": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup IP Fragmentation Parameters", Optional: true, Attributes: map[string]schema.Attribute{"forward": schema.BoolAttribute{Optional: true}, "head_session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Load Balancing Parameters", Optional: true, Attributes: map[string]schema.Attribute{"failover": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Failover part of the GsGroup Load Balancing Parameters", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds", Optional: true}, "threshold_lt_bw": schema.Int64Attribute{MarkdownDescription: "Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%", Optional: true}, "threshold_lt_pkt_rate": schema.Int64Attribute{MarkdownDescription: "Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: [500k..5M] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M", Optional: true}}}, "link_weight_type": schema.StringAttribute{MarkdownDescription: "Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links", Optional: true}, "replicate_gtpc": schema.BoolAttribute{MarkdownDescription: "Enables replication of GTP control packets (GTP-c)", Optional: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Netflow Parameters", Optional: true, Attributes: map[string]schema.Attribute{"monitor": schema.StringAttribute{MarkdownDescription: "Alias of referenced Netflow Monitor", Optional: true}}}, "node_role": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mob5_g_limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for Control 5G Node", Optional: true}, "mob_lte_limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for LTE CPN / UPN Node", Optional: true}, "stand_alone_mode": schema.BoolAttribute{MarkdownDescription: "Enables UPN Stand-alone mode", Optional: true}, "type": schema.StringAttribute{Optional: true}}}, "port_throttle_sip": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SIP Port Throttle Parameters", Optional: true, Attributes: map[string]schema.Attribute{"port_throttle": schema.StringAttribute{MarkdownDescription: "Alias of referenced Port Throttle", Optional: true}}}, "resource": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Parameters", Optional: true, Attributes: map[string]schema.Attribute{"buffer_asf_size": schema.Int64Attribute{MarkdownDescription: "Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot", Optional: true}, "cpu": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot", Optional: true}}}, "hsm_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Hsm Ssl Parameters", Optional: true, Attributes: map[string]schema.Attribute{"buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer in MB. 0 to disable", Optional: true}, "packet_buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl packet-buffer per connection", Optional: true}, "session_count": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer session count in million, 0 to disable", Optional: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "Used to configure other GS apps in addition to Inline SSL on a HC1 box", Optional: true, Attributes: map[string]schema.Attribute{"standalone": schema.BoolAttribute{MarkdownDescription: "If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory", Optional: true}}}, "metadata": schema.Int64Attribute{MarkdownDescription: "flows in millions, how many flows to support for metadata. 0 to disable", Optional: true}, "packet_buffer": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot", Optional: true}}}, "session_overload": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Session overload threshold value , Default value is 90 and 0 is disabled.", Optional: true}}}, "tunnel_overload": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Tunnel overload threshold value , Default value is 90 and 0 is disabled.", Optional: true}}}, "xpkt_match": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Cross Packet Match Parameters", Optional: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "num in 100K flows. 0 is disable", Optional: true}}}}}, "rtp_ports": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Rtp Ports Parameters", Optional: true, Attributes: map[string]schema.Attribute{"range": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Rtp Port Range Parameters", Optional: true, Attributes: map[string]schema.Attribute{"port": schema.Int64Attribute{Required: true}, "port_max": schema.Int64Attribute{MarkdownDescription: "If specified should be greater than 'port'", Optional: true}}}}}, "sa_apf": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Session Aware APF Parameters", Optional: true, Attributes: map[string]schema.Attribute{"buffer_size": schema.Int64Attribute{MarkdownDescription: "Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot", Optional: true}}}, "session_logging": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Session Logging Configuration", Optional: true, Attributes: map[string]schema.Attribute{"interface": schema.StringAttribute{MarkdownDescription: "Associated IP Interface", Optional: true}, "log_level": schema.StringAttribute{MarkdownDescription: "Log Level", Optional: true}, "remote_syslog_ip": schema.StringAttribute{MarkdownDescription: "Remote Syslog IP", Optional: true}, "remote_syslog_port": schema.Int64Attribute{MarkdownDescription: "Remote Syslog Port Number", Optional: true}}}, "sffp_profiles": schema.SingleNestedAttribute{MarkdownDescription: "3GPP CUPS sffp profile aliases", Optional: true, Attributes: map[string]schema.Attribute{"sffp_profiles": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}, "sip_media": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Media Parameters", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Sip media timeout value in seconds .Valid values: 30-300.", Optional: true}}}, "sip_ports": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Ports Parameters", Optional: true, Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{MarkdownDescription: "list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports", Optional: true, ElementType: types.Int64Type}}}, "sip_session": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Session Parameters", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Sip session inactivity timer, value in seconds .Valid values: 30-300.", Optional: true}}}, "sip_tcp_idle_timeout": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Tcp Idle Parameters", Optional: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Sip tcp idle timeout value in seconds .Valid values: 20-600.", Optional: true}}}, "sip_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SIP Whitelist Parameters", Optional: true, Attributes: map[string]schema.Attribute{"whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced SIP Whitelist", Required: true}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SSL Decrypt Parameters", Optional: true, Attributes: map[string]schema.Attribute{"decrypt_fail_action": schema.StringAttribute{Optional: true}, "enabled": schema.BoolAttribute{Optional: true}, "hsm_pkcs11": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Ssl Decrypt Hsm Pkcs11 Parameters", Optional: true, Attributes: map[string]schema.Attribute{"debug_level": schema.Int64Attribute{MarkdownDescription: "hsm pkcs11 debug level", Optional: true}, "dynamic_object": schema.BoolAttribute{MarkdownDescription: "hsm pkcs11 dynamic object", Optional: true}, "load_sharing": schema.BoolAttribute{MarkdownDescription: "hsm pkcs11 load sharing", Optional: true}}}, "hsm_timeout": schema.Int64Attribute{MarkdownDescription: "in milliseconds", Optional: true}, "key_cache_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "key_map": schema.StringAttribute{MarkdownDescription: "references one of the pre-defined 'SslDecryptionKeyMap' groups", Required: true}, "non_ssl_traffic": schema.StringAttribute{Optional: true}, "pending_session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "tcp_syn_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "ticket_cache_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}}}, "xpkt_match": schema.SingleNestedAttribute{MarkdownDescription: "cross packet match configuration", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineGsGroupParamsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineGsGroupParamsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineGsGroupParamsAction) invokeRemote(ctx context.Context, config *RedefineGsGroupParamsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/params"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineGsGroupParamsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
