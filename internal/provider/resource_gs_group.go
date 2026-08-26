package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*GsGroupResource)(nil)
	_ resource.ResourceWithImportState = (*GsGroupResource)(nil)
	_ resource.ResourceWithConfigure   = (*GsGroupResource)(nil)
)

// GsGroupResource is the generated Terraform managed resource implementation.
type GsGroupResource struct {
	client *client.Client
}

// GsGroupResourceModel describes the Terraform state and plan shape for GsGroupResource.
type GsGroupResourceModel struct {
	Alias              types.String `tfsdk:"alias"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	Hash               types.String `tfsdk:"hash"`
	HealthState        types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Params             types.Object `tfsdk:"params"`
	Ports              types.List   `tfsdk:"ports"`
}

// Metadata returns the resource type name.
func (r *GsGroupResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_gs_group"
}

// Schema returns the Terraform schema for this resource.
func (r *GsGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find GS Group by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "GS Group alias", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "hash": schema.StringAttribute{MarkdownDescription: "gsgroup hashing parameters", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "params": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Parameters", Computed: true, Attributes: map[string]schema.Attribute{"app_tcp": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup TCP Parameters", Computed: true, Attributes: map[string]schema.Attribute{"application": schema.StringAttribute{MarkdownDescription: "To choose the action on Unknown Application Data", Computed: true}, "load_balance": schema.BoolAttribute{MarkdownDescription: "Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only", Computed: true}, "tcp_control": schema.StringAttribute{MarkdownDescription: "To choose the action on TCP Control messages", Computed: true}}}, "dedup": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Dedup Parameters", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Computed: true}, "ip_tclass": schema.StringAttribute{Computed: true}, "ip_tos": schema.StringAttribute{Computed: true}, "tcp_seq": schema.StringAttribute{Computed: true}, "timer": schema.Int64Attribute{MarkdownDescription: "in microseconds", Computed: true}, "vlan": schema.StringAttribute{Computed: true}}}, "diameter_packet": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter Packet Timeout Parameters", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "timeout in seconds to remove OOO or fragmented packet", Computed: true}}}, "diameter_s6_a_session": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter s6a Session Timeout Parameters", Computed: true, Attributes: map[string]schema.Attribute{"limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for Diameter S6A", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "timeout in seconds used to clean inactive sessions", Computed: true}}}, "diameter_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Diameter Whitelist Parameters", Computed: true, Attributes: map[string]schema.Attribute{"whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced diameter Whitelist", Computed: true}}}, "eflow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Eflow Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/Disable elephant flow detection and handling", Computed: true}, "interval": schema.Int64Attribute{MarkdownDescription: "time interval in seconds", Computed: true}, "log_enabled": schema.BoolAttribute{MarkdownDescription: "Enable/Disable logging of elephant flow parameters into gs logs", Computed: true}, "packet_count": schema.Int64Attribute{MarkdownDescription: "Number of packets to be received by a flow", Computed: true}, "packet_ratio": schema.Int64Attribute{MarkdownDescription: "Percentage of packets in a flow vs overall packet count", Computed: true}}}, "engine_watchdog_timer": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup EngineWatchdogTimer Parameters", Computed: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable", Computed: true}}}, "erspan3": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup ERSPAN III Parameters", Computed: true, Attributes: map[string]schema.Attribute{"timestamp_format": schema.StringAttribute{MarkdownDescription: "erspan III tunnelDecap timestamp format", Computed: true}}}, "flow_mask": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Flow Mask Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Computed: true}, "length": schema.Int64Attribute{Computed: true}, "offset": schema.Int64Attribute{Computed: true}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Flow Sampling Parameters", Computed: true, Attributes: map[string]schema.Attribute{"rate": schema.Int64Attribute{MarkdownDescription: "in percent", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in minutes", Computed: true}, "type": schema.StringAttribute{Computed: true}}}, "generic_session_timeout": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Generic Session Timeout Parameters", Computed: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Maximum timeout for session entry", Computed: true}}}, "gpfcp_profiles": schema.SingleNestedAttribute{MarkdownDescription: "Enriched CUPS Gpfcp profile aliases", Computed: true, Attributes: map[string]schema.Attribute{"g_pfcp_profiles": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "gs_group_system": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup System Monitoring Parameters", Computed: true, Attributes: map[string]schema.Attribute{"cpu_load_alarm_threshold": schema.Int64Attribute{MarkdownDescription: "CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated", Computed: true}}}, "gta_profiles": schema.SingleNestedAttribute{MarkdownDescription: "3GPP CUPS gta profile aliases", Computed: true, Attributes: map[string]schema.Attribute{"gta_profiles": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "gtp_control_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Control Sampling Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.", Computed: true}}}, "gtp_flow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Flow Parameters", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Session Timeout. in units of 10 minutes. Default of 48 is 8 hours", Computed: true}}}, "gtp_gpfcp_delay": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp GPFCP Delay time", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{Computed: true}}}, "gtp_persistence": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "GTP Persistence Status", Computed: true}, "file_age_timeout": schema.Int64Attribute{MarkdownDescription: "GTP Persistence File Age Timeout(mins)", Computed: true}, "interval": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Interval(mins)", Computed: true}, "restart_age_time": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Restart Age Time(mins)", Computed: true}}}, "gtp_random_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Random Sampling Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, sampling of subscriber's sessions happens in random fashion", Computed: true}, "interval": schema.Int64Attribute{MarkdownDescription: "Rotation Interval in multiples of 12 (hrs)", Computed: true}}}, "gtp_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup GTP Whitelist Parameters", Computed: true, Attributes: map[string]schema.Attribute{"multi_whitelists": schema.ListAttribute{MarkdownDescription: "Alias/Aliases of referenced GTP Whitelists.", Computed: true, ElementType: types.StringType}, "whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced GTP Whitelist.Deprecated since H 5.12", Computed: true}}}, "health_check": schema.SingleNestedAttribute{MarkdownDescription: "health check configuration", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Computed: true}, "dst_port": schema.Int64Attribute{Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "interval": schema.Int64Attribute{Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "rcv_port": schema.Int64Attribute{Computed: true}, "retries": schema.Int64Attribute{Computed: true}, "round_trip_time": schema.Int64Attribute{Computed: true}, "src_port": schema.Int64Attribute{Computed: true}}}, "hsm_group": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Hsm Group Parameters", Computed: true, Attributes: map[string]schema.Attribute{"hsm_group": schema.StringAttribute{MarkdownDescription: "alias of Hsm Group", Computed: true}}}, "ip_frag": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup IP Fragmentation Parameters", Computed: true, Attributes: map[string]schema.Attribute{"forward": schema.BoolAttribute{Computed: true}, "head_session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Load Balancing Parameters", Computed: true, Attributes: map[string]schema.Attribute{"failover": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Failover part of the GsGroup Load Balancing Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds", Computed: true}, "threshold_lt_bw": schema.Int64Attribute{MarkdownDescription: "Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%", Computed: true}, "threshold_lt_pkt_rate": schema.Int64Attribute{MarkdownDescription: "Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: [500k..5M] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M", Computed: true}}}, "link_weight_type": schema.StringAttribute{MarkdownDescription: "Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links", Computed: true}, "replicate_gtpc": schema.BoolAttribute{MarkdownDescription: "Enables replication of GTP control packets (GTP-c)", Computed: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Netflow Parameters", Computed: true, Attributes: map[string]schema.Attribute{"monitor": schema.StringAttribute{MarkdownDescription: "Alias of referenced Netflow Monitor", Computed: true}}}, "node_role": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"mob5_g_limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for Control 5G Node", Computed: true}, "mob_lte_limit": schema.Int64Attribute{MarkdownDescription: "Number of sessions to allocate for LTE CPN / UPN Node", Computed: true}, "stand_alone_mode": schema.BoolAttribute{MarkdownDescription: "Enables UPN Stand-alone mode", Computed: true}, "type": schema.StringAttribute{Computed: true}}}, "port_throttle_sip": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SIP Port Throttle Parameters", Computed: true, Attributes: map[string]schema.Attribute{"port_throttle": schema.StringAttribute{MarkdownDescription: "Alias of referenced Port Throttle", Computed: true}}}, "resource": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Parameters", Computed: true, Attributes: map[string]schema.Attribute{"buffer_asf_size": schema.Int64Attribute{MarkdownDescription: "Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot", Computed: true}, "cpu": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot", Computed: true}}}, "hsm_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Hsm Ssl Parameters", Computed: true, Attributes: map[string]schema.Attribute{"buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer in MB. 0 to disable", Computed: true}, "packet_buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl packet-buffer per connection", Computed: true}, "session_count": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer session count in million, 0 to disable", Computed: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "Used to configure other GS apps in addition to Inline SSL on a HC1 box", Computed: true, Attributes: map[string]schema.Attribute{"standalone": schema.BoolAttribute{MarkdownDescription: "If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory", Computed: true}}}, "metadata": schema.Int64Attribute{MarkdownDescription: "flows in millions, how many flows to support for metadata. 0 to disable", Computed: true}, "packet_buffer": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot", Computed: true}}}, "session_overload": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Session overload threshold value , Default value is 90 and 0 is disabled.", Computed: true}}}, "tunnel_overload": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Tunnel overload threshold value , Default value is 90 and 0 is disabled.", Computed: true}}}, "xpkt_match": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Cross Packet Match Parameters", Computed: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "num in 100K flows. 0 is disable", Computed: true}}}}}, "rtp_ports": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Rtp Ports Parameters", Computed: true, Attributes: map[string]schema.Attribute{"range": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Rtp Port Range Parameters", Computed: true, Attributes: map[string]schema.Attribute{"port": schema.Int64Attribute{Computed: true}, "port_max": schema.Int64Attribute{MarkdownDescription: "If specified should be greater than 'port'", Computed: true}}}}}, "sa_apf": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Session Aware APF Parameters", Computed: true, Attributes: map[string]schema.Attribute{"buffer_size": schema.Int64Attribute{MarkdownDescription: "Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot", Computed: true}}}, "session_logging": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Session Logging Configuration", Computed: true, Attributes: map[string]schema.Attribute{"interface": schema.StringAttribute{MarkdownDescription: "Associated IP Interface", Computed: true}, "log_level": schema.StringAttribute{MarkdownDescription: "Log Level", Computed: true}, "remote_syslog_ip": schema.StringAttribute{MarkdownDescription: "Remote Syslog IP", Computed: true}, "remote_syslog_port": schema.Int64Attribute{MarkdownDescription: "Remote Syslog Port Number", Computed: true}}}, "sffp_profiles": schema.SingleNestedAttribute{MarkdownDescription: "3GPP CUPS sffp profile aliases", Computed: true, Attributes: map[string]schema.Attribute{"sffp_profiles": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "sip_media": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Media Parameters", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Sip media timeout value in seconds .Valid values: 30-300.", Computed: true}}}, "sip_ports": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Ports Parameters", Computed: true, Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{MarkdownDescription: "list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports", Computed: true, ElementType: types.Int64Type}}}, "sip_session": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Session Parameters", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "Sip session inactivity timer, value in seconds .Valid values: 30-300.", Computed: true}}}, "sip_tcp_idle_timeout": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Sip Tcp Idle Parameters", Computed: true, Attributes: map[string]schema.Attribute{"time": schema.Int64Attribute{MarkdownDescription: "Sip tcp idle timeout value in seconds .Valid values: 20-600.", Computed: true}}}, "sip_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SIP Whitelist Parameters", Computed: true, Attributes: map[string]schema.Attribute{"whitelist": schema.StringAttribute{MarkdownDescription: "Alias of referenced SIP Whitelist", Computed: true}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup SSL Decrypt Parameters", Computed: true, Attributes: map[string]schema.Attribute{"decrypt_fail_action": schema.StringAttribute{Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "hsm_pkcs11": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Ssl Decrypt Hsm Pkcs11 Parameters", Computed: true, Attributes: map[string]schema.Attribute{"debug_level": schema.Int64Attribute{MarkdownDescription: "hsm pkcs11 debug level", Computed: true}, "dynamic_object": schema.BoolAttribute{MarkdownDescription: "hsm pkcs11 dynamic object", Computed: true}, "load_sharing": schema.BoolAttribute{MarkdownDescription: "hsm pkcs11 load sharing", Computed: true}}}, "hsm_timeout": schema.Int64Attribute{MarkdownDescription: "in milliseconds", Computed: true}, "key_cache_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "key_map": schema.StringAttribute{MarkdownDescription: "references one of the pre-defined 'SslDecryptionKeyMap' groups", Computed: true}, "non_ssl_traffic": schema.StringAttribute{Computed: true}, "pending_session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "session_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "tcp_syn_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "ticket_cache_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}, "xpkt_match": schema.SingleNestedAttribute{MarkdownDescription: "cross packet match configuration", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Computed: true}}}}}, "ports": schema.ListAttribute{MarkdownDescription: "list of the member GigaSMART e-ports", Required: true, ElementType: types.StringType}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *GsGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GsGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *GsGroupResource) createRemote(ctx context.Context, plan *GsGroupResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/gsGroups"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			plan.Alias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_gs_group", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *GsGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GsGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *GsGroupResource) readRemote(ctx context.Context, state *GsGroupResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		removed = true
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gsGroup"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_gs_group", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *GsGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GsGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state GsGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	preserveStateIntoPlan(&plan, &state)
	r.updateRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *GsGroupResource) updateRemote(ctx context.Context, plan *GsGroupResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/gsGroups/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gs_group", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *GsGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GsGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *GsGroupResource) deleteRemote(ctx context.Context, state *GsGroupResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_gs_group", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *GsGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}

// ImportState imports an existing remote resource into Terraform state.
func (r *GsGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
