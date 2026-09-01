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
var _ action.Action = (*UpdateGsopAppsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateGsopAppsAction)(nil)

// UpdateGsopAppsAction is the generated Terraform action implementation.
type UpdateGsopAppsAction struct {
	client *client.Client
}

// UpdateGsopAppsActionModel describes the action configuration shape.
type UpdateGsopAppsActionModel struct {
	Alias               types.String `tfsdk:"alias"`
	Apf                 types.Object `tfsdk:"apf"`
	ClusterId           types.String `tfsdk:"cluster_id"`
	Dedup               types.Object `tfsdk:"dedup"`
	DiameterWhitelist   types.Object `tfsdk:"diameter_whitelist" json:"diameterWhitelist"`
	FlowFilter          types.Object `tfsdk:"flow_filter" json:"flowFilter"`
	FlowSampling        types.Object `tfsdk:"flow_sampling" json:"flowSampling"`
	GseriesHeaderAdd    types.Object `tfsdk:"gseries_header_add" json:"gseriesHeaderAdd"`
	GseriesHeaderRemove types.Object `tfsdk:"gseries_header_remove" json:"gseriesHeaderRemove"`
	GseriesLoadBalance  types.Object `tfsdk:"gseries_load_balance" json:"gseriesLoadBalance"`
	GseriesPatternMatch types.Object `tfsdk:"gseries_pattern_match" json:"gseriesPatternMatch"`
	GtpWhitelist        types.Object `tfsdk:"gtp_whitelist" json:"gtpWhitelist"`
	HeaderAdd           types.Object `tfsdk:"header_add" json:"headerAdd"`
	HeaderRemove        types.Object `tfsdk:"header_remove" json:"headerRemove"`
	Icap                types.Object `tfsdk:"icap"`
	InlineSsl           types.Object `tfsdk:"inline_ssl" json:"inlineSsl"`
	LoadBalance         types.Object `tfsdk:"load_balance" json:"loadBalance"`
	Masking             types.Object `tfsdk:"masking"`
	MetadataExport      types.Object `tfsdk:"metadata_export" json:"metadataExport"`
	Netflow             types.Object `tfsdk:"netflow"`
	SaApf               types.Object `tfsdk:"sa_apf" json:"saApf"`
	SipWhitelist        types.Object `tfsdk:"sip_whitelist" json:"sipWhitelist"`
	Slicing             types.Object `tfsdk:"slicing"`
	SslDecrypt          types.Object `tfsdk:"ssl_decrypt" json:"sslDecrypt"`
	TrailerAdd          types.Object `tfsdk:"trailer_add" json:"trailerAdd"`
	TrailerRemove       types.Object `tfsdk:"trailer_remove" json:"trailerRemove"`
	TunnelDecap         types.Object `tfsdk:"tunnel_decap" json:"tunnelDecap"`
	TunnelEncap         types.Object `tfsdk:"tunnel_encap" json:"tunnelEncap"`
}

// NewUpdateGsopAppsAction returns a new instance of the generated action.
func NewUpdateGsopAppsAction() action.Action {
	return &UpdateGsopAppsAction{}
}

// Metadata returns the action type name.
func (r *UpdateGsopAppsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_gsop_apps"
}

// Schema returns the action schema.
func (r *UpdateGsopAppsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update GSOP apps configuration", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GSOP", Required: true}, "apf": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dedup": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "diameter_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "flow_filter": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true}}}, "gseries_header_add": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{Required: true, ElementType: types.StringType}}}, "gseries_header_remove": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "gseries_load_balance": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Load Balancing config for G-seres devices", Optional: true, Attributes: map[string]schema.Attribute{"hash": schema.StringAttribute{Required: true}, "length": schema.Int64Attribute{Required: true}, "offset": schema.Int64Attribute{Required: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Load Balancing config for G-seres devices", Optional: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true}, "hash": schema.StringAttribute{Required: true}, "start_delim": schema.StringAttribute{Required: true}, "start_field": schema.StringAttribute{Required: true}}}}}, "gseries_pattern_match": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Pattern Match config for G-seres devices", Optional: true, Attributes: map[string]schema.Attribute{"length": schema.Int64Attribute{Required: true}, "offset": schema.Int64Attribute{Required: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Pattern Match config for G-seres devices", Optional: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true}, "start_delim": schema.StringAttribute{Required: true}}}}}, "gtp_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "header_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"vlan": schema.Int64Attribute{Required: true}}}, "header_remove": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"ah1": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.", Optional: true}, "ah2": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. next anchor header.", Optional: true}, "custom_len": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. length of unknown header.", Optional: true}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids", Optional: true}, "fp_dst_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit destination switch id", Optional: true}, "fp_src_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit source switch id", Optional: true}, "header_count": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. Number of headers to be stripped.", Optional: true}, "offset": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.", Optional: true}, "offset_range_value": schema.Int64Attribute{MarkdownDescription: "only valid and required when offset is 'offsetRange', integer within range of size of header", Optional: true}, "protocol": schema.StringAttribute{MarkdownDescription: "'gre' and 'fabricPath' are only applicable for H-series", Required: true}, "timestamp_format": schema.StringAttribute{MarkdownDescription: "Timestamp format. Only valid and required for 'fm6000Ts'", Optional: true}, "vlan_header": schema.StringAttribute{MarkdownDescription: "Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'", Optional: true}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids", Optional: true}}}, "icap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Optional: true, Attributes: map[string]schema.Attribute{"icap_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced ICAP Profile", Required: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Optional: true, Attributes: map[string]schema.Attribute{"inline_ssl_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Inline SSL Profile", Required: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"enhanced": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"elb_alias": schema.StringAttribute{MarkdownDescription: "elb app alias", Required: true}}}, "stateful": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{MarkdownDescription: "'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5", Required: true}, "diameter_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'", Optional: true}, "diameter_key_multi_hash_type": schema.ListNestedAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"avp_codevalue": schema.Int64Attribute{MarkdownDescription: "required when 'key' == 'avpCode'", Optional: true}, "key": schema.StringAttribute{Required: true}}}}, "gtp_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise", Optional: true}, "lb_type": schema.StringAttribute{MarkdownDescription: "'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'", Required: true}, "sip_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise", Optional: true}}}, "stateless": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"field_location": schema.StringAttribute{MarkdownDescription: "Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise", Optional: true}, "hash_fields": schema.StringAttribute{Required: true}}}}}, "masking": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"content_type": schema.StringAttribute{MarkdownDescription: "content type that will trigger masking, only valid and required for protocol 'sip'", Optional: true}, "length": schema.Int64Attribute{MarkdownDescription: "max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise", Optional: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise", Optional: true}, "pattern": schema.StringAttribute{MarkdownDescription: "1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise", Optional: true}, "protocol": schema.StringAttribute{Required: true}}}, "metadata_export": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"cache": schema.StringAttribute{MarkdownDescription: "metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined", Optional: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "sa_apf": schema.SingleNestedAttribute{MarkdownDescription: "Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "sip_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "slicing": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"enhanced": schema.StringAttribute{MarkdownDescription: "enhanced-slicing apps alias", Optional: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "required property till H 5.6", Required: true}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'any' port", Optional: true}, "out_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'auto' port", Optional: true}}}, "trailer_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{MarkdownDescription: "'crc' is not applicable for G-series", Required: true, ElementType: types.StringType}}}, "trailer_remove": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true}}}, "tunnel_decap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"custom": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'custom', in which case it is required.", Optional: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID", Optional: true}, "gmip_port": schema.Int64Attribute{MarkdownDescription: "only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel", Optional: true}, "l2_gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre', in which case it is required.", Optional: true}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Attributes: map[string]schema.Attribute{"decap_key": schema.StringAttribute{Optional: true}, "listener": schema.StringAttribute{Optional: true}}}, "type": schema.StringAttribute{Required: true}, "vxlan": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'vxlan', in which case it is required.", Optional: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{Required: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}, "vni": schema.Int64Attribute{Required: true}}}}}, "tunnel_encap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Optional: true, Attributes: map[string]schema.Attribute{"gmip_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for GMIP Tunnel Encapsulate GigaSmaprt App", Optional: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true}, "flow_label": schema.Int64Attribute{Optional: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true}, "src_port": schema.Int64Attribute{Required: true}, "ttl": schema.Int64Attribute{Optional: true}}}, "l2_gre_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App", Optional: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.", Optional: true}, "flow_label": schema.Int64Attribute{Optional: true}, "key": schema.Int64Attribute{Required: true}, "pg_dst": schema.StringAttribute{MarkdownDescription: "port group destination alias, mutually exclusive with 'dstIp'", Optional: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true}, "session_field": schema.StringAttribute{MarkdownDescription: "required with stateful loadBalance when 'appType' is 'tunnel'", Optional: true}, "session_pos": schema.StringAttribute{MarkdownDescription: "required if 'sessionField' is specified", Optional: true}, "ttl": schema.Int64Attribute{Optional: true}}}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Attributes: map[string]schema.Attribute{"exporter": schema.StringAttribute{Optional: true}, "exporter_group": schema.StringAttribute{Optional: true}}}, "type": schema.StringAttribute{Required: true}, "vxlan_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App", Optional: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true}, "src_port": schema.Int64Attribute{Required: true}, "ttl": schema.Int64Attribute{Optional: true}, "vni": schema.Int64Attribute{Required: true}}}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateGsopAppsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateGsopAppsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateGsopAppsAction) invokeRemote(ctx context.Context, config *UpdateGsopAppsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsops/{alias}/apps"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateGsopAppsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
