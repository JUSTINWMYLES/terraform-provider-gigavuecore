package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	regexp "regexp"
	"strings"
	"time"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	identityschema "github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*GsopResource)(nil)
	_ resource.ResourceWithIdentity    = (*GsopResource)(nil)
	_ resource.ResourceWithImportState = (*GsopResource)(nil)
	_ resource.ResourceWithConfigure   = (*GsopResource)(nil)
)

// GsopResource is the generated Terraform managed resource implementation.
type GsopResource struct {
	client *client.Client
}

// GsopResourceModel describes the Terraform state and plan shape for GsopResource.
type GsopResourceModel struct {
	Alias              types.String   `tfsdk:"alias"`
	ClusterId          types.String   `tfsdk:"cluster_id" json:"clusterId"`
	GsApps             types.Object   `tfsdk:"gs_apps" json:"gsApps"`
	GsGroup            types.String   `tfsdk:"gs_group" json:"gsGroup"`
	HealthState        types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Timeouts           timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *GsopResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_gsop"
}

// Schema returns the Terraform schema for this resource.
func (r *GsopResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find GSOP by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "gs_apps": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid", Required: true, Attributes: map[string]schema.Attribute{"apf": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "dedup": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "diameter_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "flow_filter": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp")}}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("ip", "gtp", "sip", "diameter")}}}}, "gseries_header_add": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("srcid", "timestamp"))}}}}, "gseries_header_remove": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "gseries_load_balance": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hash": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("checksum", "xor", "crc")}}, "length": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "offset": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(0)}}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "hash": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("checksum", "xor", "crc")}}, "start_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "start_field": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}}, "gseries_pattern_match": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"length": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "offset": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(0)}}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "start_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}}, "gtp_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "header_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"vlan": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}}}, "header_remove": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ah1": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "eth", "vlan", "mpls", "ipv4", "ipv6")}}, "ah2": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. next anchor header.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "eth", "vlan", "mpls", "ipv4", "ipv6", "tcp", "udp", "any")}}, "custom_len": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. length of unknown header.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 1500)}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1023)}}, "fp_dst_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit destination switch id", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}, "fp_src_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit source switch id", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}, "header_count": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. Number of headers to be stripped.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "offset": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("start", "end", "offsetRange")}}, "offset_range_value": schema.Int64Attribute{MarkdownDescription: "only valid and required when offset is 'offsetRange', integer within range of size of header", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1500)}}, "protocol": schema.StringAttribute{MarkdownDescription: "'gre' and 'fabricPath' are only applicable for H-series", Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp", "isl", "mpls", "mplsPlusVlan", "vlan", "vntag", "vxlan", "gre", "fabricPath", "fm6000Ts", "erspan", "generic")}}, "timestamp_format": schema.StringAttribute{MarkdownDescription: "Timestamp format. Only valid and required for 'fm6000Ts'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("gigasmart", "x12Ts", "none")}}, "vlan_header": schema.StringAttribute{MarkdownDescription: "Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "outer")}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 16777215)}}}}, "icap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"icap_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced ICAP Profile", Required: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inline_ssl_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Inline SSL Profile", Required: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"elb_alias": schema.StringAttribute{MarkdownDescription: "elb app alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}, "stateful": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{MarkdownDescription: "'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5", Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp", "sapf", "sip", "tunnel", "diameter")}}, "diameter_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("sessionId", "userName", "commandCode")}}, "diameter_key_multi_hash_type": schema.ListNestedAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"avp_codevalue": schema.Int64Attribute{MarkdownDescription: "required when 'key' == 'avpCode'", Optional: true, Computed: true}, "key": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("sessionId", "userName", "avpCode", "commandCode", "endToEnd", "hopByHop", "applicationId")}}}}}, "gtp_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("imsi", "imei", "msisdn")}}, "lb_type": schema.StringAttribute{MarkdownDescription: "'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'", Required: true, Validators: []validator.String{stringvalidator.OneOf("leastBw", "leastPktRate", "leastConn", "leastTotalTraffic", "roundRobin", "wtLeastBw", "wtLeastPktRate", "wtLeastConn", "wtLeastTotalTraffic", "wtRoundRobin", "wtImsi", "wtSupi", "gtpKeyHash", "sipKeyHash", "diameterKeyHash", "diameterKeyMultiHash")}}, "sip_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("callerId")}}}}, "stateless": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"field_location": schema.StringAttribute{MarkdownDescription: "Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("inner", "outer")}}, "hash_fields": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("ipOnly", "ipAndPort", "fiveTuple", "gtpuTeid")}}}}}}, "masking": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"content_type": schema.StringAttribute{MarkdownDescription: "content type that will trigger masking, only valid and required for protocol 'sip'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("message_cpim")}}, "length": schema.Int64Attribute{MarkdownDescription: "max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 9600)}}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 9000)}}, "pattern": schema.StringAttribute{MarkdownDescription: "1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(0x)?[0-9a-fA-F]{1,2}$"), "value must match pattern \"^(0x)?[0-9a-fA-F]{1,2}$\"")}}, "protocol": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "ipv4", "ipv6", "udp", "tcp", "ftp-data", "https", "ssh", "gtp", "gtp-ipv4", "gtp-udp", "gtp-tcp", "sip")}}}}, "metadata_export": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cache": schema.StringAttribute{MarkdownDescription: "metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined", Optional: true, Computed: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "sa_apf": schema.SingleNestedAttribute{MarkdownDescription: "Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "sip_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "slicing": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.StringAttribute{MarkdownDescription: "enhanced-slicing apps alias", Optional: true, Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6", Required: true, Validators: []validator.Int64{int64validator.Between(4, 9000)}}, "protocol": schema.StringAttribute{MarkdownDescription: "required property till H 5.6", Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "ipv4", "ipv6", "udp", "tcp", "ftp-data", "https", "ssh", "gtp", "gtp-ipv4", "gtp-udp", "gtp-tcp")}}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'any' port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "out_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'auto' port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}, "trailer_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{MarkdownDescription: "'crc' is not applicable for G-series", Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("crc", "srcid"))}}}}, "trailer_remove": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "tunnel_decap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"custom": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'custom', in which case it is required.", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1023)}}, "gmip_port": schema.Int64Attribute{MarkdownDescription: "only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "l2_gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre', in which case it is required.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4294967295)}}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"decap_key": schema.StringAttribute{Optional: true, Computed: true}, "listener": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gmip", "erspan", "l2gre", "custom", "vxlan", "tls-pcapng")}}, "vxlan": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'vxlan', in which case it is required.", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "vni": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 16777215)}}}}}}, "tunnel_encap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"gmip_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for GMIP Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1048575)}}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 7)}}, "src_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "l2_gre_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.", Optional: true, Computed: true}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1048575)}}, "key": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 4294967295)}}, "pg_dst": schema.StringAttribute{MarkdownDescription: "port group destination alias, mutually exclusive with 'dstIp'", Optional: true, Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 7)}}, "session_field": schema.StringAttribute{MarkdownDescription: "required with stateful loadBalance when 'appType' is 'tunnel'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("fiveTupleIpv4", "fiveTupleIpv6", "fiveTupleAny", "threeTupleIpv4", "threeTupleIpv6", "threeTupleAny", "ipv4Only", "ipv6Only", "ipAny")}}, "session_pos": schema.StringAttribute{MarkdownDescription: "required if 'sessionField' is specified", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("inner", "outer")}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"exporter": schema.StringAttribute{Optional: true, Computed: true}, "exporter_group": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gmip", "l2gre", "vxlan", "tls-pcapng")}}, "vxlan_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.OneOf(4789, 8472, 48879)}}, "src_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}, "vni": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 16777215)}}}}}}}}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PORT_LOW_UTIL", "PORT_HIGH_UTIL", "PORT_PACKET_DROP", "PORT_PACKET_ERROR", "GIGASTREAM_IMBALANCE")}}}}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *GsopResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *GsopResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GsopResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Create(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *GsopResource) createRemote(ctx context.Context, plan *GsopResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/gsops"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_gsop", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.Alias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_gsop", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *GsopResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GsopResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Read(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), state.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *GsopResource) readRemote(ctx context.Context, state *GsopResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsops/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_gsop", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_gsop", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_gsop", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_gsop", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_gsop", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gsop"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_gsop", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *GsopResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GsopResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state GsopResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Update(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *GsopResource) updateRemote(ctx context.Context, plan *GsopResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/gsops/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_gsop", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_gsop", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *GsopResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GsopResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *GsopResource) deleteRemote(ctx context.Context, state *GsopResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsops/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_gsop", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *GsopResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *GsopResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	gsopImportIDParts := strings.Split(req.ID, "/")
	if len(gsopImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}/{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), gsopImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), gsopImportIDParts[1])...)
}
